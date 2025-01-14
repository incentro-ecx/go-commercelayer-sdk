# GETCapturesCaptureId200ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsOrder**](POSTAdyenPayments201ResponseDataRelationshipsOrder.md) |  | [optional] 
**PaymentSource** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 
**ReferenceAuthorization** | Pointer to [**GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization**](GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization.md) |  | [optional] 
**Refunds** | Pointer to [**GETCapturesCaptureId200ResponseDataRelationshipsRefunds**](GETCapturesCaptureId200ResponseDataRelationshipsRefunds.md) |  | [optional] 
**Return** | Pointer to [**GETCapturesCaptureId200ResponseDataRelationshipsReturn**](GETCapturesCaptureId200ResponseDataRelationshipsReturn.md) |  | [optional] 

## Methods

### NewGETCapturesCaptureId200ResponseDataRelationships

`func NewGETCapturesCaptureId200ResponseDataRelationships() *GETCapturesCaptureId200ResponseDataRelationships`

NewGETCapturesCaptureId200ResponseDataRelationships instantiates a new GETCapturesCaptureId200ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETCapturesCaptureId200ResponseDataRelationshipsWithDefaults

`func NewGETCapturesCaptureId200ResponseDataRelationshipsWithDefaults() *GETCapturesCaptureId200ResponseDataRelationships`

NewGETCapturesCaptureId200ResponseDataRelationshipsWithDefaults instantiates a new GETCapturesCaptureId200ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETCapturesCaptureId200ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETCapturesCaptureId200ResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentSource

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetPaymentSource() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetPaymentSourceOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *GETCapturesCaptureId200ResponseDataRelationships) SetPaymentSource(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *GETCapturesCaptureId200ResponseDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetAttachments

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETCapturesCaptureId200ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETCapturesCaptureId200ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETCapturesCaptureId200ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETCapturesCaptureId200ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GETCapturesCaptureId200ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GETCapturesCaptureId200ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetReferenceAuthorization

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetReferenceAuthorization() GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization`

GetReferenceAuthorization returns the ReferenceAuthorization field if non-nil, zero value otherwise.

### GetReferenceAuthorizationOk

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetReferenceAuthorizationOk() (*GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization, bool)`

GetReferenceAuthorizationOk returns a tuple with the ReferenceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAuthorization

`func (o *GETCapturesCaptureId200ResponseDataRelationships) SetReferenceAuthorization(v GETCapturesCaptureId200ResponseDataRelationshipsReferenceAuthorization)`

SetReferenceAuthorization sets ReferenceAuthorization field to given value.

### HasReferenceAuthorization

`func (o *GETCapturesCaptureId200ResponseDataRelationships) HasReferenceAuthorization() bool`

HasReferenceAuthorization returns a boolean if a field has been set.

### GetRefunds

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetRefunds() GETCapturesCaptureId200ResponseDataRelationshipsRefunds`

GetRefunds returns the Refunds field if non-nil, zero value otherwise.

### GetRefundsOk

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetRefundsOk() (*GETCapturesCaptureId200ResponseDataRelationshipsRefunds, bool)`

GetRefundsOk returns a tuple with the Refunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunds

`func (o *GETCapturesCaptureId200ResponseDataRelationships) SetRefunds(v GETCapturesCaptureId200ResponseDataRelationshipsRefunds)`

SetRefunds sets Refunds field to given value.

### HasRefunds

`func (o *GETCapturesCaptureId200ResponseDataRelationships) HasRefunds() bool`

HasRefunds returns a boolean if a field has been set.

### GetReturn

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetReturn() GETCapturesCaptureId200ResponseDataRelationshipsReturn`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *GETCapturesCaptureId200ResponseDataRelationships) GetReturnOk() (*GETCapturesCaptureId200ResponseDataRelationshipsReturn, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturn

`func (o *GETCapturesCaptureId200ResponseDataRelationships) SetReturn(v GETCapturesCaptureId200ResponseDataRelationshipsReturn)`

SetReturn sets Return field to given value.

### HasReturn

`func (o *GETCapturesCaptureId200ResponseDataRelationships) HasReturn() bool`

HasReturn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


