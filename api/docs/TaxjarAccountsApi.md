# \TaxjarAccountsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETETaxjarAccountsTaxjarAccountId**](TaxjarAccountsApi.md#DELETETaxjarAccountsTaxjarAccountId) | **Delete** /taxjar_accounts/{taxjarAccountId} | Delete a taxjar account
[**GETTaxjarAccounts**](TaxjarAccountsApi.md#GETTaxjarAccounts) | **Get** /taxjar_accounts | List all taxjar accounts
[**GETTaxjarAccountsTaxjarAccountId**](TaxjarAccountsApi.md#GETTaxjarAccountsTaxjarAccountId) | **Get** /taxjar_accounts/{taxjarAccountId} | Retrieve a taxjar account
[**PATCHTaxjarAccountsTaxjarAccountId**](TaxjarAccountsApi.md#PATCHTaxjarAccountsTaxjarAccountId) | **Patch** /taxjar_accounts/{taxjarAccountId} | Update a taxjar account
[**POSTTaxjarAccounts**](TaxjarAccountsApi.md#POSTTaxjarAccounts) | **Post** /taxjar_accounts | Create a taxjar account



## DELETETaxjarAccountsTaxjarAccountId

> DELETETaxjarAccountsTaxjarAccountId(ctx, taxjarAccountId).Execute()

Delete a taxjar account



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
    taxjarAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaxjarAccountsApi.DELETETaxjarAccountsTaxjarAccountId(context.Background(), taxjarAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxjarAccountsApi.DELETETaxjarAccountsTaxjarAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxjarAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETETaxjarAccountsTaxjarAccountIdRequest struct via the builder pattern


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


## GETTaxjarAccounts

> GETTaxjarAccounts200Response GETTaxjarAccounts(ctx).Execute()

List all taxjar accounts



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
    resp, r, err := apiClient.TaxjarAccountsApi.GETTaxjarAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxjarAccountsApi.GETTaxjarAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxjarAccounts`: GETTaxjarAccounts200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxjarAccountsApi.GETTaxjarAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxjarAccountsRequest struct via the builder pattern


### Return type

[**GETTaxjarAccounts200Response**](GETTaxjarAccounts200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTaxjarAccountsTaxjarAccountId

> GETTaxjarAccountsTaxjarAccountId200Response GETTaxjarAccountsTaxjarAccountId(ctx, taxjarAccountId).Execute()

Retrieve a taxjar account



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
    taxjarAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxjarAccountsApi.GETTaxjarAccountsTaxjarAccountId(context.Background(), taxjarAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxjarAccountsApi.GETTaxjarAccountsTaxjarAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxjarAccountsTaxjarAccountId`: GETTaxjarAccountsTaxjarAccountId200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxjarAccountsApi.GETTaxjarAccountsTaxjarAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxjarAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxjarAccountsTaxjarAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETTaxjarAccountsTaxjarAccountId200Response**](GETTaxjarAccountsTaxjarAccountId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHTaxjarAccountsTaxjarAccountId

> PATCHTaxjarAccountsTaxjarAccountId200Response PATCHTaxjarAccountsTaxjarAccountId(ctx, taxjarAccountId).TaxjarAccountUpdate(taxjarAccountUpdate).Execute()

Update a taxjar account



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
    taxjarAccountUpdate := *openapiclient.NewTaxjarAccountUpdate(*openapiclient.NewTaxjarAccountUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHTaxjarAccountsTaxjarAccountId200ResponseDataAttributes())) // TaxjarAccountUpdate | 
    taxjarAccountId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxjarAccountsApi.PATCHTaxjarAccountsTaxjarAccountId(context.Background(), taxjarAccountId).TaxjarAccountUpdate(taxjarAccountUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxjarAccountsApi.PATCHTaxjarAccountsTaxjarAccountId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHTaxjarAccountsTaxjarAccountId`: PATCHTaxjarAccountsTaxjarAccountId200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxjarAccountsApi.PATCHTaxjarAccountsTaxjarAccountId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxjarAccountId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHTaxjarAccountsTaxjarAccountIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxjarAccountUpdate** | [**TaxjarAccountUpdate**](TaxjarAccountUpdate.md) |  | 


### Return type

[**PATCHTaxjarAccountsTaxjarAccountId200Response**](PATCHTaxjarAccountsTaxjarAccountId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTTaxjarAccounts

> POSTTaxjarAccounts201Response POSTTaxjarAccounts(ctx).TaxjarAccountCreate(taxjarAccountCreate).Execute()

Create a taxjar account



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
    taxjarAccountCreate := *openapiclient.NewTaxjarAccountCreate(*openapiclient.NewTaxjarAccountCreateData(interface{}(123), *openapiclient.NewPOSTTaxjarAccounts201ResponseDataAttributes(interface{}(Personal tax calculator), interface{}(TAXJAR_API_KEY)))) // TaxjarAccountCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxjarAccountsApi.POSTTaxjarAccounts(context.Background()).TaxjarAccountCreate(taxjarAccountCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxjarAccountsApi.POSTTaxjarAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTTaxjarAccounts`: POSTTaxjarAccounts201Response
    fmt.Fprintf(os.Stdout, "Response from `TaxjarAccountsApi.POSTTaxjarAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTaxjarAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxjarAccountCreate** | [**TaxjarAccountCreate**](TaxjarAccountCreate.md) |  | 

### Return type

[**POSTTaxjarAccounts201Response**](POSTTaxjarAccounts201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

