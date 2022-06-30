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

// FreeShippingPromotionCreateDataAttributes struct for FreeShippingPromotionCreateDataAttributes
type FreeShippingPromotionCreateDataAttributes struct {
	// The promotion's internal name.
	Name string `json:"name"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The activation date/time of this promotion.
	StartsAt string `json:"starts_at"`
	// The expiration date/time of this promotion (must be after starts_at).
	ExpiresAt string `json:"expires_at"`
	// The total number of times this promotion can be applied.
	TotalUsageLimit int32 `json:"total_usage_limit"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewFreeShippingPromotionCreateDataAttributes instantiates a new FreeShippingPromotionCreateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeShippingPromotionCreateDataAttributes(name string, startsAt string, expiresAt string, totalUsageLimit int32) *FreeShippingPromotionCreateDataAttributes {
	this := FreeShippingPromotionCreateDataAttributes{}
	this.Name = name
	this.StartsAt = startsAt
	this.ExpiresAt = expiresAt
	this.TotalUsageLimit = totalUsageLimit
	return &this
}

// NewFreeShippingPromotionCreateDataAttributesWithDefaults instantiates a new FreeShippingPromotionCreateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeShippingPromotionCreateDataAttributesWithDefaults() *FreeShippingPromotionCreateDataAttributes {
	this := FreeShippingPromotionCreateDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *FreeShippingPromotionCreateDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FreeShippingPromotionCreateDataAttributes) SetName(v string) {
	o.Name = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *FreeShippingPromotionCreateDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *FreeShippingPromotionCreateDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *FreeShippingPromotionCreateDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetStartsAt returns the StartsAt field value
func (o *FreeShippingPromotionCreateDataAttributes) GetStartsAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateDataAttributes) GetStartsAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartsAt, true
}

// SetStartsAt sets field value
func (o *FreeShippingPromotionCreateDataAttributes) SetStartsAt(v string) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *FreeShippingPromotionCreateDataAttributes) GetExpiresAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateDataAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *FreeShippingPromotionCreateDataAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value
func (o *FreeShippingPromotionCreateDataAttributes) GetTotalUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateDataAttributes) GetTotalUsageLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalUsageLimit, true
}

// SetTotalUsageLimit sets field value
func (o *FreeShippingPromotionCreateDataAttributes) SetTotalUsageLimit(v int32) {
	o.TotalUsageLimit = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *FreeShippingPromotionCreateDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *FreeShippingPromotionCreateDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *FreeShippingPromotionCreateDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *FreeShippingPromotionCreateDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *FreeShippingPromotionCreateDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *FreeShippingPromotionCreateDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FreeShippingPromotionCreateDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FreeShippingPromotionCreateDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FreeShippingPromotionCreateDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *FreeShippingPromotionCreateDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o FreeShippingPromotionCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if true {
		toSerialize["starts_at"] = o.StartsAt
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if true {
		toSerialize["total_usage_limit"] = o.TotalUsageLimit
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

type NullableFreeShippingPromotionCreateDataAttributes struct {
	value *FreeShippingPromotionCreateDataAttributes
	isSet bool
}

func (v NullableFreeShippingPromotionCreateDataAttributes) Get() *FreeShippingPromotionCreateDataAttributes {
	return v.value
}

func (v *NullableFreeShippingPromotionCreateDataAttributes) Set(val *FreeShippingPromotionCreateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeShippingPromotionCreateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeShippingPromotionCreateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeShippingPromotionCreateDataAttributes(val *FreeShippingPromotionCreateDataAttributes) *NullableFreeShippingPromotionCreateDataAttributes {
	return &NullableFreeShippingPromotionCreateDataAttributes{value: val, isSet: true}
}

func (v NullableFreeShippingPromotionCreateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeShippingPromotionCreateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
