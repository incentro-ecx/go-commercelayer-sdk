# \PriceFrequencyTiersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPriceFrequencyTiersPriceFrequencyTierId**](PriceFrequencyTiersApi.md#DELETEPriceFrequencyTiersPriceFrequencyTierId) | **Delete** /price_frequency_tiers/{priceFrequencyTierId} | Delete a price frequency tier
[**GETPriceFrequencyTiers**](PriceFrequencyTiersApi.md#GETPriceFrequencyTiers) | **Get** /price_frequency_tiers | List all price frequency tiers
[**GETPriceFrequencyTiersPriceFrequencyTierId**](PriceFrequencyTiersApi.md#GETPriceFrequencyTiersPriceFrequencyTierId) | **Get** /price_frequency_tiers/{priceFrequencyTierId} | Retrieve a price frequency tier
[**GETPriceIdPriceFrequencyTiers**](PriceFrequencyTiersApi.md#GETPriceIdPriceFrequencyTiers) | **Get** /prices/{priceId}/price_frequency_tiers | Retrieve the price frequency tiers associated to the price
[**PATCHPriceFrequencyTiersPriceFrequencyTierId**](PriceFrequencyTiersApi.md#PATCHPriceFrequencyTiersPriceFrequencyTierId) | **Patch** /price_frequency_tiers/{priceFrequencyTierId} | Update a price frequency tier
[**POSTPriceFrequencyTiers**](PriceFrequencyTiersApi.md#POSTPriceFrequencyTiers) | **Post** /price_frequency_tiers | Create a price frequency tier



## DELETEPriceFrequencyTiersPriceFrequencyTierId

> DELETEPriceFrequencyTiersPriceFrequencyTierId(ctx, priceFrequencyTierId).Execute()

Delete a price frequency tier



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
    priceFrequencyTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PriceFrequencyTiersApi.DELETEPriceFrequencyTiersPriceFrequencyTierId(context.Background(), priceFrequencyTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceFrequencyTiersApi.DELETEPriceFrequencyTiersPriceFrequencyTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceFrequencyTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPriceFrequencyTiersPriceFrequencyTierIdRequest struct via the builder pattern


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


## GETPriceFrequencyTiers

> GETPriceFrequencyTiers200Response GETPriceFrequencyTiers(ctx).Execute()

List all price frequency tiers



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
    resp, r, err := apiClient.PriceFrequencyTiersApi.GETPriceFrequencyTiers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceFrequencyTiersApi.GETPriceFrequencyTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceFrequencyTiers`: GETPriceFrequencyTiers200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceFrequencyTiersApi.GETPriceFrequencyTiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceFrequencyTiersRequest struct via the builder pattern


### Return type

[**GETPriceFrequencyTiers200Response**](GETPriceFrequencyTiers200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPriceFrequencyTiersPriceFrequencyTierId

> GETPriceFrequencyTiersPriceFrequencyTierId200Response GETPriceFrequencyTiersPriceFrequencyTierId(ctx, priceFrequencyTierId).Execute()

Retrieve a price frequency tier



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
    priceFrequencyTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceFrequencyTiersApi.GETPriceFrequencyTiersPriceFrequencyTierId(context.Background(), priceFrequencyTierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceFrequencyTiersApi.GETPriceFrequencyTiersPriceFrequencyTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPriceFrequencyTiersPriceFrequencyTierId`: GETPriceFrequencyTiersPriceFrequencyTierId200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceFrequencyTiersApi.GETPriceFrequencyTiersPriceFrequencyTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceFrequencyTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceFrequencyTiersPriceFrequencyTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPriceFrequencyTiersPriceFrequencyTierId200Response**](GETPriceFrequencyTiersPriceFrequencyTierId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPriceIdPriceFrequencyTiers

> GETPriceIdPriceFrequencyTiers(ctx, priceId).Execute()

Retrieve the price frequency tiers associated to the price



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
    priceId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PriceFrequencyTiersApi.GETPriceIdPriceFrequencyTiers(context.Background(), priceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceFrequencyTiersApi.GETPriceIdPriceFrequencyTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPriceIdPriceFrequencyTiersRequest struct via the builder pattern


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


## PATCHPriceFrequencyTiersPriceFrequencyTierId

> PATCHPriceFrequencyTiersPriceFrequencyTierId200Response PATCHPriceFrequencyTiersPriceFrequencyTierId(ctx, priceFrequencyTierId).PriceFrequencyTierUpdate(priceFrequencyTierUpdate).Execute()

Update a price frequency tier



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
    priceFrequencyTierUpdate := *openapiclient.NewPriceFrequencyTierUpdate(*openapiclient.NewPriceFrequencyTierUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPriceFrequencyTiersPriceFrequencyTierId200ResponseDataAttributes())) // PriceFrequencyTierUpdate | 
    priceFrequencyTierId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceFrequencyTiersApi.PATCHPriceFrequencyTiersPriceFrequencyTierId(context.Background(), priceFrequencyTierId).PriceFrequencyTierUpdate(priceFrequencyTierUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceFrequencyTiersApi.PATCHPriceFrequencyTiersPriceFrequencyTierId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPriceFrequencyTiersPriceFrequencyTierId`: PATCHPriceFrequencyTiersPriceFrequencyTierId200Response
    fmt.Fprintf(os.Stdout, "Response from `PriceFrequencyTiersApi.PATCHPriceFrequencyTiersPriceFrequencyTierId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**priceFrequencyTierId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPriceFrequencyTiersPriceFrequencyTierIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceFrequencyTierUpdate** | [**PriceFrequencyTierUpdate**](PriceFrequencyTierUpdate.md) |  | 


### Return type

[**PATCHPriceFrequencyTiersPriceFrequencyTierId200Response**](PATCHPriceFrequencyTiersPriceFrequencyTierId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPriceFrequencyTiers

> POSTPriceFrequencyTiers201Response POSTPriceFrequencyTiers(ctx).PriceFrequencyTierCreate(priceFrequencyTierCreate).Execute()

Create a price frequency tier



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
    priceFrequencyTierCreate := *openapiclient.NewPriceFrequencyTierCreate(*openapiclient.NewPriceFrequencyTierCreateData(interface{}(123), *openapiclient.NewPOSTPriceFrequencyTiers201ResponseDataAttributes(interface{}(six pack), interface{}(1000)))) // PriceFrequencyTierCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceFrequencyTiersApi.POSTPriceFrequencyTiers(context.Background()).PriceFrequencyTierCreate(priceFrequencyTierCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceFrequencyTiersApi.POSTPriceFrequencyTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPriceFrequencyTiers`: POSTPriceFrequencyTiers201Response
    fmt.Fprintf(os.Stdout, "Response from `PriceFrequencyTiersApi.POSTPriceFrequencyTiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPriceFrequencyTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priceFrequencyTierCreate** | [**PriceFrequencyTierCreate**](PriceFrequencyTierCreate.md) |  | 

### Return type

[**POSTPriceFrequencyTiers201Response**](POSTPriceFrequencyTiers201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

