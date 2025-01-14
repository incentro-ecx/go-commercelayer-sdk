# CarrierAccountCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCarrierAccounts201ResponseDataAttributes**](POSTCarrierAccounts201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CarrierAccountCreateDataRelationships**](CarrierAccountCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCarrierAccountCreateData

`func NewCarrierAccountCreateData(type_ interface{}, attributes POSTCarrierAccounts201ResponseDataAttributes, ) *CarrierAccountCreateData`

NewCarrierAccountCreateData instantiates a new CarrierAccountCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierAccountCreateDataWithDefaults

`func NewCarrierAccountCreateDataWithDefaults() *CarrierAccountCreateData`

NewCarrierAccountCreateDataWithDefaults instantiates a new CarrierAccountCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CarrierAccountCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CarrierAccountCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CarrierAccountCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CarrierAccountCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CarrierAccountCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CarrierAccountCreateData) GetAttributes() POSTCarrierAccounts201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CarrierAccountCreateData) GetAttributesOk() (*POSTCarrierAccounts201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CarrierAccountCreateData) SetAttributes(v POSTCarrierAccounts201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CarrierAccountCreateData) GetRelationships() CarrierAccountCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CarrierAccountCreateData) GetRelationshipsOk() (*CarrierAccountCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CarrierAccountCreateData) SetRelationships(v CarrierAccountCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CarrierAccountCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


