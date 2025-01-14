# POSTOrders201ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **interface{}** | The order identifier. Can be specified if unique within the organization (for enterprise plans only), default to numeric ID otherwise. Cannot be passed by sales channels. | [optional] 
**AffiliateCode** | Pointer to **interface{}** | The affiliate code, if any, to track commissions using any third party services. | [optional] 
**Autorefresh** | Pointer to **interface{}** | Save this attribute as &#39;false&#39; if you want prevent the order to be refreshed automatically at each change (much faster). | [optional] 
**PlaceAsync** | Pointer to **interface{}** | Save this attribute as &#39;true&#39; if you want perform the place asynchronously. Payment errors, if any, will be collected afterwards. | [optional] 
**Guest** | Pointer to **interface{}** | Indicates if the order has been placed as guest. | [optional] 
**CustomerEmail** | Pointer to **interface{}** | The email address of the associated customer. When creating or updating an order, this is a shortcut to find or create the associated customer by email. | [optional] 
**CustomerPassword** | Pointer to **interface{}** | The password of the associated customer. When creating or updating an order, this is a shortcut to sign up the associated customer. | [optional] 
**LanguageCode** | Pointer to **interface{}** | The preferred language code (ISO 639-1) to be used when communicating with the customer. This can be useful when sending the order to 3rd party marketing tools and CRMs. If the language is supported, the hosted checkout will be localized accordingly. | [optional] 
**FreightTaxable** | Pointer to **interface{}** | Indicates if taxes are applied to shipping costs. | [optional] 
**PaymentMethodTaxable** | Pointer to **interface{}** | Indicates if taxes are applied to payment methods costs. | [optional] 
**AdjustmentTaxable** | Pointer to **interface{}** | Indicates if taxes are applied to positive adjustments. | [optional] 
**GiftCardTaxable** | Pointer to **interface{}** | Indicates if taxes are applied to purchased gift cards. | [optional] 
**ShippingCountryCodeLock** | Pointer to **interface{}** | The country code that you want the shipping address to be locked to. This can be useful to make sure the shipping address belongs to a given shipping country, e.g. the one selected in a country selector page. Not relevant if order contains only digital products. | [optional] 
**CouponCode** | Pointer to **interface{}** | The coupon code to be used for the order. If valid, it triggers a promotion adding a discount line item to the order. | [optional] 
**GiftCardCode** | Pointer to **interface{}** | The gift card code (at least the first 8 characters) to be used for the order. If valid, it uses the gift card balance to pay for the order. | [optional] 
**CartUrl** | Pointer to **interface{}** | The cart url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**ReturnUrl** | Pointer to **interface{}** | The return url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**TermsUrl** | Pointer to **interface{}** | The terms and conditions url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**PrivacyUrl** | Pointer to **interface{}** | The privacy policy url on your site. If present, it will be used on our hosted checkout application. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewPOSTOrders201ResponseDataAttributes

`func NewPOSTOrders201ResponseDataAttributes() *POSTOrders201ResponseDataAttributes`

NewPOSTOrders201ResponseDataAttributes instantiates a new POSTOrders201ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTOrders201ResponseDataAttributesWithDefaults

`func NewPOSTOrders201ResponseDataAttributesWithDefaults() *POSTOrders201ResponseDataAttributes`

NewPOSTOrders201ResponseDataAttributesWithDefaults instantiates a new POSTOrders201ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *POSTOrders201ResponseDataAttributes) GetNumber() interface{}`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *POSTOrders201ResponseDataAttributes) GetNumberOk() (*interface{}, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *POSTOrders201ResponseDataAttributes) SetNumber(v interface{})`

SetNumber sets Number field to given value.

### HasNumber

