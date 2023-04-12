/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
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

// CustomerAddressesApiService CustomerAddressesApi service
type CustomerAddressesApiService service

type CustomerAddressesApiDELETECustomerAddressesCustomerAddressIdRequest struct {
	ctx               context.Context
	ApiService        *CustomerAddressesApiService
	customerAddressId interface{}
}

func (r CustomerAddressesApiDELETECustomerAddressesCustomerAddressIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETECustomerAddressesCustomerAddressIdExecute(r)
}

/*
DELETECustomerAddressesCustomerAddressId Delete a customer address

Delete a customer address

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerAddressId The resource's id
	@return CustomerAddressesApiDELETECustomerAddressesCustomerAddressIdRequest
*/
func (a *CustomerAddressesApiService) DELETECustomerAddressesCustomerAddressId(ctx context.Context, customerAddressId interface{}) CustomerAddressesApiDELETECustomerAddressesCustomerAddressIdRequest {
	return CustomerAddressesApiDELETECustomerAddressesCustomerAddressIdRequest{
		ApiService:        a,
		ctx:               ctx,
		customerAddressId: customerAddressId,
	}
}

// Execute executes the request
func (a *CustomerAddressesApiService) DELETECustomerAddressesCustomerAddressIdExecute(r CustomerAddressesApiDELETECustomerAddressesCustomerAddressIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAddressesApiService.DELETECustomerAddressesCustomerAddressId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_addresses/{customerAddressId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerAddressId"+"}", url.PathEscape(parameterValueToString(r.customerAddressId, "customerAddressId")), -1)

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

type CustomerAddressesApiGETCustomerAddressesRequest struct {
	ctx        context.Context
	ApiService *CustomerAddressesApiService
}

func (r CustomerAddressesApiGETCustomerAddressesRequest) Execute() (*GETCustomerAddresses200Response, *http.Response, error) {
	return r.ApiService.GETCustomerAddressesExecute(r)
}

/*
GETCustomerAddresses List all customer addresses

List all customer addresses

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CustomerAddressesApiGETCustomerAddressesRequest
*/
func (a *CustomerAddressesApiService) GETCustomerAddresses(ctx context.Context) CustomerAddressesApiGETCustomerAddressesRequest {
	return CustomerAddressesApiGETCustomerAddressesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETCustomerAddresses200Response
func (a *CustomerAddressesApiService) GETCustomerAddressesExecute(r CustomerAddressesApiGETCustomerAddressesRequest) (*GETCustomerAddresses200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCustomerAddresses200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAddressesApiService.GETCustomerAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_addresses"

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

type CustomerAddressesApiGETCustomerAddressesCustomerAddressIdRequest struct {
	ctx               context.Context
	ApiService        *CustomerAddressesApiService
	customerAddressId interface{}
}

func (r CustomerAddressesApiGETCustomerAddressesCustomerAddressIdRequest) Execute() (*GETCustomerAddressesCustomerAddressId200Response, *http.Response, error) {
	return r.ApiService.GETCustomerAddressesCustomerAddressIdExecute(r)
}

/*
GETCustomerAddressesCustomerAddressId Retrieve a customer address

Retrieve a customer address

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerAddressId The resource's id
	@return CustomerAddressesApiGETCustomerAddressesCustomerAddressIdRequest
*/
func (a *CustomerAddressesApiService) GETCustomerAddressesCustomerAddressId(ctx context.Context, customerAddressId interface{}) CustomerAddressesApiGETCustomerAddressesCustomerAddressIdRequest {
	return CustomerAddressesApiGETCustomerAddressesCustomerAddressIdRequest{
		ApiService:        a,
		ctx:               ctx,
		customerAddressId: customerAddressId,
	}
}

// Execute executes the request
//
//	@return GETCustomerAddressesCustomerAddressId200Response
func (a *CustomerAddressesApiService) GETCustomerAddressesCustomerAddressIdExecute(r CustomerAddressesApiGETCustomerAddressesCustomerAddressIdRequest) (*GETCustomerAddressesCustomerAddressId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCustomerAddressesCustomerAddressId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAddressesApiService.GETCustomerAddressesCustomerAddressId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_addresses/{customerAddressId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerAddressId"+"}", url.PathEscape(parameterValueToString(r.customerAddressId, "customerAddressId")), -1)

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

type CustomerAddressesApiGETCustomerIdCustomerAddressesRequest struct {
	ctx        context.Context
	ApiService *CustomerAddressesApiService
	customerId interface{}
}

func (r CustomerAddressesApiGETCustomerIdCustomerAddressesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerIdCustomerAddressesExecute(r)
}

/*
GETCustomerIdCustomerAddresses Retrieve the customer addresses associated to the customer

Retrieve the customer addresses associated to the customer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerId The resource's id
	@return CustomerAddressesApiGETCustomerIdCustomerAddressesRequest
*/
func (a *CustomerAddressesApiService) GETCustomerIdCustomerAddresses(ctx context.Context, customerId interface{}) CustomerAddressesApiGETCustomerIdCustomerAddressesRequest {
	return CustomerAddressesApiGETCustomerIdCustomerAddressesRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
	}
}

// Execute executes the request
func (a *CustomerAddressesApiService) GETCustomerIdCustomerAddressesExecute(r CustomerAddressesApiGETCustomerIdCustomerAddressesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAddressesApiService.GETCustomerIdCustomerAddresses")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/customer_addresses"
	localVarPath = strings.Replace(localVarPath, "{"+"customerId"+"}", url.PathEscape(parameterValueToString(r.customerId, "customerId")), -1)

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

type CustomerAddressesApiPATCHCustomerAddressesCustomerAddressIdRequest struct {
	ctx                   context.Context
	ApiService            *CustomerAddressesApiService
	customerAddressUpdate *CustomerAddressUpdate
	customerAddressId     interface{}
}

func (r CustomerAddressesApiPATCHCustomerAddressesCustomerAddressIdRequest) CustomerAddressUpdate(customerAddressUpdate CustomerAddressUpdate) CustomerAddressesApiPATCHCustomerAddressesCustomerAddressIdRequest {
	r.customerAddressUpdate = &customerAddressUpdate
	return r
}

func (r CustomerAddressesApiPATCHCustomerAddressesCustomerAddressIdRequest) Execute() (*PATCHCustomerAddressesCustomerAddressId200Response, *http.Response, error) {
	return r.ApiService.PATCHCustomerAddressesCustomerAddressIdExecute(r)
}

/*
PATCHCustomerAddressesCustomerAddressId Update a customer address

Update a customer address

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param customerAddressId The resource's id
	@return CustomerAddressesApiPATCHCustomerAddressesCustomerAddressIdRequest
*/
func (a *CustomerAddressesApiService) PATCHCustomerAddressesCustomerAddressId(ctx context.Context, customerAddressId interface{}) CustomerAddressesApiPATCHCustomerAddressesCustomerAddressIdRequest {
	return CustomerAddressesApiPATCHCustomerAddressesCustomerAddressIdRequest{
		ApiService:        a,
		ctx:               ctx,
		customerAddressId: customerAddressId,
	}
}

// Execute executes the request
//
//	@return PATCHCustomerAddressesCustomerAddressId200Response
func (a *CustomerAddressesApiService) PATCHCustomerAddressesCustomerAddressIdExecute(r CustomerAddressesApiPATCHCustomerAddressesCustomerAddressIdRequest) (*PATCHCustomerAddressesCustomerAddressId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHCustomerAddressesCustomerAddressId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAddressesApiService.PATCHCustomerAddressesCustomerAddressId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_addresses/{customerAddressId}"
	localVarPath = strings.Replace(localVarPath, "{"+"customerAddressId"+"}", url.PathEscape(parameterValueToString(r.customerAddressId, "customerAddressId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerAddressUpdate == nil {
		return localVarReturnValue, nil, reportError("customerAddressUpdate is required and must be specified")
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
	localVarPostBody = r.customerAddressUpdate
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

type CustomerAddressesApiPOSTCustomerAddressesRequest struct {
	ctx                   context.Context
	ApiService            *CustomerAddressesApiService
	customerAddressCreate *CustomerAddressCreate
}

func (r CustomerAddressesApiPOSTCustomerAddressesRequest) CustomerAddressCreate(customerAddressCreate CustomerAddressCreate) CustomerAddressesApiPOSTCustomerAddressesRequest {
	r.customerAddressCreate = &customerAddressCreate
	return r
}

func (r CustomerAddressesApiPOSTCustomerAddressesRequest) Execute() (*POSTCustomerAddresses201Response, *http.Response, error) {
	return r.ApiService.POSTCustomerAddressesExecute(r)
}

/*
POSTCustomerAddresses Create a customer address

Create a customer address

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CustomerAddressesApiPOSTCustomerAddressesRequest
*/
func (a *CustomerAddressesApiService) POSTCustomerAddresses(ctx context.Context) CustomerAddressesApiPOSTCustomerAddressesRequest {
	return CustomerAddressesApiPOSTCustomerAddressesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTCustomerAddresses201Response
func (a *CustomerAddressesApiService) POSTCustomerAddressesExecute(r CustomerAddressesApiPOSTCustomerAddressesRequest) (*POSTCustomerAddresses201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTCustomerAddresses201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomerAddressesApiService.POSTCustomerAddresses")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customer_addresses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.customerAddressCreate == nil {
		return localVarReturnValue, nil, reportError("customerAddressCreate is required and must be specified")
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
	localVarPostBody = r.customerAddressCreate
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
