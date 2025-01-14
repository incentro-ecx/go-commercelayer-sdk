# \InventoryReturnLocationsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEInventoryReturnLocationsInventoryReturnLocationId**](InventoryReturnLocationsApi.md#DELETEInventoryReturnLocationsInventoryReturnLocationId) | **Delete** /inventory_return_locations/{inventoryReturnLocationId} | Delete an inventory return location
[**GETInventoryModelIdInventoryReturnLocations**](InventoryReturnLocationsApi.md#GETInventoryModelIdInventoryReturnLocations) | **Get** /inventory_models/{inventoryModelId}/inventory_return_locations | Retrieve the inventory return locations associated to the inventory model
[**GETInventoryReturnLocations**](InventoryReturnLocationsApi.md#GETInventoryReturnLocations) | **Get** /inventory_return_locations | List all inventory return locations
[**GETInventoryReturnLocationsInventoryReturnLocationId**](InventoryReturnLocationsApi.md#GETInventoryReturnLocationsInventoryReturnLocationId) | **Get** /inventory_return_locations/{inventoryReturnLocationId} | Retrieve an inventory return location
[**GETStockLocationIdInventoryReturnLocations**](InventoryReturnLocationsApi.md#GETStockLocationIdInventoryReturnLocations) | **Get** /stock_locations/{stockLocationId}/inventory_return_locations | Retrieve the inventory return locations associated to the stock location
[**PATCHInventoryReturnLocationsInventoryReturnLocationId**](InventoryReturnLocationsApi.md#PATCHInventoryReturnLocationsInventoryReturnLocationId) | **Patch** /inventory_return_locations/{inventoryReturnLocationId} | Update an inventory return location
[**POSTInventoryReturnLocations**](InventoryReturnLocationsApi.md#POSTInventoryReturnLocations) | **Post** /inventory_return_locations | Create an inventory return location



## DELETEInventoryReturnLocationsInventoryReturnLocationId

> DELETEInventoryReturnLocationsInventoryReturnLocationId(ctx, inventoryReturnLocationId).Execute()

Delete an inventory return location



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
    inventoryReturnLocationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InventoryReturnLocationsApi.DELETEInventoryReturnLocationsInventoryReturnLocationId(context.Background(), inventoryReturnLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.DELETEInventoryReturnLocationsInventoryReturnLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryReturnLocationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEInventoryReturnLocationsInventoryReturnLocationIdRequest struct via the builder pattern


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


## GETInventoryModelIdInventoryReturnLocations

> GETInventoryModelIdInventoryReturnLocations(ctx, inventoryModelId).Execute()

Retrieve the inventory return locations associated to the inventory model



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
    r, err := apiClient.InventoryReturnLocationsApi.GETInventoryModelIdInventoryReturnLocations(context.Background(), inventoryModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.GETInventoryModelIdInventoryReturnLocations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETInventoryModelIdInventoryReturnLocationsRequest struct via the builder pattern


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


## GETInventoryReturnLocations

> GETInventoryReturnLocations200Response GETInventoryReturnLocations(ctx).Execute()

List all inventory return locations



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
    resp, r, err := apiClient.InventoryReturnLocationsApi.GETInventoryReturnLocations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.GETInventoryReturnLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInventoryReturnLocations`: GETInventoryReturnLocations200Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryReturnLocationsApi.GETInventoryReturnLocations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryReturnLocationsRequest struct via the builder pattern


### Return type

[**GETInventoryReturnLocations200Response**](GETInventoryReturnLocations200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETInventoryReturnLocationsInventoryReturnLocationId

> GETInventoryReturnLocationsInventoryReturnLocationId200Response GETInventoryReturnLocationsInventoryReturnLocationId(ctx, inventoryReturnLocationId).Execute()

Retrieve an inventory return location



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
    inventoryReturnLocationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryReturnLocationsApi.GETInventoryReturnLocationsInventoryReturnLocationId(context.Background(), inventoryReturnLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.GETInventoryReturnLocationsInventoryReturnLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInventoryReturnLocationsInventoryReturnLocationId`: GETInventoryReturnLocationsInventoryReturnLocationId200Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryReturnLocationsApi.GETInventoryReturnLocationsInventoryReturnLocationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryReturnLocationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryReturnLocationsInventoryReturnLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETInventoryReturnLocationsInventoryReturnLocationId200Response**](GETInventoryReturnLocationsInventoryReturnLocationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStockLocationIdInventoryReturnLocations

> GETStockLocationIdInventoryReturnLocations(ctx, stockLocationId).Execute()

Retrieve the inventory return locations associated to the stock location



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
    r, err := apiClient.InventoryReturnLocationsApi.GETStockLocationIdInventoryReturnLocations(context.Background(), stockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.GETStockLocationIdInventoryReturnLocations``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETStockLocationIdInventoryReturnLocationsRequest struct via the builder pattern


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


## PATCHInventoryReturnLocationsInventoryReturnLocationId

> PATCHInventoryReturnLocationsInventoryReturnLocationId200Response PATCHInventoryReturnLocationsInventoryReturnLocationId(ctx, inventoryReturnLocationId).InventoryReturnLocationUpdate(inventoryReturnLocationUpdate).Execute()

Update an inventory return location



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
    inventoryReturnLocationUpdate := *openapiclient.NewInventoryReturnLocationUpdate(*openapiclient.NewInventoryReturnLocationUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHInventoryReturnLocationsInventoryReturnLocationId200ResponseDataAttributes())) // InventoryReturnLocationUpdate | 
    inventoryReturnLocationId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryReturnLocationsApi.PATCHInventoryReturnLocationsInventoryReturnLocationId(context.Background(), inventoryReturnLocationId).InventoryReturnLocationUpdate(inventoryReturnLocationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.PATCHInventoryReturnLocationsInventoryReturnLocationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHInventoryReturnLocationsInventoryReturnLocationId`: PATCHInventoryReturnLocationsInventoryReturnLocationId200Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryReturnLocationsApi.PATCHInventoryReturnLocationsInventoryReturnLocationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryReturnLocationId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHInventoryReturnLocationsInventoryReturnLocationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryReturnLocationUpdate** | [**InventoryReturnLocationUpdate**](InventoryReturnLocationUpdate.md) |  | 


### Return type

[**PATCHInventoryReturnLocationsInventoryReturnLocationId200Response**](PATCHInventoryReturnLocationsInventoryReturnLocationId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTInventoryReturnLocations

> POSTInventoryReturnLocations201Response POSTInventoryReturnLocations(ctx).InventoryReturnLocationCreate(inventoryReturnLocationCreate).Execute()

Create an inventory return location



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
    inventoryReturnLocationCreate := *openapiclient.NewInventoryReturnLocationCreate(*openapiclient.NewInventoryReturnLocationCreateData(interface{}(123), *openapiclient.NewPOSTInventoryReturnLocations201ResponseDataAttributes(interface{}(1)))) // InventoryReturnLocationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryReturnLocationsApi.POSTInventoryReturnLocations(context.Background()).InventoryReturnLocationCreate(inventoryReturnLocationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryReturnLocationsApi.POSTInventoryReturnLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTInventoryReturnLocations`: POSTInventoryReturnLocations201Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryReturnLocationsApi.POSTInventoryReturnLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTInventoryReturnLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryReturnLocationCreate** | [**InventoryReturnLocationCreate**](InventoryReturnLocationCreate.md) |  | 

### Return type

[**POSTInventoryReturnLocations201Response**](POSTInventoryReturnLocations201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

