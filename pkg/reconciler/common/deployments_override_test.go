/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	mf "github.com/manifestival/manifestival"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes/scheme"
	servingv1alpha1 "knative.dev/operator/pkg/apis/operator/v1alpha1"
)

type expDeployments struct {
	expLabels              map[string]string
	expTemplateLabels      map[string]string
	expAnnotations         map[string]string
	expTemplateAnnotations map[string]string
	expReplicas            int32
}

func TestDeploymentsTransform(t *testing.T) {
	tests := []struct {
		name           string
		override       []servingv1alpha1.DeploymentOverride
		globalReplicas int32
		expDeployment  map[string]expDeployments
	}{{
		name:     "no override",
		override: nil,
		expDeployment: map[string]expDeployments{"controller": {
			expLabels:              map[string]string{"serving.knative.dev/release": "v0.13.0"},
			expTemplateLabels:      map[string]string{"serving.knative.dev/release": "v0.13.0", "app": "controller"},
			expAnnotations:         nil,
			expTemplateAnnotations: map[string]string{"cluster-autoscaler.kubernetes.io/safe-to-evict": "true"},
			expReplicas:            0,
		}},
	}, {
		name: "simple override",
		override: []servingv1alpha1.DeploymentOverride{
			{
				Name:        "controller",
				Labels:      map[string]string{"a": "b"},
				Annotations: map[string]string{"c": "d"},
				Replicas:    5,
			},
		},
		globalReplicas: 10,
		expDeployment: map[string]expDeployments{"controller": {
			expLabels:              map[string]string{"serving.knative.dev/release": "v0.13.0", "a": "b"},
			expTemplateLabels:      map[string]string{"serving.knative.dev/release": "v0.13.0", "app": "controller", "a": "b"},
			expAnnotations:         map[string]string{"c": "d"},
			expTemplateAnnotations: map[string]string{"cluster-autoscaler.kubernetes.io/safe-to-evict": "true", "c": "d"},
			expReplicas:            5,
		}},
	}, {
		name: "no replicas in deploymentoverride, use global replicas",
		override: []servingv1alpha1.DeploymentOverride{
			{Name: "controller"},
		},
		globalReplicas: 10,
		expDeployment: map[string]expDeployments{"controller": {
			expLabels:              map[string]string{"serving.knative.dev/release": "v0.13.0"},
			expTemplateLabels:      map[string]string{"serving.knative.dev/release": "v0.13.0", "app": "controller"},
			expTemplateAnnotations: map[string]string{"cluster-autoscaler.kubernetes.io/safe-to-evict": "true"},
			expReplicas:            10,
		}},
	}, {
		name: "neither replicas in deploymentoverride nor global replicas",
		override: []servingv1alpha1.DeploymentOverride{
			{Name: "controller"},
		},
		expDeployment: map[string]expDeployments{"controller": {
			expLabels:              map[string]string{"serving.knative.dev/release": "v0.13.0"},
			expTemplateLabels:      map[string]string{"serving.knative.dev/release": "v0.13.0", "app": "controller"},
			expTemplateAnnotations: map[string]string{"cluster-autoscaler.kubernetes.io/safe-to-evict": "true"},
			expReplicas:            0,
		}},
	}, {
		name: "multiple override",
		override: []servingv1alpha1.DeploymentOverride{
			{
				Name:        "controller",
				Labels:      map[string]string{"a": "b"},
				Annotations: map[string]string{"c": "d"},
				Replicas:    5,
			},
			{
				Name:        "webhook",
				Labels:      map[string]string{"e": "f"},
				Annotations: map[string]string{"g": "h"},
				Replicas:    4,
			},
		},
		globalReplicas: 10,
		expDeployment: map[string]expDeployments{
			"controller": {
				expLabels:              map[string]string{"serving.knative.dev/release": "v0.13.0", "a": "b"},
				expTemplateLabels:      map[string]string{"serving.knative.dev/release": "v0.13.0", "app": "controller", "a": "b"},
				expAnnotations:         map[string]string{"c": "d"},
				expTemplateAnnotations: map[string]string{"cluster-autoscaler.kubernetes.io/safe-to-evict": "true", "c": "d"},
				expReplicas:            5,
			},
			"webhook": {
				expLabels:              map[string]string{"serving.knative.dev/release": "v0.13.0", "e": "f"},
				expTemplateLabels:      map[string]string{"serving.knative.dev/release": "v0.13.0", "app": "webhook", "role": "webhook", "e": "f"},
				expAnnotations:         map[string]string{"g": "h"},
				expTemplateAnnotations: map[string]string{"cluster-autoscaler.kubernetes.io/safe-to-evict": "false", "g": "h"},
				expReplicas:            4,
			},
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			manifest, err := mf.NewManifest("testdata/manifest.yaml")
			if err != nil {
				t.Fatalf("Failed to create manifest: %v", err)
			}

			ks := &servingv1alpha1.KnativeServing{
				Spec: servingv1alpha1.KnativeServingSpec{
					CommonSpec: servingv1alpha1.CommonSpec{
						DeploymentOverride: test.override,
						HighAvailability: &servingv1alpha1.HighAvailability{
							Replicas: test.globalReplicas,
						},
					},
				},
			}

			//manifest, err = manifest.Transform(DeploymentsTransform(ks, log), HighAvailabilityTransform(ks, log))
			manifest, err = manifest.Transform(HighAvailabilityTransform(ks, log), DeploymentsTransform(ks, log))
			if err != nil {
				t.Fatalf("Failed to transform manifest: %v", err)
			}

			for expName, d := range test.expDeployment {
				for _, u := range manifest.Resources() {
					if u.GetKind() == "Deployment" && u.GetName() == expName {
						got := &appsv1.Deployment{}
						if err := scheme.Scheme.Convert(&u, got, nil); err != nil {
							t.Fatalf("Failed to convert unstructured to deployment: %v", err)
						}

						replicas := int32(0)
						if got.Spec.Replicas != nil {
							replicas = *got.Spec.Replicas
						}
						if diff := cmp.Diff(replicas, d.expReplicas); diff != "" {
							t.Fatalf("Unexpected replicas: %v", diff)
						}

						if diff := cmp.Diff(got.GetLabels(), d.expLabels); diff != "" {
							t.Fatalf("Unexpected labels: %v", diff)
						}
						if diff := cmp.Diff(got.Spec.Template.GetLabels(), d.expTemplateLabels); diff != "" {
							t.Fatalf("Unexpected labels in pod template: %v", diff)
						}

						if diff := cmp.Diff(got.GetAnnotations(), d.expAnnotations); diff != "" {
							t.Fatalf("Unexpected annotations: %v", diff)
						}
						if diff := cmp.Diff(got.Spec.Template.GetAnnotations(), d.expTemplateAnnotations); diff != "" {
							t.Fatalf("Unexpected annotations in pod template: %v", diff)
						}
					}
				}
			}
		})
	}
}
