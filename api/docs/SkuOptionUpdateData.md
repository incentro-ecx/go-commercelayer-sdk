# SkuOptionUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHSkuOptionsSkuOptionId200ResponseDataAttributes**](PATCHSkuOptionsSkuOptionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**SkuOptionCreateDataRelationships**](SkuOptionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewSkuOptionUpdateData

`func NewSkuOptionUpdateData(type_ interface{}, id interface{}, attributes PATCHSkuOptionsSkuOptionId200ResponseDataAttributes, ) *SkuOptionUpdateData`

NewSkuOptionUpdateData instantiates a new SkuOptionUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkuOptionUpdateDataWithDefaults

`func NewSkuOptionUpdateDataWithDefaults() *SkuOptionUpdateData`

NewSkuOptionUpdateDataWithDefaults instantiates a new SkuOptionUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SkuOptionUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SkuOptionUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SkuOptionUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *SkuOptionUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SkuOptionUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *SkuOptionUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SkuOptionUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SkuOptionUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *SkuOptionUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SkuOptionUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *SkuOptionUpdateData) GetAttributes() PATCHSkuOptionsSkuOptionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SkuOptionUpdateData) GetAttributesOk() (*PATCHSkuOptionsSkuOptionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SkuOptionUpdateData) SetAttributes(v PATCHSkuOptionsSkuOptionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SkuOptionUpdateData) GetRelationships() SkuOptionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SkuOptionUpdateData) GetRelationshipsOk() (*SkuOptionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SkuOptionUpdateData) SetRelationships(v SkuOptionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *SkuOptionUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


