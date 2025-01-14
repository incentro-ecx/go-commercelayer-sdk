# GETRefundsRefundId200ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsOrder**](POSTAdyenPayments201ResponseDataRelationshipsOrder.md) |  | [optional] 
**PaymentSource** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 
**ReferenceCapture** | Pointer to [**GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture**](GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture.md) |  | [optional] 
**Return** | Pointer to [**GETCapturesCaptureId200ResponseDataRelationshipsReturn**](GETCapturesCaptureId200ResponseDataRelationshipsReturn.md) |  | [optional] 

## Methods

### NewGETRefundsRefundId200ResponseDataRelationships

`func NewGETRefundsRefundId200ResponseDataRelationships() *GETRefundsRefundId200ResponseDataRelationships`

NewGETRefundsRefundId200ResponseDataRelationships instantiates a new GETRefundsRefundId200ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETRefundsRefundId200ResponseDataRelationshipsWithDefaults

`func NewGETRefundsRefundId200ResponseDataRelationshipsWithDefaults() *GETRefundsRefundId200ResponseDataRelationships`

NewGETRefundsRefundId200ResponseDataRelationshipsWithDefaults instantiates a new GETRefundsRefundId200ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETRefundsRefundId200ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETRefundsRefundId200ResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentSource

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetPaymentSource() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetPaymentSourceOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *GETRefundsRefundId200ResponseDataRelationships) SetPaymentSource(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *GETRefundsRefundId200ResponseDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetAttachments

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETRefundsRefundId200ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETRefundsRefundId200ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETRefundsRefundId200ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETRefundsRefundId200ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GETRefundsRefundId200ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GETRefundsRefundId200ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetReferenceCapture

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetReferenceCapture() GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture`

GetReferenceCapture returns the ReferenceCapture field if non-nil, zero value otherwise.

### GetReferenceCaptureOk

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetReferenceCaptureOk() (*GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture, bool)`

GetReferenceCaptureOk returns a tuple with the ReferenceCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCapture

`func (o *GETRefundsRefundId200ResponseDataRelationships) SetReferenceCapture(v GETRefundsRefundId200ResponseDataRelationshipsReferenceCapture)`

SetReferenceCapture sets ReferenceCapture field to given value.

### HasReferenceCapture

`func (o *GETRefundsRefundId200ResponseDataRelationships) HasReferenceCapture() bool`

HasReferenceCapture returns a boolean if a field has been set.

### GetReturn

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetReturn() GETCapturesCaptureId200ResponseDataRelationshipsReturn`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *GETRefundsRefundId200ResponseDataRelationships) GetReturnOk() (*GETCapturesCaptureId200ResponseDataRelationshipsReturn, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *GETRefundsRefundId200ResponseDataRelationships) SetReturn(v GETCapturesCaptureId200ResponseDataRelationshipsReturn)`

SetReturn sets Return field to given value.

### HasReturn

`func (o *GETRefundsRefundId200ResponseDataRelationships) HasReturn() bool`

HasReturn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


