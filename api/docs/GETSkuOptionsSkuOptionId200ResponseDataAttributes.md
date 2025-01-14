# GETSkuOptionsSkuOptionId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The SKU option&#39;s internal name. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**Description** | Pointer to **interface{}** | An internal description of the SKU option. | [optional] 
**PriceAmountCents** | Pointer to **interface{}** | The price of this shipping method, in cents. | [optional] 
**PriceAmountFloat** | Pointer to **interface{}** | The price of this shipping method, float. | [optional] 
**FormattedPriceAmount** | Pointer to **interface{}** | The price of this shipping method, formatted. | [optional] 
**DelayHours** | Pointer to **interface{}** | The delay time (in hours) that should be added to the delivery lead time when this option is purchased. | [optional] 
**DelayDays** | Pointer to **interface{}** | The delay time, in days (rounded). | [optional] 
**SkuCodeRegex** | Pointer to **interface{}** | The regex that will be evaluated to match the SKU codes, max size is 5000. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETSkuOptionsSkuOptionId200ResponseDataAttributes

`func NewGETSkuOptionsSkuOptionId200ResponseDataAttributes() *GETSkuOptionsSkuOptionId200ResponseDataAttributes`

NewGETSkuOptionsSkuOptionId200ResponseDataAttributes instantiates a new GETSkuOptionsSkuOptionId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETSkuOptionsSkuOptionId200ResponseDataAttributesWithDefaults

`func NewGETSkuOptionsSkuOptionId200ResponseDataAttributesWithDefaults() *GETSkuOptionsSkuOptionId200ResponseDataAttributes`

NewGETSkuOptionsSkuOptionId200ResponseDataAttributesWithDefaults instantiates a new GETSkuOptionsSkuOptionId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCurrencyCode

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetDescription

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetDescription() interface{}`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetDescriptionOk() (*interface{}, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetDescription(v interface{})`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceAmountCents

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetPriceAmountCents() interface{}`

GetPriceAmountCents returns the PriceAmountCents field if non-nil, zero value otherwise.

### GetPriceAmountCentsOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetPriceAmountCentsOk() (*interface{}, bool)`

GetPriceAmountCentsOk returns a tuple with the PriceAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountCents

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetPriceAmountCents(v interface{})`

SetPriceAmountCents sets PriceAmountCents field to given value.

### HasPriceAmountCents

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasPriceAmountCents() bool`

HasPriceAmountCents returns a boolean if a field has been set.

### SetPriceAmountCentsNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetPriceAmountCentsNil(b bool)`

 SetPriceAmountCentsNil sets the value for PriceAmountCents to be an explicit nil

### UnsetPriceAmountCents
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetPriceAmountCents()`

UnsetPriceAmountCents ensures that no value is present for PriceAmountCents, not even an explicit nil
### GetPriceAmountFloat

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetPriceAmountFloat() interface{}`

GetPriceAmountFloat returns the PriceAmountFloat field if non-nil, zero value otherwise.

### GetPriceAmountFloatOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetPriceAmountFloatOk() (*interface{}, bool)`

GetPriceAmountFloatOk returns a tuple with the PriceAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceAmountFloat

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetPriceAmountFloat(v interface{})`

SetPriceAmountFloat sets PriceAmountFloat field to given value.

### HasPriceAmountFloat

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasPriceAmountFloat() bool`

HasPriceAmountFloat returns a boolean if a field has been set.

### SetPriceAmountFloatNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetPriceAmountFloatNil(b bool)`

 SetPriceAmountFloatNil sets the value for PriceAmountFloat to be an explicit nil

### UnsetPriceAmountFloat
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetPriceAmountFloat()`

UnsetPriceAmountFloat ensures that no value is present for PriceAmountFloat, not even an explicit nil
### GetFormattedPriceAmount

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetFormattedPriceAmount() interface{}`

GetFormattedPriceAmount returns the FormattedPriceAmount field if non-nil, zero value otherwise.

### GetFormattedPriceAmountOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetFormattedPriceAmountOk() (*interface{}, bool)`

GetFormattedPriceAmountOk returns a tuple with the FormattedPriceAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedPriceAmount

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetFormattedPriceAmount(v interface{})`

SetFormattedPriceAmount sets FormattedPriceAmount field to given value.

### HasFormattedPriceAmount

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasFormattedPriceAmount() bool`

HasFormattedPriceAmount returns a boolean if a field has been set.

### SetFormattedPriceAmountNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetFormattedPriceAmountNil(b bool)`

 SetFormattedPriceAmountNil sets the value for FormattedPriceAmount to be an explicit nil

### UnsetFormattedPriceAmount
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetFormattedPriceAmount()`

UnsetFormattedPriceAmount ensures that no value is present for FormattedPriceAmount, not even an explicit nil
### GetDelayHours

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetDelayHours() interface{}`

GetDelayHours returns the DelayHours field if non-nil, zero value otherwise.

### GetDelayHoursOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetDelayHoursOk() (*interface{}, bool)`

GetDelayHoursOk returns a tuple with the DelayHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayHours

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetDelayHours(v interface{})`

SetDelayHours sets DelayHours field to given value.

### HasDelayHours

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasDelayHours() bool`

HasDelayHours returns a boolean if a field has been set.

### SetDelayHoursNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetDelayHoursNil(b bool)`

 SetDelayHoursNil sets the value for DelayHours to be an explicit nil

### UnsetDelayHours
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetDelayHours()`

UnsetDelayHours ensures that no value is present for DelayHours, not even an explicit nil
### GetDelayDays

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetDelayDays() interface{}`

GetDelayDays returns the DelayDays field if non-nil, zero value otherwise.

### GetDelayDaysOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetDelayDaysOk() (*interface{}, bool)`

GetDelayDaysOk returns a tuple with the DelayDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDays

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetDelayDays(v interface{})`

SetDelayDays sets DelayDays field to given value.

### HasDelayDays

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasDelayDays() bool`

HasDelayDays returns a boolean if a field has been set.

### SetDelayDaysNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetDelayDaysNil(b bool)`

 SetDelayDaysNil sets the value for DelayDays to be an explicit nil

### UnsetDelayDays
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetDelayDays()`

UnsetDelayDays ensures that no value is present for DelayDays, not even an explicit nil
### GetSkuCodeRegex

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetSkuCodeRegex() interface{}`

GetSkuCodeRegex returns the SkuCodeRegex field if non-nil, zero value otherwise.

### GetSkuCodeRegexOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetSkuCodeRegexOk() (*interface{}, bool)`

GetSkuCodeRegexOk returns a tuple with the SkuCodeRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCodeRegex

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetSkuCodeRegex(v interface{})`

SetSkuCodeRegex sets SkuCodeRegex field to given value.

### HasSkuCodeRegex

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasSkuCodeRegex() bool`

HasSkuCodeRegex returns a boolean if a field has been set.

### SetSkuCodeRegexNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetSkuCodeRegexNil(b bool)`

 SetSkuCodeRegexNil sets the value for SkuCodeRegex to be an explicit nil

### UnsetSkuCodeRegex
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetSkuCodeRegex()`

UnsetSkuCodeRegex ensures that no value is present for SkuCodeRegex, not even an explicit nil
### GetCreatedAt

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETSkuOptionsSkuOptionId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


