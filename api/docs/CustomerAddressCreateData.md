# CustomerAddressCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTCustomerAddresses201ResponseDataAttributes**](POSTCustomerAddresses201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerAddressCreateDataRelationships**](CustomerAddressCreateDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerAddressCreateData

`func NewCustomerAddressCreateData(type_ interface{}, attributes POSTCustomerAddresses201ResponseDataAttributes, ) *CustomerAddressCreateData`

NewCustomerAddressCreateData instantiates a new CustomerAddressCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressCreateDataWithDefaults

`func NewCustomerAddressCreateDataWithDefaults() *CustomerAddressCreateData`

NewCustomerAddressCreateDataWithDefaults instantiates a new CustomerAddressCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerAddressCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerAddressCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerAddressCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomerAddressCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomerAddressCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *CustomerAddressCreateData) GetAttributes() POSTCustomerAddresses201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerAddressCreateData) GetAttributesOk() (*POSTCustomerAddresses201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerAddressCreateData) SetAttributes(v POSTCustomerAddresses201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerAddressCreateData) GetRelationships() CustomerAddressCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerAddressCreateData) GetRelationshipsOk() (*CustomerAddressCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerAddressCreateData) SetRelationships(v CustomerAddressCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerAddressCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


