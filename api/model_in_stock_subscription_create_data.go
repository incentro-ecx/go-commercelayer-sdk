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

// checks if the InStockSubscriptionCreateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InStockSubscriptionCreateData{}

// InStockSubscriptionCreateData struct for InStockSubscriptionCreateData
type InStockSubscriptionCreateData struct {
	// The resource's type
	Type          interface{}                                       `json:"type"`
	Attributes    POSTInStockSubscriptions201ResponseDataAttributes `json:"attributes"`
	Relationships *InStockSubscriptionCreateDataRelationships       `json:"relationships,omitempty"`
}

// NewInStockSubscriptionCreateData instantiates a new InStockSubscriptionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInStockSubscriptionCreateData(type_ interface{}, attributes POSTInStockSubscriptions201ResponseDataAttributes) *InStockSubscriptionCreateData {
	this := InStockSubscriptionCreateData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewInStockSubscriptionCreateDataWithDefaults instantiates a new InStockSubscriptionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInStockSubscriptionCreateDataWithDefaults() *InStockSubscriptionCreateData {
	this := InStockSubscriptionCreateData{}
	return &this
}

// GetType returns the Type field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *InStockSubscriptionCreateData) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InStockSubscriptionCreateData) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InStockSubscriptionCreateData) SetType(v interface{}) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *InStockSubscriptionCreateData) GetAttributes() POSTInStockSubscriptions201ResponseDataAttributes {
	if o == nil {
		var ret POSTInStockSubscriptions201ResponseDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateData) GetAttributesOk() (*POSTInStockSubscriptions201ResponseDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *InStockSubscriptionCreateData) SetAttributes(v POSTInStockSubscriptions201ResponseDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InStockSubscriptionCreateData) GetRelationships() InStockSubscriptionCreateDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret InStockSubscriptionCreateDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InStockSubscriptionCreateData) GetRelationshipsOk() (*InStockSubscriptionCreateDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InStockSubscriptionCreateData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given InStockSubscriptionCreateDataRelationships and assigns it to the Relationships field.
func (o *InStockSubscriptionCreateData) SetRelationships(v InStockSubscriptionCreateDataRelationships) {
	o.Relationships = &v
}

func (o InStockSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InStockSubscriptionCreateData) ToMap() (map[string]interface{}, error) {
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

type NullableInStockSubscriptionCreateData struct {
	value *InStockSubscriptionCreateData
	isSet bool
}

func (v NullableInStockSubscriptionCreateData) Get() *InStockSubscriptionCreateData {
	return v.value
}

func (v *NullableInStockSubscriptionCreateData) Set(val *InStockSubscriptionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableInStockSubscriptionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableInStockSubscriptionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInStockSubscriptionCreateData(val *InStockSubscriptionCreateData) *NullableInStockSubscriptionCreateData {
	return &NullableInStockSubscriptionCreateData{value: val, isSet: true}
}

func (v NullableInStockSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInStockSubscriptionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
