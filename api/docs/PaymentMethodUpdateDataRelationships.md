# PaymentMethodUpdateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BundleCreateDataRelationshipsMarket**](BundleCreateDataRelationshipsMarket.md) |  | [optional] 
**PaymentGateway** | Pointer to [**PaymentMethodCreateDataRelationshipsPaymentGateway**](PaymentMethodCreateDataRelationshipsPaymentGateway.md) |  | [optional] 
**Store** | Pointer to [**OrderCreateDataRelationshipsStore**](OrderCreateDataRelationshipsStore.md) |  | [optional] 

## Methods

### NewPaymentMethodUpdateDataRelationships

`func NewPaymentMethodUpdateDataRelationships() *PaymentMethodUpdateDataRelationships`

NewPaymentMethodUpdateDataRelationships instantiates a new PaymentMethodUpdateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodUpdateDataRelationshipsWithDefaults

`func NewPaymentMethodUpdateDataRelationshipsWithDefaults() *PaymentMethodUpdateDataRelationships`

NewPaymentMethodUpdateDataRelationshipsWithDefaults instantiates a new PaymentMethodUpdateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *PaymentMethodUpdateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *PaymentMethodUpdateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *PaymentMethodUpdateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *PaymentMethodUpdateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *PaymentMethodUpdateDataRelationships) GetPaymentGateway() PaymentMethodCreateDataRelationshipsPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *PaymentMethodUpdateDataRelationships) GetPaymentGatewayOk() (*PaymentMethodCreateDataRelationshipsPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *PaymentMethodUpdateDataRelationships) SetPaymentGateway(v PaymentMethodCreateDataRelationshipsPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.

### HasPaymentGateway

`func (o *PaymentMethodUpdateDataRelationships) HasPaymentGateway() bool`

HasPaymentGateway returns a boolean if a field has been set.

### GetStore

`func (o *PaymentMethodUpdateDataRelationships) GetStore() OrderCreateDataRelationshipsStore`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *PaymentMethodUpdateDataRelationships) GetStoreOk() (*OrderCreateDataRelationshipsStore, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *PaymentMethodUpdateDataRelationships) SetStore(v OrderCreateDataRelationshipsStore)`

SetStore sets Store field to given value.

### HasStore

`func (o *PaymentMethodUpdateDataRelationships) HasStore() bool`

HasStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


