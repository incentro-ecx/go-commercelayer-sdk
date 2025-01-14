# \MerchantsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEMerchantsMerchantId**](MerchantsApi.md#DELETEMerchantsMerchantId) | **Delete** /merchants/{merchantId} | Delete a merchant
[**GETMarketIdMerchant**](MerchantsApi.md#GETMarketIdMerchant) | **Get** /markets/{marketId}/merchant | Retrieve the merchant associated to the market
[**GETMerchants**](MerchantsApi.md#GETMerchants) | **Get** /merchants | List all merchants
[**GETMerchantsMerchantId**](MerchantsApi.md#GETMerchantsMerchantId) | **Get** /merchants/{merchantId} | Retrieve a merchant
[**GETStoreIdMerchant**](MerchantsApi.md#GETStoreIdMerchant) | **Get** /stores/{storeId}/merchant | Retrieve the merchant associated to the store
[**PATCHMerchantsMerchantId**](MerchantsApi.md#PATCHMerchantsMerchantId) | **Patch** /merchants/{merchantId} | Update a merchant
[**POSTMerchants**](MerchantsApi.md#POSTMerchants) | **Post** /merchants | Create a merchant



## DELETEMerchantsMerchantId

> DELETEMerchantsMerchantId(ctx, merchantId).Execute()

Delete a merchant



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
    merchantId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MerchantsApi.DELETEMerchantsMerchantId(context.Background(), merchantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsApi.DELETEMerchantsMerchantId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEMerchantsMerchantIdRequest struct via the builder pattern


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


## GETMarketIdMerchant

> GETMarketIdMerchant(ctx, marketId).Execute()

Retrieve the merchant associated to the market



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
    r, err := apiClient.MerchantsApi.GETMarketIdMerchant(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsApi.GETMarketIdMerchant``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETMarketIdMerchantRequest struct via the builder pattern


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


## GETMerchants

> GETMerchants200Response GETMerchants(ctx).Execute()

List all merchants



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
    resp, r, err := apiClient.MerchantsApi.GETMerchants(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsApi.GETMerchants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMerchants`: GETMerchants200Response
    fmt.Fprintf(os.Stdout, "Response from `MerchantsApi.GETMerchants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETMerchantsRequest struct via the builder pattern


### Return type

[**GETMerchants200Response**](GETMerchants200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETMerchantsMerchantId

> GETMerchantsMerchantId200Response GETMerchantsMerchantId(ctx, merchantId).Execute()

Retrieve a merchant



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
    merchantId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerchantsApi.GETMerchantsMerchantId(context.Background(), merchantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsApi.GETMerchantsMerchantId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETMerchantsMerchantId`: GETMerchantsMerchantId200Response
    fmt.Fprintf(os.Stdout, "Response from `MerchantsApi.GETMerchantsMerchantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMerchantsMerchantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETMerchantsMerchantId200Response**](GETMerchantsMerchantId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStoreIdMerchant

> GETStoreIdMerchant(ctx, storeId).Execute()

Retrieve the merchant associated to the store



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
    storeId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MerchantsApi.GETStoreIdMerchant(context.Background(), storeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsApi.GETStoreIdMerchant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storeId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStoreIdMerchantRequest struct via the builder pattern


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


## PATCHMerchantsMerchantId

> PATCHMerchantsMerchantId200Response PATCHMerchantsMerchantId(ctx, merchantId).MerchantUpdate(merchantUpdate).Execute()

Update a merchant



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
    merchantUpdate := *openapiclient.NewMerchantUpdate(*openapiclient.NewMerchantUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHMerchantsMerchantId200ResponseDataAttributes())) // MerchantUpdate | 
    merchantId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerchantsApi.PATCHMerchantsMerchantId(context.Background(), merchantId).MerchantUpdate(merchantUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsApi.PATCHMerchantsMerchantId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHMerchantsMerchantId`: PATCHMerchantsMerchantId200Response
    fmt.Fprintf(os.Stdout, "Response from `MerchantsApi.PATCHMerchantsMerchantId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**merchantId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHMerchantsMerchantIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantUpdate** | [**MerchantUpdate**](MerchantUpdate.md) |  | 


### Return type

[**PATCHMerchantsMerchantId200Response**](PATCHMerchantsMerchantId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTMerchants

> POSTMerchants201Response POSTMerchants(ctx).MerchantCreate(merchantCreate).Execute()

Create a merchant



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
    merchantCreate := *openapiclient.NewMerchantCreate(*openapiclient.NewMerchantCreateData(interface{}(123), *openapiclient.NewPOSTMerchants201ResponseDataAttributes(interface{}(The Brand Inc.)))) // MerchantCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerchantsApi.POSTMerchants(context.Background()).MerchantCreate(merchantCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerchantsApi.POSTMerchants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTMerchants`: POSTMerchants201Response
    fmt.Fprintf(os.Stdout, "Response from `MerchantsApi.POSTMerchants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTMerchantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **merchantCreate** | [**MerchantCreate**](MerchantCreate.md) |  | 

### Return type

[**POSTMerchants201Response**](POSTMerchants201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

