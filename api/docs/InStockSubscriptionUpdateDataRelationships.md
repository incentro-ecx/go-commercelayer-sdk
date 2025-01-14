# InStockSubscriptionUpdateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BundleCreateDataRelationshipsMarket**](BundleCreateDataRelationshipsMarket.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientCreateDataRelationshipsCustomer**](CouponRecipientCreateDataRelationshipsCustomer.md) |  | [optional] 
**Sku** | Pointer to [**InStockSubscriptionCreateDataRelationshipsSku**](InStockSubscriptionCreateDataRelationshipsSku.md) |  | [optional] 

## Methods

### NewInStockSubscriptionUpdateDataRelationships

`func NewInStockSubscriptionUpdateDataRelationships() *InStockSubscriptionUpdateDataRelationships`

NewInStockSubscriptionUpdateDataRelationships instantiates a new InStockSubscriptionUpdateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInStockSubscriptionUpdateDataRelationshipsWithDefaults

`func NewInStockSubscriptionUpdateDataRelationshipsWithDefaults() *InStockSubscriptionUpdateDataRelationships`

NewInStockSubscriptionUpdateDataRelationshipsWithDefaults instantiates a new InStockSubscriptionUpdateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *InStockSubscriptionUpdateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *InStockSubscriptionUpdateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *InStockSubscriptionUpdateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *InStockSubscriptionUpdateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *InStockSubscriptionUpdateDataRelationships) GetCustomer() CouponRecipientCreateDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *InStockSubscriptionUpdateDataRelationships) GetCustomerOk() (*CouponRecipientCreateDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *InStockSubscriptionUpdateDataRelationships) SetCustomer(v CouponRecipientCreateDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *InStockSubscriptionUpdateDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetSku

`func (o *InStockSubscriptionUpdateDataRelationships) GetSku() InStockSubscriptionCreateDataRelationshipsSku`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *InStockSubscriptionUpdateDataRelationships) GetSkuOk() (*InStockSubscriptionCreateDataRelationshipsSku, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *InStockSubscriptionUpdateDataRelationships) SetSku(v InStockSubscriptionCreateDataRelationshipsSku)`

SetSku sets Sku field to given value.

### HasSku

`func (o *InStockSubscriptionUpdateDataRelationships) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


