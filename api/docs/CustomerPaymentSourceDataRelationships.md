# CustomerPaymentSourceDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customer** | Pointer to [**CouponRecipientDataRelationshipsCustomer**](CouponRecipientDataRelationshipsCustomer.md) |  | [optional] 
**PaymentMethod** | Pointer to [**AdyenGatewayDataRelationshipsPaymentMethods**](AdyenGatewayDataRelationshipsPaymentMethods.md) |  | [optional] 
**PaymentSource** | Pointer to [**AuthorizationDataRelationshipsPaymentSource**](AuthorizationDataRelationshipsPaymentSource.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewCustomerPaymentSourceDataRelationships

`func NewCustomerPaymentSourceDataRelationships() *CustomerPaymentSourceDataRelationships`

NewCustomerPaymentSourceDataRelationships instantiates a new CustomerPaymentSourceDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentSourceDataRelationshipsWithDefaults

`func NewCustomerPaymentSourceDataRelationshipsWithDefaults() *CustomerPaymentSourceDataRelationships`

NewCustomerPaymentSourceDataRelationshipsWithDefaults instantiates a new CustomerPaymentSourceDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomer

`func (o *CustomerPaymentSourceDataRelationships) GetCustomer() CouponRecipientDataRelationshipsCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *CustomerPaymentSourceDataRelationships) GetCustomerOk() (*CouponRecipientDataRelationshipsCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *CustomerPaymentSourceDataRelationships) SetCustomer(v CouponRecipientDataRelationshipsCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *CustomerPaymentSourceDataRelationships) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *CustomerPaymentSourceDataRelationships) GetPaymentMethod() AdyenGatewayDataRelationshipsPaymentMethods`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *CustomerPaymentSourceDataRelationships) GetPaymentMethodOk() (*AdyenGatewayDataRelationshipsPaymentMethods, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *CustomerPaymentSourceDataRelationships) SetPaymentMethod(v AdyenGatewayDataRelationshipsPaymentMethods)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *CustomerPaymentSourceDataRelationships) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentSource

`func (o *CustomerPaymentSourceDataRelationships) GetPaymentSource() AuthorizationDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *CustomerPaymentSourceDataRelationships) GetPaymentSourceOk() (*AuthorizationDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *CustomerPaymentSourceDataRelationships) SetPaymentSource(v AuthorizationDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *CustomerPaymentSourceDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetVersions

`func (o *CustomerPaymentSourceDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CustomerPaymentSourceDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CustomerPaymentSourceDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CustomerPaymentSourceDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