`func (o *POSTOrders201ResponseDataAttributes) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *POSTOrders201ResponseDataAttributes) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *POSTOrders201ResponseDataAttributes) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetAffiliateCode

`func (o *POSTOrders201ResponseDataAttributes) GetAffiliateCode() interface{}`

GetAffiliateCode returns the AffiliateCode field if non-nil, zero value otherwise.

### GetAffiliateCodeOk

`func (o *POSTOrders201ResponseDataAttributes) GetAffiliateCodeOk() (*interface{}, bool)`

GetAffiliateCodeOk returns a tuple with the AffiliateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliateCode

`func (o *POSTOrders201ResponseDataAttributes) SetAffiliateCode(v interface{})`

SetAffiliateCode sets AffiliateCode field to given value.

### HasAffiliateCode

`func (o *POSTOrders201ResponseDataAttributes) HasAffiliateCode() bool`

HasAffiliateCode returns a boolean if a field has been set.

### SetAffiliateCodeNil

`func (o *POSTOrders201ResponseDataAttributes) SetAffiliateCodeNil(b bool)`

 SetAffiliateCodeNil sets the value for AffiliateCode to be an explicit nil

### UnsetAffiliateCode
`func (o *POSTOrders201ResponseDataAttributes) UnsetAffiliateCode()`

UnsetAffiliateCode ensures that no value is present for AffiliateCode, not even an explicit nil
### GetAutorefresh

`func (o *POSTOrders201ResponseDataAttributes) GetAutorefresh() interface{}`

GetAutorefresh returns the Autorefresh field if non-nil, zero value otherwise.

### GetAutorefreshOk

`func (o *POSTOrders201ResponseDataAttributes) GetAutorefreshOk() (*interface{}, bool)`

GetAutorefreshOk returns a tuple with the Autorefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutorefresh

`func (o *POSTOrders201ResponseDataAttributes) SetAutorefresh(v interface{})`

SetAutorefresh sets Autorefresh field to given value.

### HasAutorefresh

`func (o *POSTOrders201ResponseDataAttributes) HasAutorefresh() bool`

HasAutorefresh returns a boolean if a field has been set.

### SetAutorefreshNil

`func (o *POSTOrders201ResponseDataAttributes) SetAutorefreshNil(b bool)`

 SetAutorefreshNil sets the value for Autorefresh to be an explicit nil

### UnsetAutorefresh
`func (o *POSTOrders201ResponseDataAttributes) UnsetAutorefresh()`

UnsetAutorefresh ensures that no value is present for Autorefresh, not even an explicit nil
### GetPlaceAsync

`func (o *POSTOrders201ResponseDataAttributes) GetPlaceAsync() interface{}`

GetPlaceAsync returns the PlaceAsync field if non-nil, zero value otherwise.

### GetPlaceAsyncOk

`func (o *POSTOrders201ResponseDataAttributes) GetPlaceAsyncOk() (*interface{}, bool)`

GetPlaceAsyncOk returns a tuple with the PlaceAsync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceAsync

`func (o *POSTOrders201ResponseDataAttributes) SetPlaceAsync(v interface{})`

SetPlaceAsync sets PlaceAsync field to given value.

### HasPlaceAsync

`func (o *POSTOrders201ResponseDataAttributes) HasPlaceAsync() bool`

HasPlaceAsync returns a boolean if a field has been set.

### SetPlaceAsyncNil

`func (o *POSTOrders201ResponseDataAttributes) SetPlaceAsyncNil(b bool)`

 SetPlaceAsyncNil sets the value for PlaceAsync to be an explicit nil

### UnsetPlaceAsync
`func (o *POSTOrders201ResponseDataAttributes) UnsetPlaceAsync()`

UnsetPlaceAsync ensures that no value is present for PlaceAsync, not even an explicit nil
### GetGuest

`func (o *POSTOrders201ResponseDataAttributes) GetGuest() interface{}`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *POSTOrders201ResponseDataAttributes) GetGuestOk() (*interface{}, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *POSTOrders201ResponseDataAttributes) SetGuest(v interface{})`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *POSTOrders201ResponseDataAttributes) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### SetGuestNil

`func (o *POSTOrders201ResponseDataAttributes) SetGuestNil(b bool)`

 SetGuestNil sets the value for Guest to be an explicit nil

### UnsetGuest
`func (o *POSTOrders201ResponseDataAttributes) UnsetGuest()`

