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

// StripePaymentsApiService StripePaymentsApi service
type StripePaymentsApiService service

type ApiDELETEStripePaymentsStripePaymentIdRequest struct {
	ctx             context.Context
	ApiService      *StripePaymentsApiService
	stripePaymentId string
}

func (r ApiDELETEStripePaymentsStripePaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEStripePaymentsStripePaymentIdExecute(r)
}

/*
DELETEStripePaymentsStripePaymentId Delete a stripe payment

Delete a stripe payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stripePaymentId The resource's id
 @return ApiDELETEStripePaymentsStripePaymentIdRequest
*/
func (a *StripePaymentsApiService) DELETEStripePaymentsStripePaymentId(ctx context.Context, stripePaymentId string) ApiDELETEStripePaymentsStripePaymentIdRequest {
	return ApiDELETEStripePaymentsStripePaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stripePaymentId: stripePaymentId,
	}
}

// Execute executes the request
func (a *StripePaymentsApiService) DELETEStripePaymentsStripePaymentIdExecute(r ApiDELETEStripePaymentsStripePaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripePaymentsApiService.DELETEStripePaymentsStripePaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_payments/{stripePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stripePaymentId"+"}", url.PathEscape(parameterToString(r.stripePaymentId, "")), -1)

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

type ApiGETStripeGatewayIdStripePaymentsRequest struct {
	ctx             context.Context
	ApiService      *StripePaymentsApiService
	stripeGatewayId string
}

func (r ApiGETStripeGatewayIdStripePaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStripeGatewayIdStripePaymentsExecute(r)
}

/*
GETStripeGatewayIdStripePayments Retrieve the stripe payments associated to the stripe gateway

Retrieve the stripe payments associated to the stripe gateway

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stripeGatewayId The resource's id
 @return ApiGETStripeGatewayIdStripePaymentsRequest
*/
func (a *StripePaymentsApiService) GETStripeGatewayIdStripePayments(ctx context.Context, stripeGatewayId string) ApiGETStripeGatewayIdStripePaymentsRequest {
	return ApiGETStripeGatewayIdStripePaymentsRequest{
		ApiService:      a,
		ctx:             ctx,
		stripeGatewayId: stripeGatewayId,
	}
}

// Execute executes the request
func (a *StripePaymentsApiService) GETStripeGatewayIdStripePaymentsExecute(r ApiGETStripeGatewayIdStripePaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripePaymentsApiService.GETStripeGatewayIdStripePayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_gateways/{stripeGatewayId}/stripe_payments"
	localVarPath = strings.Replace(localVarPath, "{"+"stripeGatewayId"+"}", url.PathEscape(parameterToString(r.stripeGatewayId, "")), -1)

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

type ApiGETStripePaymentsRequest struct {
	ctx        context.Context
	ApiService *StripePaymentsApiService
}

func (r ApiGETStripePaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStripePaymentsExecute(r)
}

/*
GETStripePayments List all stripe payments

List all stripe payments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETStripePaymentsRequest
*/
func (a *StripePaymentsApiService) GETStripePayments(ctx context.Context) ApiGETStripePaymentsRequest {
	return ApiGETStripePaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *StripePaymentsApiService) GETStripePaymentsExecute(r ApiGETStripePaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripePaymentsApiService.GETStripePayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_payments"

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

type ApiGETStripePaymentsStripePaymentIdRequest struct {
	ctx             context.Context
	ApiService      *StripePaymentsApiService
	stripePaymentId string
}

func (r ApiGETStripePaymentsStripePaymentIdRequest) Execute() (*StripePayment, *http.Response, error) {
	return r.ApiService.GETStripePaymentsStripePaymentIdExecute(r)
}

/*
GETStripePaymentsStripePaymentId Retrieve a stripe payment

Retrieve a stripe payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stripePaymentId The resource's id
 @return ApiGETStripePaymentsStripePaymentIdRequest
*/
func (a *StripePaymentsApiService) GETStripePaymentsStripePaymentId(ctx context.Context, stripePaymentId string) ApiGETStripePaymentsStripePaymentIdRequest {
	return ApiGETStripePaymentsStripePaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stripePaymentId: stripePaymentId,
	}
}

// Execute executes the request
//  @return StripePayment
func (a *StripePaymentsApiService) GETStripePaymentsStripePaymentIdExecute(r ApiGETStripePaymentsStripePaymentIdRequest) (*StripePayment, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *StripePayment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripePaymentsApiService.GETStripePaymentsStripePaymentId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_payments/{stripePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stripePaymentId"+"}", url.PathEscape(parameterToString(r.stripePaymentId, "")), -1)

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

type ApiPATCHStripePaymentsStripePaymentIdRequest struct {
	ctx                 context.Context
	ApiService          *StripePaymentsApiService
	stripePaymentUpdate *StripePaymentUpdate
	stripePaymentId     string
}

func (r ApiPATCHStripePaymentsStripePaymentIdRequest) StripePaymentUpdate(stripePaymentUpdate StripePaymentUpdate) ApiPATCHStripePaymentsStripePaymentIdRequest {
	r.stripePaymentUpdate = &stripePaymentUpdate
	return r
}

func (r ApiPATCHStripePaymentsStripePaymentIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHStripePaymentsStripePaymentIdExecute(r)
}

/*
PATCHStripePaymentsStripePaymentId Update a stripe payment

Update a stripe payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stripePaymentId The resource's id
 @return ApiPATCHStripePaymentsStripePaymentIdRequest
*/
func (a *StripePaymentsApiService) PATCHStripePaymentsStripePaymentId(ctx context.Context, stripePaymentId string) ApiPATCHStripePaymentsStripePaymentIdRequest {
	return ApiPATCHStripePaymentsStripePaymentIdRequest{
		ApiService:      a,
		ctx:             ctx,
		stripePaymentId: stripePaymentId,
	}
}

// Execute executes the request
func (a *StripePaymentsApiService) PATCHStripePaymentsStripePaymentIdExecute(r ApiPATCHStripePaymentsStripePaymentIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripePaymentsApiService.PATCHStripePaymentsStripePaymentId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_payments/{stripePaymentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"stripePaymentId"+"}", url.PathEscape(parameterToString(r.stripePaymentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stripePaymentUpdate == nil {
		return nil, reportError("stripePaymentUpdate is required and must be specified")
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
	localVarPostBody = r.stripePaymentUpdate
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

type ApiPOSTStripePaymentsRequest struct {
	ctx                 context.Context
	ApiService          *StripePaymentsApiService
	stripePaymentCreate *StripePaymentCreate
}

func (r ApiPOSTStripePaymentsRequest) StripePaymentCreate(stripePaymentCreate StripePaymentCreate) ApiPOSTStripePaymentsRequest {
	r.stripePaymentCreate = &stripePaymentCreate
	return r
}

func (r ApiPOSTStripePaymentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTStripePaymentsExecute(r)
}

/*
POSTStripePayments Create a stripe payment

Create a stripe payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTStripePaymentsRequest
*/
func (a *StripePaymentsApiService) POSTStripePayments(ctx context.Context) ApiPOSTStripePaymentsRequest {
	return ApiPOSTStripePaymentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *StripePaymentsApiService) POSTStripePaymentsExecute(r ApiPOSTStripePaymentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StripePaymentsApiService.POSTStripePayments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stripe_payments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.stripePaymentCreate == nil {
		return nil, reportError("stripePaymentCreate is required and must be specified")
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
	localVarPostBody = r.stripePaymentCreate
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
