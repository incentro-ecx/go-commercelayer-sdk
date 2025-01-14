# ImportUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHCleanupsCleanupId200ResponseDataAttributes**](PATCHCleanupsCleanupId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewImportUpdateData

`func NewImportUpdateData(type_ interface{}, id interface{}, attributes PATCHCleanupsCleanupId200ResponseDataAttributes, ) *ImportUpdateData`

NewImportUpdateData instantiates a new ImportUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportUpdateDataWithDefaults

`func NewImportUpdateDataWithDefaults() *ImportUpdateData`

NewImportUpdateDataWithDefaults instantiates a new ImportUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImportUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImportUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImportUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ImportUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ImportUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *ImportUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImportUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImportUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *ImportUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ImportUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *ImportUpdateData) GetAttributes() PATCHCleanupsCleanupId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ImportUpdateData) GetAttributesOk() (*PATCHCleanupsCleanupId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ImportUpdateData) SetAttributes(v PATCHCleanupsCleanupId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *ImportUpdateData) GetRelationships() interface{}`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ImportUpdateData) GetRelationshipsOk() (*interface{}, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ImportUpdateData) SetRelationships(v interface{})`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ImportUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### SetRelationshipsNil

`func (o *ImportUpdateData) SetRelationshipsNil(b bool)`

 SetRelationshipsNil sets the value for Relationships to be an explicit nil

### UnsetRelationships
`func (o *ImportUpdateData) UnsetRelationships()`

UnsetRelationships ensures that no value is present for Relationships, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


