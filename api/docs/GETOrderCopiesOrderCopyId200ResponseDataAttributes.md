# GETOrderCopiesOrderCopyId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **interface{}** | The order factory&#39;s type. | [optional] 
**Status** | Pointer to **interface{}** | The order factory status. One of &#39;pending&#39; (default), &#39;in_progress&#39;, &#39;failed&#39;, or &#39;completed&#39;. | [optional] 
**StartedAt** | Pointer to **interface{}** | Time at which the order copy was started. | [optional] 
**CompletedAt** | Pointer to **interface{}** | Time at which the order copy was completed. | [optional] 
**FailedAt** | Pointer to **interface{}** | Time at which the order copy has failed. | [optional] 
**ErrorsLog** | Pointer to **interface{}** | Contains the order copy errors, if any. | [optional] 
**ErrorsCount** | Pointer to **interface{}** | Indicates the number of copy errors, if any. | [optional] 
**PlaceTargetOrder** | Pointer to **interface{}** | Indicates if the target order must be placed upon copy. | [optional] 
**ReuseWallet** | Pointer to **interface{}** | Indicates if the payment source within the source order customer&#39;s wallet must be copied. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 
**CancelSourceOrder** | Pointer to **interface{}** | Indicates if the source order must be cancelled upon copy. | [optional] 
**ApplyPromotions** | Pointer to **interface{}** | Indicates if promotions got applied upon copy. | [optional] 
**SkipErrors** | Pointer to **interface{}** | Indicates to ignore any errors during copy. | [optional] 
**IgnoreInvalidCoupon** | Pointer to **interface{}** | Indicates to ignore invalid coupon code during copy. | [optional] 

## Methods

### NewGETOrderCopiesOrderCopyId200ResponseDataAttributes

`func NewGETOrderCopiesOrderCopyId200ResponseDataAttributes() *GETOrderCopiesOrderCopyId200ResponseDataAttributes`

NewGETOrderCopiesOrderCopyId200ResponseDataAttributes instantiates a new GETOrderCopiesOrderCopyId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrderCopiesOrderCopyId200ResponseDataAttributesWithDefaults

`func NewGETOrderCopiesOrderCopyId200ResponseDataAttributesWithDefaults() *GETOrderCopiesOrderCopyId200ResponseDataAttributes`

NewGETOrderCopiesOrderCopyId200ResponseDataAttributesWithDefaults instantiates a new GETOrderCopiesOrderCopyId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStatus

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetStatus() interface{}`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetStatus(v interface{})`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetStartedAt() interface{}`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetStartedAtOk() (*interface{}, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetStartedAt(v interface{})`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetCompletedAt() interface{}`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetCompletedAtOk() (*interface{}, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetCompletedAt(v interface{})`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetFailedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetFailedAt() interface{}`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetFailedAtOk() (*interface{}, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetFailedAt(v interface{})`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### SetFailedAtNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil
### GetErrorsLog

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetErrorsLog() interface{}`

GetErrorsLog returns the ErrorsLog field if non-nil, zero value otherwise.

### GetErrorsLogOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetErrorsLogOk() (*interface{}, bool)`

GetErrorsLogOk returns a tuple with the ErrorsLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsLog

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetErrorsLog(v interface{})`

SetErrorsLog sets ErrorsLog field to given value.

### HasErrorsLog

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasErrorsLog() bool`

HasErrorsLog returns a boolean if a field has been set.

### SetErrorsLogNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetErrorsLogNil(b bool)`

 SetErrorsLogNil sets the value for ErrorsLog to be an explicit nil

### UnsetErrorsLog
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetErrorsLog()`

UnsetErrorsLog ensures that no value is present for ErrorsLog, not even an explicit nil
### GetErrorsCount

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetErrorsCount() interface{}`

GetErrorsCount returns the ErrorsCount field if non-nil, zero value otherwise.

### GetErrorsCountOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetErrorsCountOk() (*interface{}, bool)`

GetErrorsCountOk returns a tuple with the ErrorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorsCount

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetErrorsCount(v interface{})`

SetErrorsCount sets ErrorsCount field to given value.

### HasErrorsCount

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasErrorsCount() bool`

HasErrorsCount returns a boolean if a field has been set.

### SetErrorsCountNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetErrorsCountNil(b bool)`

 SetErrorsCountNil sets the value for ErrorsCount to be an explicit nil

### UnsetErrorsCount
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetErrorsCount()`

UnsetErrorsCount ensures that no value is present for ErrorsCount, not even an explicit nil
### GetPlaceTargetOrder

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetPlaceTargetOrder() interface{}`

GetPlaceTargetOrder returns the PlaceTargetOrder field if non-nil, zero value otherwise.

### GetPlaceTargetOrderOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetPlaceTargetOrderOk() (*interface{}, bool)`

GetPlaceTargetOrderOk returns a tuple with the PlaceTargetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceTargetOrder

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetPlaceTargetOrder(v interface{})`

SetPlaceTargetOrder sets PlaceTargetOrder field to given value.

### HasPlaceTargetOrder

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasPlaceTargetOrder() bool`

HasPlaceTargetOrder returns a boolean if a field has been set.

### SetPlaceTargetOrderNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetPlaceTargetOrderNil(b bool)`

 SetPlaceTargetOrderNil sets the value for PlaceTargetOrder to be an explicit nil

### UnsetPlaceTargetOrder
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetPlaceTargetOrder()`

UnsetPlaceTargetOrder ensures that no value is present for PlaceTargetOrder, not even an explicit nil
### GetReuseWallet

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetReuseWallet() interface{}`

GetReuseWallet returns the ReuseWallet field if non-nil, zero value otherwise.

### GetReuseWalletOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetReuseWalletOk() (*interface{}, bool)`

GetReuseWalletOk returns a tuple with the ReuseWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseWallet

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetReuseWallet(v interface{})`

SetReuseWallet sets ReuseWallet field to given value.

### HasReuseWallet

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasReuseWallet() bool`

HasReuseWallet returns a boolean if a field has been set.

### SetReuseWalletNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetReuseWalletNil(b bool)`

 SetReuseWalletNil sets the value for ReuseWallet to be an explicit nil

### UnsetReuseWallet
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetReuseWallet()`

UnsetReuseWallet ensures that no value is present for ReuseWallet, not even an explicit nil
### GetCreatedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetCancelSourceOrder

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetCancelSourceOrder() interface{}`

GetCancelSourceOrder returns the CancelSourceOrder field if non-nil, zero value otherwise.

### GetCancelSourceOrderOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetCancelSourceOrderOk() (*interface{}, bool)`

GetCancelSourceOrderOk returns a tuple with the CancelSourceOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelSourceOrder

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetCancelSourceOrder(v interface{})`

SetCancelSourceOrder sets CancelSourceOrder field to given value.

### HasCancelSourceOrder

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasCancelSourceOrder() bool`

HasCancelSourceOrder returns a boolean if a field has been set.

### SetCancelSourceOrderNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetCancelSourceOrderNil(b bool)`

 SetCancelSourceOrderNil sets the value for CancelSourceOrder to be an explicit nil

### UnsetCancelSourceOrder
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetCancelSourceOrder()`

UnsetCancelSourceOrder ensures that no value is present for CancelSourceOrder, not even an explicit nil
### GetApplyPromotions

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetApplyPromotions() interface{}`

GetApplyPromotions returns the ApplyPromotions field if non-nil, zero value otherwise.

### GetApplyPromotionsOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetApplyPromotionsOk() (*interface{}, bool)`

GetApplyPromotionsOk returns a tuple with the ApplyPromotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPromotions

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetApplyPromotions(v interface{})`

SetApplyPromotions sets ApplyPromotions field to given value.

### HasApplyPromotions

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasApplyPromotions() bool`

HasApplyPromotions returns a boolean if a field has been set.

### SetApplyPromotionsNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetApplyPromotionsNil(b bool)`

 SetApplyPromotionsNil sets the value for ApplyPromotions to be an explicit nil

### UnsetApplyPromotions
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetApplyPromotions()`

UnsetApplyPromotions ensures that no value is present for ApplyPromotions, not even an explicit nil
### GetSkipErrors

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetSkipErrors() interface{}`

GetSkipErrors returns the SkipErrors field if non-nil, zero value otherwise.

### GetSkipErrorsOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetSkipErrorsOk() (*interface{}, bool)`

GetSkipErrorsOk returns a tuple with the SkipErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipErrors

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetSkipErrors(v interface{})`

SetSkipErrors sets SkipErrors field to given value.

### HasSkipErrors

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasSkipErrors() bool`

HasSkipErrors returns a boolean if a field has been set.

### SetSkipErrorsNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetSkipErrorsNil(b bool)`

 SetSkipErrorsNil sets the value for SkipErrors to be an explicit nil

### UnsetSkipErrors
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetSkipErrors()`

UnsetSkipErrors ensures that no value is present for SkipErrors, not even an explicit nil
### GetIgnoreInvalidCoupon

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetIgnoreInvalidCoupon() interface{}`

GetIgnoreInvalidCoupon returns the IgnoreInvalidCoupon field if non-nil, zero value otherwise.

### GetIgnoreInvalidCouponOk

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) GetIgnoreInvalidCouponOk() (*interface{}, bool)`

GetIgnoreInvalidCouponOk returns a tuple with the IgnoreInvalidCoupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreInvalidCoupon

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetIgnoreInvalidCoupon(v interface{})`

SetIgnoreInvalidCoupon sets IgnoreInvalidCoupon field to given value.

### HasIgnoreInvalidCoupon

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) HasIgnoreInvalidCoupon() bool`

HasIgnoreInvalidCoupon returns a boolean if a field has been set.

### SetIgnoreInvalidCouponNil

`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) SetIgnoreInvalidCouponNil(b bool)`

 SetIgnoreInvalidCouponNil sets the value for IgnoreInvalidCoupon to be an explicit nil

### UnsetIgnoreInvalidCoupon
`func (o *GETOrderCopiesOrderCopyId200ResponseDataAttributes) UnsetIgnoreInvalidCoupon()`

UnsetIgnoreInvalidCoupon ensures that no value is present for IgnoreInvalidCoupon, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


