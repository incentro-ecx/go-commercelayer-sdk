# \SkuOptionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETESkuOptionsSkuOptionId**](SkuOptionsApi.md#DELETESkuOptionsSkuOptionId) | **Delete** /sku_options/{skuOptionId} | Delete a SKU option
[**GETLineItemOptionIdSkuOption**](SkuOptionsApi.md#GETLineItemOptionIdSkuOption) | **Get** /line_item_options/{lineItemOptionId}/sku_option | Retrieve the sku option associated to the line item option
[**GETSkuIdSkuOptions**](SkuOptionsApi.md#GETSkuIdSkuOptions) | **Get** /skus/{skuId}/sku_options | Retrieve the sku options associated to the SKU
[**GETSkuOptions**](SkuOptionsApi.md#GETSkuOptions) | **Get** /sku_options | List all SKU options
[**GETSkuOptionsSkuOptionId**](SkuOptionsApi.md#GETSkuOptionsSkuOptionId) | **Get** /sku_options/{skuOptionId} | Retrieve a SKU option
[**PATCHSkuOptionsSkuOptionId**](SkuOptionsApi.md#PATCHSkuOptionsSkuOptionId) | **Patch** /sku_options/{skuOptionId} | Update a SKU option
[**POSTSkuOptions**](SkuOptionsApi.md#POSTSkuOptions) | **Post** /sku_options | Create a SKU option



## DELETESkuOptionsSkuOptionId

> DELETESkuOptionsSkuOptionId(ctx, skuOptionId).Execute()

Delete a SKU option



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
    skuOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SkuOptionsApi.DELETESkuOptionsSkuOptionId(context.Background(), skuOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuOptionsApi.DELETESkuOptionsSkuOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETESkuOptionsSkuOptionIdRequest struct via the builder pattern


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


## GETLineItemOptionIdSkuOption

> GETLineItemOptionIdSkuOption(ctx, lineItemOptionId).Execute()

Retrieve the sku option associated to the line item option



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
    r, err := apiClient.SkuOptionsApi.GETLineItemOptionIdSkuOption(context.Background(), lineItemOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuOptionsApi.GETLineItemOptionIdSkuOption``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETLineItemOptionIdSkuOptionRequest struct via the builder pattern


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


## GETSkuIdSkuOptions

> GETSkuIdSkuOptions(ctx, skuId).Execute()

Retrieve the sku options associated to the SKU



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
    r, err := apiClient.SkuOptionsApi.GETSkuIdSkuOptions(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuOptionsApi.GETSkuIdSkuOptions``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuIdSkuOptionsRequest struct via the builder pattern


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


## GETSkuOptions

> GETSkuOptions200Response GETSkuOptions(ctx).Execute()

List all SKU options



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
    resp, r, err := apiClient.SkuOptionsApi.GETSkuOptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuOptionsApi.GETSkuOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSkuOptions`: GETSkuOptions200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuOptionsApi.GETSkuOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuOptionsRequest struct via the builder pattern


### Return type

[**GETSkuOptions200Response**](GETSkuOptions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETSkuOptionsSkuOptionId

> GETSkuOptionsSkuOptionId200Response GETSkuOptionsSkuOptionId(ctx, skuOptionId).Execute()

Retrieve a SKU option



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
    skuOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuOptionsApi.GETSkuOptionsSkuOptionId(context.Background(), skuOptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuOptionsApi.GETSkuOptionsSkuOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETSkuOptionsSkuOptionId`: GETSkuOptionsSkuOptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuOptionsApi.GETSkuOptionsSkuOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETSkuOptionsSkuOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETSkuOptionsSkuOptionId200Response**](GETSkuOptionsSkuOptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHSkuOptionsSkuOptionId

> PATCHSkuOptionsSkuOptionId200Response PATCHSkuOptionsSkuOptionId(ctx, skuOptionId).SkuOptionUpdate(skuOptionUpdate).Execute()

Update a SKU option



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
    skuOptionUpdate := *openapiclient.NewSkuOptionUpdate(*openapiclient.NewSkuOptionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHSkuOptionsSkuOptionId200ResponseDataAttributes())) // SkuOptionUpdate | 
    skuOptionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuOptionsApi.PATCHSkuOptionsSkuOptionId(context.Background(), skuOptionId).SkuOptionUpdate(skuOptionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuOptionsApi.PATCHSkuOptionsSkuOptionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHSkuOptionsSkuOptionId`: PATCHSkuOptionsSkuOptionId200Response
    fmt.Fprintf(os.Stdout, "Response from `SkuOptionsApi.PATCHSkuOptionsSkuOptionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuOptionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHSkuOptionsSkuOptionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuOptionUpdate** | [**SkuOptionUpdate**](SkuOptionUpdate.md) |  | 


### Return type

[**PATCHSkuOptionsSkuOptionId200Response**](PATCHSkuOptionsSkuOptionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTSkuOptions

> POSTSkuOptions201Response POSTSkuOptions(ctx).SkuOptionCreate(skuOptionCreate).Execute()

Create a SKU option



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
    skuOptionCreate := *openapiclient.NewSkuOptionCreate(*openapiclient.NewSkuOptionCreateData(interface{}(123), *openapiclient.NewPOSTSkuOptions201ResponseDataAttributes(interface{}(Embossing)))) // SkuOptionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SkuOptionsApi.POSTSkuOptions(context.Background()).SkuOptionCreate(skuOptionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SkuOptionsApi.POSTSkuOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTSkuOptions`: POSTSkuOptions201Response
    fmt.Fprintf(os.Stdout, "Response from `SkuOptionsApi.POSTSkuOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTSkuOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skuOptionCreate** | [**SkuOptionCreate**](SkuOptionCreate.md) |  | 

### Return type

[**POSTSkuOptions201Response**](POSTSkuOptions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

