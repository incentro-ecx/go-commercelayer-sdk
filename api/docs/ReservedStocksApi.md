# \ReservedStocksApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETReservedStocks**](ReservedStocksApi.md#GETReservedStocks) | **Get** /reserved_stocks | List all reserved stocks
[**GETReservedStocksReservedStockId**](ReservedStocksApi.md#GETReservedStocksReservedStockId) | **Get** /reserved_stocks/{reservedStockId} | Retrieve a reserved stock
[**GETStockItemIdReservedStock**](ReservedStocksApi.md#GETStockItemIdReservedStock) | **Get** /stock_items/{stockItemId}/reserved_stock | Retrieve the reserved stock associated to the stock item
[**GETStockReservationIdReservedStock**](ReservedStocksApi.md#GETStockReservationIdReservedStock) | **Get** /stock_reservations/{stockReservationId}/reserved_stock | Retrieve the reserved stock associated to the stock reservation



## GETReservedStocks

> GETReservedStocks200Response GETReservedStocks(ctx).Execute()

List all reserved stocks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/labd/go-commercelayer-sdk/api"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservedStocksApi.GETReservedStocks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservedStocksApi.GETReservedStocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETReservedStocks`: GETReservedStocks200Response
    fmt.Fprintf(os.Stdout, "Response from `ReservedStocksApi.GETReservedStocks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETReservedStocksRequest struct via the builder pattern


### Return type

[**GETReservedStocks200Response**](GETReservedStocks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETReservedStocksReservedStockId

> GETReservedStocksReservedStockId200Response GETReservedStocksReservedStockId(ctx, reservedStockId).Execute()

Retrieve a reserved stock



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/labd/go-commercelayer-sdk/api"
)

func main() {
    reservedStockId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReservedStocksApi.GETReservedStocksReservedStockId(context.Background(), reservedStockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservedStocksApi.GETReservedStocksReservedStockId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETReservedStocksReservedStockId`: GETReservedStocksReservedStockId200Response
    fmt.Fprintf(os.Stdout, "Response from `ReservedStocksApi.GETReservedStocksReservedStockId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservedStockId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReservedStocksReservedStockIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETReservedStocksReservedStockId200Response**](GETReservedStocksReservedStockId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockItemIdReservedStock

> GETStockItemIdReservedStock(ctx, stockItemId).Execute()

Retrieve the reserved stock associated to the stock item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/labd/go-commercelayer-sdk/api"
)

func main() {
    stockItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReservedStocksApi.GETStockItemIdReservedStock(context.Background(), stockItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservedStocksApi.GETStockItemIdReservedStock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockItemIdReservedStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockReservationIdReservedStock

> GETStockReservationIdReservedStock(ctx, stockReservationId).Execute()

Retrieve the reserved stock associated to the stock reservation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/labd/go-commercelayer-sdk/api"
)

func main() {
    stockReservationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReservedStocksApi.GETStockReservationIdReservedStock(context.Background(), stockReservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReservedStocksApi.GETStockReservationIdReservedStock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockReservationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockReservationIdReservedStockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

