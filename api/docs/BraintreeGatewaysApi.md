# \BraintreeGatewaysApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBraintreeGatewaysBraintreeGatewayId**](BraintreeGatewaysApi.md#DELETEBraintreeGatewaysBraintreeGatewayId) | **Delete** /braintree_gateways/{braintreeGatewayId} | Delete a braintree gateway
[**GETBraintreeGateways**](BraintreeGatewaysApi.md#GETBraintreeGateways) | **Get** /braintree_gateways | List all braintree gateways
[**GETBraintreeGatewaysBraintreeGatewayId**](BraintreeGatewaysApi.md#GETBraintreeGatewaysBraintreeGatewayId) | **Get** /braintree_gateways/{braintreeGatewayId} | Retrieve a braintree gateway
[**PATCHBraintreeGatewaysBraintreeGatewayId**](BraintreeGatewaysApi.md#PATCHBraintreeGatewaysBraintreeGatewayId) | **Patch** /braintree_gateways/{braintreeGatewayId} | Update a braintree gateway
[**POSTBraintreeGateways**](BraintreeGatewaysApi.md#POSTBraintreeGateways) | **Post** /braintree_gateways | Create a braintree gateway



## DELETEBraintreeGatewaysBraintreeGatewayId

> DELETEBraintreeGatewaysBraintreeGatewayId(ctx, braintreeGatewayId).Execute()

Delete a braintree gateway



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
    braintreeGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BraintreeGatewaysApi.DELETEBraintreeGatewaysBraintreeGatewayId(context.Background(), braintreeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.DELETEBraintreeGatewaysBraintreeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreeGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBraintreeGatewaysBraintreeGatewayIdRequest struct via the builder pattern


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


## GETBraintreeGateways

> GETBraintreeGateways200Response GETBraintreeGateways(ctx).Execute()

List all braintree gateways



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
    resp, r, err := apiClient.BraintreeGatewaysApi.GETBraintreeGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.GETBraintreeGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBraintreeGateways`: GETBraintreeGateways200Response
    fmt.Fprintf(os.Stdout, "Response from `BraintreeGatewaysApi.GETBraintreeGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreeGatewaysRequest struct via the builder pattern


### Return type

[**GETBraintreeGateways200Response**](GETBraintreeGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETBraintreeGatewaysBraintreeGatewayId

> GETBraintreeGatewaysBraintreeGatewayId200Response GETBraintreeGatewaysBraintreeGatewayId(ctx, braintreeGatewayId).Execute()

Retrieve a braintree gateway



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
    braintreeGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreeGatewaysApi.GETBraintreeGatewaysBraintreeGatewayId(context.Background(), braintreeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.GETBraintreeGatewaysBraintreeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBraintreeGatewaysBraintreeGatewayId`: GETBraintreeGatewaysBraintreeGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `BraintreeGatewaysApi.GETBraintreeGatewaysBraintreeGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreeGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreeGatewaysBraintreeGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETBraintreeGatewaysBraintreeGatewayId200Response**](GETBraintreeGatewaysBraintreeGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHBraintreeGatewaysBraintreeGatewayId

> PATCHBraintreeGatewaysBraintreeGatewayId200Response PATCHBraintreeGatewaysBraintreeGatewayId(ctx, braintreeGatewayId).BraintreeGatewayUpdate(braintreeGatewayUpdate).Execute()

Update a braintree gateway



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
    braintreeGatewayUpdate := *openapiclient.NewBraintreeGatewayUpdate(*openapiclient.NewBraintreeGatewayUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHBraintreeGatewaysBraintreeGatewayId200ResponseDataAttributes())) // BraintreeGatewayUpdate | 
    braintreeGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreeGatewaysApi.PATCHBraintreeGatewaysBraintreeGatewayId(context.Background(), braintreeGatewayId).BraintreeGatewayUpdate(braintreeGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.PATCHBraintreeGatewaysBraintreeGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHBraintreeGatewaysBraintreeGatewayId`: PATCHBraintreeGatewaysBraintreeGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `BraintreeGatewaysApi.PATCHBraintreeGatewaysBraintreeGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreeGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBraintreeGatewaysBraintreeGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **braintreeGatewayUpdate** | [**BraintreeGatewayUpdate**](BraintreeGatewayUpdate.md) |  | 


### Return type

[**PATCHBraintreeGatewaysBraintreeGatewayId200Response**](PATCHBraintreeGatewaysBraintreeGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTBraintreeGateways

> POSTBraintreeGateways201Response POSTBraintreeGateways(ctx).BraintreeGatewayCreate(braintreeGatewayCreate).Execute()

Create a braintree gateway



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
    braintreeGatewayCreate := *openapiclient.NewBraintreeGatewayCreate(*openapiclient.NewBraintreeGatewayCreateData(interface{}(123), *openapiclient.NewPOSTBraintreeGateways201ResponseDataAttributes(interface{}(US payment gateway), interface{}(xxxx-yyyy-zzzz), interface{}(xxxx-yyyy-zzzz), interface{}(xxxx-yyyy-zzzz), interface{}(xxxx-yyyy-zzzz)))) // BraintreeGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreeGatewaysApi.POSTBraintreeGateways(context.Background()).BraintreeGatewayCreate(braintreeGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreeGatewaysApi.POSTBraintreeGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTBraintreeGateways`: POSTBraintreeGateways201Response
    fmt.Fprintf(os.Stdout, "Response from `BraintreeGatewaysApi.POSTBraintreeGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBraintreeGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **braintreeGatewayCreate** | [**BraintreeGatewayCreate**](BraintreeGatewayCreate.md) |  | 

### Return type

[**POSTBraintreeGateways201Response**](POSTBraintreeGateways201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

