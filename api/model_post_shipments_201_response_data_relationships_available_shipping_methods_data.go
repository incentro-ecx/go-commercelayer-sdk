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

// checks if the POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData{}

// POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData struct for POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData
type POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData instantiates a new POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData() *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData {
	this := POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData{}
	return &this
}

// NewPOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsDataWithDefaults instantiates a new POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsDataWithDefaults() *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData {
	this := POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData struct {
	value *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData
	isSet bool
}

func (v NullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) Get() *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData {
	return v.value
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) Set(val *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData(val *POSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) *NullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData {
	return &NullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData{value: val, isSet: true}
}

func (v NullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTShipments201ResponseDataRelationshipsAvailableShippingMethodsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}