//go:build integration

package google

import (
	"os"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const project = "acs-san-stackroxci"

func TestGoogle(t *testing.T) {
	if os.Getenv("SERVICE_ACCOUNT") == "" {
		t.Skip("SERVICE_ACCOUNT env variable required")
		return
	}
	t.Setenv("ROX_REGISTRY_RESPONSE_TIMEOUT", "90s")
	t.Setenv("ROX_REGISTRY_CLIENT_TIMEOUT", "120s")

	integration := &storage.ImageIntegration{
		IntegrationConfig: &storage.ImageIntegration_Google{
			Google: &storage.GoogleConfig{
				Endpoint:       "us.gcr.io",
				ServiceAccount: os.Getenv("SERVICE_ACCOUNT"),
			},
		},
	}

	registry, err := NewRegistry(integration, false, nil)
	require.NoError(t, err)

	metadata, err := registry.Metadata(&storage.Image{
		Name: &storage.ImageName{
			Registry: "us.gcr.io",
			Remote:   project + "/music-nginx",
			Tag:      "latest",
		},
	})
	require.NoError(t, err)
	assert.Len(t, metadata.LayerShas, 10)
}
