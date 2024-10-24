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

// checks if the POSTCoupons201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTCoupons201ResponseData{}

// POSTCoupons201ResponseData struct for POSTCoupons201ResponseData
type POSTCoupons201ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                              `json:"type,omitempty"`
	Links         *POSTAddresses201ResponseDataLinks       `json:"links,omitempty"`
	Attributes    *POSTCoupons201ResponseDataAttributes    `json:"attributes,omitempty"`
	Relationships *POSTCoupons201ResponseDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTCoupons201ResponseData instantiates a new POSTCoupons201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTCoupons201ResponseData() *POSTCoupons201ResponseData {
	this := POSTCoupons201ResponseData{}
	return &this
}

// NewPOSTCoupons201ResponseDataWithDefaults instantiates a new POSTCoupons201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTCoupons201ResponseDataWithDefaults() *POSTCoupons201ResponseData {
	this := POSTCoupons201ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCoupons201ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCoupons201ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTCoupons201ResponseData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTCoupons201ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTCoupons201ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTCoupons201ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTCoupons201ResponseData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTCoupons201ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTCoupons201ResponseData) GetLinks() POSTAddresses201ResponseDataLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCoupons201ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTCoupons201ResponseData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataLinks and assigns it to the Links field.
func (o *POSTCoupons201ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *POSTCoupons201ResponseData) GetAttributes() POSTCoupons201ResponseDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret POSTCoupons201ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCoupons201ResponseData) GetAttributesOk() (*POSTCoupons201ResponseDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *POSTCoupons201ResponseData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given POSTCoupons201ResponseDataAttributes and assigns it to the Attributes field.
func (o *POSTCoupons201ResponseData) SetAttributes(v POSTCoupons201ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTCoupons201ResponseData) GetRelationships() POSTCoupons201ResponseDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTCoupons201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTCoupons201ResponseData) GetRelationshipsOk() (*POSTCoupons201ResponseDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTCoupons201ResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTCoupons201ResponseDataRelationships and assigns it to the Relationships field.
func (o *POSTCoupons201ResponseData) SetRelationships(v POSTCoupons201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o POSTCoupons201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTCoupons201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePOSTCoupons201ResponseData struct {
	value *POSTCoupons201ResponseData
	isSet bool
}

func (v NullablePOSTCoupons201ResponseData) Get() *POSTCoupons201ResponseData {
	return v.value
}

func (v *NullablePOSTCoupons201ResponseData) Set(val *POSTCoupons201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTCoupons201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTCoupons201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTCoupons201ResponseData(val *POSTCoupons201ResponseData) *NullablePOSTCoupons201ResponseData {
	return &NullablePOSTCoupons201ResponseData{value: val, isSet: true}
}

func (v NullablePOSTCoupons201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTCoupons201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
