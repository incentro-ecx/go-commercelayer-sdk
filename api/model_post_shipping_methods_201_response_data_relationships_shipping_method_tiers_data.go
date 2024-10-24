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

// checks if the POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData{}

// POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData struct for POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData
type POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData instantiates a new POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData() *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData {
	this := POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData{}
	return &this
}

// NewPOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersDataWithDefaults instantiates a new POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersDataWithDefaults() *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData {
	this := POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData struct {
	value *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData
	isSet bool
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) Get() *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData {
	return v.value
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) Set(val *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData(val *POSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData {
	return &NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData{value: val, isSet: true}
}

func (v NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShippingMethods201ResponseDataRelationshipsShippingMethodTiersData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
