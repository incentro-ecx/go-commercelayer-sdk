# \InventoryStockLocationsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEInventoryStockLocationsInventoryStockLocationId**](InventoryStockLocationsApi.md#DELETEInventoryStockLocationsInventoryStockLocationId) | **Delete** /inventory_stock_locations/{inventoryStockLocationId} | Delete an inventory stock location
[**GETInventoryModelIdInventoryStockLocations**](InventoryStockLocationsApi.md#GETInventoryModelIdInventoryStockLocations) | **Get** /inventory_models/{inventoryModelId}/inventory_stock_locations | Retrieve the inventory stock locations associated to the inventory model
[**GETInventoryStockLocations**](InventoryStockLocationsApi.md#GETInventoryStockLocations) | **Get** /inventory_stock_locations | List all inventory stock locations
[**GETInventoryStockLocationsInventoryStockLocationId**](InventoryStockLocationsApi.md#GETInventoryStockLocationsInventoryStockLocationId) | **Get** /inventory_stock_locations/{inventoryStockLocationId} | Retrieve an inventory stock location
[**GETShipmentIdInventoryStockLocation**](InventoryStockLocationsApi.md#GETShipmentIdInventoryStockLocation) | **Get** /shipments/{shipmentId}/inventory_stock_location | Retrieve the inventory stock location associated to the shipment
[**GETStockLocationIdInventoryStockLocations**](InventoryStockLocationsApi.md#GETStockLocationIdInventoryStockLocations) | **Get** /stock_locations/{stockLocationId}/inventory_stock_locations | Retrieve the inventory stock locations associated to the stock location
[**PATCHInventoryStockLocationsInventoryStockLocationId**](InventoryStockLocationsApi.md#PATCHInventoryStockLocationsInventoryStockLocationId) | **Patch** /inventory_stock_locations/{inventoryStockLocationId} | Update an inventory stock location
[**POSTInventoryStockLocations**](InventoryStockLocationsApi.md#POSTInventoryStockLocations) | **Post** /inventory_stock_locations | Create an inventory stock location



## DELETEInventoryStockLocationsInventoryStockLocationId

> DELETEInventoryStockLocationsInventoryStockLocationId(ctx, inventoryStockLocationId).Execute()

Delete an inventory stock location



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
    inventoryStockLocationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InventoryStockLocationsApi.DELETEInventoryStockLocationsInventoryStockLocationId(context.Background(), inventoryStockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.DELETEInventoryStockLocationsInventoryStockLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryStockLocationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEInventoryStockLocationsInventoryStockLocationIdRequest struct via the builder pattern


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


## GETInventoryModelIdInventoryStockLocations

> GETInventoryModelIdInventoryStockLocations(ctx, inventoryModelId).Execute()

Retrieve the inventory stock locations associated to the inventory model



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
    inventoryModelId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InventoryStockLocationsApi.GETInventoryModelIdInventoryStockLocations(context.Background(), inventoryModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.GETInventoryModelIdInventoryStockLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryModelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryModelIdInventoryStockLocationsRequest struct via the builder pattern


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


## GETInventoryStockLocations

> GETInventoryStockLocations200Response GETInventoryStockLocations(ctx).Execute()

List all inventory stock locations



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
    resp, r, err := apiClient.InventoryStockLocationsApi.GETInventoryStockLocations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.GETInventoryStockLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInventoryStockLocations`: GETInventoryStockLocations200Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryStockLocationsApi.GETInventoryStockLocations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryStockLocationsRequest struct via the builder pattern


### Return type

[**GETInventoryStockLocations200Response**](GETInventoryStockLocations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETInventoryStockLocationsInventoryStockLocationId

> GETInventoryStockLocationsInventoryStockLocationId200Response GETInventoryStockLocationsInventoryStockLocationId(ctx, inventoryStockLocationId).Execute()

Retrieve an inventory stock location



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
    inventoryStockLocationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryStockLocationsApi.GETInventoryStockLocationsInventoryStockLocationId(context.Background(), inventoryStockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.GETInventoryStockLocationsInventoryStockLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInventoryStockLocationsInventoryStockLocationId`: GETInventoryStockLocationsInventoryStockLocationId200Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryStockLocationsApi.GETInventoryStockLocationsInventoryStockLocationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryStockLocationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryStockLocationsInventoryStockLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETInventoryStockLocationsInventoryStockLocationId200Response**](GETInventoryStockLocationsInventoryStockLocationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShipmentIdInventoryStockLocation

> GETShipmentIdInventoryStockLocation(ctx, shipmentId).Execute()

Retrieve the inventory stock location associated to the shipment



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
    r, err := apiClient.InventoryStockLocationsApi.GETShipmentIdInventoryStockLocation(context.Background(), shipmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.GETShipmentIdInventoryStockLocation``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShipmentIdInventoryStockLocationRequest struct via the builder pattern


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


## GETStockLocationIdInventoryStockLocations

> GETStockLocationIdInventoryStockLocations(ctx, stockLocationId).Execute()

Retrieve the inventory stock locations associated to the stock location



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
    r, err := apiClient.InventoryStockLocationsApi.GETStockLocationIdInventoryStockLocations(context.Background(), stockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.GETStockLocationIdInventoryStockLocations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockLocationIdInventoryStockLocationsRequest struct via the builder pattern


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


## PATCHInventoryStockLocationsInventoryStockLocationId

> PATCHInventoryStockLocationsInventoryStockLocationId200Response PATCHInventoryStockLocationsInventoryStockLocationId(ctx, inventoryStockLocationId).InventoryStockLocationUpdate(inventoryStockLocationUpdate).Execute()

Update an inventory stock location



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
    inventoryStockLocationUpdate := *openapiclient.NewInventoryStockLocationUpdate(*openapiclient.NewInventoryStockLocationUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHInventoryStockLocationsInventoryStockLocationId200ResponseDataAttributes())) // InventoryStockLocationUpdate | 
    inventoryStockLocationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryStockLocationsApi.PATCHInventoryStockLocationsInventoryStockLocationId(context.Background(), inventoryStockLocationId).InventoryStockLocationUpdate(inventoryStockLocationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.PATCHInventoryStockLocationsInventoryStockLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHInventoryStockLocationsInventoryStockLocationId`: PATCHInventoryStockLocationsInventoryStockLocationId200Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryStockLocationsApi.PATCHInventoryStockLocationsInventoryStockLocationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryStockLocationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHInventoryStockLocationsInventoryStockLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryStockLocationUpdate** | [**InventoryStockLocationUpdate**](InventoryStockLocationUpdate.md) |  | 


### Return type

[**PATCHInventoryStockLocationsInventoryStockLocationId200Response**](PATCHInventoryStockLocationsInventoryStockLocationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTInventoryStockLocations

> POSTInventoryStockLocations201Response POSTInventoryStockLocations(ctx).InventoryStockLocationCreate(inventoryStockLocationCreate).Execute()

Create an inventory stock location



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
    inventoryStockLocationCreate := *openapiclient.NewInventoryStockLocationCreate(*openapiclient.NewInventoryStockLocationCreateData(interface{}(123), *openapiclient.NewPOSTInventoryStockLocations201ResponseDataAttributes(interface{}(1)))) // InventoryStockLocationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryStockLocationsApi.POSTInventoryStockLocations(context.Background()).InventoryStockLocationCreate(inventoryStockLocationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockLocationsApi.POSTInventoryStockLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTInventoryStockLocations`: POSTInventoryStockLocations201Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryStockLocationsApi.POSTInventoryStockLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTInventoryStockLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryStockLocationCreate** | [**InventoryStockLocationCreate**](InventoryStockLocationCreate.md) |  | 

### Return type

[**POSTInventoryStockLocations201Response**](POSTInventoryStockLocations201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

