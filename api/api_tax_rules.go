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

// TaxRulesApiService TaxRulesApi service
type TaxRulesApiService service

type TaxRulesApiDELETETaxRulesTaxRuleIdRequest struct {
	ctx        context.Context
	ApiService *TaxRulesApiService
	taxRuleId  interface{}
}

func (r TaxRulesApiDELETETaxRulesTaxRuleIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETETaxRulesTaxRuleIdExecute(r)
}

/*
DELETETaxRulesTaxRuleId Delete a tax rule

Delete a tax rule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxRuleId The resource's id
	@return TaxRulesApiDELETETaxRulesTaxRuleIdRequest
*/
func (a *TaxRulesApiService) DELETETaxRulesTaxRuleId(ctx context.Context, taxRuleId interface{}) TaxRulesApiDELETETaxRulesTaxRuleIdRequest {
	return TaxRulesApiDELETETaxRulesTaxRuleIdRequest{
		ApiService: a,
		ctx:        ctx,
		taxRuleId:  taxRuleId,
	}
}

// Execute executes the request
func (a *TaxRulesApiService) DELETETaxRulesTaxRuleIdExecute(r TaxRulesApiDELETETaxRulesTaxRuleIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxRulesApiService.DELETETaxRulesTaxRuleId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_rules/{taxRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxRuleId"+"}", url.PathEscape(parameterValueToString(r.taxRuleId, "taxRuleId")), -1)

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

type TaxRulesApiGETManualTaxCalculatorIdTaxRulesRequest struct {
	ctx                   context.Context
	ApiService            *TaxRulesApiService
	manualTaxCalculatorId interface{}
}

func (r TaxRulesApiGETManualTaxCalculatorIdTaxRulesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETManualTaxCalculatorIdTaxRulesExecute(r)
}

/*
GETManualTaxCalculatorIdTaxRules Retrieve the tax rules associated to the manual tax calculator

Retrieve the tax rules associated to the manual tax calculator

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param manualTaxCalculatorId The resource's id
	@return TaxRulesApiGETManualTaxCalculatorIdTaxRulesRequest
*/
func (a *TaxRulesApiService) GETManualTaxCalculatorIdTaxRules(ctx context.Context, manualTaxCalculatorId interface{}) TaxRulesApiGETManualTaxCalculatorIdTaxRulesRequest {
	return TaxRulesApiGETManualTaxCalculatorIdTaxRulesRequest{
		ApiService:            a,
		ctx:                   ctx,
		manualTaxCalculatorId: manualTaxCalculatorId,
	}
}

// Execute executes the request
func (a *TaxRulesApiService) GETManualTaxCalculatorIdTaxRulesExecute(r TaxRulesApiGETManualTaxCalculatorIdTaxRulesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxRulesApiService.GETManualTaxCalculatorIdTaxRules")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/manual_tax_calculators/{manualTaxCalculatorId}/tax_rules"
	localVarPath = strings.Replace(localVarPath, "{"+"manualTaxCalculatorId"+"}", url.PathEscape(parameterValueToString(r.manualTaxCalculatorId, "manualTaxCalculatorId")), -1)

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

type TaxRulesApiGETTaxRulesRequest struct {
	ctx        context.Context
	ApiService *TaxRulesApiService
}

func (r TaxRulesApiGETTaxRulesRequest) Execute() (*GETTaxRules200Response, *http.Response, error) {
	return r.ApiService.GETTaxRulesExecute(r)
}

/*
GETTaxRules List all tax rules

List all tax rules

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TaxRulesApiGETTaxRulesRequest
*/
func (a *TaxRulesApiService) GETTaxRules(ctx context.Context) TaxRulesApiGETTaxRulesRequest {
	return TaxRulesApiGETTaxRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETTaxRules200Response
func (a *TaxRulesApiService) GETTaxRulesExecute(r TaxRulesApiGETTaxRulesRequest) (*GETTaxRules200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETTaxRules200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxRulesApiService.GETTaxRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_rules"

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

type TaxRulesApiGETTaxRulesTaxRuleIdRequest struct {
	ctx        context.Context
	ApiService *TaxRulesApiService
	taxRuleId  interface{}
}

func (r TaxRulesApiGETTaxRulesTaxRuleIdRequest) Execute() (*GETTaxRulesTaxRuleId200Response, *http.Response, error) {
	return r.ApiService.GETTaxRulesTaxRuleIdExecute(r)
}

/*
GETTaxRulesTaxRuleId Retrieve a tax rule

Retrieve a tax rule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxRuleId The resource's id
	@return TaxRulesApiGETTaxRulesTaxRuleIdRequest
*/
func (a *TaxRulesApiService) GETTaxRulesTaxRuleId(ctx context.Context, taxRuleId interface{}) TaxRulesApiGETTaxRulesTaxRuleIdRequest {
	return TaxRulesApiGETTaxRulesTaxRuleIdRequest{
		ApiService: a,
		ctx:        ctx,
		taxRuleId:  taxRuleId,
	}
}

// Execute executes the request
//
//	@return GETTaxRulesTaxRuleId200Response
func (a *TaxRulesApiService) GETTaxRulesTaxRuleIdExecute(r TaxRulesApiGETTaxRulesTaxRuleIdRequest) (*GETTaxRulesTaxRuleId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETTaxRulesTaxRuleId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxRulesApiService.GETTaxRulesTaxRuleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_rules/{taxRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxRuleId"+"}", url.PathEscape(parameterValueToString(r.taxRuleId, "taxRuleId")), -1)

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

type TaxRulesApiPATCHTaxRulesTaxRuleIdRequest struct {
	ctx           context.Context
	ApiService    *TaxRulesApiService
	taxRuleUpdate *TaxRuleUpdate
	taxRuleId     interface{}
}

func (r TaxRulesApiPATCHTaxRulesTaxRuleIdRequest) TaxRuleUpdate(taxRuleUpdate TaxRuleUpdate) TaxRulesApiPATCHTaxRulesTaxRuleIdRequest {
	r.taxRuleUpdate = &taxRuleUpdate
	return r
}

func (r TaxRulesApiPATCHTaxRulesTaxRuleIdRequest) Execute() (*PATCHTaxRulesTaxRuleId200Response, *http.Response, error) {
	return r.ApiService.PATCHTaxRulesTaxRuleIdExecute(r)
}

/*
PATCHTaxRulesTaxRuleId Update a tax rule

Update a tax rule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taxRuleId The resource's id
	@return TaxRulesApiPATCHTaxRulesTaxRuleIdRequest
*/
func (a *TaxRulesApiService) PATCHTaxRulesTaxRuleId(ctx context.Context, taxRuleId interface{}) TaxRulesApiPATCHTaxRulesTaxRuleIdRequest {
	return TaxRulesApiPATCHTaxRulesTaxRuleIdRequest{
		ApiService: a,
		ctx:        ctx,
		taxRuleId:  taxRuleId,
	}
}

// Execute executes the request
//
//	@return PATCHTaxRulesTaxRuleId200Response
func (a *TaxRulesApiService) PATCHTaxRulesTaxRuleIdExecute(r TaxRulesApiPATCHTaxRulesTaxRuleIdRequest) (*PATCHTaxRulesTaxRuleId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHTaxRulesTaxRuleId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxRulesApiService.PATCHTaxRulesTaxRuleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_rules/{taxRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taxRuleId"+"}", url.PathEscape(parameterValueToString(r.taxRuleId, "taxRuleId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.taxRuleUpdate == nil {
		return localVarReturnValue, nil, reportError("taxRuleUpdate is required and must be specified")
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
	localVarPostBody = r.taxRuleUpdate
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

type TaxRulesApiPOSTTaxRulesRequest struct {
	ctx           context.Context
	ApiService    *TaxRulesApiService
	taxRuleCreate *TaxRuleCreate
}

func (r TaxRulesApiPOSTTaxRulesRequest) TaxRuleCreate(taxRuleCreate TaxRuleCreate) TaxRulesApiPOSTTaxRulesRequest {
	r.taxRuleCreate = &taxRuleCreate
	return r
}

func (r TaxRulesApiPOSTTaxRulesRequest) Execute() (*POSTTaxRules201Response, *http.Response, error) {
	return r.ApiService.POSTTaxRulesExecute(r)
}

/*
POSTTaxRules Create a tax rule

Create a tax rule

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TaxRulesApiPOSTTaxRulesRequest
*/
func (a *TaxRulesApiService) POSTTaxRules(ctx context.Context) TaxRulesApiPOSTTaxRulesRequest {
	return TaxRulesApiPOSTTaxRulesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTTaxRules201Response
func (a *TaxRulesApiService) POSTTaxRulesExecute(r TaxRulesApiPOSTTaxRulesRequest) (*POSTTaxRules201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTTaxRules201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaxRulesApiService.POSTTaxRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tax_rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.taxRuleCreate == nil {
		return localVarReturnValue, nil, reportError("taxRuleCreate is required and must be specified")
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
	localVarPostBody = r.taxRuleCreate
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
