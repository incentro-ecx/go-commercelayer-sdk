# GETPricesPriceId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, inherited from the associated price list. | [optional] 
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. When creating a price, either a valid sku_code or a SKU relationship must be present. | [optional] 
**AmountCents** | Pointer to **interface{}** | The SKU price amount for the associated price list, in cents. | [optional] 
**AmountFloat** | Pointer to **interface{}** | The SKU price amount for the associated price list, float. | [optional] 
**FormattedAmount** | Pointer to **interface{}** | The SKU price amount for the associated price list, formatted. | [optional] 
**OriginalAmountCents** | Pointer to **interface{}** | The SKU price amount for the associated price list, in cents before any applied rule. | [optional] 
**FormattedOriginalAmount** | Pointer to **interface{}** | The SKU price amount for the associated price list, in cents before any applied rule, formatted. | [optional] 
**CompareAtAmountCents** | Pointer to **interface{}** | The compared price amount, in cents. Useful to display a percentage discount. | [optional] 
**CompareAtAmountFloat** | Pointer to **interface{}** | The compared price amount, float. | [optional] 
**FormattedCompareAtAmount** | Pointer to **interface{}** | The compared price amount, formatted. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Rules** | Pointer to **interface{}** | The rules (using Rules Engine) to be applied. | [optional] 
**RuleOutcomes** | Pointer to **interface{}** | The rule outcomes. | [optional] 
**ResourcePayload** | Pointer to **interface{}** | The payload used to evaluate the rules. | [optional] 
**JwtCustomClaim** | Pointer to **interface{}** | The custom_claim attached to the current JWT (if any). | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETPricesPriceId200ResponseDataAttributes

`func NewGETPricesPriceId200ResponseDataAttributes() *GETPricesPriceId200ResponseDataAttributes`

NewGETPricesPriceId200ResponseDataAttributes instantiates a new GETPricesPriceId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETPricesPriceId200ResponseDataAttributesWithDefaults

`func NewGETPricesPriceId200ResponseDataAttributesWithDefaults() *GETPricesPriceId200ResponseDataAttributes`

NewGETPricesPriceId200ResponseDataAttributesWithDefaults instantiates a new GETPricesPriceId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *GETPricesPriceId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETPricesPriceId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETPricesPriceId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetSkuCode

`func (o *GETPricesPriceId200ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *GETPricesPriceId200ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *GETPricesPriceId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetAmountCents

`func (o *GETPricesPriceId200ResponseDataAttributes) GetAmountCents() interface{}`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetAmountCentsOk() (*interface{}, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GETPricesPriceId200ResponseDataAttributes) SetAmountCents(v interface{})`

SetAmountCents sets AmountCents field to given value.

### HasAmountCents

`func (o *GETPricesPriceId200ResponseDataAttributes) HasAmountCents() bool`

HasAmountCents returns a boolean if a field has been set.

### SetAmountCentsNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetAmountCentsNil(b bool)`

 SetAmountCentsNil sets the value for AmountCents to be an explicit nil

### UnsetAmountCents
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetAmountCents()`

UnsetAmountCents ensures that no value is present for AmountCents, not even an explicit nil
### GetAmountFloat

`func (o *GETPricesPriceId200ResponseDataAttributes) GetAmountFloat() interface{}`

GetAmountFloat returns the AmountFloat field if non-nil, zero value otherwise.

### GetAmountFloatOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetAmountFloatOk() (*interface{}, bool)`

GetAmountFloatOk returns a tuple with the AmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountFloat

`func (o *GETPricesPriceId200ResponseDataAttributes) SetAmountFloat(v interface{})`

SetAmountFloat sets AmountFloat field to given value.

### HasAmountFloat

`func (o *GETPricesPriceId200ResponseDataAttributes) HasAmountFloat() bool`

HasAmountFloat returns a boolean if a field has been set.

### SetAmountFloatNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetAmountFloatNil(b bool)`

 SetAmountFloatNil sets the value for AmountFloat to be an explicit nil

### UnsetAmountFloat
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetAmountFloat()`

UnsetAmountFloat ensures that no value is present for AmountFloat, not even an explicit nil
### GetFormattedAmount

`func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedAmount() interface{}`

GetFormattedAmount returns the FormattedAmount field if non-nil, zero value otherwise.

### GetFormattedAmountOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedAmountOk() (*interface{}, bool)`

GetFormattedAmountOk returns a tuple with the FormattedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedAmount

`func (o *GETPricesPriceId200ResponseDataAttributes) SetFormattedAmount(v interface{})`

SetFormattedAmount sets FormattedAmount field to given value.

### HasFormattedAmount

`func (o *GETPricesPriceId200ResponseDataAttributes) HasFormattedAmount() bool`

HasFormattedAmount returns a boolean if a field has been set.

### SetFormattedAmountNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetFormattedAmountNil(b bool)`

 SetFormattedAmountNil sets the value for FormattedAmount to be an explicit nil

### UnsetFormattedAmount
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetFormattedAmount()`

UnsetFormattedAmount ensures that no value is present for FormattedAmount, not even an explicit nil
### GetOriginalAmountCents

`func (o *GETPricesPriceId200ResponseDataAttributes) GetOriginalAmountCents() interface{}`

GetOriginalAmountCents returns the OriginalAmountCents field if non-nil, zero value otherwise.

### GetOriginalAmountCentsOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetOriginalAmountCentsOk() (*interface{}, bool)`

GetOriginalAmountCentsOk returns a tuple with the OriginalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmountCents

`func (o *GETPricesPriceId200ResponseDataAttributes) SetOriginalAmountCents(v interface{})`

SetOriginalAmountCents sets OriginalAmountCents field to given value.

### HasOriginalAmountCents

`func (o *GETPricesPriceId200ResponseDataAttributes) HasOriginalAmountCents() bool`

HasOriginalAmountCents returns a boolean if a field has been set.

### SetOriginalAmountCentsNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetOriginalAmountCentsNil(b bool)`

 SetOriginalAmountCentsNil sets the value for OriginalAmountCents to be an explicit nil

### UnsetOriginalAmountCents
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetOriginalAmountCents()`

UnsetOriginalAmountCents ensures that no value is present for OriginalAmountCents, not even an explicit nil
### GetFormattedOriginalAmount

`func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedOriginalAmount() interface{}`

GetFormattedOriginalAmount returns the FormattedOriginalAmount field if non-nil, zero value otherwise.

### GetFormattedOriginalAmountOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedOriginalAmountOk() (*interface{}, bool)`

GetFormattedOriginalAmountOk returns a tuple with the FormattedOriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedOriginalAmount

`func (o *GETPricesPriceId200ResponseDataAttributes) SetFormattedOriginalAmount(v interface{})`

SetFormattedOriginalAmount sets FormattedOriginalAmount field to given value.

### HasFormattedOriginalAmount

`func (o *GETPricesPriceId200ResponseDataAttributes) HasFormattedOriginalAmount() bool`

HasFormattedOriginalAmount returns a boolean if a field has been set.

### SetFormattedOriginalAmountNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetFormattedOriginalAmountNil(b bool)`

 SetFormattedOriginalAmountNil sets the value for FormattedOriginalAmount to be an explicit nil

