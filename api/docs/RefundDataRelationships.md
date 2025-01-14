# RefundDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**PaymentSource** | Pointer to [**AuthorizationDataRelationshipsPaymentSource**](AuthorizationDataRelationshipsPaymentSource.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 
**ReferenceCapture** | Pointer to [**AuthorizationDataRelationshipsCaptures**](AuthorizationDataRelationshipsCaptures.md) |  | [optional] 
**Return** | Pointer to [**CaptureDataRelationshipsReturn**](CaptureDataRelationshipsReturn.md) |  | [optional] 

## Methods

### NewRefundDataRelationships

`func NewRefundDataRelationships() *RefundDataRelationships`

NewRefundDataRelationships instantiates a new RefundDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundDataRelationshipsWithDefaults

`func NewRefundDataRelationshipsWithDefaults() *RefundDataRelationships`

NewRefundDataRelationshipsWithDefaults instantiates a new RefundDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *RefundDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *RefundDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *RefundDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *RefundDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentSource

`func (o *RefundDataRelationships) GetPaymentSource() AuthorizationDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *RefundDataRelationships) GetPaymentSourceOk() (*AuthorizationDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *RefundDataRelationships) SetPaymentSource(v AuthorizationDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *RefundDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetAttachments

`func (o *RefundDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *RefundDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *RefundDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *RefundDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *RefundDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RefundDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RefundDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RefundDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *RefundDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *RefundDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *RefundDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *RefundDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetReferenceCapture

`func (o *RefundDataRelationships) GetReferenceCapture() AuthorizationDataRelationshipsCaptures`

GetReferenceCapture returns the ReferenceCapture field if non-nil, zero value otherwise.

### GetReferenceCaptureOk

`func (o *RefundDataRelationships) GetReferenceCaptureOk() (*AuthorizationDataRelationshipsCaptures, bool)`

GetReferenceCaptureOk returns a tuple with the ReferenceCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCapture

`func (o *RefundDataRelationships) SetReferenceCapture(v AuthorizationDataRelationshipsCaptures)`

SetReferenceCapture sets ReferenceCapture field to given value.

### HasReferenceCapture

`func (o *RefundDataRelationships) HasReferenceCapture() bool`

HasReferenceCapture returns a boolean if a field has been set.

### GetReturn

`func (o *RefundDataRelationships) GetReturn() CaptureDataRelationshipsReturn`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *RefundDataRelationships) GetReturnOk() (*CaptureDataRelationshipsReturn, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *RefundDataRelationships) SetReturn(v CaptureDataRelationshipsReturn)`

SetReturn sets Return field to given value.

### HasReturn

`func (o *RefundDataRelationships) HasReturn() bool`

HasReturn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


