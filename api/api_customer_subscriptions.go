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

// CustomerSubscriptionsApiService CustomerSubscriptionsApi service
type CustomerSubscriptionsApiService service

type CustomerSubscriptionsApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest struct {
	ctx                    context.Context
	ApiService             *CustomerSubscriptionsApiService
	customerSubscriptionId string
}

func (r CustomerSubscriptionsApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECustomerSubscriptionsCustomerSubscriptionIdExecute(r)
}

/*
DELETECustomerSubscriptionsCustomerSubscriptionId Delete a customer subscription

Delete a customer subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerSubscriptionId The resource's id
	@return CustomerSubscriptionsApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest
*/
func (a *CustomerSubscriptionsApiService) DELETECustomerSubscriptionsCustomerSubscriptionId(ctx context.Context, customerSubscriptionId string) CustomerSubscriptionsApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest {
	return CustomerSubscriptionsApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest{
		ApiService:             a,
		ctx:                    ctx,
		customerSubscriptionId: customerSubscriptionId,
	}
}

// Execute executes the request
func (a *CustomerSubscriptionsApiService) DELETECustomerSubscriptionsCustomerSubscriptionIdExecute(r CustomerSubscriptionsApiDELETECustomerSubscriptionsCustomerSubscriptionIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.DELETECustomerSubscriptionsCustomerSubscriptionId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions/{customerSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerSubscriptionId"+"}", url.PathEscape(parameterToString(r.customerSubscriptionId, "")), -1)

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

type CustomerSubscriptionsApiGETCustomerIdCustomerSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *CustomerSubscriptionsApiService
	customerId string
}

func (r CustomerSubscriptionsApiGETCustomerIdCustomerSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerIdCustomerSubscriptionsExecute(r)
}

/*
GETCustomerIdCustomerSubscriptions Retrieve the customer subscriptions associated to the customer

Retrieve the customer subscriptions associated to the customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerId The resource's id
	@return CustomerSubscriptionsApiGETCustomerIdCustomerSubscriptionsRequest
*/
func (a *CustomerSubscriptionsApiService) GETCustomerIdCustomerSubscriptions(ctx context.Context, customerId string) CustomerSubscriptionsApiGETCustomerIdCustomerSubscriptionsRequest {
	return CustomerSubscriptionsApiGETCustomerIdCustomerSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
	}
}

// Execute executes the request
func (a *CustomerSubscriptionsApiService) GETCustomerIdCustomerSubscriptionsExecute(r CustomerSubscriptionsApiGETCustomerIdCustomerSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.GETCustomerIdCustomerSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/customer_subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterToString(r.customerId, "")), -1)

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

type CustomerSubscriptionsApiGETCustomerSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *CustomerSubscriptionsApiService
}

func (r CustomerSubscriptionsApiGETCustomerSubscriptionsRequest) Execute() (*GETCustomerSubscriptions200Response, *http.Response, error) {
	return r.ApiService.GETCustomerSubscriptionsExecute(r)
}

