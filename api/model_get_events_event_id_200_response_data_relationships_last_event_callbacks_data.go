/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 4.1.3
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData{}

// GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData struct for GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData
type GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData instantiates a new GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData() *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData {
	this := GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData{}
	return &this
}

// NewGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksDataWithDefaults instantiates a new GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksDataWithDefaults() *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData {
	this := GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) SetId(v interface{}) {
	o.Id = v
}

func (o GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData struct {
	value *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData
	isSet bool
}

func (v NullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) Get() *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData {
	return v.value
}

func (v *NullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) Set(val *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) {
	v.value = val
	v.isSet = true
}

func (v NullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) IsSet() bool {
	return v.isSet
}

func (v *NullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData(val *GETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) *NullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData {
	return &NullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData{value: val, isSet: true}
}

func (v NullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETEventsEventId200ResponseDataRelationshipsLastEventCallbacksData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}