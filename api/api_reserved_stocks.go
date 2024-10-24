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

// ReservedStocksApiService ReservedStocksApi service
type ReservedStocksApiService service

type ReservedStocksApiGETReservedStocksRequest struct {
	ctx        context.Context
	ApiService *ReservedStocksApiService
}

func (r ReservedStocksApiGETReservedStocksRequest) Execute() (*GETReservedStocks200Response, *http.Response, error) {
	return r.ApiService.GETReservedStocksExecute(r)
}

/*
GETReservedStocks List all reserved stocks

List all reserved stocks

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ReservedStocksApiGETReservedStocksRequest
*/
func (a *ReservedStocksApiService) GETReservedStocks(ctx context.Context) ReservedStocksApiGETReservedStocksRequest {
	return ReservedStocksApiGETReservedStocksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GETReservedStocks200Response
func (a *ReservedStocksApiService) GETReservedStocksExecute(r ReservedStocksApiGETReservedStocksRequest) (*GETReservedStocks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETReservedStocks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservedStocksApiService.GETReservedStocks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reserved_stocks"

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

type ReservedStocksApiGETReservedStocksReservedStockIdRequest struct {
	ctx             context.Context
	ApiService      *ReservedStocksApiService
	reservedStockId interface{}
}

func (r ReservedStocksApiGETReservedStocksReservedStockIdRequest) Execute() (*GETReservedStocksReservedStockId200Response, *http.Response, error) {
	return r.ApiService.GETReservedStocksReservedStockIdExecute(r)
}

/*
GETReservedStocksReservedStockId Retrieve a reserved stock

Retrieve a reserved stock

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reservedStockId The resource's id
	@return ReservedStocksApiGETReservedStocksReservedStockIdRequest
*/
func (a *ReservedStocksApiService) GETReservedStocksReservedStockId(ctx context.Context, reservedStockId interface{}) ReservedStocksApiGETReservedStocksReservedStockIdRequest {
	return ReservedStocksApiGETReservedStocksReservedStockIdRequest{
		ApiService:      a,
		ctx:             ctx,
		reservedStockId: reservedStockId,
	}
}

// Execute executes the request
//
//	@return GETReservedStocksReservedStockId200Response
func (a *ReservedStocksApiService) GETReservedStocksReservedStockIdExecute(r ReservedStocksApiGETReservedStocksReservedStockIdRequest) (*GETReservedStocksReservedStockId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GETReservedStocksReservedStockId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservedStocksApiService.GETReservedStocksReservedStockId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reserved_stocks/{reservedStockId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reservedStockId"+"}", url.PathEscape(parameterValueToString(r.reservedStockId, "reservedStockId")), -1)

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

type ReservedStocksApiGETStockItemIdReservedStockRequest struct {
	ctx         context.Context
	ApiService  *ReservedStocksApiService
	stockItemId interface{}
}

func (r ReservedStocksApiGETStockItemIdReservedStockRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockItemIdReservedStockExecute(r)
}

/*
GETStockItemIdReservedStock Retrieve the reserved stock associated to the stock item

Retrieve the reserved stock associated to the stock item

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockItemId The resource's id
	@return ReservedStocksApiGETStockItemIdReservedStockRequest
*/
func (a *ReservedStocksApiService) GETStockItemIdReservedStock(ctx context.Context, stockItemId interface{}) ReservedStocksApiGETStockItemIdReservedStockRequest {
	return ReservedStocksApiGETStockItemIdReservedStockRequest{
		ApiService:  a,
		ctx:         ctx,
		stockItemId: stockItemId,
	}
}

// Execute executes the request
func (a *ReservedStocksApiService) GETStockItemIdReservedStockExecute(r ReservedStocksApiGETStockItemIdReservedStockRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservedStocksApiService.GETStockItemIdReservedStock")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_items/{stockItemId}/reserved_stock"
	localVarPath = strings.Replace(localVarPath, "{"+"stockItemId"+"}", url.PathEscape(parameterValueToString(r.stockItemId, "stockItemId")), -1)

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

type ReservedStocksApiGETStockReservationIdReservedStockRequest struct {
	ctx                context.Context
	ApiService         *ReservedStocksApiService
	stockReservationId interface{}
}

func (r ReservedStocksApiGETStockReservationIdReservedStockRequest) Execute() (*http.Response, error) {
	return r.ApiService.GETStockReservationIdReservedStockExecute(r)
}

/*
GETStockReservationIdReservedStock Retrieve the reserved stock associated to the stock reservation

Retrieve the reserved stock associated to the stock reservation

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param stockReservationId The resource's id
	@return ReservedStocksApiGETStockReservationIdReservedStockRequest
*/
func (a *ReservedStocksApiService) GETStockReservationIdReservedStock(ctx context.Context, stockReservationId interface{}) ReservedStocksApiGETStockReservationIdReservedStockRequest {
	return ReservedStocksApiGETStockReservationIdReservedStockRequest{
		ApiService:         a,
		ctx:                ctx,
		stockReservationId: stockReservationId,
	}
}

// Execute executes the request
func (a *ReservedStocksApiService) GETStockReservationIdReservedStockExecute(r ReservedStocksApiGETStockReservationIdReservedStockRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReservedStocksApiService.GETStockReservationIdReservedStock")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stock_reservations/{stockReservationId}/reserved_stock"
	localVarPath = strings.Replace(localVarPath, "{"+"stockReservationId"+"}", url.PathEscape(parameterValueToString(r.stockReservationId, "stockReservationId")), -1)

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
