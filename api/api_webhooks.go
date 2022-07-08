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

// WebhooksApiService WebhooksApi service
type WebhooksApiService service

type ApiDELETEWebhooksWebhookIdRequest struct {
	ctx        context.Context
	ApiService *WebhooksApiService
	webhookId  string
}

func (r ApiDELETEWebhooksWebhookIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEWebhooksWebhookIdExecute(r)
}

/*
DELETEWebhooksWebhookId Delete a webhook

Delete a webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookId The resource's id
 @return ApiDELETEWebhooksWebhookIdRequest
*/
func (a *WebhooksApiService) DELETEWebhooksWebhookId(ctx context.Context, webhookId string) ApiDELETEWebhooksWebhookIdRequest {
	return ApiDELETEWebhooksWebhookIdRequest{
		ApiService: a,
		ctx:        ctx,
		webhookId:  webhookId,
	}
}

// Execute executes the request
func (a *WebhooksApiService) DELETEWebhooksWebhookIdExecute(r ApiDELETEWebhooksWebhookIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.DELETEWebhooksWebhookId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

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

type ApiGETEventCallbackIdWebhookRequest struct {
	ctx             context.Context
	ApiService      *WebhooksApiService
	eventCallbackId string
}

func (r ApiGETEventCallbackIdWebhookRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETEventCallbackIdWebhookExecute(r)
}

/*
GETEventCallbackIdWebhook Retrieve the webhook associated to the event callback

Retrieve the webhook associated to the event callback

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventCallbackId The resource's id
 @return ApiGETEventCallbackIdWebhookRequest
*/
func (a *WebhooksApiService) GETEventCallbackIdWebhook(ctx context.Context, eventCallbackId string) ApiGETEventCallbackIdWebhookRequest {
	return ApiGETEventCallbackIdWebhookRequest{
		ApiService:      a,
		ctx:             ctx,
		eventCallbackId: eventCallbackId,
	}
}

// Execute executes the request
func (a *WebhooksApiService) GETEventCallbackIdWebhookExecute(r ApiGETEventCallbackIdWebhookRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.GETEventCallbackIdWebhook")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/event_callbacks/{eventCallbackId}/webhook"
	localVarPath = strings.Replace(localVarPath, "{"+"eventCallbackId"+"}", url.PathEscape(parameterToString(r.eventCallbackId, "")), -1)

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

type ApiGETEventIdWebhooksRequest struct {
	ctx        context.Context
	ApiService *WebhooksApiService
	eventId    string
}

func (r ApiGETEventIdWebhooksRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETEventIdWebhooksExecute(r)
}

/*
GETEventIdWebhooks Retrieve the webhooks associated to the event

Retrieve the webhooks associated to the event

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param eventId The resource's id
 @return ApiGETEventIdWebhooksRequest
*/
func (a *WebhooksApiService) GETEventIdWebhooks(ctx context.Context, eventId string) ApiGETEventIdWebhooksRequest {
	return ApiGETEventIdWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
		eventId:    eventId,
	}
}

// Execute executes the request
func (a *WebhooksApiService) GETEventIdWebhooksExecute(r ApiGETEventIdWebhooksRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.GETEventIdWebhooks")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/events/{eventId}/webhooks"
	localVarPath = strings.Replace(localVarPath, "{"+"eventId"+"}", url.PathEscape(parameterToString(r.eventId, "")), -1)

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

type ApiGETWebhooksRequest struct {
	ctx        context.Context
	ApiService *WebhooksApiService
}

func (r ApiGETWebhooksRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETWebhooksExecute(r)
}

/*
GETWebhooks List all webhooks

List all webhooks

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGETWebhooksRequest
*/
func (a *WebhooksApiService) GETWebhooks(ctx context.Context) ApiGETWebhooksRequest {
	return ApiGETWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *WebhooksApiService) GETWebhooksExecute(r ApiGETWebhooksRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.GETWebhooks")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks"

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

type ApiGETWebhooksWebhookIdRequest struct {
	ctx        context.Context
	ApiService *WebhooksApiService
	webhookId  string
}

func (r ApiGETWebhooksWebhookIdRequest) Execute() (*Webhook, *http.Response, error) {
	return r.ApiService.GETWebhooksWebhookIdExecute(r)
}

/*
GETWebhooksWebhookId Retrieve a webhook

Retrieve a webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookId The resource's id
 @return ApiGETWebhooksWebhookIdRequest
*/
func (a *WebhooksApiService) GETWebhooksWebhookId(ctx context.Context, webhookId string) ApiGETWebhooksWebhookIdRequest {
	return ApiGETWebhooksWebhookIdRequest{
		ApiService: a,
		ctx:        ctx,
		webhookId:  webhookId,
	}
}

// Execute executes the request
//  @return Webhook
func (a *WebhooksApiService) GETWebhooksWebhookIdExecute(r ApiGETWebhooksWebhookIdRequest) (*Webhook, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Webhook
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.GETWebhooksWebhookId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

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

type ApiPATCHWebhooksWebhookIdRequest struct {
	ctx           context.Context
	ApiService    *WebhooksApiService
	webhookUpdate *WebhookUpdate
	webhookId     string
}

func (r ApiPATCHWebhooksWebhookIdRequest) WebhookUpdate(webhookUpdate WebhookUpdate) ApiPATCHWebhooksWebhookIdRequest {
	r.webhookUpdate = &webhookUpdate
	return r
}

func (r ApiPATCHWebhooksWebhookIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PATCHWebhooksWebhookIdExecute(r)
}

/*
PATCHWebhooksWebhookId Update a webhook

Update a webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param webhookId The resource's id
 @return ApiPATCHWebhooksWebhookIdRequest
*/
func (a *WebhooksApiService) PATCHWebhooksWebhookId(ctx context.Context, webhookId string) ApiPATCHWebhooksWebhookIdRequest {
	return ApiPATCHWebhooksWebhookIdRequest{
		ApiService: a,
		ctx:        ctx,
		webhookId:  webhookId,
	}
}

// Execute executes the request
func (a *WebhooksApiService) PATCHWebhooksWebhookIdExecute(r ApiPATCHWebhooksWebhookIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.PATCHWebhooksWebhookId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks/{webhookId}"
	localVarPath = strings.Replace(localVarPath, "{"+"webhookId"+"}", url.PathEscape(parameterToString(r.webhookId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookUpdate == nil {
		return nil, reportError("webhookUpdate is required and must be specified")
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
	localVarPostBody = r.webhookUpdate
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

type ApiPOSTWebhooksRequest struct {
	ctx           context.Context
	ApiService    *WebhooksApiService
	webhookCreate *WebhookCreate
}

func (r ApiPOSTWebhooksRequest) WebhookCreate(webhookCreate WebhookCreate) ApiPOSTWebhooksRequest {
	r.webhookCreate = &webhookCreate
	return r
}

func (r ApiPOSTWebhooksRequest) Execute() (*http.Response, error) {
	return r.ApiService.POSTWebhooksExecute(r)
}

/*
POSTWebhooks Create a webhook

Create a webhook

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPOSTWebhooksRequest
*/
func (a *WebhooksApiService) POSTWebhooks(ctx context.Context) ApiPOSTWebhooksRequest {
	return ApiPOSTWebhooksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *WebhooksApiService) POSTWebhooksExecute(r ApiPOSTWebhooksRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WebhooksApiService.POSTWebhooks")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/webhooks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.webhookCreate == nil {
		return nil, reportError("webhookCreate is required and must be specified")
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
	localVarPostBody = r.webhookCreate
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
