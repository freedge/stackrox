import * as yup from 'yup';
import React, { ReactElement, useEffect, useState } from 'react';
import {
    Alert,
    AlertVariant,
    Button,
    Flex,
    FlexItem,
    Form,
    FormGroup,
    FormSection,
    PageSection,
    Popover,
    SelectOption,
    TextInput,
} from '@patternfly/react-core';
import { FieldArray, FormikProvider } from 'formik';
import { ArrowRightIcon, HelpIcon, PlusCircleIcon, TrashIcon } from '@patternfly/react-icons';
import { IntegrationFormProps } from 'Containers/Integrations/IntegrationForm/integrationFormTypes';
import useIntegrationForm from 'Containers/Integrations/IntegrationForm/useIntegrationForm';
import FormMessage from 'Components/PatternFly/FormMessage';
import FormLabelGroup from 'Containers/Integrations/IntegrationForm/FormLabelGroup';
import IntegrationFormActions from 'Containers/Integrations/IntegrationForm/IntegrationFormActions';
import FormSaveButton from 'Components/PatternFly/FormSaveButton';
import FormCancelButton from 'Components/PatternFly/FormCancelButton';
import SelectSingle from 'Components/SelectSingle';
import { fetchRolesAsArray, Role } from 'services/RolesService';
import { MachineConfigType } from 'services/MachineAccessService';
import { getAxiosErrorMessage } from '../../../../utils/responseErrorUtils';

export type MachineAccessConfig = {
    id: string;
    type: MachineConfigType;
    tokenExpirationDuration: string;
    mappings: {
        key: string;
        valueExpression: string;
        role: string;
    }[];
    issuer: string;
};

export const validationSchema = yup.object().shape({
    type: yup.string().trim().required('Type is required.'),
    tokenExpirationDuration: yup.string().trim().required('Token lifetime is required'),
    mappings: yup
        .array()
        .of(
            yup.object().shape({
                key: yup.string().trim().required('Key is required.'),
                valueExpression: yup.string().trim().required('Value expression is required.'),
                role: yup.string().trim().required('Role is required.'),
            })
        )
        .min(1),
    issuer: yup.string().trim().required('Issuer is required.'),
});

export const defaultValues: MachineAccessConfig = {
    issuer: '',
    mappings: [],
    tokenExpirationDuration: '',
    type: 'GENERIC',
    id: '',
};

