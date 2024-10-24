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

// checks if the PriceFrequencyTierCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceFrequencyTierCreateData{}

// PriceFrequencyTierCreateData struct for PriceFrequencyTierCreateData
type PriceFrequencyTierCreateData struct {
	// The resource's type
	Type          interface{}                                      `json:"type"`
	Attributes    POSTPriceFrequencyTiers201ResponseDataAttributes `json:"attributes"`
	Relationships *PriceFrequencyTierCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewPriceFrequencyTierCreateData instantiates a new PriceFrequencyTierCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceFrequencyTierCreateData(type_ interface{}, attributes POSTPriceFrequencyTiers201ResponseDataAttributes) *PriceFrequencyTierCreateData {
	this := PriceFrequencyTierCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewPriceFrequencyTierCreateDataWithDefaults instantiates a new PriceFrequencyTierCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceFrequencyTierCreateDataWithDefaults() *PriceFrequencyTierCreateData {
	this := PriceFrequencyTierCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *PriceFrequencyTierCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PriceFrequencyTierCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PriceFrequencyTierCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PriceFrequencyTierCreateData) GetAttributes() POSTPriceFrequencyTiers201ResponseDataAttributes {
	if o == nil {
		var ret POSTPriceFrequencyTiers201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PriceFrequencyTierCreateData) GetAttributesOk() (*POSTPriceFrequencyTiers201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PriceFrequencyTierCreateData) SetAttributes(v POSTPriceFrequencyTiers201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PriceFrequencyTierCreateData) GetRelationships() PriceFrequencyTierCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret PriceFrequencyTierCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceFrequencyTierCreateData) GetRelationshipsOk() (*PriceFrequencyTierCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PriceFrequencyTierCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given PriceFrequencyTierCreateDataRelationships and assigns it to the Relationships field.
func (o *PriceFrequencyTierCreateData) SetRelationships(v PriceFrequencyTierCreateDataRelationships) {
	o.Relationships = &v
}

func (o PriceFrequencyTierCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceFrequencyTierCreateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullablePriceFrequencyTierCreateData struct {
	value *PriceFrequencyTierCreateData
	isSet bool
}

func (v NullablePriceFrequencyTierCreateData) Get() *PriceFrequencyTierCreateData {
	return v.value
}

func (v *NullablePriceFrequencyTierCreateData) Set(val *PriceFrequencyTierCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceFrequencyTierCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceFrequencyTierCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceFrequencyTierCreateData(val *PriceFrequencyTierCreateData) *NullablePriceFrequencyTierCreateData {
	return &NullablePriceFrequencyTierCreateData{value: val, isSet: true}
}

func (v NullablePriceFrequencyTierCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceFrequencyTierCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