UnsetGuest ensures that no value is present for Guest, not even an explicit nil
### GetCustomerEmail

`func (o *POSTOrders201ResponseDataAttributes) GetCustomerEmail() interface{}`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *POSTOrders201ResponseDataAttributes) GetCustomerEmailOk() (*interface{}, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *POSTOrders201ResponseDataAttributes) SetCustomerEmail(v interface{})`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *POSTOrders201ResponseDataAttributes) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### SetCustomerEmailNil

`func (o *POSTOrders201ResponseDataAttributes) SetCustomerEmailNil(b bool)`

 SetCustomerEmailNil sets the value for CustomerEmail to be an explicit nil

### UnsetCustomerEmail
`func (o *POSTOrders201ResponseDataAttributes) UnsetCustomerEmail()`

UnsetCustomerEmail ensures that no value is present for CustomerEmail, not even an explicit nil
### GetCustomerPassword

`func (o *POSTOrders201ResponseDataAttributes) GetCustomerPassword() interface{}`

GetCustomerPassword returns the CustomerPassword field if non-nil, zero value otherwise.

### GetCustomerPasswordOk

`func (o *POSTOrders201ResponseDataAttributes) GetCustomerPasswordOk() (*interface{}, bool)`

GetCustomerPasswordOk returns a tuple with the CustomerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPassword

`func (o *POSTOrders201ResponseDataAttributes) SetCustomerPassword(v interface{})`

SetCustomerPassword sets CustomerPassword field to given value.

### HasCustomerPassword

`func (o *POSTOrders201ResponseDataAttributes) HasCustomerPassword() bool`

HasCustomerPassword returns a boolean if a field has been set.

### SetCustomerPasswordNil

`func (o *POSTOrders201ResponseDataAttributes) SetCustomerPasswordNil(b bool)`

 SetCustomerPasswordNil sets the value for CustomerPassword to be an explicit nil

### UnsetCustomerPassword
`func (o *POSTOrders201ResponseDataAttributes) UnsetCustomerPassword()`

UnsetCustomerPassword ensures that no value is present for CustomerPassword, not even an explicit nil
### GetLanguageCode

`func (o *POSTOrders201ResponseDataAttributes) GetLanguageCode() interface{}`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *POSTOrders201ResponseDataAttributes) GetLanguageCodeOk() (*interface{}, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *POSTOrders201ResponseDataAttributes) SetLanguageCode(v interface{})`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *POSTOrders201ResponseDataAttributes) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### SetLanguageCodeNil

`func (o *POSTOrders201ResponseDataAttributes) SetLanguageCodeNil(b bool)`

 SetLanguageCodeNil sets the value for LanguageCode to be an explicit nil

### UnsetLanguageCode
`func (o *POSTOrders201ResponseDataAttributes) UnsetLanguageCode()`

UnsetLanguageCode ensures that no value is present for LanguageCode, not even an explicit nil
### GetFreightTaxable

`func (o *POSTOrders201ResponseDataAttributes) GetFreightTaxable() interface{}`

GetFreightTaxable returns the FreightTaxable field if non-nil, zero value otherwise.

### GetFreightTaxableOk

`func (o *POSTOrders201ResponseDataAttributes) GetFreightTaxableOk() (*interface{}, bool)`

GetFreightTaxableOk returns a tuple with the FreightTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreightTaxable

`func (o *POSTOrders201ResponseDataAttributes) SetFreightTaxable(v interface{})`

SetFreightTaxable sets FreightTaxable field to given value.

### HasFreightTaxable

`func (o *POSTOrders201ResponseDataAttributes) HasFreightTaxable() bool`

HasFreightTaxable returns a boolean if a field has been set.

### SetFreightTaxableNil

`func (o *POSTOrders201ResponseDataAttributes) SetFreightTaxableNil(b bool)`

 SetFreightTaxableNil sets the value for FreightTaxable to be an explicit nil

### UnsetFreightTaxable
`func (o *POSTOrders201ResponseDataAttributes) UnsetFreightTaxable()`

