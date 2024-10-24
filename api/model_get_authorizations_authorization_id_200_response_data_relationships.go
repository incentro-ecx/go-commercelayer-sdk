/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 7.3.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETAuthorizationsAuthorizationId200ResponseDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETAuthorizationsAuthorizationId200ResponseDataRelationships{}

// GETAuthorizationsAuthorizationId200ResponseDataRelationships struct for GETAuthorizationsAuthorizationId200ResponseDataRelationships
type GETAuthorizationsAuthorizationId200ResponseDataRelationships struct {
	Order       *POSTAdyenPayments201ResponseDataRelationshipsOrder                      `json:"order,omitempty"`
	Attachments *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments `json:"attachments,omitempty"`
	Events      *POSTAddresses201ResponseDataRelationshipsEvents                         `json:"events,omitempty"`
	Versions    *POSTAddresses201ResponseDataRelationshipsVersions                       `json:"versions,omitempty"`
	Captures    *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures    `json:"captures,omitempty"`
	Voids       *GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids       `json:"voids,omitempty"`
}

// NewGETAuthorizationsAuthorizationId200ResponseDataRelationships instantiates a new GETAuthorizationsAuthorizationId200ResponseDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETAuthorizationsAuthorizationId200ResponseDataRelationships() *GETAuthorizationsAuthorizationId200ResponseDataRelationships {
	this := GETAuthorizationsAuthorizationId200ResponseDataRelationships{}
	return &this
}

// NewGETAuthorizationsAuthorizationId200ResponseDataRelationshipsWithDefaults instantiates a new GETAuthorizationsAuthorizationId200ResponseDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETAuthorizationsAuthorizationId200ResponseDataRelationshipsWithDefaults() *GETAuthorizationsAuthorizationId200ResponseDataRelationships {
	this := GETAuthorizationsAuthorizationId200ResponseDataRelationships{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetOrder() POSTAdyenPayments201ResponseDataRelationshipsOrder {
	if o == nil || IsNil(o.Order) {
		var ret POSTAdyenPayments201ResponseDataRelationshipsOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetOrderOk() (*POSTAdyenPayments201ResponseDataRelationshipsOrder, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given POSTAdyenPayments201ResponseDataRelationshipsOrder and assigns it to the Order field.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetOrder(v POSTAdyenPayments201ResponseDataRelationshipsOrder) {
	o.Order = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments {
	if o == nil || IsNil(o.Attachments) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments) {
	o.Attachments = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetEvents() POSTAddresses201ResponseDataRelationshipsEvents {
	if o == nil || IsNil(o.Events) {
		var ret POSTAddresses201ResponseDataRelationshipsEvents
		return ret
	}
	return *o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetEventsOk() (*POSTAddresses201ResponseDataRelationshipsEvents, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given POSTAddresses201ResponseDataRelationshipsEvents and assigns it to the Events field.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetEvents(v POSTAddresses201ResponseDataRelationshipsEvents) {
	o.Events = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions {
	if o == nil || IsNil(o.Versions) {
		var ret POSTAddresses201ResponseDataRelationshipsVersions
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given POSTAddresses201ResponseDataRelationshipsVersions and assigns it to the Versions field.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions) {
	o.Versions = &v
}

// GetCaptures returns the Captures field value if set, zero value otherwise.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetCaptures() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures {
	if o == nil || IsNil(o.Captures) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures
		return ret
	}
	return *o.Captures
}

// GetCapturesOk returns a tuple with the Captures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetCapturesOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures, bool) {
	if o == nil || IsNil(o.Captures) {
		return nil, false
	}
	return o.Captures, true
}

// HasCaptures returns a boolean if a field has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasCaptures() bool {
	if o != nil && !IsNil(o.Captures) {
		return true
	}

	return false
}

// SetCaptures gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures and assigns it to the Captures field.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetCaptures(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsCaptures) {
	o.Captures = &v
}

// GetVoids returns the Voids field value if set, zero value otherwise.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetVoids() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids {
	if o == nil || IsNil(o.Voids) {
		var ret GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids
		return ret
	}
	return *o.Voids
}

// GetVoidsOk returns a tuple with the Voids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) GetVoidsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids, bool) {
	if o == nil || IsNil(o.Voids) {
		return nil, false
	}
	return o.Voids, true
}

// HasVoids returns a boolean if a field has been set.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) HasVoids() bool {
	if o != nil && !IsNil(o.Voids) {
		return true
	}

	return false
}

// SetVoids gets a reference to the given GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids and assigns it to the Voids field.
func (o *GETAuthorizationsAuthorizationId200ResponseDataRelationships) SetVoids(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsVoids) {
	o.Voids = &v
}

func (o GETAuthorizationsAuthorizationId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETAuthorizationsAuthorizationId200ResponseDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.Captures) {
		toSerialize["captures"] = o.Captures
	}
	if !IsNil(o.Voids) {
		toSerialize["voids"] = o.Voids
	}
	return toSerialize, nil
}

type NullableGETAuthorizationsAuthorizationId200ResponseDataRelationships struct {
	value *GETAuthorizationsAuthorizationId200ResponseDataRelationships
	isSet bool
}

func (v NullableGETAuthorizationsAuthorizationId200ResponseDataRelationships) Get() *GETAuthorizationsAuthorizationId200ResponseDataRelationships {
	return v.value
}

func (v *NullableGETAuthorizationsAuthorizationId200ResponseDataRelationships) Set(val *GETAuthorizationsAuthorizationId200ResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGETAuthorizationsAuthorizationId200ResponseDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGETAuthorizationsAuthorizationId200ResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETAuthorizationsAuthorizationId200ResponseDataRelationships(val *GETAuthorizationsAuthorizationId200ResponseDataRelationships) *NullableGETAuthorizationsAuthorizationId200ResponseDataRelationships {
	return &NullableGETAuthorizationsAuthorizationId200ResponseDataRelationships{value: val, isSet: true}
}

func (v NullableGETAuthorizationsAuthorizationId200ResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETAuthorizationsAuthorizationId200ResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
