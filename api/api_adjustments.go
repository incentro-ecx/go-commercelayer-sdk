/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AdjustmentsApiService AdjustmentsApi service
type AdjustmentsApiService service

type AdjustmentsApiDELETEAdjustmentsAdjustmentIdRequest struct {
	ctx          context.Context
	ApiService   *AdjustmentsApiService
	adjustmentId interface{}
}

func (r AdjustmentsApiDELETEAdjustmentsAdjustmentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEAdjustmentsAdjustmentIdExecute(r)
}

/*
DELETEAdjustmentsAdjustmentId Delete an adjustment

Delete an adjustment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param adjustmentId The resource's id
	@return AdjustmentsApiDELETEAdjustmentsAdjustmentIdRequest
*/
func (a *AdjustmentsApiService) DELETEAdjustmentsAdjustmentId(ctx context.Context, adjustmentId interface{}) AdjustmentsApiDELETEAdjustmentsAdjustmentIdRequest {
	return AdjustmentsApiDELETEAdjustmentsAdjustmentIdRequest{
		ApiService:   a,
		ctx:          ctx,
		adjustmentId: adjustmentId,
	}
}

// Execute executes the request
func (a *AdjustmentsApiService) DELETEAdjustmentsAdjustmentIdExecute(r AdjustmentsApiDELETEAdjustmentsAdjustmentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdjustmentsApiService.DELETEAdjustmentsAdjustmentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adjustments/{adjustmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adjustmentId"+"}", url.PathEscape(parameterValueToString(r.adjustmentId, "adjustmentId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type AdjustmentsApiGETAdjustmentsRequest struct {
	ctx        context.Context
	ApiService *AdjustmentsApiService
}

func (r AdjustmentsApiGETAdjustmentsRequest) Execute() (*GETAdjustments200Response, *http.Response, error) {
	return r.ApiService.GETAdjustmentsExecute(r)
}

/*
GETAdjustments List all adjustments

List all adjustments

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AdjustmentsApiGETAdjustmentsRequest
*/
func (a *AdjustmentsApiService) GETAdjustments(ctx context.Context) AdjustmentsApiGETAdjustmentsRequest {
	return AdjustmentsApiGETAdjustmentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETAdjustments200Response
func (a *AdjustmentsApiService) GETAdjustmentsExecute(r AdjustmentsApiGETAdjustmentsRequest) (*GETAdjustments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAdjustments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdjustmentsApiService.GETAdjustments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adjustments"

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type AdjustmentsApiGETAdjustmentsAdjustmentIdRequest struct {
	ctx          context.Context
	ApiService   *AdjustmentsApiService
	adjustmentId interface{}
}

func (r AdjustmentsApiGETAdjustmentsAdjustmentIdRequest) Execute() (*GETAdjustmentsAdjustmentId200Response, *http.Response, error) {
	return r.ApiService.GETAdjustmentsAdjustmentIdExecute(r)
}

/*
GETAdjustmentsAdjustmentId Retrieve an adjustment

Retrieve an adjustment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param adjustmentId The resource's id
	@return AdjustmentsApiGETAdjustmentsAdjustmentIdRequest
*/
func (a *AdjustmentsApiService) GETAdjustmentsAdjustmentId(ctx context.Context, adjustmentId interface{}) AdjustmentsApiGETAdjustmentsAdjustmentIdRequest {
	return AdjustmentsApiGETAdjustmentsAdjustmentIdRequest{
		ApiService:   a,
		ctx:          ctx,
		adjustmentId: adjustmentId,
	}
}

// Execute executes the request
//
//	@return GETAdjustmentsAdjustmentId200Response
func (a *AdjustmentsApiService) GETAdjustmentsAdjustmentIdExecute(r AdjustmentsApiGETAdjustmentsAdjustmentIdRequest) (*GETAdjustmentsAdjustmentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAdjustmentsAdjustmentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdjustmentsApiService.GETAdjustmentsAdjustmentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adjustments/{adjustmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adjustmentId"+"}", url.PathEscape(parameterValueToString(r.adjustmentId, "adjustmentId")), -1)

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type AdjustmentsApiPATCHAdjustmentsAdjustmentIdRequest struct {
	ctx              context.Context
	ApiService       *AdjustmentsApiService
	adjustmentUpdate *AdjustmentUpdate
	adjustmentId     interface{}
}

func (r AdjustmentsApiPATCHAdjustmentsAdjustmentIdRequest) AdjustmentUpdate(adjustmentUpdate AdjustmentUpdate) AdjustmentsApiPATCHAdjustmentsAdjustmentIdRequest {
	r.adjustmentUpdate = &adjustmentUpdate
	return r
}

func (r AdjustmentsApiPATCHAdjustmentsAdjustmentIdRequest) Execute() (*PATCHAdjustmentsAdjustmentId200Response, *http.Response, error) {
	return r.ApiService.PATCHAdjustmentsAdjustmentIdExecute(r)
}

/*
PATCHAdjustmentsAdjustmentId Update an adjustment

Update an adjustment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param adjustmentId The resource's id
	@return AdjustmentsApiPATCHAdjustmentsAdjustmentIdRequest
*/
func (a *AdjustmentsApiService) PATCHAdjustmentsAdjustmentId(ctx context.Context, adjustmentId interface{}) AdjustmentsApiPATCHAdjustmentsAdjustmentIdRequest {
	return AdjustmentsApiPATCHAdjustmentsAdjustmentIdRequest{
		ApiService:   a,
		ctx:          ctx,
		adjustmentId: adjustmentId,
	}
}

// Execute executes the request
//
//	@return PATCHAdjustmentsAdjustmentId200Response
func (a *AdjustmentsApiService) PATCHAdjustmentsAdjustmentIdExecute(r AdjustmentsApiPATCHAdjustmentsAdjustmentIdRequest) (*PATCHAdjustmentsAdjustmentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHAdjustmentsAdjustmentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdjustmentsApiService.PATCHAdjustmentsAdjustmentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adjustments/{adjustmentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"adjustmentId"+"}", url.PathEscape(parameterValueToString(r.adjustmentId, "adjustmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adjustmentUpdate == nil {
		return localVarReturnValue, nil, reportError("adjustmentUpdate is required and must be specified")
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
	localVarPostBody = r.adjustmentUpdate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type AdjustmentsApiPOSTAdjustmentsRequest struct {
	ctx              context.Context
	ApiService       *AdjustmentsApiService
	adjustmentCreate *AdjustmentCreate
}

func (r AdjustmentsApiPOSTAdjustmentsRequest) AdjustmentCreate(adjustmentCreate AdjustmentCreate) AdjustmentsApiPOSTAdjustmentsRequest {
	r.adjustmentCreate = &adjustmentCreate
	return r
}

func (r AdjustmentsApiPOSTAdjustmentsRequest) Execute() (*POSTAdjustments201Response, *http.Response, error) {
	return r.ApiService.POSTAdjustmentsExecute(r)
}

/*
POSTAdjustments Create an adjustment

Create an adjustment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AdjustmentsApiPOSTAdjustmentsRequest
*/
func (a *AdjustmentsApiService) POSTAdjustments(ctx context.Context) AdjustmentsApiPOSTAdjustmentsRequest {
	return AdjustmentsApiPOSTAdjustmentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTAdjustments201Response
func (a *AdjustmentsApiService) POSTAdjustmentsExecute(r AdjustmentsApiPOSTAdjustmentsRequest) (*POSTAdjustments201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTAdjustments201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdjustmentsApiService.POSTAdjustments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/adjustments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.adjustmentCreate == nil {
		return localVarReturnValue, nil, reportError("adjustmentCreate is required and must be specified")
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
	localVarPostBody = r.adjustmentCreate
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
