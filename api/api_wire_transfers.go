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

// WireTransfersApiService WireTransfersApi service
type WireTransfersApiService service

type ApiDELETEWireTransfersWireTransferIdRequest struct {
	ctx            context.Context
	ApiService     *WireTransfersApiService
	wireTransferId string
}

func (r ApiDELETEWireTransfersWireTransferIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEWireTransfersWireTransferIdExecute(r)
}

/*
DELETEWireTransfersWireTransferId Delete a wire transfer

Delete a wire transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wireTransferId The resource's id
 @return ApiDELETEWireTransfersWireTransferIdRequest
*/
func (a *WireTransfersApiService) DELETEWireTransfersWireTransferId(ctx context.Context, wireTransferId string) ApiDELETEWireTransfersWireTransferIdRequest {
	return ApiDELETEWireTransfersWireTransferIdRequest{
		ApiService:     a,
		ctx:            ctx,
		wireTransferId: wireTransferId,
	}
}

// Execute executes the request
func (a *WireTransfersApiService) DELETEWireTransfersWireTransferIdExecute(r ApiDELETEWireTransfersWireTransferIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.DELETEWireTransfersWireTransferId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers/{wireTransferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"wireTransferId"+"}", url.PathEscape(parameterToString(r.wireTransferId, "")), -1)

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

type ApiGETWireTransfersRequest struct {
	ctx        context.Context
	ApiService *WireTransfersApiService
}

func (r ApiGETWireTransfersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETWireTransfersExecute(r)
}

/*
GETWireTransfers List all wire transfers

List all wire transfers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETWireTransfersRequest
*/
func (a *WireTransfersApiService) GETWireTransfers(ctx context.Context) ApiGETWireTransfersRequest {
	return ApiGETWireTransfersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *WireTransfersApiService) GETWireTransfersExecute(r ApiGETWireTransfersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.GETWireTransfers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers"

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

type ApiGETWireTransfersWireTransferIdRequest struct {
	ctx            context.Context
	ApiService     *WireTransfersApiService
	wireTransferId string
}

func (r ApiGETWireTransfersWireTransferIdRequest) Execute() (*WireTransfer, *http.Response, error) {
	return r.ApiService.GETWireTransfersWireTransferIdExecute(r)
}

/*
GETWireTransfersWireTransferId Retrieve a wire transfer

Retrieve a wire transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wireTransferId The resource's id
 @return ApiGETWireTransfersWireTransferIdRequest
*/
func (a *WireTransfersApiService) GETWireTransfersWireTransferId(ctx context.Context, wireTransferId string) ApiGETWireTransfersWireTransferIdRequest {
	return ApiGETWireTransfersWireTransferIdRequest{
		ApiService:     a,
		ctx:            ctx,
		wireTransferId: wireTransferId,
	}
}

// Execute executes the request
//  @return WireTransfer
func (a *WireTransfersApiService) GETWireTransfersWireTransferIdExecute(r ApiGETWireTransfersWireTransferIdRequest) (*WireTransfer, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WireTransfer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.GETWireTransfersWireTransferId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers/{wireTransferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"wireTransferId"+"}", url.PathEscape(parameterToString(r.wireTransferId, "")), -1)

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

type ApiPATCHWireTransfersWireTransferIdRequest struct {
	ctx                context.Context
	ApiService         *WireTransfersApiService
	wireTransferUpdate *WireTransferUpdate
	wireTransferId     string
}

func (r ApiPATCHWireTransfersWireTransferIdRequest) WireTransferUpdate(wireTransferUpdate WireTransferUpdate) ApiPATCHWireTransfersWireTransferIdRequest {
	r.wireTransferUpdate = &wireTransferUpdate
	return r
}

func (r ApiPATCHWireTransfersWireTransferIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHWireTransfersWireTransferIdExecute(r)
}

/*
PATCHWireTransfersWireTransferId Update a wire transfer

Update a wire transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param wireTransferId The resource's id
 @return ApiPATCHWireTransfersWireTransferIdRequest
*/
func (a *WireTransfersApiService) PATCHWireTransfersWireTransferId(ctx context.Context, wireTransferId string) ApiPATCHWireTransfersWireTransferIdRequest {
	return ApiPATCHWireTransfersWireTransferIdRequest{
		ApiService:     a,
		ctx:            ctx,
		wireTransferId: wireTransferId,
	}
}

// Execute executes the request
func (a *WireTransfersApiService) PATCHWireTransfersWireTransferIdExecute(r ApiPATCHWireTransfersWireTransferIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.PATCHWireTransfersWireTransferId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers/{wireTransferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"wireTransferId"+"}", url.PathEscape(parameterToString(r.wireTransferId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.wireTransferUpdate == nil {
		return nil, reportError("wireTransferUpdate is required and must be specified")
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
	localVarPostBody = r.wireTransferUpdate
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

type ApiPOSTWireTransfersRequest struct {
	ctx                context.Context
	ApiService         *WireTransfersApiService
	wireTransferCreate *WireTransferCreate
}

func (r ApiPOSTWireTransfersRequest) WireTransferCreate(wireTransferCreate WireTransferCreate) ApiPOSTWireTransfersRequest {
	r.wireTransferCreate = &wireTransferCreate
	return r
}

func (r ApiPOSTWireTransfersRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTWireTransfersExecute(r)
}

/*
POSTWireTransfers Create a wire transfer

Create a wire transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTWireTransfersRequest
*/
func (a *WireTransfersApiService) POSTWireTransfers(ctx context.Context) ApiPOSTWireTransfersRequest {
	return ApiPOSTWireTransfersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *WireTransfersApiService) POSTWireTransfersExecute(r ApiPOSTWireTransfersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WireTransfersApiService.POSTWireTransfers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wire_transfers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.wireTransferCreate == nil {
		return nil, reportError("wireTransferCreate is required and must be specified")
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
	localVarPostBody = r.wireTransferCreate
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
