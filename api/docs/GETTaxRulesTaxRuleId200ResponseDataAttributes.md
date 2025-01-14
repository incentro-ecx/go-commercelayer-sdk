# GETTaxRulesTaxRuleId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The tax rule internal name. | [optional] 
**TaxRate** | Pointer to **interface{}** | The tax rate for this rule. | [optional] 
**FreightTaxable** | Pointer to **interface{}** | Indicates if the freight is taxable. | [optional] 
**PaymentMethodTaxable** | Pointer to **interface{}** | Indicates if the payment method is taxable. | [optional] 
**GiftCardTaxable** | Pointer to **interface{}** | Indicates if gift cards are taxable. | [optional] 
**AdjustmentTaxable** | Pointer to **interface{}** | Indicates if adjustemnts are taxable. | [optional] 
**Breakdown** | Pointer to **interface{}** | The breakdown for this tax rule (if calculated). | [optional] 
**CountryCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the shipping address country code, max size is 5000. | [optional] 
**NotCountryCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated as negative match for the shipping address country code, max size is 5000. | [optional] 
**StateCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the shipping address state code, max size is 5000. | [optional] 
**NotStateCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated as negative match for the shipping address state code, max size is 5000. | [optional] 
**ZipCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the shipping address zip code, max size is 5000. | [optional] 
**NotZipCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated as negative match for the shipping zip country code, max size is 5000. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETTaxRulesTaxRuleId200ResponseDataAttributes

`func NewGETTaxRulesTaxRuleId200ResponseDataAttributes() *GETTaxRulesTaxRuleId200ResponseDataAttributes`

NewGETTaxRulesTaxRuleId200ResponseDataAttributes instantiates a new GETTaxRulesTaxRuleId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETTaxRulesTaxRuleId200ResponseDataAttributesWithDefaults

`func NewGETTaxRulesTaxRuleId200ResponseDataAttributesWithDefaults() *GETTaxRulesTaxRuleId200ResponseDataAttributes`

NewGETTaxRulesTaxRuleId200ResponseDataAttributesWithDefaults instantiates a new GETTaxRulesTaxRuleId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTaxRate

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetTaxRate() interface{}`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetTaxRateOk() (*interface{}, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetTaxRate(v interface{})`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetFreightTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetFreightTaxable() interface{}`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetFreightTaxableOk() (*interface{}, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetFreightTaxable(v interface{})`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### SetFreightTaxableNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetFreightTaxableNil(b bool)`

 SetFreightTaxableNil sets the value for FreightTaxable to be an explicit nil

### UnsetFreightTaxable
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetFreightTaxable()`

UnsetFreightTaxable ensures that no value is present for FreightTaxable, not even an explicit nil
### GetPaymentMethodTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetPaymentMethodTaxable() interface{}`

GetPaymentMethodTaxable returns the PaymentMethodTaxable field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetPaymentMethodTaxableOk() (*interface{}, bool)`

GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetPaymentMethodTaxable(v interface{})`

SetPaymentMethodTaxable sets PaymentMethodTaxable field to given value.

### HasPaymentMethodTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasPaymentMethodTaxable() bool`

HasPaymentMethodTaxable returns a boolean if a field has been set.

### SetPaymentMethodTaxableNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetPaymentMethodTaxableNil(b bool)`

 SetPaymentMethodTaxableNil sets the value for PaymentMethodTaxable to be an explicit nil

### UnsetPaymentMethodTaxable
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetPaymentMethodTaxable()`

UnsetPaymentMethodTaxable ensures that no value is present for PaymentMethodTaxable, not even an explicit nil
### GetGiftCardTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetGiftCardTaxable() interface{}`

GetGiftCardTaxable returns the GiftCardTaxable field if non-nil, zero value otherwise.

### GetGiftCardTaxableOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetGiftCardTaxableOk() (*interface{}, bool)`

GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetGiftCardTaxable(v interface{})`

SetGiftCardTaxable sets GiftCardTaxable field to given value.

### HasGiftCardTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasGiftCardTaxable() bool`

HasGiftCardTaxable returns a boolean if a field has been set.

### SetGiftCardTaxableNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetGiftCardTaxableNil(b bool)`

 SetGiftCardTaxableNil sets the value for GiftCardTaxable to be an explicit nil

### UnsetGiftCardTaxable
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetGiftCardTaxable()`

UnsetGiftCardTaxable ensures that no value is present for GiftCardTaxable, not even an explicit nil
### GetAdjustmentTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetAdjustmentTaxable() interface{}`

GetAdjustmentTaxable returns the AdjustmentTaxable field if non-nil, zero value otherwise.

### GetAdjustmentTaxableOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetAdjustmentTaxableOk() (*interface{}, bool)`

GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetAdjustmentTaxable(v interface{})`

SetAdjustmentTaxable sets AdjustmentTaxable field to given value.

### HasAdjustmentTaxable

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasAdjustmentTaxable() bool`

HasAdjustmentTaxable returns a boolean if a field has been set.

### SetAdjustmentTaxableNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetAdjustmentTaxableNil(b bool)`

 SetAdjustmentTaxableNil sets the value for AdjustmentTaxable to be an explicit nil

