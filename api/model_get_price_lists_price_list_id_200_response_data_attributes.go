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

// checks if the GETPriceListsPriceListId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETPriceListsPriceListId200ResponseDataAttributes{}

// GETPriceListsPriceListId200ResponseDataAttributes struct for GETPriceListsPriceListId200ResponseDataAttributes
type GETPriceListsPriceListId200ResponseDataAttributes struct {
	// The price list's internal name.
	Name interface{} `json:"name,omitempty"`
	// A string that you can use to identify the price list (must be unique within the environment).
	Code interface{} `json:"code,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode interface{} `json:"currency_code,omitempty"`
	// Indicates if the associated prices include taxes.
	TaxIncluded interface{} `json:"tax_included,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// The rules (using Rules Engine) to be applied.
	Rules interface{} `json:"rules,omitempty"`
	// The rule outcomes.
	RuleOutcomes interface{} `json:"rule_outcomes,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETPriceListsPriceListId200ResponseDataAttributes instantiates a new GETPriceListsPriceListId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETPriceListsPriceListId200ResponseDataAttributes() *GETPriceListsPriceListId200ResponseDataAttributes {
	this := GETPriceListsPriceListId200ResponseDataAttributes{}
	return &this
}

// NewGETPriceListsPriceListId200ResponseDataAttributesWithDefaults instantiates a new GETPriceListsPriceListId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETPriceListsPriceListId200ResponseDataAttributesWithDefaults() *GETPriceListsPriceListId200ResponseDataAttributes {
	this := GETPriceListsPriceListId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasCode() bool {
	if o != nil && IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetCode(v interface{}) {
	o.Code = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCurrencyCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasCurrencyCode() bool {
	if o != nil && IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given interface{} and assigns it to the CurrencyCode field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetCurrencyCode(v interface{}) {
	o.CurrencyCode = v
}

// GetTaxIncluded returns the TaxIncluded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetTaxIncluded() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.TaxIncluded
}

// GetTaxIncludedOk returns a tuple with the TaxIncluded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetTaxIncludedOk() (*interface{}, bool) {
	if o == nil || IsNil(o.TaxIncluded) {
		return nil, false
	}
	return &o.TaxIncluded, true
}

// HasTaxIncluded returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasTaxIncluded() bool {
	if o != nil && IsNil(o.TaxIncluded) {
		return true
	}

	return false
}

// SetTaxIncluded gets a reference to the given interface{} and assigns it to the TaxIncluded field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetTaxIncluded(v interface{}) {
	o.TaxIncluded = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetRules returns the Rules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetRules() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetRulesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return &o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasRules() bool {
	if o != nil && IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given interface{} and assigns it to the Rules field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetRules(v interface{}) {
	o.Rules = v
}

// GetRuleOutcomes returns the RuleOutcomes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetRuleOutcomes() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RuleOutcomes
}

// GetRuleOutcomesOk returns a tuple with the RuleOutcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetRuleOutcomesOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RuleOutcomes) {
		return nil, false
	}
	return &o.RuleOutcomes, true
}

// HasRuleOutcomes returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasRuleOutcomes() bool {
	if o != nil && IsNil(o.RuleOutcomes) {
		return true
	}

	return false
}

// SetRuleOutcomes gets a reference to the given interface{} and assigns it to the RuleOutcomes field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetRuleOutcomes(v interface{}) {
	o.RuleOutcomes = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETPriceListsPriceListId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETPriceListsPriceListId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETPriceListsPriceListId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETPriceListsPriceListId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.TaxIncluded != nil {
		toSerialize["tax_included"] = o.TaxIncluded
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.ReferenceOrigin != nil {
		toSerialize["reference_origin"] = o.ReferenceOrigin
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	if o.RuleOutcomes != nil {
		toSerialize["rule_outcomes"] = o.RuleOutcomes
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableGETPriceListsPriceListId200ResponseDataAttributes struct {
	value *GETPriceListsPriceListId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETPriceListsPriceListId200ResponseDataAttributes) Get() *GETPriceListsPriceListId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETPriceListsPriceListId200ResponseDataAttributes) Set(val *GETPriceListsPriceListId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETPriceListsPriceListId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETPriceListsPriceListId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETPriceListsPriceListId200ResponseDataAttributes(val *GETPriceListsPriceListId200ResponseDataAttributes) *NullableGETPriceListsPriceListId200ResponseDataAttributes {
	return &NullableGETPriceListsPriceListId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETPriceListsPriceListId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETPriceListsPriceListId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
