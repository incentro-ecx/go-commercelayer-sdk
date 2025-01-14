# GETFlexPromotionsFlexPromotionId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The promotion&#39;s internal name. | [optional] 
**Type** | Pointer to **interface{}** | The promotion&#39;s type. | [optional] 
**Exclusive** | Pointer to **interface{}** | Indicates if the promotion will be applied exclusively, based on its priority score. | [optional] 
**Priority** | Pointer to **interface{}** | The priority assigned to the promotion (lower means higher priority). | [optional] 
**StartsAt** | Pointer to **interface{}** | The activation date/time of this promotion. | [optional] 
**ExpiresAt** | Pointer to **interface{}** | The expiration date/time of this promotion (must be after starts_at). | [optional] 
**TotalUsageLimit** | Pointer to **interface{}** | The total number of times this promotion can be applied. When &#39;null&#39; it means promotion can be applied infinite times. | [optional] 
**TotalUsageCount** | Pointer to **interface{}** | The number of times this promotion has been applied. | [optional] 
**Active** | Pointer to **interface{}** | Indicates if the promotion is active (enabled and not expired). | [optional] 
**Status** | Pointer to **interface{}** | The promotion status. One of &#39;disabled&#39;, &#39;expired&#39;, &#39;pending&#39;, &#39;active&#39;, or &#39;inactive&#39;. | [optional] 
**Rules** | Pointer to **interface{}** | The discount rule to be applied. | [optional] 
**DisabledAt** | Pointer to **interface{}** | Time at which this resource was disabled. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**RuleOutcomes** | Pointer to **interface{}** | The rule outcomes. | [optional] 
**ResourcePayload** | Pointer to **interface{}** | The payload used to evaluate the rules. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETFlexPromotionsFlexPromotionId200ResponseDataAttributes

`func NewGETFlexPromotionsFlexPromotionId200ResponseDataAttributes() *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes`

NewGETFlexPromotionsFlexPromotionId200ResponseDataAttributes instantiates a new GETFlexPromotionsFlexPromotionId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETFlexPromotionsFlexPromotionId200ResponseDataAttributesWithDefaults

`func NewGETFlexPromotionsFlexPromotionId200ResponseDataAttributesWithDefaults() *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes`

NewGETFlexPromotionsFlexPromotionId200ResponseDataAttributesWithDefaults instantiates a new GETFlexPromotionsFlexPromotionId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetExclusive

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetExclusive() interface{}`

GetExclusive returns the Exclusive field if non-nil, zero value otherwise.

### GetExclusiveOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetExclusiveOk() (*interface{}, bool)`

GetExclusiveOk returns a tuple with the Exclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusive

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetExclusive(v interface{})`

SetExclusive sets Exclusive field to given value.

### HasExclusive

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasExclusive() bool`

HasExclusive returns a boolean if a field has been set.

### SetExclusiveNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetExclusiveNil(b bool)`

 SetExclusiveNil sets the value for Exclusive to be an explicit nil

