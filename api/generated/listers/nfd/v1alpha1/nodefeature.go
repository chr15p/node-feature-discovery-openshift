/*
Copyright 2024 The Kubernetes Authors.

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
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "github.com/openshift/node-feature-discovery/api/nfd/v1alpha1"
)

// NodeFeatureLister helps list NodeFeatures.
// All objects returned here must be treated as read-only.
type NodeFeatureLister interface {
	// List lists all NodeFeatures in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeFeature, err error)
	// NodeFeatures returns an object that can list and get NodeFeatures.
	NodeFeatures(namespace string) NodeFeatureNamespaceLister
	NodeFeatureListerExpansion
}

// nodeFeatureLister implements the NodeFeatureLister interface.
type nodeFeatureLister struct {
	listers.ResourceIndexer[*v1alpha1.NodeFeature]
}

// NewNodeFeatureLister returns a new NodeFeatureLister.
func NewNodeFeatureLister(indexer cache.Indexer) NodeFeatureLister {
	return &nodeFeatureLister{listers.New[*v1alpha1.NodeFeature](indexer, v1alpha1.Resource("nodefeature"))}
}

// NodeFeatures returns an object that can list and get NodeFeatures.
func (s *nodeFeatureLister) NodeFeatures(namespace string) NodeFeatureNamespaceLister {
	return nodeFeatureNamespaceLister{listers.NewNamespaced[*v1alpha1.NodeFeature](s.ResourceIndexer, namespace)}
}

// NodeFeatureNamespaceLister helps list and get NodeFeatures.
// All objects returned here must be treated as read-only.
type NodeFeatureNamespaceLister interface {
	// List lists all NodeFeatures in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeFeature, err error)
	// Get retrieves the NodeFeature from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NodeFeature, error)
	NodeFeatureNamespaceListerExpansion
}

// nodeFeatureNamespaceLister implements the NodeFeatureNamespaceLister
// interface.
type nodeFeatureNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.NodeFeature]
}
