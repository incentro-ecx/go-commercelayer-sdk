# \ShipmentsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEShipmentsShipmentId**](ShipmentsApi.md#DELETEShipmentsShipmentId) | **Delete** /shipments/{shipmentId} | Delete a shipment
[**GETOrderIdShipments**](ShipmentsApi.md#GETOrderIdShipments) | **Get** /orders/{orderId}/shipments | Retrieve the shipments associated to the order
[**GETParcelIdShipment**](ShipmentsApi.md#GETParcelIdShipment) | **Get** /parcels/{parcelId}/shipment | Retrieve the shipment associated to the parcel
[**GETShipments**](ShipmentsApi.md#GETShipments) | **Get** /shipments | List all shipments
[**GETShipmentsShipmentId**](ShipmentsApi.md#GETShipmentsShipmentId) | **Get** /shipments/{shipmentId} | Retrieve a shipment
[**GETStockLineItemIdShipment**](ShipmentsApi.md#GETStockLineItemIdShipment) | **Get** /stock_line_items/{stockLineItemId}/shipment | Retrieve the shipment associated to the stock line item
[**GETStockTransferIdShipment**](ShipmentsApi.md#GETStockTransferIdShipment) | **Get** /stock_transfers/{stockTransferId}/shipment | Retrieve the shipment associated to the stock transfer
[**PATCHShipmentsShipmentId**](ShipmentsApi.md#PATCHShipmentsShipmentId) | **Patch** /shipments/{shipmentId} | Update a shipment
[**POSTShipments**](ShipmentsApi.md#POSTShipments) | **Post** /shipments | Create a shipment



## DELETEShipmentsShipmentId

> DELETEShipmentsShipmentId(ctx, shipmentId).Execute()

Delete a shipment



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
    r, err := apiClient.ShipmentsApi.DELETEShipmentsShipmentId(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsApi.DELETEShipmentsShipmentId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEShipmentsShipmentIdRequest struct via the builder pattern


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


## GETOrderIdShipments

> GETOrderIdShipments(ctx, orderId).Execute()

Retrieve the shipments associated to the order



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
    r, err := apiClient.ShipmentsApi.GETOrderIdShipments(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsApi.GETOrderIdShipments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdShipmentsRequest struct via the builder pattern


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


## GETParcelIdShipment

> GETParcelIdShipment(ctx, parcelId).Execute()

Retrieve the shipment associated to the parcel



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
    parcelId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShipmentsApi.GETParcelIdShipment(context.Background(), parcelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsApi.GETParcelIdShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parcelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETParcelIdShipmentRequest struct via the builder pattern


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


## GETShipments

> GETShipments200Response GETShipments(ctx).Execute()

List all shipments



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
    resp, r, err := apiClient.ShipmentsApi.GETShipments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsApi.GETShipments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShipments`: GETShipments200Response
    fmt.Fprintf(os.Stdout, "Response from `ShipmentsApi.GETShipments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETShipmentsRequest struct via the builder pattern


### Return type

[**GETShipments200Response**](GETShipments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShipmentsShipmentId

> GETShipmentsShipmentId200Response GETShipmentsShipmentId(ctx, shipmentId).Execute()

Retrieve a shipment



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
    resp, r, err := apiClient.ShipmentsApi.GETShipmentsShipmentId(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsApi.GETShipmentsShipmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShipmentsShipmentId`: GETShipmentsShipmentId200Response
    fmt.Fprintf(os.Stdout, "Response from `ShipmentsApi.GETShipmentsShipmentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShipmentsShipmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETShipmentsShipmentId200Response**](GETShipmentsShipmentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockLineItemIdShipment

> GETStockLineItemIdShipment(ctx, stockLineItemId).Execute()

Retrieve the shipment associated to the stock line item



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
    r, err := apiClient.ShipmentsApi.GETStockLineItemIdShipment(context.Background(), stockLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsApi.GETStockLineItemIdShipment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockLineItemIdShipmentRequest struct via the builder pattern


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


## GETStockTransferIdShipment

> GETStockTransferIdShipment(ctx, stockTransferId).Execute()

Retrieve the shipment associated to the stock transfer



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
    r, err := apiClient.ShipmentsApi.GETStockTransferIdShipment(context.Background(), stockTransferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsApi.GETStockTransferIdShipment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockTransferIdShipmentRequest struct via the builder pattern


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


## PATCHShipmentsShipmentId

> PATCHShipmentsShipmentId200Response PATCHShipmentsShipmentId(ctx, shipmentId).ShipmentUpdate(shipmentUpdate).Execute()

Update a shipment



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
    shipmentUpdate := *openapiclient.NewShipmentUpdate(*openapiclient.NewShipmentUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHShipmentsShipmentId200ResponseDataAttributes())) // ShipmentUpdate | 
    shipmentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShipmentsApi.PATCHShipmentsShipmentId(context.Background(), shipmentId).ShipmentUpdate(shipmentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsApi.PATCHShipmentsShipmentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHShipmentsShipmentId`: PATCHShipmentsShipmentId200Response
    fmt.Fprintf(os.Stdout, "Response from `ShipmentsApi.PATCHShipmentsShipmentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHShipmentsShipmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentUpdate** | [**ShipmentUpdate**](ShipmentUpdate.md) |  | 


### Return type

[**PATCHShipmentsShipmentId200Response**](PATCHShipmentsShipmentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTShipments

> POSTShipments201Response POSTShipments(ctx).ShipmentCreate(shipmentCreate).Execute()

Create a shipment



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
    shipmentCreate := *openapiclient.NewShipmentCreate(*openapiclient.NewShipmentCreateData(interface{}(123), *openapiclient.NewPOSTAdyenPayments201ResponseDataAttributes())) // ShipmentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShipmentsApi.POSTShipments(context.Background()).ShipmentCreate(shipmentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShipmentsApi.POSTShipments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTShipments`: POSTShipments201Response
    fmt.Fprintf(os.Stdout, "Response from `ShipmentsApi.POSTShipments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTShipmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentCreate** | [**ShipmentCreate**](ShipmentCreate.md) |  | 

### Return type

[**POSTShipments201Response**](POSTShipments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

