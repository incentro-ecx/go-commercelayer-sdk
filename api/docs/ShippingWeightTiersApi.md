# \ShippingWeightTiersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEShippingWeightTiersShippingWeightTierId**](ShippingWeightTiersApi.md#DELETEShippingWeightTiersShippingWeightTierId) | **Delete** /shipping_weight_tiers/{shippingWeightTierId} | Delete a shipping weight tier
[**GETShippingMethodIdShippingWeightTiers**](ShippingWeightTiersApi.md#GETShippingMethodIdShippingWeightTiers) | **Get** /shipping_methods/{shippingMethodId}/shipping_weight_tiers | Retrieve the shipping weight tiers associated to the shipping method
[**GETShippingWeightTiers**](ShippingWeightTiersApi.md#GETShippingWeightTiers) | **Get** /shipping_weight_tiers | List all shipping weight tiers
[**GETShippingWeightTiersShippingWeightTierId**](ShippingWeightTiersApi.md#GETShippingWeightTiersShippingWeightTierId) | **Get** /shipping_weight_tiers/{shippingWeightTierId} | Retrieve a shipping weight tier
[**PATCHShippingWeightTiersShippingWeightTierId**](ShippingWeightTiersApi.md#PATCHShippingWeightTiersShippingWeightTierId) | **Patch** /shipping_weight_tiers/{shippingWeightTierId} | Update a shipping weight tier
[**POSTShippingWeightTiers**](ShippingWeightTiersApi.md#POSTShippingWeightTiers) | **Post** /shipping_weight_tiers | Create a shipping weight tier



## DELETEShippingWeightTiersShippingWeightTierId

> DELETEShippingWeightTiersShippingWeightTierId(ctx, shippingWeightTierId).Execute()

Delete a shipping weight tier



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
    shippingWeightTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShippingWeightTiersApi.DELETEShippingWeightTiersShippingWeightTierId(context.Background(), shippingWeightTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.DELETEShippingWeightTiersShippingWeightTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingWeightTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEShippingWeightTiersShippingWeightTierIdRequest struct via the builder pattern


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


## GETShippingMethodIdShippingWeightTiers

> GETShippingMethodIdShippingWeightTiers(ctx, shippingMethodId).Execute()

Retrieve the shipping weight tiers associated to the shipping method



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
    shippingMethodId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShippingWeightTiersApi.GETShippingMethodIdShippingWeightTiers(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.GETShippingMethodIdShippingWeightTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingMethodId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodIdShippingWeightTiersRequest struct via the builder pattern


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


## GETShippingWeightTiers

> GETShippingWeightTiers200Response GETShippingWeightTiers(ctx).Execute()

List all shipping weight tiers



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
    resp, r, err := apiClient.ShippingWeightTiersApi.GETShippingWeightTiers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.GETShippingWeightTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingWeightTiers`: GETShippingWeightTiers200Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingWeightTiersApi.GETShippingWeightTiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingWeightTiersRequest struct via the builder pattern


### Return type

[**GETShippingWeightTiers200Response**](GETShippingWeightTiers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShippingWeightTiersShippingWeightTierId

> GETShippingWeightTiersShippingWeightTierId200Response GETShippingWeightTiersShippingWeightTierId(ctx, shippingWeightTierId).Execute()

Retrieve a shipping weight tier



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
    shippingWeightTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingWeightTiersApi.GETShippingWeightTiersShippingWeightTierId(context.Background(), shippingWeightTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.GETShippingWeightTiersShippingWeightTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingWeightTiersShippingWeightTierId`: GETShippingWeightTiersShippingWeightTierId200Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingWeightTiersApi.GETShippingWeightTiersShippingWeightTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingWeightTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingWeightTiersShippingWeightTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETShippingWeightTiersShippingWeightTierId200Response**](GETShippingWeightTiersShippingWeightTierId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHShippingWeightTiersShippingWeightTierId

> PATCHShippingWeightTiersShippingWeightTierId200Response PATCHShippingWeightTiersShippingWeightTierId(ctx, shippingWeightTierId).ShippingWeightTierUpdate(shippingWeightTierUpdate).Execute()

Update a shipping weight tier



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
    shippingWeightTierUpdate := *openapiclient.NewShippingWeightTierUpdate(*openapiclient.NewShippingWeightTierUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHShippingWeightTiersShippingWeightTierId200ResponseDataAttributes())) // ShippingWeightTierUpdate | 
    shippingWeightTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingWeightTiersApi.PATCHShippingWeightTiersShippingWeightTierId(context.Background(), shippingWeightTierId).ShippingWeightTierUpdate(shippingWeightTierUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.PATCHShippingWeightTiersShippingWeightTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHShippingWeightTiersShippingWeightTierId`: PATCHShippingWeightTiersShippingWeightTierId200Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingWeightTiersApi.PATCHShippingWeightTiersShippingWeightTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingWeightTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHShippingWeightTiersShippingWeightTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingWeightTierUpdate** | [**ShippingWeightTierUpdate**](ShippingWeightTierUpdate.md) |  | 


### Return type

[**PATCHShippingWeightTiersShippingWeightTierId200Response**](PATCHShippingWeightTiersShippingWeightTierId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTShippingWeightTiers

> POSTShippingWeightTiers201Response POSTShippingWeightTiers(ctx).ShippingWeightTierCreate(shippingWeightTierCreate).Execute()

Create a shipping weight tier



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
    shippingWeightTierCreate := *openapiclient.NewShippingWeightTierCreate(*openapiclient.NewShippingWeightTierCreateData(interface{}(123), *openapiclient.NewPOSTShippingWeightTiers201ResponseDataAttributes(interface{}(Light shipping under 3kg), interface{}(1000)))) // ShippingWeightTierCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingWeightTiersApi.POSTShippingWeightTiers(context.Background()).ShippingWeightTierCreate(shippingWeightTierCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingWeightTiersApi.POSTShippingWeightTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTShippingWeightTiers`: POSTShippingWeightTiers201Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingWeightTiersApi.POSTShippingWeightTiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTShippingWeightTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shippingWeightTierCreate** | [**ShippingWeightTierCreate**](ShippingWeightTierCreate.md) |  | 

### Return type

[**POSTShippingWeightTiers201Response**](POSTShippingWeightTiers201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

