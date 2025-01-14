# \PaypalGatewaysApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEPaypalGatewaysPaypalGatewayId**](PaypalGatewaysApi.md#DELETEPaypalGatewaysPaypalGatewayId) | **Delete** /paypal_gateways/{paypalGatewayId} | Delete a paypal gateway
[**GETPaypalGateways**](PaypalGatewaysApi.md#GETPaypalGateways) | **Get** /paypal_gateways | List all paypal gateways
[**GETPaypalGatewaysPaypalGatewayId**](PaypalGatewaysApi.md#GETPaypalGatewaysPaypalGatewayId) | **Get** /paypal_gateways/{paypalGatewayId} | Retrieve a paypal gateway
[**PATCHPaypalGatewaysPaypalGatewayId**](PaypalGatewaysApi.md#PATCHPaypalGatewaysPaypalGatewayId) | **Patch** /paypal_gateways/{paypalGatewayId} | Update a paypal gateway
[**POSTPaypalGateways**](PaypalGatewaysApi.md#POSTPaypalGateways) | **Post** /paypal_gateways | Create a paypal gateway



## DELETEPaypalGatewaysPaypalGatewayId

> DELETEPaypalGatewaysPaypalGatewayId(ctx, paypalGatewayId).Execute()

Delete a paypal gateway



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
    paypalGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaypalGatewaysApi.DELETEPaypalGatewaysPaypalGatewayId(context.Background(), paypalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalGatewaysApi.DELETEPaypalGatewaysPaypalGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEPaypalGatewaysPaypalGatewayIdRequest struct via the builder pattern


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


## GETPaypalGateways

> GETPaypalGateways200Response GETPaypalGateways(ctx).Execute()

List all paypal gateways



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
    resp, r, err := apiClient.PaypalGatewaysApi.GETPaypalGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalGatewaysApi.GETPaypalGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaypalGateways`: GETPaypalGateways200Response
    fmt.Fprintf(os.Stdout, "Response from `PaypalGatewaysApi.GETPaypalGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaypalGatewaysRequest struct via the builder pattern


### Return type

[**GETPaypalGateways200Response**](GETPaypalGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPaypalGatewaysPaypalGatewayId

> GETPaypalGatewaysPaypalGatewayId200Response GETPaypalGatewaysPaypalGatewayId(ctx, paypalGatewayId).Execute()

Retrieve a paypal gateway



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
    paypalGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalGatewaysApi.GETPaypalGatewaysPaypalGatewayId(context.Background(), paypalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalGatewaysApi.GETPaypalGatewaysPaypalGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPaypalGatewaysPaypalGatewayId`: GETPaypalGatewaysPaypalGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `PaypalGatewaysApi.GETPaypalGatewaysPaypalGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPaypalGatewaysPaypalGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPaypalGatewaysPaypalGatewayId200Response**](GETPaypalGatewaysPaypalGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHPaypalGatewaysPaypalGatewayId

> PATCHPaypalGatewaysPaypalGatewayId200Response PATCHPaypalGatewaysPaypalGatewayId(ctx, paypalGatewayId).PaypalGatewayUpdate(paypalGatewayUpdate).Execute()

Update a paypal gateway



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
    paypalGatewayUpdate := *openapiclient.NewPaypalGatewayUpdate(*openapiclient.NewPaypalGatewayUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHPaypalGatewaysPaypalGatewayId200ResponseDataAttributes())) // PaypalGatewayUpdate | 
    paypalGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalGatewaysApi.PATCHPaypalGatewaysPaypalGatewayId(context.Background(), paypalGatewayId).PaypalGatewayUpdate(paypalGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalGatewaysApi.PATCHPaypalGatewaysPaypalGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHPaypalGatewaysPaypalGatewayId`: PATCHPaypalGatewaysPaypalGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `PaypalGatewaysApi.PATCHPaypalGatewaysPaypalGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paypalGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHPaypalGatewaysPaypalGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paypalGatewayUpdate** | [**PaypalGatewayUpdate**](PaypalGatewayUpdate.md) |  | 


### Return type

[**PATCHPaypalGatewaysPaypalGatewayId200Response**](PATCHPaypalGatewaysPaypalGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTPaypalGateways

> POSTPaypalGateways201Response POSTPaypalGateways(ctx).PaypalGatewayCreate(paypalGatewayCreate).Execute()

Create a paypal gateway



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
    paypalGatewayCreate := *openapiclient.NewPaypalGatewayCreate(*openapiclient.NewPaypalGatewayCreateData(interface{}(123), *openapiclient.NewPOSTPaypalGateways201ResponseDataAttributes(interface{}(US payment gateway), interface{}(xxxx-yyyy-zzzz), interface{}(xxxx-yyyy-zzzz)))) // PaypalGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaypalGatewaysApi.POSTPaypalGateways(context.Background()).PaypalGatewayCreate(paypalGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaypalGatewaysApi.POSTPaypalGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTPaypalGateways`: POSTPaypalGateways201Response
    fmt.Fprintf(os.Stdout, "Response from `PaypalGatewaysApi.POSTPaypalGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTPaypalGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paypalGatewayCreate** | [**PaypalGatewayCreate**](PaypalGatewayCreate.md) |  | 

### Return type

[**POSTPaypalGateways201Response**](POSTPaypalGateways201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