UnsetFreightTaxable ensures that no value is present for FreightTaxable, not even an explicit nil
### GetPaymentMethodTaxable

`func (o *POSTOrders201ResponseDataAttributes) GetPaymentMethodTaxable() interface{}`

GetPaymentMethodTaxable returns the PaymentMethodTaxable field if non-nil, zero value otherwise.

### GetPaymentMethodTaxableOk

`func (o *POSTOrders201ResponseDataAttributes) GetPaymentMethodTaxableOk() (*interface{}, bool)`

GetPaymentMethodTaxableOk returns a tuple with the PaymentMethodTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodTaxable

`func (o *POSTOrders201ResponseDataAttributes) SetPaymentMethodTaxable(v interface{})`

SetPaymentMethodTaxable sets PaymentMethodTaxable field to given value.

### HasPaymentMethodTaxable

`func (o *POSTOrders201ResponseDataAttributes) HasPaymentMethodTaxable() bool`

HasPaymentMethodTaxable returns a boolean if a field has been set.

### SetPaymentMethodTaxableNil

`func (o *POSTOrders201ResponseDataAttributes) SetPaymentMethodTaxableNil(b bool)`

 SetPaymentMethodTaxableNil sets the value for PaymentMethodTaxable to be an explicit nil

### UnsetPaymentMethodTaxable
`func (o *POSTOrders201ResponseDataAttributes) UnsetPaymentMethodTaxable()`

UnsetPaymentMethodTaxable ensures that no value is present for PaymentMethodTaxable, not even an explicit nil
### GetAdjustmentTaxable

`func (o *POSTOrders201ResponseDataAttributes) GetAdjustmentTaxable() interface{}`

GetAdjustmentTaxable returns the AdjustmentTaxable field if non-nil, zero value otherwise.

### GetAdjustmentTaxableOk

`func (o *POSTOrders201ResponseDataAttributes) GetAdjustmentTaxableOk() (*interface{}, bool)`

GetAdjustmentTaxableOk returns a tuple with the AdjustmentTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTaxable

`func (o *POSTOrders201ResponseDataAttributes) SetAdjustmentTaxable(v interface{})`

SetAdjustmentTaxable sets AdjustmentTaxable field to given value.

### HasAdjustmentTaxable

`func (o *POSTOrders201ResponseDataAttributes) HasAdjustmentTaxable() bool`

HasAdjustmentTaxable returns a boolean if a field has been set.

### SetAdjustmentTaxableNil

`func (o *POSTOrders201ResponseDataAttributes) SetAdjustmentTaxableNil(b bool)`

 SetAdjustmentTaxableNil sets the value for AdjustmentTaxable to be an explicit nil

### UnsetAdjustmentTaxable
`func (o *POSTOrders201ResponseDataAttributes) UnsetAdjustmentTaxable()`

UnsetAdjustmentTaxable ensures that no value is present for AdjustmentTaxable, not even an explicit nil
### GetGiftCardTaxable

`func (o *POSTOrders201ResponseDataAttributes) GetGiftCardTaxable() interface{}`

GetGiftCardTaxable returns the GiftCardTaxable field if non-nil, zero value otherwise.

### GetGiftCardTaxableOk

`func (o *POSTOrders201ResponseDataAttributes) GetGiftCardTaxableOk() (*interface{}, bool)`

GetGiftCardTaxableOk returns a tuple with the GiftCardTaxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardTaxable

`func (o *POSTOrders201ResponseDataAttributes) SetGiftCardTaxable(v interface{})`

SetGiftCardTaxable sets GiftCardTaxable field to given value.

### HasGiftCardTaxable

`func (o *POSTOrders201ResponseDataAttributes) HasGiftCardTaxable() bool`

HasGiftCardTaxable returns a boolean if a field has been set.

### SetGiftCardTaxableNil

`func (o *POSTOrders201ResponseDataAttributes) SetGiftCardTaxableNil(b bool)`

 SetGiftCardTaxableNil sets the value for GiftCardTaxable to be an explicit nil

