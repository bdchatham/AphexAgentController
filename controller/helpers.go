package controller

import (
	"k8s.io/apimachinery/pkg/api/resource"
)

// mustParseQuantity parses a resource quantity string and panics on error.
// Safe to use for hardcoded values in controller logic.
func mustParseQuantity(s string) resource.Quantity {
	q, err := resource.ParseQuantity(s)
	if err != nil {
		panic(err)
	}
	return q
}
