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

// checks if the EventData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventData{}

// EventData struct for EventData
type EventData struct {
	// The resource's type
	Type          interface{}                               `json:"type"`
	Attributes    GETEventsEventId200ResponseDataAttributes `json:"attributes"`
	Relationships *EventDataRelationships                   `json:"relationships,omitempty"`
}

// NewEventData instantiates a new EventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventData(type_ interface{}, attributes GETEventsEventId200ResponseDataAttributes) *EventData {
	this := EventData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewEventDataWithDefaults instantiates a new EventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDataWithDefaults() *EventData {
	this := EventData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *EventData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *EventData) GetAttributes() GETEventsEventId200ResponseDataAttributes {
	if o == nil {
		var ret GETEventsEventId200ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EventData) GetAttributesOk() (*GETEventsEventId200ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *EventData) SetAttributes(v GETEventsEventId200ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *EventData) GetRelationships() EventDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret EventDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventData) GetRelationshipsOk() (*EventDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *EventData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given EventDataRelationships and assigns it to the Relationships field.
func (o *EventData) SetRelationships(v EventDataRelationships) {
	o.Relationships = &v
}

func (o EventData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableEventData struct {
	value *EventData
	isSet bool
}

func (v NullableEventData) Get() *EventData {
	return v.value
}

func (v *NullableEventData) Set(val *EventData) {
	v.value = val
	v.isSet = true
}

func (v NullableEventData) IsSet() bool {
	return v.isSet
}

func (v *NullableEventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventData(val *EventData) *NullableEventData {
	return &NullableEventData{value: val, isSet: true}
}

func (v NullableEventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