### UnsetAdjustmentTaxable
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetAdjustmentTaxable()`

UnsetAdjustmentTaxable ensures that no value is present for AdjustmentTaxable, not even an explicit nil
### GetBreakdown

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetBreakdown() interface{}`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetBreakdownOk() (*interface{}, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetBreakdown(v interface{})`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### SetBreakdownNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetBreakdownNil(b bool)`

 SetBreakdownNil sets the value for Breakdown to be an explicit nil

### UnsetBreakdown
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetBreakdown()`

UnsetBreakdown ensures that no value is present for Breakdown, not even an explicit nil
### GetCountryCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetCountryCodeRegex() interface{}`

GetCountryCodeRegex returns the CountryCodeRegex field if non-nil, zero value otherwise.

### GetCountryCodeRegexOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetCountryCodeRegexOk() (*interface{}, bool)`

GetCountryCodeRegexOk returns a tuple with the CountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetCountryCodeRegex(v interface{})`

SetCountryCodeRegex sets CountryCodeRegex field to given value.

### HasCountryCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasCountryCodeRegex() bool`

HasCountryCodeRegex returns a boolean if a field has been set.

### SetCountryCodeRegexNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetCountryCodeRegexNil(b bool)`

 SetCountryCodeRegexNil sets the value for CountryCodeRegex to be an explicit nil

### UnsetCountryCodeRegex
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetCountryCodeRegex()`

UnsetCountryCodeRegex ensures that no value is present for CountryCodeRegex, not even an explicit nil
### GetNotCountryCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetNotCountryCodeRegex() interface{}`

GetNotCountryCodeRegex returns the NotCountryCodeRegex field if non-nil, zero value otherwise.

### GetNotCountryCodeRegexOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetNotCountryCodeRegexOk() (*interface{}, bool)`

GetNotCountryCodeRegexOk returns a tuple with the NotCountryCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCountryCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetNotCountryCodeRegex(v interface{})`

SetNotCountryCodeRegex sets NotCountryCodeRegex field to given value.

### HasNotCountryCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasNotCountryCodeRegex() bool`

HasNotCountryCodeRegex returns a boolean if a field has been set.

### SetNotCountryCodeRegexNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetNotCountryCodeRegexNil(b bool)`

 SetNotCountryCodeRegexNil sets the value for NotCountryCodeRegex to be an explicit nil

### UnsetNotCountryCodeRegex
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetNotCountryCodeRegex()`

UnsetNotCountryCodeRegex ensures that no value is present for NotCountryCodeRegex, not even an explicit nil
### GetStateCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetStateCodeRegex() interface{}`

GetStateCodeRegex returns the StateCodeRegex field if non-nil, zero value otherwise.

### GetStateCodeRegexOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetStateCodeRegexOk() (*interface{}, bool)`

GetStateCodeRegexOk returns a tuple with the StateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetStateCodeRegex(v interface{})`

SetStateCodeRegex sets StateCodeRegex field to given value.

### HasStateCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasStateCodeRegex() bool`

HasStateCodeRegex returns a boolean if a field has been set.

### SetStateCodeRegexNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetStateCodeRegexNil(b bool)`

 SetStateCodeRegexNil sets the value for StateCodeRegex to be an explicit nil

### UnsetStateCodeRegex
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetStateCodeRegex()`

UnsetStateCodeRegex ensures that no value is present for StateCodeRegex, not even an explicit nil
### GetNotStateCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetNotStateCodeRegex() interface{}`

GetNotStateCodeRegex returns the NotStateCodeRegex field if non-nil, zero value otherwise.

### GetNotStateCodeRegexOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetNotStateCodeRegexOk() (*interface{}, bool)`

GetNotStateCodeRegexOk returns a tuple with the NotStateCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotStateCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetNotStateCodeRegex(v interface{})`

SetNotStateCodeRegex sets NotStateCodeRegex field to given value.

### HasNotStateCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasNotStateCodeRegex() bool`

HasNotStateCodeRegex returns a boolean if a field has been set.

### SetNotStateCodeRegexNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetNotStateCodeRegexNil(b bool)`

 SetNotStateCodeRegexNil sets the value for NotStateCodeRegex to be an explicit nil

### UnsetNotStateCodeRegex
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetNotStateCodeRegex()`

UnsetNotStateCodeRegex ensures that no value is present for NotStateCodeRegex, not even an explicit nil
### GetZipCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetZipCodeRegex() interface{}`

GetZipCodeRegex returns the ZipCodeRegex field if non-nil, zero value otherwise.

### GetZipCodeRegexOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetZipCodeRegexOk() (*interface{}, bool)`

GetZipCodeRegexOk returns a tuple with the ZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetZipCodeRegex(v interface{})`

SetZipCodeRegex sets ZipCodeRegex field to given value.

### HasZipCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasZipCodeRegex() bool`

HasZipCodeRegex returns a boolean if a field has been set.

### SetZipCodeRegexNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetZipCodeRegexNil(b bool)`

 SetZipCodeRegexNil sets the value for ZipCodeRegex to be an explicit nil

### UnsetZipCodeRegex
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetZipCodeRegex()`

UnsetZipCodeRegex ensures that no value is present for ZipCodeRegex, not even an explicit nil
### GetNotZipCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetNotZipCodeRegex() interface{}`

GetNotZipCodeRegex returns the NotZipCodeRegex field if non-nil, zero value otherwise.

### GetNotZipCodeRegexOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetNotZipCodeRegexOk() (*interface{}, bool)`

GetNotZipCodeRegexOk returns a tuple with the NotZipCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotZipCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetNotZipCodeRegex(v interface{})`

SetNotZipCodeRegex sets NotZipCodeRegex field to given value.

### HasNotZipCodeRegex

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasNotZipCodeRegex() bool`

HasNotZipCodeRegex returns a boolean if a field has been set.

### SetNotZipCodeRegexNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetNotZipCodeRegexNil(b bool)`

 SetNotZipCodeRegexNil sets the value for NotZipCodeRegex to be an explicit nil

### UnsetNotZipCodeRegex
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetNotZipCodeRegex()`

UnsetNotZipCodeRegex ensures that no value is present for NotZipCodeRegex, not even an explicit nil
### GetCreatedAt

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETTaxRulesTaxRuleId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


