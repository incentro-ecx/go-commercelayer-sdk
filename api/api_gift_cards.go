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

// GiftCardsApiService GiftCardsApi service
type GiftCardsApiService service

type GiftCardsApiDELETEGiftCardsGiftCardIdRequest struct {
	ctx        context.Context
	ApiService *GiftCardsApiService
	giftCardId string
}

func (r GiftCardsApiDELETEGiftCardsGiftCardIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DELETEGiftCardsGiftCardIdExecute(r)
}

/*
DELETEGiftCardsGiftCardId Delete a gift card

Delete a gift card

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param giftCardId The resource's id
	@return GiftCardsApiDELETEGiftCardsGiftCardIdRequest
*/
func (a *GiftCardsApiService) DELETEGiftCardsGiftCardId(ctx context.Context, giftCardId string) GiftCardsApiDELETEGiftCardsGiftCardIdRequest {
	return GiftCardsApiDELETEGiftCardsGiftCardIdRequest{
		ApiService: a,
		ctx:        ctx,
		giftCardId: giftCardId,
	}
}

// Execute executes the request
func (a *GiftCardsApiService) DELETEGiftCardsGiftCardIdExecute(r GiftCardsApiDELETEGiftCardsGiftCardIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardsApiService.DELETEGiftCardsGiftCardId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_cards/{giftCardId}"
	localVarPath = strings.Replace(localVarPath, "{"+"giftCardId"+"}", url.PathEscape(parameterToString(r.giftCardId, "")), -1)

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

type GiftCardsApiGETGiftCardsRequest struct {
	ctx        context.Context
	ApiService *GiftCardsApiService
}

func (r GiftCardsApiGETGiftCardsRequest) Execute() (*GETGiftCards200Response, *http.Response, error) {
	return r.ApiService.GETGiftCardsExecute(r)
}

/*
GETGiftCards List all gift cards

List all gift cards

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GiftCardsApiGETGiftCardsRequest
*/
func (a *GiftCardsApiService) GETGiftCards(ctx context.Context) GiftCardsApiGETGiftCardsRequest {
	return GiftCardsApiGETGiftCardsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETGiftCards200Response
func (a *GiftCardsApiService) GETGiftCardsExecute(r GiftCardsApiGETGiftCardsRequest) (*GETGiftCards200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETGiftCards200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardsApiService.GETGiftCards")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_cards"

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

type GiftCardsApiGETGiftCardsGiftCardIdRequest struct {
	ctx        context.Context
	ApiService *GiftCardsApiService
	giftCardId string
}

func (r GiftCardsApiGETGiftCardsGiftCardIdRequest) Execute() (*GETGiftCardsGiftCardId200Response, *http.Response, error) {
	return r.ApiService.GETGiftCardsGiftCardIdExecute(r)
}

/*
GETGiftCardsGiftCardId Retrieve a gift card

Retrieve a gift card

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param giftCardId The resource's id
	@return GiftCardsApiGETGiftCardsGiftCardIdRequest
*/
func (a *GiftCardsApiService) GETGiftCardsGiftCardId(ctx context.Context, giftCardId string) GiftCardsApiGETGiftCardsGiftCardIdRequest {
	return GiftCardsApiGETGiftCardsGiftCardIdRequest{
		ApiService: a,
		ctx:        ctx,
		giftCardId: giftCardId,
	}
}

// Execute executes the request
//
//	@return GETGiftCardsGiftCardId200Response
func (a *GiftCardsApiService) GETGiftCardsGiftCardIdExecute(r GiftCardsApiGETGiftCardsGiftCardIdRequest) (*GETGiftCardsGiftCardId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETGiftCardsGiftCardId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardsApiService.GETGiftCardsGiftCardId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_cards/{giftCardId}"
	localVarPath = strings.Replace(localVarPath, "{"+"giftCardId"+"}", url.PathEscape(parameterToString(r.giftCardId, "")), -1)

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

type GiftCardsApiPATCHGiftCardsGiftCardIdRequest struct {
	ctx            context.Context
	ApiService     *GiftCardsApiService
	giftCardUpdate *GiftCardUpdate
	giftCardId     string
}

func (r GiftCardsApiPATCHGiftCardsGiftCardIdRequest) GiftCardUpdate(giftCardUpdate GiftCardUpdate) GiftCardsApiPATCHGiftCardsGiftCardIdRequest {
	r.giftCardUpdate = &giftCardUpdate
	return r
}

func (r GiftCardsApiPATCHGiftCardsGiftCardIdRequest) Execute() (*PATCHGiftCardsGiftCardId200Response, *http.Response, error) {
	return r.ApiService.PATCHGiftCardsGiftCardIdExecute(r)
}

/*
PATCHGiftCardsGiftCardId Update a gift card

Update a gift card

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param giftCardId The resource's id
	@return GiftCardsApiPATCHGiftCardsGiftCardIdRequest
*/
func (a *GiftCardsApiService) PATCHGiftCardsGiftCardId(ctx context.Context, giftCardId string) GiftCardsApiPATCHGiftCardsGiftCardIdRequest {
	return GiftCardsApiPATCHGiftCardsGiftCardIdRequest{
		ApiService: a,
		ctx:        ctx,
		giftCardId: giftCardId,
	}
}

// Execute executes the request
//
//	@return PATCHGiftCardsGiftCardId200Response
func (a *GiftCardsApiService) PATCHGiftCardsGiftCardIdExecute(r GiftCardsApiPATCHGiftCardsGiftCardIdRequest) (*PATCHGiftCardsGiftCardId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PATCHGiftCardsGiftCardId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardsApiService.PATCHGiftCardsGiftCardId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_cards/{giftCardId}"
	localVarPath = strings.Replace(localVarPath, "{"+"giftCardId"+"}", url.PathEscape(parameterToString(r.giftCardId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.giftCardUpdate == nil {
		return localVarReturnValue, nil, reportError("giftCardUpdate is required and must be specified")
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
	localVarPostBody = r.giftCardUpdate
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

type GiftCardsApiPOSTGiftCardsRequest struct {
	ctx            context.Context
	ApiService     *GiftCardsApiService
	giftCardCreate *GiftCardCreate
}

func (r GiftCardsApiPOSTGiftCardsRequest) GiftCardCreate(giftCardCreate GiftCardCreate) GiftCardsApiPOSTGiftCardsRequest {
	r.giftCardCreate = &giftCardCreate
	return r
}

func (r GiftCardsApiPOSTGiftCardsRequest) Execute() (*POSTGiftCards201Response, *http.Response, error) {
	return r.ApiService.POSTGiftCardsExecute(r)
}

/*
POSTGiftCards Create a gift card

Create a gift card

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GiftCardsApiPOSTGiftCardsRequest
*/
func (a *GiftCardsApiService) POSTGiftCards(ctx context.Context) GiftCardsApiPOSTGiftCardsRequest {
	return GiftCardsApiPOSTGiftCardsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return POSTGiftCards201Response
func (a *GiftCardsApiService) POSTGiftCardsExecute(r GiftCardsApiPOSTGiftCardsRequest) (*POSTGiftCards201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *POSTGiftCards201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GiftCardsApiService.POSTGiftCards")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/gift_cards"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.giftCardCreate == nil {
		return localVarReturnValue, nil, reportError("giftCardCreate is required and must be specified")
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
	localVarPostBody = r.giftCardCreate
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
