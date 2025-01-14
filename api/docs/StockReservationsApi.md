# \StockReservationsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEStockReservationsStockReservationId**](StockReservationsApi.md#DELETEStockReservationsStockReservationId) | **Delete** /stock_reservations/{stockReservationId} | Delete a stock reservation
[**GETLineItemIdStockReservations**](StockReservationsApi.md#GETLineItemIdStockReservations) | **Get** /line_items/{lineItemId}/stock_reservations | Retrieve the stock reservations associated to the line item
[**GETOrderIdStockReservations**](StockReservationsApi.md#GETOrderIdStockReservations) | **Get** /orders/{orderId}/stock_reservations | Retrieve the stock reservations associated to the order
[**GETReservedStockIdStockReservations**](StockReservationsApi.md#GETReservedStockIdStockReservations) | **Get** /reserved_stocks/{reservedStockId}/stock_reservations | Retrieve the stock reservations associated to the reserved stock
[**GETSkuIdStockReservations**](StockReservationsApi.md#GETSkuIdStockReservations) | **Get** /skus/{skuId}/stock_reservations | Retrieve the stock reservations associated to the SKU
[**GETStockItemIdStockReservations**](StockReservationsApi.md#GETStockItemIdStockReservations) | **Get** /stock_items/{stockItemId}/stock_reservations | Retrieve the stock reservations associated to the stock item
[**GETStockLineItemIdStockReservation**](StockReservationsApi.md#GETStockLineItemIdStockReservation) | **Get** /stock_line_items/{stockLineItemId}/stock_reservation | Retrieve the stock reservation associated to the stock line item
[**GETStockReservations**](StockReservationsApi.md#GETStockReservations) | **Get** /stock_reservations | List all stock reservations
[**GETStockReservationsStockReservationId**](StockReservationsApi.md#GETStockReservationsStockReservationId) | **Get** /stock_reservations/{stockReservationId} | Retrieve a stock reservation
[**GETStockTransferIdStockReservation**](StockReservationsApi.md#GETStockTransferIdStockReservation) | **Get** /stock_transfers/{stockTransferId}/stock_reservation | Retrieve the stock reservation associated to the stock transfer
[**PATCHStockReservationsStockReservationId**](StockReservationsApi.md#PATCHStockReservationsStockReservationId) | **Patch** /stock_reservations/{stockReservationId} | Update a stock reservation
[**POSTStockReservations**](StockReservationsApi.md#POSTStockReservations) | **Post** /stock_reservations | Create a stock reservation



## DELETEStockReservationsStockReservationId

> DELETEStockReservationsStockReservationId(ctx, stockReservationId).Execute()

Delete a stock reservation



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
    r, err := apiClient.StockReservationsApi.DELETEStockReservationsStockReservationId(context.Background(), stockReservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.DELETEStockReservationsStockReservationId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEStockReservationsStockReservationIdRequest struct via the builder pattern


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


## GETLineItemIdStockReservations

> GETLineItemIdStockReservations(ctx, lineItemId).Execute()

Retrieve the stock reservations associated to the line item



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
    lineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StockReservationsApi.GETLineItemIdStockReservations(context.Background(), lineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.GETLineItemIdStockReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemIdStockReservationsRequest struct via the builder pattern


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


## GETOrderIdStockReservations

> GETOrderIdStockReservations(ctx, orderId).Execute()

Retrieve the stock reservations associated to the order



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
    orderId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StockReservationsApi.GETOrderIdStockReservations(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.GETOrderIdStockReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderIdStockReservationsRequest struct via the builder pattern


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


## GETReservedStockIdStockReservations

> GETReservedStockIdStockReservations(ctx, reservedStockId).Execute()

Retrieve the stock reservations associated to the reserved stock



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
    r, err := apiClient.StockReservationsApi.GETReservedStockIdStockReservations(context.Background(), reservedStockId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.GETReservedStockIdStockReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reservedStockId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReservedStockIdStockReservationsRequest struct via the builder pattern


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


## GETSkuIdStockReservations

> GETSkuIdStockReservations(ctx, skuId).Execute()

Retrieve the stock reservations associated to the SKU



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
    skuId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StockReservationsApi.GETSkuIdStockReservations(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.GETSkuIdStockReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuIdStockReservationsRequest struct via the builder pattern


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


## GETStockItemIdStockReservations

> GETStockItemIdStockReservations(ctx, stockItemId).Execute()

Retrieve the stock reservations associated to the stock item



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
    r, err := apiClient.StockReservationsApi.GETStockItemIdStockReservations(context.Background(), stockItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.GETStockItemIdStockReservations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockItemIdStockReservationsRequest struct via the builder pattern


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


## GETStockLineItemIdStockReservation

> GETStockLineItemIdStockReservation(ctx, stockLineItemId).Execute()

Retrieve the stock reservation associated to the stock line item



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
    stockLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StockReservationsApi.GETStockLineItemIdStockReservation(context.Background(), stockLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.GETStockLineItemIdStockReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockLineItemIdStockReservationRequest struct via the builder pattern


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


## GETStockReservations

> GETStockReservations200Response GETStockReservations(ctx).Execute()

List all stock reservations



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
    resp, r, err := apiClient.StockReservationsApi.GETStockReservations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.GETStockReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStockReservations`: GETStockReservations200Response
    fmt.Fprintf(os.Stdout, "Response from `StockReservationsApi.GETStockReservations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockReservationsRequest struct via the builder pattern


### Return type

[**GETStockReservations200Response**](GETStockReservations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockReservationsStockReservationId

> GETStockReservationsStockReservationId200Response GETStockReservationsStockReservationId(ctx, stockReservationId).Execute()

Retrieve a stock reservation



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
    resp, r, err := apiClient.StockReservationsApi.GETStockReservationsStockReservationId(context.Background(), stockReservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.GETStockReservationsStockReservationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStockReservationsStockReservationId`: GETStockReservationsStockReservationId200Response
    fmt.Fprintf(os.Stdout, "Response from `StockReservationsApi.GETStockReservationsStockReservationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockReservationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockReservationsStockReservationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETStockReservationsStockReservationId200Response**](GETStockReservationsStockReservationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockTransferIdStockReservation

> GETStockTransferIdStockReservation(ctx, stockTransferId).Execute()

Retrieve the stock reservation associated to the stock transfer



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
    stockTransferId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StockReservationsApi.GETStockTransferIdStockReservation(context.Background(), stockTransferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.GETStockTransferIdStockReservation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockTransferId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockTransferIdStockReservationRequest struct via the builder pattern


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


## PATCHStockReservationsStockReservationId

> PATCHStockReservationsStockReservationId200Response PATCHStockReservationsStockReservationId(ctx, stockReservationId).StockReservationUpdate(stockReservationUpdate).Execute()

Update a stock reservation



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
    stockReservationUpdate := *openapiclient.NewStockReservationUpdate(*openapiclient.NewStockReservationUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHStockReservationsStockReservationId200ResponseDataAttributes())) // StockReservationUpdate | 
    stockReservationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockReservationsApi.PATCHStockReservationsStockReservationId(context.Background(), stockReservationId).StockReservationUpdate(stockReservationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.PATCHStockReservationsStockReservationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHStockReservationsStockReservationId`: PATCHStockReservationsStockReservationId200Response
    fmt.Fprintf(os.Stdout, "Response from `StockReservationsApi.PATCHStockReservationsStockReservationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockReservationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHStockReservationsStockReservationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockReservationUpdate** | [**StockReservationUpdate**](StockReservationUpdate.md) |  | 


### Return type

[**PATCHStockReservationsStockReservationId200Response**](PATCHStockReservationsStockReservationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTStockReservations

> POSTStockReservations201Response POSTStockReservations(ctx).StockReservationCreate(stockReservationCreate).Execute()

Create a stock reservation



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
    stockReservationCreate := *openapiclient.NewStockReservationCreate(*openapiclient.NewStockReservationCreateData(interface{}(123), *openapiclient.NewPOSTStockReservations201ResponseDataAttributes(interface{}(4)))) // StockReservationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockReservationsApi.POSTStockReservations(context.Background()).StockReservationCreate(stockReservationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockReservationsApi.POSTStockReservations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTStockReservations`: POSTStockReservations201Response
    fmt.Fprintf(os.Stdout, "Response from `StockReservationsApi.POSTStockReservations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTStockReservationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockReservationCreate** | [**StockReservationCreate**](StockReservationCreate.md) |  | 

### Return type

[**POSTStockReservations201Response**](POSTStockReservations201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

