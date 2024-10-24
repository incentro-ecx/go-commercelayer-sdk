# CustomPromotionRuleUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes**](PATCHCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomPromotionRuleUpdateDataRelationships**](CustomPromotionRuleUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewCustomPromotionRuleUpdateData

`func NewCustomPromotionRuleUpdateData(type_ interface{}, id interface{}, attributes PATCHCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes, ) *CustomPromotionRuleUpdateData`

NewCustomPromotionRuleUpdateData instantiates a new CustomPromotionRuleUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomPromotionRuleUpdateDataWithDefaults

`func NewCustomPromotionRuleUpdateDataWithDefaults() *CustomPromotionRuleUpdateData`

NewCustomPromotionRuleUpdateDataWithDefaults instantiates a new CustomPromotionRuleUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomPromotionRuleUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomPromotionRuleUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomPromotionRuleUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomPromotionRuleUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomPromotionRuleUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *CustomPromotionRuleUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomPromotionRuleUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomPromotionRuleUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *CustomPromotionRuleUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CustomPromotionRuleUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *CustomPromotionRuleUpdateData) GetAttributes() PATCHCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomPromotionRuleUpdateData) GetAttributesOk() (*PATCHCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomPromotionRuleUpdateData) SetAttributes(v PATCHCustomPromotionRulesCustomPromotionRuleId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomPromotionRuleUpdateData) GetRelationships() CustomPromotionRuleUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomPromotionRuleUpdateData) GetRelationshipsOk() (*CustomPromotionRuleUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomPromotionRuleUpdateData) SetRelationships(v CustomPromotionRuleUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomPromotionRuleUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


