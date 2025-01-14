# \WebhooksApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DELETEWebhooksWebhookId**](WebhooksApi.md#DELETEWebhooksWebhookId) | **Delete** /webhooks/{webhookId} | Delete a webhook
[**GETEventCallbackIdWebhook**](WebhooksApi.md#GETEventCallbackIdWebhook) | **Get** /event_callbacks/{eventCallbackId}/webhook | Retrieve the webhook associated to the event callback
[**GETEventIdWebhooks**](WebhooksApi.md#GETEventIdWebhooks) | **Get** /events/{eventId}/webhooks | Retrieve the webhooks associated to the event
[**GETWebhooks**](WebhooksApi.md#GETWebhooks) | **Get** /webhooks | List all webhooks
[**GETWebhooksWebhookId**](WebhooksApi.md#GETWebhooksWebhookId) | **Get** /webhooks/{webhookId} | Retrieve a webhook
[**PATCHWebhooksWebhookId**](WebhooksApi.md#PATCHWebhooksWebhookId) | **Patch** /webhooks/{webhookId} | Update a webhook
[**POSTWebhooks**](WebhooksApi.md#POSTWebhooks) | **Post** /webhooks | Create a webhook



## DELETEWebhooksWebhookId

> DELETEWebhooksWebhookId(ctx, webhookId).Execute()

Delete a webhook



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
    webhookId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.DELETEWebhooksWebhookId(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DELETEWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDELETEWebhooksWebhookIdRequest struct via the builder pattern


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


## GETEventCallbackIdWebhook

> GETEventCallbackIdWebhook(ctx, eventCallbackId).Execute()

Retrieve the webhook associated to the event callback



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
    eventCallbackId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.GETEventCallbackIdWebhook(context.Background(), eventCallbackId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GETEventCallbackIdWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventCallbackId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETEventCallbackIdWebhookRequest struct via the builder pattern


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


## GETEventIdWebhooks

> GETEventIdWebhooks(ctx, eventId).Execute()

Retrieve the webhooks associated to the event



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
    eventId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.GETEventIdWebhooks(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GETEventIdWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETEventIdWebhooksRequest struct via the builder pattern


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


## GETWebhooks

> GETWebhooks200Response GETWebhooks(ctx).Execute()

List all webhooks



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
    resp, r, err := apiClient.WebhooksApi.GETWebhooks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GETWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETWebhooks`: GETWebhooks200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GETWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETWebhooksRequest struct via the builder pattern


### Return type

[**GETWebhooks200Response**](GETWebhooks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETWebhooksWebhookId

> GETWebhooksWebhookId200Response GETWebhooksWebhookId(ctx, webhookId).Execute()

Retrieve a webhook



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
    webhookId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GETWebhooksWebhookId(context.Background(), webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GETWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETWebhooksWebhookId`: GETWebhooksWebhookId200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GETWebhooksWebhookId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETWebhooksWebhookIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETWebhooksWebhookId200Response**](GETWebhooksWebhookId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PATCHWebhooksWebhookId

> PATCHWebhooksWebhookId200Response PATCHWebhooksWebhookId(ctx, webhookId).WebhookUpdate(webhookUpdate).Execute()

Update a webhook



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
    webhookUpdate := *openapiclient.NewWebhookUpdate(*openapiclient.NewWebhookUpdateData(interface{}(123), interface{}(XGZwpOSrWL), *openapiclient.NewPATCHWebhooksWebhookId200ResponseDataAttributes())) // WebhookUpdate | 
    webhookId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.PATCHWebhooksWebhookId(context.Background(), webhookId).WebhookUpdate(webhookUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.PATCHWebhooksWebhookId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PATCHWebhooksWebhookId`: PATCHWebhooksWebhookId200Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.PATCHWebhooksWebhookId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPATCHWebhooksWebhookIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookUpdate** | [**WebhookUpdate**](WebhookUpdate.md) |  | 


### Return type

[**PATCHWebhooksWebhookId200Response**](PATCHWebhooksWebhookId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## POSTWebhooks

> POSTWebhooks201Response POSTWebhooks(ctx).WebhookCreate(webhookCreate).Execute()

Create a webhook



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
    webhookCreate := *openapiclient.NewWebhookCreate(*openapiclient.NewWebhookCreateData(interface{}(123), *openapiclient.NewPOSTWebhooks201ResponseDataAttributes(interface{}(orders.place), interface{}(https://yourapp.com/webhooks)))) // WebhookCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.POSTWebhooks(context.Background()).WebhookCreate(webhookCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.POSTWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `POSTWebhooks`: POSTWebhooks201Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.POSTWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPOSTWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookCreate** | [**WebhookCreate**](WebhookCreate.md) |  | 

### Return type

[**POSTWebhooks201Response**](POSTWebhooks201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.api+json
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

