package scanner

import (
	"bytes"
	"encoding/json"
	"net/http"
	"sort"

	"github.com/docker/distribution/reference"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/apiparams"
	"github.com/stackrox/rox/pkg/errorhelpers"
	"github.com/stackrox/rox/pkg/httputil"
	"github.com/stackrox/rox/pkg/images/defaults"
	"github.com/stackrox/rox/pkg/mtls"
	"github.com/stackrox/rox/pkg/renderer"
	"github.com/stackrox/rox/pkg/stringutils"
	"github.com/stackrox/rox/pkg/zip"
	"google.golang.org/grpc/codes"
)

func validateImageName(imageName, whichImage string) error {
	if imageName != "" {
		if _, err := reference.ParseAnyReference(imageName); err != nil {
			return errors.Wrapf(err, "invalid %s image", whichImage)
		}
	}
	return nil
}

func validateParamsForScannerV1(p *apiparams.Scanner) (errs []error) {
	if err := validateImageName(p.ScannerImage, "scanner"); err != nil {
		errs = append(errs, err)
	}
	return
}

func validateParamsAndNormalizeClusterType(p *apiparams.Scanner) (storage.ClusterType, error) {
	errorList := errorhelpers.NewErrorList("invalid params:")

	clusterType := storage.ClusterType(storage.ClusterType_value[p.ClusterType])

	if int32(clusterType) == 0 {
		var validClusterTypes []string
		for clusterString, value := range storage.ClusterType_value {
			if value > 0 {
				validClusterTypes = append(validClusterTypes, clusterString)
			}
		}
		sort.Strings(validClusterTypes)
		errorList.AddStringf("invalid cluster type: %q; valid options are %+v", p.ClusterType, validClusterTypes)
	}

	errorList.AddErrors(validateParamsForScannerV1(p)...)

	return clusterType, errorList.ToError()
}

func generateFilesForScannerV1(params *apiparams.Scanner, clusterType storage.ClusterType) ([]*zip.File, error) {
	centralCA, err := mtls.CACertPEM()
	if err != nil {
		return nil, errors.Wrap(err, "could not load central CA")
	}

	cert, err := mtls.IssueNewCert(mtls.ScannerSubject)
	if err != nil {
		return nil, errors.Wrap(err, "could not issue scanner cert")
	}

	scannerDBCert, err := mtls.IssueNewCert(mtls.ScannerDBSubject)
	if err != nil {
		return nil, errors.Wrap(err, "could not issue scanner db cert")
	}
	dbPassword := []byte(renderer.CreatePassword())
	v4DBPassword := []byte(renderer.CreatePassword())

	scannerV4IndexerCert, err := mtls.IssueNewCert(mtls.ScannerV4IndexerSubject)
	if err != nil {
		return nil, errors.Wrap(err, "could not issue certifiate for Scanner V4 Indexer")
	}
	scannerV4MatcherCert, err := mtls.IssueNewCert(mtls.ScannerV4MatcherSubject)
	if err != nil {
		return nil, errors.Wrap(err, "could not issue certifiate for Scanner V4 Matcher")
	}
	scannerV4DBCert, err := mtls.IssueNewCert(mtls.ScannerV4DBSubject)
	if err != nil {
		return nil, errors.Wrap(err, "could not issue certifiate for Scanner V4 DB")
	}

	flavor := defaults.GetImageFlavorFromEnv()
	config := renderer.Config{
		ClusterType: clusterType,
		K8sConfig: &renderer.K8sConfig{
			CommonConfig: renderer.CommonConfig{
				ScannerImage:     stringutils.OrDefault(params.ScannerImage, flavor.ScannerImage()),
				ScannerDBImage:   stringutils.OrDefault(params.ScannerDBImage, flavor.ScannerDBImage()),
				ScannerV4Image:   stringutils.OrDefault(params.ScannerV4Image, flavor.ScannerV4Image()),
				ScannerV4DBImage: stringutils.OrDefault(params.ScannerV4DBImage, flavor.ScannerV4DBImage()),
			},
			OfflineMode:  params.OfflineMode,
			IstioVersion: params.IstioVersion,
		},
		SecretsByteMap: map[string][]byte{
			"ca.pem":           centralCA,
			"scanner-cert.pem": cert.CertPEM,
			"scanner-key.pem":  cert.KeyPEM,

			"scanner-db-cert.pem": scannerDBCert.CertPEM,
			"scanner-db-key.pem":  scannerDBCert.KeyPEM,
			"scanner-db-password": dbPassword,

			"scanner-v4-db-cert.pem": scannerV4DBCert.CertPEM,
			"scanner-v4-db-key.pem":  scannerV4DBCert.KeyPEM,
			"scanner-v4-db-password": v4DBPassword,

			"scanner-v4-indexer-cert.pem": scannerV4IndexerCert.CertPEM,
			"scanner-v4-indexer-key.pem":  scannerV4IndexerCert.KeyPEM,

			"scanner-v4-matcher-cert.pem": scannerV4MatcherCert.CertPEM,
			"scanner-v4-matcher-key.pem":  scannerV4MatcherCert.KeyPEM,
		},
		EnablePodSecurityPolicies: !params.DisablePodSecurityPolicies,
	}

	return renderer.RenderScannerOnly(config, flavor)
}

func serveHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var params apiparams.Scanner
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		httputil.WriteGRPCStyleError(w, codes.InvalidArgument, err)
		return
	}
	err = json.Unmarshal(buf.Bytes(), &params)
	if err != nil {
		httputil.WriteGRPCStyleError(w, codes.Internal, err)
		return
	}

	clusterType, err := validateParamsAndNormalizeClusterType(&params)
	if err != nil {
		httputil.WriteGRPCStyleError(w, codes.InvalidArgument, err)
		return
	}

	files, err := generateFilesForScannerV1(&params, clusterType)
	if err != nil {
		httputil.WriteGRPCStyleError(w, codes.Internal, err)
		return
	}

	wrapper := zip.NewWrapper()
	wrapper.AddFiles(files...)
	bytes, err := wrapper.Zip()
	if err != nil {
		httputil.WriteGRPCStyleError(w, codes.Internal, err)
	}

	// Tell the browser this is a download.
	w.Header().Add("Content-Disposition", `attachment; filename="scanner-bundle.zip"`)
	_, _ = w.Write(bytes)
}

// Handler returns the handler that serves scanner zip files.
func Handler() http.Handler {
	return http.HandlerFunc(serveHTTP)
}
