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

// checks if the AxerveGatewayCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AxerveGatewayCreateData{}

// AxerveGatewayCreateData struct for AxerveGatewayCreateData
type AxerveGatewayCreateData struct {
	// The resource's type
	Type          interface{}                                 `json:"type"`
	Attributes    POSTAxerveGateways201ResponseDataAttributes `json:"attributes"`
	Relationships *AxerveGatewayCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewAxerveGatewayCreateData instantiates a new AxerveGatewayCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAxerveGatewayCreateData(type_ interface{}, attributes POSTAxerveGateways201ResponseDataAttributes) *AxerveGatewayCreateData {
	this := AxerveGatewayCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAxerveGatewayCreateDataWithDefaults instantiates a new AxerveGatewayCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAxerveGatewayCreateDataWithDefaults() *AxerveGatewayCreateData {
	this := AxerveGatewayCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *AxerveGatewayCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AxerveGatewayCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AxerveGatewayCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AxerveGatewayCreateData) GetAttributes() POSTAxerveGateways201ResponseDataAttributes {
	if o == nil {
		var ret POSTAxerveGateways201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AxerveGatewayCreateData) GetAttributesOk() (*POSTAxerveGateways201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AxerveGatewayCreateData) SetAttributes(v POSTAxerveGateways201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AxerveGatewayCreateData) GetRelationships() AxerveGatewayCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AxerveGatewayCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AxerveGatewayCreateData) GetRelationshipsOk() (*AxerveGatewayCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AxerveGatewayCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AxerveGatewayCreateDataRelationships and assigns it to the Relationships field.
func (o *AxerveGatewayCreateData) SetRelationships(v AxerveGatewayCreateDataRelationships) {
	o.Relationships = &v
}

func (o AxerveGatewayCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AxerveGatewayCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableAxerveGatewayCreateData struct {
	value *AxerveGatewayCreateData
	isSet bool
}

func (v NullableAxerveGatewayCreateData) Get() *AxerveGatewayCreateData {
	return v.value
}

func (v *NullableAxerveGatewayCreateData) Set(val *AxerveGatewayCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAxerveGatewayCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAxerveGatewayCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAxerveGatewayCreateData(val *AxerveGatewayCreateData) *NullableAxerveGatewayCreateData {
	return &NullableAxerveGatewayCreateData{value: val, isSet: true}
}

func (v NullableAxerveGatewayCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAxerveGatewayCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
