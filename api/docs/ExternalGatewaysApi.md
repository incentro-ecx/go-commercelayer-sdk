# \ExternalGatewaysApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEExternalGatewaysExternalGatewayId**](ExternalGatewaysApi.md#DELETEExternalGatewaysExternalGatewayId) | **Delete** /external_gateways/{externalGatewayId} | Delete an external gateway
[**GETExternalGateways**](ExternalGatewaysApi.md#GETExternalGateways) | **Get** /external_gateways | List all external gateways
[**GETExternalGatewaysExternalGatewayId**](ExternalGatewaysApi.md#GETExternalGatewaysExternalGatewayId) | **Get** /external_gateways/{externalGatewayId} | Retrieve an external gateway
[**PATCHExternalGatewaysExternalGatewayId**](ExternalGatewaysApi.md#PATCHExternalGatewaysExternalGatewayId) | **Patch** /external_gateways/{externalGatewayId} | Update an external gateway
[**POSTExternalGateways**](ExternalGatewaysApi.md#POSTExternalGateways) | **Post** /external_gateways | Create an external gateway



## DELETEExternalGatewaysExternalGatewayId

> DELETEExternalGatewaysExternalGatewayId(ctx, externalGatewayId).Execute()

Delete an external gateway



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
    externalGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExternalGatewaysApi.DELETEExternalGatewaysExternalGatewayId(context.Background(), externalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.DELETEExternalGatewaysExternalGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEExternalGatewaysExternalGatewayIdRequest struct via the builder pattern


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


## GETExternalGateways

> GETExternalGateways200Response GETExternalGateways(ctx).Execute()

List all external gateways



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
    resp, r, err := apiClient.ExternalGatewaysApi.GETExternalGateways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.GETExternalGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalGateways`: GETExternalGateways200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalGatewaysApi.GETExternalGateways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalGatewaysRequest struct via the builder pattern


### Return type

[**GETExternalGateways200Response**](GETExternalGateways200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExternalGatewaysExternalGatewayId

> GETExternalGatewaysExternalGatewayId200Response GETExternalGatewaysExternalGatewayId(ctx, externalGatewayId).Execute()

Retrieve an external gateway



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
    externalGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalGatewaysApi.GETExternalGatewaysExternalGatewayId(context.Background(), externalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.GETExternalGatewaysExternalGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalGatewaysExternalGatewayId`: GETExternalGatewaysExternalGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalGatewaysApi.GETExternalGatewaysExternalGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalGatewaysExternalGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETExternalGatewaysExternalGatewayId200Response**](GETExternalGatewaysExternalGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHExternalGatewaysExternalGatewayId

> PATCHExternalGatewaysExternalGatewayId200Response PATCHExternalGatewaysExternalGatewayId(ctx, externalGatewayId).ExternalGatewayUpdate(externalGatewayUpdate).Execute()

Update an external gateway



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
    externalGatewayUpdate := *openapiclient.NewExternalGatewayUpdate(*openapiclient.NewExternalGatewayUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHExternalGatewaysExternalGatewayId200ResponseDataAttributes())) // ExternalGatewayUpdate | 
    externalGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalGatewaysApi.PATCHExternalGatewaysExternalGatewayId(context.Background(), externalGatewayId).ExternalGatewayUpdate(externalGatewayUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.PATCHExternalGatewaysExternalGatewayId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHExternalGatewaysExternalGatewayId`: PATCHExternalGatewaysExternalGatewayId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalGatewaysApi.PATCHExternalGatewaysExternalGatewayId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHExternalGatewaysExternalGatewayIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalGatewayUpdate** | [**ExternalGatewayUpdate**](ExternalGatewayUpdate.md) |  | 


### Return type

[**PATCHExternalGatewaysExternalGatewayId200Response**](PATCHExternalGatewaysExternalGatewayId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTExternalGateways

> POSTExternalGateways201Response POSTExternalGateways(ctx).ExternalGatewayCreate(externalGatewayCreate).Execute()

Create an external gateway



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
    externalGatewayCreate := *openapiclient.NewExternalGatewayCreate(*openapiclient.NewExternalGatewayCreateData(interface{}(123), *openapiclient.NewPOSTExternalGateways201ResponseDataAttributes(interface{}(US payment gateway)))) // ExternalGatewayCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalGatewaysApi.POSTExternalGateways(context.Background()).ExternalGatewayCreate(externalGatewayCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalGatewaysApi.POSTExternalGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTExternalGateways`: POSTExternalGateways201Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalGatewaysApi.POSTExternalGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTExternalGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalGatewayCreate** | [**ExternalGatewayCreate**](ExternalGatewayCreate.md) |  | 

### Return type

[**POSTExternalGateways201Response**](POSTExternalGateways201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

