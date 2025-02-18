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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	operatorv1beta1 "knative.dev/operator/pkg/apis/operator/v1beta1"
)

// KnativeEventingLister helps list KnativeEventings.
// All objects returned here must be treated as read-only.
type KnativeEventingLister interface {
	// List lists all KnativeEventings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*operatorv1beta1.KnativeEventing, err error)
	// KnativeEventings returns an object that can list and get KnativeEventings.
	KnativeEventings(namespace string) KnativeEventingNamespaceLister
	KnativeEventingListerExpansion
}

// knativeEventingLister implements the KnativeEventingLister interface.
type knativeEventingLister struct {
	listers.ResourceIndexer[*operatorv1beta1.KnativeEventing]
}

// NewKnativeEventingLister returns a new KnativeEventingLister.
func NewKnativeEventingLister(indexer cache.Indexer) KnativeEventingLister {
	return &knativeEventingLister{listers.New[*operatorv1beta1.KnativeEventing](indexer, operatorv1beta1.Resource("knativeeventing"))}
}

// KnativeEventings returns an object that can list and get KnativeEventings.
func (s *knativeEventingLister) KnativeEventings(namespace string) KnativeEventingNamespaceLister {
	return knativeEventingNamespaceLister{listers.NewNamespaced[*operatorv1beta1.KnativeEventing](s.ResourceIndexer, namespace)}
}

// KnativeEventingNamespaceLister helps list and get KnativeEventings.
// All objects returned here must be treated as read-only.
type KnativeEventingNamespaceLister interface {
	// List lists all KnativeEventings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*operatorv1beta1.KnativeEventing, err error)
	// Get retrieves the KnativeEventing from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*operatorv1beta1.KnativeEventing, error)
	KnativeEventingNamespaceListerExpansion
}

// knativeEventingNamespaceLister implements the KnativeEventingNamespaceLister
// interface.
type knativeEventingNamespaceLister struct {
	listers.ResourceIndexer[*operatorv1beta1.KnativeEventing]
}
