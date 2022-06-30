/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.5
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// LineItemOptionCreateDataAttributes struct for LineItemOptionCreateDataAttributes
type LineItemOptionCreateDataAttributes struct {
	// The name of the line item option. When blank, it gets populated with the name of the associated sku option.
	Name *string `json:"name,omitempty"`
	// The line item option's quantity
	Quantity int32 `json:"quantity"`
	// Set of key-value pairs that represent the selected options.
	Options map[string]interface{} `json:"options"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewLineItemOptionCreateDataAttributes instantiates a new LineItemOptionCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLineItemOptionCreateDataAttributes(quantity int32, options map[string]interface{}) *LineItemOptionCreateDataAttributes {
	this := LineItemOptionCreateDataAttributes{}
	this.Quantity = quantity
	this.Options = options
	return &this
}

// NewLineItemOptionCreateDataAttributesWithDefaults instantiates a new LineItemOptionCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineItemOptionCreateDataAttributesWithDefaults() *LineItemOptionCreateDataAttributes {
	this := LineItemOptionCreateDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LineItemOptionCreateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LineItemOptionCreateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LineItemOptionCreateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetQuantity returns the Quantity field value
func (o *LineItemOptionCreateDataAttributes) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateDataAttributes) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *LineItemOptionCreateDataAttributes) SetQuantity(v int32) {
	o.Quantity = v
}

// GetOptions returns the Options field value
func (o *LineItemOptionCreateDataAttributes) GetOptions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateDataAttributes) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *LineItemOptionCreateDataAttributes) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *LineItemOptionCreateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *LineItemOptionCreateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *LineItemOptionCreateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *LineItemOptionCreateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *LineItemOptionCreateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *LineItemOptionCreateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LineItemOptionCreateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LineItemOptionCreateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LineItemOptionCreateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *LineItemOptionCreateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o LineItemOptionCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["options"] = o.Options
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableLineItemOptionCreateDataAttributes struct {
	value *LineItemOptionCreateDataAttributes
	isSet bool
}

func (v NullableLineItemOptionCreateDataAttributes) Get() *LineItemOptionCreateDataAttributes {
	return v.value
}

func (v *NullableLineItemOptionCreateDataAttributes) Set(val *LineItemOptionCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableLineItemOptionCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableLineItemOptionCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineItemOptionCreateDataAttributes(val *LineItemOptionCreateDataAttributes) *NullableLineItemOptionCreateDataAttributes {
	return &NullableLineItemOptionCreateDataAttributes{value: val, isSet: true}
}

func (v NullableLineItemOptionCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineItemOptionCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
