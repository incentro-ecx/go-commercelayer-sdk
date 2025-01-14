# AuthorizationDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AdyenPaymentDataRelationshipsOrder**](AdyenPaymentDataRelationshipsOrder.md) |  | [optional] 
**PaymentSource** | Pointer to [**AuthorizationDataRelationshipsPaymentSource**](AuthorizationDataRelationshipsPaymentSource.md) |  | [optional] 
**Attachments** | Pointer to [**AuthorizationDataRelationshipsAttachments**](AuthorizationDataRelationshipsAttachments.md) |  | [optional] 
**Events** | Pointer to [**AddressDataRelationshipsEvents**](AddressDataRelationshipsEvents.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 
**Captures** | Pointer to [**AuthorizationDataRelationshipsCaptures**](AuthorizationDataRelationshipsCaptures.md) |  | [optional] 
**Voids** | Pointer to [**AuthorizationDataRelationshipsVoids**](AuthorizationDataRelationshipsVoids.md) |  | [optional] 

## Methods

### NewAuthorizationDataRelationships

`func NewAuthorizationDataRelationships() *AuthorizationDataRelationships`

NewAuthorizationDataRelationships instantiates a new AuthorizationDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationDataRelationshipsWithDefaults

`func NewAuthorizationDataRelationshipsWithDefaults() *AuthorizationDataRelationships`

NewAuthorizationDataRelationshipsWithDefaults instantiates a new AuthorizationDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *AuthorizationDataRelationships) GetOrder() AdyenPaymentDataRelationshipsOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AuthorizationDataRelationships) GetOrderOk() (*AdyenPaymentDataRelationshipsOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AuthorizationDataRelationships) SetOrder(v AdyenPaymentDataRelationshipsOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AuthorizationDataRelationships) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPaymentSource

`func (o *AuthorizationDataRelationships) GetPaymentSource() AuthorizationDataRelationshipsPaymentSource`

GetPaymentSource returns the PaymentSource field if non-nil, zero value otherwise.

### GetPaymentSourceOk

`func (o *AuthorizationDataRelationships) GetPaymentSourceOk() (*AuthorizationDataRelationshipsPaymentSource, bool)`

GetPaymentSourceOk returns a tuple with the PaymentSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentSource

`func (o *AuthorizationDataRelationships) SetPaymentSource(v AuthorizationDataRelationshipsPaymentSource)`

SetPaymentSource sets PaymentSource field to given value.

### HasPaymentSource

`func (o *AuthorizationDataRelationships) HasPaymentSource() bool`

HasPaymentSource returns a boolean if a field has been set.

### GetAttachments

`func (o *AuthorizationDataRelationships) GetAttachments() AuthorizationDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *AuthorizationDataRelationships) GetAttachmentsOk() (*AuthorizationDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *AuthorizationDataRelationships) SetAttachments(v AuthorizationDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *AuthorizationDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetEvents

`func (o *AuthorizationDataRelationships) GetEvents() AddressDataRelationshipsEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AuthorizationDataRelationships) GetEventsOk() (*AddressDataRelationshipsEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AuthorizationDataRelationships) SetEvents(v AddressDataRelationshipsEvents)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AuthorizationDataRelationships) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetVersions

`func (o *AuthorizationDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AuthorizationDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AuthorizationDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AuthorizationDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetCaptures

`func (o *AuthorizationDataRelationships) GetCaptures() AuthorizationDataRelationshipsCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *AuthorizationDataRelationships) GetCapturesOk() (*AuthorizationDataRelationshipsCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *AuthorizationDataRelationships) SetCaptures(v AuthorizationDataRelationshipsCaptures)`

SetCaptures sets Captures field to given value.

### HasCaptures

`func (o *AuthorizationDataRelationships) HasCaptures() bool`

HasCaptures returns a boolean if a field has been set.

### GetVoids

`func (o *AuthorizationDataRelationships) GetVoids() AuthorizationDataRelationshipsVoids`

GetVoids returns the Voids field if non-nil, zero value otherwise.

### GetVoidsOk

`func (o *AuthorizationDataRelationships) GetVoidsOk() (*AuthorizationDataRelationshipsVoids, bool)`

GetVoidsOk returns a tuple with the Voids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoids

`func (o *AuthorizationDataRelationships) SetVoids(v AuthorizationDataRelationshipsVoids)`

SetVoids sets Voids field to given value.

### HasVoids

`func (o *AuthorizationDataRelationships) HasVoids() bool`

HasVoids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


