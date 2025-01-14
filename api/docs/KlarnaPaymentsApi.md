# \KlarnaPaymentsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEKlarnaPaymentsKlarnaPaymentId**](KlarnaPaymentsApi.md#DELETEKlarnaPaymentsKlarnaPaymentId) | **Delete** /klarna_payments/{klarnaPaymentId} | Delete a klarna payment
[**GETKlarnaGatewayIdKlarnaPayments**](KlarnaPaymentsApi.md#GETKlarnaGatewayIdKlarnaPayments) | **Get** /klarna_gateways/{klarnaGatewayId}/klarna_payments | Retrieve the klarna payments associated to the klarna gateway
[**GETKlarnaPayments**](KlarnaPaymentsApi.md#GETKlarnaPayments) | **Get** /klarna_payments | List all klarna payments
[**GETKlarnaPaymentsKlarnaPaymentId**](KlarnaPaymentsApi.md#GETKlarnaPaymentsKlarnaPaymentId) | **Get** /klarna_payments/{klarnaPaymentId} | Retrieve a klarna payment
[**PATCHKlarnaPaymentsKlarnaPaymentId**](KlarnaPaymentsApi.md#PATCHKlarnaPaymentsKlarnaPaymentId) | **Patch** /klarna_payments/{klarnaPaymentId} | Update a klarna payment
[**POSTKlarnaPayments**](KlarnaPaymentsApi.md#POSTKlarnaPayments) | **Post** /klarna_payments | Create a klarna payment



## DELETEKlarnaPaymentsKlarnaPaymentId

> DELETEKlarnaPaymentsKlarnaPaymentId(ctx, klarnaPaymentId).Execute()

Delete a klarna payment



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
    klarnaPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KlarnaPaymentsApi.DELETEKlarnaPaymentsKlarnaPaymentId(context.Background(), klarnaPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.DELETEKlarnaPaymentsKlarnaPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEKlarnaPaymentsKlarnaPaymentIdRequest struct via the builder pattern


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


## GETKlarnaGatewayIdKlarnaPayments

> GETKlarnaGatewayIdKlarnaPayments(ctx, klarnaGatewayId).Execute()

Retrieve the klarna payments associated to the klarna gateway



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
    klarnaGatewayId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KlarnaPaymentsApi.GETKlarnaGatewayIdKlarnaPayments(context.Background(), klarnaGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.GETKlarnaGatewayIdKlarnaPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaGatewayId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaGatewayIdKlarnaPaymentsRequest struct via the builder pattern


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


## GETKlarnaPayments

> GETKlarnaPayments200Response GETKlarnaPayments(ctx).Execute()

List all klarna payments



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
    resp, r, err := apiClient.KlarnaPaymentsApi.GETKlarnaPayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.GETKlarnaPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETKlarnaPayments`: GETKlarnaPayments200Response
    fmt.Fprintf(os.Stdout, "Response from `KlarnaPaymentsApi.GETKlarnaPayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaPaymentsRequest struct via the builder pattern


### Return type

[**GETKlarnaPayments200Response**](GETKlarnaPayments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETKlarnaPaymentsKlarnaPaymentId

> GETKlarnaPaymentsKlarnaPaymentId200Response GETKlarnaPaymentsKlarnaPaymentId(ctx, klarnaPaymentId).Execute()

Retrieve a klarna payment



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
    klarnaPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaPaymentsApi.GETKlarnaPaymentsKlarnaPaymentId(context.Background(), klarnaPaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.GETKlarnaPaymentsKlarnaPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETKlarnaPaymentsKlarnaPaymentId`: GETKlarnaPaymentsKlarnaPaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `KlarnaPaymentsApi.GETKlarnaPaymentsKlarnaPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETKlarnaPaymentsKlarnaPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETKlarnaPaymentsKlarnaPaymentId200Response**](GETKlarnaPaymentsKlarnaPaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHKlarnaPaymentsKlarnaPaymentId

> PATCHKlarnaPaymentsKlarnaPaymentId200Response PATCHKlarnaPaymentsKlarnaPaymentId(ctx, klarnaPaymentId).KlarnaPaymentUpdate(klarnaPaymentUpdate).Execute()

Update a klarna payment



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
    klarnaPaymentUpdate := *openapiclient.NewKlarnaPaymentUpdate(*openapiclient.NewKlarnaPaymentUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHKlarnaPaymentsKlarnaPaymentId200ResponseDataAttributes())) // KlarnaPaymentUpdate | 
    klarnaPaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaPaymentsApi.PATCHKlarnaPaymentsKlarnaPaymentId(context.Background(), klarnaPaymentId).KlarnaPaymentUpdate(klarnaPaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.PATCHKlarnaPaymentsKlarnaPaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHKlarnaPaymentsKlarnaPaymentId`: PATCHKlarnaPaymentsKlarnaPaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `KlarnaPaymentsApi.PATCHKlarnaPaymentsKlarnaPaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**klarnaPaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHKlarnaPaymentsKlarnaPaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **klarnaPaymentUpdate** | [**KlarnaPaymentUpdate**](KlarnaPaymentUpdate.md) |  | 


### Return type

[**PATCHKlarnaPaymentsKlarnaPaymentId200Response**](PATCHKlarnaPaymentsKlarnaPaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTKlarnaPayments

> POSTKlarnaPayments201Response POSTKlarnaPayments(ctx).KlarnaPaymentCreate(klarnaPaymentCreate).Execute()

Create a klarna payment



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
    klarnaPaymentCreate := *openapiclient.NewKlarnaPaymentCreate(*openapiclient.NewKlarnaPaymentCreateData(interface{}(123), *openapiclient.NewPOSTAdyenPayments201ResponseDataAttributes())) // KlarnaPaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KlarnaPaymentsApi.POSTKlarnaPayments(context.Background()).KlarnaPaymentCreate(klarnaPaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KlarnaPaymentsApi.POSTKlarnaPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTKlarnaPayments`: POSTKlarnaPayments201Response
    fmt.Fprintf(os.Stdout, "Response from `KlarnaPaymentsApi.POSTKlarnaPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTKlarnaPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **klarnaPaymentCreate** | [**KlarnaPaymentCreate**](KlarnaPaymentCreate.md) |  | 

### Return type

[**POSTKlarnaPayments201Response**](POSTKlarnaPayments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

