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

type ReturnsApi interface {

	/*
		DELETEReturnsReturnId Delete a return

		Delete a return

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param returnId The resource's id
		@return ReturnsApiDELETEReturnsReturnIdRequest
	*/
	DELETEReturnsReturnId(ctx context.Context, returnId string) ReturnsApiDELETEReturnsReturnIdRequest

	// DELETEReturnsReturnIdExecute executes the request
	DELETEReturnsReturnIdExecute(r ReturnsApiDELETEReturnsReturnIdRequest) (*http.Response, error)

	/*
		GETCustomerIdReturns Retrieve the returns associated to the customer

		Retrieve the returns associated to the customer

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param customerId The resource's id
		@return ReturnsApiGETCustomerIdReturnsRequest
	*/
	GETCustomerIdReturns(ctx context.Context, customerId string) ReturnsApiGETCustomerIdReturnsRequest

	// GETCustomerIdReturnsExecute executes the request
	GETCustomerIdReturnsExecute(r ReturnsApiGETCustomerIdReturnsRequest) (*http.Response, error)

	/*
		GETReturnLineItemIdReturn Retrieve the return associated to the return line item

		Retrieve the return associated to the return line item

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param returnLineItemId The resource's id
		@return ReturnsApiGETReturnLineItemIdReturnRequest
	*/
	GETReturnLineItemIdReturn(ctx context.Context, returnLineItemId string) ReturnsApiGETReturnLineItemIdReturnRequest

	// GETReturnLineItemIdReturnExecute executes the request
	GETReturnLineItemIdReturnExecute(r ReturnsApiGETReturnLineItemIdReturnRequest) (*http.Response, error)

	/*
		GETReturns List all returns

		List all returns

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ReturnsApiGETReturnsRequest
	*/
	GETReturns(ctx context.Context) ReturnsApiGETReturnsRequest

	// GETReturnsExecute executes the request
	GETReturnsExecute(r ReturnsApiGETReturnsRequest) (*http.Response, error)

	/*
		GETReturnsReturnId Retrieve a return

		Retrieve a return

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param returnId The resource's id
		@return ReturnsApiGETReturnsReturnIdRequest
	*/
	GETReturnsReturnId(ctx context.Context, returnId string) ReturnsApiGETReturnsReturnIdRequest

	// GETReturnsReturnIdExecute executes the request
	//  @return ModelReturn
	GETReturnsReturnIdExecute(r ReturnsApiGETReturnsReturnIdRequest) (*ModelReturn, *http.Response, error)

	/*
		PATCHReturnsReturnId Update a return

		Update a return

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param returnId The resource's id
		@return ReturnsApiPATCHReturnsReturnIdRequest
	*/
	PATCHReturnsReturnId(ctx context.Context, returnId string) ReturnsApiPATCHReturnsReturnIdRequest

	// PATCHReturnsReturnIdExecute executes the request
	PATCHReturnsReturnIdExecute(r ReturnsApiPATCHReturnsReturnIdRequest) (*http.Response, error)

	/*
		POSTReturns Create a return

		Create a return

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ReturnsApiPOSTReturnsRequest
	*/
	POSTReturns(ctx context.Context) ReturnsApiPOSTReturnsRequest

	// POSTReturnsExecute executes the request
	POSTReturnsExecute(r ReturnsApiPOSTReturnsRequest) (*http.Response, error)
}

// ReturnsApiService ReturnsApi service
type ReturnsApiService service

type ReturnsApiDELETEReturnsReturnIdRequest struct {
	ctx        context.Context
	ApiService ReturnsApi
	returnId   string
}

func (r ReturnsApiDELETEReturnsReturnIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEReturnsReturnIdExecute(r)
}

/*
DELETEReturnsReturnId Delete a return

Delete a return

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnId The resource's id
 @return ReturnsApiDELETEReturnsReturnIdRequest
*/
func (a *ReturnsApiService) DELETEReturnsReturnId(ctx context.Context, returnId string) ReturnsApiDELETEReturnsReturnIdRequest {
	return ReturnsApiDELETEReturnsReturnIdRequest{
		ApiService: a,
		ctx:        ctx,
		returnId:   returnId,
	}
}

