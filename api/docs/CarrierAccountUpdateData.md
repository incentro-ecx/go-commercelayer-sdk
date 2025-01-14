# CarrierAccountUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes**](PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CarrierAccountCreateDataRelationships**](CarrierAccountCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCarrierAccountUpdateData

`func NewCarrierAccountUpdateData(type_ interface{}, id interface{}, attributes PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes, ) *CarrierAccountUpdateData`

NewCarrierAccountUpdateData instantiates a new CarrierAccountUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierAccountUpdateDataWithDefaults

`func NewCarrierAccountUpdateDataWithDefaults() *CarrierAccountUpdateData`

NewCarrierAccountUpdateDataWithDefaults instantiates a new CarrierAccountUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CarrierAccountUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CarrierAccountUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CarrierAccountUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CarrierAccountUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CarrierAccountUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *CarrierAccountUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CarrierAccountUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CarrierAccountUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *CarrierAccountUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CarrierAccountUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *CarrierAccountUpdateData) GetAttributes() PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CarrierAccountUpdateData) GetAttributesOk() (*PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CarrierAccountUpdateData) SetAttributes(v PATCHCarrierAccountsCarrierAccountId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CarrierAccountUpdateData) GetRelationships() CarrierAccountCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CarrierAccountUpdateData) GetRelationshipsOk() (*CarrierAccountCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CarrierAccountUpdateData) SetRelationships(v CarrierAccountCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CarrierAccountUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


