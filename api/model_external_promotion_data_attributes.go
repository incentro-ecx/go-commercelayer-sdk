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

// ExternalPromotionDataAttributes struct for ExternalPromotionDataAttributes
type ExternalPromotionDataAttributes struct {
	// The promotion's internal name.
	Name *string `json:"name,omitempty"`
	// The international 3-letter currency code as defined by the ISO 4217 standard.
	CurrencyCode *string `json:"currency_code,omitempty"`
	// The activation date/time of this promotion.
	StartsAt *string `json:"starts_at,omitempty"`
	// The expiration date/time of this promotion (must be after starts_at).
	ExpiresAt *string `json:"expires_at,omitempty"`
	// The total number of times this promotion can be applied.
	TotalUsageLimit *int32 `json:"total_usage_limit,omitempty"`
	// The number of times this promotion has been applied.
	TotalUsageCount *int32 `json:"total_usage_count,omitempty"`
	// Indicates if the promotion is active.
	Active *bool `json:"active,omitempty"`
	// Unique identifier for the resource (hash).
	Id *string `json:"id,omitempty"`
	// Time at which the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// The URL to the service that will compute the discount.
	PromotionUrl *string `json:"promotion_url,omitempty"`
}

// NewExternalPromotionDataAttributes instantiates a new ExternalPromotionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPromotionDataAttributes() *ExternalPromotionDataAttributes {
	this := ExternalPromotionDataAttributes{}
	return &this
}

// NewExternalPromotionDataAttributesWithDefaults instantiates a new ExternalPromotionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPromotionDataAttributesWithDefaults() *ExternalPromotionDataAttributes {
	this := ExternalPromotionDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalPromotionDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetCurrencyCode() string {
	if o == nil || o.CurrencyCode == nil {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || o.CurrencyCode == nil {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasCurrencyCode() bool {
	if o != nil && o.CurrencyCode != nil {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *ExternalPromotionDataAttributes) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetStartsAt() string {
	if o == nil || o.StartsAt == nil {
		var ret string
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetStartsAtOk() (*string, bool) {
	if o == nil || o.StartsAt == nil {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasStartsAt() bool {
	if o != nil && o.StartsAt != nil {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given string and assigns it to the StartsAt field.
func (o *ExternalPromotionDataAttributes) SetStartsAt(v string) {
	o.StartsAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *ExternalPromotionDataAttributes) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetTotalUsageLimit returns the TotalUsageLimit field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetTotalUsageLimit() int32 {
	if o == nil || o.TotalUsageLimit == nil {
		var ret int32
		return ret
	}
	return *o.TotalUsageLimit
}

// GetTotalUsageLimitOk returns a tuple with the TotalUsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetTotalUsageLimitOk() (*int32, bool) {
	if o == nil || o.TotalUsageLimit == nil {
		return nil, false
	}
	return o.TotalUsageLimit, true
}

// HasTotalUsageLimit returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasTotalUsageLimit() bool {
	if o != nil && o.TotalUsageLimit != nil {
		return true
	}

	return false
}

// SetTotalUsageLimit gets a reference to the given int32 and assigns it to the TotalUsageLimit field.
func (o *ExternalPromotionDataAttributes) SetTotalUsageLimit(v int32) {
	o.TotalUsageLimit = &v
}

// GetTotalUsageCount returns the TotalUsageCount field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetTotalUsageCount() int32 {
	if o == nil || o.TotalUsageCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalUsageCount
}

// GetTotalUsageCountOk returns a tuple with the TotalUsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetTotalUsageCountOk() (*int32, bool) {
	if o == nil || o.TotalUsageCount == nil {
		return nil, false
	}
	return o.TotalUsageCount, true
}

// HasTotalUsageCount returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasTotalUsageCount() bool {
	if o != nil && o.TotalUsageCount != nil {
		return true
	}

	return false
}

// SetTotalUsageCount gets a reference to the given int32 and assigns it to the TotalUsageCount field.
func (o *ExternalPromotionDataAttributes) SetTotalUsageCount(v int32) {
	o.TotalUsageCount = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ExternalPromotionDataAttributes) SetActive(v bool) {
	o.Active = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExternalPromotionDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ExternalPromotionDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ExternalPromotionDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ExternalPromotionDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *ExternalPromotionDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ExternalPromotionDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetPromotionUrl returns the PromotionUrl field value if set, zero value otherwise.
func (o *ExternalPromotionDataAttributes) GetPromotionUrl() string {
	if o == nil || o.PromotionUrl == nil {
		var ret string
		return ret
	}
	return *o.PromotionUrl
}

// GetPromotionUrlOk returns a tuple with the PromotionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPromotionDataAttributes) GetPromotionUrlOk() (*string, bool) {
	if o == nil || o.PromotionUrl == nil {
		return nil, false
	}
	return o.PromotionUrl, true
}

// HasPromotionUrl returns a boolean if a field has been set.
func (o *ExternalPromotionDataAttributes) HasPromotionUrl() bool {
	if o != nil && o.PromotionUrl != nil {
		return true
	}

	return false
}

// SetPromotionUrl gets a reference to the given string and assigns it to the PromotionUrl field.
func (o *ExternalPromotionDataAttributes) SetPromotionUrl(v string) {
	o.PromotionUrl = &v
}

func (o ExternalPromotionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CurrencyCode != nil {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	if o.StartsAt != nil {
		toSerialize["starts_at"] = o.StartsAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.TotalUsageLimit != nil {
		toSerialize["total_usage_limit"] = o.TotalUsageLimit
	}
	if o.TotalUsageCount != nil {
		toSerialize["total_usage_count"] = o.TotalUsageCount
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PromotionUrl != nil {
		toSerialize["promotion_url"] = o.PromotionUrl
	}
	return json.Marshal(toSerialize)
}

type NullableExternalPromotionDataAttributes struct {
	value *ExternalPromotionDataAttributes
	isSet bool
}

func (v NullableExternalPromotionDataAttributes) Get() *ExternalPromotionDataAttributes {
	return v.value
}

func (v *NullableExternalPromotionDataAttributes) Set(val *ExternalPromotionDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPromotionDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPromotionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPromotionDataAttributes(val *ExternalPromotionDataAttributes) *NullableExternalPromotionDataAttributes {
	return &NullableExternalPromotionDataAttributes{value: val, isSet: true}
}

func (v NullableExternalPromotionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPromotionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
