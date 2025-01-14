# OrderSubscriptionCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BundleCreateDataRelationshipsMarket**](BundleCreateDataRelationshipsMarket.md) |  | [optional] 
**SourceOrder** | [**AdyenPaymentCreateDataRelationshipsOrder**](AdyenPaymentCreateDataRelationshipsOrder.md) |  | 
**Tags** | Pointer to [**AddressCreateDataRelationshipsTags**](AddressCreateDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewOrderSubscriptionCreateDataRelationships

`func NewOrderSubscriptionCreateDataRelationships(sourceOrder AdyenPaymentCreateDataRelationshipsOrder, ) *OrderSubscriptionCreateDataRelationships`

NewOrderSubscriptionCreateDataRelationships instantiates a new OrderSubscriptionCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderSubscriptionCreateDataRelationshipsWithDefaults

`func NewOrderSubscriptionCreateDataRelationshipsWithDefaults() *OrderSubscriptionCreateDataRelationships`

NewOrderSubscriptionCreateDataRelationshipsWithDefaults instantiates a new OrderSubscriptionCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *OrderSubscriptionCreateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderSubscriptionCreateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderSubscriptionCreateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *OrderSubscriptionCreateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetSourceOrder

`func (o *OrderSubscriptionCreateDataRelationships) GetSourceOrder() AdyenPaymentCreateDataRelationshipsOrder`

GetSourceOrder returns the SourceOrder field if non-nil, zero value otherwise.

### GetSourceOrderOk

`func (o *OrderSubscriptionCreateDataRelationships) GetSourceOrderOk() (*AdyenPaymentCreateDataRelationshipsOrder, bool)`

GetSourceOrderOk returns a tuple with the SourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOrder

`func (o *OrderSubscriptionCreateDataRelationships) SetSourceOrder(v AdyenPaymentCreateDataRelationshipsOrder)`

SetSourceOrder sets SourceOrder field to given value.


### GetTags

`func (o *OrderSubscriptionCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrderSubscriptionCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrderSubscriptionCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrderSubscriptionCreateDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


