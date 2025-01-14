# \InventoryModelsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEInventoryModelsInventoryModelId**](InventoryModelsApi.md#DELETEInventoryModelsInventoryModelId) | **Delete** /inventory_models/{inventoryModelId} | Delete an inventory model
[**GETInventoryModels**](InventoryModelsApi.md#GETInventoryModels) | **Get** /inventory_models | List all inventory models
[**GETInventoryModelsInventoryModelId**](InventoryModelsApi.md#GETInventoryModelsInventoryModelId) | **Get** /inventory_models/{inventoryModelId} | Retrieve an inventory model
[**GETInventoryReturnLocationIdInventoryModel**](InventoryModelsApi.md#GETInventoryReturnLocationIdInventoryModel) | **Get** /inventory_return_locations/{inventoryReturnLocationId}/inventory_model | Retrieve the inventory model associated to the inventory return location
[**GETInventoryStockLocationIdInventoryModel**](InventoryModelsApi.md#GETInventoryStockLocationIdInventoryModel) | **Get** /inventory_stock_locations/{inventoryStockLocationId}/inventory_model | Retrieve the inventory model associated to the inventory stock location
[**GETMarketIdInventoryModel**](InventoryModelsApi.md#GETMarketIdInventoryModel) | **Get** /markets/{marketId}/inventory_model | Retrieve the inventory model associated to the market
[**PATCHInventoryModelsInventoryModelId**](InventoryModelsApi.md#PATCHInventoryModelsInventoryModelId) | **Patch** /inventory_models/{inventoryModelId} | Update an inventory model
[**POSTInventoryModels**](InventoryModelsApi.md#POSTInventoryModels) | **Post** /inventory_models | Create an inventory model



## DELETEInventoryModelsInventoryModelId

> DELETEInventoryModelsInventoryModelId(ctx, inventoryModelId).Execute()

Delete an inventory model



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
    r, err := apiClient.InventoryModelsApi.DELETEInventoryModelsInventoryModelId(context.Background(), inventoryModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryModelsApi.DELETEInventoryModelsInventoryModelId``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDELETEInventoryModelsInventoryModelIdRequest struct via the builder pattern


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


## GETInventoryModels

> GETInventoryModels200Response GETInventoryModels(ctx).Execute()

List all inventory models



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
    resp, r, err := apiClient.InventoryModelsApi.GETInventoryModels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryModelsApi.GETInventoryModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInventoryModels`: GETInventoryModels200Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryModelsApi.GETInventoryModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryModelsRequest struct via the builder pattern


### Return type

[**GETInventoryModels200Response**](GETInventoryModels200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETInventoryModelsInventoryModelId

> GETInventoryModelsInventoryModelId200Response GETInventoryModelsInventoryModelId(ctx, inventoryModelId).Execute()

Retrieve an inventory model



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
    resp, r, err := apiClient.InventoryModelsApi.GETInventoryModelsInventoryModelId(context.Background(), inventoryModelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryModelsApi.GETInventoryModelsInventoryModelId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETInventoryModelsInventoryModelId`: GETInventoryModelsInventoryModelId200Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryModelsApi.GETInventoryModelsInventoryModelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryModelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETInventoryModelsInventoryModelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETInventoryModelsInventoryModelId200Response**](GETInventoryModelsInventoryModelId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETInventoryReturnLocationIdInventoryModel

> GETInventoryReturnLocationIdInventoryModel(ctx, inventoryReturnLocationId).Execute()

Retrieve the inventory model associated to the inventory return location



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
    r, err := apiClient.InventoryModelsApi.GETInventoryReturnLocationIdInventoryModel(context.Background(), inventoryReturnLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryModelsApi.GETInventoryReturnLocationIdInventoryModel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETInventoryReturnLocationIdInventoryModelRequest struct via the builder pattern


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


## GETInventoryStockLocationIdInventoryModel

> GETInventoryStockLocationIdInventoryModel(ctx, inventoryStockLocationId).Execute()

Retrieve the inventory model associated to the inventory stock location



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
    r, err := apiClient.InventoryModelsApi.GETInventoryStockLocationIdInventoryModel(context.Background(), inventoryStockLocationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryModelsApi.GETInventoryStockLocationIdInventoryModel``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETInventoryStockLocationIdInventoryModelRequest struct via the builder pattern


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


## GETMarketIdInventoryModel

> GETMarketIdInventoryModel(ctx, marketId).Execute()

Retrieve the inventory model associated to the market



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
    marketId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InventoryModelsApi.GETMarketIdInventoryModel(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryModelsApi.GETMarketIdInventoryModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketIdInventoryModelRequest struct via the builder pattern


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


## PATCHInventoryModelsInventoryModelId

> PATCHInventoryModelsInventoryModelId200Response PATCHInventoryModelsInventoryModelId(ctx, inventoryModelId).InventoryModelUpdate(inventoryModelUpdate).Execute()

Update an inventory model



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
    inventoryModelUpdate := *openapiclient.NewInventoryModelUpdate(*openapiclient.NewInventoryModelUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHInventoryModelsInventoryModelId200ResponseDataAttributes())) // InventoryModelUpdate | 
    inventoryModelId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryModelsApi.PATCHInventoryModelsInventoryModelId(context.Background(), inventoryModelId).InventoryModelUpdate(inventoryModelUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryModelsApi.PATCHInventoryModelsInventoryModelId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHInventoryModelsInventoryModelId`: PATCHInventoryModelsInventoryModelId200Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryModelsApi.PATCHInventoryModelsInventoryModelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryModelId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHInventoryModelsInventoryModelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryModelUpdate** | [**InventoryModelUpdate**](InventoryModelUpdate.md) |  | 


### Return type

[**PATCHInventoryModelsInventoryModelId200Response**](PATCHInventoryModelsInventoryModelId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTInventoryModels

> POSTInventoryModels201Response POSTInventoryModels(ctx).InventoryModelCreate(inventoryModelCreate).Execute()

Create an inventory model



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
    inventoryModelCreate := *openapiclient.NewInventoryModelCreate(*openapiclient.NewInventoryModelCreateData(interface{}(123), *openapiclient.NewPOSTInventoryModels201ResponseDataAttributes(interface{}(EU Inventory Model)))) // InventoryModelCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InventoryModelsApi.POSTInventoryModels(context.Background()).InventoryModelCreate(inventoryModelCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryModelsApi.POSTInventoryModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTInventoryModels`: POSTInventoryModels201Response
    fmt.Fprintf(os.Stdout, "Response from `InventoryModelsApi.POSTInventoryModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTInventoryModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inventoryModelCreate** | [**InventoryModelCreate**](InventoryModelCreate.md) |  | 

### Return type

[**POSTInventoryModels201Response**](POSTInventoryModels201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

