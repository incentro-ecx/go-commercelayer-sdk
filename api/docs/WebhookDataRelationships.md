# WebhookDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastEventCallbacks** | Pointer to [**EventDataRelationshipsLastEventCallbacks**](EventDataRelationshipsLastEventCallbacks.md) |  | [optional] 
**Versions** | Pointer to [**AddressDataRelationshipsVersions**](AddressDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewWebhookDataRelationships

`func NewWebhookDataRelationships() *WebhookDataRelationships`

NewWebhookDataRelationships instantiates a new WebhookDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDataRelationshipsWithDefaults

`func NewWebhookDataRelationshipsWithDefaults() *WebhookDataRelationships`

NewWebhookDataRelationshipsWithDefaults instantiates a new WebhookDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastEventCallbacks

`func (o *WebhookDataRelationships) GetLastEventCallbacks() EventDataRelationshipsLastEventCallbacks`

GetLastEventCallbacks returns the LastEventCallbacks field if non-nil, zero value otherwise.

### GetLastEventCallbacksOk

`func (o *WebhookDataRelationships) GetLastEventCallbacksOk() (*EventDataRelationshipsLastEventCallbacks, bool)`

GetLastEventCallbacksOk returns a tuple with the LastEventCallbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEventCallbacks

`func (o *WebhookDataRelationships) SetLastEventCallbacks(v EventDataRelationshipsLastEventCallbacks)`

SetLastEventCallbacks sets LastEventCallbacks field to given value.

### HasLastEventCallbacks

`func (o *WebhookDataRelationships) HasLastEventCallbacks() bool`

HasLastEventCallbacks returns a boolean if a field has been set.

### GetVersions

`func (o *WebhookDataRelationships) GetVersions() AddressDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *WebhookDataRelationships) GetVersionsOk() (*AddressDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *WebhookDataRelationships) SetVersions(v AddressDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *WebhookDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


