# PercentageDiscountPromotionCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTPercentageDiscountPromotions201ResponseDataAttributes**](POSTPercentageDiscountPromotions201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BuyXPayYPromotionUpdateDataRelationships**](BuyXPayYPromotionUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewPercentageDiscountPromotionCreateData

`func NewPercentageDiscountPromotionCreateData(type_ interface{}, attributes POSTPercentageDiscountPromotions201ResponseDataAttributes, ) *PercentageDiscountPromotionCreateData`

NewPercentageDiscountPromotionCreateData instantiates a new PercentageDiscountPromotionCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPercentageDiscountPromotionCreateDataWithDefaults

`func NewPercentageDiscountPromotionCreateDataWithDefaults() *PercentageDiscountPromotionCreateData`

NewPercentageDiscountPromotionCreateDataWithDefaults instantiates a new PercentageDiscountPromotionCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PercentageDiscountPromotionCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PercentageDiscountPromotionCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PercentageDiscountPromotionCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *PercentageDiscountPromotionCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PercentageDiscountPromotionCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *PercentageDiscountPromotionCreateData) GetAttributes() POSTPercentageDiscountPromotions201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PercentageDiscountPromotionCreateData) GetAttributesOk() (*POSTPercentageDiscountPromotions201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PercentageDiscountPromotionCreateData) SetAttributes(v POSTPercentageDiscountPromotions201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *PercentageDiscountPromotionCreateData) GetRelationships() BuyXPayYPromotionUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PercentageDiscountPromotionCreateData) GetRelationshipsOk() (*BuyXPayYPromotionUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PercentageDiscountPromotionCreateData) SetRelationships(v BuyXPayYPromotionUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PercentageDiscountPromotionCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


