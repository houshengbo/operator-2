// +build preupgrade postdowngrade

/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"os"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/operator/pkg/reconciler/common"
	util "knative.dev/operator/pkg/reconciler/common/testing"
	"knative.dev/operator/test"
	"knative.dev/operator/test/client"
	"knative.dev/operator/test/resources"
	"knative.dev/pkg/test/logstream"
)

// TestKnativeEventingPreviousVersion verifies the KnativeEventing creation in previous version.
// This test case is called before upgrading and after downgrading.
func TestKnativeEventingPreviousVersion(t *testing.T) {
	cancel := logstream.Start(t)
	defer cancel()
	clients := client.Setup(t)

	names := test.ResourceNames{
		KnativeEventing: test.OperatorName,
		Namespace:       test.EventingOperatorNamespace,
	}

	// Create a KnativeEventing
	if _, err := resources.EnsureKnativeEventingExists(clients.KnativeEventing(), names); err != nil {
		t.Fatalf("KnativeEventing %q failed to create: %v", names.KnativeEventing, err)
	}

	// Verify if resources match the requirement for the previous version before upgrade or after downgrade
	t.Run("verify resources", func(t *testing.T) {
		resources.AssertKEOperatorCRReadyStatus(t, clients, names)
		keventing, err := clients.KnativeEventing().Get(names.KnativeEventing, metav1.GetOptions{})
		if err != nil {
			t.Fatalf("Failed to get KnativeEventing CR: %v", err)
		}
		resources.SetKodataDir()
		defer os.Unsetenv(common.KoEnvKey)
		// Based on the status.version, get the deployment resources.
		defer os.Unsetenv(common.KoEnvKey)
		_, expectedDeployments := resources.GetExpectedDeployments(t, keventing)
		util.AssertEqual(t, len(expectedDeployments) > 0, true)
		resources.AssertKnativeDeploymentStatus(t, clients, names.Namespace, expectedDeployments)
	})
}

// TestKnativeServingPreviousVersion verifies the KnativeServing creation in previous version.
// This test case is called before upgrading and after downgrading.
func TestKnativeServingPreviousVersion(t *testing.T) {
	cancel := logstream.Start(t)
	defer cancel()
	clients := client.Setup(t)

	names := test.ResourceNames{
		KnativeServing: test.OperatorName,
		Namespace:      test.ServingOperatorNamespace,
	}

	// Create a KnativeServing
	if _, err := resources.EnsureKnativeServingExists(clients.KnativeServing(), names); err != nil {
		t.Fatalf("KnativeServing %q failed to create: %v", names.KnativeServing, err)
	}

	// Verify if resources match the requirement for the previous version before upgrade or after downgrade
	t.Run("verify resources", func(t *testing.T) {
		resources.AssertKSOperatorCRReadyStatus(t, clients, names)
		ks, err := clients.KnativeServing().Get(names.KnativeServing, metav1.GetOptions{})
		if err != nil {
			t.Fatalf("Failed to get KnativeServing CR: %v", err)
		}
		resources.SetKodataDir()
		defer os.Unsetenv(common.KoEnvKey)
		// Based on the status.version, get the deployment resources.
		defer os.Unsetenv(common.KoEnvKey)
		_, expectedDeployments := resources.GetExpectedDeployments(t, ks)
		util.AssertEqual(t, len(expectedDeployments) > 0, true)
		resources.AssertKnativeDeploymentStatus(t, clients, names.Namespace, expectedDeployments)
	})
}
