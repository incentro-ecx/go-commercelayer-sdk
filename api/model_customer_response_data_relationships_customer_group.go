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

// CustomerResponseDataRelationshipsCustomerGroup struct for CustomerResponseDataRelationshipsCustomerGroup
type CustomerResponseDataRelationshipsCustomerGroup struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks      `json:"links,omitempty"`
	Data  *CustomerResponseDataRelationshipsCustomerGroupData `json:"data,omitempty"`
}

// NewCustomerResponseDataRelationshipsCustomerGroup instantiates a new CustomerResponseDataRelationshipsCustomerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerResponseDataRelationshipsCustomerGroup() *CustomerResponseDataRelationshipsCustomerGroup {
	this := CustomerResponseDataRelationshipsCustomerGroup{}
	return &this
}

// NewCustomerResponseDataRelationshipsCustomerGroupWithDefaults instantiates a new CustomerResponseDataRelationshipsCustomerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerResponseDataRelationshipsCustomerGroupWithDefaults() *CustomerResponseDataRelationshipsCustomerGroup {
	this := CustomerResponseDataRelationshipsCustomerGroup{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomerResponseDataRelationshipsCustomerGroup) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponseDataRelationshipsCustomerGroup) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomerResponseDataRelationshipsCustomerGroup) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *CustomerResponseDataRelationshipsCustomerGroup) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerResponseDataRelationshipsCustomerGroup) GetData() CustomerResponseDataRelationshipsCustomerGroupData {
	if o == nil || o.Data == nil {
		var ret CustomerResponseDataRelationshipsCustomerGroupData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerResponseDataRelationshipsCustomerGroup) GetDataOk() (*CustomerResponseDataRelationshipsCustomerGroupData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerResponseDataRelationshipsCustomerGroup) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CustomerResponseDataRelationshipsCustomerGroupData and assigns it to the Data field.
func (o *CustomerResponseDataRelationshipsCustomerGroup) SetData(v CustomerResponseDataRelationshipsCustomerGroupData) {
	o.Data = &v
}

func (o CustomerResponseDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerResponseDataRelationshipsCustomerGroup struct {
	value *CustomerResponseDataRelationshipsCustomerGroup
	isSet bool
}

func (v NullableCustomerResponseDataRelationshipsCustomerGroup) Get() *CustomerResponseDataRelationshipsCustomerGroup {
	return v.value
}

func (v *NullableCustomerResponseDataRelationshipsCustomerGroup) Set(val *CustomerResponseDataRelationshipsCustomerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerResponseDataRelationshipsCustomerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerResponseDataRelationshipsCustomerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerResponseDataRelationshipsCustomerGroup(val *CustomerResponseDataRelationshipsCustomerGroup) *NullableCustomerResponseDataRelationshipsCustomerGroup {
	return &NullableCustomerResponseDataRelationshipsCustomerGroup{value: val, isSet: true}
}

func (v NullableCustomerResponseDataRelationshipsCustomerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerResponseDataRelationshipsCustomerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}