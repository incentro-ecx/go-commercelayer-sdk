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

// AdjustmentUpdateDataAttributes struct for AdjustmentUpdateDataAttributes
type AdjustmentUpdateDataAttributes struct {
	// The adjustment name
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The adjustment amount, in cents.
	AmountCents *int32 `json:"amount_cents,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewAdjustmentUpdateDataAttributes instantiates a new AdjustmentUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdjustmentUpdateDataAttributes() *AdjustmentUpdateDataAttributes {
	this := AdjustmentUpdateDataAttributes{}
	return &this
}

// NewAdjustmentUpdateDataAttributesWithDefaults instantiates a new AdjustmentUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdjustmentUpdateDataAttributesWithDefaults() *AdjustmentUpdateDataAttributes {
	this := AdjustmentUpdateDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdjustmentUpdateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentUpdateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdjustmentUpdateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdjustmentUpdateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *AdjustmentUpdateDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentUpdateDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *AdjustmentUpdateDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *AdjustmentUpdateDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetAmountCents returns the AmountCents field value if set, zero value otherwise.
func (o *AdjustmentUpdateDataAttributes) GetAmountCents() int32 {
	if o == nil || o.AmountCents == nil {
		var ret int32
		return ret
	}
	return *o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentUpdateDataAttributes) GetAmountCentsOk() (*int32, bool) {
	if o == nil || o.AmountCents == nil {
		return nil, false
	}
	return o.AmountCents, true
}

// HasAmountCents returns a boolean if a field has been set.
func (o *AdjustmentUpdateDataAttributes) HasAmountCents() bool {
	if o != nil && o.AmountCents != nil {
		return true
	}

	return false
}

// SetAmountCents gets a reference to the given int32 and assigns it to the AmountCents field.
func (o *AdjustmentUpdateDataAttributes) SetAmountCents(v int32) {
	o.AmountCents = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *AdjustmentUpdateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentUpdateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *AdjustmentUpdateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *AdjustmentUpdateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *AdjustmentUpdateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentUpdateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *AdjustmentUpdateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *AdjustmentUpdateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AdjustmentUpdateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdjustmentUpdateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AdjustmentUpdateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *AdjustmentUpdateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o AdjustmentUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.AmountCents != nil {
		toSerialize["amount_cents"] = o.AmountCents
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

type NullableAdjustmentUpdateDataAttributes struct {
	value *AdjustmentUpdateDataAttributes
	isSet bool
}

func (v NullableAdjustmentUpdateDataAttributes) Get() *AdjustmentUpdateDataAttributes {
	return v.value
}

func (v *NullableAdjustmentUpdateDataAttributes) Set(val *AdjustmentUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAdjustmentUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAdjustmentUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdjustmentUpdateDataAttributes(val *AdjustmentUpdateDataAttributes) *NullableAdjustmentUpdateDataAttributes {
	return &NullableAdjustmentUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableAdjustmentUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdjustmentUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
