# \TaxCategoriesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETETaxCategoriesTaxCategoryId**](TaxCategoriesApi.md#DELETETaxCategoriesTaxCategoryId) | **Delete** /tax_categories/{taxCategoryId} | Delete a tax category
[**GETAvalaraAccountIdTaxCategories**](TaxCategoriesApi.md#GETAvalaraAccountIdTaxCategories) | **Get** /avalara_accounts/{avalaraAccountId}/tax_categories | Retrieve the tax categories associated to the avalara account
[**GETTaxCategories**](TaxCategoriesApi.md#GETTaxCategories) | **Get** /tax_categories | List all tax categories
[**GETTaxCategoriesTaxCategoryId**](TaxCategoriesApi.md#GETTaxCategoriesTaxCategoryId) | **Get** /tax_categories/{taxCategoryId} | Retrieve a tax category
[**GETTaxjarAccountIdTaxCategories**](TaxCategoriesApi.md#GETTaxjarAccountIdTaxCategories) | **Get** /taxjar_accounts/{taxjarAccountId}/tax_categories | Retrieve the tax categories associated to the taxjar account
[**PATCHTaxCategoriesTaxCategoryId**](TaxCategoriesApi.md#PATCHTaxCategoriesTaxCategoryId) | **Patch** /tax_categories/{taxCategoryId} | Update a tax category
[**POSTTaxCategories**](TaxCategoriesApi.md#POSTTaxCategories) | **Post** /tax_categories | Create a tax category



## DELETETaxCategoriesTaxCategoryId

> DELETETaxCategoriesTaxCategoryId(ctx, taxCategoryId).Execute()

Delete a tax category



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
    taxCategoryId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaxCategoriesApi.DELETETaxCategoriesTaxCategoryId(context.Background(), taxCategoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.DELETETaxCategoriesTaxCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCategoryId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETETaxCategoriesTaxCategoryIdRequest struct via the builder pattern


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


## GETAvalaraAccountIdTaxCategories

> GETAvalaraAccountIdTaxCategories(ctx, avalaraAccountId).Execute()

Retrieve the tax categories associated to the avalara account



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
    r, err := apiClient.TaxCategoriesApi.GETAvalaraAccountIdTaxCategories(context.Background(), avalaraAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETAvalaraAccountIdTaxCategories``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETAvalaraAccountIdTaxCategoriesRequest struct via the builder pattern


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


## GETTaxCategories

> GETTaxCategories200Response GETTaxCategories(ctx).Execute()

List all tax categories



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
    resp, r, err := apiClient.TaxCategoriesApi.GETTaxCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETTaxCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxCategories`: GETTaxCategories200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxCategoriesApi.GETTaxCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxCategoriesRequest struct via the builder pattern


### Return type

[**GETTaxCategories200Response**](GETTaxCategories200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTaxCategoriesTaxCategoryId

> GETTaxCategoriesTaxCategoryId200Response GETTaxCategoriesTaxCategoryId(ctx, taxCategoryId).Execute()

Retrieve a tax category



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
    taxCategoryId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.GETTaxCategoriesTaxCategoryId(context.Background(), taxCategoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETTaxCategoriesTaxCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETTaxCategoriesTaxCategoryId`: GETTaxCategoriesTaxCategoryId200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxCategoriesApi.GETTaxCategoriesTaxCategoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCategoryId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETTaxCategoriesTaxCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETTaxCategoriesTaxCategoryId200Response**](GETTaxCategoriesTaxCategoryId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETTaxjarAccountIdTaxCategories

> GETTaxjarAccountIdTaxCategories(ctx, taxjarAccountId).Execute()

Retrieve the tax categories associated to the taxjar account



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
    r, err := apiClient.TaxCategoriesApi.GETTaxjarAccountIdTaxCategories(context.Background(), taxjarAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.GETTaxjarAccountIdTaxCategories``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGETTaxjarAccountIdTaxCategoriesRequest struct via the builder pattern


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


## PATCHTaxCategoriesTaxCategoryId

> PATCHTaxCategoriesTaxCategoryId200Response PATCHTaxCategoriesTaxCategoryId(ctx, taxCategoryId).TaxCategoryUpdate(taxCategoryUpdate).Execute()

Update a tax category



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
    taxCategoryUpdate := *openapiclient.NewTaxCategoryUpdate(*openapiclient.NewTaxCategoryUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHTaxCategoriesTaxCategoryId200ResponseDataAttributes())) // TaxCategoryUpdate | 
    taxCategoryId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.PATCHTaxCategoriesTaxCategoryId(context.Background(), taxCategoryId).TaxCategoryUpdate(taxCategoryUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.PATCHTaxCategoriesTaxCategoryId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHTaxCategoriesTaxCategoryId`: PATCHTaxCategoriesTaxCategoryId200Response
    fmt.Fprintf(os.Stdout, "Response from `TaxCategoriesApi.PATCHTaxCategoriesTaxCategoryId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxCategoryId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHTaxCategoriesTaxCategoryIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxCategoryUpdate** | [**TaxCategoryUpdate**](TaxCategoryUpdate.md) |  | 


### Return type

[**PATCHTaxCategoriesTaxCategoryId200Response**](PATCHTaxCategoriesTaxCategoryId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTTaxCategories

> POSTTaxCategories201Response POSTTaxCategories(ctx).TaxCategoryCreate(taxCategoryCreate).Execute()

Create a tax category



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
    taxCategoryCreate := *openapiclient.NewTaxCategoryCreate(*openapiclient.NewTaxCategoryCreateData(interface{}(123), *openapiclient.NewPOSTTaxCategories201ResponseDataAttributes(interface{}(31000)))) // TaxCategoryCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxCategoriesApi.POSTTaxCategories(context.Background()).TaxCategoryCreate(taxCategoryCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxCategoriesApi.POSTTaxCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTTaxCategories`: POSTTaxCategories201Response
    fmt.Fprintf(os.Stdout, "Response from `TaxCategoriesApi.POSTTaxCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTTaxCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxCategoryCreate** | [**TaxCategoryCreate**](TaxCategoryCreate.md) |  | 

### Return type

[**POSTTaxCategories201Response**](POSTTaxCategories201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

