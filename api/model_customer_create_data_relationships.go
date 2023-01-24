/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.2.0
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CustomerCreateDataRelationships struct for CustomerCreateDataRelationships
type CustomerCreateDataRelationships struct {
	CustomerGroup *CustomerCreateDataRelationshipsCustomerGroup `json:"customer_group,omitempty"`
}

// NewCustomerCreateDataRelationships instantiates a new CustomerCreateDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCreateDataRelationships() *CustomerCreateDataRelationships {
	this := CustomerCreateDataRelationships{}
	return &this
}

// NewCustomerCreateDataRelationshipsWithDefaults instantiates a new CustomerCreateDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCreateDataRelationshipsWithDefaults() *CustomerCreateDataRelationships {
	this := CustomerCreateDataRelationships{}
	return &this
}

// GetCustomerGroup returns the CustomerGroup field value if set, zero value otherwise.
func (o *CustomerCreateDataRelationships) GetCustomerGroup() CustomerCreateDataRelationshipsCustomerGroup {
	if o == nil || o.CustomerGroup == nil {
		var ret CustomerCreateDataRelationshipsCustomerGroup
		return ret
	}
	return *o.CustomerGroup
}

// GetCustomerGroupOk returns a tuple with the CustomerGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCreateDataRelationships) GetCustomerGroupOk() (*CustomerCreateDataRelationshipsCustomerGroup, bool) {
	if o == nil || o.CustomerGroup == nil {
		return nil, false
	}
	return o.CustomerGroup, true
}

// HasCustomerGroup returns a boolean if a field has been set.
func (o *CustomerCreateDataRelationships) HasCustomerGroup() bool {
	if o != nil && o.CustomerGroup != nil {
		return true
	}

	return false
}

// SetCustomerGroup gets a reference to the given CustomerCreateDataRelationshipsCustomerGroup and assigns it to the CustomerGroup field.
func (o *CustomerCreateDataRelationships) SetCustomerGroup(v CustomerCreateDataRelationshipsCustomerGroup) {
	o.CustomerGroup = &v
}

func (o CustomerCreateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomerGroup != nil {
		toSerialize["customer_group"] = o.CustomerGroup
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerCreateDataRelationships struct {
	value *CustomerCreateDataRelationships
	isSet bool
}

func (v NullableCustomerCreateDataRelationships) Get() *CustomerCreateDataRelationships {
	return v.value
}

func (v *NullableCustomerCreateDataRelationships) Set(val *CustomerCreateDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCreateDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCreateDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCreateDataRelationships(val *CustomerCreateDataRelationships) *NullableCustomerCreateDataRelationships {
	return &NullableCustomerCreateDataRelationships{value: val, isSet: true}
}

func (v NullableCustomerCreateDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCreateDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
