# \StockTransfersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEStockTransfersStockTransferId**](StockTransfersApi.md#DELETEStockTransfersStockTransferId) | **Delete** /stock_transfers/{stockTransferId} | Delete a stock transfer
[**GETLineItemIdStockTransfers**](StockTransfersApi.md#GETLineItemIdStockTransfers) | **Get** /line_items/{lineItemId}/stock_transfers | Retrieve the stock transfers associated to the line item
[**GETOrderIdStockTransfers**](StockTransfersApi.md#GETOrderIdStockTransfers) | **Get** /orders/{orderId}/stock_transfers | Retrieve the stock transfers associated to the order
[**GETShipmentIdStockTransfers**](StockTransfersApi.md#GETShipmentIdStockTransfers) | **Get** /shipments/{shipmentId}/stock_transfers | Retrieve the stock transfers associated to the shipment
[**GETStockLocationIdStockTransfers**](StockTransfersApi.md#GETStockLocationIdStockTransfers) | **Get** /stock_locations/{stockLocationId}/stock_transfers | Retrieve the stock transfers associated to the stock location
[**GETStockReservationIdStockTransfer**](StockTransfersApi.md#GETStockReservationIdStockTransfer) | **Get** /stock_reservations/{stockReservationId}/stock_transfer | Retrieve the stock transfer associated to the stock reservation
[**GETStockTransfers**](StockTransfersApi.md#GETStockTransfers) | **Get** /stock_transfers | List all stock transfers
[**GETStockTransfersStockTransferId**](StockTransfersApi.md#GETStockTransfersStockTransferId) | **Get** /stock_transfers/{stockTransferId} | Retrieve a stock transfer
[**PATCHStockTransfersStockTransferId**](StockTransfersApi.md#PATCHStockTransfersStockTransferId) | **Patch** /stock_transfers/{stockTransferId} | Update a stock transfer
[**POSTStockTransfers**](StockTransfersApi.md#POSTStockTransfers) | **Post** /stock_transfers | Create a stock transfer



## DELETEStockTransfersStockTransferId

> DELETEStockTransfersStockTransferId(ctx, stockTransferId).Execute()

Delete a stock transfer



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
    r, err := apiClient.StockTransfersApi.DELETEStockTransfersStockTransferId(context.Background(), stockTransferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.DELETEStockTransfersStockTransferId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEStockTransfersStockTransferIdRequest struct via the builder pattern


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


## GETLineItemIdStockTransfers

> GETLineItemIdStockTransfers(ctx, lineItemId).Execute()

Retrieve the stock transfers associated to the line item



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
    r, err := apiClient.StockTransfersApi.GETLineItemIdStockTransfers(context.Background(), lineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.GETLineItemIdStockTransfers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETLineItemIdStockTransfersRequest struct via the builder pattern


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


## GETOrderIdStockTransfers

> GETOrderIdStockTransfers(ctx, orderId).Execute()

Retrieve the stock transfers associated to the order



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
    r, err := apiClient.StockTransfersApi.GETOrderIdStockTransfers(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.GETOrderIdStockTransfers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdStockTransfersRequest struct via the builder pattern


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


## GETShipmentIdStockTransfers

> GETShipmentIdStockTransfers(ctx, shipmentId).Execute()

Retrieve the stock transfers associated to the shipment



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
    shipmentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StockTransfersApi.GETShipmentIdStockTransfers(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.GETShipmentIdStockTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShipmentIdStockTransfersRequest struct via the builder pattern


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


## GETStockLocationIdStockTransfers

> GETStockLocationIdStockTransfers(ctx, stockLocationId).Execute()

Retrieve the stock transfers associated to the stock location



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
    stockLocationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StockTransfersApi.GETStockLocationIdStockTransfers(context.Background(), stockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.GETStockLocationIdStockTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockLocationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockLocationIdStockTransfersRequest struct via the builder pattern


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


## GETStockReservationIdStockTransfer

> GETStockReservationIdStockTransfer(ctx, stockReservationId).Execute()

Retrieve the stock transfer associated to the stock reservation



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
    r, err := apiClient.StockTransfersApi.GETStockReservationIdStockTransfer(context.Background(), stockReservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.GETStockReservationIdStockTransfer``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockReservationIdStockTransferRequest struct via the builder pattern


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


## GETStockTransfers

> GETStockTransfers200Response GETStockTransfers(ctx).Execute()

List all stock transfers



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
    resp, r, err := apiClient.StockTransfersApi.GETStockTransfers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.GETStockTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStockTransfers`: GETStockTransfers200Response
    fmt.Fprintf(os.Stdout, "Response from `StockTransfersApi.GETStockTransfers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockTransfersRequest struct via the builder pattern


### Return type

[**GETStockTransfers200Response**](GETStockTransfers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockTransfersStockTransferId

> GETStockTransfersStockTransferId200Response GETStockTransfersStockTransferId(ctx, stockTransferId).Execute()

Retrieve a stock transfer



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
    resp, r, err := apiClient.StockTransfersApi.GETStockTransfersStockTransferId(context.Background(), stockTransferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.GETStockTransfersStockTransferId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStockTransfersStockTransferId`: GETStockTransfersStockTransferId200Response
    fmt.Fprintf(os.Stdout, "Response from `StockTransfersApi.GETStockTransfersStockTransferId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockTransferId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockTransfersStockTransferIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETStockTransfersStockTransferId200Response**](GETStockTransfersStockTransferId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHStockTransfersStockTransferId

> PATCHStockTransfersStockTransferId200Response PATCHStockTransfersStockTransferId(ctx, stockTransferId).StockTransferUpdate(stockTransferUpdate).Execute()

Update a stock transfer



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
    stockTransferUpdate := *openapiclient.NewStockTransferUpdate(*openapiclient.NewStockTransferUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHStockTransfersStockTransferId200ResponseDataAttributes())) // StockTransferUpdate | 
    stockTransferId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockTransfersApi.PATCHStockTransfersStockTransferId(context.Background(), stockTransferId).StockTransferUpdate(stockTransferUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.PATCHStockTransfersStockTransferId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHStockTransfersStockTransferId`: PATCHStockTransfersStockTransferId200Response
    fmt.Fprintf(os.Stdout, "Response from `StockTransfersApi.PATCHStockTransfersStockTransferId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockTransferId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHStockTransfersStockTransferIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockTransferUpdate** | [**StockTransferUpdate**](StockTransferUpdate.md) |  | 


### Return type

[**PATCHStockTransfersStockTransferId200Response**](PATCHStockTransfersStockTransferId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTStockTransfers

> POSTStockTransfers201Response POSTStockTransfers(ctx).StockTransferCreate(stockTransferCreate).Execute()

Create a stock transfer



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
    stockTransferCreate := *openapiclient.NewStockTransferCreate(*openapiclient.NewStockTransferCreateData(interface{}(123), *openapiclient.NewPOSTStockTransfers201ResponseDataAttributes(interface{}(2)))) // StockTransferCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockTransfersApi.POSTStockTransfers(context.Background()).StockTransferCreate(stockTransferCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockTransfersApi.POSTStockTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTStockTransfers`: POSTStockTransfers201Response
    fmt.Fprintf(os.Stdout, "Response from `StockTransfersApi.POSTStockTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTStockTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockTransferCreate** | [**StockTransferCreate**](StockTransferCreate.md) |  | 

### Return type

[**POSTStockTransfers201Response**](POSTStockTransfers201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

