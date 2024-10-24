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

// ExternalTaxCalculatorsApiService ExternalTaxCalculatorsApi service
type ExternalTaxCalculatorsApiService service

type ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest struct {
	ctx                     context.Context
	ApiService              *ExternalTaxCalculatorsApiService
	externalTaxCalculatorId interface{}
}

func (r ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r)
}

/*
DELETEExternalTaxCalculatorsExternalTaxCalculatorId Delete an external tax calculator

Delete an external tax calculator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalTaxCalculatorId The resource's id
	@return ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest
*/
func (a *ExternalTaxCalculatorsApiService) DELETEExternalTaxCalculatorsExternalTaxCalculatorId(ctx context.Context, externalTaxCalculatorId interface{}) ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest {
	return ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		externalTaxCalculatorId: externalTaxCalculatorId,
	}
}

// Execute executes the request
func (a *ExternalTaxCalculatorsApiService) DELETEExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalTaxCalculatorsApiService.DELETEExternalTaxCalculatorsExternalTaxCalculatorId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_tax_calculators/{externalTaxCalculatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalTaxCalculatorId"+"}", url.PathEscape(parameterValueToString(r.externalTaxCalculatorId, "externalTaxCalculatorId")), -1)

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

type ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest struct {
	ctx        context.Context
	ApiService *ExternalTaxCalculatorsApiService
}

func (r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest) Execute() (*GETExternalTaxCalculators200Response, *http.Response, error) {
	return r.ApiService.GETExternalTaxCalculatorsExecute(r)
}

/*
GETExternalTaxCalculators List all external tax calculators

List all external tax calculators

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest
*/
func (a *ExternalTaxCalculatorsApiService) GETExternalTaxCalculators(ctx context.Context) ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest {
	return ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETExternalTaxCalculators200Response
func (a *ExternalTaxCalculatorsApiService) GETExternalTaxCalculatorsExecute(r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest) (*GETExternalTaxCalculators200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETExternalTaxCalculators200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalTaxCalculatorsApiService.GETExternalTaxCalculators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_tax_calculators"

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

type ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest struct {
	ctx                     context.Context
	ApiService              *ExternalTaxCalculatorsApiService
	externalTaxCalculatorId interface{}
}

func (r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest) Execute() (*GETExternalTaxCalculatorsExternalTaxCalculatorId200Response, *http.Response, error) {
	return r.ApiService.GETExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r)
}

/*
GETExternalTaxCalculatorsExternalTaxCalculatorId Retrieve an external tax calculator

Retrieve an external tax calculator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalTaxCalculatorId The resource's id
	@return ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest
*/
func (a *ExternalTaxCalculatorsApiService) GETExternalTaxCalculatorsExternalTaxCalculatorId(ctx context.Context, externalTaxCalculatorId interface{}) ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest {
	return ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		externalTaxCalculatorId: externalTaxCalculatorId,
	}
}

// Execute executes the request
//
//	@return GETExternalTaxCalculatorsExternalTaxCalculatorId200Response
func (a *ExternalTaxCalculatorsApiService) GETExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest) (*GETExternalTaxCalculatorsExternalTaxCalculatorId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETExternalTaxCalculatorsExternalTaxCalculatorId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalTaxCalculatorsApiService.GETExternalTaxCalculatorsExternalTaxCalculatorId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_tax_calculators/{externalTaxCalculatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalTaxCalculatorId"+"}", url.PathEscape(parameterValueToString(r.externalTaxCalculatorId, "externalTaxCalculatorId")), -1)

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

type ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest struct {
	ctx                         context.Context
	ApiService                  *ExternalTaxCalculatorsApiService
	externalTaxCalculatorUpdate *ExternalTaxCalculatorUpdate
	externalTaxCalculatorId     interface{}
}

func (r ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest) ExternalTaxCalculatorUpdate(externalTaxCalculatorUpdate ExternalTaxCalculatorUpdate) ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest {
	r.externalTaxCalculatorUpdate = &externalTaxCalculatorUpdate
	return r
}

func (r ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest) Execute() (*PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response, *http.Response, error) {
	return r.ApiService.PATCHExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r)
}

/*
PATCHExternalTaxCalculatorsExternalTaxCalculatorId Update an external tax calculator

Update an external tax calculator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalTaxCalculatorId The resource's id
	@return ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest
*/
func (a *ExternalTaxCalculatorsApiService) PATCHExternalTaxCalculatorsExternalTaxCalculatorId(ctx context.Context, externalTaxCalculatorId interface{}) ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest {
	return ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		externalTaxCalculatorId: externalTaxCalculatorId,
	}
}

// Execute executes the request
//
//	@return PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response
func (a *ExternalTaxCalculatorsApiService) PATCHExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest) (*PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHExternalTaxCalculatorsExternalTaxCalculatorId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalTaxCalculatorsApiService.PATCHExternalTaxCalculatorsExternalTaxCalculatorId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_tax_calculators/{externalTaxCalculatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalTaxCalculatorId"+"}", url.PathEscape(parameterValueToString(r.externalTaxCalculatorId, "externalTaxCalculatorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalTaxCalculatorUpdate == nil {
		return localVarReturnValue, nil, reportError("externalTaxCalculatorUpdate is required and must be specified")
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
	localVarPostBody = r.externalTaxCalculatorUpdate
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

type ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest struct {
	ctx                         context.Context
	ApiService                  *ExternalTaxCalculatorsApiService
	externalTaxCalculatorCreate *ExternalTaxCalculatorCreate
}

func (r ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest) ExternalTaxCalculatorCreate(externalTaxCalculatorCreate ExternalTaxCalculatorCreate) ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest {
	r.externalTaxCalculatorCreate = &externalTaxCalculatorCreate
	return r
}

func (r ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest) Execute() (*POSTExternalTaxCalculators201Response, *http.Response, error) {
	return r.ApiService.POSTExternalTaxCalculatorsExecute(r)
}

/*
POSTExternalTaxCalculators Create an external tax calculator

Create an external tax calculator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest
*/
func (a *ExternalTaxCalculatorsApiService) POSTExternalTaxCalculators(ctx context.Context) ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest {
	return ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTExternalTaxCalculators201Response
func (a *ExternalTaxCalculatorsApiService) POSTExternalTaxCalculatorsExecute(r ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest) (*POSTExternalTaxCalculators201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTExternalTaxCalculators201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalTaxCalculatorsApiService.POSTExternalTaxCalculators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_tax_calculators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalTaxCalculatorCreate == nil {
		return localVarReturnValue, nil, reportError("externalTaxCalculatorCreate is required and must be specified")
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
	localVarPostBody = r.externalTaxCalculatorCreate
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
