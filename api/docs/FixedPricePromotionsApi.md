# \FixedPricePromotionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFixedPricePromotionsFixedPricePromotionId**](FixedPricePromotionsApi.md#DELETEFixedPricePromotionsFixedPricePromotionId) | **Delete** /fixed_price_promotions/{fixedPricePromotionId} | Delete a fixed price promotion
[**GETFixedPricePromotions**](FixedPricePromotionsApi.md#GETFixedPricePromotions) | **Get** /fixed_price_promotions | List all fixed price promotions
[**GETFixedPricePromotionsFixedPricePromotionId**](FixedPricePromotionsApi.md#GETFixedPricePromotionsFixedPricePromotionId) | **Get** /fixed_price_promotions/{fixedPricePromotionId} | Retrieve a fixed price promotion
[**PATCHFixedPricePromotionsFixedPricePromotionId**](FixedPricePromotionsApi.md#PATCHFixedPricePromotionsFixedPricePromotionId) | **Patch** /fixed_price_promotions/{fixedPricePromotionId} | Update a fixed price promotion
[**POSTFixedPricePromotions**](FixedPricePromotionsApi.md#POSTFixedPricePromotions) | **Post** /fixed_price_promotions | Create a fixed price promotion



## DELETEFixedPricePromotionsFixedPricePromotionId

> DELETEFixedPricePromotionsFixedPricePromotionId(ctx, fixedPricePromotionId).Execute()

Delete a fixed price promotion



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
    fixedPricePromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FixedPricePromotionsApi.DELETEFixedPricePromotionsFixedPricePromotionId(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.DELETEFixedPricePromotionsFixedPricePromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedPricePromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEFixedPricePromotionsFixedPricePromotionIdRequest struct via the builder pattern


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


## GETFixedPricePromotions

> GETFixedPricePromotions200Response GETFixedPricePromotions(ctx).Execute()

List all fixed price promotions



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
    resp, r, err := apiClient.FixedPricePromotionsApi.GETFixedPricePromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.GETFixedPricePromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFixedPricePromotions`: GETFixedPricePromotions200Response
    fmt.Fprintf(os.Stdout, "Response from `FixedPricePromotionsApi.GETFixedPricePromotions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedPricePromotionsRequest struct via the builder pattern


### Return type

[**GETFixedPricePromotions200Response**](GETFixedPricePromotions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFixedPricePromotionsFixedPricePromotionId

> GETFixedPricePromotionsFixedPricePromotionId200Response GETFixedPricePromotionsFixedPricePromotionId(ctx, fixedPricePromotionId).Execute()

Retrieve a fixed price promotion



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
    fixedPricePromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedPricePromotionsApi.GETFixedPricePromotionsFixedPricePromotionId(context.Background(), fixedPricePromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.GETFixedPricePromotionsFixedPricePromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFixedPricePromotionsFixedPricePromotionId`: GETFixedPricePromotionsFixedPricePromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `FixedPricePromotionsApi.GETFixedPricePromotionsFixedPricePromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedPricePromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedPricePromotionsFixedPricePromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETFixedPricePromotionsFixedPricePromotionId200Response**](GETFixedPricePromotionsFixedPricePromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHFixedPricePromotionsFixedPricePromotionId

> PATCHFixedPricePromotionsFixedPricePromotionId200Response PATCHFixedPricePromotionsFixedPricePromotionId(ctx, fixedPricePromotionId).FixedPricePromotionUpdate(fixedPricePromotionUpdate).Execute()

Update a fixed price promotion



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
    fixedPricePromotionUpdate := *openapiclient.NewFixedPricePromotionUpdate(*openapiclient.NewFixedPricePromotionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHFixedPricePromotionsFixedPricePromotionId200ResponseDataAttributes())) // FixedPricePromotionUpdate | 
    fixedPricePromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedPricePromotionsApi.PATCHFixedPricePromotionsFixedPricePromotionId(context.Background(), fixedPricePromotionId).FixedPricePromotionUpdate(fixedPricePromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.PATCHFixedPricePromotionsFixedPricePromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHFixedPricePromotionsFixedPricePromotionId`: PATCHFixedPricePromotionsFixedPricePromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `FixedPricePromotionsApi.PATCHFixedPricePromotionsFixedPricePromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedPricePromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHFixedPricePromotionsFixedPricePromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fixedPricePromotionUpdate** | [**FixedPricePromotionUpdate**](FixedPricePromotionUpdate.md) |  | 


### Return type

[**PATCHFixedPricePromotionsFixedPricePromotionId200Response**](PATCHFixedPricePromotionsFixedPricePromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTFixedPricePromotions

> POSTFixedPricePromotions201Response POSTFixedPricePromotions(ctx).FixedPricePromotionCreate(fixedPricePromotionCreate).Execute()

Create a fixed price promotion



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
    fixedPricePromotionCreate := *openapiclient.NewFixedPricePromotionCreate(*openapiclient.NewFixedPricePromotionCreateData(interface{}(123), *openapiclient.NewPOSTFixedPricePromotions201ResponseDataAttributes(interface{}(Personal promotion), interface{}(2018-01-01T12:00:00.000Z), interface{}(2018-01-02T12:00:00.000Z), interface{}(1000)))) // FixedPricePromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedPricePromotionsApi.POSTFixedPricePromotions(context.Background()).FixedPricePromotionCreate(fixedPricePromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedPricePromotionsApi.POSTFixedPricePromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTFixedPricePromotions`: POSTFixedPricePromotions201Response
    fmt.Fprintf(os.Stdout, "Response from `FixedPricePromotionsApi.POSTFixedPricePromotions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTFixedPricePromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fixedPricePromotionCreate** | [**FixedPricePromotionCreate**](FixedPricePromotionCreate.md) |  | 

### Return type

[**POSTFixedPricePromotions201Response**](POSTFixedPricePromotions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

