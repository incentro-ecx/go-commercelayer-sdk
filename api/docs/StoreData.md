# StoreData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETStoresStoreId200ResponseDataAttributes**](GETStoresStoreId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**StoreDataRelationships**](StoreDataRelationships.md) |  | [optional] 

## Methods

### NewStoreData

`func NewStoreData(type_ interface{}, attributes GETStoresStoreId200ResponseDataAttributes, ) *StoreData`

NewStoreData instantiates a new StoreData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDataWithDefaults

`func NewStoreDataWithDefaults() *StoreData`

NewStoreDataWithDefaults instantiates a new StoreData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StoreData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StoreData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StoreData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *StoreData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StoreData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *StoreData) GetAttributes() GETStoresStoreId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *StoreData) GetAttributesOk() (*GETStoresStoreId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *StoreData) SetAttributes(v GETStoresStoreId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *StoreData) GetRelationships() StoreDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StoreData) GetRelationshipsOk() (*StoreDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StoreData) SetRelationships(v StoreDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StoreData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


