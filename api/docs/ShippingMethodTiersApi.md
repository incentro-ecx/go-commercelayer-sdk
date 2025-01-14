# \ShippingMethodTiersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETShippingMethodIdShippingMethodTiers**](ShippingMethodTiersApi.md#GETShippingMethodIdShippingMethodTiers) | **Get** /shipping_methods/{shippingMethodId}/shipping_method_tiers | Retrieve the shipping method tiers associated to the shipping method
[**GETShippingMethodTiers**](ShippingMethodTiersApi.md#GETShippingMethodTiers) | **Get** /shipping_method_tiers | List all shipping method tiers
[**GETShippingMethodTiersShippingMethodTierId**](ShippingMethodTiersApi.md#GETShippingMethodTiersShippingMethodTierId) | **Get** /shipping_method_tiers/{shippingMethodTierId} | Retrieve a shipping method tier



## GETShippingMethodIdShippingMethodTiers

> GETShippingMethodIdShippingMethodTiers(ctx, shippingMethodId).Execute()

Retrieve the shipping method tiers associated to the shipping method



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
    r, err := apiClient.ShippingMethodTiersApi.GETShippingMethodIdShippingMethodTiers(context.Background(), shippingMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodTiersApi.GETShippingMethodIdShippingMethodTiers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETShippingMethodIdShippingMethodTiersRequest struct via the builder pattern


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


## GETShippingMethodTiers

> GETShippingMethodTiers200Response GETShippingMethodTiers(ctx).Execute()

List all shipping method tiers



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
    resp, r, err := apiClient.ShippingMethodTiersApi.GETShippingMethodTiers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodTiersApi.GETShippingMethodTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingMethodTiers`: GETShippingMethodTiers200Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingMethodTiersApi.GETShippingMethodTiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodTiersRequest struct via the builder pattern


### Return type

[**GETShippingMethodTiers200Response**](GETShippingMethodTiers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETShippingMethodTiersShippingMethodTierId

> GETShippingMethodTiersShippingMethodTierId200Response GETShippingMethodTiersShippingMethodTierId(ctx, shippingMethodTierId).Execute()

Retrieve a shipping method tier



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
    shippingMethodTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingMethodTiersApi.GETShippingMethodTiersShippingMethodTierId(context.Background(), shippingMethodTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingMethodTiersApi.GETShippingMethodTiersShippingMethodTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETShippingMethodTiersShippingMethodTierId`: GETShippingMethodTiersShippingMethodTierId200Response
    fmt.Fprintf(os.Stdout, "Response from `ShippingMethodTiersApi.GETShippingMethodTiersShippingMethodTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shippingMethodTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETShippingMethodTiersShippingMethodTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETShippingMethodTiersShippingMethodTierId200Response**](GETShippingMethodTiersShippingMethodTierId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

