// Package model provides OpenAPI model definitions
package model

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// File is an alias to openapi3.File for backwards compatibility
type File = openapi3.File

// Register types that need aliasing with OpenAPI
func init() {
	// This ensures that the alias is used properly
	var _ = openapi3.NewServer("").Description
}
