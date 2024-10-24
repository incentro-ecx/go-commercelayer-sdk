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

// checks if the PATCHCouponsCouponId200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PATCHCouponsCouponId200ResponseData{}

// PATCHCouponsCouponId200ResponseData struct for PATCHCouponsCouponId200ResponseData
type PATCHCouponsCouponId200ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                    `json:"type,omitempty"`
	Links         *POSTAddresses201ResponseDataLinks             `json:"links,omitempty"`
	Attributes    *PATCHCouponsCouponId200ResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *POSTCoupons201ResponseDataRelationships       `json:"relationships,omitempty"`
}

// NewPATCHCouponsCouponId200ResponseData instantiates a new PATCHCouponsCouponId200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHCouponsCouponId200ResponseData() *PATCHCouponsCouponId200ResponseData {
	this := PATCHCouponsCouponId200ResponseData{}
	return &this
}

// NewPATCHCouponsCouponId200ResponseDataWithDefaults instantiates a new PATCHCouponsCouponId200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHCouponsCouponId200ResponseDataWithDefaults() *PATCHCouponsCouponId200ResponseData {
	this := PATCHCouponsCouponId200ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponsCouponId200ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponsCouponId200ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *PATCHCouponsCouponId200ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PATCHCouponsCouponId200ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PATCHCouponsCouponId200ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *PATCHCouponsCouponId200ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseData) GetLinks() POSTAddresses201ResponseDataLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataLinks and assigns it to the Links field.
func (o *PATCHCouponsCouponId200ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseData) GetAttributes() PATCHCouponsCouponId200ResponseDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret PATCHCouponsCouponId200ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseData) GetAttributesOk() (*PATCHCouponsCouponId200ResponseDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given PATCHCouponsCouponId200ResponseDataAttributes and assigns it to the Attributes field.
func (o *PATCHCouponsCouponId200ResponseData) SetAttributes(v PATCHCouponsCouponId200ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PATCHCouponsCouponId200ResponseData) GetRelationships() POSTCoupons201ResponseDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTCoupons201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHCouponsCouponId200ResponseData) GetRelationshipsOk() (*POSTCoupons201ResponseDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PATCHCouponsCouponId200ResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTCoupons201ResponseDataRelationships and assigns it to the Relationships field.
func (o *PATCHCouponsCouponId200ResponseData) SetRelationships(v POSTCoupons201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o PATCHCouponsCouponId200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PATCHCouponsCouponId200ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullablePATCHCouponsCouponId200ResponseData struct {
	value *PATCHCouponsCouponId200ResponseData
	isSet bool
}

func (v NullablePATCHCouponsCouponId200ResponseData) Get() *PATCHCouponsCouponId200ResponseData {
	return v.value
}

func (v *NullablePATCHCouponsCouponId200ResponseData) Set(val *PATCHCouponsCouponId200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHCouponsCouponId200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHCouponsCouponId200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHCouponsCouponId200ResponseData(val *PATCHCouponsCouponId200ResponseData) *NullablePATCHCouponsCouponId200ResponseData {
	return &NullablePATCHCouponsCouponId200ResponseData{value: val, isSet: true}
}

func (v NullablePATCHCouponsCouponId200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHCouponsCouponId200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
