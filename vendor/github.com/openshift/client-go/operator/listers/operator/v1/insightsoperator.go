// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// InsightsOperatorLister helps list InsightsOperators.
// All objects returned here must be treated as read-only.
type InsightsOperatorLister interface {
	// List lists all InsightsOperators in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*operatorv1.InsightsOperator, err error)
	// Get retrieves the InsightsOperator from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*operatorv1.InsightsOperator, error)
	InsightsOperatorListerExpansion
}

// insightsOperatorLister implements the InsightsOperatorLister interface.
type insightsOperatorLister struct {
	listers.ResourceIndexer[*operatorv1.InsightsOperator]
}

// NewInsightsOperatorLister returns a new InsightsOperatorLister.
func NewInsightsOperatorLister(indexer cache.Indexer) InsightsOperatorLister {
	return &insightsOperatorLister{listers.New[*operatorv1.InsightsOperator](indexer, operatorv1.Resource("insightsoperator"))}
}
