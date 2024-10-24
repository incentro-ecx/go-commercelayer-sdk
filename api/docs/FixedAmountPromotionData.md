# FixedAmountPromotionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes**](GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**BuyXPayYPromotionDataRelationships**](BuyXPayYPromotionDataRelationships.md) |  | [optional] 

## Methods

### NewFixedAmountPromotionData

`func NewFixedAmountPromotionData(type_ interface{}, attributes GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes, ) *FixedAmountPromotionData`

NewFixedAmountPromotionData instantiates a new FixedAmountPromotionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedAmountPromotionDataWithDefaults

`func NewFixedAmountPromotionDataWithDefaults() *FixedAmountPromotionData`

NewFixedAmountPromotionDataWithDefaults instantiates a new FixedAmountPromotionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FixedAmountPromotionData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FixedAmountPromotionData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FixedAmountPromotionData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *FixedAmountPromotionData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *FixedAmountPromotionData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *FixedAmountPromotionData) GetAttributes() GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FixedAmountPromotionData) GetAttributesOk() (*GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FixedAmountPromotionData) SetAttributes(v GETFixedAmountPromotionsFixedAmountPromotionId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *FixedAmountPromotionData) GetRelationships() BuyXPayYPromotionDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FixedAmountPromotionData) GetRelationshipsOk() (*BuyXPayYPromotionDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FixedAmountPromotionData) SetRelationships(v BuyXPayYPromotionDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FixedAmountPromotionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


