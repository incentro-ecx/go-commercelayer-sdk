# \ReturnLineItemsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEReturnLineItemsReturnLineItemId**](ReturnLineItemsApi.md#DELETEReturnLineItemsReturnLineItemId) | **Delete** /return_line_items/{returnLineItemId} | Delete a return line item
[**GETLineItemIdReturnLineItems**](ReturnLineItemsApi.md#GETLineItemIdReturnLineItems) | **Get** /line_items/{lineItemId}/return_line_items | Retrieve the return line items associated to the line item
[**GETReturnIdReturnLineItems**](ReturnLineItemsApi.md#GETReturnIdReturnLineItems) | **Get** /returns/{returnId}/return_line_items | Retrieve the return line items associated to the return
[**GETReturnLineItems**](ReturnLineItemsApi.md#GETReturnLineItems) | **Get** /return_line_items | List all return line items
[**GETReturnLineItemsReturnLineItemId**](ReturnLineItemsApi.md#GETReturnLineItemsReturnLineItemId) | **Get** /return_line_items/{returnLineItemId} | Retrieve a return line item
[**PATCHReturnLineItemsReturnLineItemId**](ReturnLineItemsApi.md#PATCHReturnLineItemsReturnLineItemId) | **Patch** /return_line_items/{returnLineItemId} | Update a return line item
[**POSTReturnLineItems**](ReturnLineItemsApi.md#POSTReturnLineItems) | **Post** /return_line_items | Create a return line item



## DELETEReturnLineItemsReturnLineItemId

> DELETEReturnLineItemsReturnLineItemId(ctx, returnLineItemId).Execute()

Delete a return line item



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
    returnLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReturnLineItemsApi.DELETEReturnLineItemsReturnLineItemId(context.Background(), returnLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.DELETEReturnLineItemsReturnLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEReturnLineItemsReturnLineItemIdRequest struct via the builder pattern


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


## GETLineItemIdReturnLineItems

> GETLineItemIdReturnLineItems(ctx, lineItemId).Execute()

Retrieve the return line items associated to the line item



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
    lineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReturnLineItemsApi.GETLineItemIdReturnLineItems(context.Background(), lineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.GETLineItemIdReturnLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETLineItemIdReturnLineItemsRequest struct via the builder pattern


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


## GETReturnIdReturnLineItems

> GETReturnIdReturnLineItems(ctx, returnId).Execute()

Retrieve the return line items associated to the return



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
    returnId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReturnLineItemsApi.GETReturnIdReturnLineItems(context.Background(), returnId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.GETReturnIdReturnLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnIdReturnLineItemsRequest struct via the builder pattern


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


## GETReturnLineItems

> GETReturnLineItems200Response GETReturnLineItems(ctx).Execute()

List all return line items



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
    resp, r, err := apiClient.ReturnLineItemsApi.GETReturnLineItems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.GETReturnLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETReturnLineItems`: GETReturnLineItems200Response
    fmt.Fprintf(os.Stdout, "Response from `ReturnLineItemsApi.GETReturnLineItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnLineItemsRequest struct via the builder pattern


### Return type

[**GETReturnLineItems200Response**](GETReturnLineItems200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETReturnLineItemsReturnLineItemId

> GETReturnLineItemsReturnLineItemId200Response GETReturnLineItemsReturnLineItemId(ctx, returnLineItemId).Execute()

Retrieve a return line item



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
    returnLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnLineItemsApi.GETReturnLineItemsReturnLineItemId(context.Background(), returnLineItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.GETReturnLineItemsReturnLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETReturnLineItemsReturnLineItemId`: GETReturnLineItemsReturnLineItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `ReturnLineItemsApi.GETReturnLineItemsReturnLineItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETReturnLineItemsReturnLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETReturnLineItemsReturnLineItemId200Response**](GETReturnLineItemsReturnLineItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHReturnLineItemsReturnLineItemId

> PATCHReturnLineItemsReturnLineItemId200Response PATCHReturnLineItemsReturnLineItemId(ctx, returnLineItemId).ReturnLineItemUpdate(returnLineItemUpdate).Execute()

Update a return line item



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
    returnLineItemUpdate := *openapiclient.NewReturnLineItemUpdate(*openapiclient.NewReturnLineItemUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHReturnLineItemsReturnLineItemId200ResponseDataAttributes())) // ReturnLineItemUpdate | 
    returnLineItemId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnLineItemsApi.PATCHReturnLineItemsReturnLineItemId(context.Background(), returnLineItemId).ReturnLineItemUpdate(returnLineItemUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.PATCHReturnLineItemsReturnLineItemId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHReturnLineItemsReturnLineItemId`: PATCHReturnLineItemsReturnLineItemId200Response
    fmt.Fprintf(os.Stdout, "Response from `ReturnLineItemsApi.PATCHReturnLineItemsReturnLineItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**returnLineItemId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHReturnLineItemsReturnLineItemIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnLineItemUpdate** | [**ReturnLineItemUpdate**](ReturnLineItemUpdate.md) |  | 


### Return type

[**PATCHReturnLineItemsReturnLineItemId200Response**](PATCHReturnLineItemsReturnLineItemId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTReturnLineItems

> POSTReturnLineItems201Response POSTReturnLineItems(ctx).ReturnLineItemCreate(returnLineItemCreate).Execute()

Create a return line item



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
    returnLineItemCreate := *openapiclient.NewReturnLineItemCreate(*openapiclient.NewReturnLineItemCreateData(interface{}(123), *openapiclient.NewPOSTReturnLineItems201ResponseDataAttributes(interface{}(4)))) // ReturnLineItemCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnLineItemsApi.POSTReturnLineItems(context.Background()).ReturnLineItemCreate(returnLineItemCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnLineItemsApi.POSTReturnLineItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTReturnLineItems`: POSTReturnLineItems201Response
    fmt.Fprintf(os.Stdout, "Response from `ReturnLineItemsApi.POSTReturnLineItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTReturnLineItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **returnLineItemCreate** | [**ReturnLineItemCreate**](ReturnLineItemCreate.md) |  | 

### Return type

[**POSTReturnLineItems201Response**](POSTReturnLineItems201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

