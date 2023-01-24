/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
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

// CleanupsApiService CleanupsApi service
type CleanupsApiService service

type CleanupsApiDELETECleanupsCleanupIdRequest struct {
	ctx        context.Context
	ApiService *CleanupsApiService
	cleanupId  string
}

func (r CleanupsApiDELETECleanupsCleanupIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECleanupsCleanupIdExecute(r)
}

/*
DELETECleanupsCleanupId Delete a cleanup

Delete a cleanup

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cleanupId The resource's id
	@return CleanupsApiDELETECleanupsCleanupIdRequest
*/
func (a *CleanupsApiService) DELETECleanupsCleanupId(ctx context.Context, cleanupId string) CleanupsApiDELETECleanupsCleanupIdRequest {
	return CleanupsApiDELETECleanupsCleanupIdRequest{
		ApiService: a,
		ctx:        ctx,
		cleanupId:  cleanupId,
	}
}

// Execute executes the request
func (a *CleanupsApiService) DELETECleanupsCleanupIdExecute(r CleanupsApiDELETECleanupsCleanupIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CleanupsApiService.DELETECleanupsCleanupId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cleanups/{cleanupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"cleanupId"+"}", url.PathEscape(parameterToString(r.cleanupId, "")), -1)

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

type CleanupsApiGETCleanupsRequest struct {
	ctx        context.Context
	ApiService *CleanupsApiService
}

func (r CleanupsApiGETCleanupsRequest) Execute() (*GETCleanups200Response, *http.Response, error) {
	return r.ApiService.GETCleanupsExecute(r)
}

/*
GETCleanups List all cleanups

List all cleanups

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CleanupsApiGETCleanupsRequest
*/
func (a *CleanupsApiService) GETCleanups(ctx context.Context) CleanupsApiGETCleanupsRequest {
	return CleanupsApiGETCleanupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETCleanups200Response
func (a *CleanupsApiService) GETCleanupsExecute(r CleanupsApiGETCleanupsRequest) (*GETCleanups200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCleanups200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CleanupsApiService.GETCleanups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cleanups"

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

type CleanupsApiGETCleanupsCleanupIdRequest struct {
	ctx        context.Context
	ApiService *CleanupsApiService
	cleanupId  string
}

func (r CleanupsApiGETCleanupsCleanupIdRequest) Execute() (*GETCleanupsCleanupId200Response, *http.Response, error) {
	return r.ApiService.GETCleanupsCleanupIdExecute(r)
}

/*
GETCleanupsCleanupId Retrieve a cleanup

Retrieve a cleanup

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cleanupId The resource's id
	@return CleanupsApiGETCleanupsCleanupIdRequest
*/
func (a *CleanupsApiService) GETCleanupsCleanupId(ctx context.Context, cleanupId string) CleanupsApiGETCleanupsCleanupIdRequest {
	return CleanupsApiGETCleanupsCleanupIdRequest{
		ApiService: a,
		ctx:        ctx,
		cleanupId:  cleanupId,
	}
}

// Execute executes the request
//
//	@return GETCleanupsCleanupId200Response
func (a *CleanupsApiService) GETCleanupsCleanupIdExecute(r CleanupsApiGETCleanupsCleanupIdRequest) (*GETCleanupsCleanupId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCleanupsCleanupId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CleanupsApiService.GETCleanupsCleanupId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cleanups/{cleanupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"cleanupId"+"}", url.PathEscape(parameterToString(r.cleanupId, "")), -1)

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

type CleanupsApiPOSTCleanupsRequest struct {
	ctx           context.Context
	ApiService    *CleanupsApiService
	cleanupCreate *CleanupCreate
}

func (r CleanupsApiPOSTCleanupsRequest) CleanupCreate(cleanupCreate CleanupCreate) CleanupsApiPOSTCleanupsRequest {
	r.cleanupCreate = &cleanupCreate
	return r
}

func (r CleanupsApiPOSTCleanupsRequest) Execute() (*POSTCleanups201Response, *http.Response, error) {
	return r.ApiService.POSTCleanupsExecute(r)
}

/*
POSTCleanups Create a cleanup

Create a cleanup

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CleanupsApiPOSTCleanupsRequest
*/
func (a *CleanupsApiService) POSTCleanups(ctx context.Context) CleanupsApiPOSTCleanupsRequest {
	return CleanupsApiPOSTCleanupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTCleanups201Response
func (a *CleanupsApiService) POSTCleanupsExecute(r CleanupsApiPOSTCleanupsRequest) (*POSTCleanups201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTCleanups201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CleanupsApiService.POSTCleanups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cleanups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cleanupCreate == nil {
		return localVarReturnValue, nil, reportError("cleanupCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.api+json"}

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
	// body params
	localVarPostBody = r.cleanupCreate
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
