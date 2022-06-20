/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
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

type BingGeocodersApi interface {

	/*
		DELETEBingGeocodersBingGeocoderId Delete a bing geocoder

		Delete a bing geocoder

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param bingGeocoderId The resource's id
		@return ApiDELETEBingGeocodersBingGeocoderIdRequest
	*/
	DELETEBingGeocodersBingGeocoderId(ctx context.Context, bingGeocoderId string) ApiDELETEBingGeocodersBingGeocoderIdRequest

	// DELETEBingGeocodersBingGeocoderIdExecute executes the request
	DELETEBingGeocodersBingGeocoderIdExecute(r ApiDELETEBingGeocodersBingGeocoderIdRequest) (*http.Response, error)

	/*
		GETBingGeocoders List all bing geocoders

		List all bing geocoders

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGETBingGeocodersRequest
	*/
	GETBingGeocoders(ctx context.Context) ApiGETBingGeocodersRequest

	// GETBingGeocodersExecute executes the request
	GETBingGeocodersExecute(r ApiGETBingGeocodersRequest) (*http.Response, error)

	/*
		GETBingGeocodersBingGeocoderId Retrieve a bing geocoder

		Retrieve a bing geocoder

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param bingGeocoderId The resource's id
		@return ApiGETBingGeocodersBingGeocoderIdRequest
	*/
	GETBingGeocodersBingGeocoderId(ctx context.Context, bingGeocoderId string) ApiGETBingGeocodersBingGeocoderIdRequest

	// GETBingGeocodersBingGeocoderIdExecute executes the request
	//  @return BingGeocoder
	GETBingGeocodersBingGeocoderIdExecute(r ApiGETBingGeocodersBingGeocoderIdRequest) (*BingGeocoder, *http.Response, error)

	/*
		PATCHBingGeocodersBingGeocoderId Update a bing geocoder

		Update a bing geocoder

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param bingGeocoderId The resource's id
		@return ApiPATCHBingGeocodersBingGeocoderIdRequest
	*/
	PATCHBingGeocodersBingGeocoderId(ctx context.Context, bingGeocoderId string) ApiPATCHBingGeocodersBingGeocoderIdRequest

	// PATCHBingGeocodersBingGeocoderIdExecute executes the request
	PATCHBingGeocodersBingGeocoderIdExecute(r ApiPATCHBingGeocodersBingGeocoderIdRequest) (*http.Response, error)

	/*
		POSTBingGeocoders Create a bing geocoder

		Create a bing geocoder

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiPOSTBingGeocodersRequest
	*/
	POSTBingGeocoders(ctx context.Context) ApiPOSTBingGeocodersRequest

	// POSTBingGeocodersExecute executes the request
	POSTBingGeocodersExecute(r ApiPOSTBingGeocodersRequest) (*http.Response, error)
}

// BingGeocodersApiService BingGeocodersApi service
type BingGeocodersApiService service

type ApiDELETEBingGeocodersBingGeocoderIdRequest struct {
	ctx            context.Context
	ApiService     BingGeocodersApi
	bingGeocoderId string
}

func (r ApiDELETEBingGeocodersBingGeocoderIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEBingGeocodersBingGeocoderIdExecute(r)
}

/*
DELETEBingGeocodersBingGeocoderId Delete a bing geocoder

Delete a bing geocoder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bingGeocoderId The resource's id
 @return ApiDELETEBingGeocodersBingGeocoderIdRequest
*/
func (a *BingGeocodersApiService) DELETEBingGeocodersBingGeocoderId(ctx context.Context, bingGeocoderId string) ApiDELETEBingGeocodersBingGeocoderIdRequest {
	return ApiDELETEBingGeocodersBingGeocoderIdRequest{
		ApiService:     a,
		ctx:            ctx,
		bingGeocoderId: bingGeocoderId,
	}
}

