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

// checks if the POSTPriceLists201ResponseDataRelationshipsPricesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTPriceLists201ResponseDataRelationshipsPricesData{}

// POSTPriceLists201ResponseDataRelationshipsPricesData struct for POSTPriceLists201ResponseDataRelationshipsPricesData
type POSTPriceLists201ResponseDataRelationshipsPricesData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTPriceLists201ResponseDataRelationshipsPricesData instantiates a new POSTPriceLists201ResponseDataRelationshipsPricesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTPriceLists201ResponseDataRelationshipsPricesData() *POSTPriceLists201ResponseDataRelationshipsPricesData {
	this := POSTPriceLists201ResponseDataRelationshipsPricesData{}
	return &this
}

// NewPOSTPriceLists201ResponseDataRelationshipsPricesDataWithDefaults instantiates a new POSTPriceLists201ResponseDataRelationshipsPricesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTPriceLists201ResponseDataRelationshipsPricesDataWithDefaults() *POSTPriceLists201ResponseDataRelationshipsPricesData {
	this := POSTPriceLists201ResponseDataRelationshipsPricesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPriceLists201ResponseDataRelationshipsPricesData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceLists201ResponseDataRelationshipsPricesData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTPriceLists201ResponseDataRelationshipsPricesData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTPriceLists201ResponseDataRelationshipsPricesData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTPriceLists201ResponseDataRelationshipsPricesData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTPriceLists201ResponseDataRelationshipsPricesData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTPriceLists201ResponseDataRelationshipsPricesData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTPriceLists201ResponseDataRelationshipsPricesData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTPriceLists201ResponseDataRelationshipsPricesData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTPriceLists201ResponseDataRelationshipsPricesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTPriceLists201ResponseDataRelationshipsPricesData struct {
	value *POSTPriceLists201ResponseDataRelationshipsPricesData
	isSet bool
}

func (v NullablePOSTPriceLists201ResponseDataRelationshipsPricesData) Get() *POSTPriceLists201ResponseDataRelationshipsPricesData {
	return v.value
}

func (v *NullablePOSTPriceLists201ResponseDataRelationshipsPricesData) Set(val *POSTPriceLists201ResponseDataRelationshipsPricesData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTPriceLists201ResponseDataRelationshipsPricesData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTPriceLists201ResponseDataRelationshipsPricesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTPriceLists201ResponseDataRelationshipsPricesData(val *POSTPriceLists201ResponseDataRelationshipsPricesData) *NullablePOSTPriceLists201ResponseDataRelationshipsPricesData {
	return &NullablePOSTPriceLists201ResponseDataRelationshipsPricesData{value: val, isSet: true}
}

func (v NullablePOSTPriceLists201ResponseDataRelationshipsPricesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTPriceLists201ResponseDataRelationshipsPricesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
