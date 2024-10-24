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

// checks if the AdyenGatewayCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdyenGatewayCreateData{}

// AdyenGatewayCreateData struct for AdyenGatewayCreateData
type AdyenGatewayCreateData struct {
	// The resource's type
	Type          interface{}                                `json:"type"`
	Attributes    POSTAdyenGateways201ResponseDataAttributes `json:"attributes"`
	Relationships *AdyenGatewayCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewAdyenGatewayCreateData instantiates a new AdyenGatewayCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdyenGatewayCreateData(type_ interface{}, attributes POSTAdyenGateways201ResponseDataAttributes) *AdyenGatewayCreateData {
	this := AdyenGatewayCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAdyenGatewayCreateDataWithDefaults instantiates a new AdyenGatewayCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdyenGatewayCreateDataWithDefaults() *AdyenGatewayCreateData {
	this := AdyenGatewayCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AdyenGatewayCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdyenGatewayCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdyenGatewayCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AdyenGatewayCreateData) GetAttributes() POSTAdyenGateways201ResponseDataAttributes {
	if o == nil {
		var ret POSTAdyenGateways201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AdyenGatewayCreateData) GetAttributesOk() (*POSTAdyenGateways201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AdyenGatewayCreateData) SetAttributes(v POSTAdyenGateways201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AdyenGatewayCreateData) GetRelationships() AdyenGatewayCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AdyenGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdyenGatewayCreateData) GetRelationshipsOk() (*AdyenGatewayCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AdyenGatewayCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AdyenGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *AdyenGatewayCreateData) SetRelationships(v AdyenGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o AdyenGatewayCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdyenGatewayCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableAdyenGatewayCreateData struct {
	value *AdyenGatewayCreateData
	isSet bool
}

func (v NullableAdyenGatewayCreateData) Get() *AdyenGatewayCreateData {
	return v.value
}

func (v *NullableAdyenGatewayCreateData) Set(val *AdyenGatewayCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAdyenGatewayCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAdyenGatewayCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdyenGatewayCreateData(val *AdyenGatewayCreateData) *NullableAdyenGatewayCreateData {
	return &NullableAdyenGatewayCreateData{value: val, isSet: true}
}

func (v NullableAdyenGatewayCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdyenGatewayCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
