/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// WebhookResponseDataRelationships struct for WebhookResponseDataRelationships
type WebhookResponseDataRelationships struct {
	LastEventCallbacks *EventResponseDataRelationshipsLastEventCallbacks `json:"last_event_callbacks,omitempty"`
}

// NewWebhookResponseDataRelationships instantiates a new WebhookResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponseDataRelationships() *WebhookResponseDataRelationships {
	this := WebhookResponseDataRelationships{}
	return &this
}

// NewWebhookResponseDataRelationshipsWithDefaults instantiates a new WebhookResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseDataRelationshipsWithDefaults() *WebhookResponseDataRelationships {
	this := WebhookResponseDataRelationships{}
	return &this
}

// GetLastEventCallbacks returns the LastEventCallbacks field value if set, zero value otherwise.
func (o *WebhookResponseDataRelationships) GetLastEventCallbacks() EventResponseDataRelationshipsLastEventCallbacks {
	if o == nil || o.LastEventCallbacks == nil {
		var ret EventResponseDataRelationshipsLastEventCallbacks
		return ret
	}
	return *o.LastEventCallbacks
}

// GetLastEventCallbacksOk returns a tuple with the LastEventCallbacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseDataRelationships) GetLastEventCallbacksOk() (*EventResponseDataRelationshipsLastEventCallbacks, bool) {
	if o == nil || o.LastEventCallbacks == nil {
		return nil, false
	}
	return o.LastEventCallbacks, true
}

// HasLastEventCallbacks returns a boolean if a field has been set.
func (o *WebhookResponseDataRelationships) HasLastEventCallbacks() bool {
	if o != nil && o.LastEventCallbacks != nil {
		return true
	}

	return false
}

// SetLastEventCallbacks gets a reference to the given EventResponseDataRelationshipsLastEventCallbacks and assigns it to the LastEventCallbacks field.
func (o *WebhookResponseDataRelationships) SetLastEventCallbacks(v EventResponseDataRelationshipsLastEventCallbacks) {
	o.LastEventCallbacks = &v
}

func (o WebhookResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastEventCallbacks != nil {
		toSerialize["last_event_callbacks"] = o.LastEventCallbacks
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookResponseDataRelationships struct {
	value *WebhookResponseDataRelationships
	isSet bool
}

func (v NullableWebhookResponseDataRelationships) Get() *WebhookResponseDataRelationships {
	return v.value
}

func (v *NullableWebhookResponseDataRelationships) Set(val *WebhookResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponseDataRelationships(val *WebhookResponseDataRelationships) *NullableWebhookResponseDataRelationships {
	return &NullableWebhookResponseDataRelationships{value: val, isSet: true}
}

func (v NullableWebhookResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}