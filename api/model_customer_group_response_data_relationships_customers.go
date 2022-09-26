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

// CustomerGroupResponseDataRelationshipsCustomers struct for CustomerGroupResponseDataRelationshipsCustomers
type CustomerGroupResponseDataRelationshipsCustomers struct {
	Links *AddressResponseDataRelationshipsGeocoderLinks             `json:"links,omitempty"`
	Data  []CustomerGroupResponseDataRelationshipsCustomersDataInner `json:"data,omitempty"`
}

// NewCustomerGroupResponseDataRelationshipsCustomers instantiates a new CustomerGroupResponseDataRelationshipsCustomers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerGroupResponseDataRelationshipsCustomers() *CustomerGroupResponseDataRelationshipsCustomers {
	this := CustomerGroupResponseDataRelationshipsCustomers{}
	return &this
}

// NewCustomerGroupResponseDataRelationshipsCustomersWithDefaults instantiates a new CustomerGroupResponseDataRelationshipsCustomers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerGroupResponseDataRelationshipsCustomersWithDefaults() *CustomerGroupResponseDataRelationshipsCustomers {
	this := CustomerGroupResponseDataRelationshipsCustomers{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CustomerGroupResponseDataRelationshipsCustomers) GetLinks() AddressResponseDataRelationshipsGeocoderLinks {
	if o == nil || o.Links == nil {
		var ret AddressResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponseDataRelationshipsCustomers) GetLinksOk() (*AddressResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CustomerGroupResponseDataRelationshipsCustomers) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AddressResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *CustomerGroupResponseDataRelationshipsCustomers) SetLinks(v AddressResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerGroupResponseDataRelationshipsCustomers) GetData() []CustomerGroupResponseDataRelationshipsCustomersDataInner {
	if o == nil || o.Data == nil {
		var ret []CustomerGroupResponseDataRelationshipsCustomersDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerGroupResponseDataRelationshipsCustomers) GetDataOk() ([]CustomerGroupResponseDataRelationshipsCustomersDataInner, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerGroupResponseDataRelationshipsCustomers) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []CustomerGroupResponseDataRelationshipsCustomersDataInner and assigns it to the Data field.
func (o *CustomerGroupResponseDataRelationshipsCustomers) SetData(v []CustomerGroupResponseDataRelationshipsCustomersDataInner) {
	o.Data = v
}

func (o CustomerGroupResponseDataRelationshipsCustomers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerGroupResponseDataRelationshipsCustomers struct {
	value *CustomerGroupResponseDataRelationshipsCustomers
	isSet bool
}

func (v NullableCustomerGroupResponseDataRelationshipsCustomers) Get() *CustomerGroupResponseDataRelationshipsCustomers {
	return v.value
}

func (v *NullableCustomerGroupResponseDataRelationshipsCustomers) Set(val *CustomerGroupResponseDataRelationshipsCustomers) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerGroupResponseDataRelationshipsCustomers) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerGroupResponseDataRelationshipsCustomers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerGroupResponseDataRelationshipsCustomers(val *CustomerGroupResponseDataRelationshipsCustomers) *NullableCustomerGroupResponseDataRelationshipsCustomers {
	return &NullableCustomerGroupResponseDataRelationshipsCustomers{value: val, isSet: true}
}

func (v NullableCustomerGroupResponseDataRelationshipsCustomers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerGroupResponseDataRelationshipsCustomers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}