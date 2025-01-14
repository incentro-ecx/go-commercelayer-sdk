# POSTOrderCopies201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaceTargetOrder** | Pointer to **interface{}** | Indicates if the target order must be placed upon copy. | [optional] 
**ReuseWallet** | Pointer to **interface{}** | Indicates if the payment source within the source order customer&#39;s wallet must be copied. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**CancelSourceOrder** | Pointer to **interface{}** | Indicates if the source order must be cancelled upon copy. | [optional] 
**ApplyPromotions** | Pointer to **interface{}** | Indicates if promotions got applied upon copy. | [optional] 
**SkipErrors** | Pointer to **interface{}** | Indicates to ignore any errors during copy. | [optional] 
**IgnoreInvalidCoupon** | Pointer to **interface{}** | Indicates to ignore invalid coupon code during copy. | [optional] 

## Methods

### NewPOSTOrderCopies201ResponseDataAttributes

`func NewPOSTOrderCopies201ResponseDataAttributes() *POSTOrderCopies201ResponseDataAttributes`

NewPOSTOrderCopies201ResponseDataAttributes instantiates a new POSTOrderCopies201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrderCopies201ResponseDataAttributesWithDefaults

`func NewPOSTOrderCopies201ResponseDataAttributesWithDefaults() *POSTOrderCopies201ResponseDataAttributes`

NewPOSTOrderCopies201ResponseDataAttributesWithDefaults instantiates a new POSTOrderCopies201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaceTargetOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) GetPlaceTargetOrder() interface{}`

GetPlaceTargetOrder returns the PlaceTargetOrder field if non-nil, zero value otherwise.

### GetPlaceTargetOrderOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetPlaceTargetOrderOk() (*interface{}, bool)`

GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTargetOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) SetPlaceTargetOrder(v interface{})`

SetPlaceTargetOrder sets PlaceTargetOrder field to given value.

### HasPlaceTargetOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) HasPlaceTargetOrder() bool`

HasPlaceTargetOrder returns a boolean if a field has been set.

### SetPlaceTargetOrderNil

`func (o *POSTOrderCopies201ResponseDataAttributes) SetPlaceTargetOrderNil(b bool)`

 SetPlaceTargetOrderNil sets the value for PlaceTargetOrder to be an explicit nil

### UnsetPlaceTargetOrder
`func (o *POSTOrderCopies201ResponseDataAttributes) UnsetPlaceTargetOrder()`

UnsetPlaceTargetOrder ensures that no value is present for PlaceTargetOrder, not even an explicit nil
### GetReuseWallet

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReuseWallet() interface{}`

GetReuseWallet returns the ReuseWallet field if non-nil, zero value otherwise.

### GetReuseWalletOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReuseWalletOk() (*interface{}, bool)`

GetReuseWalletOk returns a tuple with the ReuseWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseWallet

`func (o *POSTOrderCopies201ResponseDataAttributes) SetReuseWallet(v interface{})`

SetReuseWallet sets ReuseWallet field to given value.

### HasReuseWallet

`func (o *POSTOrderCopies201ResponseDataAttributes) HasReuseWallet() bool`

HasReuseWallet returns a boolean if a field has been set.

### SetReuseWalletNil

`func (o *POSTOrderCopies201ResponseDataAttributes) SetReuseWalletNil(b bool)`

 SetReuseWalletNil sets the value for ReuseWallet to be an explicit nil

### UnsetReuseWallet
`func (o *POSTOrderCopies201ResponseDataAttributes) UnsetReuseWallet()`

UnsetReuseWallet ensures that no value is present for ReuseWallet, not even an explicit nil
### GetReference

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTOrderCopies201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTOrderCopies201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTOrderCopies201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTOrderCopies201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTOrderCopies201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTOrderCopies201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTOrderCopies201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTOrderCopies201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTOrderCopies201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTOrderCopies201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTOrderCopies201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTOrderCopies201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTOrderCopies201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCancelSourceOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) GetCancelSourceOrder() interface{}`

GetCancelSourceOrder returns the CancelSourceOrder field if non-nil, zero value otherwise.

### GetCancelSourceOrderOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetCancelSourceOrderOk() (*interface{}, bool)`

