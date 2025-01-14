# GETPriceListsPriceListId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The price list&#39;s internal name. | [optional] 
**Code** | Pointer to **interface{}** | A string that you can use to identify the price list (must be unique within the environment). | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard. | [optional] 
**TaxIncluded** | Pointer to **interface{}** | Indicates if the associated prices include taxes. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Rules** | Pointer to **interface{}** | The rules (using Rules Engine) to be applied. | [optional] 
**RuleOutcomes** | Pointer to **interface{}** | The rule outcomes. | [optional] 
**ResourcePayload** | Pointer to **interface{}** | The payload used to evaluate the rules. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPriceListsPriceListId200ResponseDataAttributes

`func NewGETPriceListsPriceListId200ResponseDataAttributes() *GETPriceListsPriceListId200ResponseDataAttributes`

NewGETPriceListsPriceListId200ResponseDataAttributes instantiates a new GETPriceListsPriceListId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPriceListsPriceListId200ResponseDataAttributesWithDefaults

`func NewGETPriceListsPriceListId200ResponseDataAttributesWithDefaults() *GETPriceListsPriceListId200ResponseDataAttributes`

NewGETPriceListsPriceListId200ResponseDataAttributesWithDefaults instantiates a new GETPriceListsPriceListId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCode

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCurrencyCode

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetTaxIncluded

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetTaxIncluded() interface{}`

GetTaxIncluded returns the TaxIncluded field if non-nil, zero value otherwise.

### GetTaxIncludedOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetTaxIncludedOk() (*interface{}, bool)`

GetTaxIncludedOk returns a tuple with the TaxIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIncluded

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetTaxIncluded(v interface{})`

SetTaxIncluded sets TaxIncluded field to given value.

### HasTaxIncluded

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasTaxIncluded() bool`

HasTaxIncluded returns a boolean if a field has been set.

### SetTaxIncludedNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetTaxIncludedNil(b bool)`

 SetTaxIncludedNil sets the value for TaxIncluded to be an explicit nil

### UnsetTaxIncluded
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetTaxIncluded()`

UnsetTaxIncluded ensures that no value is present for TaxIncluded, not even an explicit nil
### GetCreatedAt

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetRules

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetRules() interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetRulesOk() (*interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetRules(v interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetRuleOutcomes

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetRuleOutcomes() interface{}`

GetRuleOutcomes returns the RuleOutcomes field if non-nil, zero value otherwise.

### GetRuleOutcomesOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetRuleOutcomesOk() (*interface{}, bool)`

GetRuleOutcomesOk returns a tuple with the RuleOutcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOutcomes

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetRuleOutcomes(v interface{})`

SetRuleOutcomes sets RuleOutcomes field to given value.

### HasRuleOutcomes

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasRuleOutcomes() bool`

HasRuleOutcomes returns a boolean if a field has been set.

### SetRuleOutcomesNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetRuleOutcomesNil(b bool)`

 SetRuleOutcomesNil sets the value for RuleOutcomes to be an explicit nil

### UnsetRuleOutcomes
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetRuleOutcomes()`

UnsetRuleOutcomes ensures that no value is present for RuleOutcomes, not even an explicit nil
### GetResourcePayload

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetResourcePayload() interface{}`

GetResourcePayload returns the ResourcePayload field if non-nil, zero value otherwise.

### GetResourcePayloadOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetResourcePayloadOk() (*interface{}, bool)`

GetResourcePayloadOk returns a tuple with the ResourcePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePayload

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetResourcePayload(v interface{})`

SetResourcePayload sets ResourcePayload field to given value.

### HasResourcePayload

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasResourcePayload() bool`

HasResourcePayload returns a boolean if a field has been set.

### SetResourcePayloadNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetResourcePayloadNil(b bool)`

 SetResourcePayloadNil sets the value for ResourcePayload to be an explicit nil

### UnsetResourcePayload
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetResourcePayload()`

UnsetResourcePayload ensures that no value is present for ResourcePayload, not even an explicit nil
### GetMetadata

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPriceListsPriceListId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


