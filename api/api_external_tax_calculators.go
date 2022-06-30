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

type ExternalTaxCalculatorsApi interface {

	/*
		DELETEExternalTaxCalculatorsExternalTaxCalculatorId Delete an external tax calculator

		Delete an external tax calculator

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalTaxCalculatorId The resource's id
		@return ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest
	*/
	DELETEExternalTaxCalculatorsExternalTaxCalculatorId(ctx context.Context, externalTaxCalculatorId string) ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest

	// DELETEExternalTaxCalculatorsExternalTaxCalculatorIdExecute executes the request
	DELETEExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest) (*http.Response, error)

	/*
		GETExternalTaxCalculators List all external tax calculators

		List all external tax calculators

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest
	*/
	GETExternalTaxCalculators(ctx context.Context) ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest

	// GETExternalTaxCalculatorsExecute executes the request
	GETExternalTaxCalculatorsExecute(r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest) (*http.Response, error)

	/*
		GETExternalTaxCalculatorsExternalTaxCalculatorId Retrieve an external tax calculator

		Retrieve an external tax calculator

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalTaxCalculatorId The resource's id
		@return ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest
	*/
	GETExternalTaxCalculatorsExternalTaxCalculatorId(ctx context.Context, externalTaxCalculatorId string) ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest

	// GETExternalTaxCalculatorsExternalTaxCalculatorIdExecute executes the request
	//  @return ExternalTaxCalculator
	GETExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest) (*ExternalTaxCalculator, *http.Response, error)

	/*
		PATCHExternalTaxCalculatorsExternalTaxCalculatorId Update an external tax calculator

		Update an external tax calculator

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param externalTaxCalculatorId The resource's id
		@return ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest
	*/
	PATCHExternalTaxCalculatorsExternalTaxCalculatorId(ctx context.Context, externalTaxCalculatorId string) ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest

	// PATCHExternalTaxCalculatorsExternalTaxCalculatorIdExecute executes the request
	PATCHExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest) (*http.Response, error)

	/*
		POSTExternalTaxCalculators Create an external tax calculator

		Create an external tax calculator

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest
	*/
	POSTExternalTaxCalculators(ctx context.Context) ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest

	// POSTExternalTaxCalculatorsExecute executes the request
	POSTExternalTaxCalculatorsExecute(r ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest) (*http.Response, error)
}

// ExternalTaxCalculatorsApiService ExternalTaxCalculatorsApi service
type ExternalTaxCalculatorsApiService service

type ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest struct {
	ctx                     context.Context
	ApiService              ExternalTaxCalculatorsApi
	externalTaxCalculatorId string
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
func (a *ExternalTaxCalculatorsApiService) DELETEExternalTaxCalculatorsExternalTaxCalculatorId(ctx context.Context, externalTaxCalculatorId string) ExternalTaxCalculatorsApiDELETEExternalTaxCalculatorsExternalTaxCalculatorIdRequest {
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
	localVarPath = strings.Replace(localVarPath, "{"+"externalTaxCalculatorId"+"}", url.PathEscape(parameterToString(r.externalTaxCalculatorId, "")), -1)

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

type ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest struct {
	ctx        context.Context
	ApiService ExternalTaxCalculatorsApi
}

func (r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest) Execute() (*http.Response, error) {
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
func (a *ExternalTaxCalculatorsApiService) GETExternalTaxCalculatorsExecute(r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalTaxCalculatorsApiService.GETExternalTaxCalculators")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
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

type ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest struct {
	ctx                     context.Context
	ApiService              ExternalTaxCalculatorsApi
	externalTaxCalculatorId string
}

func (r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest) Execute() (*ExternalTaxCalculator, *http.Response, error) {
	return r.ApiService.GETExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r)
}

/*
GETExternalTaxCalculatorsExternalTaxCalculatorId Retrieve an external tax calculator

Retrieve an external tax calculator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalTaxCalculatorId The resource's id
 @return ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest
*/
func (a *ExternalTaxCalculatorsApiService) GETExternalTaxCalculatorsExternalTaxCalculatorId(ctx context.Context, externalTaxCalculatorId string) ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest {
	return ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		externalTaxCalculatorId: externalTaxCalculatorId,
	}
}

// Execute executes the request
//  @return ExternalTaxCalculator
func (a *ExternalTaxCalculatorsApiService) GETExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r ExternalTaxCalculatorsApiGETExternalTaxCalculatorsExternalTaxCalculatorIdRequest) (*ExternalTaxCalculator, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExternalTaxCalculator
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalTaxCalculatorsApiService.GETExternalTaxCalculatorsExternalTaxCalculatorId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_tax_calculators/{externalTaxCalculatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalTaxCalculatorId"+"}", url.PathEscape(parameterToString(r.externalTaxCalculatorId, "")), -1)

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

type ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest struct {
	ctx                         context.Context
	ApiService                  ExternalTaxCalculatorsApi
	externalTaxCalculatorId     string
	externalTaxCalculatorUpdate *ExternalTaxCalculatorUpdate
}

func (r ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest) ExternalTaxCalculatorUpdate(externalTaxCalculatorUpdate ExternalTaxCalculatorUpdate) ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest {
	r.externalTaxCalculatorUpdate = &externalTaxCalculatorUpdate
	return r
}

func (r ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r)
}

/*
PATCHExternalTaxCalculatorsExternalTaxCalculatorId Update an external tax calculator

Update an external tax calculator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalTaxCalculatorId The resource's id
 @return ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest
*/
func (a *ExternalTaxCalculatorsApiService) PATCHExternalTaxCalculatorsExternalTaxCalculatorId(ctx context.Context, externalTaxCalculatorId string) ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest {
	return ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest{
		ApiService:              a,
		ctx:                     ctx,
		externalTaxCalculatorId: externalTaxCalculatorId,
	}
}

// Execute executes the request
func (a *ExternalTaxCalculatorsApiService) PATCHExternalTaxCalculatorsExternalTaxCalculatorIdExecute(r ExternalTaxCalculatorsApiPATCHExternalTaxCalculatorsExternalTaxCalculatorIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalTaxCalculatorsApiService.PATCHExternalTaxCalculatorsExternalTaxCalculatorId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_tax_calculators/{externalTaxCalculatorId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalTaxCalculatorId"+"}", url.PathEscape(parameterToString(r.externalTaxCalculatorId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalTaxCalculatorUpdate == nil {
		return nil, reportError("externalTaxCalculatorUpdate is required and must be specified")
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
	localVarPostBody = r.externalTaxCalculatorUpdate
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

type ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest struct {
	ctx                         context.Context
	ApiService                  ExternalTaxCalculatorsApi
	externalTaxCalculatorCreate *ExternalTaxCalculatorCreate
}

func (r ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest) ExternalTaxCalculatorCreate(externalTaxCalculatorCreate ExternalTaxCalculatorCreate) ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest {
	r.externalTaxCalculatorCreate = &externalTaxCalculatorCreate
	return r
}

func (r ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest) Execute() (*http.Response, error) {
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
func (a *ExternalTaxCalculatorsApiService) POSTExternalTaxCalculatorsExecute(r ExternalTaxCalculatorsApiPOSTExternalTaxCalculatorsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalTaxCalculatorsApiService.POSTExternalTaxCalculators")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/external_tax_calculators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.externalTaxCalculatorCreate == nil {
		return nil, reportError("externalTaxCalculatorCreate is required and must be specified")
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
	localVarPostBody = r.externalTaxCalculatorCreate
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
