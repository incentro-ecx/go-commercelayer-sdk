# \StockLineItemsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEStockLineItemsStockLineItemId**](StockLineItemsApi.md#DELETEStockLineItemsStockLineItemId) | **Delete** /stock_line_items/{stockLineItemId} | Delete a stock line item
[**GETLineItemIdStockLineItems**](StockLineItemsApi.md#GETLineItemIdStockLineItems) | **Get** /line_items/{lineItemId}/stock_line_items | Retrieve the stock line items associated to the line item
[**GETOrderIdStockLineItems**](StockLineItemsApi.md#GETOrderIdStockLineItems) | **Get** /orders/{orderId}/stock_line_items | Retrieve the stock line items associated to the order
[**GETParcelLineItemIdStockLineItem**](StockLineItemsApi.md#GETParcelLineItemIdStockLineItem) | **Get** /parcel_line_items/{parcelLineItemId}/stock_line_item | Retrieve the stock line item associated to the parcel line item
[**GETShipmentIdStockLineItems**](StockLineItemsApi.md#GETShipmentIdStockLineItems) | **Get** /shipments/{shipmentId}/stock_line_items | Retrieve the stock line items associated to the shipment
[**GETStockLineItems**](StockLineItemsApi.md#GETStockLineItems) | **Get** /stock_line_items | List all stock line items
[**GETStockLineItemsStockLineItemId**](StockLineItemsApi.md#GETStockLineItemsStockLineItemId) | **Get** /stock_line_items/{stockLineItemId} | Retrieve a stock line item
[**GETStockReservationIdStockLineItem**](StockLineItemsApi.md#GETStockReservationIdStockLineItem) | **Get** /stock_reservations/{stockReservationId}/stock_line_item | Retrieve the stock line item associated to the stock reservation
[**PATCHStockLineItemsStockLineItemId**](StockLineItemsApi.md#PATCHStockLineItemsStockLineItemId) | **Patch** /stock_line_items/{stockLineItemId} | Update a stock line item
[**POSTStockLineItems**](StockLineItemsApi.md#POSTStockLineItems) | **Post** /stock_line_items | Create a stock line item



## DELETEStockLineItemsStockLineItemId

> DELETEStockLineItemsStockLineItemId(ctx, stockLineItemId).Execute()

Delete a stock line item



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
    r, err := apiClient.StockLineItemsApi.DELETEStockLineItemsStockLineItemId(context.Background(), stockLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.DELETEStockLineItemsStockLineItemId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEStockLineItemsStockLineItemIdRequest struct via the builder pattern


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


## GETLineItemIdStockLineItems

> GETLineItemIdStockLineItems(ctx, lineItemId).Execute()

Retrieve the stock line items associated to the line item



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
    r, err := apiClient.StockLineItemsApi.GETLineItemIdStockLineItems(context.Background(), lineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.GETLineItemIdStockLineItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETLineItemIdStockLineItemsRequest struct via the builder pattern


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


## GETOrderIdStockLineItems

> GETOrderIdStockLineItems(ctx, orderId).Execute()

Retrieve the stock line items associated to the order



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
    r, err := apiClient.StockLineItemsApi.GETOrderIdStockLineItems(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.GETOrderIdStockLineItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdStockLineItemsRequest struct via the builder pattern


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


## GETParcelLineItemIdStockLineItem

> GETParcelLineItemIdStockLineItem(ctx, parcelLineItemId).Execute()

Retrieve the stock line item associated to the parcel line item



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
    parcelLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StockLineItemsApi.GETParcelLineItemIdStockLineItem(context.Background(), parcelLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.GETParcelLineItemIdStockLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parcelLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETParcelLineItemIdStockLineItemRequest struct via the builder pattern


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


## GETShipmentIdStockLineItems

> GETShipmentIdStockLineItems(ctx, shipmentId).Execute()

Retrieve the stock line items associated to the shipment



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
    r, err := apiClient.StockLineItemsApi.GETShipmentIdStockLineItems(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.GETShipmentIdStockLineItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShipmentIdStockLineItemsRequest struct via the builder pattern


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


## GETStockLineItems

> GETStockLineItems200Response GETStockLineItems(ctx).Execute()

List all stock line items



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
    resp, r, err := apiClient.StockLineItemsApi.GETStockLineItems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.GETStockLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStockLineItems`: GETStockLineItems200Response
    fmt.Fprintf(os.Stdout, "Response from `StockLineItemsApi.GETStockLineItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockLineItemsRequest struct via the builder pattern


### Return type

[**GETStockLineItems200Response**](GETStockLineItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockLineItemsStockLineItemId

> GETStockLineItemsStockLineItemId200Response GETStockLineItemsStockLineItemId(ctx, stockLineItemId).Execute()

Retrieve a stock line item



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
    resp, r, err := apiClient.StockLineItemsApi.GETStockLineItemsStockLineItemId(context.Background(), stockLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.GETStockLineItemsStockLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStockLineItemsStockLineItemId`: GETStockLineItemsStockLineItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `StockLineItemsApi.GETStockLineItemsStockLineItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStockLineItemsStockLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETStockLineItemsStockLineItemId200Response**](GETStockLineItemsStockLineItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockReservationIdStockLineItem

> GETStockReservationIdStockLineItem(ctx, stockReservationId).Execute()

Retrieve the stock line item associated to the stock reservation



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
    r, err := apiClient.StockLineItemsApi.GETStockReservationIdStockLineItem(context.Background(), stockReservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.GETStockReservationIdStockLineItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockReservationIdStockLineItemRequest struct via the builder pattern


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


## PATCHStockLineItemsStockLineItemId

> PATCHStockLineItemsStockLineItemId200Response PATCHStockLineItemsStockLineItemId(ctx, stockLineItemId).StockLineItemUpdate(stockLineItemUpdate).Execute()

Update a stock line item



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
    stockLineItemUpdate := *openapiclient.NewStockLineItemUpdate(*openapiclient.NewStockLineItemUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHStockLineItemsStockLineItemId200ResponseDataAttributes())) // StockLineItemUpdate | 
    stockLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLineItemsApi.PATCHStockLineItemsStockLineItemId(context.Background(), stockLineItemId).StockLineItemUpdate(stockLineItemUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.PATCHStockLineItemsStockLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHStockLineItemsStockLineItemId`: PATCHStockLineItemsStockLineItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `StockLineItemsApi.PATCHStockLineItemsStockLineItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stockLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHStockLineItemsStockLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockLineItemUpdate** | [**StockLineItemUpdate**](StockLineItemUpdate.md) |  | 


### Return type

[**PATCHStockLineItemsStockLineItemId200Response**](PATCHStockLineItemsStockLineItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTStockLineItems

> POSTStockLineItems201Response POSTStockLineItems(ctx).StockLineItemCreate(stockLineItemCreate).Execute()

Create a stock line item



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
    stockLineItemCreate := *openapiclient.NewStockLineItemCreate(*openapiclient.NewStockLineItemCreateData(interface{}(123), *openapiclient.NewPOSTStockLineItems201ResponseDataAttributes(interface{}(4)))) // StockLineItemCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StockLineItemsApi.POSTStockLineItems(context.Background()).StockLineItemCreate(stockLineItemCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StockLineItemsApi.POSTStockLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTStockLineItems`: POSTStockLineItems201Response
    fmt.Fprintf(os.Stdout, "Response from `StockLineItemsApi.POSTStockLineItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTStockLineItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stockLineItemCreate** | [**StockLineItemCreate**](StockLineItemCreate.md) |  | 

### Return type

[**POSTStockLineItems201Response**](POSTStockLineItems201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

