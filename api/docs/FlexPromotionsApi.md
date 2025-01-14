# \FlexPromotionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEFlexPromotionsFlexPromotionId**](FlexPromotionsApi.md#DELETEFlexPromotionsFlexPromotionId) | **Delete** /flex_promotions/{flexPromotionId} | Delete a flex promotion
[**GETFlexPromotions**](FlexPromotionsApi.md#GETFlexPromotions) | **Get** /flex_promotions | List all flex promotions
[**GETFlexPromotionsFlexPromotionId**](FlexPromotionsApi.md#GETFlexPromotionsFlexPromotionId) | **Get** /flex_promotions/{flexPromotionId} | Retrieve a flex promotion
[**PATCHFlexPromotionsFlexPromotionId**](FlexPromotionsApi.md#PATCHFlexPromotionsFlexPromotionId) | **Patch** /flex_promotions/{flexPromotionId} | Update a flex promotion
[**POSTFlexPromotions**](FlexPromotionsApi.md#POSTFlexPromotions) | **Post** /flex_promotions | Create a flex promotion



## DELETEFlexPromotionsFlexPromotionId

> DELETEFlexPromotionsFlexPromotionId(ctx, flexPromotionId).Execute()

Delete a flex promotion



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
    flexPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FlexPromotionsApi.DELETEFlexPromotionsFlexPromotionId(context.Background(), flexPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexPromotionsApi.DELETEFlexPromotionsFlexPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flexPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEFlexPromotionsFlexPromotionIdRequest struct via the builder pattern


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


## GETFlexPromotions

> GETFlexPromotions200Response GETFlexPromotions(ctx).Execute()

List all flex promotions



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
    resp, r, err := apiClient.FlexPromotionsApi.GETFlexPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexPromotionsApi.GETFlexPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFlexPromotions`: GETFlexPromotions200Response
    fmt.Fprintf(os.Stdout, "Response from `FlexPromotionsApi.GETFlexPromotions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETFlexPromotionsRequest struct via the builder pattern


### Return type

[**GETFlexPromotions200Response**](GETFlexPromotions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETFlexPromotionsFlexPromotionId

> GETFlexPromotionsFlexPromotionId200Response GETFlexPromotionsFlexPromotionId(ctx, flexPromotionId).Execute()

Retrieve a flex promotion



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
    flexPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlexPromotionsApi.GETFlexPromotionsFlexPromotionId(context.Background(), flexPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexPromotionsApi.GETFlexPromotionsFlexPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETFlexPromotionsFlexPromotionId`: GETFlexPromotionsFlexPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `FlexPromotionsApi.GETFlexPromotionsFlexPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flexPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETFlexPromotionsFlexPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETFlexPromotionsFlexPromotionId200Response**](GETFlexPromotionsFlexPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHFlexPromotionsFlexPromotionId

> PATCHFlexPromotionsFlexPromotionId200Response PATCHFlexPromotionsFlexPromotionId(ctx, flexPromotionId).FlexPromotionUpdate(flexPromotionUpdate).Execute()

Update a flex promotion



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
    flexPromotionUpdate := *openapiclient.NewFlexPromotionUpdate(*openapiclient.NewFlexPromotionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHFlexPromotionsFlexPromotionId200ResponseDataAttributes())) // FlexPromotionUpdate | 
    flexPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlexPromotionsApi.PATCHFlexPromotionsFlexPromotionId(context.Background(), flexPromotionId).FlexPromotionUpdate(flexPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexPromotionsApi.PATCHFlexPromotionsFlexPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHFlexPromotionsFlexPromotionId`: PATCHFlexPromotionsFlexPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `FlexPromotionsApi.PATCHFlexPromotionsFlexPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flexPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHFlexPromotionsFlexPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flexPromotionUpdate** | [**FlexPromotionUpdate**](FlexPromotionUpdate.md) |  | 


### Return type

[**PATCHFlexPromotionsFlexPromotionId200Response**](PATCHFlexPromotionsFlexPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTFlexPromotions

> POSTFlexPromotions201Response POSTFlexPromotions(ctx).FlexPromotionCreate(flexPromotionCreate).Execute()

Create a flex promotion



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
    flexPromotionCreate := *openapiclient.NewFlexPromotionCreate(*openapiclient.NewFlexPromotionCreateData(interface{}(123), *openapiclient.NewPOSTFlexPromotions201ResponseDataAttributes(interface{}(Personal promotion), interface{}(2018-01-01T12:00:00.000Z), interface{}(2018-01-02T12:00:00.000Z), interface{}({})))) // FlexPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FlexPromotionsApi.POSTFlexPromotions(context.Background()).FlexPromotionCreate(flexPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FlexPromotionsApi.POSTFlexPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTFlexPromotions`: POSTFlexPromotions201Response
    fmt.Fprintf(os.Stdout, "Response from `FlexPromotionsApi.POSTFlexPromotions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTFlexPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flexPromotionCreate** | [**FlexPromotionCreate**](FlexPromotionCreate.md) |  | 

### Return type

[**POSTFlexPromotions201Response**](POSTFlexPromotions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

