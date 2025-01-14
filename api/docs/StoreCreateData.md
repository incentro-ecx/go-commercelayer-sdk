# StoreCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTStores201ResponseDataAttributes**](POSTStores201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**StoreCreateDataRelationships**](StoreCreateDataRelationships.md) |  | [optional] 

## Methods

### NewStoreCreateData

`func NewStoreCreateData(type_ interface{}, attributes POSTStores201ResponseDataAttributes, ) *StoreCreateData`

NewStoreCreateData instantiates a new StoreCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreCreateDataWithDefaults

`func NewStoreCreateDataWithDefaults() *StoreCreateData`

NewStoreCreateDataWithDefaults instantiates a new StoreCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StoreCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoreCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoreCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StoreCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StoreCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *StoreCreateData) GetAttributes() POSTStores201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StoreCreateData) GetAttributesOk() (*POSTStores201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StoreCreateData) SetAttributes(v POSTStores201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StoreCreateData) GetRelationships() StoreCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StoreCreateData) GetRelationshipsOk() (*StoreCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StoreCreateData) SetRelationships(v StoreCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StoreCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


