# \LineItemsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETELineItemsLineItemId**](LineItemsApi.md#DELETELineItemsLineItemId) | **Delete** /line_items/{lineItemId} | Delete a line item
[**GETLineItemOptionIdLineItem**](LineItemsApi.md#GETLineItemOptionIdLineItem) | **Get** /line_item_options/{lineItemOptionId}/line_item | Retrieve the line item associated to the line item option
[**GETLineItems**](LineItemsApi.md#GETLineItems) | **Get** /line_items | List all line items
[**GETLineItemsLineItemId**](LineItemsApi.md#GETLineItemsLineItemId) | **Get** /line_items/{lineItemId} | Retrieve a line item
[**GETOrderIdLineItems**](LineItemsApi.md#GETOrderIdLineItems) | **Get** /orders/{orderId}/line_items | Retrieve the line items associated to the order
[**GETOrderSubscriptionItemIdSourceLineItem**](LineItemsApi.md#GETOrderSubscriptionItemIdSourceLineItem) | **Get** /order_subscription_items/{orderSubscriptionItemId}/source_line_item | Retrieve the source line item associated to the order subscription item
[**GETReturnLineItemIdLineItem**](LineItemsApi.md#GETReturnLineItemIdLineItem) | **Get** /return_line_items/{returnLineItemId}/line_item | Retrieve the line item associated to the return line item
[**GETShipmentIdLineItems**](LineItemsApi.md#GETShipmentIdLineItems) | **Get** /shipments/{shipmentId}/line_items | Retrieve the line items associated to the shipment
[**GETStockLineItemIdLineItem**](LineItemsApi.md#GETStockLineItemIdLineItem) | **Get** /stock_line_items/{stockLineItemId}/line_item | Retrieve the line item associated to the stock line item
[**GETStockReservationIdLineItem**](LineItemsApi.md#GETStockReservationIdLineItem) | **Get** /stock_reservations/{stockReservationId}/line_item | Retrieve the line item associated to the stock reservation
[**GETStockTransferIdLineItem**](LineItemsApi.md#GETStockTransferIdLineItem) | **Get** /stock_transfers/{stockTransferId}/line_item | Retrieve the line item associated to the stock transfer
[**PATCHLineItemsLineItemId**](LineItemsApi.md#PATCHLineItemsLineItemId) | **Patch** /line_items/{lineItemId} | Update a line item
[**POSTLineItems**](LineItemsApi.md#POSTLineItems) | **Post** /line_items | Create a line item



## DELETELineItemsLineItemId

> DELETELineItemsLineItemId(ctx, lineItemId).Execute()

Delete a line item



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
    r, err := apiClient.LineItemsApi.DELETELineItemsLineItemId(context.Background(), lineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.DELETELineItemsLineItemId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETELineItemsLineItemIdRequest struct via the builder pattern


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


## GETLineItemOptionIdLineItem

> GETLineItemOptionIdLineItem(ctx, lineItemOptionId).Execute()

Retrieve the line item associated to the line item option



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
    lineItemOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LineItemsApi.GETLineItemOptionIdLineItem(context.Background(), lineItemOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETLineItemOptionIdLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemOptionIdLineItemRequest struct via the builder pattern


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


## GETLineItems

> GETLineItems200Response GETLineItems(ctx).Execute()

List all line items



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
    resp, r, err := apiClient.LineItemsApi.GETLineItems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETLineItems`: GETLineItems200Response
    fmt.Fprintf(os.Stdout, "Response from `LineItemsApi.GETLineItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemsRequest struct via the builder pattern


### Return type

[**GETLineItems200Response**](GETLineItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETLineItemsLineItemId

> GETLineItemsLineItemId200Response GETLineItemsLineItemId(ctx, lineItemId).Execute()

Retrieve a line item



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
    resp, r, err := apiClient.LineItemsApi.GETLineItemsLineItemId(context.Background(), lineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETLineItemsLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETLineItemsLineItemId`: GETLineItemsLineItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `LineItemsApi.GETLineItemsLineItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemsLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETLineItemsLineItemId200Response**](GETLineItemsLineItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderIdLineItems

> GETOrderIdLineItems(ctx, orderId).Execute()

Retrieve the line items associated to the order



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
    r, err := apiClient.LineItemsApi.GETOrderIdLineItems(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETOrderIdLineItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdLineItemsRequest struct via the builder pattern


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


## GETOrderSubscriptionItemIdSourceLineItem

> GETOrderSubscriptionItemIdSourceLineItem(ctx, orderSubscriptionItemId).Execute()

Retrieve the source line item associated to the order subscription item



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
    orderSubscriptionItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LineItemsApi.GETOrderSubscriptionItemIdSourceLineItem(context.Background(), orderSubscriptionItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETOrderSubscriptionItemIdSourceLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderSubscriptionItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETOrderSubscriptionItemIdSourceLineItemRequest struct via the builder pattern


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


## GETReturnLineItemIdLineItem

> GETReturnLineItemIdLineItem(ctx, returnLineItemId).Execute()

Retrieve the line item associated to the return line item



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
    returnLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LineItemsApi.GETReturnLineItemIdLineItem(context.Background(), returnLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETReturnLineItemIdLineItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnLineItemIdLineItemRequest struct via the builder pattern


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


## GETShipmentIdLineItems

> GETShipmentIdLineItems(ctx, shipmentId).Execute()

Retrieve the line items associated to the shipment



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
    r, err := apiClient.LineItemsApi.GETShipmentIdLineItems(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETShipmentIdLineItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShipmentIdLineItemsRequest struct via the builder pattern


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


## GETStockLineItemIdLineItem

> GETStockLineItemIdLineItem(ctx, stockLineItemId).Execute()

Retrieve the line item associated to the stock line item



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
    r, err := apiClient.LineItemsApi.GETStockLineItemIdLineItem(context.Background(), stockLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETStockLineItemIdLineItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockLineItemIdLineItemRequest struct via the builder pattern


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


## GETStockReservationIdLineItem

> GETStockReservationIdLineItem(ctx, stockReservationId).Execute()

Retrieve the line item associated to the stock reservation



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
    r, err := apiClient.LineItemsApi.GETStockReservationIdLineItem(context.Background(), stockReservationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETStockReservationIdLineItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockReservationIdLineItemRequest struct via the builder pattern


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


## GETStockTransferIdLineItem

> GETStockTransferIdLineItem(ctx, stockTransferId).Execute()

Retrieve the line item associated to the stock transfer



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
    r, err := apiClient.LineItemsApi.GETStockTransferIdLineItem(context.Background(), stockTransferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.GETStockTransferIdLineItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockTransferIdLineItemRequest struct via the builder pattern


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


## PATCHLineItemsLineItemId

> PATCHLineItemsLineItemId200Response PATCHLineItemsLineItemId(ctx, lineItemId).LineItemUpdate(lineItemUpdate).Execute()

Update a line item



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
    lineItemUpdate := *openapiclient.NewLineItemUpdate(*openapiclient.NewLineItemUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHLineItemsLineItemId200ResponseDataAttributes())) // LineItemUpdate | 
    lineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemsApi.PATCHLineItemsLineItemId(context.Background(), lineItemId).LineItemUpdate(lineItemUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.PATCHLineItemsLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHLineItemsLineItemId`: PATCHLineItemsLineItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `LineItemsApi.PATCHLineItemsLineItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHLineItemsLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lineItemUpdate** | [**LineItemUpdate**](LineItemUpdate.md) |  | 


### Return type

[**PATCHLineItemsLineItemId200Response**](PATCHLineItemsLineItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTLineItems

> POSTLineItems201Response POSTLineItems(ctx).LineItemCreate(lineItemCreate).Execute()

Create a line item



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
    lineItemCreate := *openapiclient.NewLineItemCreate(*openapiclient.NewLineItemCreateData(interface{}(123), *openapiclient.NewPOSTLineItems201ResponseDataAttributes(interface{}(4)))) // LineItemCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LineItemsApi.POSTLineItems(context.Background()).LineItemCreate(lineItemCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LineItemsApi.POSTLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTLineItems`: POSTLineItems201Response
    fmt.Fprintf(os.Stdout, "Response from `LineItemsApi.POSTLineItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTLineItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lineItemCreate** | [**LineItemCreate**](LineItemCreate.md) |  | 

### Return type

[**POSTLineItems201Response**](POSTLineItems201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

