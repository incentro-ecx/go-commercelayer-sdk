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

// checks if the POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData{}

// POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData struct for POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData
type POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData struct {
	// The resource's type
	Type interface{} `json:"type,omitempty"`
	// The resource ID
	Id interface{} `json:"id,omitempty"`
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData() *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData{}
	return &this
}

// NewPOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesDataWithDefaults instantiates a new POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesDataWithDefaults() *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData {
	this := POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) SetType(v interface{}) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) SetId(v interface{}) {
	o.Id = v
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData struct {
	value *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData
	isSet bool
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) Get() *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData {
	return v.value
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) Set(val *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData(val *POSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData {
	return &NullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData{value: val, isSet: true}
}

func (v NullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrderSubscriptions201ResponseDataRelationshipsRecurringOrderCopiesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