### UnsetGiftCardTaxable
`func (o *POSTOrders201ResponseDataAttributes) UnsetGiftCardTaxable()`

UnsetGiftCardTaxable ensures that no value is present for GiftCardTaxable, not even an explicit nil
### GetShippingCountryCodeLock

`func (o *POSTOrders201ResponseDataAttributes) GetShippingCountryCodeLock() interface{}`

GetShippingCountryCodeLock returns the ShippingCountryCodeLock field if non-nil, zero value otherwise.

### GetShippingCountryCodeLockOk

`func (o *POSTOrders201ResponseDataAttributes) GetShippingCountryCodeLockOk() (*interface{}, bool)`

GetShippingCountryCodeLockOk returns a tuple with the ShippingCountryCodeLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCountryCodeLock

`func (o *POSTOrders201ResponseDataAttributes) SetShippingCountryCodeLock(v interface{})`

SetShippingCountryCodeLock sets ShippingCountryCodeLock field to given value.

### HasShippingCountryCodeLock

`func (o *POSTOrders201ResponseDataAttributes) HasShippingCountryCodeLock() bool`

HasShippingCountryCodeLock returns a boolean if a field has been set.

### SetShippingCountryCodeLockNil

`func (o *POSTOrders201ResponseDataAttributes) SetShippingCountryCodeLockNil(b bool)`

 SetShippingCountryCodeLockNil sets the value for ShippingCountryCodeLock to be an explicit nil

### UnsetShippingCountryCodeLock
`func (o *POSTOrders201ResponseDataAttributes) UnsetShippingCountryCodeLock()`

UnsetShippingCountryCodeLock ensures that no value is present for ShippingCountryCodeLock, not even an explicit nil
### GetCouponCode

`func (o *POSTOrders201ResponseDataAttributes) GetCouponCode() interface{}`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *POSTOrders201ResponseDataAttributes) GetCouponCodeOk() (*interface{}, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *POSTOrders201ResponseDataAttributes) SetCouponCode(v interface{})`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *POSTOrders201ResponseDataAttributes) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### SetCouponCodeNil

`func (o *POSTOrders201ResponseDataAttributes) SetCouponCodeNil(b bool)`

 SetCouponCodeNil sets the value for CouponCode to be an explicit nil

### UnsetCouponCode
`func (o *POSTOrders201ResponseDataAttributes) UnsetCouponCode()`

UnsetCouponCode ensures that no value is present for CouponCode, not even an explicit nil
### GetGiftCardCode

`func (o *POSTOrders201ResponseDataAttributes) GetGiftCardCode() interface{}`

GetGiftCardCode returns the GiftCardCode field if non-nil, zero value otherwise.

### GetGiftCardCodeOk

`func (o *POSTOrders201ResponseDataAttributes) GetGiftCardCodeOk() (*interface{}, bool)`

GetGiftCardCodeOk returns a tuple with the GiftCardCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardCode

`func (o *POSTOrders201ResponseDataAttributes) SetGiftCardCode(v interface{})`

SetGiftCardCode sets GiftCardCode field to given value.

### HasGiftCardCode

`func (o *POSTOrders201ResponseDataAttributes) HasGiftCardCode() bool`

HasGiftCardCode returns a boolean if a field has been set.

### SetGiftCardCodeNil

`func (o *POSTOrders201ResponseDataAttributes) SetGiftCardCodeNil(b bool)`

 SetGiftCardCodeNil sets the value for GiftCardCode to be an explicit nil

### UnsetGiftCardCode
`func (o *POSTOrders201ResponseDataAttributes) UnsetGiftCardCode()`

UnsetGiftCardCode ensures that no value is present for GiftCardCode, not even an explicit nil
### GetCartUrl

`func (o *POSTOrders201ResponseDataAttributes) GetCartUrl() interface{}`

GetCartUrl returns the CartUrl field if non-nil, zero value otherwise.

### GetCartUrlOk

`func (o *POSTOrders201ResponseDataAttributes) GetCartUrlOk() (*interface{}, bool)`

GetCartUrlOk returns a tuple with the CartUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartUrl

`func (o *POSTOrders201ResponseDataAttributes) SetCartUrl(v interface{})`

SetCartUrl sets CartUrl field to given value.

### HasCartUrl

`func (o *POSTOrders201ResponseDataAttributes) HasCartUrl() bool`

HasCartUrl returns a boolean if a field has been set.

### SetCartUrlNil

`func (o *POSTOrders201ResponseDataAttributes) SetCartUrlNil(b bool)`

 SetCartUrlNil sets the value for CartUrl to be an explicit nil

### UnsetCartUrl
`func (o *POSTOrders201ResponseDataAttributes) UnsetCartUrl()`

UnsetCartUrl ensures that no value is present for CartUrl, not even an explicit nil
### GetReturnUrl

`func (o *POSTOrders201ResponseDataAttributes) GetReturnUrl() interface{}`

GetReturnUrl returns the ReturnUrl field if non-nil, zero value otherwise.

### GetReturnUrlOk

`func (o *POSTOrders201ResponseDataAttributes) GetReturnUrlOk() (*interface{}, bool)`

GetReturnUrlOk returns a tuple with the ReturnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnUrl

`func (o *POSTOrders201ResponseDataAttributes) SetReturnUrl(v interface{})`

SetReturnUrl sets ReturnUrl field to given value.

### HasReturnUrl

`func (o *POSTOrders201ResponseDataAttributes) HasReturnUrl() bool`

HasReturnUrl returns a boolean if a field has been set.

### SetReturnUrlNil

`func (o *POSTOrders201ResponseDataAttributes) SetReturnUrlNil(b bool)`

 SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil

### UnsetReturnUrl
`func (o *POSTOrders201ResponseDataAttributes) UnsetReturnUrl()`

UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
### GetTermsUrl

`func (o *POSTOrders201ResponseDataAttributes) GetTermsUrl() interface{}`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *POSTOrders201ResponseDataAttributes) GetTermsUrlOk() (*interface{}, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *POSTOrders201ResponseDataAttributes) SetTermsUrl(v interface{})`

