package portworx_test

import (
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	test_utils "github.com/portworx/helm/test/utils"
	"github.com/stretchr/testify/require"
)

func TestStorageClusterHelmTemplate(t *testing.T) {

	t.Parallel()

	// Path to the helm chart we will test
	helmChartPath, err := filepath.Abs("../../charts/portworx/")
	// name of template that we want to test
	templateFileName := "storage-cluster.yaml"
	require.NoError(t, err)

	testCases := []struct {
		name           string
		helmOption     *helm.Options
		resultFileName string
	}{
		{
			name:           "TestAllComponentsEnabled",
			resultFileName: "storagecluster_all_compenents_enabled.yaml",
			helmOption: &helm.Options{
				ValuesFiles: []string{"./testValues/storagecluster_all_components_enabled.yaml"},
			},
		},
		{
			name:           "TestDefaultChartValues",
			resultFileName: "storagecluster_with_default_values.yaml",
			helmOption: &helm.Options{
				SetValues: map[string]string{"internalKVDB": "true"},
			},
		},
		{
			name:           "TestCustomRegistry",
			resultFileName: "storagecluster_custom_registry.yaml",
			helmOption: &helm.Options{
				ValuesFiles: []string{"./testValues/storagecluster_custom_registry.yaml"},
			},
		},
		{
			name:           "TestExternalETCD",
			resultFileName: "storagecluster_external_etcd.yaml",
			helmOption: &helm.Options{
				ValuesFiles: []string{"./testValues/storagecluster_external_etcd.yaml"},
			},
		},
		{
			name:           "TestPlacementTolerations",
			resultFileName: "storagecluster_placement.yaml",
			helmOption: &helm.Options{
				ValuesFiles: []string{"./testValues/storagecluster_placement.yaml"},
			},
		},
	}

	for _, testCase := range testCases {

		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			test_utils.TestRenderedHelmTemplate(t, testCase.helmOption, helmChartPath, templateFileName, testCase.resultFileName)
		})
	}
}