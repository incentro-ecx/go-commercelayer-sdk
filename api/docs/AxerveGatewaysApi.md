# \AxerveGatewaysApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAxerveGatewaysAxerveGatewayId**](AxerveGatewaysApi.md#DELETEAxerveGatewaysAxerveGatewayId) | **Delete** /axerve_gateways/{axerveGatewayId} | Delete an axerve gateway
[**GETAxerveGateways**](AxerveGatewaysApi.md#GETAxerveGateways) | **Get** /axerve_gateways | List all axerve gateways
[**GETAxerveGatewaysAxerveGatewayId**](AxerveGatewaysApi.md#GETAxerveGatewaysAxerveGatewayId) | **Get** /axerve_gateways/{axerveGatewayId} | Retrieve an axerve gateway
[**PATCHAxerveGatewaysAxerveGatewayId**](AxerveGatewaysApi.md#PATCHAxerveGatewaysAxerveGatewayId) | **Patch** /axerve_gateways/{axerveGatewayId} | Update an axerve gateway
[**POSTAxerveGateways**](AxerveGatewaysApi.md#POSTAxerveGateways) | **Post** /axerve_gateways | Create an axerve gateway



## DELETEAxerveGatewaysAxerveGatewayId

> DELETEAxerveGatewaysAxerveGatewayId(ctx, axerveGatewayId).Execute()

Delete an axerve gateway



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
    axerveGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AxerveGatewaysApi.DELETEAxerveGatewaysAxerveGatewayId(context.Background(), axerveGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxerveGatewaysApi.DELETEAxerveGatewaysAxerveGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**axerveGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEAxerveGatewaysAxerveGatewayIdRequest struct via the builder pattern


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


## GETAxerveGateways

> GETAxerveGateways200Response GETAxerveGateways(ctx).Execute()

List all axerve gateways



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
    resp, r, err := apiClient.AxerveGatewaysApi.GETAxerveGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxerveGatewaysApi.GETAxerveGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAxerveGateways`: GETAxerveGateways200Response
    fmt.Fprintf(os.Stdout, "Response from `AxerveGatewaysApi.GETAxerveGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETAxerveGatewaysRequest struct via the builder pattern


### Return type

[**GETAxerveGateways200Response**](GETAxerveGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETAxerveGatewaysAxerveGatewayId

> GETAxerveGatewaysAxerveGatewayId200Response GETAxerveGatewaysAxerveGatewayId(ctx, axerveGatewayId).Execute()

Retrieve an axerve gateway



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
    axerveGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AxerveGatewaysApi.GETAxerveGatewaysAxerveGatewayId(context.Background(), axerveGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxerveGatewaysApi.GETAxerveGatewaysAxerveGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAxerveGatewaysAxerveGatewayId`: GETAxerveGatewaysAxerveGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `AxerveGatewaysApi.GETAxerveGatewaysAxerveGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**axerveGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAxerveGatewaysAxerveGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETAxerveGatewaysAxerveGatewayId200Response**](GETAxerveGatewaysAxerveGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHAxerveGatewaysAxerveGatewayId

> PATCHAxerveGatewaysAxerveGatewayId200Response PATCHAxerveGatewaysAxerveGatewayId(ctx, axerveGatewayId).AxerveGatewayUpdate(axerveGatewayUpdate).Execute()

Update an axerve gateway



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
    axerveGatewayUpdate := *openapiclient.NewAxerveGatewayUpdate(*openapiclient.NewAxerveGatewayUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHAxerveGatewaysAxerveGatewayId200ResponseDataAttributes())) // AxerveGatewayUpdate | 
    axerveGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AxerveGatewaysApi.PATCHAxerveGatewaysAxerveGatewayId(context.Background(), axerveGatewayId).AxerveGatewayUpdate(axerveGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxerveGatewaysApi.PATCHAxerveGatewaysAxerveGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHAxerveGatewaysAxerveGatewayId`: PATCHAxerveGatewaysAxerveGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `AxerveGatewaysApi.PATCHAxerveGatewaysAxerveGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**axerveGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHAxerveGatewaysAxerveGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **axerveGatewayUpdate** | [**AxerveGatewayUpdate**](AxerveGatewayUpdate.md) |  | 


### Return type

[**PATCHAxerveGatewaysAxerveGatewayId200Response**](PATCHAxerveGatewaysAxerveGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTAxerveGateways

> POSTAxerveGateways201Response POSTAxerveGateways(ctx).AxerveGatewayCreate(axerveGatewayCreate).Execute()

Create an axerve gateway



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
    axerveGatewayCreate := *openapiclient.NewAxerveGatewayCreate(*openapiclient.NewAxerveGatewayCreateData(interface{}(123), *openapiclient.NewPOSTAxerveGateways201ResponseDataAttributes(interface{}(US payment gateway), interface{}(xxxx-yyyy-zzzz), interface{}(xxxx-yyyy-zzzz)))) // AxerveGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AxerveGatewaysApi.POSTAxerveGateways(context.Background()).AxerveGatewayCreate(axerveGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxerveGatewaysApi.POSTAxerveGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTAxerveGateways`: POSTAxerveGateways201Response
    fmt.Fprintf(os.Stdout, "Response from `AxerveGatewaysApi.POSTAxerveGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTAxerveGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **axerveGatewayCreate** | [**AxerveGatewayCreate**](AxerveGatewayCreate.md) |  | 

### Return type

[**POSTAxerveGateways201Response**](POSTAxerveGateways201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

