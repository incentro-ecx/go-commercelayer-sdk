# \GeocodersApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETAddressIdGeocoder**](GeocodersApi.md#GETAddressIdGeocoder) | **Get** /addresses/{addressId}/geocoder | Retrieve the geocoder associated to the address
[**GETGeocoders**](GeocodersApi.md#GETGeocoders) | **Get** /geocoders | List all geocoders
[**GETGeocodersGeocoderId**](GeocodersApi.md#GETGeocodersGeocoderId) | **Get** /geocoders/{geocoderId} | Retrieve a geocoder
[**GETMarketIdGeocoder**](GeocodersApi.md#GETMarketIdGeocoder) | **Get** /markets/{marketId}/geocoder | Retrieve the geocoder associated to the market



## GETAddressIdGeocoder

> GETAddressIdGeocoder(ctx, addressId).Execute()

Retrieve the geocoder associated to the address



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
    addressId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GeocodersApi.GETAddressIdGeocoder(context.Background(), addressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeocodersApi.GETAddressIdGeocoder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addressId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAddressIdGeocoderRequest struct via the builder pattern


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


## GETGeocoders

> GETGeocoders200Response GETGeocoders(ctx).Execute()

List all geocoders



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
    resp, r, err := apiClient.GeocodersApi.GETGeocoders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeocodersApi.GETGeocoders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETGeocoders`: GETGeocoders200Response
    fmt.Fprintf(os.Stdout, "Response from `GeocodersApi.GETGeocoders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETGeocodersRequest struct via the builder pattern


### Return type

[**GETGeocoders200Response**](GETGeocoders200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETGeocodersGeocoderId

> GETGeocodersGeocoderId200Response GETGeocodersGeocoderId(ctx, geocoderId).Execute()

Retrieve a geocoder



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
    geocoderId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeocodersApi.GETGeocodersGeocoderId(context.Background(), geocoderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeocodersApi.GETGeocodersGeocoderId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETGeocodersGeocoderId`: GETGeocodersGeocoderId200Response
    fmt.Fprintf(os.Stdout, "Response from `GeocodersApi.GETGeocodersGeocoderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geocoderId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETGeocodersGeocoderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETGeocodersGeocoderId200Response**](GETGeocodersGeocoderId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETMarketIdGeocoder

> GETMarketIdGeocoder(ctx, marketId).Execute()

Retrieve the geocoder associated to the market



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
    marketId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GeocodersApi.GETMarketIdGeocoder(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeocodersApi.GETMarketIdGeocoder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketIdGeocoderRequest struct via the builder pattern


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