### UnsetFormattedOriginalAmount
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetFormattedOriginalAmount()`

UnsetFormattedOriginalAmount ensures that no value is present for FormattedOriginalAmount, not even an explicit nil
### GetCompareAtAmountCents

`func (o *GETPricesPriceId200ResponseDataAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *GETPricesPriceId200ResponseDataAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.

### HasCompareAtAmountCents

`func (o *GETPricesPriceId200ResponseDataAttributes) HasCompareAtAmountCents() bool`

HasCompareAtAmountCents returns a boolean if a field has been set.

### SetCompareAtAmountCentsNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetCompareAtAmountFloat

`func (o *GETPricesPriceId200ResponseDataAttributes) GetCompareAtAmountFloat() interface{}`

GetCompareAtAmountFloat returns the CompareAtAmountFloat field if non-nil, zero value otherwise.

### GetCompareAtAmountFloatOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetCompareAtAmountFloatOk() (*interface{}, bool)`

GetCompareAtAmountFloatOk returns a tuple with the CompareAtAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountFloat

`func (o *GETPricesPriceId200ResponseDataAttributes) SetCompareAtAmountFloat(v interface{})`

SetCompareAtAmountFloat sets CompareAtAmountFloat field to given value.

### HasCompareAtAmountFloat

`func (o *GETPricesPriceId200ResponseDataAttributes) HasCompareAtAmountFloat() bool`

HasCompareAtAmountFloat returns a boolean if a field has been set.

### SetCompareAtAmountFloatNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetCompareAtAmountFloatNil(b bool)`

 SetCompareAtAmountFloatNil sets the value for CompareAtAmountFloat to be an explicit nil

### UnsetCompareAtAmountFloat
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetCompareAtAmountFloat()`

UnsetCompareAtAmountFloat ensures that no value is present for CompareAtAmountFloat, not even an explicit nil
### GetFormattedCompareAtAmount

`func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedCompareAtAmount() interface{}`

GetFormattedCompareAtAmount returns the FormattedCompareAtAmount field if non-nil, zero value otherwise.

### GetFormattedCompareAtAmountOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetFormattedCompareAtAmountOk() (*interface{}, bool)`

GetFormattedCompareAtAmountOk returns a tuple with the FormattedCompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCompareAtAmount

`func (o *GETPricesPriceId200ResponseDataAttributes) SetFormattedCompareAtAmount(v interface{})`

SetFormattedCompareAtAmount sets FormattedCompareAtAmount field to given value.

### HasFormattedCompareAtAmount

`func (o *GETPricesPriceId200ResponseDataAttributes) HasFormattedCompareAtAmount() bool`

HasFormattedCompareAtAmount returns a boolean if a field has been set.

### SetFormattedCompareAtAmountNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetFormattedCompareAtAmountNil(b bool)`

 SetFormattedCompareAtAmountNil sets the value for FormattedCompareAtAmount to be an explicit nil

### UnsetFormattedCompareAtAmount
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetFormattedCompareAtAmount()`

UnsetFormattedCompareAtAmount ensures that no value is present for FormattedCompareAtAmount, not even an explicit nil
### GetCreatedAt

`func (o *GETPricesPriceId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETPricesPriceId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETPricesPriceId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETPricesPriceId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETPricesPriceId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETPricesPriceId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETPricesPriceId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETPricesPriceId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETPricesPriceId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETPricesPriceId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETPricesPriceId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETPricesPriceId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetRules

`func (o *GETPricesPriceId200ResponseDataAttributes) GetRules() interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetRulesOk() (*interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GETPricesPriceId200ResponseDataAttributes) SetRules(v interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *GETPricesPriceId200ResponseDataAttributes) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetRuleOutcomes

`func (o *GETPricesPriceId200ResponseDataAttributes) GetRuleOutcomes() interface{}`

GetRuleOutcomes returns the RuleOutcomes field if non-nil, zero value otherwise.

### GetRuleOutcomesOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetRuleOutcomesOk() (*interface{}, bool)`

GetRuleOutcomesOk returns a tuple with the RuleOutcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOutcomes

`func (o *GETPricesPriceId200ResponseDataAttributes) SetRuleOutcomes(v interface{})`

SetRuleOutcomes sets RuleOutcomes field to given value.

### HasRuleOutcomes

`func (o *GETPricesPriceId200ResponseDataAttributes) HasRuleOutcomes() bool`

HasRuleOutcomes returns a boolean if a field has been set.

### SetRuleOutcomesNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetRuleOutcomesNil(b bool)`

 SetRuleOutcomesNil sets the value for RuleOutcomes to be an explicit nil

### UnsetRuleOutcomes
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetRuleOutcomes()`

UnsetRuleOutcomes ensures that no value is present for RuleOutcomes, not even an explicit nil
### GetResourcePayload

`func (o *GETPricesPriceId200ResponseDataAttributes) GetResourcePayload() interface{}`

GetResourcePayload returns the ResourcePayload field if non-nil, zero value otherwise.

### GetResourcePayloadOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetResourcePayloadOk() (*interface{}, bool)`

GetResourcePayloadOk returns a tuple with the ResourcePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePayload

`func (o *GETPricesPriceId200ResponseDataAttributes) SetResourcePayload(v interface{})`

SetResourcePayload sets ResourcePayload field to given value.

### HasResourcePayload

`func (o *GETPricesPriceId200ResponseDataAttributes) HasResourcePayload() bool`

HasResourcePayload returns a boolean if a field has been set.

### SetResourcePayloadNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetResourcePayloadNil(b bool)`

 SetResourcePayloadNil sets the value for ResourcePayload to be an explicit nil

### UnsetResourcePayload
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetResourcePayload()`

UnsetResourcePayload ensures that no value is present for ResourcePayload, not even an explicit nil
### GetJwtCustomClaim

`func (o *GETPricesPriceId200ResponseDataAttributes) GetJwtCustomClaim() interface{}`

GetJwtCustomClaim returns the JwtCustomClaim field if non-nil, zero value otherwise.

### GetJwtCustomClaimOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetJwtCustomClaimOk() (*interface{}, bool)`

GetJwtCustomClaimOk returns a tuple with the JwtCustomClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtCustomClaim

`func (o *GETPricesPriceId200ResponseDataAttributes) SetJwtCustomClaim(v interface{})`

SetJwtCustomClaim sets JwtCustomClaim field to given value.

### HasJwtCustomClaim

`func (o *GETPricesPriceId200ResponseDataAttributes) HasJwtCustomClaim() bool`

HasJwtCustomClaim returns a boolean if a field has been set.

### SetJwtCustomClaimNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetJwtCustomClaimNil(b bool)`

 SetJwtCustomClaimNil sets the value for JwtCustomClaim to be an explicit nil

### UnsetJwtCustomClaim
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetJwtCustomClaim()`

UnsetJwtCustomClaim ensures that no value is present for JwtCustomClaim, not even an explicit nil
### GetMetadata

`func (o *GETPricesPriceId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETPricesPriceId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETPricesPriceId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETPricesPriceId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETPricesPriceId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETPricesPriceId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


