# CustomerPaymentSourceCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCustomerPaymentSources201ResponseDataAttributes**](POSTCustomerPaymentSources201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerPaymentSourceCreateDataRelationships**](CustomerPaymentSourceCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerPaymentSourceCreateData

`func NewCustomerPaymentSourceCreateData(type_ interface{}, attributes POSTCustomerPaymentSources201ResponseDataAttributes, ) *CustomerPaymentSourceCreateData`

NewCustomerPaymentSourceCreateData instantiates a new CustomerPaymentSourceCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerPaymentSourceCreateDataWithDefaults

`func NewCustomerPaymentSourceCreateDataWithDefaults() *CustomerPaymentSourceCreateData`

NewCustomerPaymentSourceCreateDataWithDefaults instantiates a new CustomerPaymentSourceCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerPaymentSourceCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerPaymentSourceCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerPaymentSourceCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomerPaymentSourceCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomerPaymentSourceCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CustomerPaymentSourceCreateData) GetAttributes() POSTCustomerPaymentSources201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerPaymentSourceCreateData) GetAttributesOk() (*POSTCustomerPaymentSources201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerPaymentSourceCreateData) SetAttributes(v POSTCustomerPaymentSources201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerPaymentSourceCreateData) GetRelationships() CustomerPaymentSourceCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerPaymentSourceCreateData) GetRelationshipsOk() (*CustomerPaymentSourceCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerPaymentSourceCreateData) SetRelationships(v CustomerPaymentSourceCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerPaymentSourceCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


