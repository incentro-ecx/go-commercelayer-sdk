/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// ReturnResponseDataRelationshipsReturnLineItems struct for ReturnResponseDataRelationshipsReturnLineItems
type ReturnResponseDataRelationshipsReturnLineItems struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks            `json:"links,omitempty"`
	Data  []ReturnResponseDataRelationshipsReturnLineItemsDataInner `json:"data,omitempty"`
}

// NewReturnResponseDataRelationshipsReturnLineItems instantiates a new ReturnResponseDataRelationshipsReturnLineItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnResponseDataRelationshipsReturnLineItems() *ReturnResponseDataRelationshipsReturnLineItems {
	this := ReturnResponseDataRelationshipsReturnLineItems{}
	return &this
}

// NewReturnResponseDataRelationshipsReturnLineItemsWithDefaults instantiates a new ReturnResponseDataRelationshipsReturnLineItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnResponseDataRelationshipsReturnLineItemsWithDefaults() *ReturnResponseDataRelationshipsReturnLineItems {
	this := ReturnResponseDataRelationshipsReturnLineItems{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReturnResponseDataRelationshipsReturnLineItems) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnResponseDataRelationshipsReturnLineItems) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReturnResponseDataRelationshipsReturnLineItems) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *ReturnResponseDataRelationshipsReturnLineItems) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ReturnResponseDataRelationshipsReturnLineItems) GetData() []ReturnResponseDataRelationshipsReturnLineItemsDataInner {
	if o == nil || o.Data == nil {
		var ret []ReturnResponseDataRelationshipsReturnLineItemsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnResponseDataRelationshipsReturnLineItems) GetDataOk() ([]ReturnResponseDataRelationshipsReturnLineItemsDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ReturnResponseDataRelationshipsReturnLineItems) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ReturnResponseDataRelationshipsReturnLineItemsDataInner and assigns it to the Data field.
func (o *ReturnResponseDataRelationshipsReturnLineItems) SetData(v []ReturnResponseDataRelationshipsReturnLineItemsDataInner) {
	o.Data = v
}

func (o ReturnResponseDataRelationshipsReturnLineItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableReturnResponseDataRelationshipsReturnLineItems struct {
	value *ReturnResponseDataRelationshipsReturnLineItems
	isSet bool
}

func (v NullableReturnResponseDataRelationshipsReturnLineItems) Get() *ReturnResponseDataRelationshipsReturnLineItems {
	return v.value
}

func (v *NullableReturnResponseDataRelationshipsReturnLineItems) Set(val *ReturnResponseDataRelationshipsReturnLineItems) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnResponseDataRelationshipsReturnLineItems) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnResponseDataRelationshipsReturnLineItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnResponseDataRelationshipsReturnLineItems(val *ReturnResponseDataRelationshipsReturnLineItems) *NullableReturnResponseDataRelationshipsReturnLineItems {
	return &NullableReturnResponseDataRelationshipsReturnLineItems{value: val, isSet: true}
}

func (v NullableReturnResponseDataRelationshipsReturnLineItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnResponseDataRelationshipsReturnLineItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}