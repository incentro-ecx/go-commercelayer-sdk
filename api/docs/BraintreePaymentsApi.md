# \BraintreePaymentsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEBraintreePaymentsBraintreePaymentId**](BraintreePaymentsApi.md#DELETEBraintreePaymentsBraintreePaymentId) | **Delete** /braintree_payments/{braintreePaymentId} | Delete a braintree payment
[**GETBraintreeGatewayIdBraintreePayments**](BraintreePaymentsApi.md#GETBraintreeGatewayIdBraintreePayments) | **Get** /braintree_gateways/{braintreeGatewayId}/braintree_payments | Retrieve the braintree payments associated to the braintree gateway
[**GETBraintreePayments**](BraintreePaymentsApi.md#GETBraintreePayments) | **Get** /braintree_payments | List all braintree payments
[**GETBraintreePaymentsBraintreePaymentId**](BraintreePaymentsApi.md#GETBraintreePaymentsBraintreePaymentId) | **Get** /braintree_payments/{braintreePaymentId} | Retrieve a braintree payment
[**PATCHBraintreePaymentsBraintreePaymentId**](BraintreePaymentsApi.md#PATCHBraintreePaymentsBraintreePaymentId) | **Patch** /braintree_payments/{braintreePaymentId} | Update a braintree payment
[**POSTBraintreePayments**](BraintreePaymentsApi.md#POSTBraintreePayments) | **Post** /braintree_payments | Create a braintree payment



## DELETEBraintreePaymentsBraintreePaymentId

> DELETEBraintreePaymentsBraintreePaymentId(ctx, braintreePaymentId).Execute()

Delete a braintree payment



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
    braintreePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BraintreePaymentsApi.DELETEBraintreePaymentsBraintreePaymentId(context.Background(), braintreePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.DELETEBraintreePaymentsBraintreePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEBraintreePaymentsBraintreePaymentIdRequest struct via the builder pattern


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


## GETBraintreeGatewayIdBraintreePayments

> GETBraintreeGatewayIdBraintreePayments(ctx, braintreeGatewayId).Execute()

Retrieve the braintree payments associated to the braintree gateway



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
    r, err := apiClient.BraintreePaymentsApi.GETBraintreeGatewayIdBraintreePayments(context.Background(), braintreeGatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.GETBraintreeGatewayIdBraintreePayments``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETBraintreeGatewayIdBraintreePaymentsRequest struct via the builder pattern


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


## GETBraintreePayments

> GETBraintreePayments200Response GETBraintreePayments(ctx).Execute()

List all braintree payments



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
    resp, r, err := apiClient.BraintreePaymentsApi.GETBraintreePayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.GETBraintreePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBraintreePayments`: GETBraintreePayments200Response
    fmt.Fprintf(os.Stdout, "Response from `BraintreePaymentsApi.GETBraintreePayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreePaymentsRequest struct via the builder pattern


### Return type

[**GETBraintreePayments200Response**](GETBraintreePayments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETBraintreePaymentsBraintreePaymentId

> GETBraintreePaymentsBraintreePaymentId200Response GETBraintreePaymentsBraintreePaymentId(ctx, braintreePaymentId).Execute()

Retrieve a braintree payment



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
    braintreePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreePaymentsApi.GETBraintreePaymentsBraintreePaymentId(context.Background(), braintreePaymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.GETBraintreePaymentsBraintreePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETBraintreePaymentsBraintreePaymentId`: GETBraintreePaymentsBraintreePaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `BraintreePaymentsApi.GETBraintreePaymentsBraintreePaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETBraintreePaymentsBraintreePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETBraintreePaymentsBraintreePaymentId200Response**](GETBraintreePaymentsBraintreePaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHBraintreePaymentsBraintreePaymentId

> PATCHBraintreePaymentsBraintreePaymentId200Response PATCHBraintreePaymentsBraintreePaymentId(ctx, braintreePaymentId).BraintreePaymentUpdate(braintreePaymentUpdate).Execute()

Update a braintree payment



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
    braintreePaymentUpdate := *openapiclient.NewBraintreePaymentUpdate(*openapiclient.NewBraintreePaymentUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHBraintreePaymentsBraintreePaymentId200ResponseDataAttributes())) // BraintreePaymentUpdate | 
    braintreePaymentId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreePaymentsApi.PATCHBraintreePaymentsBraintreePaymentId(context.Background(), braintreePaymentId).BraintreePaymentUpdate(braintreePaymentUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.PATCHBraintreePaymentsBraintreePaymentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHBraintreePaymentsBraintreePaymentId`: PATCHBraintreePaymentsBraintreePaymentId200Response
    fmt.Fprintf(os.Stdout, "Response from `BraintreePaymentsApi.PATCHBraintreePaymentsBraintreePaymentId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**braintreePaymentId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHBraintreePaymentsBraintreePaymentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **braintreePaymentUpdate** | [**BraintreePaymentUpdate**](BraintreePaymentUpdate.md) |  | 


### Return type

[**PATCHBraintreePaymentsBraintreePaymentId200Response**](PATCHBraintreePaymentsBraintreePaymentId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTBraintreePayments

> POSTBraintreePayments201Response POSTBraintreePayments(ctx).BraintreePaymentCreate(braintreePaymentCreate).Execute()

Create a braintree payment



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
    braintreePaymentCreate := *openapiclient.NewBraintreePaymentCreate(*openapiclient.NewBraintreePaymentCreateData(interface{}(123), *openapiclient.NewPOSTBraintreePayments201ResponseDataAttributes())) // BraintreePaymentCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BraintreePaymentsApi.POSTBraintreePayments(context.Background()).BraintreePaymentCreate(braintreePaymentCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BraintreePaymentsApi.POSTBraintreePayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTBraintreePayments`: POSTBraintreePayments201Response
    fmt.Fprintf(os.Stdout, "Response from `BraintreePaymentsApi.POSTBraintreePayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTBraintreePaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **braintreePaymentCreate** | [**BraintreePaymentCreate**](BraintreePaymentCreate.md) |  | 

### Return type

[**POSTBraintreePayments201Response**](POSTBraintreePayments201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