GetCancelSourceOrderOk returns a tuple with the CancelSourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSourceOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) SetCancelSourceOrder(v interface{})`

SetCancelSourceOrder sets CancelSourceOrder field to given value.

### HasCancelSourceOrder

`func (o *POSTOrderCopies201ResponseDataAttributes) HasCancelSourceOrder() bool`

HasCancelSourceOrder returns a boolean if a field has been set.

### SetCancelSourceOrderNil

`func (o *POSTOrderCopies201ResponseDataAttributes) SetCancelSourceOrderNil(b bool)`

 SetCancelSourceOrderNil sets the value for CancelSourceOrder to be an explicit nil

### UnsetCancelSourceOrder
`func (o *POSTOrderCopies201ResponseDataAttributes) UnsetCancelSourceOrder()`

UnsetCancelSourceOrder ensures that no value is present for CancelSourceOrder, not even an explicit nil
### GetApplyPromotions

`func (o *POSTOrderCopies201ResponseDataAttributes) GetApplyPromotions() interface{}`

GetApplyPromotions returns the ApplyPromotions field if non-nil, zero value otherwise.

### GetApplyPromotionsOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetApplyPromotionsOk() (*interface{}, bool)`

GetApplyPromotionsOk returns a tuple with the ApplyPromotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromotions

`func (o *POSTOrderCopies201ResponseDataAttributes) SetApplyPromotions(v interface{})`

SetApplyPromotions sets ApplyPromotions field to given value.

### HasApplyPromotions

`func (o *POSTOrderCopies201ResponseDataAttributes) HasApplyPromotions() bool`

HasApplyPromotions returns a boolean if a field has been set.

### SetApplyPromotionsNil

`func (o *POSTOrderCopies201ResponseDataAttributes) SetApplyPromotionsNil(b bool)`

 SetApplyPromotionsNil sets the value for ApplyPromotions to be an explicit nil

### UnsetApplyPromotions
`func (o *POSTOrderCopies201ResponseDataAttributes) UnsetApplyPromotions()`

UnsetApplyPromotions ensures that no value is present for ApplyPromotions, not even an explicit nil
### GetSkipErrors

`func (o *POSTOrderCopies201ResponseDataAttributes) GetSkipErrors() interface{}`

GetSkipErrors returns the SkipErrors field if non-nil, zero value otherwise.

### GetSkipErrorsOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetSkipErrorsOk() (*interface{}, bool)`

GetSkipErrorsOk returns a tuple with the SkipErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipErrors

`func (o *POSTOrderCopies201ResponseDataAttributes) SetSkipErrors(v interface{})`

SetSkipErrors sets SkipErrors field to given value.

### HasSkipErrors

`func (o *POSTOrderCopies201ResponseDataAttributes) HasSkipErrors() bool`

HasSkipErrors returns a boolean if a field has been set.

### SetSkipErrorsNil

`func (o *POSTOrderCopies201ResponseDataAttributes) SetSkipErrorsNil(b bool)`

 SetSkipErrorsNil sets the value for SkipErrors to be an explicit nil

### UnsetSkipErrors
`func (o *POSTOrderCopies201ResponseDataAttributes) UnsetSkipErrors()`

UnsetSkipErrors ensures that no value is present for SkipErrors, not even an explicit nil
### GetIgnoreInvalidCoupon

`func (o *POSTOrderCopies201ResponseDataAttributes) GetIgnoreInvalidCoupon() interface{}`

GetIgnoreInvalidCoupon returns the IgnoreInvalidCoupon field if non-nil, zero value otherwise.

### GetIgnoreInvalidCouponOk

`func (o *POSTOrderCopies201ResponseDataAttributes) GetIgnoreInvalidCouponOk() (*interface{}, bool)`

GetIgnoreInvalidCouponOk returns a tuple with the IgnoreInvalidCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreInvalidCoupon

`func (o *POSTOrderCopies201ResponseDataAttributes) SetIgnoreInvalidCoupon(v interface{})`

SetIgnoreInvalidCoupon sets IgnoreInvalidCoupon field to given value.

### HasIgnoreInvalidCoupon

`func (o *POSTOrderCopies201ResponseDataAttributes) HasIgnoreInvalidCoupon() bool`

HasIgnoreInvalidCoupon returns a boolean if a field has been set.

### SetIgnoreInvalidCouponNil

`func (o *POSTOrderCopies201ResponseDataAttributes) SetIgnoreInvalidCouponNil(b bool)`

 SetIgnoreInvalidCouponNil sets the value for IgnoreInvalidCoupon to be an explicit nil

### UnsetIgnoreInvalidCoupon
`func (o *POSTOrderCopies201ResponseDataAttributes) UnsetIgnoreInvalidCoupon()`

UnsetIgnoreInvalidCoupon ensures that no value is present for IgnoreInvalidCoupon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


