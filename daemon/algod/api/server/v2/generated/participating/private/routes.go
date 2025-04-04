// Package private provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/DePINNetwork/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/DePINNetwork/depin-sdk/daemon/algod/api/server/v2/generated/model"
	"github.com/DePINNetwork/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Gets the merged config file.
	// (GET /debug/settings/config)
	GetConfig(ctx echo.Context) error

	// (GET /debug/settings/pprof)
	GetDebugSettingsProf(ctx echo.Context) error

	// (PUT /debug/settings/pprof)
	PutDebugSettingsProf(ctx echo.Context) error
	// Return a list of participation keys
	// (GET /v2/participation)
	GetParticipationKeys(ctx echo.Context) error
	// Add a participation key to the node
	// (POST /v2/participation)
	AddParticipationKey(ctx echo.Context) error
	// Generate and install participation keys to the node.
	// (POST /v2/participation/generate/{address})
	GenerateParticipationKeys(ctx echo.Context, address string, params GenerateParticipationKeysParams) error
	// Delete a given participation key by ID
	// (DELETE /v2/participation/{participation-id})
	DeleteParticipationKeyByID(ctx echo.Context, participationId string) error
	// Get participation key info given a participation ID
	// (GET /v2/participation/{participation-id})
	GetParticipationKeyByID(ctx echo.Context, participationId string) error
	// Append state proof keys to a participation key
	// (POST /v2/participation/{participation-id})
	AppendKeys(ctx echo.Context, participationId string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetConfig converts echo context to params.
func (w *ServerInterfaceWrapper) GetConfig(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetConfig(ctx)
	return err
}

// GetDebugSettingsProf converts echo context to params.
func (w *ServerInterfaceWrapper) GetDebugSettingsProf(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDebugSettingsProf(ctx)
	return err
}

// PutDebugSettingsProf converts echo context to params.
func (w *ServerInterfaceWrapper) PutDebugSettingsProf(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutDebugSettingsProf(ctx)
	return err
}

// GetParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeys(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeys(ctx)
	return err
}

// AddParticipationKey converts echo context to params.
func (w *ServerInterfaceWrapper) AddParticipationKey(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddParticipationKey(ctx)
	return err
}

// GenerateParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) GenerateParticipationKeys(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameterWithLocation("simple", false, "address", runtime.ParamLocationPath, ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GenerateParticipationKeysParams
	// ------------- Optional query parameter "dilution" -------------

	err = runtime.BindQueryParameter("form", true, false, "dilution", ctx.QueryParams(), &params.Dilution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter dilution: %s", err))
	}

	// ------------- Required query parameter "first" -------------

	err = runtime.BindQueryParameter("form", true, true, "first", ctx.QueryParams(), &params.First)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter first: %s", err))
	}

	// ------------- Required query parameter "last" -------------

	err = runtime.BindQueryParameter("form", true, true, "last", ctx.QueryParams(), &params.Last)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter last: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GenerateParticipationKeys(ctx, address, params)
	return err
}

// DeleteParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteParticipationKeyByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteParticipationKeyByID(ctx, participationId)
	return err
}

// GetParticipationKeyByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetParticipationKeyByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetParticipationKeyByID(ctx, participationId)
	return err
}

// AppendKeys converts echo context to params.
func (w *ServerInterfaceWrapper) AppendKeys(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "participation-id" -------------
	var participationId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "participation-id", runtime.ParamLocationPath, ctx.Param("participation-id"), &participationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter participation-id: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AppendKeys(ctx, participationId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/debug/settings/config", wrapper.GetConfig, m...)
	router.GET(baseURL+"/debug/settings/pprof", wrapper.GetDebugSettingsProf, m...)
	router.PUT(baseURL+"/debug/settings/pprof", wrapper.PutDebugSettingsProf, m...)
	router.GET(baseURL+"/v2/participation", wrapper.GetParticipationKeys, m...)
	router.POST(baseURL+"/v2/participation", wrapper.AddParticipationKey, m...)
	router.POST(baseURL+"/v2/participation/generate/:address", wrapper.GenerateParticipationKeys, m...)
	router.DELETE(baseURL+"/v2/participation/:participation-id", wrapper.DeleteParticipationKeyByID, m...)
	router.GET(baseURL+"/v2/participation/:participation-id", wrapper.GetParticipationKeyByID, m...)
	router.POST(baseURL+"/v2/participation/:participation-id", wrapper.AppendKeys, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
