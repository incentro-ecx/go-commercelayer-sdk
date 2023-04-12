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

// checks if the POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData{}

// POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData struct for POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData
type POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData instantiates a new POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData() *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData {
	this := POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData{}
	return &this
}

// NewPOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationDataWithDefaults instantiates a new POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationDataWithDefaults() *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData {
	this := POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData struct {
	value *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData
	isSet bool
}

func (v NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) Get() *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData {
	return v.value
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) Set(val *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData(val *POSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) *NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData {
	return &NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData{value: val, isSet: true}
}

func (v NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockTransfers201ResponseDataRelationshipsDestinationStockLocationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
