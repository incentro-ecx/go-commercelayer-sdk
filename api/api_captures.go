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

// CapturesApiService CapturesApi service
type CapturesApiService service

type CapturesApiGETAuthorizationIdCapturesRequest struct {
	ctx             context.Context
	ApiService      *CapturesApiService
	authorizationId interface{}
}

func (r CapturesApiGETAuthorizationIdCapturesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETAuthorizationIdCapturesExecute(r)
}

/*
GETAuthorizationIdCaptures Retrieve the captures associated to the authorization

Retrieve the captures associated to the authorization

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param authorizationId The resource's id
	@return CapturesApiGETAuthorizationIdCapturesRequest
*/
func (a *CapturesApiService) GETAuthorizationIdCaptures(ctx context.Context, authorizationId interface{}) CapturesApiGETAuthorizationIdCapturesRequest {
	return CapturesApiGETAuthorizationIdCapturesRequest{
		ApiService:      a,
		ctx:             ctx,
		authorizationId: authorizationId,
	}
}

// Execute executes the request
func (a *CapturesApiService) GETAuthorizationIdCapturesExecute(r CapturesApiGETAuthorizationIdCapturesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CapturesApiService.GETAuthorizationIdCaptures")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/authorizations/{authorizationId}/captures"
	localVarPath = strings.Replace(localVarPath, "{"+"authorizationId"+"}", url.PathEscape(parameterValueToString(r.authorizationId, "authorizationId")), -1)

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

type CapturesApiGETCapturesRequest struct {
	ctx        context.Context
	ApiService *CapturesApiService
}

func (r CapturesApiGETCapturesRequest) Execute() (*GETCaptures200Response, *http.Response, error) {
	return r.ApiService.GETCapturesExecute(r)
}

/*
GETCaptures List all captures

List all captures

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CapturesApiGETCapturesRequest
*/
func (a *CapturesApiService) GETCaptures(ctx context.Context) CapturesApiGETCapturesRequest {
	return CapturesApiGETCapturesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETCaptures200Response
func (a *CapturesApiService) GETCapturesExecute(r CapturesApiGETCapturesRequest) (*GETCaptures200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCaptures200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CapturesApiService.GETCaptures")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/captures"

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

type CapturesApiGETCapturesCaptureIdRequest struct {
	ctx        context.Context
	ApiService *CapturesApiService
	captureId  interface{}
}

func (r CapturesApiGETCapturesCaptureIdRequest) Execute() (*GETCapturesCaptureId200Response, *http.Response, error) {
	return r.ApiService.GETCapturesCaptureIdExecute(r)
}

/*
GETCapturesCaptureId Retrieve a capture

Retrieve a capture

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param captureId The resource's id
	@return CapturesApiGETCapturesCaptureIdRequest
*/
func (a *CapturesApiService) GETCapturesCaptureId(ctx context.Context, captureId interface{}) CapturesApiGETCapturesCaptureIdRequest {
	return CapturesApiGETCapturesCaptureIdRequest{
		ApiService: a,
		ctx:        ctx,
		captureId:  captureId,
	}
}

// Execute executes the request
//
//	@return GETCapturesCaptureId200Response
func (a *CapturesApiService) GETCapturesCaptureIdExecute(r CapturesApiGETCapturesCaptureIdRequest) (*GETCapturesCaptureId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETCapturesCaptureId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CapturesApiService.GETCapturesCaptureId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/captures/{captureId}"
	localVarPath = strings.Replace(localVarPath, "{"+"captureId"+"}", url.PathEscape(parameterValueToString(r.captureId, "captureId")), -1)

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

type CapturesApiGETOrderIdCapturesRequest struct {
	ctx        context.Context
	ApiService *CapturesApiService
	orderId    interface{}
}

func (r CapturesApiGETOrderIdCapturesRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETOrderIdCapturesExecute(r)
}

/*
GETOrderIdCaptures Retrieve the captures associated to the order

Retrieve the captures associated to the order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orderId The resource's id
	@return CapturesApiGETOrderIdCapturesRequest
*/
func (a *CapturesApiService) GETOrderIdCaptures(ctx context.Context, orderId interface{}) CapturesApiGETOrderIdCapturesRequest {
	return CapturesApiGETOrderIdCapturesRequest{
		ApiService: a,
		ctx:        ctx,
		orderId:    orderId,
	}
}

// Execute executes the request
func (a *CapturesApiService) GETOrderIdCapturesExecute(r CapturesApiGETOrderIdCapturesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CapturesApiService.GETOrderIdCaptures")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/orders/{orderId}/captures"
	localVarPath = strings.Replace(localVarPath, "{"+"orderId"+"}", url.PathEscape(parameterValueToString(r.orderId, "orderId")), -1)

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

type CapturesApiGETRefundIdReferenceCaptureRequest struct {
	ctx        context.Context
	ApiService *CapturesApiService
	refundId   interface{}
}

func (r CapturesApiGETRefundIdReferenceCaptureRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETRefundIdReferenceCaptureExecute(r)
}

/*
GETRefundIdReferenceCapture Retrieve the reference capture associated to the refund

Retrieve the reference capture associated to the refund

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param refundId The resource's id
	@return CapturesApiGETRefundIdReferenceCaptureRequest
*/
func (a *CapturesApiService) GETRefundIdReferenceCapture(ctx context.Context, refundId interface{}) CapturesApiGETRefundIdReferenceCaptureRequest {
	return CapturesApiGETRefundIdReferenceCaptureRequest{
		ApiService: a,
		ctx:        ctx,
		refundId:   refundId,
	}
}

// Execute executes the request
func (a *CapturesApiService) GETRefundIdReferenceCaptureExecute(r CapturesApiGETRefundIdReferenceCaptureRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CapturesApiService.GETRefundIdReferenceCapture")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/refunds/{refundId}/reference_capture"
	localVarPath = strings.Replace(localVarPath, "{"+"refundId"+"}", url.PathEscape(parameterValueToString(r.refundId, "refundId")), -1)

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

type CapturesApiGETReturnIdReferenceCaptureRequest struct {
	ctx        context.Context
	ApiService *CapturesApiService
	returnId   interface{}
}

func (r CapturesApiGETReturnIdReferenceCaptureRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETReturnIdReferenceCaptureExecute(r)
}

/*
GETReturnIdReferenceCapture Retrieve the reference capture associated to the return

Retrieve the reference capture associated to the return

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param returnId The resource's id
	@return CapturesApiGETReturnIdReferenceCaptureRequest
*/
func (a *CapturesApiService) GETReturnIdReferenceCapture(ctx context.Context, returnId interface{}) CapturesApiGETReturnIdReferenceCaptureRequest {
	return CapturesApiGETReturnIdReferenceCaptureRequest{
		ApiService: a,
		ctx:        ctx,
		returnId:   returnId,
	}
}

// Execute executes the request
func (a *CapturesApiService) GETReturnIdReferenceCaptureExecute(r CapturesApiGETReturnIdReferenceCaptureRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CapturesApiService.GETReturnIdReferenceCapture")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/returns/{returnId}/reference_capture"
	localVarPath = strings.Replace(localVarPath, "{"+"returnId"+"}", url.PathEscape(parameterValueToString(r.returnId, "returnId")), -1)

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

type CapturesApiPATCHCapturesCaptureIdRequest struct {
	ctx           context.Context
	ApiService    *CapturesApiService
	captureUpdate *CaptureUpdate
	captureId     interface{}
}

func (r CapturesApiPATCHCapturesCaptureIdRequest) CaptureUpdate(captureUpdate CaptureUpdate) CapturesApiPATCHCapturesCaptureIdRequest {
	r.captureUpdate = &captureUpdate
	return r
}

func (r CapturesApiPATCHCapturesCaptureIdRequest) Execute() (*PATCHCapturesCaptureId200Response, *http.Response, error) {
	return r.ApiService.PATCHCapturesCaptureIdExecute(r)
}

/*
PATCHCapturesCaptureId Update a capture

Update a capture

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param captureId The resource's id
	@return CapturesApiPATCHCapturesCaptureIdRequest
*/
func (a *CapturesApiService) PATCHCapturesCaptureId(ctx context.Context, captureId interface{}) CapturesApiPATCHCapturesCaptureIdRequest {
	return CapturesApiPATCHCapturesCaptureIdRequest{
		ApiService: a,
		ctx:        ctx,
		captureId:  captureId,
	}
}

// Execute executes the request
//
//	@return PATCHCapturesCaptureId200Response
func (a *CapturesApiService) PATCHCapturesCaptureIdExecute(r CapturesApiPATCHCapturesCaptureIdRequest) (*PATCHCapturesCaptureId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHCapturesCaptureId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CapturesApiService.PATCHCapturesCaptureId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/captures/{captureId}"
	localVarPath = strings.Replace(localVarPath, "{"+"captureId"+"}", url.PathEscape(parameterValueToString(r.captureId, "captureId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.captureUpdate == nil {
		return localVarReturnValue, nil, reportError("captureUpdate is required and must be specified")
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
	localVarPostBody = r.captureUpdate
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
