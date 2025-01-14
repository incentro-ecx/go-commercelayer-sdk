# \BundlesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBundlesBundleId**](BundlesApi.md#DELETEBundlesBundleId) | **Delete** /bundles/{bundleId} | Delete a bundle
[**GETBundles**](BundlesApi.md#GETBundles) | **Get** /bundles | List all bundles
[**GETBundlesBundleId**](BundlesApi.md#GETBundlesBundleId) | **Get** /bundles/{bundleId} | Retrieve a bundle
[**GETOrderIdAvailableFreeBundles**](BundlesApi.md#GETOrderIdAvailableFreeBundles) | **Get** /orders/{orderId}/available_free_bundles | Retrieve the available free bundles associated to the order
[**GETSkuListIdBundles**](BundlesApi.md#GETSkuListIdBundles) | **Get** /sku_lists/{skuListId}/bundles | Retrieve the bundles associated to the SKU list
[**PATCHBundlesBundleId**](BundlesApi.md#PATCHBundlesBundleId) | **Patch** /bundles/{bundleId} | Update a bundle
[**POSTBundles**](BundlesApi.md#POSTBundles) | **Post** /bundles | Create a bundle



## DELETEBundlesBundleId

> DELETEBundlesBundleId(ctx, bundleId).Execute()

Delete a bundle



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
    bundleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BundlesApi.DELETEBundlesBundleId(context.Background(), bundleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.DELETEBundlesBundleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBundlesBundleIdRequest struct via the builder pattern


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


## GETBundles

> GETBundles200Response GETBundles(ctx).Execute()

List all bundles



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
    resp, r, err := apiClient.BundlesApi.GETBundles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.GETBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBundles`: GETBundles200Response
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.GETBundles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBundlesRequest struct via the builder pattern


### Return type

[**GETBundles200Response**](GETBundles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETBundlesBundleId

> GETBundlesBundleId200Response GETBundlesBundleId(ctx, bundleId).Execute()

Retrieve a bundle



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
    bundleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesApi.GETBundlesBundleId(context.Background(), bundleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.GETBundlesBundleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBundlesBundleId`: GETBundlesBundleId200Response
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.GETBundlesBundleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBundlesBundleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETBundlesBundleId200Response**](GETBundlesBundleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderIdAvailableFreeBundles

> GETOrderIdAvailableFreeBundles(ctx, orderId).Execute()

Retrieve the available free bundles associated to the order



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
    r, err := apiClient.BundlesApi.GETOrderIdAvailableFreeBundles(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.GETOrderIdAvailableFreeBundles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdAvailableFreeBundlesRequest struct via the builder pattern


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


## GETSkuListIdBundles

> GETSkuListIdBundles(ctx, skuListId).Execute()

Retrieve the bundles associated to the SKU list



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
    r, err := apiClient.BundlesApi.GETSkuListIdBundles(context.Background(), skuListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.GETSkuListIdBundles``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuListIdBundlesRequest struct via the builder pattern


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


## PATCHBundlesBundleId

> PATCHBundlesBundleId200Response PATCHBundlesBundleId(ctx, bundleId).BundleUpdate(bundleUpdate).Execute()

Update a bundle



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
    bundleUpdate := *openapiclient.NewBundleUpdate(*openapiclient.NewBundleUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHBundlesBundleId200ResponseDataAttributes())) // BundleUpdate | 
    bundleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesApi.PATCHBundlesBundleId(context.Background(), bundleId).BundleUpdate(bundleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.PATCHBundlesBundleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHBundlesBundleId`: PATCHBundlesBundleId200Response
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.PATCHBundlesBundleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bundleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBundlesBundleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleUpdate** | [**BundleUpdate**](BundleUpdate.md) |  | 


### Return type

[**PATCHBundlesBundleId200Response**](PATCHBundlesBundleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTBundles

> POSTBundles201Response POSTBundles(ctx).BundleCreate(bundleCreate).Execute()

Create a bundle



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
    bundleCreate := *openapiclient.NewBundleCreate(*openapiclient.NewBundleCreateData(interface{}(123), *openapiclient.NewPOSTBundles201ResponseDataAttributes(interface{}(BUNDMM000000FFFFFFXLXX), interface{}(Men's Black T-shirt (XL) with Black Cap and Socks, all with White Logo)))) // BundleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundlesApi.POSTBundles(context.Background()).BundleCreate(bundleCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundlesApi.POSTBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTBundles`: POSTBundles201Response
    fmt.Fprintf(os.Stdout, "Response from `BundlesApi.POSTBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleCreate** | [**BundleCreate**](BundleCreate.md) |  | 

### Return type

[**POSTBundles201Response**](POSTBundles201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

