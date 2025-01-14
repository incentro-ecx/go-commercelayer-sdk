# GETAuthorizationsAuthorizationId200ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsOrder**](POSTAdyenPayments201ResponseDataRelationshipsOrder.md) |  | [optional] 
**PaymentSource** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**POSTAddresses201ResponseDataRelationshipsEvents**](POSTAddresses201ResponseDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 
**Captures** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures.md) |  | [optional] 
**Voids** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids.md) |  | [optional] 

## Methods

### NewGETAuthorizationsAuthorizationId200ResponseDataRelationships

`func NewGETAuthorizationsAuthorizationId200ResponseDataRelationships() *GETAuthorizationsAuthorizationId200ResponseDataRelationships`

NewGETAuthorizationsAuthorizationId200ResponseDataRelationships instantiates a new GETAuthorizationsAuthorizationId200ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETAuthorizationsAuthorizationId200ResponseDataRelationshipsWithDefaults

`func NewGETAuthorizationsAuthorizationId200ResponseDataRelationshipsWithDefaults() *GETAuthorizationsAuthorizationId200ResponseDataRelationships`

NewGETAuthorizationsAuthorizationId200ResponseDataRelationshipsWithDefaults instantiates a new GETAuthorizationsAuthorizationId200ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentSource

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetPaymentSource() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetPaymentSourceOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetPaymentSource(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetAttachments

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetCaptures

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetCaptures() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetCapturesOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetCaptures(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetVoids() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetVoidsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetVoids(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


