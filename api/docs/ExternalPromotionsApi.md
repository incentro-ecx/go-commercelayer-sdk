# \ExternalPromotionsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEExternalPromotionsExternalPromotionId**](ExternalPromotionsApi.md#DELETEExternalPromotionsExternalPromotionId) | **Delete** /external_promotions/{externalPromotionId} | Delete an external promotion
[**GETExternalPromotions**](ExternalPromotionsApi.md#GETExternalPromotions) | **Get** /external_promotions | List all external promotions
[**GETExternalPromotionsExternalPromotionId**](ExternalPromotionsApi.md#GETExternalPromotionsExternalPromotionId) | **Get** /external_promotions/{externalPromotionId} | Retrieve an external promotion
[**PATCHExternalPromotionsExternalPromotionId**](ExternalPromotionsApi.md#PATCHExternalPromotionsExternalPromotionId) | **Patch** /external_promotions/{externalPromotionId} | Update an external promotion
[**POSTExternalPromotions**](ExternalPromotionsApi.md#POSTExternalPromotions) | **Post** /external_promotions | Create an external promotion



## DELETEExternalPromotionsExternalPromotionId

> DELETEExternalPromotionsExternalPromotionId(ctx, externalPromotionId).Execute()

Delete an external promotion



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
    externalPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExternalPromotionsApi.DELETEExternalPromotionsExternalPromotionId(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.DELETEExternalPromotionsExternalPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEExternalPromotionsExternalPromotionIdRequest struct via the builder pattern


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


## GETExternalPromotions

> GETExternalPromotions200Response GETExternalPromotions(ctx).Execute()

List all external promotions



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
    resp, r, err := apiClient.ExternalPromotionsApi.GETExternalPromotions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.GETExternalPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalPromotions`: GETExternalPromotions200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalPromotionsApi.GETExternalPromotions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPromotionsRequest struct via the builder pattern


### Return type

[**GETExternalPromotions200Response**](GETExternalPromotions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExternalPromotionsExternalPromotionId

> GETExternalPromotionsExternalPromotionId200Response GETExternalPromotionsExternalPromotionId(ctx, externalPromotionId).Execute()

Retrieve an external promotion



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
    externalPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPromotionsApi.GETExternalPromotionsExternalPromotionId(context.Background(), externalPromotionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.GETExternalPromotionsExternalPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalPromotionsExternalPromotionId`: GETExternalPromotionsExternalPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalPromotionsApi.GETExternalPromotionsExternalPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPromotionsExternalPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETExternalPromotionsExternalPromotionId200Response**](GETExternalPromotionsExternalPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHExternalPromotionsExternalPromotionId

> PATCHExternalPromotionsExternalPromotionId200Response PATCHExternalPromotionsExternalPromotionId(ctx, externalPromotionId).ExternalPromotionUpdate(externalPromotionUpdate).Execute()

Update an external promotion



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
    externalPromotionUpdate := *openapiclient.NewExternalPromotionUpdate(*openapiclient.NewExternalPromotionUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHExternalPromotionsExternalPromotionId200ResponseDataAttributes())) // ExternalPromotionUpdate | 
    externalPromotionId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPromotionsApi.PATCHExternalPromotionsExternalPromotionId(context.Background(), externalPromotionId).ExternalPromotionUpdate(externalPromotionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.PATCHExternalPromotionsExternalPromotionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHExternalPromotionsExternalPromotionId`: PATCHExternalPromotionsExternalPromotionId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalPromotionsApi.PATCHExternalPromotionsExternalPromotionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPromotionId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHExternalPromotionsExternalPromotionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalPromotionUpdate** | [**ExternalPromotionUpdate**](ExternalPromotionUpdate.md) |  | 


### Return type

[**PATCHExternalPromotionsExternalPromotionId200Response**](PATCHExternalPromotionsExternalPromotionId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTExternalPromotions

> POSTExternalPromotions201Response POSTExternalPromotions(ctx).ExternalPromotionCreate(externalPromotionCreate).Execute()

Create an external promotion



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
    externalPromotionCreate := *openapiclient.NewExternalPromotionCreate(*openapiclient.NewExternalPromotionCreateData(interface{}(123), *openapiclient.NewPOSTExternalPromotions201ResponseDataAttributes(interface{}(Personal promotion), interface{}(2018-01-01T12:00:00.000Z), interface{}(2018-01-02T12:00:00.000Z), interface{}(https://external_promotion.yourbrand.com)))) // ExternalPromotionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPromotionsApi.POSTExternalPromotions(context.Background()).ExternalPromotionCreate(externalPromotionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPromotionsApi.POSTExternalPromotions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTExternalPromotions`: POSTExternalPromotions201Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalPromotionsApi.POSTExternalPromotions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTExternalPromotionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalPromotionCreate** | [**ExternalPromotionCreate**](ExternalPromotionCreate.md) |  | 

### Return type

[**POSTExternalPromotions201Response**](POSTExternalPromotions201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

