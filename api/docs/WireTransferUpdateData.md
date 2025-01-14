# WireTransferUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes**](PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**AdyenPaymentUpdateDataRelationships**](AdyenPaymentUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewWireTransferUpdateData

`func NewWireTransferUpdateData(type_ interface{}, id interface{}, attributes PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes, ) *WireTransferUpdateData`

NewWireTransferUpdateData instantiates a new WireTransferUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireTransferUpdateDataWithDefaults

`func NewWireTransferUpdateDataWithDefaults() *WireTransferUpdateData`

NewWireTransferUpdateDataWithDefaults instantiates a new WireTransferUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WireTransferUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WireTransferUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WireTransferUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *WireTransferUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WireTransferUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *WireTransferUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireTransferUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireTransferUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *WireTransferUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WireTransferUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *WireTransferUpdateData) GetAttributes() PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WireTransferUpdateData) GetAttributesOk() (*PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WireTransferUpdateData) SetAttributes(v PATCHCouponCodesPromotionRulesCouponCodesPromotionRuleId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *WireTransferUpdateData) GetRelationships() AdyenPaymentUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WireTransferUpdateData) GetRelationshipsOk() (*AdyenPaymentUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WireTransferUpdateData) SetRelationships(v AdyenPaymentUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WireTransferUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


