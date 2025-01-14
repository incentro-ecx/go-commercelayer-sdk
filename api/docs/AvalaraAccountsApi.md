# \AvalaraAccountsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEAvalaraAccountsAvalaraAccountId**](AvalaraAccountsApi.md#DELETEAvalaraAccountsAvalaraAccountId) | **Delete** /avalara_accounts/{avalaraAccountId} | Delete an avalara account
[**GETAvalaraAccounts**](AvalaraAccountsApi.md#GETAvalaraAccounts) | **Get** /avalara_accounts | List all avalara accounts
[**GETAvalaraAccountsAvalaraAccountId**](AvalaraAccountsApi.md#GETAvalaraAccountsAvalaraAccountId) | **Get** /avalara_accounts/{avalaraAccountId} | Retrieve an avalara account
[**PATCHAvalaraAccountsAvalaraAccountId**](AvalaraAccountsApi.md#PATCHAvalaraAccountsAvalaraAccountId) | **Patch** /avalara_accounts/{avalaraAccountId} | Update an avalara account
[**POSTAvalaraAccounts**](AvalaraAccountsApi.md#POSTAvalaraAccounts) | **Post** /avalara_accounts | Create an avalara account



## DELETEAvalaraAccountsAvalaraAccountId

> DELETEAvalaraAccountsAvalaraAccountId(ctx, avalaraAccountId).Execute()

Delete an avalara account



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
    avalaraAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AvalaraAccountsApi.DELETEAvalaraAccountsAvalaraAccountId(context.Background(), avalaraAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.DELETEAvalaraAccountsAvalaraAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avalaraAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEAvalaraAccountsAvalaraAccountIdRequest struct via the builder pattern


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


## GETAvalaraAccounts

> GETAvalaraAccounts200Response GETAvalaraAccounts(ctx).Execute()

List all avalara accounts



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
    resp, r, err := apiClient.AvalaraAccountsApi.GETAvalaraAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.GETAvalaraAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAvalaraAccounts`: GETAvalaraAccounts200Response
    fmt.Fprintf(os.Stdout, "Response from `AvalaraAccountsApi.GETAvalaraAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETAvalaraAccountsRequest struct via the builder pattern


### Return type

[**GETAvalaraAccounts200Response**](GETAvalaraAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETAvalaraAccountsAvalaraAccountId

> GETAvalaraAccountsAvalaraAccountId200Response GETAvalaraAccountsAvalaraAccountId(ctx, avalaraAccountId).Execute()

Retrieve an avalara account



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
    avalaraAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvalaraAccountsApi.GETAvalaraAccountsAvalaraAccountId(context.Background(), avalaraAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.GETAvalaraAccountsAvalaraAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETAvalaraAccountsAvalaraAccountId`: GETAvalaraAccountsAvalaraAccountId200Response
    fmt.Fprintf(os.Stdout, "Response from `AvalaraAccountsApi.GETAvalaraAccountsAvalaraAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avalaraAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETAvalaraAccountsAvalaraAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETAvalaraAccountsAvalaraAccountId200Response**](GETAvalaraAccountsAvalaraAccountId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHAvalaraAccountsAvalaraAccountId

> PATCHAvalaraAccountsAvalaraAccountId200Response PATCHAvalaraAccountsAvalaraAccountId(ctx, avalaraAccountId).AvalaraAccountUpdate(avalaraAccountUpdate).Execute()

Update an avalara account



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
    avalaraAccountUpdate := *openapiclient.NewAvalaraAccountUpdate(*openapiclient.NewAvalaraAccountUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHAvalaraAccountsAvalaraAccountId200ResponseDataAttributes())) // AvalaraAccountUpdate | 
    avalaraAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvalaraAccountsApi.PATCHAvalaraAccountsAvalaraAccountId(context.Background(), avalaraAccountId).AvalaraAccountUpdate(avalaraAccountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.PATCHAvalaraAccountsAvalaraAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHAvalaraAccountsAvalaraAccountId`: PATCHAvalaraAccountsAvalaraAccountId200Response
    fmt.Fprintf(os.Stdout, "Response from `AvalaraAccountsApi.PATCHAvalaraAccountsAvalaraAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**avalaraAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHAvalaraAccountsAvalaraAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avalaraAccountUpdate** | [**AvalaraAccountUpdate**](AvalaraAccountUpdate.md) |  | 


### Return type

[**PATCHAvalaraAccountsAvalaraAccountId200Response**](PATCHAvalaraAccountsAvalaraAccountId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTAvalaraAccounts

> POSTAvalaraAccounts201Response POSTAvalaraAccounts(ctx).AvalaraAccountCreate(avalaraAccountCreate).Execute()

Create an avalara account



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
    avalaraAccountCreate := *openapiclient.NewAvalaraAccountCreate(*openapiclient.NewAvalaraAccountCreateData(interface{}(123), *openapiclient.NewPOSTAvalaraAccounts201ResponseDataAttributes(interface{}(Personal tax calculator), interface{}(user@mydomain.com), interface{}(secret), interface{}(MYCOMPANY)))) // AvalaraAccountCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AvalaraAccountsApi.POSTAvalaraAccounts(context.Background()).AvalaraAccountCreate(avalaraAccountCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AvalaraAccountsApi.POSTAvalaraAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTAvalaraAccounts`: POSTAvalaraAccounts201Response
    fmt.Fprintf(os.Stdout, "Response from `AvalaraAccountsApi.POSTAvalaraAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTAvalaraAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avalaraAccountCreate** | [**AvalaraAccountCreate**](AvalaraAccountCreate.md) |  | 

### Return type

[**POSTAvalaraAccounts201Response**](POSTAvalaraAccounts201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

