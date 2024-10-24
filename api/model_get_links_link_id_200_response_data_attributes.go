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

// checks if the GETLinksLinkId200ResponseDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GETLinksLinkId200ResponseDataAttributes{}

// GETLinksLinkId200ResponseDataAttributes struct for GETLinksLinkId200ResponseDataAttributes
type GETLinksLinkId200ResponseDataAttributes struct {
	// The link internal name.
	Name interface{} `json:"name,omitempty"`
	// The link application client id, used to fetch JWT.
	ClientId interface{} `json:"client_id,omitempty"`
	// The link application scope, used to fetch JWT.
	Scope interface{} `json:"scope,omitempty"`
	// The activation date/time of this link.
	StartsAt interface{} `json:"starts_at,omitempty"`
	// The expiration date/time of this link (must be after starts_at).
	ExpiresAt interface{} `json:"expires_at,omitempty"`
	// Indicates if the link is active (enabled and not expired).
	Active interface{} `json:"active,omitempty"`
	// The link status. One of 'disabled', 'expired', 'pending', or 'active'.
	Status interface{} `json:"status,omitempty"`
	// The link URL second level domain.
	Domain interface{} `json:"domain,omitempty"`
	// The link URL.
	Url interface{} `json:"url,omitempty"`
	// The type of the associated item. One of 'orders', 'skus', or 'sku_lists'.
	ItemType interface{} `json:"item_type,omitempty"`
	// The link params to be passed in URL the query string.
	Params interface{} `json:"params,omitempty"`
	// Time at which this resource was disabled.
	DisabledAt interface{} `json:"disabled_at,omitempty"`
	// Time at which the resource was created.
	CreatedAt interface{} `json:"created_at,omitempty"`
	// Time at which the resource was last updated.
	UpdatedAt interface{} `json:"updated_at,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference interface{} `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code.
	ReferenceOrigin interface{} `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata interface{} `json:"metadata,omitempty"`
}

// NewGETLinksLinkId200ResponseDataAttributes instantiates a new GETLinksLinkId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGETLinksLinkId200ResponseDataAttributes() *GETLinksLinkId200ResponseDataAttributes {
	this := GETLinksLinkId200ResponseDataAttributes{}
	return &this
}

// NewGETLinksLinkId200ResponseDataAttributesWithDefaults instantiates a new GETLinksLinkId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGETLinksLinkId200ResponseDataAttributesWithDefaults() *GETLinksLinkId200ResponseDataAttributes {
	this := GETLinksLinkId200ResponseDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasName() bool {
	if o != nil && IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given interface{} and assigns it to the Name field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetName(v interface{}) {
	o.Name = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetClientId() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetClientIdOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return &o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasClientId() bool {
	if o != nil && IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given interface{} and assigns it to the ClientId field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetClientId(v interface{}) {
	o.ClientId = v
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetScope() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetScopeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasScope() bool {
	if o != nil && IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given interface{} and assigns it to the Scope field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetScope(v interface{}) {
	o.Scope = v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetStartsAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetStartsAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return &o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasStartsAt() bool {
	if o != nil && IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given interface{} and assigns it to the StartsAt field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetStartsAt(v interface{}) {
	o.StartsAt = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetExpiresAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetExpiresAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasExpiresAt() bool {
	if o != nil && IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given interface{} and assigns it to the ExpiresAt field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetExpiresAt(v interface{}) {
	o.ExpiresAt = v
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetActive() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetActiveOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return &o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasActive() bool {
	if o != nil && IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given interface{} and assigns it to the Active field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetActive(v interface{}) {
	o.Active = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetStatus(v interface{}) {
	o.Status = v
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetDomain() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetDomainOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return &o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasDomain() bool {
	if o != nil && IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given interface{} and assigns it to the Domain field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetDomain(v interface{}) {
	o.Domain = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetUrl() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetUrlOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return &o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasUrl() bool {
	if o != nil && IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given interface{} and assigns it to the Url field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetUrl(v interface{}) {
	o.Url = v
}

// GetItemType returns the ItemType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetItemType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetItemTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ItemType) {
		return nil, false
	}
	return &o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasItemType() bool {
	if o != nil && IsNil(o.ItemType) {
		return true
	}

	return false
}

// SetItemType gets a reference to the given interface{} and assigns it to the ItemType field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetItemType(v interface{}) {
	o.ItemType = v
}

// GetParams returns the Params field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetParams() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetParamsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return &o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasParams() bool {
	if o != nil && IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given interface{} and assigns it to the Params field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetParams(v interface{}) {
	o.Params = v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetDisabledAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetDisabledAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.DisabledAt) {
		return nil, false
	}
	return &o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasDisabledAt() bool {
	if o != nil && IsNil(o.DisabledAt) {
		return true
	}

	return false
}

// SetDisabledAt gets a reference to the given interface{} and assigns it to the DisabledAt field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetDisabledAt(v interface{}) {
	o.DisabledAt = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetCreatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return &o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasCreatedAt() bool {
	if o != nil && IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given interface{} and assigns it to the CreatedAt field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetCreatedAt(v interface{}) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetUpdatedAt() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasUpdatedAt() bool {
	if o != nil && IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given interface{} and assigns it to the UpdatedAt field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetUpdatedAt(v interface{}) {
	o.UpdatedAt = v
}

// GetReference returns the Reference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetReference() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return &o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasReference() bool {
	if o != nil && IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given interface{} and assigns it to the Reference field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetReference(v interface{}) {
	o.Reference = v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetReferenceOrigin() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ReferenceOrigin) {
		return nil, false
	}
	return &o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && IsNil(o.ReferenceOrigin) {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given interface{} and assigns it to the ReferenceOrigin field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetReferenceOrigin(v interface{}) {
	o.ReferenceOrigin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GETLinksLinkId200ResponseDataAttributes) GetMetadata() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GETLinksLinkId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GETLinksLinkId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given interface{} and assigns it to the Metadata field.
func (o *GETLinksLinkId200ResponseDataAttributes) SetMetadata(v interface{}) {
	o.Metadata = v
}

func (o GETLinksLinkId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GETLinksLinkId200ResponseDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.StartsAt != nil {
		toSerialize["starts_at"] = o.StartsAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.ItemType != nil {
		toSerialize["item_type"] = o.ItemType
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
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
	return toSerialize, nil
}

type NullableGETLinksLinkId200ResponseDataAttributes struct {
	value *GETLinksLinkId200ResponseDataAttributes
	isSet bool
}

func (v NullableGETLinksLinkId200ResponseDataAttributes) Get() *GETLinksLinkId200ResponseDataAttributes {
	return v.value
}

func (v *NullableGETLinksLinkId200ResponseDataAttributes) Set(val *GETLinksLinkId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGETLinksLinkId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGETLinksLinkId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGETLinksLinkId200ResponseDataAttributes(val *GETLinksLinkId200ResponseDataAttributes) *NullableGETLinksLinkId200ResponseDataAttributes {
	return &NullableGETLinksLinkId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullableGETLinksLinkId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGETLinksLinkId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
