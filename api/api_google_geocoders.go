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

// GoogleGeocodersApiService GoogleGeocodersApi service
type GoogleGeocodersApiService service

type ApiDELETEGoogleGeocodersGoogleGeocoderIdRequest struct {
	ctx              context.Context
	ApiService       *GoogleGeocodersApiService
	googleGeocoderId string
}

func (r ApiDELETEGoogleGeocodersGoogleGeocoderIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEGoogleGeocodersGoogleGeocoderIdExecute(r)
}

/*
DELETEGoogleGeocodersGoogleGeocoderId Delete a google geocoder

Delete a google geocoder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param googleGeocoderId The resource's id
 @return ApiDELETEGoogleGeocodersGoogleGeocoderIdRequest
*/
func (a *GoogleGeocodersApiService) DELETEGoogleGeocodersGoogleGeocoderId(ctx context.Context, googleGeocoderId string) ApiDELETEGoogleGeocodersGoogleGeocoderIdRequest {
	return ApiDELETEGoogleGeocodersGoogleGeocoderIdRequest{
		ApiService:       a,
		ctx:              ctx,
		googleGeocoderId: googleGeocoderId,
	}
}

// Execute executes the request
func (a *GoogleGeocodersApiService) DELETEGoogleGeocodersGoogleGeocoderIdExecute(r ApiDELETEGoogleGeocodersGoogleGeocoderIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.DELETEGoogleGeocodersGoogleGeocoderId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders/{googleGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"googleGeocoderId"+"}", url.PathEscape(parameterToString(r.googleGeocoderId, "")), -1)

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

type ApiGETGoogleGeocodersRequest struct {
	ctx        context.Context
	ApiService *GoogleGeocodersApiService
}

func (r ApiGETGoogleGeocodersRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETGoogleGeocodersExecute(r)
}

/*
GETGoogleGeocoders List all google geocoders

List all google geocoders

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETGoogleGeocodersRequest
*/
func (a *GoogleGeocodersApiService) GETGoogleGeocoders(ctx context.Context) ApiGETGoogleGeocodersRequest {
	return ApiGETGoogleGeocodersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *GoogleGeocodersApiService) GETGoogleGeocodersExecute(r ApiGETGoogleGeocodersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.GETGoogleGeocoders")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders"

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

type ApiGETGoogleGeocodersGoogleGeocoderIdRequest struct {
	ctx              context.Context
	ApiService       *GoogleGeocodersApiService
	googleGeocoderId string
}

func (r ApiGETGoogleGeocodersGoogleGeocoderIdRequest) Execute() (*GoogleGeocoder, *http.Response, error) {
	return r.ApiService.GETGoogleGeocodersGoogleGeocoderIdExecute(r)
}

/*
GETGoogleGeocodersGoogleGeocoderId Retrieve a google geocoder

Retrieve a google geocoder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param googleGeocoderId The resource's id
 @return ApiGETGoogleGeocodersGoogleGeocoderIdRequest
*/
func (a *GoogleGeocodersApiService) GETGoogleGeocodersGoogleGeocoderId(ctx context.Context, googleGeocoderId string) ApiGETGoogleGeocodersGoogleGeocoderIdRequest {
	return ApiGETGoogleGeocodersGoogleGeocoderIdRequest{
		ApiService:       a,
		ctx:              ctx,
		googleGeocoderId: googleGeocoderId,
	}
}

// Execute executes the request
//  @return GoogleGeocoder
func (a *GoogleGeocodersApiService) GETGoogleGeocodersGoogleGeocoderIdExecute(r ApiGETGoogleGeocodersGoogleGeocoderIdRequest) (*GoogleGeocoder, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GoogleGeocoder
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.GETGoogleGeocodersGoogleGeocoderId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders/{googleGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"googleGeocoderId"+"}", url.PathEscape(parameterToString(r.googleGeocoderId, "")), -1)

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

type ApiPATCHGoogleGeocodersGoogleGeocoderIdRequest struct {
	ctx                  context.Context
	ApiService           *GoogleGeocodersApiService
	googleGeocoderUpdate *GoogleGeocoderUpdate
	googleGeocoderId     string
}

func (r ApiPATCHGoogleGeocodersGoogleGeocoderIdRequest) GoogleGeocoderUpdate(googleGeocoderUpdate GoogleGeocoderUpdate) ApiPATCHGoogleGeocodersGoogleGeocoderIdRequest {
	r.googleGeocoderUpdate = &googleGeocoderUpdate
	return r
}

func (r ApiPATCHGoogleGeocodersGoogleGeocoderIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHGoogleGeocodersGoogleGeocoderIdExecute(r)
}

/*
PATCHGoogleGeocodersGoogleGeocoderId Update a google geocoder

Update a google geocoder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param googleGeocoderId The resource's id
 @return ApiPATCHGoogleGeocodersGoogleGeocoderIdRequest
*/
func (a *GoogleGeocodersApiService) PATCHGoogleGeocodersGoogleGeocoderId(ctx context.Context, googleGeocoderId string) ApiPATCHGoogleGeocodersGoogleGeocoderIdRequest {
	return ApiPATCHGoogleGeocodersGoogleGeocoderIdRequest{
		ApiService:       a,
		ctx:              ctx,
		googleGeocoderId: googleGeocoderId,
	}
}

// Execute executes the request
func (a *GoogleGeocodersApiService) PATCHGoogleGeocodersGoogleGeocoderIdExecute(r ApiPATCHGoogleGeocodersGoogleGeocoderIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.PATCHGoogleGeocodersGoogleGeocoderId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders/{googleGeocoderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"googleGeocoderId"+"}", url.PathEscape(parameterToString(r.googleGeocoderId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.googleGeocoderUpdate == nil {
		return nil, reportError("googleGeocoderUpdate is required and must be specified")
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
	localVarPostBody = r.googleGeocoderUpdate
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

type ApiPOSTGoogleGeocodersRequest struct {
	ctx                  context.Context
	ApiService           *GoogleGeocodersApiService
	googleGeocoderCreate *GoogleGeocoderCreate
}

func (r ApiPOSTGoogleGeocodersRequest) GoogleGeocoderCreate(googleGeocoderCreate GoogleGeocoderCreate) ApiPOSTGoogleGeocodersRequest {
	r.googleGeocoderCreate = &googleGeocoderCreate
	return r
}

func (r ApiPOSTGoogleGeocodersRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTGoogleGeocodersExecute(r)
}

/*
POSTGoogleGeocoders Create a google geocoder

Create a google geocoder

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTGoogleGeocodersRequest
*/
func (a *GoogleGeocodersApiService) POSTGoogleGeocoders(ctx context.Context) ApiPOSTGoogleGeocodersRequest {
	return ApiPOSTGoogleGeocodersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *GoogleGeocodersApiService) POSTGoogleGeocodersExecute(r ApiPOSTGoogleGeocodersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoogleGeocodersApiService.POSTGoogleGeocoders")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/google_geocoders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.googleGeocoderCreate == nil {
		return nil, reportError("googleGeocoderCreate is required and must be specified")
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
	localVarPostBody = r.googleGeocoderCreate
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