/*
GETCustomerSubscriptions List all customer subscriptions

List all customer subscriptions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CustomerSubscriptionsApiGETCustomerSubscriptionsRequest
*/
func (a *CustomerSubscriptionsApiService) GETCustomerSubscriptions(ctx context.Context) CustomerSubscriptionsApiGETCustomerSubscriptionsRequest {
	return CustomerSubscriptionsApiGETCustomerSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETCustomerSubscriptions200Response
func (a *CustomerSubscriptionsApiService) GETCustomerSubscriptionsExecute(r CustomerSubscriptionsApiGETCustomerSubscriptionsRequest) (*GETCustomerSubscriptions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCustomerSubscriptions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.GETCustomerSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions"

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

type CustomerSubscriptionsApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest struct {
	ctx                    context.Context
	ApiService             *CustomerSubscriptionsApiService
	customerSubscriptionId string
}

func (r CustomerSubscriptionsApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest) Execute() (*GETCustomerSubscriptionsCustomerSubscriptionId200Response, *http.Response, error) {
	return r.ApiService.GETCustomerSubscriptionsCustomerSubscriptionIdExecute(r)
}

/*
GETCustomerSubscriptionsCustomerSubscriptionId Retrieve a customer subscription

Retrieve a customer subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerSubscriptionId The resource's id
	@return CustomerSubscriptionsApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest
*/
func (a *CustomerSubscriptionsApiService) GETCustomerSubscriptionsCustomerSubscriptionId(ctx context.Context, customerSubscriptionId string) CustomerSubscriptionsApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest {
	return CustomerSubscriptionsApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest{
		ApiService:             a,
		ctx:                    ctx,
		customerSubscriptionId: customerSubscriptionId,
	}
}

// Execute executes the request
//
//	@return GETCustomerSubscriptionsCustomerSubscriptionId200Response
func (a *CustomerSubscriptionsApiService) GETCustomerSubscriptionsCustomerSubscriptionIdExecute(r CustomerSubscriptionsApiGETCustomerSubscriptionsCustomerSubscriptionIdRequest) (*GETCustomerSubscriptionsCustomerSubscriptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCustomerSubscriptionsCustomerSubscriptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.GETCustomerSubscriptionsCustomerSubscriptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions/{customerSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerSubscriptionId"+"}", url.PathEscape(parameterToString(r.customerSubscriptionId, "")), -1)

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

type CustomerSubscriptionsApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest struct {
	ctx                        context.Context
	ApiService                 *CustomerSubscriptionsApiService
	customerSubscriptionUpdate *CustomerSubscriptionUpdate
	customerSubscriptionId     string
}

func (r CustomerSubscriptionsApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest) CustomerSubscriptionUpdate(customerSubscriptionUpdate CustomerSubscriptionUpdate) CustomerSubscriptionsApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest {
	r.customerSubscriptionUpdate = &customerSubscriptionUpdate
	return r
}

func (r CustomerSubscriptionsApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest) Execute() (*PATCHCustomerSubscriptionsCustomerSubscriptionId200Response, *http.Response, error) {
	return r.ApiService.PATCHCustomerSubscriptionsCustomerSubscriptionIdExecute(r)
}

/*
PATCHCustomerSubscriptionsCustomerSubscriptionId Update a customer subscription

Update a customer subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerSubscriptionId The resource's id
	@return CustomerSubscriptionsApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest
*/
func (a *CustomerSubscriptionsApiService) PATCHCustomerSubscriptionsCustomerSubscriptionId(ctx context.Context, customerSubscriptionId string) CustomerSubscriptionsApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest {
	return CustomerSubscriptionsApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest{
		ApiService:             a,
		ctx:                    ctx,
		customerSubscriptionId: customerSubscriptionId,
	}
}

// Execute executes the request
//
//	@return PATCHCustomerSubscriptionsCustomerSubscriptionId200Response
func (a *CustomerSubscriptionsApiService) PATCHCustomerSubscriptionsCustomerSubscriptionIdExecute(r CustomerSubscriptionsApiPATCHCustomerSubscriptionsCustomerSubscriptionIdRequest) (*PATCHCustomerSubscriptionsCustomerSubscriptionId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHCustomerSubscriptionsCustomerSubscriptionId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.PATCHCustomerSubscriptionsCustomerSubscriptionId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions/{customerSubscriptionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerSubscriptionId"+"}", url.PathEscape(parameterToString(r.customerSubscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerSubscriptionUpdate == nil {
		return localVarReturnValue, nil, reportError("customerSubscriptionUpdate is required and must be specified")
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
	localVarPostBody = r.customerSubscriptionUpdate
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

type CustomerSubscriptionsApiPOSTCustomerSubscriptionsRequest struct {
	ctx                        context.Context
	ApiService                 *CustomerSubscriptionsApiService
	customerSubscriptionCreate *CustomerSubscriptionCreate
}

func (r CustomerSubscriptionsApiPOSTCustomerSubscriptionsRequest) CustomerSubscriptionCreate(customerSubscriptionCreate CustomerSubscriptionCreate) CustomerSubscriptionsApiPOSTCustomerSubscriptionsRequest {
	r.customerSubscriptionCreate = &customerSubscriptionCreate
	return r
}

func (r CustomerSubscriptionsApiPOSTCustomerSubscriptionsRequest) Execute() (*POSTCustomerSubscriptions201Response, *http.Response, error) {
	return r.ApiService.POSTCustomerSubscriptionsExecute(r)
}

/*
POSTCustomerSubscriptions Create a customer subscription

Create a customer subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CustomerSubscriptionsApiPOSTCustomerSubscriptionsRequest
*/
func (a *CustomerSubscriptionsApiService) POSTCustomerSubscriptions(ctx context.Context) CustomerSubscriptionsApiPOSTCustomerSubscriptionsRequest {
	return CustomerSubscriptionsApiPOSTCustomerSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTCustomerSubscriptions201Response
func (a *CustomerSubscriptionsApiService) POSTCustomerSubscriptionsExecute(r CustomerSubscriptionsApiPOSTCustomerSubscriptionsRequest) (*POSTCustomerSubscriptions201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTCustomerSubscriptions201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerSubscriptionsApiService.POSTCustomerSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerSubscriptionCreate == nil {
		return localVarReturnValue, nil, reportError("customerSubscriptionCreate is required and must be specified")
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
	localVarPostBody = r.customerSubscriptionCreate
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
