# \ExternalPaymentsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEExternalPaymentsExternalPaymentId**](ExternalPaymentsApi.md#DELETEExternalPaymentsExternalPaymentId) | **Delete** /external_payments/{externalPaymentId} | Delete an external payment
[**GETExternalGatewayIdExternalPayments**](ExternalPaymentsApi.md#GETExternalGatewayIdExternalPayments) | **Get** /external_gateways/{externalGatewayId}/external_payments | Retrieve the external payments associated to the external gateway
[**GETExternalPayments**](ExternalPaymentsApi.md#GETExternalPayments) | **Get** /external_payments | List all external payments
[**GETExternalPaymentsExternalPaymentId**](ExternalPaymentsApi.md#GETExternalPaymentsExternalPaymentId) | **Get** /external_payments/{externalPaymentId} | Retrieve an external payment
[**PATCHExternalPaymentsExternalPaymentId**](ExternalPaymentsApi.md#PATCHExternalPaymentsExternalPaymentId) | **Patch** /external_payments/{externalPaymentId} | Update an external payment
[**POSTExternalPayments**](ExternalPaymentsApi.md#POSTExternalPayments) | **Post** /external_payments | Create an external payment



## DELETEExternalPaymentsExternalPaymentId

> DELETEExternalPaymentsExternalPaymentId(ctx, externalPaymentId).Execute()

Delete an external payment



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
    externalPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExternalPaymentsApi.DELETEExternalPaymentsExternalPaymentId(context.Background(), externalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.DELETEExternalPaymentsExternalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEExternalPaymentsExternalPaymentIdRequest struct via the builder pattern


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


## GETExternalGatewayIdExternalPayments

> GETExternalGatewayIdExternalPayments(ctx, externalGatewayId).Execute()

Retrieve the external payments associated to the external gateway



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
    r, err := apiClient.ExternalPaymentsApi.GETExternalGatewayIdExternalPayments(context.Background(), externalGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.GETExternalGatewayIdExternalPayments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETExternalGatewayIdExternalPaymentsRequest struct via the builder pattern


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


## GETExternalPayments

> GETExternalPayments200Response GETExternalPayments(ctx).Execute()

List all external payments



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
    resp, r, err := apiClient.ExternalPaymentsApi.GETExternalPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.GETExternalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalPayments`: GETExternalPayments200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalPaymentsApi.GETExternalPayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPaymentsRequest struct via the builder pattern


### Return type

[**GETExternalPayments200Response**](GETExternalPayments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETExternalPaymentsExternalPaymentId

> GETExternalPaymentsExternalPaymentId200Response GETExternalPaymentsExternalPaymentId(ctx, externalPaymentId).Execute()

Retrieve an external payment



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
    externalPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPaymentsApi.GETExternalPaymentsExternalPaymentId(context.Background(), externalPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.GETExternalPaymentsExternalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETExternalPaymentsExternalPaymentId`: GETExternalPaymentsExternalPaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalPaymentsApi.GETExternalPaymentsExternalPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETExternalPaymentsExternalPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETExternalPaymentsExternalPaymentId200Response**](GETExternalPaymentsExternalPaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHExternalPaymentsExternalPaymentId

> PATCHExternalPaymentsExternalPaymentId200Response PATCHExternalPaymentsExternalPaymentId(ctx, externalPaymentId).ExternalPaymentUpdate(externalPaymentUpdate).Execute()

Update an external payment



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
    externalPaymentUpdate := *openapiclient.NewExternalPaymentUpdate(*openapiclient.NewExternalPaymentUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHExternalPaymentsExternalPaymentId200ResponseDataAttributes())) // ExternalPaymentUpdate | 
    externalPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPaymentsApi.PATCHExternalPaymentsExternalPaymentId(context.Background(), externalPaymentId).ExternalPaymentUpdate(externalPaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.PATCHExternalPaymentsExternalPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHExternalPaymentsExternalPaymentId`: PATCHExternalPaymentsExternalPaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalPaymentsApi.PATCHExternalPaymentsExternalPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHExternalPaymentsExternalPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalPaymentUpdate** | [**ExternalPaymentUpdate**](ExternalPaymentUpdate.md) |  | 


### Return type

[**PATCHExternalPaymentsExternalPaymentId200Response**](PATCHExternalPaymentsExternalPaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTExternalPayments

> POSTExternalPayments201Response POSTExternalPayments(ctx).ExternalPaymentCreate(externalPaymentCreate).Execute()

Create an external payment



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
    externalPaymentCreate := *openapiclient.NewExternalPaymentCreate(*openapiclient.NewExternalPaymentCreateData(interface{}(123), *openapiclient.NewPOSTExternalPayments201ResponseDataAttributes(interface{}(xxxx.yyyy.zzzz)))) // ExternalPaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExternalPaymentsApi.POSTExternalPayments(context.Background()).ExternalPaymentCreate(externalPaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExternalPaymentsApi.POSTExternalPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTExternalPayments`: POSTExternalPayments201Response
    fmt.Fprintf(os.Stdout, "Response from `ExternalPaymentsApi.POSTExternalPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTExternalPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **externalPaymentCreate** | [**ExternalPaymentCreate**](ExternalPaymentCreate.md) |  | 

### Return type

[**POSTExternalPayments201Response**](POSTExternalPayments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

