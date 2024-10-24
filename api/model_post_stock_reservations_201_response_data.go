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

// checks if the POSTStockReservations201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &POSTStockReservations201ResponseData{}

// POSTStockReservations201ResponseData struct for POSTStockReservations201ResponseData
type POSTStockReservations201ResponseData struct {
	// The resource's id
	Id interface{} `json:"id,omitempty"`
	// The resource's type
	Type          interface{}                                        `json:"type,omitempty"`
	Links         *POSTAddresses201ResponseDataLinks                 `json:"links,omitempty"`
	Attributes    *POSTStockReservations201ResponseDataAttributes    `json:"attributes,omitempty"`
	Relationships *POSTStockReservations201ResponseDataRelationships `json:"relationships,omitempty"`
}

// NewPOSTStockReservations201ResponseData instantiates a new POSTStockReservations201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPOSTStockReservations201ResponseData() *POSTStockReservations201ResponseData {
	this := POSTStockReservations201ResponseData{}
	return &this
}

// NewPOSTStockReservations201ResponseDataWithDefaults instantiates a new POSTStockReservations201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPOSTStockReservations201ResponseDataWithDefaults() *POSTStockReservations201ResponseData {
	this := POSTStockReservations201ResponseData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockReservations201ResponseData) GetId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockReservations201ResponseData) GetIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return &o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseData) HasId() bool {
	if o != nil && IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given interface{} and assigns it to the Id field.
func (o *POSTStockReservations201ResponseData) SetId(v interface{}) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *POSTStockReservations201ResponseData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *POSTStockReservations201ResponseData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseData) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *POSTStockReservations201ResponseData) SetType(v interface{}) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseData) GetLinks() POSTAddresses201ResponseDataLinks {
	if o == nil || IsNil(o.Links) {
		var ret POSTAddresses201ResponseDataLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseData) GetLinksOk() (*POSTAddresses201ResponseDataLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseData) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given POSTAddresses201ResponseDataLinks and assigns it to the Links field.
func (o *POSTStockReservations201ResponseData) SetLinks(v POSTAddresses201ResponseDataLinks) {
	o.Links = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseData) GetAttributes() POSTStockReservations201ResponseDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret POSTStockReservations201ResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseData) GetAttributesOk() (*POSTStockReservations201ResponseDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given POSTStockReservations201ResponseDataAttributes and assigns it to the Attributes field.
func (o *POSTStockReservations201ResponseData) SetAttributes(v POSTStockReservations201ResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *POSTStockReservations201ResponseData) GetRelationships() POSTStockReservations201ResponseDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret POSTStockReservations201ResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *POSTStockReservations201ResponseData) GetRelationshipsOk() (*POSTStockReservations201ResponseDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *POSTStockReservations201ResponseData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given POSTStockReservations201ResponseDataRelationships and assigns it to the Relationships field.
func (o *POSTStockReservations201ResponseData) SetRelationships(v POSTStockReservations201ResponseDataRelationships) {
	o.Relationships = &v
}

func (o POSTStockReservations201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o POSTStockReservations201ResponseData) ToMap() (map[string]interface{}, error) {
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

type NullablePOSTStockReservations201ResponseData struct {
	value *POSTStockReservations201ResponseData
	isSet bool
}

func (v NullablePOSTStockReservations201ResponseData) Get() *POSTStockReservations201ResponseData {
	return v.value
}

func (v *NullablePOSTStockReservations201ResponseData) Set(val *POSTStockReservations201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullablePOSTStockReservations201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullablePOSTStockReservations201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePOSTStockReservations201ResponseData(val *POSTStockReservations201ResponseData) *NullablePOSTStockReservations201ResponseData {
	return &NullablePOSTStockReservations201ResponseData{value: val, isSet: true}
}

func (v NullablePOSTStockReservations201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePOSTStockReservations201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
