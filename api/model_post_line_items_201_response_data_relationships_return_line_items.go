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

// checks if the POSTLineItems201ResponseDataRelationshipsReturnLineItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTLineItems201ResponseDataRelationshipsReturnLineItems{}

// POSTLineItems201ResponseDataRelationshipsReturnLineItems struct for POSTLineItems201ResponseDataRelationshipsReturnLineItems
type POSTLineItems201ResponseDataRelationshipsReturnLineItems struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks       `json:"links,omitempty"`
	Data  *POSTLineItems201ResponseDataRelationshipsReturnLineItemsData `json:"data,omitempty"`
}

// NewPOSTLineItems201ResponseDataRelationshipsReturnLineItems instantiates a new POSTLineItems201ResponseDataRelationshipsReturnLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTLineItems201ResponseDataRelationshipsReturnLineItems() *POSTLineItems201ResponseDataRelationshipsReturnLineItems {
	this := POSTLineItems201ResponseDataRelationshipsReturnLineItems{}
	return &this
}

// NewPOSTLineItems201ResponseDataRelationshipsReturnLineItemsWithDefaults instantiates a new POSTLineItems201ResponseDataRelationshipsReturnLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTLineItems201ResponseDataRelationshipsReturnLineItemsWithDefaults() *POSTLineItems201ResponseDataRelationshipsReturnLineItems {
	this := POSTLineItems201ResponseDataRelationshipsReturnLineItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsReturnLineItems) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsReturnLineItems) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsReturnLineItems) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTLineItems201ResponseDataRelationshipsReturnLineItems) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTLineItems201ResponseDataRelationshipsReturnLineItems) GetData() POSTLineItems201ResponseDataRelationshipsReturnLineItemsData {
	if o == nil || IsNil(o.Data) {
		var ret POSTLineItems201ResponseDataRelationshipsReturnLineItemsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTLineItems201ResponseDataRelationshipsReturnLineItems) GetDataOk() (*POSTLineItems201ResponseDataRelationshipsReturnLineItemsData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTLineItems201ResponseDataRelationshipsReturnLineItems) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTLineItems201ResponseDataRelationshipsReturnLineItemsData and assigns it to the Data field.
func (o *POSTLineItems201ResponseDataRelationshipsReturnLineItems) SetData(v POSTLineItems201ResponseDataRelationshipsReturnLineItemsData) {
	o.Data = &v
}

func (o POSTLineItems201ResponseDataRelationshipsReturnLineItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTLineItems201ResponseDataRelationshipsReturnLineItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems struct {
	value *POSTLineItems201ResponseDataRelationshipsReturnLineItems
	isSet bool
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems) Get() *POSTLineItems201ResponseDataRelationshipsReturnLineItems {
	return v.value
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems) Set(val *POSTLineItems201ResponseDataRelationshipsReturnLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems(val *POSTLineItems201ResponseDataRelationshipsReturnLineItems) *NullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems {
	return &NullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems{value: val, isSet: true}
}

func (v NullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTLineItems201ResponseDataRelationshipsReturnLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
