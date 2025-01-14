# CustomerAddressUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes**](PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**CustomerAddressUpdateDataRelationships**](CustomerAddressUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewCustomerAddressUpdateData

`func NewCustomerAddressUpdateData(type_ interface{}, id interface{}, attributes PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes, ) *CustomerAddressUpdateData`

NewCustomerAddressUpdateData instantiates a new CustomerAddressUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAddressUpdateDataWithDefaults

`func NewCustomerAddressUpdateDataWithDefaults() *CustomerAddressUpdateData`

NewCustomerAddressUpdateDataWithDefaults instantiates a new CustomerAddressUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerAddressUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerAddressUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerAddressUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CustomerAddressUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomerAddressUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *CustomerAddressUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerAddressUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerAddressUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *CustomerAddressUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CustomerAddressUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *CustomerAddressUpdateData) GetAttributes() PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerAddressUpdateData) GetAttributesOk() (*PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerAddressUpdateData) SetAttributes(v PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *CustomerAddressUpdateData) GetRelationships() CustomerAddressUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerAddressUpdateData) GetRelationshipsOk() (*CustomerAddressUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerAddressUpdateData) SetRelationships(v CustomerAddressUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerAddressUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


