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

// checks if the MarketCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarketCreateData{}

// MarketCreateData struct for MarketCreateData
type MarketCreateData struct {
	// The resource's type
	Type          interface{}                          `json:"type"`
	Attributes    POSTMarkets201ResponseDataAttributes `json:"attributes"`
	Relationships *MarketCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewMarketCreateData instantiates a new MarketCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketCreateData(type_ interface{}, attributes POSTMarkets201ResponseDataAttributes) *MarketCreateData {
	this := MarketCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewMarketCreateDataWithDefaults instantiates a new MarketCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketCreateDataWithDefaults() *MarketCreateData {
	this := MarketCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *MarketCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MarketCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MarketCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *MarketCreateData) GetAttributes() POSTMarkets201ResponseDataAttributes {
	if o == nil {
		var ret POSTMarkets201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *MarketCreateData) GetAttributesOk() (*POSTMarkets201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *MarketCreateData) SetAttributes(v POSTMarkets201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *MarketCreateData) GetRelationships() MarketCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret MarketCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketCreateData) GetRelationshipsOk() (*MarketCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *MarketCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given MarketCreateDataRelationships and assigns it to the Relationships field.
func (o *MarketCreateData) SetRelationships(v MarketCreateDataRelationships) {
	o.Relationships = &v
}

func (o MarketCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarketCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableMarketCreateData struct {
	value *MarketCreateData
	isSet bool
}

func (v NullableMarketCreateData) Get() *MarketCreateData {
	return v.value
}

func (v *NullableMarketCreateData) Set(val *MarketCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketCreateData(val *MarketCreateData) *NullableMarketCreateData {
	return &NullableMarketCreateData{value: val, isSet: true}
}

func (v NullableMarketCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
