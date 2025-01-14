# \FreeShippingPromotionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFreeShippingPromotionsFreeShippingPromotionId**](FreeShippingPromotionsApi.md#DELETEFreeShippingPromotionsFreeShippingPromotionId) | **Delete** /free_shipping_promotions/{freeShippingPromotionId} | Delete a free shipping promotion
[**GETFreeShippingPromotions**](FreeShippingPromotionsApi.md#GETFreeShippingPromotions) | **Get** /free_shipping_promotions | List all free shipping promotions
[**GETFreeShippingPromotionsFreeShippingPromotionId**](FreeShippingPromotionsApi.md#GETFreeShippingPromotionsFreeShippingPromotionId) | **Get** /free_shipping_promotions/{freeShippingPromotionId} | Retrieve a free shipping promotion
[**PATCHFreeShippingPromotionsFreeShippingPromotionId**](FreeShippingPromotionsApi.md#PATCHFreeShippingPromotionsFreeShippingPromotionId) | **Patch** /free_shipping_promotions/{freeShippingPromotionId} | Update a free shipping promotion
[**POSTFreeShippingPromotions**](FreeShippingPromotionsApi.md#POSTFreeShippingPromotions) | **Post** /free_shipping_promotions | Create a free shipping promotion



## DELETEFreeShippingPromotionsFreeShippingPromotionId

> DELETEFreeShippingPromotionsFreeShippingPromotionId(ctx, freeShippingPromotionId).Execute()

Delete a free shipping promotion



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
    freeShippingPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FreeShippingPromotionsApi.DELETEFreeShippingPromotionsFreeShippingPromotionId(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.DELETEFreeShippingPromotionsFreeShippingPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeShippingPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEFreeShippingPromotionsFreeShippingPromotionIdRequest struct via the builder pattern


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


## GETFreeShippingPromotions

> GETFreeShippingPromotions200Response GETFreeShippingPromotions(ctx).Execute()

List all free shipping promotions



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
    resp, r, err := apiClient.FreeShippingPromotionsApi.GETFreeShippingPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.GETFreeShippingPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFreeShippingPromotions`: GETFreeShippingPromotions200Response
    fmt.Fprintf(os.Stdout, "Response from `FreeShippingPromotionsApi.GETFreeShippingPromotions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionsRequest struct via the builder pattern


### Return type

[**GETFreeShippingPromotions200Response**](GETFreeShippingPromotions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFreeShippingPromotionsFreeShippingPromotionId

> GETFreeShippingPromotionsFreeShippingPromotionId200Response GETFreeShippingPromotionsFreeShippingPromotionId(ctx, freeShippingPromotionId).Execute()

Retrieve a free shipping promotion



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
    freeShippingPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FreeShippingPromotionsApi.GETFreeShippingPromotionsFreeShippingPromotionId(context.Background(), freeShippingPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.GETFreeShippingPromotionsFreeShippingPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFreeShippingPromotionsFreeShippingPromotionId`: GETFreeShippingPromotionsFreeShippingPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `FreeShippingPromotionsApi.GETFreeShippingPromotionsFreeShippingPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeShippingPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFreeShippingPromotionsFreeShippingPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETFreeShippingPromotionsFreeShippingPromotionId200Response**](GETFreeShippingPromotionsFreeShippingPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHFreeShippingPromotionsFreeShippingPromotionId

> PATCHFreeShippingPromotionsFreeShippingPromotionId200Response PATCHFreeShippingPromotionsFreeShippingPromotionId(ctx, freeShippingPromotionId).FreeShippingPromotionUpdate(freeShippingPromotionUpdate).Execute()

Update a free shipping promotion



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
    freeShippingPromotionUpdate := *openapiclient.NewFreeShippingPromotionUpdate(*openapiclient.NewFreeShippingPromotionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHFreeShippingPromotionsFreeShippingPromotionId200ResponseDataAttributes())) // FreeShippingPromotionUpdate | 
    freeShippingPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FreeShippingPromotionsApi.PATCHFreeShippingPromotionsFreeShippingPromotionId(context.Background(), freeShippingPromotionId).FreeShippingPromotionUpdate(freeShippingPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.PATCHFreeShippingPromotionsFreeShippingPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHFreeShippingPromotionsFreeShippingPromotionId`: PATCHFreeShippingPromotionsFreeShippingPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `FreeShippingPromotionsApi.PATCHFreeShippingPromotionsFreeShippingPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**freeShippingPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHFreeShippingPromotionsFreeShippingPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **freeShippingPromotionUpdate** | [**FreeShippingPromotionUpdate**](FreeShippingPromotionUpdate.md) |  | 


### Return type

[**PATCHFreeShippingPromotionsFreeShippingPromotionId200Response**](PATCHFreeShippingPromotionsFreeShippingPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTFreeShippingPromotions

> POSTFreeShippingPromotions201Response POSTFreeShippingPromotions(ctx).FreeShippingPromotionCreate(freeShippingPromotionCreate).Execute()

Create a free shipping promotion



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
    freeShippingPromotionCreate := *openapiclient.NewFreeShippingPromotionCreate(*openapiclient.NewFreeShippingPromotionCreateData(interface{}(123), *openapiclient.NewPOSTFreeShippingPromotions201ResponseDataAttributes(interface{}(Personal promotion), interface{}(2018-01-01T12:00:00.000Z), interface{}(2018-01-02T12:00:00.000Z)))) // FreeShippingPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FreeShippingPromotionsApi.POSTFreeShippingPromotions(context.Background()).FreeShippingPromotionCreate(freeShippingPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeShippingPromotionsApi.POSTFreeShippingPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTFreeShippingPromotions`: POSTFreeShippingPromotions201Response
    fmt.Fprintf(os.Stdout, "Response from `FreeShippingPromotionsApi.POSTFreeShippingPromotions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTFreeShippingPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **freeShippingPromotionCreate** | [**FreeShippingPromotionCreate**](FreeShippingPromotionCreate.md) |  | 

### Return type

[**POSTFreeShippingPromotions201Response**](POSTFreeShippingPromotions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

