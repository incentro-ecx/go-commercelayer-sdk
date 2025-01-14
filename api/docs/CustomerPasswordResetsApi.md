# \CustomerPasswordResetsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomerPasswordResetsCustomerPasswordResetId**](CustomerPasswordResetsApi.md#DELETECustomerPasswordResetsCustomerPasswordResetId) | **Delete** /customer_password_resets/{customerPasswordResetId} | Delete a customer password reset
[**GETCustomerPasswordResets**](CustomerPasswordResetsApi.md#GETCustomerPasswordResets) | **Get** /customer_password_resets | List all customer password resets
[**GETCustomerPasswordResetsCustomerPasswordResetId**](CustomerPasswordResetsApi.md#GETCustomerPasswordResetsCustomerPasswordResetId) | **Get** /customer_password_resets/{customerPasswordResetId} | Retrieve a customer password reset
[**PATCHCustomerPasswordResetsCustomerPasswordResetId**](CustomerPasswordResetsApi.md#PATCHCustomerPasswordResetsCustomerPasswordResetId) | **Patch** /customer_password_resets/{customerPasswordResetId} | Update a customer password reset
[**POSTCustomerPasswordResets**](CustomerPasswordResetsApi.md#POSTCustomerPasswordResets) | **Post** /customer_password_resets | Create a customer password reset



## DELETECustomerPasswordResetsCustomerPasswordResetId

> DELETECustomerPasswordResetsCustomerPasswordResetId(ctx, customerPasswordResetId).Execute()

Delete a customer password reset



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
    customerPasswordResetId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomerPasswordResetsApi.DELETECustomerPasswordResetsCustomerPasswordResetId(context.Background(), customerPasswordResetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.DELETECustomerPasswordResetsCustomerPasswordResetId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPasswordResetId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECustomerPasswordResetsCustomerPasswordResetIdRequest struct via the builder pattern


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


## GETCustomerPasswordResets

> GETCustomerPasswordResets200Response GETCustomerPasswordResets(ctx).Execute()

List all customer password resets



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
    resp, r, err := apiClient.CustomerPasswordResetsApi.GETCustomerPasswordResets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.GETCustomerPasswordResets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerPasswordResets`: GETCustomerPasswordResets200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerPasswordResetsApi.GETCustomerPasswordResets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerPasswordResetsRequest struct via the builder pattern


### Return type

[**GETCustomerPasswordResets200Response**](GETCustomerPasswordResets200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCustomerPasswordResetsCustomerPasswordResetId

> GETCustomerPasswordResetsCustomerPasswordResetId200Response GETCustomerPasswordResetsCustomerPasswordResetId(ctx, customerPasswordResetId).Execute()

Retrieve a customer password reset



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
    customerPasswordResetId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPasswordResetsApi.GETCustomerPasswordResetsCustomerPasswordResetId(context.Background(), customerPasswordResetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.GETCustomerPasswordResetsCustomerPasswordResetId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerPasswordResetsCustomerPasswordResetId`: GETCustomerPasswordResetsCustomerPasswordResetId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerPasswordResetsApi.GETCustomerPasswordResetsCustomerPasswordResetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPasswordResetId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerPasswordResetsCustomerPasswordResetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCustomerPasswordResetsCustomerPasswordResetId200Response**](GETCustomerPasswordResetsCustomerPasswordResetId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHCustomerPasswordResetsCustomerPasswordResetId

> PATCHCustomerPasswordResetsCustomerPasswordResetId200Response PATCHCustomerPasswordResetsCustomerPasswordResetId(ctx, customerPasswordResetId).CustomerPasswordResetUpdate(customerPasswordResetUpdate).Execute()

Update a customer password reset



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
    customerPasswordResetUpdate := *openapiclient.NewCustomerPasswordResetUpdate(*openapiclient.NewCustomerPasswordResetUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCustomerPasswordResetsCustomerPasswordResetId200ResponseDataAttributes())) // CustomerPasswordResetUpdate | 
    customerPasswordResetId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPasswordResetsApi.PATCHCustomerPasswordResetsCustomerPasswordResetId(context.Background(), customerPasswordResetId).CustomerPasswordResetUpdate(customerPasswordResetUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.PATCHCustomerPasswordResetsCustomerPasswordResetId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCustomerPasswordResetsCustomerPasswordResetId`: PATCHCustomerPasswordResetsCustomerPasswordResetId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerPasswordResetsApi.PATCHCustomerPasswordResetsCustomerPasswordResetId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerPasswordResetId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomerPasswordResetsCustomerPasswordResetIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerPasswordResetUpdate** | [**CustomerPasswordResetUpdate**](CustomerPasswordResetUpdate.md) |  | 


### Return type

[**PATCHCustomerPasswordResetsCustomerPasswordResetId200Response**](PATCHCustomerPasswordResetsCustomerPasswordResetId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCustomerPasswordResets

> POSTCustomerPasswordResets201Response POSTCustomerPasswordResets(ctx).CustomerPasswordResetCreate(customerPasswordResetCreate).Execute()

Create a customer password reset



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
    customerPasswordResetCreate := *openapiclient.NewCustomerPasswordResetCreate(*openapiclient.NewCustomerPasswordResetCreateData(interface{}(123), *openapiclient.NewPOSTCustomerPasswordResets201ResponseDataAttributes(interface{}(john@example.com)))) // CustomerPasswordResetCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerPasswordResetsApi.POSTCustomerPasswordResets(context.Background()).CustomerPasswordResetCreate(customerPasswordResetCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerPasswordResetsApi.POSTCustomerPasswordResets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCustomerPasswordResets`: POSTCustomerPasswordResets201Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerPasswordResetsApi.POSTCustomerPasswordResets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomerPasswordResetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerPasswordResetCreate** | [**CustomerPasswordResetCreate**](CustomerPasswordResetCreate.md) |  | 

### Return type

[**POSTCustomerPasswordResets201Response**](POSTCustomerPasswordResets201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

