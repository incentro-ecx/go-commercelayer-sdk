# \CustomerGroupsApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETECustomerGroupsCustomerGroupId**](CustomerGroupsApi.md#DELETECustomerGroupsCustomerGroupId) | **Delete** /customer_groups/{customerGroupId} | Delete a customer group
[**GETCustomerGroups**](CustomerGroupsApi.md#GETCustomerGroups) | **Get** /customer_groups | List all customer groups
[**GETCustomerGroupsCustomerGroupId**](CustomerGroupsApi.md#GETCustomerGroupsCustomerGroupId) | **Get** /customer_groups/{customerGroupId} | Retrieve a customer group
[**GETCustomerIdCustomerGroup**](CustomerGroupsApi.md#GETCustomerIdCustomerGroup) | **Get** /customers/{customerId}/customer_group | Retrieve the customer group associated to the customer
[**GETMarketIdCustomerGroup**](CustomerGroupsApi.md#GETMarketIdCustomerGroup) | **Get** /markets/{marketId}/customer_group | Retrieve the customer group associated to the market
[**PATCHCustomerGroupsCustomerGroupId**](CustomerGroupsApi.md#PATCHCustomerGroupsCustomerGroupId) | **Patch** /customer_groups/{customerGroupId} | Update a customer group
[**POSTCustomerGroups**](CustomerGroupsApi.md#POSTCustomerGroups) | **Post** /customer_groups | Create a customer group



## DELETECustomerGroupsCustomerGroupId

> DELETECustomerGroupsCustomerGroupId(ctx, customerGroupId).Execute()

Delete a customer group



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
    customerGroupId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomerGroupsApi.DELETECustomerGroupsCustomerGroupId(context.Background(), customerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.DELETECustomerGroupsCustomerGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerGroupId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETECustomerGroupsCustomerGroupIdRequest struct via the builder pattern


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


## GETCustomerGroups

> GETCustomerGroups200Response GETCustomerGroups(ctx).Execute()

List all customer groups



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
    resp, r, err := apiClient.CustomerGroupsApi.GETCustomerGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.GETCustomerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerGroups`: GETCustomerGroups200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsApi.GETCustomerGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerGroupsRequest struct via the builder pattern


### Return type

[**GETCustomerGroups200Response**](GETCustomerGroups200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCustomerGroupsCustomerGroupId

> GETCustomerGroupsCustomerGroupId200Response GETCustomerGroupsCustomerGroupId(ctx, customerGroupId).Execute()

Retrieve a customer group



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
    customerGroupId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerGroupsApi.GETCustomerGroupsCustomerGroupId(context.Background(), customerGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.GETCustomerGroupsCustomerGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETCustomerGroupsCustomerGroupId`: GETCustomerGroupsCustomerGroupId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsApi.GETCustomerGroupsCustomerGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerGroupId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerGroupsCustomerGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETCustomerGroupsCustomerGroupId200Response**](GETCustomerGroupsCustomerGroupId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETCustomerIdCustomerGroup

> GETCustomerIdCustomerGroup(ctx, customerId).Execute()

Retrieve the customer group associated to the customer



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
    customerId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomerGroupsApi.GETCustomerIdCustomerGroup(context.Background(), customerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.GETCustomerIdCustomerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETCustomerIdCustomerGroupRequest struct via the builder pattern


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


## GETMarketIdCustomerGroup

> GETMarketIdCustomerGroup(ctx, marketId).Execute()

Retrieve the customer group associated to the market



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
    marketId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CustomerGroupsApi.GETMarketIdCustomerGroup(context.Background(), marketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.GETMarketIdCustomerGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**marketId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETMarketIdCustomerGroupRequest struct via the builder pattern


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


## PATCHCustomerGroupsCustomerGroupId

> PATCHCustomerGroupsCustomerGroupId200Response PATCHCustomerGroupsCustomerGroupId(ctx, customerGroupId).CustomerGroupUpdate(customerGroupUpdate).Execute()

Update a customer group



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
    customerGroupUpdate := *openapiclient.NewCustomerGroupUpdate(*openapiclient.NewCustomerGroupUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHCustomerGroupsCustomerGroupId200ResponseDataAttributes())) // CustomerGroupUpdate | 
    customerGroupId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerGroupsApi.PATCHCustomerGroupsCustomerGroupId(context.Background(), customerGroupId).CustomerGroupUpdate(customerGroupUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.PATCHCustomerGroupsCustomerGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHCustomerGroupsCustomerGroupId`: PATCHCustomerGroupsCustomerGroupId200Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsApi.PATCHCustomerGroupsCustomerGroupId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerGroupId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHCustomerGroupsCustomerGroupIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerGroupUpdate** | [**CustomerGroupUpdate**](CustomerGroupUpdate.md) |  | 


### Return type

[**PATCHCustomerGroupsCustomerGroupId200Response**](PATCHCustomerGroupsCustomerGroupId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTCustomerGroups

> POSTCustomerGroups201Response POSTCustomerGroups(ctx).CustomerGroupCreate(customerGroupCreate).Execute()

Create a customer group



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
    customerGroupCreate := *openapiclient.NewCustomerGroupCreate(*openapiclient.NewCustomerGroupCreateData(interface{}(123), *openapiclient.NewPOSTCustomerGroups201ResponseDataAttributes(interface{}(VIP)))) // CustomerGroupCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomerGroupsApi.POSTCustomerGroups(context.Background()).CustomerGroupCreate(customerGroupCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerGroupsApi.POSTCustomerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTCustomerGroups`: POSTCustomerGroups201Response
    fmt.Fprintf(os.Stdout, "Response from `CustomerGroupsApi.POSTCustomerGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTCustomerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerGroupCreate** | [**CustomerGroupCreate**](CustomerGroupCreate.md) |  | 

### Return type

[**POSTCustomerGroups201Response**](POSTCustomerGroups201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

