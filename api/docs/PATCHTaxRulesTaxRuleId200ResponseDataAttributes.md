# PATCHTaxRulesTaxRuleId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The tax rule internal name. | [optional] 
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

### NewPATCHTaxRulesTaxRuleId200ResponseDataAttributes

`func NewPATCHTaxRulesTaxRuleId200ResponseDataAttributes() *PATCHTaxRulesTaxRuleId200ResponseDataAttributes`

NewPATCHTaxRulesTaxRuleId200ResponseDataAttributes instantiates a new PATCHTaxRulesTaxRuleId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPATCHTaxRulesTaxRuleId200ResponseDataAttributesWithDefaults

`func NewPATCHTaxRulesTaxRuleId200ResponseDataAttributesWithDefaults() *PATCHTaxRulesTaxRuleId200ResponseDataAttributes`

NewPATCHTaxRulesTaxRuleId200ResponseDataAttributesWithDefaults instantiates a new PATCHTaxRulesTaxRuleId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaxRate

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetTaxRate() interface{}`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetTaxRateOk() (*interface{}, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetTaxRate(v interface{})`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetFreightTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetFreightTaxable() interface{}`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetFreightTaxableOk() (*interface{}, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetFreightTaxable(v interface{})`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### SetFreightTaxableNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetFreightTaxableNil(b bool)`

 SetFreightTaxableNil sets the value for FreightTaxable to be an explicit nil

### UnsetFreightTaxable
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetFreightTaxable()`

UnsetFreightTaxable ensures that no value is present for FreightTaxable, not even an explicit nil
### GetPaymentMethodTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetPaymentMethodTaxable() interface{}`

GetPaymentMethodTaxable returns the PaymentMethodTaxable field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetPaymentMethodTaxableOk() (*interface{}, bool)`

GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetPaymentMethodTaxable(v interface{})`

SetPaymentMethodTaxable sets PaymentMethodTaxable field to given value.

### HasPaymentMethodTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasPaymentMethodTaxable() bool`

HasPaymentMethodTaxable returns a boolean if a field has been set.

### SetPaymentMethodTaxableNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetPaymentMethodTaxableNil(b bool)`

 SetPaymentMethodTaxableNil sets the value for PaymentMethodTaxable to be an explicit nil

### UnsetPaymentMethodTaxable
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetPaymentMethodTaxable()`

UnsetPaymentMethodTaxable ensures that no value is present for PaymentMethodTaxable, not even an explicit nil
### GetGiftCardTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetGiftCardTaxable() interface{}`

GetGiftCardTaxable returns the GiftCardTaxable field if non-nil, zero value otherwise.

### GetGiftCardTaxableOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetGiftCardTaxableOk() (*interface{}, bool)`

GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetGiftCardTaxable(v interface{})`

SetGiftCardTaxable sets GiftCardTaxable field to given value.

### HasGiftCardTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasGiftCardTaxable() bool`

HasGiftCardTaxable returns a boolean if a field has been set.

### SetGiftCardTaxableNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetGiftCardTaxableNil(b bool)`

 SetGiftCardTaxableNil sets the value for GiftCardTaxable to be an explicit nil

### UnsetGiftCardTaxable
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetGiftCardTaxable()`

UnsetGiftCardTaxable ensures that no value is present for GiftCardTaxable, not even an explicit nil
### GetAdjustmentTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetAdjustmentTaxable() interface{}`

GetAdjustmentTaxable returns the AdjustmentTaxable field if non-nil, zero value otherwise.

### GetAdjustmentTaxableOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetAdjustmentTaxableOk() (*interface{}, bool)`

GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetAdjustmentTaxable(v interface{})`

SetAdjustmentTaxable sets AdjustmentTaxable field to given value.

### HasAdjustmentTaxable

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasAdjustmentTaxable() bool`

HasAdjustmentTaxable returns a boolean if a field has been set.

### SetAdjustmentTaxableNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetAdjustmentTaxableNil(b bool)`

 SetAdjustmentTaxableNil sets the value for AdjustmentTaxable to be an explicit nil

### UnsetAdjustmentTaxable
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetAdjustmentTaxable()`

UnsetAdjustmentTaxable ensures that no value is present for AdjustmentTaxable, not even an explicit nil
### GetCountryCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetCountryCodeRegex() interface{}`

GetCountryCodeRegex returns the CountryCodeRegex field if non-nil, zero value otherwise.

### GetCountryCodeRegexOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetCountryCodeRegexOk() (*interface{}, bool)`

GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetCountryCodeRegex(v interface{})`

SetCountryCodeRegex sets CountryCodeRegex field to given value.

### HasCountryCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasCountryCodeRegex() bool`

HasCountryCodeRegex returns a boolean if a field has been set.

### SetCountryCodeRegexNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetCountryCodeRegexNil(b bool)`

 SetCountryCodeRegexNil sets the value for CountryCodeRegex to be an explicit nil

### UnsetCountryCodeRegex
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetCountryCodeRegex()`

UnsetCountryCodeRegex ensures that no value is present for CountryCodeRegex, not even an explicit nil
### GetNotCountryCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetNotCountryCodeRegex() interface{}`

GetNotCountryCodeRegex returns the NotCountryCodeRegex field if non-nil, zero value otherwise.

### GetNotCountryCodeRegexOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetNotCountryCodeRegexOk() (*interface{}, bool)`

GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCountryCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetNotCountryCodeRegex(v interface{})`

SetNotCountryCodeRegex sets NotCountryCodeRegex field to given value.

### HasNotCountryCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasNotCountryCodeRegex() bool`

HasNotCountryCodeRegex returns a boolean if a field has been set.

### SetNotCountryCodeRegexNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetNotCountryCodeRegexNil(b bool)`

 SetNotCountryCodeRegexNil sets the value for NotCountryCodeRegex to be an explicit nil

### UnsetNotCountryCodeRegex
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetNotCountryCodeRegex()`

UnsetNotCountryCodeRegex ensures that no value is present for NotCountryCodeRegex, not even an explicit nil
### GetStateCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetStateCodeRegex() interface{}`

GetStateCodeRegex returns the StateCodeRegex field if non-nil, zero value otherwise.

### GetStateCodeRegexOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetStateCodeRegexOk() (*interface{}, bool)`

GetStateCodeRegexOk returns a tuple with the StateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetStateCodeRegex(v interface{})`

SetStateCodeRegex sets StateCodeRegex field to given value.

### HasStateCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasStateCodeRegex() bool`

HasStateCodeRegex returns a boolean if a field has been set.

### SetStateCodeRegexNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetStateCodeRegexNil(b bool)`

 SetStateCodeRegexNil sets the value for StateCodeRegex to be an explicit nil

### UnsetStateCodeRegex
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetStateCodeRegex()`

UnsetStateCodeRegex ensures that no value is present for StateCodeRegex, not even an explicit nil
### GetNotStateCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetNotStateCodeRegex() interface{}`

GetNotStateCodeRegex returns the NotStateCodeRegex field if non-nil, zero value otherwise.

### GetNotStateCodeRegexOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetNotStateCodeRegexOk() (*interface{}, bool)`

GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStateCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetNotStateCodeRegex(v interface{})`

SetNotStateCodeRegex sets NotStateCodeRegex field to given value.

### HasNotStateCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasNotStateCodeRegex() bool`

HasNotStateCodeRegex returns a boolean if a field has been set.

### SetNotStateCodeRegexNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetNotStateCodeRegexNil(b bool)`

 SetNotStateCodeRegexNil sets the value for NotStateCodeRegex to be an explicit nil

### UnsetNotStateCodeRegex
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetNotStateCodeRegex()`

UnsetNotStateCodeRegex ensures that no value is present for NotStateCodeRegex, not even an explicit nil
### GetZipCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetZipCodeRegex() interface{}`

GetZipCodeRegex returns the ZipCodeRegex field if non-nil, zero value otherwise.

### GetZipCodeRegexOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetZipCodeRegexOk() (*interface{}, bool)`

GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetZipCodeRegex(v interface{})`

SetZipCodeRegex sets ZipCodeRegex field to given value.

### HasZipCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasZipCodeRegex() bool`

HasZipCodeRegex returns a boolean if a field has been set.

### SetZipCodeRegexNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetZipCodeRegexNil(b bool)`

 SetZipCodeRegexNil sets the value for ZipCodeRegex to be an explicit nil

### UnsetZipCodeRegex
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetZipCodeRegex()`

UnsetZipCodeRegex ensures that no value is present for ZipCodeRegex, not even an explicit nil
### GetNotZipCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetNotZipCodeRegex() interface{}`

GetNotZipCodeRegex returns the NotZipCodeRegex field if non-nil, zero value otherwise.

### GetNotZipCodeRegexOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetNotZipCodeRegexOk() (*interface{}, bool)`

GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotZipCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetNotZipCodeRegex(v interface{})`

SetNotZipCodeRegex sets NotZipCodeRegex field to given value.

### HasNotZipCodeRegex

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasNotZipCodeRegex() bool`

HasNotZipCodeRegex returns a boolean if a field has been set.

### SetNotZipCodeRegexNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetNotZipCodeRegexNil(b bool)`

 SetNotZipCodeRegexNil sets the value for NotZipCodeRegex to be an explicit nil

### UnsetNotZipCodeRegex
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetNotZipCodeRegex()`

UnsetNotZipCodeRegex ensures that no value is present for NotZipCodeRegex, not even an explicit nil
### GetReference

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PATCHTaxRulesTaxRuleId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