// Execute executes the request
func (a *ReturnsApiService) DELETEReturnsReturnIdExecute(r ReturnsApiDELETEReturnsReturnIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.DELETEReturnsReturnId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/returns/{returnId}"
	localVarPath = strings.Replace(localVarPath, "{"+"returnId"+"}", url.PathEscape(parameterToString(r.returnId, "")), -1)

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

type ReturnsApiGETCustomerIdReturnsRequest struct {
	ctx        context.Context
	ApiService ReturnsApi
	customerId string
}

func (r ReturnsApiGETCustomerIdReturnsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETCustomerIdReturnsExecute(r)
}

/*
GETCustomerIdReturns Retrieve the returns associated to the customer

Retrieve the returns associated to the customer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customerId The resource's id
 @return ReturnsApiGETCustomerIdReturnsRequest
*/
func (a *ReturnsApiService) GETCustomerIdReturns(ctx context.Context, customerId string) ReturnsApiGETCustomerIdReturnsRequest {
	return ReturnsApiGETCustomerIdReturnsRequest{
		ApiService: a,
		ctx:        ctx,
		customerId: customerId,
	}
}

// Execute executes the request
func (a *ReturnsApiService) GETCustomerIdReturnsExecute(r ReturnsApiGETCustomerIdReturnsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.GETCustomerIdReturns")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/customers/{customerId}/returns"
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

type ReturnsApiGETReturnLineItemIdReturnRequest struct {
	ctx              context.Context
	ApiService       ReturnsApi
	returnLineItemId string
}

func (r ReturnsApiGETReturnLineItemIdReturnRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETReturnLineItemIdReturnExecute(r)
}

/*
GETReturnLineItemIdReturn Retrieve the return associated to the return line item

Retrieve the return associated to the return line item

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnLineItemId The resource's id
 @return ReturnsApiGETReturnLineItemIdReturnRequest
*/
func (a *ReturnsApiService) GETReturnLineItemIdReturn(ctx context.Context, returnLineItemId string) ReturnsApiGETReturnLineItemIdReturnRequest {
	return ReturnsApiGETReturnLineItemIdReturnRequest{
		ApiService:       a,
		ctx:              ctx,
		returnLineItemId: returnLineItemId,
	}
}

// Execute executes the request
func (a *ReturnsApiService) GETReturnLineItemIdReturnExecute(r ReturnsApiGETReturnLineItemIdReturnRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.GETReturnLineItemIdReturn")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/return_line_items/{returnLineItemId}/return"
	localVarPath = strings.Replace(localVarPath, "{"+"returnLineItemId"+"}", url.PathEscape(parameterToString(r.returnLineItemId, "")), -1)

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

type ReturnsApiGETReturnsRequest struct {
	ctx        context.Context
	ApiService ReturnsApi
}

func (r ReturnsApiGETReturnsRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETReturnsExecute(r)
}

/*
GETReturns List all returns

List all returns

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ReturnsApiGETReturnsRequest
*/
func (a *ReturnsApiService) GETReturns(ctx context.Context) ReturnsApiGETReturnsRequest {
	return ReturnsApiGETReturnsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ReturnsApiService) GETReturnsExecute(r ReturnsApiGETReturnsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.GETReturns")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/returns"

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

type ReturnsApiGETReturnsReturnIdRequest struct {
	ctx        context.Context
	ApiService ReturnsApi
	returnId   string
}

func (r ReturnsApiGETReturnsReturnIdRequest) Execute() (*ModelReturn, *http.Response, error) {
	return r.ApiService.GETReturnsReturnIdExecute(r)
}

/*
GETReturnsReturnId Retrieve a return

Retrieve a return

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnId The resource's id
 @return ReturnsApiGETReturnsReturnIdRequest
*/
func (a *ReturnsApiService) GETReturnsReturnId(ctx context.Context, returnId string) ReturnsApiGETReturnsReturnIdRequest {
	return ReturnsApiGETReturnsReturnIdRequest{
		ApiService: a,
		ctx:        ctx,
		returnId:   returnId,
	}
}

// Execute executes the request
//  @return ModelReturn
func (a *ReturnsApiService) GETReturnsReturnIdExecute(r ReturnsApiGETReturnsReturnIdRequest) (*ModelReturn, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ModelReturn
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.GETReturnsReturnId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/returns/{returnId}"
	localVarPath = strings.Replace(localVarPath, "{"+"returnId"+"}", url.PathEscape(parameterToString(r.returnId, "")), -1)

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

type ReturnsApiPATCHReturnsReturnIdRequest struct {
	ctx          context.Context
	ApiService   ReturnsApi
	returnId     string
	returnUpdate *ReturnUpdate
}

func (r ReturnsApiPATCHReturnsReturnIdRequest) ReturnUpdate(returnUpdate ReturnUpdate) ReturnsApiPATCHReturnsReturnIdRequest {
	r.returnUpdate = &returnUpdate
	return r
}

func (r ReturnsApiPATCHReturnsReturnIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHReturnsReturnIdExecute(r)
}

/*
PATCHReturnsReturnId Update a return

Update a return

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param returnId The resource's id
 @return ReturnsApiPATCHReturnsReturnIdRequest
*/
func (a *ReturnsApiService) PATCHReturnsReturnId(ctx context.Context, returnId string) ReturnsApiPATCHReturnsReturnIdRequest {
	return ReturnsApiPATCHReturnsReturnIdRequest{
		ApiService: a,
		ctx:        ctx,
		returnId:   returnId,
	}
}

// Execute executes the request
func (a *ReturnsApiService) PATCHReturnsReturnIdExecute(r ReturnsApiPATCHReturnsReturnIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.PATCHReturnsReturnId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/returns/{returnId}"
	localVarPath = strings.Replace(localVarPath, "{"+"returnId"+"}", url.PathEscape(parameterToString(r.returnId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.returnUpdate == nil {
		return nil, reportError("returnUpdate is required and must be specified")
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
	localVarPostBody = r.returnUpdate
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

type ReturnsApiPOSTReturnsRequest struct {
	ctx          context.Context
	ApiService   ReturnsApi
	returnCreate *ReturnCreate
}

func (r ReturnsApiPOSTReturnsRequest) ReturnCreate(returnCreate ReturnCreate) ReturnsApiPOSTReturnsRequest {
	r.returnCreate = &returnCreate
	return r
}

func (r ReturnsApiPOSTReturnsRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTReturnsExecute(r)
}

/*
POSTReturns Create a return

Create a return

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ReturnsApiPOSTReturnsRequest
*/
func (a *ReturnsApiService) POSTReturns(ctx context.Context) ReturnsApiPOSTReturnsRequest {
	return ReturnsApiPOSTReturnsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ReturnsApiService) POSTReturnsExecute(r ReturnsApiPOSTReturnsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReturnsApiService.POSTReturns")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/returns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.returnCreate == nil {
		return nil, reportError("returnCreate is required and must be specified")
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
	localVarPostBody = r.returnCreate
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
