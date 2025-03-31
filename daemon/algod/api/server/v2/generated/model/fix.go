// Package model provides OpenAPI model definitions
package model

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// These aliases ensure proper compatibility with OpenAPI types
var (
	_ = openapi3.NewServer
	_ = openapi3.File{}
)
