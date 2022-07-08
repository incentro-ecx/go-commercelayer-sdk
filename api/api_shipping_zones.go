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

// ShippingZonesApiService ShippingZonesApi service
type ShippingZonesApiService service

type ApiDELETEShippingZonesShippingZoneIdRequest struct {
	ctx            context.Context
	ApiService     *ShippingZonesApiService
	shippingZoneId string
}

func (r ApiDELETEShippingZonesShippingZoneIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEShippingZonesShippingZoneIdExecute(r)
}

/*
DELETEShippingZonesShippingZoneId Delete a shipping zone

Delete a shipping zone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingZoneId The resource's id
 @return ApiDELETEShippingZonesShippingZoneIdRequest
*/
func (a *ShippingZonesApiService) DELETEShippingZonesShippingZoneId(ctx context.Context, shippingZoneId string) ApiDELETEShippingZonesShippingZoneIdRequest {
	return ApiDELETEShippingZonesShippingZoneIdRequest{
		ApiService:     a,
		ctx:            ctx,
		shippingZoneId: shippingZoneId,
	}
}

// Execute executes the request
func (a *ShippingZonesApiService) DELETEShippingZonesShippingZoneIdExecute(r ApiDELETEShippingZonesShippingZoneIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.DELETEShippingZonesShippingZoneId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones/{shippingZoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingZoneId"+"}", url.PathEscape(parameterToString(r.shippingZoneId, "")), -1)

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

type ApiGETShippingMethodIdShippingZoneRequest struct {
	ctx              context.Context
	ApiService       *ShippingZonesApiService
	shippingMethodId string
}

func (r ApiGETShippingMethodIdShippingZoneRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingMethodIdShippingZoneExecute(r)
}

/*
GETShippingMethodIdShippingZone Retrieve the shipping zone associated to the shipping method

Retrieve the shipping zone associated to the shipping method

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingMethodId The resource's id
 @return ApiGETShippingMethodIdShippingZoneRequest
*/
func (a *ShippingZonesApiService) GETShippingMethodIdShippingZone(ctx context.Context, shippingMethodId string) ApiGETShippingMethodIdShippingZoneRequest {
	return ApiGETShippingMethodIdShippingZoneRequest{
		ApiService:       a,
		ctx:              ctx,
		shippingMethodId: shippingMethodId,
	}
}

// Execute executes the request
func (a *ShippingZonesApiService) GETShippingMethodIdShippingZoneExecute(r ApiGETShippingMethodIdShippingZoneRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.GETShippingMethodIdShippingZone")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_methods/{shippingMethodId}/shipping_zone"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingMethodId"+"}", url.PathEscape(parameterToString(r.shippingMethodId, "")), -1)

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

type ApiGETShippingZonesRequest struct {
	ctx        context.Context
	ApiService *ShippingZonesApiService
}

func (r ApiGETShippingZonesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETShippingZonesExecute(r)
}

/*
GETShippingZones List all shipping zones

List all shipping zones

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETShippingZonesRequest
*/
func (a *ShippingZonesApiService) GETShippingZones(ctx context.Context) ApiGETShippingZonesRequest {
	return ApiGETShippingZonesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ShippingZonesApiService) GETShippingZonesExecute(r ApiGETShippingZonesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.GETShippingZones")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones"

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

type ApiGETShippingZonesShippingZoneIdRequest struct {
	ctx            context.Context
	ApiService     *ShippingZonesApiService
	shippingZoneId string
}

func (r ApiGETShippingZonesShippingZoneIdRequest) Execute() (*ShippingZone, *http.Response, error) {
	return r.ApiService.GETShippingZonesShippingZoneIdExecute(r)
}

/*
GETShippingZonesShippingZoneId Retrieve a shipping zone

Retrieve a shipping zone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingZoneId The resource's id
 @return ApiGETShippingZonesShippingZoneIdRequest
*/
func (a *ShippingZonesApiService) GETShippingZonesShippingZoneId(ctx context.Context, shippingZoneId string) ApiGETShippingZonesShippingZoneIdRequest {
	return ApiGETShippingZonesShippingZoneIdRequest{
		ApiService:     a,
		ctx:            ctx,
		shippingZoneId: shippingZoneId,
	}
}

// Execute executes the request
//  @return ShippingZone
func (a *ShippingZonesApiService) GETShippingZonesShippingZoneIdExecute(r ApiGETShippingZonesShippingZoneIdRequest) (*ShippingZone, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ShippingZone
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.GETShippingZonesShippingZoneId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones/{shippingZoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingZoneId"+"}", url.PathEscape(parameterToString(r.shippingZoneId, "")), -1)

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

type ApiPATCHShippingZonesShippingZoneIdRequest struct {
	ctx                context.Context
	ApiService         *ShippingZonesApiService
	shippingZoneUpdate *ShippingZoneUpdate
	shippingZoneId     string
}

func (r ApiPATCHShippingZonesShippingZoneIdRequest) ShippingZoneUpdate(shippingZoneUpdate ShippingZoneUpdate) ApiPATCHShippingZonesShippingZoneIdRequest {
	r.shippingZoneUpdate = &shippingZoneUpdate
	return r
}

func (r ApiPATCHShippingZonesShippingZoneIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHShippingZonesShippingZoneIdExecute(r)
}

/*
PATCHShippingZonesShippingZoneId Update a shipping zone

Update a shipping zone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shippingZoneId The resource's id
 @return ApiPATCHShippingZonesShippingZoneIdRequest
*/
func (a *ShippingZonesApiService) PATCHShippingZonesShippingZoneId(ctx context.Context, shippingZoneId string) ApiPATCHShippingZonesShippingZoneIdRequest {
	return ApiPATCHShippingZonesShippingZoneIdRequest{
		ApiService:     a,
		ctx:            ctx,
		shippingZoneId: shippingZoneId,
	}
}

// Execute executes the request
func (a *ShippingZonesApiService) PATCHShippingZonesShippingZoneIdExecute(r ApiPATCHShippingZonesShippingZoneIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.PATCHShippingZonesShippingZoneId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones/{shippingZoneId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shippingZoneId"+"}", url.PathEscape(parameterToString(r.shippingZoneId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingZoneUpdate == nil {
		return nil, reportError("shippingZoneUpdate is required and must be specified")
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
	localVarPostBody = r.shippingZoneUpdate
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

type ApiPOSTShippingZonesRequest struct {
	ctx                context.Context
	ApiService         *ShippingZonesApiService
	shippingZoneCreate *ShippingZoneCreate
}

func (r ApiPOSTShippingZonesRequest) ShippingZoneCreate(shippingZoneCreate ShippingZoneCreate) ApiPOSTShippingZonesRequest {
	r.shippingZoneCreate = &shippingZoneCreate
	return r
}

func (r ApiPOSTShippingZonesRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTShippingZonesExecute(r)
}

/*
POSTShippingZones Create a shipping zone

Create a shipping zone

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTShippingZonesRequest
*/
func (a *ShippingZonesApiService) POSTShippingZones(ctx context.Context) ApiPOSTShippingZonesRequest {
	return ApiPOSTShippingZonesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ShippingZonesApiService) POSTShippingZonesExecute(r ApiPOSTShippingZonesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShippingZonesApiService.POSTShippingZones")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/shipping_zones"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shippingZoneCreate == nil {
		return nil, reportError("shippingZoneCreate is required and must be specified")
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
	localVarPostBody = r.shippingZoneCreate
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