function MachineAccessIntegrationForm({
    initialValues = null,
    isEditable = false,
}: IntegrationFormProps<MachineAccessConfig>): ReactElement {
    const formInitialValues = { ...defaultValues, ...initialValues };
    const formik = useIntegrationForm<MachineAccessConfig>({
        initialValues: formInitialValues,
        validationSchema,
    });
    const {
        values,
        touched,
        errors,
        dirty,
        isValid,
        setFieldValue,
        handleBlur,
        isSubmitting,
        isTesting,
        onSave,
        onCancel,
        message,
    } = formik;

    function onChange(value, event) {
        return setFieldValue(event.target.id, value);
    }

    const [roles, setRoles] = useState<Role[]>([]);
    const [alertRoles, setAlertRoles] = useState<string>('');

    useEffect(() => {
        fetchRolesAsArray()
            .then((rolesFetched) => {
                setRoles(rolesFetched);
                setAlertRoles('');
            })
            .catch((error) => {
                setAlertRoles(getAxiosErrorMessage(error));
            });
    }, []);

    return (
        <>
            {alertRoles && (
                <Alert title="Fetch roles failed" variant={AlertVariant.warning} isInline>
                    {alertRoles}
                </Alert>
            )}
            <PageSection variant="light" isFilled hasOverflowScroll>
                <FormMessage message={message} />
                <Form isWidthLimited>
                    <FormikProvider value={formik}>
                        <FormLabelGroup
                            isRequired
                            label="Select configuration type"
                            fieldId="type"
                            touched={touched}
                            errors={errors}
                        >
                            <SelectSingle
                                id="type"
                                value={values.type}
                                handleSelect={(name, value) => {
                                    setFieldValue(name, value);
                                    setFieldValue(
                                        'issuer',
                                        value === 'GITHUB_ACTIONS'
                                            ? 'https://token.actions.githubusercontent.com'
                                            : ''
                                    );
                                }}
                                direction="down"
                                isDisabled={!isEditable}
                            >
                                <SelectOption key={'GENERIC'} value={'GENERIC'}>
                                    Generic
                                </SelectOption>
                                <SelectOption key={'GITHUB_ACTIONS'} value={'GITHUB_ACTIONS'}>
                                    Github action
                                </SelectOption>
                            </SelectSingle>
                        </FormLabelGroup>
                        <FormLabelGroup
                            isRequired
                            label="Issuer"
                            fieldId="issuer"
                            touched={touched}
                            errors={errors}
                        >
                            <TextInput
                                isRequired
                                type="text"
                                id="issuer"
                                value={values.issuer}
                                onChange={onChange}
                                onBlur={handleBlur}
                                isDisabled={!isEditable || values.type === 'GITHUB_ACTIONS'}
                            />
                        </FormLabelGroup>
                        <FormLabelGroup
                            isRequired
                            label="Token lifetime"
                            fieldId="tokenExpirationDuration"
                            touched={touched}
                            errors={errors}
                            helperText="For example, 3h20m20s"
                        >
                            <TextInput
                                isRequired
                                type="text"
                                id="tokenExpirationDuration"
                                value={values.tokenExpirationDuration}
                                onChange={onChange}
                                onBlur={handleBlur}
                                isDisabled={!isEditable}
                            />
                        </FormLabelGroup>
                        <FormSection title="Rules" titleElement="h3" className="pf-u-mt-0">
                            <FieldArray
                                name="mappings"
                                render={(arrayHelpers) => (
                                    <>
                                        {values.mappings.length === 0 && <p>No rules defined</p>}
                                        {values.mappings.length > 0 &&
                                            values.mappings.map((_mapping, index: number) => (
                                                // eslint-disable-next-line react/no-array-index-key
                                                <Flex key={`mapping_${index}`}>
                                                    <FlexItem>
                                                        <FormLabelGroup
                                                            label="Key"
                                                            fieldId={`mappings[${index}].key`}
                                                            touched={touched}
                                                            errors={errors}
                                                        >
                                                            <TextInput
                                                                isRequired
                                                                type="text"
                                                                id={`mappings[${index}].key`}
                                                                value={
                                                                    values.mappings[`${index}`].key
                                                                }
                                                                onChange={onChange}
                                                                onBlur={handleBlur}
                                                                isDisabled={!isEditable}
                                                            />
                                                        </FormLabelGroup>
                                                    </FlexItem>
                                                    <FlexItem>
                                                        <FormLabelGroup
                                                            label="Value"
                                                            labelIcon={
                                                                <Popover
                                                                    bodyContent={
                                                                        <div>
                                                                            <a
                                                                                href="https://golang.org/s/re2syntax"
                                                                                target="_blank"
                                                                                rel="noreferrer"
                                                                            >
                                                                                Learn how to use
                                                                                regex here
                                                                            </a>
                                                                        </div>
                                                                    }
                                                                    headerContent={
                                                                        'Use regex to enter values'
                                                                    }
                                                                >
                                                                    {
                                                                        <Button
                                                                            type="button"
                                                                            aria-label="More info for name field"
                                                                            className="pf-c-form__group-label-help"
                                                                            style={{
                                                                                backgroundColor:
                                                                                    'transparent',
                                                                            }}
                                                                        >
                                                                            <HelpIcon
                                                                                style={{
                                                                                    color: 'black',
                                                                                }}
                                                                                noVerticalAlign
                                                                            />
                                                                        </Button>
                                                                    }
                                                                </Popover>
                                                            }
                                                            fieldId={`mappings[${index}].valueExpression`}
                                                            touched={touched}
                                                            errors={errors}
                                                        >
                                                            <TextInput
                                                                isRequired
                                                                type="text"
                                                                id={`mappings[${index}].valueExpression`}
                                                                value={
                                                                    values.mappings[`${index}`]
                                                                        .valueExpression
                                                                }
                                                                onChange={onChange}
                                                                onBlur={handleBlur}
                                                                isDisabled={!isEditable}
                                                            />
                                                        </FormLabelGroup>
                                                    </FlexItem>
                                                    <FlexItem>
                                                        <ArrowRightIcon
                                                            style={{
                                                                transform: 'translate(0, 42px)',
                                                            }}
                                                        />
                                                    </FlexItem>
                                                    <FlexItem>
                                                        <FormGroup
                                                            label="Role"
                                                            fieldId={`mappings[${index}].role`}
                                                            helperTextInvalid={
                                                                errors[index]?.role || ''
                                                            }
                                                            validated={
                                                                errors[index]?.role
                                                                    ? 'error'
                                                                    : 'default'
                                                            }
                                                        >
                                                            <SelectSingle
                                                                id={`mappings[${index}].role`}
                                                                value={
                                                                    values.mappings[`${index}`].role
                                                                }
                                                                isDisabled={!isEditable}
                                                                handleSelect={setFieldValue}
                                                                direction="up"
                                                                placeholderText="Select a role"
                                                            >
                                                                {roles.map(({ name }) => (
                                                                    <SelectOption
                                                                        key={name}
                                                                        value={name}
                                                                    />
                                                                ))}
                                                            </SelectSingle>
                                                        </FormGroup>
                                                    </FlexItem>
                                                    {isEditable && (
                                                        <FlexItem>
                                                            <Button
                                                                variant="plain"
                                                                aria-label="Delete rule"
                                                                style={{
                                                                    transform: 'translate(0, 42px)',
                                                                }}
                                                                onClick={() =>
                                                                    arrayHelpers.remove(index)
                                                                }
                                                            >
                                                                <TrashIcon />
                                                            </Button>
                                                        </FlexItem>
                                                    )}
                                                </Flex>
                                            ))}
                                        {isEditable && (
                                            <Flex>
                                                <FlexItem>
                                                    <Button
                                                        variant="link"
                                                        isInline
                                                        icon={
                                                            <PlusCircleIcon className="pf-u-mr-sm" />
                                                        }
                                                        onClick={() =>
                                                            arrayHelpers.push({
                                                                key: '',
                                                                valueExpression: '',
                                                                role: '',
                                                            })
                                                        }
                                                    >
                                                        Add new rule
                                                    </Button>
                                                </FlexItem>
                                            </Flex>
                                        )}
                                    </>
                                )}
                            />
                        </FormSection>
                    </FormikProvider>
                </Form>
            </PageSection>
            {isEditable && (
                <IntegrationFormActions>
                    <FormSaveButton
                        onSave={onSave}
                        isSubmitting={isSubmitting}
                        isTesting={isTesting}
                        isDisabled={!dirty || !isValid}
                    >
                        Save
                    </FormSaveButton>
                    <FormCancelButton onCancel={onCancel}>Cancel</FormCancelButton>
                </IntegrationFormActions>
            )}
        </>
    );
}

export default MachineAccessIntegrationForm;