### UnsetExclusive
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetExclusive()`

UnsetExclusive ensures that no value is present for Exclusive, not even an explicit nil
### GetPriority

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetPriority() interface{}`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetPriorityOk() (*interface{}, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetPriority(v interface{})`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetStartsAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetStartsAt() interface{}`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetStartsAt(v interface{})`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### SetStartsAtNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetStartsAtNil(b bool)`

 SetStartsAtNil sets the value for StartsAt to be an explicit nil

### UnsetStartsAt
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetStartsAt()`

UnsetStartsAt ensures that no value is present for StartsAt, not even an explicit nil
### GetExpiresAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetTotalUsageLimit

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetTotalUsageLimit() interface{}`

GetTotalUsageLimit returns the TotalUsageLimit field if non-nil, zero value otherwise.

### GetTotalUsageLimitOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetTotalUsageLimitOk() (*interface{}, bool)`

GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageLimit

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetTotalUsageLimit(v interface{})`

SetTotalUsageLimit sets TotalUsageLimit field to given value.

### HasTotalUsageLimit

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasTotalUsageLimit() bool`

HasTotalUsageLimit returns a boolean if a field has been set.

### SetTotalUsageLimitNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetTotalUsageLimitNil(b bool)`

 SetTotalUsageLimitNil sets the value for TotalUsageLimit to be an explicit nil

### UnsetTotalUsageLimit
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetTotalUsageLimit()`

UnsetTotalUsageLimit ensures that no value is present for TotalUsageLimit, not even an explicit nil
### GetTotalUsageCount

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetTotalUsageCount() interface{}`

GetTotalUsageCount returns the TotalUsageCount field if non-nil, zero value otherwise.

### GetTotalUsageCountOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetTotalUsageCountOk() (*interface{}, bool)`

GetTotalUsageCountOk returns a tuple with the TotalUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsageCount

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetTotalUsageCount(v interface{})`

SetTotalUsageCount sets TotalUsageCount field to given value.

### HasTotalUsageCount

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasTotalUsageCount() bool`

HasTotalUsageCount returns a boolean if a field has been set.

### SetTotalUsageCountNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetTotalUsageCountNil(b bool)`

 SetTotalUsageCountNil sets the value for TotalUsageCount to be an explicit nil

### UnsetTotalUsageCount
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetTotalUsageCount()`

UnsetTotalUsageCount ensures that no value is present for TotalUsageCount, not even an explicit nil
### GetActive

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetActive() interface{}`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetActive(v interface{})`

SetActive sets Active field to given value.

### HasActive

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetStatus

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRules

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetRules() interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetRulesOk() (*interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetRules(v interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetDisabledAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetDisabledAt() interface{}`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### GetDisabledAtOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool)`

GetDisabledAtOk returns a tuple with the DisabledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetDisabledAt(v interface{})`

SetDisabledAt sets DisabledAt field to given value.

### HasDisabledAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasDisabledAt() bool`

HasDisabledAt returns a boolean if a field has been set.

### SetDisabledAtNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetDisabledAtNil(b bool)`

 SetDisabledAtNil sets the value for DisabledAt to be an explicit nil

### UnsetDisabledAt
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetDisabledAt()`

UnsetDisabledAt ensures that no value is present for DisabledAt, not even an explicit nil
### GetCreatedAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetRuleOutcomes

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetRuleOutcomes() interface{}`

GetRuleOutcomes returns the RuleOutcomes field if non-nil, zero value otherwise.

### GetRuleOutcomesOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetRuleOutcomesOk() (*interface{}, bool)`

GetRuleOutcomesOk returns a tuple with the RuleOutcomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOutcomes

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetRuleOutcomes(v interface{})`

SetRuleOutcomes sets RuleOutcomes field to given value.

### HasRuleOutcomes

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasRuleOutcomes() bool`

HasRuleOutcomes returns a boolean if a field has been set.

### SetRuleOutcomesNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetRuleOutcomesNil(b bool)`

 SetRuleOutcomesNil sets the value for RuleOutcomes to be an explicit nil

### UnsetRuleOutcomes
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetRuleOutcomes()`

UnsetRuleOutcomes ensures that no value is present for RuleOutcomes, not even an explicit nil
### GetResourcePayload

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetResourcePayload() interface{}`

GetResourcePayload returns the ResourcePayload field if non-nil, zero value otherwise.

### GetResourcePayloadOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetResourcePayloadOk() (*interface{}, bool)`

GetResourcePayloadOk returns a tuple with the ResourcePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePayload

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetResourcePayload(v interface{})`

SetResourcePayload sets ResourcePayload field to given value.

### HasResourcePayload

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasResourcePayload() bool`

HasResourcePayload returns a boolean if a field has been set.

### SetResourcePayloadNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetResourcePayloadNil(b bool)`

 SetResourcePayloadNil sets the value for ResourcePayload to be an explicit nil

### UnsetResourcePayload
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetResourcePayload()`

UnsetResourcePayload ensures that no value is present for ResourcePayload, not even an explicit nil
### GetMetadata

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETFlexPromotionsFlexPromotionId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


