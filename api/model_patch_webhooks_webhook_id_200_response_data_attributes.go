/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHWebhooksWebhookId200ResponseDataAttributes struct for PATCHWebhooksWebhookId200ResponseDataAttributes
type PATCHWebhooksWebhookId200ResponseDataAttributes struct {
	// Unique name for the webhook.
	Name *string `json:"name,omitempty"`
	// The identifier of the resource/event that will trigger the webhook.
	Topic *string `json:"topic,omitempty"`
	// URI where the webhook subscription should send the POST request when the event occurs.
	CallbackUrl *string `json:"callback_url,omitempty"`
	// List of related resources that should be included in the webhook body.
	IncludeResources []string `json:"include_resources,omitempty"`
	// Send this attribute if you want to reset the circuit breaker associated to this webhook to 'closed' state and zero failures count.
	ResetCircuit *bool `json:"_reset_circuit,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHWebhooksWebhookId200ResponseDataAttributes instantiates a new PATCHWebhooksWebhookId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHWebhooksWebhookId200ResponseDataAttributes() *PATCHWebhooksWebhookId200ResponseDataAttributes {
	this := PATCHWebhooksWebhookId200ResponseDataAttributes{}
	return &this
}

// NewPATCHWebhooksWebhookId200ResponseDataAttributesWithDefaults instantiates a new PATCHWebhooksWebhookId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHWebhooksWebhookId200ResponseDataAttributesWithDefaults() *PATCHWebhooksWebhookId200ResponseDataAttributes {
	this := PATCHWebhooksWebhookId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetTopic(v string) {
	o.Topic = &v
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrl() string {
	if o == nil || o.CallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetCallbackUrlOk() (*string, bool) {
	if o == nil || o.CallbackUrl == nil {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasCallbackUrl() bool {
	if o != nil && o.CallbackUrl != nil {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetIncludeResources returns the IncludeResources field value if set, zero value otherwise.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetIncludeResources() []string {
	if o == nil || o.IncludeResources == nil {
		var ret []string
		return ret
	}
	return o.IncludeResources
}

// GetIncludeResourcesOk returns a tuple with the IncludeResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetIncludeResourcesOk() ([]string, bool) {
	if o == nil || o.IncludeResources == nil {
		return nil, false
	}
	return o.IncludeResources, true
}

// HasIncludeResources returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasIncludeResources() bool {
	if o != nil && o.IncludeResources != nil {
		return true
	}

	return false
}

// SetIncludeResources gets a reference to the given []string and assigns it to the IncludeResources field.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetIncludeResources(v []string) {
	o.IncludeResources = v
}

// GetResetCircuit returns the ResetCircuit field value if set, zero value otherwise.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetResetCircuit() bool {
	if o == nil || o.ResetCircuit == nil {
		var ret bool
		return ret
	}
	return *o.ResetCircuit
}

// GetResetCircuitOk returns a tuple with the ResetCircuit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetResetCircuitOk() (*bool, bool) {
	if o == nil || o.ResetCircuit == nil {
		return nil, false
	}
	return o.ResetCircuit, true
}

// HasResetCircuit returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasResetCircuit() bool {
	if o != nil && o.ResetCircuit != nil {
		return true
	}

	return false
}

// SetResetCircuit gets a reference to the given bool and assigns it to the ResetCircuit field.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetResetCircuit(v bool) {
	o.ResetCircuit = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHWebhooksWebhookId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHWebhooksWebhookId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.CallbackUrl != nil {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if o.IncludeResources != nil {
		toSerialize["include_resources"] = o.IncludeResources
	}
	if o.ResetCircuit != nil {
		toSerialize["_reset_circuit"] = o.ResetCircuit
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullablePATCHWebhooksWebhookId200ResponseDataAttributes struct {
	value *PATCHWebhooksWebhookId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHWebhooksWebhookId200ResponseDataAttributes) Get() *PATCHWebhooksWebhookId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHWebhooksWebhookId200ResponseDataAttributes) Set(val *PATCHWebhooksWebhookId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHWebhooksWebhookId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHWebhooksWebhookId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHWebhooksWebhookId200ResponseDataAttributes(val *PATCHWebhooksWebhookId200ResponseDataAttributes) *NullablePATCHWebhooksWebhookId200ResponseDataAttributes {
	return &NullablePATCHWebhooksWebhookId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHWebhooksWebhookId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHWebhooksWebhookId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
