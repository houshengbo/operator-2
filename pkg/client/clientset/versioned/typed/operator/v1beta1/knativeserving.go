/*
Copyright 2025 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	operatorv1beta1 "knative.dev/operator/pkg/apis/operator/v1beta1"
	scheme "knative.dev/operator/pkg/client/clientset/versioned/scheme"
)

// KnativeServingsGetter has a method to return a KnativeServingInterface.
// A group's client should implement this interface.
type KnativeServingsGetter interface {
	KnativeServings(namespace string) KnativeServingInterface
}

// KnativeServingInterface has methods to work with KnativeServing resources.
type KnativeServingInterface interface {
	Create(ctx context.Context, knativeServing *operatorv1beta1.KnativeServing, opts v1.CreateOptions) (*operatorv1beta1.KnativeServing, error)
	Update(ctx context.Context, knativeServing *operatorv1beta1.KnativeServing, opts v1.UpdateOptions) (*operatorv1beta1.KnativeServing, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, knativeServing *operatorv1beta1.KnativeServing, opts v1.UpdateOptions) (*operatorv1beta1.KnativeServing, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*operatorv1beta1.KnativeServing, error)
	List(ctx context.Context, opts v1.ListOptions) (*operatorv1beta1.KnativeServingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1beta1.KnativeServing, err error)
	KnativeServingExpansion
}

// knativeServings implements KnativeServingInterface
type knativeServings struct {
	*gentype.ClientWithList[*operatorv1beta1.KnativeServing, *operatorv1beta1.KnativeServingList]
}

// newKnativeServings returns a KnativeServings
func newKnativeServings(c *OperatorV1beta1Client, namespace string) *knativeServings {
	return &knativeServings{
		gentype.NewClientWithList[*operatorv1beta1.KnativeServing, *operatorv1beta1.KnativeServingList](
			"knativeservings",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *operatorv1beta1.KnativeServing { return &operatorv1beta1.KnativeServing{} },
			func() *operatorv1beta1.KnativeServingList { return &operatorv1beta1.KnativeServingList{} },
		),
	}
}
