# \FixedAmountPromotionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFixedAmountPromotionsFixedAmountPromotionId**](FixedAmountPromotionsApi.md#DELETEFixedAmountPromotionsFixedAmountPromotionId) | **Delete** /fixed_amount_promotions/{fixedAmountPromotionId} | Delete a fixed amount promotion
[**GETFixedAmountPromotions**](FixedAmountPromotionsApi.md#GETFixedAmountPromotions) | **Get** /fixed_amount_promotions | List all fixed amount promotions
[**GETFixedAmountPromotionsFixedAmountPromotionId**](FixedAmountPromotionsApi.md#GETFixedAmountPromotionsFixedAmountPromotionId) | **Get** /fixed_amount_promotions/{fixedAmountPromotionId} | Retrieve a fixed amount promotion
[**PATCHFixedAmountPromotionsFixedAmountPromotionId**](FixedAmountPromotionsApi.md#PATCHFixedAmountPromotionsFixedAmountPromotionId) | **Patch** /fixed_amount_promotions/{fixedAmountPromotionId} | Update a fixed amount promotion
[**POSTFixedAmountPromotions**](FixedAmountPromotionsApi.md#POSTFixedAmountPromotions) | **Post** /fixed_amount_promotions | Create a fixed amount promotion



## DELETEFixedAmountPromotionsFixedAmountPromotionId

> DELETEFixedAmountPromotionsFixedAmountPromotionId(ctx, fixedAmountPromotionId).Execute()

Delete a fixed amount promotion



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
    fixedAmountPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FixedAmountPromotionsApi.DELETEFixedAmountPromotionsFixedAmountPromotionId(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.DELETEFixedAmountPromotionsFixedAmountPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedAmountPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEFixedAmountPromotionsFixedAmountPromotionIdRequest struct via the builder pattern


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


## GETFixedAmountPromotions

> GETFixedAmountPromotions200Response GETFixedAmountPromotions(ctx).Execute()

List all fixed amount promotions



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
    resp, r, err := apiClient.FixedAmountPromotionsApi.GETFixedAmountPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.GETFixedAmountPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFixedAmountPromotions`: GETFixedAmountPromotions200Response
    fmt.Fprintf(os.Stdout, "Response from `FixedAmountPromotionsApi.GETFixedAmountPromotions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionsRequest struct via the builder pattern


### Return type

[**GETFixedAmountPromotions200Response**](GETFixedAmountPromotions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFixedAmountPromotionsFixedAmountPromotionId

> GETFixedAmountPromotionsFixedAmountPromotionId200Response GETFixedAmountPromotionsFixedAmountPromotionId(ctx, fixedAmountPromotionId).Execute()

Retrieve a fixed amount promotion



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
    fixedAmountPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedAmountPromotionsApi.GETFixedAmountPromotionsFixedAmountPromotionId(context.Background(), fixedAmountPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.GETFixedAmountPromotionsFixedAmountPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFixedAmountPromotionsFixedAmountPromotionId`: GETFixedAmountPromotionsFixedAmountPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `FixedAmountPromotionsApi.GETFixedAmountPromotionsFixedAmountPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedAmountPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFixedAmountPromotionsFixedAmountPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETFixedAmountPromotionsFixedAmountPromotionId200Response**](GETFixedAmountPromotionsFixedAmountPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHFixedAmountPromotionsFixedAmountPromotionId

> PATCHFixedAmountPromotionsFixedAmountPromotionId200Response PATCHFixedAmountPromotionsFixedAmountPromotionId(ctx, fixedAmountPromotionId).FixedAmountPromotionUpdate(fixedAmountPromotionUpdate).Execute()

Update a fixed amount promotion



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
    fixedAmountPromotionUpdate := *openapiclient.NewFixedAmountPromotionUpdate(*openapiclient.NewFixedAmountPromotionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes())) // FixedAmountPromotionUpdate | 
    fixedAmountPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedAmountPromotionsApi.PATCHFixedAmountPromotionsFixedAmountPromotionId(context.Background(), fixedAmountPromotionId).FixedAmountPromotionUpdate(fixedAmountPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.PATCHFixedAmountPromotionsFixedAmountPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHFixedAmountPromotionsFixedAmountPromotionId`: PATCHFixedAmountPromotionsFixedAmountPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `FixedAmountPromotionsApi.PATCHFixedAmountPromotionsFixedAmountPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fixedAmountPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHFixedAmountPromotionsFixedAmountPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fixedAmountPromotionUpdate** | [**FixedAmountPromotionUpdate**](FixedAmountPromotionUpdate.md) |  | 


### Return type

[**PATCHFixedAmountPromotionsFixedAmountPromotionId200Response**](PATCHFixedAmountPromotionsFixedAmountPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTFixedAmountPromotions

> POSTFixedAmountPromotions201Response POSTFixedAmountPromotions(ctx).FixedAmountPromotionCreate(fixedAmountPromotionCreate).Execute()

Create a fixed amount promotion



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
    fixedAmountPromotionCreate := *openapiclient.NewFixedAmountPromotionCreate(*openapiclient.NewFixedAmountPromotionCreateData(interface{}(123), *openapiclient.NewPOSTFixedAmountPromotions201ResponseDataAttributes(interface{}(Personal promotion), interface{}(2018-01-01T12:00:00.000Z), interface{}(2018-01-02T12:00:00.000Z), interface{}(1000)))) // FixedAmountPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FixedAmountPromotionsApi.POSTFixedAmountPromotions(context.Background()).FixedAmountPromotionCreate(fixedAmountPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FixedAmountPromotionsApi.POSTFixedAmountPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTFixedAmountPromotions`: POSTFixedAmountPromotions201Response
    fmt.Fprintf(os.Stdout, "Response from `FixedAmountPromotionsApi.POSTFixedAmountPromotions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTFixedAmountPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fixedAmountPromotionCreate** | [**FixedAmountPromotionCreate**](FixedAmountPromotionCreate.md) |  | 

### Return type

[**POSTFixedAmountPromotions201Response**](POSTFixedAmountPromotions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

