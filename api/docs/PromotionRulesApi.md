# \PromotionRulesApi

All URIs are relative to *https://.commercelayer.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GETPromotionRules**](PromotionRulesApi.md#GETPromotionRules) | **Get** /promotion_rules | List all promotion rules
[**GETPromotionRulesPromotionRuleId**](PromotionRulesApi.md#GETPromotionRulesPromotionRuleId) | **Get** /promotion_rules/{promotionRuleId} | Retrieve a promotion rule



## GETPromotionRules

> GETPromotionRules200Response GETPromotionRules(ctx).Execute()

List all promotion rules



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
    resp, r, err := apiClient.PromotionRulesApi.GETPromotionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionRulesApi.GETPromotionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPromotionRules`: GETPromotionRules200Response
    fmt.Fprintf(os.Stdout, "Response from `PromotionRulesApi.GETPromotionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGETPromotionRulesRequest struct via the builder pattern


### Return type

[**GETPromotionRules200Response**](GETPromotionRules200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GETPromotionRulesPromotionRuleId

> GETPromotionRulesPromotionRuleId200Response GETPromotionRulesPromotionRuleId(ctx, promotionRuleId).Execute()

Retrieve a promotion rule



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
    promotionRuleId := TODO // interface{} | The resource's id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotionRulesApi.GETPromotionRulesPromotionRuleId(context.Background(), promotionRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotionRulesApi.GETPromotionRulesPromotionRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GETPromotionRulesPromotionRuleId`: GETPromotionRulesPromotionRuleId200Response
    fmt.Fprintf(os.Stdout, "Response from `PromotionRulesApi.GETPromotionRulesPromotionRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**promotionRuleId** | [**interface{}**](.md) | The resource&#39;s id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGETPromotionRulesPromotionRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GETPromotionRulesPromotionRuleId200Response**](GETPromotionRulesPromotionRuleId200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.api+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

