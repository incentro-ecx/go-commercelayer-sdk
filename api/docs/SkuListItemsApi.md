# \SkuListItemsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESkuListItemsSkuListItemId**](SkuListItemsApi.md#DELETESkuListItemsSkuListItemId) | **Delete** /sku_list_items/{skuListItemId} | Delete a SKU list item
[**GETSkuIdSkuListItems**](SkuListItemsApi.md#GETSkuIdSkuListItems) | **Get** /skus/{skuId}/sku_list_items | Retrieve the sku list items associated to the SKU
[**GETSkuListIdSkuListItems**](SkuListItemsApi.md#GETSkuListIdSkuListItems) | **Get** /sku_lists/{skuListId}/sku_list_items | Retrieve the sku list items associated to the SKU list
[**GETSkuListItems**](SkuListItemsApi.md#GETSkuListItems) | **Get** /sku_list_items | List all SKU list items
[**GETSkuListItemsSkuListItemId**](SkuListItemsApi.md#GETSkuListItemsSkuListItemId) | **Get** /sku_list_items/{skuListItemId} | Retrieve a SKU list item
[**PATCHSkuListItemsSkuListItemId**](SkuListItemsApi.md#PATCHSkuListItemsSkuListItemId) | **Patch** /sku_list_items/{skuListItemId} | Update a SKU list item
[**POSTSkuListItems**](SkuListItemsApi.md#POSTSkuListItems) | **Post** /sku_list_items | Create a SKU list item



## DELETESkuListItemsSkuListItemId

> DELETESkuListItemsSkuListItemId(ctx, skuListItemId).Execute()

Delete a SKU list item



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
    skuListItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SkuListItemsApi.DELETESkuListItemsSkuListItemId(context.Background(), skuListItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListItemsApi.DELETESkuListItemsSkuListItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETESkuListItemsSkuListItemIdRequest struct via the builder pattern


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


## GETSkuIdSkuListItems

> GETSkuIdSkuListItems(ctx, skuId).Execute()

Retrieve the sku list items associated to the SKU



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
    r, err := apiClient.SkuListItemsApi.GETSkuIdSkuListItems(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListItemsApi.GETSkuIdSkuListItems``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuIdSkuListItemsRequest struct via the builder pattern


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


## GETSkuListIdSkuListItems

> GETSkuListIdSkuListItems(ctx, skuListId).Execute()

Retrieve the sku list items associated to the SKU list



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
    skuListId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SkuListItemsApi.GETSkuListIdSkuListItems(context.Background(), skuListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListItemsApi.GETSkuListIdSkuListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListIdSkuListItemsRequest struct via the builder pattern


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


## GETSkuListItems

> GETSkuListItems200Response GETSkuListItems(ctx).Execute()

List all SKU list items



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
    resp, r, err := apiClient.SkuListItemsApi.GETSkuListItems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListItemsApi.GETSkuListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSkuListItems`: GETSkuListItems200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListItemsApi.GETSkuListItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListItemsRequest struct via the builder pattern


### Return type

[**GETSkuListItems200Response**](GETSkuListItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETSkuListItemsSkuListItemId

> GETSkuListItemsSkuListItemId200Response GETSkuListItemsSkuListItemId(ctx, skuListItemId).Execute()

Retrieve a SKU list item



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
    skuListItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListItemsApi.GETSkuListItemsSkuListItemId(context.Background(), skuListItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListItemsApi.GETSkuListItemsSkuListItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSkuListItemsSkuListItemId`: GETSkuListItemsSkuListItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListItemsApi.GETSkuListItemsSkuListItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuListItemsSkuListItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETSkuListItemsSkuListItemId200Response**](GETSkuListItemsSkuListItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHSkuListItemsSkuListItemId

> PATCHSkuListItemsSkuListItemId200Response PATCHSkuListItemsSkuListItemId(ctx, skuListItemId).SkuListItemUpdate(skuListItemUpdate).Execute()

Update a SKU list item



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
    skuListItemUpdate := *openapiclient.NewSkuListItemUpdate(*openapiclient.NewSkuListItemUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHSkuListItemsSkuListItemId200ResponseDataAttributes())) // SkuListItemUpdate | 
    skuListItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListItemsApi.PATCHSkuListItemsSkuListItemId(context.Background(), skuListItemId).SkuListItemUpdate(skuListItemUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListItemsApi.PATCHSkuListItemsSkuListItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHSkuListItemsSkuListItemId`: PATCHSkuListItemsSkuListItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListItemsApi.PATCHSkuListItemsSkuListItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuListItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHSkuListItemsSkuListItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuListItemUpdate** | [**SkuListItemUpdate**](SkuListItemUpdate.md) |  | 


### Return type

[**PATCHSkuListItemsSkuListItemId200Response**](PATCHSkuListItemsSkuListItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTSkuListItems

> POSTSkuListItems201Response POSTSkuListItems(ctx).SkuListItemCreate(skuListItemCreate).Execute()

Create a SKU list item



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
    skuListItemCreate := *openapiclient.NewSkuListItemCreate(*openapiclient.NewSkuListItemCreateData(interface{}(123), *openapiclient.NewPOSTSkuListItems201ResponseDataAttributes())) // SkuListItemCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuListItemsApi.POSTSkuListItems(context.Background()).SkuListItemCreate(skuListItemCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuListItemsApi.POSTSkuListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTSkuListItems`: POSTSkuListItems201Response
    fmt.Fprintf(os.Stdout, "Response from `SkuListItemsApi.POSTSkuListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTSkuListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuListItemCreate** | [**SkuListItemCreate**](SkuListItemCreate.md) |  | 

### Return type

[**POSTSkuListItems201Response**](POSTSkuListItems201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

