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

// checks if the POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData{}

// POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData struct for POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData
type POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData instantiates a new POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData() *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData {
	this := POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData{}
	return &this
}

// NewPOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodDataWithDefaults instantiates a new POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodDataWithDefaults() *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData {
	this := POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData struct {
	value *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData
	isSet bool
}

func (v NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) Get() *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData {
	return v.value
}

func (v *NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) Set(val *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData(val *POSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) *NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData {
	return &NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData{value: val, isSet: true}
}

func (v NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTDeliveryLeadTimes201ResponseDataRelationshipsShippingMethodData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}