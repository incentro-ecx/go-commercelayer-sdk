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

// checks if the POSTOrders201ResponseDataRelationshipsPaymentMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTOrders201ResponseDataRelationshipsPaymentMethod{}

// POSTOrders201ResponseDataRelationshipsPaymentMethod struct for POSTOrders201ResponseDataRelationshipsPaymentMethod
type POSTOrders201ResponseDataRelationshipsPaymentMethod struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks  `json:"links,omitempty"`
	Data  *POSTOrders201ResponseDataRelationshipsPaymentMethodData `json:"data,omitempty"`
}

// NewPOSTOrders201ResponseDataRelationshipsPaymentMethod instantiates a new POSTOrders201ResponseDataRelationshipsPaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTOrders201ResponseDataRelationshipsPaymentMethod() *POSTOrders201ResponseDataRelationshipsPaymentMethod {
	this := POSTOrders201ResponseDataRelationshipsPaymentMethod{}
	return &this
}

// NewPOSTOrders201ResponseDataRelationshipsPaymentMethodWithDefaults instantiates a new POSTOrders201ResponseDataRelationshipsPaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTOrders201ResponseDataRelationshipsPaymentMethodWithDefaults() *POSTOrders201ResponseDataRelationshipsPaymentMethod {
	this := POSTOrders201ResponseDataRelationshipsPaymentMethod{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsPaymentMethod) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentMethod) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentMethod) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTOrders201ResponseDataRelationshipsPaymentMethod) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTOrders201ResponseDataRelationshipsPaymentMethod) GetData() POSTOrders201ResponseDataRelationshipsPaymentMethodData {
	if o == nil || IsNil(o.Data) {
		var ret POSTOrders201ResponseDataRelationshipsPaymentMethodData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentMethod) GetDataOk() (*POSTOrders201ResponseDataRelationshipsPaymentMethodData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTOrders201ResponseDataRelationshipsPaymentMethod) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTOrders201ResponseDataRelationshipsPaymentMethodData and assigns it to the Data field.
func (o *POSTOrders201ResponseDataRelationshipsPaymentMethod) SetData(v POSTOrders201ResponseDataRelationshipsPaymentMethodData) {
	o.Data = &v
}

func (o POSTOrders201ResponseDataRelationshipsPaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTOrders201ResponseDataRelationshipsPaymentMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTOrders201ResponseDataRelationshipsPaymentMethod struct {
	value *POSTOrders201ResponseDataRelationshipsPaymentMethod
	isSet bool
}

func (v NullablePOSTOrders201ResponseDataRelationshipsPaymentMethod) Get() *POSTOrders201ResponseDataRelationshipsPaymentMethod {
	return v.value
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsPaymentMethod) Set(val *POSTOrders201ResponseDataRelationshipsPaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTOrders201ResponseDataRelationshipsPaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsPaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTOrders201ResponseDataRelationshipsPaymentMethod(val *POSTOrders201ResponseDataRelationshipsPaymentMethod) *NullablePOSTOrders201ResponseDataRelationshipsPaymentMethod {
	return &NullablePOSTOrders201ResponseDataRelationshipsPaymentMethod{value: val, isSet: true}
}

func (v NullablePOSTOrders201ResponseDataRelationshipsPaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTOrders201ResponseDataRelationshipsPaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
