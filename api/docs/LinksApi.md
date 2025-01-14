# \LinksApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETELinksLinkId**](LinksApi.md#DELETELinksLinkId) | **Delete** /links/{linkId} | Delete a link
[**GETLinks**](LinksApi.md#GETLinks) | **Get** /links | List all links
[**GETLinksLinkId**](LinksApi.md#GETLinksLinkId) | **Get** /links/{linkId} | Retrieve a link
[**GETOrderIdLinks**](LinksApi.md#GETOrderIdLinks) | **Get** /orders/{orderId}/links | Retrieve the links associated to the order
[**GETSkuIdLinks**](LinksApi.md#GETSkuIdLinks) | **Get** /skus/{skuId}/links | Retrieve the links associated to the SKU
[**GETSkuListIdLinks**](LinksApi.md#GETSkuListIdLinks) | **Get** /sku_lists/{skuListId}/links | Retrieve the links associated to the SKU list
[**PATCHLinksLinkId**](LinksApi.md#PATCHLinksLinkId) | **Patch** /links/{linkId} | Update a link
[**POSTLinks**](LinksApi.md#POSTLinks) | **Post** /links | Create a link



## DELETELinksLinkId

> DELETELinksLinkId(ctx, linkId).Execute()

Delete a link



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
    linkId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LinksApi.DELETELinksLinkId(context.Background(), linkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.DELETELinksLinkId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETELinksLinkIdRequest struct via the builder pattern


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


## GETLinks

> GETLinks200Response GETLinks(ctx).Execute()

List all links



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
    resp, r, err := apiClient.LinksApi.GETLinks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.GETLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETLinks`: GETLinks200Response
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.GETLinks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETLinksRequest struct via the builder pattern


### Return type

[**GETLinks200Response**](GETLinks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETLinksLinkId

> GETLinksLinkId200Response GETLinksLinkId(ctx, linkId).Execute()

Retrieve a link



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
    linkId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinksApi.GETLinksLinkId(context.Background(), linkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.GETLinksLinkId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETLinksLinkId`: GETLinksLinkId200Response
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.GETLinksLinkId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLinksLinkIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETLinksLinkId200Response**](GETLinksLinkId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETOrderIdLinks

> GETOrderIdLinks(ctx, orderId).Execute()

Retrieve the links associated to the order



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
    r, err := apiClient.LinksApi.GETOrderIdLinks(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.GETOrderIdLinks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETOrderIdLinksRequest struct via the builder pattern


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


## GETSkuIdLinks

> GETSkuIdLinks(ctx, skuId).Execute()

Retrieve the links associated to the SKU



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
    r, err := apiClient.LinksApi.GETSkuIdLinks(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.GETSkuIdLinks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuIdLinksRequest struct via the builder pattern


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


## GETSkuListIdLinks

> GETSkuListIdLinks(ctx, skuListId).Execute()

Retrieve the links associated to the SKU list



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
    r, err := apiClient.LinksApi.GETSkuListIdLinks(context.Background(), skuListId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.GETSkuListIdLinks``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETSkuListIdLinksRequest struct via the builder pattern


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


## PATCHLinksLinkId

> PATCHLinksLinkId200Response PATCHLinksLinkId(ctx, linkId).LinkUpdate(linkUpdate).Execute()

Update a link



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
    linkUpdate := *openapiclient.NewLinkUpdate(*openapiclient.NewLinkUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHLinksLinkId200ResponseDataAttributes())) // LinkUpdate | 
    linkId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinksApi.PATCHLinksLinkId(context.Background(), linkId).LinkUpdate(linkUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.PATCHLinksLinkId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHLinksLinkId`: PATCHLinksLinkId200Response
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.PATCHLinksLinkId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHLinksLinkIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkUpdate** | [**LinkUpdate**](LinkUpdate.md) |  | 


### Return type

[**PATCHLinksLinkId200Response**](PATCHLinksLinkId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTLinks

> POSTLinks201Response POSTLinks(ctx).LinkCreate(linkCreate).Execute()

Create a link



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
    linkCreate := *openapiclient.NewLinkCreate(*openapiclient.NewLinkCreateData(interface{}(123), *openapiclient.NewPOSTLinks201ResponseDataAttributes(interface{}(FW SALE 2023), interface{}(xxxx-yyyy-zzzz), interface{}(market:id:GhvCxsElAQ,market:id:kJhgVcxZDr)))) // LinkCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinksApi.POSTLinks(context.Background()).LinkCreate(linkCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinksApi.POSTLinks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTLinks`: POSTLinks201Response
    fmt.Fprintf(os.Stdout, "Response from `LinksApi.POSTLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkCreate** | [**LinkCreate**](LinkCreate.md) |  | 

### Return type

[**POSTLinks201Response**](POSTLinks201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

