# POSTTaxRules201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The tax rule internal name. | 
**TaxRate** | Pointer to **interface{}** | The tax rate for this rule. | [optional] 
**FreightTaxable** | Pointer to **interface{}** | Indicates if the freight is taxable. | [optional] 
**PaymentMethodTaxable** | Pointer to **interface{}** | Indicates if the payment method is taxable. | [optional] 
**GiftCardTaxable** | Pointer to **interface{}** | Indicates if gift cards are taxable. | [optional] 
**AdjustmentTaxable** | Pointer to **interface{}** | Indicates if adjustemnts are taxable. | [optional] 
**CountryCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the shipping address country code, max size is 5000. | [optional] 
**NotCountryCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated as negative match for the shipping address country code, max size is 5000. | [optional] 
**StateCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the shipping address state code, max size is 5000. | [optional] 
**NotStateCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated as negative match for the shipping address state code, max size is 5000. | [optional] 
**ZipCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the shipping address zip code, max size is 5000. | [optional] 
**NotZipCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated as negative match for the shipping zip country code, max size is 5000. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTTaxRules201ResponseDataAttributes

`func NewPOSTTaxRules201ResponseDataAttributes(name interface{}, ) *POSTTaxRules201ResponseDataAttributes`

NewPOSTTaxRules201ResponseDataAttributes instantiates a new POSTTaxRules201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTTaxRules201ResponseDataAttributesWithDefaults

`func NewPOSTTaxRules201ResponseDataAttributesWithDefaults() *POSTTaxRules201ResponseDataAttributes`

NewPOSTTaxRules201ResponseDataAttributesWithDefaults instantiates a new POSTTaxRules201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTTaxRules201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTTaxRules201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaxRate

`func (o *POSTTaxRules201ResponseDataAttributes) GetTaxRate() interface{}`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetTaxRateOk() (*interface{}, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *POSTTaxRules201ResponseDataAttributes) SetTaxRate(v interface{})`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *POSTTaxRules201ResponseDataAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetFreightTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) GetFreightTaxable() interface{}`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetFreightTaxableOk() (*interface{}, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) SetFreightTaxable(v interface{})`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### SetFreightTaxableNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetFreightTaxableNil(b bool)`

 SetFreightTaxableNil sets the value for FreightTaxable to be an explicit nil

### UnsetFreightTaxable
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetFreightTaxable()`

UnsetFreightTaxable ensures that no value is present for FreightTaxable, not even an explicit nil
### GetPaymentMethodTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) GetPaymentMethodTaxable() interface{}`

GetPaymentMethodTaxable returns the PaymentMethodTaxable field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetPaymentMethodTaxableOk() (*interface{}, bool)`

GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) SetPaymentMethodTaxable(v interface{})`

SetPaymentMethodTaxable sets PaymentMethodTaxable field to given value.

### HasPaymentMethodTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) HasPaymentMethodTaxable() bool`

HasPaymentMethodTaxable returns a boolean if a field has been set.

### SetPaymentMethodTaxableNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetPaymentMethodTaxableNil(b bool)`

 SetPaymentMethodTaxableNil sets the value for PaymentMethodTaxable to be an explicit nil

### UnsetPaymentMethodTaxable
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetPaymentMethodTaxable()`

UnsetPaymentMethodTaxable ensures that no value is present for PaymentMethodTaxable, not even an explicit nil
### GetGiftCardTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) GetGiftCardTaxable() interface{}`

GetGiftCardTaxable returns the GiftCardTaxable field if non-nil, zero value otherwise.

### GetGiftCardTaxableOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetGiftCardTaxableOk() (*interface{}, bool)`

GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) SetGiftCardTaxable(v interface{})`

SetGiftCardTaxable sets GiftCardTaxable field to given value.

### HasGiftCardTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) HasGiftCardTaxable() bool`

HasGiftCardTaxable returns a boolean if a field has been set.

### SetGiftCardTaxableNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetGiftCardTaxableNil(b bool)`

 SetGiftCardTaxableNil sets the value for GiftCardTaxable to be an explicit nil

### UnsetGiftCardTaxable
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetGiftCardTaxable()`

UnsetGiftCardTaxable ensures that no value is present for GiftCardTaxable, not even an explicit nil
### GetAdjustmentTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) GetAdjustmentTaxable() interface{}`

GetAdjustmentTaxable returns the AdjustmentTaxable field if non-nil, zero value otherwise.

### GetAdjustmentTaxableOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetAdjustmentTaxableOk() (*interface{}, bool)`

GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) SetAdjustmentTaxable(v interface{})`

SetAdjustmentTaxable sets AdjustmentTaxable field to given value.

### HasAdjustmentTaxable

`func (o *POSTTaxRules201ResponseDataAttributes) HasAdjustmentTaxable() bool`

HasAdjustmentTaxable returns a boolean if a field has been set.

### SetAdjustmentTaxableNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetAdjustmentTaxableNil(b bool)`

 SetAdjustmentTaxableNil sets the value for AdjustmentTaxable to be an explicit nil

### UnsetAdjustmentTaxable
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetAdjustmentTaxable()`

UnsetAdjustmentTaxable ensures that no value is present for AdjustmentTaxable, not even an explicit nil
### GetCountryCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) GetCountryCodeRegex() interface{}`

GetCountryCodeRegex returns the CountryCodeRegex field if non-nil, zero value otherwise.

### GetCountryCodeRegexOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetCountryCodeRegexOk() (*interface{}, bool)`

GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) SetCountryCodeRegex(v interface{})`

SetCountryCodeRegex sets CountryCodeRegex field to given value.

### HasCountryCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) HasCountryCodeRegex() bool`

HasCountryCodeRegex returns a boolean if a field has been set.

### SetCountryCodeRegexNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetCountryCodeRegexNil(b bool)`

 SetCountryCodeRegexNil sets the value for CountryCodeRegex to be an explicit nil

### UnsetCountryCodeRegex
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetCountryCodeRegex()`

UnsetCountryCodeRegex ensures that no value is present for CountryCodeRegex, not even an explicit nil
### GetNotCountryCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) GetNotCountryCodeRegex() interface{}`

GetNotCountryCodeRegex returns the NotCountryCodeRegex field if non-nil, zero value otherwise.

### GetNotCountryCodeRegexOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetNotCountryCodeRegexOk() (*interface{}, bool)`

GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCountryCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) SetNotCountryCodeRegex(v interface{})`

SetNotCountryCodeRegex sets NotCountryCodeRegex field to given value.

### HasNotCountryCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) HasNotCountryCodeRegex() bool`

HasNotCountryCodeRegex returns a boolean if a field has been set.

### SetNotCountryCodeRegexNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetNotCountryCodeRegexNil(b bool)`

 SetNotCountryCodeRegexNil sets the value for NotCountryCodeRegex to be an explicit nil

### UnsetNotCountryCodeRegex
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetNotCountryCodeRegex()`

UnsetNotCountryCodeRegex ensures that no value is present for NotCountryCodeRegex, not even an explicit nil
### GetStateCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) GetStateCodeRegex() interface{}`

GetStateCodeRegex returns the StateCodeRegex field if non-nil, zero value otherwise.

### GetStateCodeRegexOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetStateCodeRegexOk() (*interface{}, bool)`

GetStateCodeRegexOk returns a tuple with the StateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) SetStateCodeRegex(v interface{})`

SetStateCodeRegex sets StateCodeRegex field to given value.

### HasStateCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) HasStateCodeRegex() bool`

HasStateCodeRegex returns a boolean if a field has been set.

### SetStateCodeRegexNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetStateCodeRegexNil(b bool)`

 SetStateCodeRegexNil sets the value for StateCodeRegex to be an explicit nil

### UnsetStateCodeRegex
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetStateCodeRegex()`

UnsetStateCodeRegex ensures that no value is present for StateCodeRegex, not even an explicit nil
### GetNotStateCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) GetNotStateCodeRegex() interface{}`

GetNotStateCodeRegex returns the NotStateCodeRegex field if non-nil, zero value otherwise.

### GetNotStateCodeRegexOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetNotStateCodeRegexOk() (*interface{}, bool)`

GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStateCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) SetNotStateCodeRegex(v interface{})`

SetNotStateCodeRegex sets NotStateCodeRegex field to given value.

### HasNotStateCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) HasNotStateCodeRegex() bool`

HasNotStateCodeRegex returns a boolean if a field has been set.

### SetNotStateCodeRegexNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetNotStateCodeRegexNil(b bool)`

 SetNotStateCodeRegexNil sets the value for NotStateCodeRegex to be an explicit nil

### UnsetNotStateCodeRegex
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetNotStateCodeRegex()`

UnsetNotStateCodeRegex ensures that no value is present for NotStateCodeRegex, not even an explicit nil
### GetZipCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) GetZipCodeRegex() interface{}`

GetZipCodeRegex returns the ZipCodeRegex field if non-nil, zero value otherwise.

### GetZipCodeRegexOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetZipCodeRegexOk() (*interface{}, bool)`

GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) SetZipCodeRegex(v interface{})`

SetZipCodeRegex sets ZipCodeRegex field to given value.

### HasZipCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) HasZipCodeRegex() bool`

HasZipCodeRegex returns a boolean if a field has been set.

### SetZipCodeRegexNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetZipCodeRegexNil(b bool)`

 SetZipCodeRegexNil sets the value for ZipCodeRegex to be an explicit nil

### UnsetZipCodeRegex
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetZipCodeRegex()`

UnsetZipCodeRegex ensures that no value is present for ZipCodeRegex, not even an explicit nil
### GetNotZipCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) GetNotZipCodeRegex() interface{}`

GetNotZipCodeRegex returns the NotZipCodeRegex field if non-nil, zero value otherwise.

### GetNotZipCodeRegexOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetNotZipCodeRegexOk() (*interface{}, bool)`

GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotZipCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) SetNotZipCodeRegex(v interface{})`

SetNotZipCodeRegex sets NotZipCodeRegex field to given value.

### HasNotZipCodeRegex

`func (o *POSTTaxRules201ResponseDataAttributes) HasNotZipCodeRegex() bool`

HasNotZipCodeRegex returns a boolean if a field has been set.

### SetNotZipCodeRegexNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetNotZipCodeRegexNil(b bool)`

 SetNotZipCodeRegexNil sets the value for NotZipCodeRegex to be an explicit nil

### UnsetNotZipCodeRegex
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetNotZipCodeRegex()`

UnsetNotZipCodeRegex ensures that no value is present for NotZipCodeRegex, not even an explicit nil
### GetReference

`func (o *POSTTaxRules201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTTaxRules201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTTaxRules201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTTaxRules201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTTaxRules201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTTaxRules201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTTaxRules201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTTaxRules201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTTaxRules201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTTaxRules201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTTaxRules201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTTaxRules201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


