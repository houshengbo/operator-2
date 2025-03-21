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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	eventingv1alpha1 "knative.dev/eventing/pkg/apis/eventing/v1alpha1"
)

// EventTransformLister helps list EventTransforms.
// All objects returned here must be treated as read-only.
type EventTransformLister interface {
	// List lists all EventTransforms in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*eventingv1alpha1.EventTransform, err error)
	// EventTransforms returns an object that can list and get EventTransforms.
	EventTransforms(namespace string) EventTransformNamespaceLister
	EventTransformListerExpansion
}

// eventTransformLister implements the EventTransformLister interface.
type eventTransformLister struct {
	listers.ResourceIndexer[*eventingv1alpha1.EventTransform]
}

// NewEventTransformLister returns a new EventTransformLister.
func NewEventTransformLister(indexer cache.Indexer) EventTransformLister {
	return &eventTransformLister{listers.New[*eventingv1alpha1.EventTransform](indexer, eventingv1alpha1.Resource("eventtransform"))}
}

// EventTransforms returns an object that can list and get EventTransforms.
func (s *eventTransformLister) EventTransforms(namespace string) EventTransformNamespaceLister {
	return eventTransformNamespaceLister{listers.NewNamespaced[*eventingv1alpha1.EventTransform](s.ResourceIndexer, namespace)}
}

// EventTransformNamespaceLister helps list and get EventTransforms.
// All objects returned here must be treated as read-only.
type EventTransformNamespaceLister interface {
	// List lists all EventTransforms in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*eventingv1alpha1.EventTransform, err error)
	// Get retrieves the EventTransform from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*eventingv1alpha1.EventTransform, error)
	EventTransformNamespaceListerExpansion
}

// eventTransformNamespaceLister implements the EventTransformNamespaceLister
// interface.
type eventTransformNamespaceLister struct {
	listers.ResourceIndexer[*eventingv1alpha1.EventTransform]
}
