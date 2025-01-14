# \StripeGatewaysApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEStripeGatewaysStripeGatewayId**](StripeGatewaysApi.md#DELETEStripeGatewaysStripeGatewayId) | **Delete** /stripe_gateways/{stripeGatewayId} | Delete a stripe gateway
[**GETStripeGateways**](StripeGatewaysApi.md#GETStripeGateways) | **Get** /stripe_gateways | List all stripe gateways
[**GETStripeGatewaysStripeGatewayId**](StripeGatewaysApi.md#GETStripeGatewaysStripeGatewayId) | **Get** /stripe_gateways/{stripeGatewayId} | Retrieve a stripe gateway
[**PATCHStripeGatewaysStripeGatewayId**](StripeGatewaysApi.md#PATCHStripeGatewaysStripeGatewayId) | **Patch** /stripe_gateways/{stripeGatewayId} | Update a stripe gateway
[**POSTStripeGateways**](StripeGatewaysApi.md#POSTStripeGateways) | **Post** /stripe_gateways | Create a stripe gateway



## DELETEStripeGatewaysStripeGatewayId

> DELETEStripeGatewaysStripeGatewayId(ctx, stripeGatewayId).Execute()

Delete a stripe gateway



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
    stripeGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StripeGatewaysApi.DELETEStripeGatewaysStripeGatewayId(context.Background(), stripeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.DELETEStripeGatewaysStripeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripeGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEStripeGatewaysStripeGatewayIdRequest struct via the builder pattern


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


## GETStripeGateways

> GETStripeGateways200Response GETStripeGateways(ctx).Execute()

List all stripe gateways



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
    resp, r, err := apiClient.StripeGatewaysApi.GETStripeGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.GETStripeGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStripeGateways`: GETStripeGateways200Response
    fmt.Fprintf(os.Stdout, "Response from `StripeGatewaysApi.GETStripeGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETStripeGatewaysRequest struct via the builder pattern


### Return type

[**GETStripeGateways200Response**](GETStripeGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETStripeGatewaysStripeGatewayId

> GETStripeGatewaysStripeGatewayId200Response GETStripeGatewaysStripeGatewayId(ctx, stripeGatewayId).Execute()

Retrieve a stripe gateway



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
    stripeGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripeGatewaysApi.GETStripeGatewaysStripeGatewayId(context.Background(), stripeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.GETStripeGatewaysStripeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETStripeGatewaysStripeGatewayId`: GETStripeGatewaysStripeGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `StripeGatewaysApi.GETStripeGatewaysStripeGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripeGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETStripeGatewaysStripeGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETStripeGatewaysStripeGatewayId200Response**](GETStripeGatewaysStripeGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHStripeGatewaysStripeGatewayId

> PATCHStripeGatewaysStripeGatewayId200Response PATCHStripeGatewaysStripeGatewayId(ctx, stripeGatewayId).StripeGatewayUpdate(stripeGatewayUpdate).Execute()

Update a stripe gateway



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
    stripeGatewayUpdate := *openapiclient.NewStripeGatewayUpdate(*openapiclient.NewStripeGatewayUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHStripeGatewaysStripeGatewayId200ResponseDataAttributes())) // StripeGatewayUpdate | 
    stripeGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripeGatewaysApi.PATCHStripeGatewaysStripeGatewayId(context.Background(), stripeGatewayId).StripeGatewayUpdate(stripeGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.PATCHStripeGatewaysStripeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHStripeGatewaysStripeGatewayId`: PATCHStripeGatewaysStripeGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `StripeGatewaysApi.PATCHStripeGatewaysStripeGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stripeGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHStripeGatewaysStripeGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stripeGatewayUpdate** | [**StripeGatewayUpdate**](StripeGatewayUpdate.md) |  | 


### Return type

[**PATCHStripeGatewaysStripeGatewayId200Response**](PATCHStripeGatewaysStripeGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTStripeGateways

> POSTStripeGateways201Response POSTStripeGateways(ctx).StripeGatewayCreate(stripeGatewayCreate).Execute()

Create a stripe gateway



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
    stripeGatewayCreate := *openapiclient.NewStripeGatewayCreate(*openapiclient.NewStripeGatewayCreateData(interface{}(123), *openapiclient.NewPOSTStripeGateways201ResponseDataAttributes(interface{}(US payment gateway), interface{}(sk_live_xxxx-yyyy-zzzz)))) // StripeGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StripeGatewaysApi.POSTStripeGateways(context.Background()).StripeGatewayCreate(stripeGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StripeGatewaysApi.POSTStripeGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTStripeGateways`: POSTStripeGateways201Response
    fmt.Fprintf(os.Stdout, "Response from `StripeGatewaysApi.POSTStripeGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTStripeGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stripeGatewayCreate** | [**StripeGatewayCreate**](StripeGatewayCreate.md) |  | 

### Return type

[**POSTStripeGateways201Response**](POSTStripeGateways201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