SetTermsUrl sets TermsUrl field to given value.

### HasTermsUrl

`func (o *POSTOrders201ResponseDataAttributes) HasTermsUrl() bool`

HasTermsUrl returns a boolean if a field has been set.

### SetTermsUrlNil

`func (o *POSTOrders201ResponseDataAttributes) SetTermsUrlNil(b bool)`

 SetTermsUrlNil sets the value for TermsUrl to be an explicit nil

### UnsetTermsUrl
`func (o *POSTOrders201ResponseDataAttributes) UnsetTermsUrl()`

UnsetTermsUrl ensures that no value is present for TermsUrl, not even an explicit nil
### GetPrivacyUrl

`func (o *POSTOrders201ResponseDataAttributes) GetPrivacyUrl() interface{}`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *POSTOrders201ResponseDataAttributes) GetPrivacyUrlOk() (*interface{}, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *POSTOrders201ResponseDataAttributes) SetPrivacyUrl(v interface{})`

SetPrivacyUrl sets PrivacyUrl field to given value.

### HasPrivacyUrl

`func (o *POSTOrders201ResponseDataAttributes) HasPrivacyUrl() bool`

HasPrivacyUrl returns a boolean if a field has been set.

### SetPrivacyUrlNil

`func (o *POSTOrders201ResponseDataAttributes) SetPrivacyUrlNil(b bool)`

 SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil

### UnsetPrivacyUrl
`func (o *POSTOrders201ResponseDataAttributes) UnsetPrivacyUrl()`

UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
### GetReference

`func (o *POSTOrders201ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *POSTOrders201ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *POSTOrders201ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *POSTOrders201ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *POSTOrders201ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *POSTOrders201ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *POSTOrders201ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *POSTOrders201ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *POSTOrders201ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *POSTOrders201ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *POSTOrders201ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *POSTOrders201ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *POSTOrders201ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *POSTOrders201ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *POSTOrders201ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *POSTOrders201ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *POSTOrders201ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *POSTOrders201ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


