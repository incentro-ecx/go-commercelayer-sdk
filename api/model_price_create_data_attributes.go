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

// PriceCreateDataAttributes struct for PriceCreateDataAttributes
type PriceCreateDataAttributes struct {
	// The code of the associated SKU. When creating a price, either a valid sku_code or a SKU relationship must be present.
	SkuCode *string `json:"sku_code,omitempty"`
	// The SKU price amount for the associated price list, in cents.
	AmountCents int32 `json:"amount_cents"`
	// The compared price amount, in cents. Useful to display a percentage discount.
	CompareAtAmountCents int32 `json:"compare_at_amount_cents"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPriceCreateDataAttributes instantiates a new PriceCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceCreateDataAttributes(amountCents int32, compareAtAmountCents int32) *PriceCreateDataAttributes {
	this := PriceCreateDataAttributes{}
	this.AmountCents = amountCents
	this.CompareAtAmountCents = compareAtAmountCents
	return &this
}

// NewPriceCreateDataAttributesWithDefaults instantiates a new PriceCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceCreateDataAttributesWithDefaults() *PriceCreateDataAttributes {
	this := PriceCreateDataAttributes{}
	return &this
}

// GetSkuCode returns the SkuCode field value if set, zero value otherwise.
func (o *PriceCreateDataAttributes) GetSkuCode() string {
	if o == nil || o.SkuCode == nil {
		var ret string
		return ret
	}
	return *o.SkuCode
}

// GetSkuCodeOk returns a tuple with the SkuCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceCreateDataAttributes) GetSkuCodeOk() (*string, bool) {
	if o == nil || o.SkuCode == nil {
		return nil, false
	}
	return o.SkuCode, true
}

// HasSkuCode returns a boolean if a field has been set.
func (o *PriceCreateDataAttributes) HasSkuCode() bool {
	if o != nil && o.SkuCode != nil {
		return true
	}

	return false
}

// SetSkuCode gets a reference to the given string and assigns it to the SkuCode field.
func (o *PriceCreateDataAttributes) SetSkuCode(v string) {
	o.SkuCode = &v
}

// GetAmountCents returns the AmountCents field value
func (o *PriceCreateDataAttributes) GetAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AmountCents
}

// GetAmountCentsOk returns a tuple with the AmountCents field value
// and a boolean to check if the value has been set.
func (o *PriceCreateDataAttributes) GetAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountCents, true
}

// SetAmountCents sets field value
func (o *PriceCreateDataAttributes) SetAmountCents(v int32) {
	o.AmountCents = v
}

// GetCompareAtAmountCents returns the CompareAtAmountCents field value
func (o *PriceCreateDataAttributes) GetCompareAtAmountCents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CompareAtAmountCents
}

// GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field value
// and a boolean to check if the value has been set.
func (o *PriceCreateDataAttributes) GetCompareAtAmountCentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompareAtAmountCents, true
}

// SetCompareAtAmountCents sets field value
func (o *PriceCreateDataAttributes) SetCompareAtAmountCents(v int32) {
	o.CompareAtAmountCents = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PriceCreateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceCreateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PriceCreateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PriceCreateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PriceCreateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceCreateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PriceCreateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PriceCreateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PriceCreateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceCreateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PriceCreateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PriceCreateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PriceCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SkuCode != nil {
		toSerialize["sku_code"] = o.SkuCode
	}
	if true {
		toSerialize["amount_cents"] = o.AmountCents
	}
	if true {
		toSerialize["compare_at_amount_cents"] = o.CompareAtAmountCents
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

type NullablePriceCreateDataAttributes struct {
	value *PriceCreateDataAttributes
	isSet bool
}

func (v NullablePriceCreateDataAttributes) Get() *PriceCreateDataAttributes {
	return v.value
}

func (v *NullablePriceCreateDataAttributes) Set(val *PriceCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceCreateDataAttributes(val *PriceCreateDataAttributes) *NullablePriceCreateDataAttributes {
	return &NullablePriceCreateDataAttributes{value: val, isSet: true}
}

func (v NullablePriceCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
