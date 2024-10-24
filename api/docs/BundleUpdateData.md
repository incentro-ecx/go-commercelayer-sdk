# BundleUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHBundlesBundleId200ResponseDataAttributes**](PATCHBundlesBundleId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BundleUpdateDataRelationships**](BundleUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewBundleUpdateData

`func NewBundleUpdateData(type_ interface{}, id interface{}, attributes PATCHBundlesBundleId200ResponseDataAttributes, ) *BundleUpdateData`

NewBundleUpdateData instantiates a new BundleUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleUpdateDataWithDefaults

`func NewBundleUpdateDataWithDefaults() *BundleUpdateData`

NewBundleUpdateDataWithDefaults instantiates a new BundleUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BundleUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BundleUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BundleUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *BundleUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BundleUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *BundleUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BundleUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BundleUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *BundleUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BundleUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *BundleUpdateData) GetAttributes() PATCHBundlesBundleId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BundleUpdateData) GetAttributesOk() (*PATCHBundlesBundleId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BundleUpdateData) SetAttributes(v PATCHBundlesBundleId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *BundleUpdateData) GetRelationships() BundleUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BundleUpdateData) GetRelationshipsOk() (*BundleUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BundleUpdateData) SetRelationships(v BundleUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BundleUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


