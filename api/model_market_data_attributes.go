/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 2.9.4
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MarketDataAttributes struct for MarketDataAttributes
type MarketDataAttributes struct {
	// Unique identifier for the market (numeric)
	Number *int32 `json:"number,omitempty"`
	// The market's internal name
	Name *string `json:"name,omitempty"`
	// The Facebook Pixed ID
	FacebookPixelId *string `json:"facebook_pixel_id,omitempty"`
	// The checkout URL for this market
	CheckoutUrl *string `json:"checkout_url,omitempty"`
	// The URL used to fetch prices from an external source
	ExternalPricesUrl *string `json:"external_prices_url,omitempty"`
	// Indicates if market belongs to a customer_group.
	Private *bool `json:"private,omitempty"`
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
}

// NewMarketDataAttributes instantiates a new MarketDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketDataAttributes() *MarketDataAttributes {
	this := MarketDataAttributes{}
	return &this
}

// NewMarketDataAttributesWithDefaults instantiates a new MarketDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketDataAttributesWithDefaults() *MarketDataAttributes {
	this := MarketDataAttributes{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *MarketDataAttributes) SetNumber(v int32) {
	o.Number = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MarketDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetFacebookPixelId returns the FacebookPixelId field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetFacebookPixelId() string {
	if o == nil || o.FacebookPixelId == nil {
		var ret string
		return ret
	}
	return *o.FacebookPixelId
}

// GetFacebookPixelIdOk returns a tuple with the FacebookPixelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetFacebookPixelIdOk() (*string, bool) {
	if o == nil || o.FacebookPixelId == nil {
		return nil, false
	}
	return o.FacebookPixelId, true
}

// HasFacebookPixelId returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasFacebookPixelId() bool {
	if o != nil && o.FacebookPixelId != nil {
		return true
	}

	return false
}

// SetFacebookPixelId gets a reference to the given string and assigns it to the FacebookPixelId field.
func (o *MarketDataAttributes) SetFacebookPixelId(v string) {
	o.FacebookPixelId = &v
}

// GetCheckoutUrl returns the CheckoutUrl field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetCheckoutUrl() string {
	if o == nil || o.CheckoutUrl == nil {
		var ret string
		return ret
	}
	return *o.CheckoutUrl
}

// GetCheckoutUrlOk returns a tuple with the CheckoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetCheckoutUrlOk() (*string, bool) {
	if o == nil || o.CheckoutUrl == nil {
		return nil, false
	}
	return o.CheckoutUrl, true
}

// HasCheckoutUrl returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasCheckoutUrl() bool {
	if o != nil && o.CheckoutUrl != nil {
		return true
	}

	return false
}

// SetCheckoutUrl gets a reference to the given string and assigns it to the CheckoutUrl field.
func (o *MarketDataAttributes) SetCheckoutUrl(v string) {
	o.CheckoutUrl = &v
}

// GetExternalPricesUrl returns the ExternalPricesUrl field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetExternalPricesUrl() string {
	if o == nil || o.ExternalPricesUrl == nil {
		var ret string
		return ret
	}
	return *o.ExternalPricesUrl
}

// GetExternalPricesUrlOk returns a tuple with the ExternalPricesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetExternalPricesUrlOk() (*string, bool) {
	if o == nil || o.ExternalPricesUrl == nil {
		return nil, false
	}
	return o.ExternalPricesUrl, true
}

// HasExternalPricesUrl returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasExternalPricesUrl() bool {
	if o != nil && o.ExternalPricesUrl != nil {
		return true
	}

	return false
}

// SetExternalPricesUrl gets a reference to the given string and assigns it to the ExternalPricesUrl field.
func (o *MarketDataAttributes) SetExternalPricesUrl(v string) {
	o.ExternalPricesUrl = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetPrivate() bool {
	if o == nil || o.Private == nil {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetPrivateOk() (*bool, bool) {
	if o == nil || o.Private == nil {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasPrivate() bool {
	if o != nil && o.Private != nil {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *MarketDataAttributes) SetPrivate(v bool) {
	o.Private = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MarketDataAttributes) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *MarketDataAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MarketDataAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *MarketDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *MarketDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MarketDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MarketDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *MarketDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o MarketDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.FacebookPixelId != nil {
		toSerialize["facebook_pixel_id"] = o.FacebookPixelId
	}
	if o.CheckoutUrl != nil {
		toSerialize["checkout_url"] = o.CheckoutUrl
	}
	if o.ExternalPricesUrl != nil {
		toSerialize["external_prices_url"] = o.ExternalPricesUrl
	}
	if o.Private != nil {
		toSerialize["private"] = o.Private
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
	return json.Marshal(toSerialize)
}

type NullableMarketDataAttributes struct {
	value *MarketDataAttributes
	isSet bool
}

func (v NullableMarketDataAttributes) Get() *MarketDataAttributes {
	return v.value
}

func (v *NullableMarketDataAttributes) Set(val *MarketDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketDataAttributes(val *MarketDataAttributes) *NullableMarketDataAttributes {
	return &NullableMarketDataAttributes{value: val, isSet: true}
}

func (v NullableMarketDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}