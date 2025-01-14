# OrderCreateDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**BundleCreateDataRelationshipsMarket**](BundleCreateDataRelationshipsMarket.md) |  | [optional] 
**Customer** | Pointer to [**CouponRecipientCreateDataRelationshipsCustomer**](CouponRecipientCreateDataRelationshipsCustomer.md) |  | [optional] 
**ShippingAddress** | Pointer to [**CustomerAddressCreateDataRelationshipsAddress**](CustomerAddressCreateDataRelationshipsAddress.md) |  | [optional] 
**BillingAddress** | Pointer to [**CustomerAddressCreateDataRelationshipsAddress**](CustomerAddressCreateDataRelationshipsAddress.md) |  | [optional] 
**Store** | Pointer to [**OrderCreateDataRelationshipsStore**](OrderCreateDataRelationshipsStore.md) |  | [optional] 
**PaymentMethod** | Pointer to [**CustomerPaymentSourceCreateDataRelationshipsPaymentMethod**](CustomerPaymentSourceCreateDataRelationshipsPaymentMethod.md) |  | [optional] 
**PaymentSource** | Pointer to [**CustomerPaymentSourceCreateDataRelationshipsPaymentSource**](CustomerPaymentSourceCreateDataRelationshipsPaymentSource.md) |  | [optional] 
**Tags** | Pointer to [**AddressCreateDataRelationshipsTags**](AddressCreateDataRelationshipsTags.md) |  | [optional] 

## Methods

### NewOrderCreateDataRelationships

`func NewOrderCreateDataRelationships() *OrderCreateDataRelationships`

NewOrderCreateDataRelationships instantiates a new OrderCreateDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreateDataRelationshipsWithDefaults

`func NewOrderCreateDataRelationshipsWithDefaults() *OrderCreateDataRelationships`

NewOrderCreateDataRelationshipsWithDefaults instantiates a new OrderCreateDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *OrderCreateDataRelationships) GetMarket() BundleCreateDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *OrderCreateDataRelationships) GetMarketOk() (*BundleCreateDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *OrderCreateDataRelationships) SetMarket(v BundleCreateDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *OrderCreateDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderCreateDataRelationships) GetCustomer() CouponRecipientCreateDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderCreateDataRelationships) GetCustomerOk() (*CouponRecipientCreateDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderCreateDataRelationships) SetCustomer(v CouponRecipientCreateDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderCreateDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetShippingAddress

`func (o *OrderCreateDataRelationships) GetShippingAddress() CustomerAddressCreateDataRelationshipsAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *OrderCreateDataRelationships) GetShippingAddressOk() (*CustomerAddressCreateDataRelationshipsAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *OrderCreateDataRelationships) SetShippingAddress(v CustomerAddressCreateDataRelationshipsAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *OrderCreateDataRelationships) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetBillingAddress

`func (o *OrderCreateDataRelationships) GetBillingAddress() CustomerAddressCreateDataRelationshipsAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *OrderCreateDataRelationships) GetBillingAddressOk() (*CustomerAddressCreateDataRelationshipsAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *OrderCreateDataRelationships) SetBillingAddress(v CustomerAddressCreateDataRelationshipsAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *OrderCreateDataRelationships) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetStore

`func (o *OrderCreateDataRelationships) GetStore() OrderCreateDataRelationshipsStore`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *OrderCreateDataRelationships) GetStoreOk() (*OrderCreateDataRelationshipsStore, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *OrderCreateDataRelationships) SetStore(v OrderCreateDataRelationshipsStore)`

SetStore sets Store field to given value.

### HasStore

`func (o *OrderCreateDataRelationships) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *OrderCreateDataRelationships) GetPaymentMethod() CustomerPaymentSourceCreateDataRelationshipsPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *OrderCreateDataRelationships) GetPaymentMethodOk() (*CustomerPaymentSourceCreateDataRelationshipsPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *OrderCreateDataRelationships) SetPaymentMethod(v CustomerPaymentSourceCreateDataRelationshipsPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *OrderCreateDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *OrderCreateDataRelationships) GetPaymentSource() CustomerPaymentSourceCreateDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *OrderCreateDataRelationships) GetPaymentSourceOk() (*CustomerPaymentSourceCreateDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *OrderCreateDataRelationships) SetPaymentSource(v CustomerPaymentSourceCreateDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *OrderCreateDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetTags

`func (o *OrderCreateDataRelationships) GetTags() AddressCreateDataRelationshipsTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrderCreateDataRelationships) GetTagsOk() (*AddressCreateDataRelationshipsTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrderCreateDataRelationships) SetTags(v AddressCreateDataRelationshipsTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrderCreateDataRelationships) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


