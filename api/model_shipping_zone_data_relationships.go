/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ShippingZoneDataRelationships struct for ShippingZoneDataRelationships
type ShippingZoneDataRelationships struct {
	Attachments *AvalaraAccountDataRelationshipsAttachments `json:"attachments,omitempty"`
}

// NewShippingZoneDataRelationships instantiates a new ShippingZoneDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShippingZoneDataRelationships() *ShippingZoneDataRelationships {
	this := ShippingZoneDataRelationships{}
	return &this
}

// NewShippingZoneDataRelationshipsWithDefaults instantiates a new ShippingZoneDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShippingZoneDataRelationshipsWithDefaults() *ShippingZoneDataRelationships {
	this := ShippingZoneDataRelationships{}
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ShippingZoneDataRelationships) GetAttachments() AvalaraAccountDataRelationshipsAttachments {
	if o == nil || o.Attachments == nil {
		var ret AvalaraAccountDataRelationshipsAttachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShippingZoneDataRelationships) GetAttachmentsOk() (*AvalaraAccountDataRelationshipsAttachments, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ShippingZoneDataRelationships) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given AvalaraAccountDataRelationshipsAttachments and assigns it to the Attachments field.
func (o *ShippingZoneDataRelationships) SetAttachments(v AvalaraAccountDataRelationshipsAttachments) {
	o.Attachments = &v
}

func (o ShippingZoneDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	return json.Marshal(toSerialize)
}

type NullableShippingZoneDataRelationships struct {
	value *ShippingZoneDataRelationships
	isSet bool
}

func (v NullableShippingZoneDataRelationships) Get() *ShippingZoneDataRelationships {
	return v.value
}

func (v *NullableShippingZoneDataRelationships) Set(val *ShippingZoneDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableShippingZoneDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableShippingZoneDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShippingZoneDataRelationships(val *ShippingZoneDataRelationships) *NullableShippingZoneDataRelationships {
	return &NullableShippingZoneDataRelationships{value: val, isSet: true}
}

func (v NullableShippingZoneDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShippingZoneDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}