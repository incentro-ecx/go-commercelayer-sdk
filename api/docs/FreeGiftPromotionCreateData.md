# FreeGiftPromotionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTFreeGiftPromotions201ResponseDataAttributes**](POSTFreeGiftPromotions201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BuyXPayYPromotionCreateDataRelationships**](BuyXPayYPromotionCreateDataRelationships.md) |  | [optional] 

## Methods

### NewFreeGiftPromotionCreateData

`func NewFreeGiftPromotionCreateData(type_ interface{}, attributes POSTFreeGiftPromotions201ResponseDataAttributes, ) *FreeGiftPromotionCreateData`

NewFreeGiftPromotionCreateData instantiates a new FreeGiftPromotionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFreeGiftPromotionCreateDataWithDefaults

`func NewFreeGiftPromotionCreateDataWithDefaults() *FreeGiftPromotionCreateData`

NewFreeGiftPromotionCreateDataWithDefaults instantiates a new FreeGiftPromotionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FreeGiftPromotionCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FreeGiftPromotionCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FreeGiftPromotionCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FreeGiftPromotionCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FreeGiftPromotionCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *FreeGiftPromotionCreateData) GetAttributes() POSTFreeGiftPromotions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FreeGiftPromotionCreateData) GetAttributesOk() (*POSTFreeGiftPromotions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FreeGiftPromotionCreateData) SetAttributes(v POSTFreeGiftPromotions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FreeGiftPromotionCreateData) GetRelationships() BuyXPayYPromotionCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FreeGiftPromotionCreateData) GetRelationshipsOk() (*BuyXPayYPromotionCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FreeGiftPromotionCreateData) SetRelationships(v BuyXPayYPromotionCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FreeGiftPromotionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


