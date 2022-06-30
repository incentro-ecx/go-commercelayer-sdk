/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ImportsApi interface {

	/*
		DELETEImportsImportId Delete an import

		Delete an import

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param importId The resource's id
		@return ImportsApiDELETEImportsImportIdRequest
	*/
	DELETEImportsImportId(ctx context.Context, importId string) ImportsApiDELETEImportsImportIdRequest

	// DELETEImportsImportIdExecute executes the request
	DELETEImportsImportIdExecute(r ImportsApiDELETEImportsImportIdRequest) (*http.Response, error)

	/*
		GETImports List all imports

		List all imports

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ImportsApiGETImportsRequest
	*/
	GETImports(ctx context.Context) ImportsApiGETImportsRequest

	// GETImportsExecute executes the request
	GETImportsExecute(r ImportsApiGETImportsRequest) (*http.Response, error)

	/*
		GETImportsImportId Retrieve an import

		Retrieve an import

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param importId The resource's id
		@return ImportsApiGETImportsImportIdRequest
	*/
	GETImportsImportId(ctx context.Context, importId string) ImportsApiGETImportsImportIdRequest

	// GETImportsImportIdExecute executes the request
	//  @return ModelImport
	GETImportsImportIdExecute(r ImportsApiGETImportsImportIdRequest) (*ModelImport, *http.Response, error)

	/*
		POSTImports Create an import

		Create an import

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ImportsApiPOSTImportsRequest
	*/
	POSTImports(ctx context.Context) ImportsApiPOSTImportsRequest

	// POSTImportsExecute executes the request
	POSTImportsExecute(r ImportsApiPOSTImportsRequest) (*http.Response, error)
}

// ImportsApiService ImportsApi service
type ImportsApiService service

type ImportsApiDELETEImportsImportIdRequest struct {
	ctx        context.Context
	ApiService ImportsApi
	importId   string
}

func (r ImportsApiDELETEImportsImportIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEImportsImportIdExecute(r)
}

/*
DELETEImportsImportId Delete an import

Delete an import

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param importId The resource's id
 @return ImportsApiDELETEImportsImportIdRequest
*/
func (a *ImportsApiService) DELETEImportsImportId(ctx context.Context, importId string) ImportsApiDELETEImportsImportIdRequest {
	return ImportsApiDELETEImportsImportIdRequest{
		ApiService: a,
		ctx:        ctx,
		importId:   importId,
	}
}

// Execute executes the request
func (a *ImportsApiService) DELETEImportsImportIdExecute(r ImportsApiDELETEImportsImportIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.DELETEImportsImportId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports/{importId}"
	localVarPath = strings.Replace(localVarPath, "{"+"importId"+"}", url.PathEscape(parameterToString(r.importId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ImportsApiGETImportsRequest struct {
	ctx        context.Context
	ApiService ImportsApi
}

func (r ImportsApiGETImportsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETImportsExecute(r)
}

/*
GETImports List all imports

List all imports

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ImportsApiGETImportsRequest
*/
func (a *ImportsApiService) GETImports(ctx context.Context) ImportsApiGETImportsRequest {
	return ImportsApiGETImportsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ImportsApiService) GETImportsExecute(r ImportsApiGETImportsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.GETImports")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ImportsApiGETImportsImportIdRequest struct {
	ctx        context.Context
	ApiService ImportsApi
	importId   string
}

func (r ImportsApiGETImportsImportIdRequest) Execute() (*ModelImport, *http.Response, error) {
	return r.ApiService.GETImportsImportIdExecute(r)
}

/*
GETImportsImportId Retrieve an import

Retrieve an import

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param importId The resource's id
 @return ImportsApiGETImportsImportIdRequest
*/
func (a *ImportsApiService) GETImportsImportId(ctx context.Context, importId string) ImportsApiGETImportsImportIdRequest {
	return ImportsApiGETImportsImportIdRequest{
		ApiService: a,
		ctx:        ctx,
		importId:   importId,
	}
}

// Execute executes the request
//  @return ModelImport
func (a *ImportsApiService) GETImportsImportIdExecute(r ImportsApiGETImportsImportIdRequest) (*ModelImport, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ModelImport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.GETImportsImportId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports/{importId}"
	localVarPath = strings.Replace(localVarPath, "{"+"importId"+"}", url.PathEscape(parameterToString(r.importId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.api+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ImportsApiPOSTImportsRequest struct {
	ctx          context.Context
	ApiService   ImportsApi
	importCreate *ImportCreate
}

func (r ImportsApiPOSTImportsRequest) ImportCreate(importCreate ImportCreate) ImportsApiPOSTImportsRequest {
	r.importCreate = &importCreate
	return r
}

func (r ImportsApiPOSTImportsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTImportsExecute(r)
}

/*
POSTImports Create an import

Create an import

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ImportsApiPOSTImportsRequest
*/
func (a *ImportsApiService) POSTImports(ctx context.Context) ImportsApiPOSTImportsRequest {
	return ImportsApiPOSTImportsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ImportsApiService) POSTImportsExecute(r ImportsApiPOSTImportsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.POSTImports")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importCreate == nil {
		return nil, reportError("importCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.importCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
