# CustomerPaymentSourceUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes**](PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerPaymentSourceUpdateDataRelationships**](CustomerPaymentSourceUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerPaymentSourceUpdateData

`func NewCustomerPaymentSourceUpdateData(type_ interface{}, id interface{}, attributes PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes, ) *CustomerPaymentSourceUpdateData`

NewCustomerPaymentSourceUpdateData instantiates a new CustomerPaymentSourceUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentSourceUpdateDataWithDefaults

`func NewCustomerPaymentSourceUpdateDataWithDefaults() *CustomerPaymentSourceUpdateData`

NewCustomerPaymentSourceUpdateDataWithDefaults instantiates a new CustomerPaymentSourceUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerPaymentSourceUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerPaymentSourceUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerPaymentSourceUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomerPaymentSourceUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomerPaymentSourceUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *CustomerPaymentSourceUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerPaymentSourceUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerPaymentSourceUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *CustomerPaymentSourceUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CustomerPaymentSourceUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *CustomerPaymentSourceUpdateData) GetAttributes() PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerPaymentSourceUpdateData) GetAttributesOk() (*PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerPaymentSourceUpdateData) SetAttributes(v PATCHCustomerPaymentSourcesCustomerPaymentSourceId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerPaymentSourceUpdateData) GetRelationships() CustomerPaymentSourceUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerPaymentSourceUpdateData) GetRelationshipsOk() (*CustomerPaymentSourceUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerPaymentSourceUpdateData) SetRelationships(v CustomerPaymentSourceUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerPaymentSourceUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