// Execute executes the request
func (a *BingGeocodersApiService) DELETEBingGeocodersBingGeocoderIdExecute(r ApiDELETEBingGeocodersBingGeocoderIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.DELETEBingGeocodersBingGeocoderId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders/{bingGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bingGeocoderId"+"}", url.PathEscape(parameterToString(r.bingGeocoderId, "")), -1)

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

type ApiGETBingGeocodersRequest struct {
	ctx        context.Context
	ApiService BingGeocodersApi
}

func (r ApiGETBingGeocodersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETBingGeocodersExecute(r)
}

/*
GETBingGeocoders List all bing geocoders

List all bing geocoders

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETBingGeocodersRequest
*/
func (a *BingGeocodersApiService) GETBingGeocoders(ctx context.Context) ApiGETBingGeocodersRequest {
	return ApiGETBingGeocodersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BingGeocodersApiService) GETBingGeocodersExecute(r ApiGETBingGeocodersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.GETBingGeocoders")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders"

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

type ApiGETBingGeocodersBingGeocoderIdRequest struct {
	ctx            context.Context
	ApiService     BingGeocodersApi
	bingGeocoderId string
}

func (r ApiGETBingGeocodersBingGeocoderIdRequest) Execute() (*BingGeocoder, *http.Response, error) {
	return r.ApiService.GETBingGeocodersBingGeocoderIdExecute(r)
}

/*
GETBingGeocodersBingGeocoderId Retrieve a bing geocoder

Retrieve a bing geocoder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bingGeocoderId The resource's id
 @return ApiGETBingGeocodersBingGeocoderIdRequest
*/
func (a *BingGeocodersApiService) GETBingGeocodersBingGeocoderId(ctx context.Context, bingGeocoderId string) ApiGETBingGeocodersBingGeocoderIdRequest {
	return ApiGETBingGeocodersBingGeocoderIdRequest{
		ApiService:     a,
		ctx:            ctx,
		bingGeocoderId: bingGeocoderId,
	}
}

// Execute executes the request
//  @return BingGeocoder
func (a *BingGeocodersApiService) GETBingGeocodersBingGeocoderIdExecute(r ApiGETBingGeocodersBingGeocoderIdRequest) (*BingGeocoder, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BingGeocoder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.GETBingGeocodersBingGeocoderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders/{bingGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bingGeocoderId"+"}", url.PathEscape(parameterToString(r.bingGeocoderId, "")), -1)

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

type ApiPATCHBingGeocodersBingGeocoderIdRequest struct {
	ctx                context.Context
	ApiService         BingGeocodersApi
	bingGeocoderId     string
	bingGeocoderUpdate *BingGeocoderUpdate
}

func (r ApiPATCHBingGeocodersBingGeocoderIdRequest) BingGeocoderUpdate(bingGeocoderUpdate BingGeocoderUpdate) ApiPATCHBingGeocodersBingGeocoderIdRequest {
	r.bingGeocoderUpdate = &bingGeocoderUpdate
	return r
}

func (r ApiPATCHBingGeocodersBingGeocoderIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHBingGeocodersBingGeocoderIdExecute(r)
}

/*
PATCHBingGeocodersBingGeocoderId Update a bing geocoder

Update a bing geocoder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param bingGeocoderId The resource's id
 @return ApiPATCHBingGeocodersBingGeocoderIdRequest
*/
func (a *BingGeocodersApiService) PATCHBingGeocodersBingGeocoderId(ctx context.Context, bingGeocoderId string) ApiPATCHBingGeocodersBingGeocoderIdRequest {
	return ApiPATCHBingGeocodersBingGeocoderIdRequest{
		ApiService:     a,
		ctx:            ctx,
		bingGeocoderId: bingGeocoderId,
	}
}

// Execute executes the request
func (a *BingGeocodersApiService) PATCHBingGeocodersBingGeocoderIdExecute(r ApiPATCHBingGeocodersBingGeocoderIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.PATCHBingGeocodersBingGeocoderId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders/{bingGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"bingGeocoderId"+"}", url.PathEscape(parameterToString(r.bingGeocoderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bingGeocoderUpdate == nil {
		return nil, reportError("bingGeocoderUpdate is required and must be specified")
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
	localVarPostBody = r.bingGeocoderUpdate
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

type ApiPOSTBingGeocodersRequest struct {
	ctx                context.Context
	ApiService         BingGeocodersApi
	bingGeocoderCreate *BingGeocoderCreate
}

func (r ApiPOSTBingGeocodersRequest) BingGeocoderCreate(bingGeocoderCreate BingGeocoderCreate) ApiPOSTBingGeocodersRequest {
	r.bingGeocoderCreate = &bingGeocoderCreate
	return r
}

func (r ApiPOSTBingGeocodersRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTBingGeocodersExecute(r)
}

/*
POSTBingGeocoders Create a bing geocoder

Create a bing geocoder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTBingGeocodersRequest
*/
func (a *BingGeocodersApiService) POSTBingGeocoders(ctx context.Context) ApiPOSTBingGeocodersRequest {
	return ApiPOSTBingGeocodersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *BingGeocodersApiService) POSTBingGeocodersExecute(r ApiPOSTBingGeocodersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BingGeocodersApiService.POSTBingGeocoders")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bing_geocoders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bingGeocoderCreate == nil {
		return nil, reportError("bingGeocoderCreate is required and must be specified")
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
	localVarPostBody = r.bingGeocoderCreate
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