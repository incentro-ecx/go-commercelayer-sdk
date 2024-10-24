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

// AxervePaymentsApiService AxervePaymentsApi service
type AxervePaymentsApiService service

type AxervePaymentsApiDELETEAxervePaymentsAxervePaymentIdRequest struct {
	ctx             context.Context
	ApiService      *AxervePaymentsApiService
	axervePaymentId interface{}
}

func (r AxervePaymentsApiDELETEAxervePaymentsAxervePaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEAxervePaymentsAxervePaymentIdExecute(r)
}

/*
DELETEAxervePaymentsAxervePaymentId Delete an axerve payment

Delete an axerve payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param axervePaymentId The resource's id
	@return AxervePaymentsApiDELETEAxervePaymentsAxervePaymentIdRequest
*/
func (a *AxervePaymentsApiService) DELETEAxervePaymentsAxervePaymentId(ctx context.Context, axervePaymentId interface{}) AxervePaymentsApiDELETEAxervePaymentsAxervePaymentIdRequest {
	return AxervePaymentsApiDELETEAxervePaymentsAxervePaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		axervePaymentId: axervePaymentId,
	}
}

// Execute executes the request
func (a *AxervePaymentsApiService) DELETEAxervePaymentsAxervePaymentIdExecute(r AxervePaymentsApiDELETEAxervePaymentsAxervePaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxervePaymentsApiService.DELETEAxervePaymentsAxervePaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_payments/{axervePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"axervePaymentId"+"}", url.PathEscape(parameterValueToString(r.axervePaymentId, "axervePaymentId")), -1)

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

type AxervePaymentsApiGETAxerveGatewayIdAxervePaymentsRequest struct {
	ctx             context.Context
	ApiService      *AxervePaymentsApiService
	axerveGatewayId interface{}
}

func (r AxervePaymentsApiGETAxerveGatewayIdAxervePaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETAxerveGatewayIdAxervePaymentsExecute(r)
}

/*
GETAxerveGatewayIdAxervePayments Retrieve the axerve payments associated to the axerve gateway

Retrieve the axerve payments associated to the axerve gateway

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param axerveGatewayId The resource's id
	@return AxervePaymentsApiGETAxerveGatewayIdAxervePaymentsRequest
*/
func (a *AxervePaymentsApiService) GETAxerveGatewayIdAxervePayments(ctx context.Context, axerveGatewayId interface{}) AxervePaymentsApiGETAxerveGatewayIdAxervePaymentsRequest {
	return AxervePaymentsApiGETAxerveGatewayIdAxervePaymentsRequest{
		ApiService:      a,
		ctx:             ctx,
		axerveGatewayId: axerveGatewayId,
	}
}

// Execute executes the request
func (a *AxervePaymentsApiService) GETAxerveGatewayIdAxervePaymentsExecute(r AxervePaymentsApiGETAxerveGatewayIdAxervePaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxervePaymentsApiService.GETAxerveGatewayIdAxervePayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_gateways/{axerveGatewayId}/axerve_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"axerveGatewayId"+"}", url.PathEscape(parameterValueToString(r.axerveGatewayId, "axerveGatewayId")), -1)

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

type AxervePaymentsApiGETAxervePaymentsRequest struct {
	ctx        context.Context
	ApiService *AxervePaymentsApiService
}

func (r AxervePaymentsApiGETAxervePaymentsRequest) Execute() (*GETAxervePayments200Response, *http.Response, error) {
	return r.ApiService.GETAxervePaymentsExecute(r)
}

/*
GETAxervePayments List all axerve payments

List all axerve payments

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AxervePaymentsApiGETAxervePaymentsRequest
*/
func (a *AxervePaymentsApiService) GETAxervePayments(ctx context.Context) AxervePaymentsApiGETAxervePaymentsRequest {
	return AxervePaymentsApiGETAxervePaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETAxervePayments200Response
func (a *AxervePaymentsApiService) GETAxervePaymentsExecute(r AxervePaymentsApiGETAxervePaymentsRequest) (*GETAxervePayments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAxervePayments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxervePaymentsApiService.GETAxervePayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_payments"

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

type AxervePaymentsApiGETAxervePaymentsAxervePaymentIdRequest struct {
	ctx             context.Context
	ApiService      *AxervePaymentsApiService
	axervePaymentId interface{}
}

func (r AxervePaymentsApiGETAxervePaymentsAxervePaymentIdRequest) Execute() (*GETAxervePaymentsAxervePaymentId200Response, *http.Response, error) {
	return r.ApiService.GETAxervePaymentsAxervePaymentIdExecute(r)
}

/*
GETAxervePaymentsAxervePaymentId Retrieve an axerve payment

Retrieve an axerve payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param axervePaymentId The resource's id
	@return AxervePaymentsApiGETAxervePaymentsAxervePaymentIdRequest
*/
func (a *AxervePaymentsApiService) GETAxervePaymentsAxervePaymentId(ctx context.Context, axervePaymentId interface{}) AxervePaymentsApiGETAxervePaymentsAxervePaymentIdRequest {
	return AxervePaymentsApiGETAxervePaymentsAxervePaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		axervePaymentId: axervePaymentId,
	}
}

// Execute executes the request
//
//	@return GETAxervePaymentsAxervePaymentId200Response
func (a *AxervePaymentsApiService) GETAxervePaymentsAxervePaymentIdExecute(r AxervePaymentsApiGETAxervePaymentsAxervePaymentIdRequest) (*GETAxervePaymentsAxervePaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETAxervePaymentsAxervePaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxervePaymentsApiService.GETAxervePaymentsAxervePaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_payments/{axervePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"axervePaymentId"+"}", url.PathEscape(parameterValueToString(r.axervePaymentId, "axervePaymentId")), -1)

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

type AxervePaymentsApiPATCHAxervePaymentsAxervePaymentIdRequest struct {
	ctx                 context.Context
	ApiService          *AxervePaymentsApiService
	axervePaymentUpdate *AxervePaymentUpdate
	axervePaymentId     interface{}
}

func (r AxervePaymentsApiPATCHAxervePaymentsAxervePaymentIdRequest) AxervePaymentUpdate(axervePaymentUpdate AxervePaymentUpdate) AxervePaymentsApiPATCHAxervePaymentsAxervePaymentIdRequest {
	r.axervePaymentUpdate = &axervePaymentUpdate
	return r
}

func (r AxervePaymentsApiPATCHAxervePaymentsAxervePaymentIdRequest) Execute() (*PATCHAxervePaymentsAxervePaymentId200Response, *http.Response, error) {
	return r.ApiService.PATCHAxervePaymentsAxervePaymentIdExecute(r)
}

/*
PATCHAxervePaymentsAxervePaymentId Update an axerve payment

Update an axerve payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param axervePaymentId The resource's id
	@return AxervePaymentsApiPATCHAxervePaymentsAxervePaymentIdRequest
*/
func (a *AxervePaymentsApiService) PATCHAxervePaymentsAxervePaymentId(ctx context.Context, axervePaymentId interface{}) AxervePaymentsApiPATCHAxervePaymentsAxervePaymentIdRequest {
	return AxervePaymentsApiPATCHAxervePaymentsAxervePaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		axervePaymentId: axervePaymentId,
	}
}

// Execute executes the request
//
//	@return PATCHAxervePaymentsAxervePaymentId200Response
func (a *AxervePaymentsApiService) PATCHAxervePaymentsAxervePaymentIdExecute(r AxervePaymentsApiPATCHAxervePaymentsAxervePaymentIdRequest) (*PATCHAxervePaymentsAxervePaymentId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHAxervePaymentsAxervePaymentId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxervePaymentsApiService.PATCHAxervePaymentsAxervePaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_payments/{axervePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"axervePaymentId"+"}", url.PathEscape(parameterValueToString(r.axervePaymentId, "axervePaymentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.axervePaymentUpdate == nil {
		return localVarReturnValue, nil, reportError("axervePaymentUpdate is required and must be specified")
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
	localVarPostBody = r.axervePaymentUpdate
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

type AxervePaymentsApiPOSTAxervePaymentsRequest struct {
	ctx                 context.Context
	ApiService          *AxervePaymentsApiService
	axervePaymentCreate *AxervePaymentCreate
}

func (r AxervePaymentsApiPOSTAxervePaymentsRequest) AxervePaymentCreate(axervePaymentCreate AxervePaymentCreate) AxervePaymentsApiPOSTAxervePaymentsRequest {
	r.axervePaymentCreate = &axervePaymentCreate
	return r
}

func (r AxervePaymentsApiPOSTAxervePaymentsRequest) Execute() (*POSTAxervePayments201Response, *http.Response, error) {
	return r.ApiService.POSTAxervePaymentsExecute(r)
}

/*
POSTAxervePayments Create an axerve payment

Create an axerve payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AxervePaymentsApiPOSTAxervePaymentsRequest
*/
func (a *AxervePaymentsApiService) POSTAxervePayments(ctx context.Context) AxervePaymentsApiPOSTAxervePaymentsRequest {
	return AxervePaymentsApiPOSTAxervePaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTAxervePayments201Response
func (a *AxervePaymentsApiService) POSTAxervePaymentsExecute(r AxervePaymentsApiPOSTAxervePaymentsRequest) (*POSTAxervePayments201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTAxervePayments201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AxervePaymentsApiService.POSTAxervePayments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/axerve_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.axervePaymentCreate == nil {
		return localVarReturnValue, nil, reportError("axervePaymentCreate is required and must be specified")
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
	localVarPostBody = r.axervePaymentCreate
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
