# BuyXPayYPromotionCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BundleCreateDataRelationshipsMarket**](BundleCreateDataRelationshipsMarket.md) |  | [optional] 
**OrderAmountPromotionRule** | Pointer to [**BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule**](BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule.md) |  | [optional] 
**SkuListPromotionRule** | Pointer to [**BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule**](BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule.md) |  | [optional] 
**CouponCodesPromotionRule** | Pointer to [**BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule**](BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule.md) |  | [optional] 
**CustomPromotionRule** | Pointer to [**BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule**](BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule.md) |  | [optional] 
**SkuList** | [**BundleCreateDataRelationshipsSkuList**](BundleCreateDataRelationshipsSkuList.md) |  | 
**Tags** | Pointer to [**AddressCreateDataRelationshipsTags**](AddressCreateDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewBuyXPayYPromotionCreateDataRelationships

`func NewBuyXPayYPromotionCreateDataRelationships(skuList BundleCreateDataRelationshipsSkuList, ) *BuyXPayYPromotionCreateDataRelationships`

NewBuyXPayYPromotionCreateDataRelationships instantiates a new BuyXPayYPromotionCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuyXPayYPromotionCreateDataRelationshipsWithDefaults

`func NewBuyXPayYPromotionCreateDataRelationshipsWithDefaults() *BuyXPayYPromotionCreateDataRelationships`

NewBuyXPayYPromotionCreateDataRelationshipsWithDefaults instantiates a new BuyXPayYPromotionCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *BuyXPayYPromotionCreateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *BuyXPayYPromotionCreateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *BuyXPayYPromotionCreateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *BuyXPayYPromotionCreateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetOrderAmountPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) GetOrderAmountPromotionRule() BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule`

GetOrderAmountPromotionRule returns the OrderAmountPromotionRule field if non-nil, zero value otherwise.

### GetOrderAmountPromotionRuleOk

`func (o *BuyXPayYPromotionCreateDataRelationships) GetOrderAmountPromotionRuleOk() (*BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule, bool)`

GetOrderAmountPromotionRuleOk returns a tuple with the OrderAmountPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAmountPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) SetOrderAmountPromotionRule(v BuyXPayYPromotionCreateDataRelationshipsOrderAmountPromotionRule)`

SetOrderAmountPromotionRule sets OrderAmountPromotionRule field to given value.

### HasOrderAmountPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) HasOrderAmountPromotionRule() bool`

HasOrderAmountPromotionRule returns a boolean if a field has been set.

### GetSkuListPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) GetSkuListPromotionRule() BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule`

GetSkuListPromotionRule returns the SkuListPromotionRule field if non-nil, zero value otherwise.

### GetSkuListPromotionRuleOk

`func (o *BuyXPayYPromotionCreateDataRelationships) GetSkuListPromotionRuleOk() (*BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule, bool)`

GetSkuListPromotionRuleOk returns a tuple with the SkuListPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuListPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) SetSkuListPromotionRule(v BuyXPayYPromotionCreateDataRelationshipsSkuListPromotionRule)`

SetSkuListPromotionRule sets SkuListPromotionRule field to given value.

### HasSkuListPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) HasSkuListPromotionRule() bool`

HasSkuListPromotionRule returns a boolean if a field has been set.

### GetCouponCodesPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) GetCouponCodesPromotionRule() BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule`

GetCouponCodesPromotionRule returns the CouponCodesPromotionRule field if non-nil, zero value otherwise.

### GetCouponCodesPromotionRuleOk

`func (o *BuyXPayYPromotionCreateDataRelationships) GetCouponCodesPromotionRuleOk() (*BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule, bool)`

GetCouponCodesPromotionRuleOk returns a tuple with the CouponCodesPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCodesPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) SetCouponCodesPromotionRule(v BuyXPayYPromotionCreateDataRelationshipsCouponCodesPromotionRule)`

SetCouponCodesPromotionRule sets CouponCodesPromotionRule field to given value.

### HasCouponCodesPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) HasCouponCodesPromotionRule() bool`

HasCouponCodesPromotionRule returns a boolean if a field has been set.

### GetCustomPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) GetCustomPromotionRule() BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule`

GetCustomPromotionRule returns the CustomPromotionRule field if non-nil, zero value otherwise.

### GetCustomPromotionRuleOk

`func (o *BuyXPayYPromotionCreateDataRelationships) GetCustomPromotionRuleOk() (*BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule, bool)`

GetCustomPromotionRuleOk returns a tuple with the CustomPromotionRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) SetCustomPromotionRule(v BuyXPayYPromotionCreateDataRelationshipsCustomPromotionRule)`

SetCustomPromotionRule sets CustomPromotionRule field to given value.

### HasCustomPromotionRule

`func (o *BuyXPayYPromotionCreateDataRelationships) HasCustomPromotionRule() bool`

HasCustomPromotionRule returns a boolean if a field has been set.

### GetSkuList

`func (o *BuyXPayYPromotionCreateDataRelationships) GetSkuList() BundleCreateDataRelationshipsSkuList`

GetSkuList returns the SkuList field if non-nil, zero value otherwise.

### GetSkuListOk

`func (o *BuyXPayYPromotionCreateDataRelationships) GetSkuListOk() (*BundleCreateDataRelationshipsSkuList, bool)`

GetSkuListOk returns a tuple with the SkuList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuList

`func (o *BuyXPayYPromotionCreateDataRelationships) SetSkuList(v BundleCreateDataRelationshipsSkuList)`

SetSkuList sets SkuList field to given value.


### GetTags

`func (o *BuyXPayYPromotionCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BuyXPayYPromotionCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BuyXPayYPromotionCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BuyXPayYPromotionCreateDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


