# \AxervePaymentsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAxervePaymentsAxervePaymentId**](AxervePaymentsApi.md#DELETEAxervePaymentsAxervePaymentId) | **Delete** /axerve_payments/{axervePaymentId} | Delete an axerve payment
[**GETAxerveGatewayIdAxervePayments**](AxervePaymentsApi.md#GETAxerveGatewayIdAxervePayments) | **Get** /axerve_gateways/{axerveGatewayId}/axerve_payments | Retrieve the axerve payments associated to the axerve gateway
[**GETAxervePayments**](AxervePaymentsApi.md#GETAxervePayments) | **Get** /axerve_payments | List all axerve payments
[**GETAxervePaymentsAxervePaymentId**](AxervePaymentsApi.md#GETAxervePaymentsAxervePaymentId) | **Get** /axerve_payments/{axervePaymentId} | Retrieve an axerve payment
[**PATCHAxervePaymentsAxervePaymentId**](AxervePaymentsApi.md#PATCHAxervePaymentsAxervePaymentId) | **Patch** /axerve_payments/{axervePaymentId} | Update an axerve payment
[**POSTAxervePayments**](AxervePaymentsApi.md#POSTAxervePayments) | **Post** /axerve_payments | Create an axerve payment



## DELETEAxervePaymentsAxervePaymentId

> DELETEAxervePaymentsAxervePaymentId(ctx, axervePaymentId).Execute()

Delete an axerve payment



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
    axervePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AxervePaymentsApi.DELETEAxervePaymentsAxervePaymentId(context.Background(), axervePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxervePaymentsApi.DELETEAxervePaymentsAxervePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**axervePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEAxervePaymentsAxervePaymentIdRequest struct via the builder pattern


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


## GETAxerveGatewayIdAxervePayments

> GETAxerveGatewayIdAxervePayments(ctx, axerveGatewayId).Execute()

Retrieve the axerve payments associated to the axerve gateway



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
    r, err := apiClient.AxervePaymentsApi.GETAxerveGatewayIdAxervePayments(context.Background(), axerveGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxervePaymentsApi.GETAxerveGatewayIdAxervePayments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETAxerveGatewayIdAxervePaymentsRequest struct via the builder pattern


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


## GETAxervePayments

> GETAxervePayments200Response GETAxervePayments(ctx).Execute()

List all axerve payments



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
    resp, r, err := apiClient.AxervePaymentsApi.GETAxervePayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxervePaymentsApi.GETAxervePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAxervePayments`: GETAxervePayments200Response
    fmt.Fprintf(os.Stdout, "Response from `AxervePaymentsApi.GETAxervePayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETAxervePaymentsRequest struct via the builder pattern


### Return type

[**GETAxervePayments200Response**](GETAxervePayments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETAxervePaymentsAxervePaymentId

> GETAxervePaymentsAxervePaymentId200Response GETAxervePaymentsAxervePaymentId(ctx, axervePaymentId).Execute()

Retrieve an axerve payment



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
    axervePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AxervePaymentsApi.GETAxervePaymentsAxervePaymentId(context.Background(), axervePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxervePaymentsApi.GETAxervePaymentsAxervePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAxervePaymentsAxervePaymentId`: GETAxervePaymentsAxervePaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `AxervePaymentsApi.GETAxervePaymentsAxervePaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**axervePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAxervePaymentsAxervePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETAxervePaymentsAxervePaymentId200Response**](GETAxervePaymentsAxervePaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHAxervePaymentsAxervePaymentId

> PATCHAxervePaymentsAxervePaymentId200Response PATCHAxervePaymentsAxervePaymentId(ctx, axervePaymentId).AxervePaymentUpdate(axervePaymentUpdate).Execute()

Update an axerve payment



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
    axervePaymentUpdate := *openapiclient.NewAxervePaymentUpdate(*openapiclient.NewAxervePaymentUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHAxervePaymentsAxervePaymentId200ResponseDataAttributes())) // AxervePaymentUpdate | 
    axervePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AxervePaymentsApi.PATCHAxervePaymentsAxervePaymentId(context.Background(), axervePaymentId).AxervePaymentUpdate(axervePaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxervePaymentsApi.PATCHAxervePaymentsAxervePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHAxervePaymentsAxervePaymentId`: PATCHAxervePaymentsAxervePaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `AxervePaymentsApi.PATCHAxervePaymentsAxervePaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**axervePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHAxervePaymentsAxervePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **axervePaymentUpdate** | [**AxervePaymentUpdate**](AxervePaymentUpdate.md) |  | 


### Return type

[**PATCHAxervePaymentsAxervePaymentId200Response**](PATCHAxervePaymentsAxervePaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTAxervePayments

> POSTAxervePayments201Response POSTAxervePayments(ctx).AxervePaymentCreate(axervePaymentCreate).Execute()

Create an axerve payment



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
    axervePaymentCreate := *openapiclient.NewAxervePaymentCreate(*openapiclient.NewAxervePaymentCreateData(interface{}(123), *openapiclient.NewPOSTAxervePayments201ResponseDataAttributes(interface{}(https://yourdomain.com/thankyou)))) // AxervePaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AxervePaymentsApi.POSTAxervePayments(context.Background()).AxervePaymentCreate(axervePaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AxervePaymentsApi.POSTAxervePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTAxervePayments`: POSTAxervePayments201Response
    fmt.Fprintf(os.Stdout, "Response from `AxervePaymentsApi.POSTAxervePayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTAxervePaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **axervePaymentCreate** | [**AxervePaymentCreate**](AxervePaymentCreate.md) |  | 

### Return type

[**POSTAxervePayments201Response**](POSTAxervePayments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

