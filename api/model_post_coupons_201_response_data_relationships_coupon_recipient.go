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

// checks if the POSTCoupons201ResponseDataRelationshipsCouponRecipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCoupons201ResponseDataRelationshipsCouponRecipient{}

// POSTCoupons201ResponseDataRelationshipsCouponRecipient struct for POSTCoupons201ResponseDataRelationshipsCouponRecipient
type POSTCoupons201ResponseDataRelationshipsCouponRecipient struct {
	Links *POSTAddresses201ResponseDataRelationshipsGeocoderLinks     `json:"links,omitempty"`
	Data  *POSTCoupons201ResponseDataRelationshipsCouponRecipientData `json:"data,omitempty"`
}

// NewPOSTCoupons201ResponseDataRelationshipsCouponRecipient instantiates a new POSTCoupons201ResponseDataRelationshipsCouponRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCoupons201ResponseDataRelationshipsCouponRecipient() *POSTCoupons201ResponseDataRelationshipsCouponRecipient {
	this := POSTCoupons201ResponseDataRelationshipsCouponRecipient{}
	return &this
}

// NewPOSTCoupons201ResponseDataRelationshipsCouponRecipientWithDefaults instantiates a new POSTCoupons201ResponseDataRelationshipsCouponRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCoupons201ResponseDataRelationshipsCouponRecipientWithDefaults() *POSTCoupons201ResponseDataRelationshipsCouponRecipient {
	this := POSTCoupons201ResponseDataRelationshipsCouponRecipient{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCoupons201ResponseDataRelationshipsCouponRecipient) GetLinks() POSTAddresses201ResponseDataRelationshipsGeocoderLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataRelationshipsGeocoderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCoupons201ResponseDataRelationshipsCouponRecipient) GetLinksOk() (*POSTAddresses201ResponseDataRelationshipsGeocoderLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCoupons201ResponseDataRelationshipsCouponRecipient) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataRelationshipsGeocoderLinks and assigns it to the Links field.
func (o *POSTCoupons201ResponseDataRelationshipsCouponRecipient) SetLinks(v POSTAddresses201ResponseDataRelationshipsGeocoderLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *POSTCoupons201ResponseDataRelationshipsCouponRecipient) GetData() POSTCoupons201ResponseDataRelationshipsCouponRecipientData {
	if o == nil || IsNil(o.Data) {
		var ret POSTCoupons201ResponseDataRelationshipsCouponRecipientData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCoupons201ResponseDataRelationshipsCouponRecipient) GetDataOk() (*POSTCoupons201ResponseDataRelationshipsCouponRecipientData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *POSTCoupons201ResponseDataRelationshipsCouponRecipient) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given POSTCoupons201ResponseDataRelationshipsCouponRecipientData and assigns it to the Data field.
func (o *POSTCoupons201ResponseDataRelationshipsCouponRecipient) SetData(v POSTCoupons201ResponseDataRelationshipsCouponRecipientData) {
	o.Data = &v
}

func (o POSTCoupons201ResponseDataRelationshipsCouponRecipient) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCoupons201ResponseDataRelationshipsCouponRecipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient struct {
	value *POSTCoupons201ResponseDataRelationshipsCouponRecipient
	isSet bool
}

func (v NullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient) Get() *POSTCoupons201ResponseDataRelationshipsCouponRecipient {
	return v.value
}

func (v *NullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient) Set(val *POSTCoupons201ResponseDataRelationshipsCouponRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient(val *POSTCoupons201ResponseDataRelationshipsCouponRecipient) *NullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient {
	return &NullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient{value: val, isSet: true}
}

func (v NullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCoupons201ResponseDataRelationshipsCouponRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}