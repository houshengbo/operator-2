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
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1beta1 "knative.dev/operator/pkg/apis/operator/v1beta1"
)

// KnativeServingLister helps list KnativeServings.
// All objects returned here must be treated as read-only.
type KnativeServingLister interface {
	// List lists all KnativeServings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.KnativeServing, err error)
	// KnativeServings returns an object that can list and get KnativeServings.
	KnativeServings(namespace string) KnativeServingNamespaceLister
	KnativeServingListerExpansion
}

// knativeServingLister implements the KnativeServingLister interface.
type knativeServingLister struct {
	listers.ResourceIndexer[*v1beta1.KnativeServing]
}

// NewKnativeServingLister returns a new KnativeServingLister.
func NewKnativeServingLister(indexer cache.Indexer) KnativeServingLister {
	return &knativeServingLister{listers.New[*v1beta1.KnativeServing](indexer, v1beta1.Resource("knativeserving"))}
}

// KnativeServings returns an object that can list and get KnativeServings.
func (s *knativeServingLister) KnativeServings(namespace string) KnativeServingNamespaceLister {
	return knativeServingNamespaceLister{listers.NewNamespaced[*v1beta1.KnativeServing](s.ResourceIndexer, namespace)}
}

// KnativeServingNamespaceLister helps list and get KnativeServings.
// All objects returned here must be treated as read-only.
type KnativeServingNamespaceLister interface {
	// List lists all KnativeServings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.KnativeServing, err error)
	// Get retrieves the KnativeServing from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.KnativeServing, error)
	KnativeServingNamespaceListerExpansion
}

// knativeServingNamespaceLister implements the KnativeServingNamespaceLister
// interface.
type knativeServingNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.KnativeServing]
}
