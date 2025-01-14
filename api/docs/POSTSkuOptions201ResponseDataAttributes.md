# POSTSkuOptions201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **interface{}** | The SKU option&#39;s internal name. | 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Description** | Pointer to **interface{}** | An internal description of the SKU option. | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The price of this shipping method, in cents. | [optional] 
**DelayHours** | Pointer to **interface{}** | The delay time (in hours) that should be added to the delivery lead time when this option is purchased. | [optional] 
**SkuCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the SKU codes, max size is 5000. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTSkuOptions201ResponseDataAttributes

`func NewPOSTSkuOptions201ResponseDataAttributes(name interface{}, ) *POSTSkuOptions201ResponseDataAttributes`

NewPOSTSkuOptions201ResponseDataAttributes instantiates a new POSTSkuOptions201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTSkuOptions201ResponseDataAttributesWithDefaults

`func NewPOSTSkuOptions201ResponseDataAttributesWithDefaults() *POSTSkuOptions201ResponseDataAttributes`

NewPOSTSkuOptions201ResponseDataAttributesWithDefaults instantiates a new POSTSkuOptions201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *POSTSkuOptions201ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *POSTSkuOptions201ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *POSTSkuOptions201ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.


### SetNameNil

`func (o *POSTSkuOptions201ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *POSTSkuOptions201ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *POSTSkuOptions201ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *POSTSkuOptions201ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *POSTSkuOptions201ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *POSTSkuOptions201ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *POSTSkuOptions201ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *POSTSkuOptions201ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDescription

`func (o *POSTSkuOptions201ResponseDataAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *POSTSkuOptions201ResponseDataAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *POSTSkuOptions201ResponseDataAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *POSTSkuOptions201ResponseDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *POSTSkuOptions201ResponseDataAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *POSTSkuOptions201ResponseDataAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceAmountCents

`func (o *POSTSkuOptions201ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *POSTSkuOptions201ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *POSTSkuOptions201ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *POSTSkuOptions201ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *POSTSkuOptions201ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *POSTSkuOptions201ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetDelayHours

`func (o *POSTSkuOptions201ResponseDataAttributes) GetDelayHours() interface{}`

GetDelayHours returns the DelayHours field if non-nil, zero value otherwise.

### GetDelayHoursOk

`func (o *POSTSkuOptions201ResponseDataAttributes) GetDelayHoursOk() (*interface{}, bool)`

GetDelayHoursOk returns a tuple with the DelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayHours

`func (o *POSTSkuOptions201ResponseDataAttributes) SetDelayHours(v interface{})`

SetDelayHours sets DelayHours field to given value.

### HasDelayHours

`func (o *POSTSkuOptions201ResponseDataAttributes) HasDelayHours() bool`

HasDelayHours returns a boolean if a field has been set.

### SetDelayHoursNil

`func (o *POSTSkuOptions201ResponseDataAttributes) SetDelayHoursNil(b bool)`

 SetDelayHoursNil sets the value for DelayHours to be an explicit nil

### UnsetDelayHours
`func (o *POSTSkuOptions201ResponseDataAttributes) UnsetDelayHours()`

UnsetDelayHours ensures that no value is present for DelayHours, not even an explicit nil
### GetSkuCodeRegex

`func (o *POSTSkuOptions201ResponseDataAttributes) GetSkuCodeRegex() interface{}`

GetSkuCodeRegex returns the SkuCodeRegex field if non-nil, zero value otherwise.

### GetSkuCodeRegexOk

`func (o *POSTSkuOptions201ResponseDataAttributes) GetSkuCodeRegexOk() (*interface{}, bool)`

GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCodeRegex

`func (o *POSTSkuOptions201ResponseDataAttributes) SetSkuCodeRegex(v interface{})`

SetSkuCodeRegex sets SkuCodeRegex field to given value.

### HasSkuCodeRegex

`func (o *POSTSkuOptions201ResponseDataAttributes) HasSkuCodeRegex() bool`

HasSkuCodeRegex returns a boolean if a field has been set.

### SetSkuCodeRegexNil

`func (o *POSTSkuOptions201ResponseDataAttributes) SetSkuCodeRegexNil(b bool)`

 SetSkuCodeRegexNil sets the value for SkuCodeRegex to be an explicit nil

### UnsetSkuCodeRegex
`func (o *POSTSkuOptions201ResponseDataAttributes) UnsetSkuCodeRegex()`

UnsetSkuCodeRegex ensures that no value is present for SkuCodeRegex, not even an explicit nil
### GetReference

`func (o *POSTSkuOptions201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTSkuOptions201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTSkuOptions201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTSkuOptions201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTSkuOptions201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTSkuOptions201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTSkuOptions201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTSkuOptions201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTSkuOptions201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTSkuOptions201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTSkuOptions201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTSkuOptions201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTSkuOptions201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTSkuOptions201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTSkuOptions201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTSkuOptions201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTSkuOptions201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTSkuOptions201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


