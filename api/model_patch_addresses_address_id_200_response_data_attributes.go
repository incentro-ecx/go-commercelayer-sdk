/*
Commerce Layer API

Headless Commerce for Global Brands.

API version: 3.0.1
Contact: support@commercelayer.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PATCHAddressesAddressId200ResponseDataAttributes struct for PATCHAddressesAddressId200ResponseDataAttributes
type PATCHAddressesAddressId200ResponseDataAttributes struct {
	// Indicates if it's a business or a personal address
	Business *bool `json:"business,omitempty"`
	// Address first name (personal)
	FirstName *string `json:"first_name,omitempty"`
	// Address last name (personal)
	LastName *string `json:"last_name,omitempty"`
	// Address company name (business)
	Company *string `json:"company,omitempty"`
	// Address line 1, i.e. Street address, PO Box
	Line1 *string `json:"line_1,omitempty"`
	// Address line 2, i.e. Apartment, Suite, Building
	Line2 *string `json:"line_2,omitempty"`
	// Address city
	City *string `json:"city,omitempty"`
	// ZIP or postal code
	ZipCode *string `json:"zip_code,omitempty"`
	// State, province or region code.
	StateCode *string `json:"state_code,omitempty"`
	// The international 2-letter country code as defined by the ISO 3166-1 standard
	CountryCode *string `json:"country_code,omitempty"`
	// Phone number (including extension).
	Phone *string `json:"phone,omitempty"`
	// Email address.
	Email *string `json:"email,omitempty"`
	// A free notes attached to the address. When used as a shipping address, this can be useful to let the customers add specific delivery instructions.
	Notes *string `json:"notes,omitempty"`
	// The address geocoded latitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order's market.
	Lat *float32 `json:"lat,omitempty"`
	// The address geocoded longitude. This is automatically generated when creating a shipping/billing address for an order and a valid geocoder is attached to the order's market.
	Lng *float32 `json:"lng,omitempty"`
	// Customer's billing information (i.e. VAT number, codice fiscale)
	BillingInfo *string `json:"billing_info,omitempty"`
	// A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever.
	Reference *string `json:"reference,omitempty"`
	// Any identifier of the third party system that defines the reference code
	ReferenceOrigin *string `json:"reference_origin,omitempty"`
	// Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewPATCHAddressesAddressId200ResponseDataAttributes instantiates a new PATCHAddressesAddressId200ResponseDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPATCHAddressesAddressId200ResponseDataAttributes() *PATCHAddressesAddressId200ResponseDataAttributes {
	this := PATCHAddressesAddressId200ResponseDataAttributes{}
	return &this
}

// NewPATCHAddressesAddressId200ResponseDataAttributesWithDefaults instantiates a new PATCHAddressesAddressId200ResponseDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPATCHAddressesAddressId200ResponseDataAttributesWithDefaults() *PATCHAddressesAddressId200ResponseDataAttributes {
	this := PATCHAddressesAddressId200ResponseDataAttributes{}
	return &this
}

// GetBusiness returns the Business field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetBusiness() bool {
	if o == nil || o.Business == nil {
		var ret bool
		return ret
	}
	return *o.Business
}

// GetBusinessOk returns a tuple with the Business field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetBusinessOk() (*bool, bool) {
	if o == nil || o.Business == nil {
		return nil, false
	}
	return o.Business, true
}

// HasBusiness returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasBusiness() bool {
	if o != nil && o.Business != nil {
		return true
	}

	return false
}

// SetBusiness gets a reference to the given bool and assigns it to the Business field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetBusiness(v bool) {
	o.Business = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLastName(v string) {
	o.LastName = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetCompany(v string) {
	o.Company = &v
}

// GetLine1 returns the Line1 field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLine1() string {
	if o == nil || o.Line1 == nil {
		var ret string
		return ret
	}
	return *o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLine1Ok() (*string, bool) {
	if o == nil || o.Line1 == nil {
		return nil, false
	}
	return o.Line1, true
}

// HasLine1 returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLine1() bool {
	if o != nil && o.Line1 != nil {
		return true
	}

	return false
}

// SetLine1 gets a reference to the given string and assigns it to the Line1 field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLine1(v string) {
	o.Line1 = &v
}

// GetLine2 returns the Line2 field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLine2() string {
	if o == nil || o.Line2 == nil {
		var ret string
		return ret
	}
	return *o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLine2Ok() (*string, bool) {
	if o == nil || o.Line2 == nil {
		return nil, false
	}
	return o.Line2, true
}

// HasLine2 returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLine2() bool {
	if o != nil && o.Line2 != nil {
		return true
	}

	return false
}

// SetLine2 gets a reference to the given string and assigns it to the Line2 field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLine2(v string) {
	o.Line2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetCity(v string) {
	o.City = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetZipCode() string {
	if o == nil || o.ZipCode == nil {
		var ret string
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetZipCodeOk() (*string, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given string and assigns it to the ZipCode field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetZipCode(v string) {
	o.ZipCode = &v
}

// GetStateCode returns the StateCode field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetStateCode() string {
	if o == nil || o.StateCode == nil {
		var ret string
		return ret
	}
	return *o.StateCode
}

// GetStateCodeOk returns a tuple with the StateCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetStateCodeOk() (*string, bool) {
	if o == nil || o.StateCode == nil {
		return nil, false
	}
	return o.StateCode, true
}

// HasStateCode returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasStateCode() bool {
	if o != nil && o.StateCode != nil {
		return true
	}

	return false
}

// SetStateCode gets a reference to the given string and assigns it to the StateCode field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetStateCode(v string) {
	o.StateCode = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCountryCode() string {
	if o == nil || o.CountryCode == nil {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetCountryCodeOk() (*string, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetPhone() string {
	if o == nil || o.Phone == nil {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetPhoneOk() (*string, bool) {
	if o == nil || o.Phone == nil {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasPhone() bool {
	if o != nil && o.Phone != nil {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetNotes(v string) {
	o.Notes = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLat() float32 {
	if o == nil || o.Lat == nil {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLatOk() (*float32, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLng() float32 {
	if o == nil || o.Lng == nil {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetLngOk() (*float32, bool) {
	if o == nil || o.Lng == nil {
		return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasLng() bool {
	if o != nil && o.Lng != nil {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetLng(v float32) {
	o.Lng = &v
}

// GetBillingInfo returns the BillingInfo field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetBillingInfo() string {
	if o == nil || o.BillingInfo == nil {
		var ret string
		return ret
	}
	return *o.BillingInfo
}

// GetBillingInfoOk returns a tuple with the BillingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetBillingInfoOk() (*string, bool) {
	if o == nil || o.BillingInfo == nil {
		return nil, false
	}
	return o.BillingInfo, true
}

// HasBillingInfo returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasBillingInfo() bool {
	if o != nil && o.BillingInfo != nil {
		return true
	}

	return false
}

// SetBillingInfo gets a reference to the given string and assigns it to the BillingInfo field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetBillingInfo(v string) {
	o.BillingInfo = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetReference(v string) {
	o.Reference = &v
}

// GetReferenceOrigin returns the ReferenceOrigin field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetReferenceOrigin() string {
	if o == nil || o.ReferenceOrigin == nil {
		var ret string
		return ret
	}
	return *o.ReferenceOrigin
}

// GetReferenceOriginOk returns a tuple with the ReferenceOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetReferenceOriginOk() (*string, bool) {
	if o == nil || o.ReferenceOrigin == nil {
		return nil, false
	}
	return o.ReferenceOrigin, true
}

// HasReferenceOrigin returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasReferenceOrigin() bool {
	if o != nil && o.ReferenceOrigin != nil {
		return true
	}

	return false
}

// SetReferenceOrigin gets a reference to the given string and assigns it to the ReferenceOrigin field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetReferenceOrigin(v string) {
	o.ReferenceOrigin = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *PATCHAddressesAddressId200ResponseDataAttributes) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o PATCHAddressesAddressId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Business != nil {
		toSerialize["business"] = o.Business
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Line1 != nil {
		toSerialize["line_1"] = o.Line1
	}
	if o.Line2 != nil {
		toSerialize["line_2"] = o.Line2
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.ZipCode != nil {
		toSerialize["zip_code"] = o.ZipCode
	}
	if o.StateCode != nil {
		toSerialize["state_code"] = o.StateCode
	}
	if o.CountryCode != nil {
		toSerialize["country_code"] = o.CountryCode
	}
	if o.Phone != nil {
		toSerialize["phone"] = o.Phone
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Lng != nil {
		toSerialize["lng"] = o.Lng
	}
	if o.BillingInfo != nil {
		toSerialize["billing_info"] = o.BillingInfo
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

type NullablePATCHAddressesAddressId200ResponseDataAttributes struct {
	value *PATCHAddressesAddressId200ResponseDataAttributes
	isSet bool
}

func (v NullablePATCHAddressesAddressId200ResponseDataAttributes) Get() *PATCHAddressesAddressId200ResponseDataAttributes {
	return v.value
}

func (v *NullablePATCHAddressesAddressId200ResponseDataAttributes) Set(val *PATCHAddressesAddressId200ResponseDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullablePATCHAddressesAddressId200ResponseDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullablePATCHAddressesAddressId200ResponseDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePATCHAddressesAddressId200ResponseDataAttributes(val *PATCHAddressesAddressId200ResponseDataAttributes) *NullablePATCHAddressesAddressId200ResponseDataAttributes {
	return &NullablePATCHAddressesAddressId200ResponseDataAttributes{value: val, isSet: true}
}

func (v NullablePATCHAddressesAddressId200ResponseDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePATCHAddressesAddressId200ResponseDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}