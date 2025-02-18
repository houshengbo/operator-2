/*
Copyright 2021 The Knative Authors

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

package v1alpha1

import (
	context "context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	eventingv1alpha1 "knative.dev/eventing/pkg/apis/eventing/v1alpha1"
	scheme "knative.dev/eventing/pkg/client/clientset/versioned/scheme"
)

// EventTransformsGetter has a method to return a EventTransformInterface.
// A group's client should implement this interface.
type EventTransformsGetter interface {
	EventTransforms(namespace string) EventTransformInterface
}

// EventTransformInterface has methods to work with EventTransform resources.
type EventTransformInterface interface {
	Create(ctx context.Context, eventTransform *eventingv1alpha1.EventTransform, opts v1.CreateOptions) (*eventingv1alpha1.EventTransform, error)
	Update(ctx context.Context, eventTransform *eventingv1alpha1.EventTransform, opts v1.UpdateOptions) (*eventingv1alpha1.EventTransform, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, eventTransform *eventingv1alpha1.EventTransform, opts v1.UpdateOptions) (*eventingv1alpha1.EventTransform, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*eventingv1alpha1.EventTransform, error)
	List(ctx context.Context, opts v1.ListOptions) (*eventingv1alpha1.EventTransformList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *eventingv1alpha1.EventTransform, err error)
	EventTransformExpansion
}

// eventTransforms implements EventTransformInterface
type eventTransforms struct {
	*gentype.ClientWithList[*eventingv1alpha1.EventTransform, *eventingv1alpha1.EventTransformList]
}

// newEventTransforms returns a EventTransforms
func newEventTransforms(c *EventingV1alpha1Client, namespace string) *eventTransforms {
	return &eventTransforms{
		gentype.NewClientWithList[*eventingv1alpha1.EventTransform, *eventingv1alpha1.EventTransformList](
			"eventtransforms",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *eventingv1alpha1.EventTransform { return &eventingv1alpha1.EventTransform{} },
			func() *eventingv1alpha1.EventTransformList { return &eventingv1alpha1.EventTransformList{} },
		),
	}
}
