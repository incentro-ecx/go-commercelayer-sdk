# GETOrganizationOrganizationId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **interface{}** | The organization&#39;s internal name. | [optional] 
**Slug** | Pointer to **interface{}** | The organization&#39;s slug name. | [optional] 
**Domain** | Pointer to **interface{}** | The organization&#39;s domain. | [optional] 
**SupportPhone** | Pointer to **interface{}** | The organization&#39;s support phone. | [optional] 
**SupportEmail** | Pointer to **interface{}** | The organization&#39;s support email. | [optional] 
**LogoUrl** | Pointer to **interface{}** | The URL to the organization&#39;s logo. | [optional] 
**FaviconUrl** | Pointer to **interface{}** | The URL to the organization&#39;s favicon. | [optional] 
**PrimaryColor** | Pointer to **interface{}** | The organization&#39;s primary color. | [optional] 
**GtmId** | Pointer to **interface{}** | The organization&#39;s Google Tag Manager ID. | [optional] 
**GtmIdTest** | Pointer to **interface{}** | The organization&#39;s Google Tag Manager ID for test. | [optional] 
**Config** | Pointer to **interface{}** | The organization&#39;s configuration. | [optional] 
**ApiAuthRedirect** | Pointer to **interface{}** | Enables the redirect on the new Auth API. | [optional] 
**ApiRulesEngine** | Pointer to **interface{}** | Enables the rules engine for flex promotions and price list rules. | [optional] 
**ApiNewAuth** | Pointer to **interface{}** | Forces the usage of the new Authentication API. | [optional] 
**ApiPurgeSingleResource** | Pointer to **interface{}** | Enables the purge of cached single resources when list is purged. | [optional] 
**ApiMaxRegexLength** | Pointer to **interface{}** | The maximum length for the regular expressions, default is 5000. | [optional] 
**AddressesPhoneRequired** | Pointer to **interface{}** | Indicates if the phone attribute is required for addresses, default is true. | [optional] 
**OrdersAutorefreshCutoffTest** | Pointer to **interface{}** | The maximum number line items allowed for a test order before disabling the autorefresh option. | [optional] 
**OrdersAutorefreshCutoffLive** | Pointer to **interface{}** | The maximum number line items allowed for a live order before disabling the autorefresh option. | [optional] 
**OrdersNumberEditableTest** | Pointer to **interface{}** | Enables orders number editing as a string in test (for enterprise plans only). | [optional] 
**OrdersNumberEditableLive** | Pointer to **interface{}** | Enables orders number editing as a string in live (for enterprise plans only). | [optional] 
**OrdersNumberAsReference** | Pointer to **interface{}** | Enables to use the order number as payment reference on supported gateways. | [optional] 
**BundlesMaxItemsCount** | Pointer to **interface{}** | The maximum number of SKUs allowed for bundles, default is 10. | [optional] 
**CouponsMinCodeLength** | Pointer to **interface{}** | The minimum length for coupon code, default is 8. | [optional] 
**CouponsMaxCodeLength** | Pointer to **interface{}** | The maximum length for coupon code, default is 40. | [optional] 
**GiftCardsMinCodeLength** | Pointer to **interface{}** | The minimum length for gift card code, default is 8. | [optional] 
**GiftCardsMaxCodeLength** | Pointer to **interface{}** | The maximum length for gift card code, default is 40. | [optional] 
**CleanupsMaxConcurrentCount** | Pointer to **interface{}** | The maximum number of concurrent cleanups allowed for your organization, default is 10. | [optional] 
**ExportsMaxConcurrentCount** | Pointer to **interface{}** | The maximum number of concurrent exports allowed for your organization, default is 10. | [optional] 
**ImportsMaxConcurrentCount** | Pointer to **interface{}** | The maximum number of concurrent imports allowed for your organization, default is 10. | [optional] 
**ImportsPurgeCache** | Pointer to **interface{}** | Enables purging of cached resources upon succeeded imports. | [optional] 
**ImportsSkipErrors** | Pointer to **interface{}** | Disables the interruption of the import in case its errors exceeds the 10% threshold, default is false. | [optional] 
**PromotionsMaxConcurrentCount** | Pointer to **interface{}** | The maximum number of active concurrent promotions allowed for your organization, default is 10. | [optional] 
**ImportsTriggerWebhooks** | Pointer to **interface{}** | Enables triggering of webhooks during imports, default is false. | [optional] 
**DiscountEnginesEnabled** | Pointer to **interface{}** | Enables the use of an external discount engine in place of the standard one, default is false. | [optional] 
**DiscountEnginesErrors** | Pointer to **interface{}** | Enables raising of API errors in case of discount engine failure, default is false. | [optional] 
**TagsMaxNameLength** | Pointer to **interface{}** | The maximum length for the tag name, default is 25. | [optional] 
**TaxCalculatorsErrors** | Pointer to **interface{}** | Enables raising of API errors in case of tax calculation failure, default is false. | [optional] 
**ExternalPromotionsErrors** | Pointer to **interface{}** | Enables raising of API errors in case of external promotion failure, default is false. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETOrganizationOrganizationId200ResponseDataAttributes

`func NewGETOrganizationOrganizationId200ResponseDataAttributes() *GETOrganizationOrganizationId200ResponseDataAttributes`

NewGETOrganizationOrganizationId200ResponseDataAttributes instantiates a new GETOrganizationOrganizationId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETOrganizationOrganizationId200ResponseDataAttributesWithDefaults

`func NewGETOrganizationOrganizationId200ResponseDataAttributesWithDefaults() *GETOrganizationOrganizationId200ResponseDataAttributes`

NewGETOrganizationOrganizationId200ResponseDataAttributesWithDefaults instantiates a new GETOrganizationOrganizationId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSlug

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSlug() interface{}`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSlugOk() (*interface{}, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSlug(v interface{})`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetDomain

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDomain() interface{}`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDomainOk() (*interface{}, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDomain(v interface{})`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### SetDomainNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDomainNil(b bool)`

 SetDomainNil sets the value for Domain to be an explicit nil

### UnsetDomain
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetDomain()`

UnsetDomain ensures that no value is present for Domain, not even an explicit nil
### GetSupportPhone

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportPhone() interface{}`

GetSupportPhone returns the SupportPhone field if non-nil, zero value otherwise.

### GetSupportPhoneOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportPhoneOk() (*interface{}, bool)`

GetSupportPhoneOk returns a tuple with the SupportPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPhone

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSupportPhone(v interface{})`

SetSupportPhone sets SupportPhone field to given value.

### HasSupportPhone

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasSupportPhone() bool`

HasSupportPhone returns a boolean if a field has been set.

### SetSupportPhoneNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSupportPhoneNil(b bool)`

 SetSupportPhoneNil sets the value for SupportPhone to be an explicit nil

### UnsetSupportPhone
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetSupportPhone()`

UnsetSupportPhone ensures that no value is present for SupportPhone, not even an explicit nil
### GetSupportEmail

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportEmail() interface{}`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetSupportEmailOk() (*interface{}, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSupportEmail(v interface{})`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### SetSupportEmailNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetSupportEmailNil(b bool)`

 SetSupportEmailNil sets the value for SupportEmail to be an explicit nil

### UnsetSupportEmail
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetSupportEmail()`

UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
### GetLogoUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetLogoUrl() interface{}`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetLogoUrlOk() (*interface{}, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetLogoUrl(v interface{})`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### SetLogoUrlNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetFaviconUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetFaviconUrl() interface{}`

GetFaviconUrl returns the FaviconUrl field if non-nil, zero value otherwise.

### GetFaviconUrlOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetFaviconUrlOk() (*interface{}, bool)`

GetFaviconUrlOk returns a tuple with the FaviconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaviconUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetFaviconUrl(v interface{})`

SetFaviconUrl sets FaviconUrl field to given value.

### HasFaviconUrl

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasFaviconUrl() bool`

HasFaviconUrl returns a boolean if a field has been set.

### SetFaviconUrlNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetFaviconUrlNil(b bool)`

 SetFaviconUrlNil sets the value for FaviconUrl to be an explicit nil

### UnsetFaviconUrl
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetFaviconUrl()`

UnsetFaviconUrl ensures that no value is present for FaviconUrl, not even an explicit nil
### GetPrimaryColor

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetPrimaryColor() interface{}`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetPrimaryColorOk() (*interface{}, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetPrimaryColor(v interface{})`

SetPrimaryColor sets PrimaryColor field to given value.

### HasPrimaryColor

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasPrimaryColor() bool`

HasPrimaryColor returns a boolean if a field has been set.

### SetPrimaryColorNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetPrimaryColorNil(b bool)`

 SetPrimaryColorNil sets the value for PrimaryColor to be an explicit nil

### UnsetPrimaryColor
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetPrimaryColor()`

UnsetPrimaryColor ensures that no value is present for PrimaryColor, not even an explicit nil
### GetGtmId

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmId() interface{}`

GetGtmId returns the GtmId field if non-nil, zero value otherwise.

### GetGtmIdOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmIdOk() (*interface{}, bool)`

GetGtmIdOk returns a tuple with the GtmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtmId

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGtmId(v interface{})`

SetGtmId sets GtmId field to given value.

### HasGtmId

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasGtmId() bool`

HasGtmId returns a boolean if a field has been set.

### SetGtmIdNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGtmIdNil(b bool)`

 SetGtmIdNil sets the value for GtmId to be an explicit nil

### UnsetGtmId
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetGtmId()`

UnsetGtmId ensures that no value is present for GtmId, not even an explicit nil
### GetGtmIdTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmIdTest() interface{}`

GetGtmIdTest returns the GtmIdTest field if non-nil, zero value otherwise.

### GetGtmIdTestOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGtmIdTestOk() (*interface{}, bool)`

GetGtmIdTestOk returns a tuple with the GtmIdTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGtmIdTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGtmIdTest(v interface{})`

SetGtmIdTest sets GtmIdTest field to given value.

### HasGtmIdTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasGtmIdTest() bool`

HasGtmIdTest returns a boolean if a field has been set.

### SetGtmIdTestNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGtmIdTestNil(b bool)`

 SetGtmIdTestNil sets the value for GtmIdTest to be an explicit nil

### UnsetGtmIdTest
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetGtmIdTest()`

UnsetGtmIdTest ensures that no value is present for GtmIdTest, not even an explicit nil
### GetConfig

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetApiAuthRedirect

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiAuthRedirect() interface{}`

GetApiAuthRedirect returns the ApiAuthRedirect field if non-nil, zero value otherwise.

### GetApiAuthRedirectOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiAuthRedirectOk() (*interface{}, bool)`

GetApiAuthRedirectOk returns a tuple with the ApiAuthRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAuthRedirect

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiAuthRedirect(v interface{})`

SetApiAuthRedirect sets ApiAuthRedirect field to given value.

### HasApiAuthRedirect

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasApiAuthRedirect() bool`

HasApiAuthRedirect returns a boolean if a field has been set.

### SetApiAuthRedirectNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiAuthRedirectNil(b bool)`

 SetApiAuthRedirectNil sets the value for ApiAuthRedirect to be an explicit nil

### UnsetApiAuthRedirect
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetApiAuthRedirect()`

UnsetApiAuthRedirect ensures that no value is present for ApiAuthRedirect, not even an explicit nil
### GetApiRulesEngine

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiRulesEngine() interface{}`

GetApiRulesEngine returns the ApiRulesEngine field if non-nil, zero value otherwise.

### GetApiRulesEngineOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiRulesEngineOk() (*interface{}, bool)`

GetApiRulesEngineOk returns a tuple with the ApiRulesEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRulesEngine

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiRulesEngine(v interface{})`

SetApiRulesEngine sets ApiRulesEngine field to given value.

### HasApiRulesEngine

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasApiRulesEngine() bool`

HasApiRulesEngine returns a boolean if a field has been set.

### SetApiRulesEngineNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiRulesEngineNil(b bool)`

 SetApiRulesEngineNil sets the value for ApiRulesEngine to be an explicit nil

### UnsetApiRulesEngine
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetApiRulesEngine()`

UnsetApiRulesEngine ensures that no value is present for ApiRulesEngine, not even an explicit nil
### GetApiNewAuth

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiNewAuth() interface{}`

GetApiNewAuth returns the ApiNewAuth field if non-nil, zero value otherwise.

### GetApiNewAuthOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiNewAuthOk() (*interface{}, bool)`

GetApiNewAuthOk returns a tuple with the ApiNewAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiNewAuth

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiNewAuth(v interface{})`

SetApiNewAuth sets ApiNewAuth field to given value.

### HasApiNewAuth

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasApiNewAuth() bool`

HasApiNewAuth returns a boolean if a field has been set.

### SetApiNewAuthNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiNewAuthNil(b bool)`

 SetApiNewAuthNil sets the value for ApiNewAuth to be an explicit nil

### UnsetApiNewAuth
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetApiNewAuth()`

UnsetApiNewAuth ensures that no value is present for ApiNewAuth, not even an explicit nil
### GetApiPurgeSingleResource

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiPurgeSingleResource() interface{}`

GetApiPurgeSingleResource returns the ApiPurgeSingleResource field if non-nil, zero value otherwise.

### GetApiPurgeSingleResourceOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiPurgeSingleResourceOk() (*interface{}, bool)`

GetApiPurgeSingleResourceOk returns a tuple with the ApiPurgeSingleResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPurgeSingleResource

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiPurgeSingleResource(v interface{})`

SetApiPurgeSingleResource sets ApiPurgeSingleResource field to given value.

### HasApiPurgeSingleResource

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasApiPurgeSingleResource() bool`

HasApiPurgeSingleResource returns a boolean if a field has been set.

### SetApiPurgeSingleResourceNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiPurgeSingleResourceNil(b bool)`

 SetApiPurgeSingleResourceNil sets the value for ApiPurgeSingleResource to be an explicit nil

### UnsetApiPurgeSingleResource
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetApiPurgeSingleResource()`

UnsetApiPurgeSingleResource ensures that no value is present for ApiPurgeSingleResource, not even an explicit nil
### GetApiMaxRegexLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiMaxRegexLength() interface{}`

GetApiMaxRegexLength returns the ApiMaxRegexLength field if non-nil, zero value otherwise.

### GetApiMaxRegexLengthOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetApiMaxRegexLengthOk() (*interface{}, bool)`

GetApiMaxRegexLengthOk returns a tuple with the ApiMaxRegexLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMaxRegexLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiMaxRegexLength(v interface{})`

SetApiMaxRegexLength sets ApiMaxRegexLength field to given value.

### HasApiMaxRegexLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasApiMaxRegexLength() bool`

HasApiMaxRegexLength returns a boolean if a field has been set.

### SetApiMaxRegexLengthNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetApiMaxRegexLengthNil(b bool)`

 SetApiMaxRegexLengthNil sets the value for ApiMaxRegexLength to be an explicit nil

### UnsetApiMaxRegexLength
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetApiMaxRegexLength()`

UnsetApiMaxRegexLength ensures that no value is present for ApiMaxRegexLength, not even an explicit nil
### GetAddressesPhoneRequired

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAddressesPhoneRequired() interface{}`

GetAddressesPhoneRequired returns the AddressesPhoneRequired field if non-nil, zero value otherwise.

### GetAddressesPhoneRequiredOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetAddressesPhoneRequiredOk() (*interface{}, bool)`

GetAddressesPhoneRequiredOk returns a tuple with the AddressesPhoneRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressesPhoneRequired

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetAddressesPhoneRequired(v interface{})`

SetAddressesPhoneRequired sets AddressesPhoneRequired field to given value.

### HasAddressesPhoneRequired

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasAddressesPhoneRequired() bool`

HasAddressesPhoneRequired returns a boolean if a field has been set.

### SetAddressesPhoneRequiredNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetAddressesPhoneRequiredNil(b bool)`

 SetAddressesPhoneRequiredNil sets the value for AddressesPhoneRequired to be an explicit nil

### UnsetAddressesPhoneRequired
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetAddressesPhoneRequired()`

UnsetAddressesPhoneRequired ensures that no value is present for AddressesPhoneRequired, not even an explicit nil
### GetOrdersAutorefreshCutoffTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersAutorefreshCutoffTest() interface{}`

GetOrdersAutorefreshCutoffTest returns the OrdersAutorefreshCutoffTest field if non-nil, zero value otherwise.

### GetOrdersAutorefreshCutoffTestOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersAutorefreshCutoffTestOk() (*interface{}, bool)`

GetOrdersAutorefreshCutoffTestOk returns a tuple with the OrdersAutorefreshCutoffTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersAutorefreshCutoffTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersAutorefreshCutoffTest(v interface{})`

SetOrdersAutorefreshCutoffTest sets OrdersAutorefreshCutoffTest field to given value.

### HasOrdersAutorefreshCutoffTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasOrdersAutorefreshCutoffTest() bool`

HasOrdersAutorefreshCutoffTest returns a boolean if a field has been set.

### SetOrdersAutorefreshCutoffTestNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersAutorefreshCutoffTestNil(b bool)`

 SetOrdersAutorefreshCutoffTestNil sets the value for OrdersAutorefreshCutoffTest to be an explicit nil

### UnsetOrdersAutorefreshCutoffTest
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetOrdersAutorefreshCutoffTest()`

UnsetOrdersAutorefreshCutoffTest ensures that no value is present for OrdersAutorefreshCutoffTest, not even an explicit nil
### GetOrdersAutorefreshCutoffLive

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersAutorefreshCutoffLive() interface{}`

GetOrdersAutorefreshCutoffLive returns the OrdersAutorefreshCutoffLive field if non-nil, zero value otherwise.

### GetOrdersAutorefreshCutoffLiveOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersAutorefreshCutoffLiveOk() (*interface{}, bool)`

GetOrdersAutorefreshCutoffLiveOk returns a tuple with the OrdersAutorefreshCutoffLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersAutorefreshCutoffLive

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersAutorefreshCutoffLive(v interface{})`

SetOrdersAutorefreshCutoffLive sets OrdersAutorefreshCutoffLive field to given value.

### HasOrdersAutorefreshCutoffLive

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasOrdersAutorefreshCutoffLive() bool`

HasOrdersAutorefreshCutoffLive returns a boolean if a field has been set.

### SetOrdersAutorefreshCutoffLiveNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersAutorefreshCutoffLiveNil(b bool)`

 SetOrdersAutorefreshCutoffLiveNil sets the value for OrdersAutorefreshCutoffLive to be an explicit nil

### UnsetOrdersAutorefreshCutoffLive
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetOrdersAutorefreshCutoffLive()`

UnsetOrdersAutorefreshCutoffLive ensures that no value is present for OrdersAutorefreshCutoffLive, not even an explicit nil
### GetOrdersNumberEditableTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersNumberEditableTest() interface{}`

GetOrdersNumberEditableTest returns the OrdersNumberEditableTest field if non-nil, zero value otherwise.

### GetOrdersNumberEditableTestOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersNumberEditableTestOk() (*interface{}, bool)`

GetOrdersNumberEditableTestOk returns a tuple with the OrdersNumberEditableTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersNumberEditableTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersNumberEditableTest(v interface{})`

SetOrdersNumberEditableTest sets OrdersNumberEditableTest field to given value.

### HasOrdersNumberEditableTest

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasOrdersNumberEditableTest() bool`

HasOrdersNumberEditableTest returns a boolean if a field has been set.

### SetOrdersNumberEditableTestNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersNumberEditableTestNil(b bool)`

 SetOrdersNumberEditableTestNil sets the value for OrdersNumberEditableTest to be an explicit nil

### UnsetOrdersNumberEditableTest
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetOrdersNumberEditableTest()`

UnsetOrdersNumberEditableTest ensures that no value is present for OrdersNumberEditableTest, not even an explicit nil
### GetOrdersNumberEditableLive

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersNumberEditableLive() interface{}`

GetOrdersNumberEditableLive returns the OrdersNumberEditableLive field if non-nil, zero value otherwise.

### GetOrdersNumberEditableLiveOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersNumberEditableLiveOk() (*interface{}, bool)`

GetOrdersNumberEditableLiveOk returns a tuple with the OrdersNumberEditableLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersNumberEditableLive

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersNumberEditableLive(v interface{})`

SetOrdersNumberEditableLive sets OrdersNumberEditableLive field to given value.

### HasOrdersNumberEditableLive

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasOrdersNumberEditableLive() bool`

HasOrdersNumberEditableLive returns a boolean if a field has been set.

### SetOrdersNumberEditableLiveNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersNumberEditableLiveNil(b bool)`

 SetOrdersNumberEditableLiveNil sets the value for OrdersNumberEditableLive to be an explicit nil

### UnsetOrdersNumberEditableLive
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetOrdersNumberEditableLive()`

UnsetOrdersNumberEditableLive ensures that no value is present for OrdersNumberEditableLive, not even an explicit nil
### GetOrdersNumberAsReference

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersNumberAsReference() interface{}`

GetOrdersNumberAsReference returns the OrdersNumberAsReference field if non-nil, zero value otherwise.

### GetOrdersNumberAsReferenceOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetOrdersNumberAsReferenceOk() (*interface{}, bool)`

GetOrdersNumberAsReferenceOk returns a tuple with the OrdersNumberAsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersNumberAsReference

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersNumberAsReference(v interface{})`

SetOrdersNumberAsReference sets OrdersNumberAsReference field to given value.

### HasOrdersNumberAsReference

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasOrdersNumberAsReference() bool`

HasOrdersNumberAsReference returns a boolean if a field has been set.

### SetOrdersNumberAsReferenceNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetOrdersNumberAsReferenceNil(b bool)`

 SetOrdersNumberAsReferenceNil sets the value for OrdersNumberAsReference to be an explicit nil

### UnsetOrdersNumberAsReference
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetOrdersNumberAsReference()`

UnsetOrdersNumberAsReference ensures that no value is present for OrdersNumberAsReference, not even an explicit nil
### GetBundlesMaxItemsCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetBundlesMaxItemsCount() interface{}`

GetBundlesMaxItemsCount returns the BundlesMaxItemsCount field if non-nil, zero value otherwise.

### GetBundlesMaxItemsCountOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetBundlesMaxItemsCountOk() (*interface{}, bool)`

GetBundlesMaxItemsCountOk returns a tuple with the BundlesMaxItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundlesMaxItemsCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetBundlesMaxItemsCount(v interface{})`

SetBundlesMaxItemsCount sets BundlesMaxItemsCount field to given value.

### HasBundlesMaxItemsCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasBundlesMaxItemsCount() bool`

HasBundlesMaxItemsCount returns a boolean if a field has been set.

### SetBundlesMaxItemsCountNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetBundlesMaxItemsCountNil(b bool)`

 SetBundlesMaxItemsCountNil sets the value for BundlesMaxItemsCount to be an explicit nil

### UnsetBundlesMaxItemsCount
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetBundlesMaxItemsCount()`

UnsetBundlesMaxItemsCount ensures that no value is present for BundlesMaxItemsCount, not even an explicit nil
### GetCouponsMinCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCouponsMinCodeLength() interface{}`

GetCouponsMinCodeLength returns the CouponsMinCodeLength field if non-nil, zero value otherwise.

### GetCouponsMinCodeLengthOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCouponsMinCodeLengthOk() (*interface{}, bool)`

GetCouponsMinCodeLengthOk returns a tuple with the CouponsMinCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponsMinCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCouponsMinCodeLength(v interface{})`

SetCouponsMinCodeLength sets CouponsMinCodeLength field to given value.

### HasCouponsMinCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasCouponsMinCodeLength() bool`

HasCouponsMinCodeLength returns a boolean if a field has been set.

### SetCouponsMinCodeLengthNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCouponsMinCodeLengthNil(b bool)`

 SetCouponsMinCodeLengthNil sets the value for CouponsMinCodeLength to be an explicit nil

### UnsetCouponsMinCodeLength
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetCouponsMinCodeLength()`

UnsetCouponsMinCodeLength ensures that no value is present for CouponsMinCodeLength, not even an explicit nil
### GetCouponsMaxCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCouponsMaxCodeLength() interface{}`

GetCouponsMaxCodeLength returns the CouponsMaxCodeLength field if non-nil, zero value otherwise.

### GetCouponsMaxCodeLengthOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCouponsMaxCodeLengthOk() (*interface{}, bool)`

GetCouponsMaxCodeLengthOk returns a tuple with the CouponsMaxCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponsMaxCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCouponsMaxCodeLength(v interface{})`

SetCouponsMaxCodeLength sets CouponsMaxCodeLength field to given value.

### HasCouponsMaxCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasCouponsMaxCodeLength() bool`

HasCouponsMaxCodeLength returns a boolean if a field has been set.

### SetCouponsMaxCodeLengthNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCouponsMaxCodeLengthNil(b bool)`

 SetCouponsMaxCodeLengthNil sets the value for CouponsMaxCodeLength to be an explicit nil

### UnsetCouponsMaxCodeLength
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetCouponsMaxCodeLength()`

UnsetCouponsMaxCodeLength ensures that no value is present for CouponsMaxCodeLength, not even an explicit nil
### GetGiftCardsMinCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGiftCardsMinCodeLength() interface{}`

GetGiftCardsMinCodeLength returns the GiftCardsMinCodeLength field if non-nil, zero value otherwise.

### GetGiftCardsMinCodeLengthOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGiftCardsMinCodeLengthOk() (*interface{}, bool)`

GetGiftCardsMinCodeLengthOk returns a tuple with the GiftCardsMinCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardsMinCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGiftCardsMinCodeLength(v interface{})`

SetGiftCardsMinCodeLength sets GiftCardsMinCodeLength field to given value.

### HasGiftCardsMinCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasGiftCardsMinCodeLength() bool`

HasGiftCardsMinCodeLength returns a boolean if a field has been set.

### SetGiftCardsMinCodeLengthNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGiftCardsMinCodeLengthNil(b bool)`

 SetGiftCardsMinCodeLengthNil sets the value for GiftCardsMinCodeLength to be an explicit nil

### UnsetGiftCardsMinCodeLength
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetGiftCardsMinCodeLength()`

UnsetGiftCardsMinCodeLength ensures that no value is present for GiftCardsMinCodeLength, not even an explicit nil
### GetGiftCardsMaxCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGiftCardsMaxCodeLength() interface{}`

GetGiftCardsMaxCodeLength returns the GiftCardsMaxCodeLength field if non-nil, zero value otherwise.

### GetGiftCardsMaxCodeLengthOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetGiftCardsMaxCodeLengthOk() (*interface{}, bool)`

GetGiftCardsMaxCodeLengthOk returns a tuple with the GiftCardsMaxCodeLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiftCardsMaxCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGiftCardsMaxCodeLength(v interface{})`

SetGiftCardsMaxCodeLength sets GiftCardsMaxCodeLength field to given value.

### HasGiftCardsMaxCodeLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasGiftCardsMaxCodeLength() bool`

HasGiftCardsMaxCodeLength returns a boolean if a field has been set.

### SetGiftCardsMaxCodeLengthNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetGiftCardsMaxCodeLengthNil(b bool)`

 SetGiftCardsMaxCodeLengthNil sets the value for GiftCardsMaxCodeLength to be an explicit nil

### UnsetGiftCardsMaxCodeLength
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetGiftCardsMaxCodeLength()`

UnsetGiftCardsMaxCodeLength ensures that no value is present for GiftCardsMaxCodeLength, not even an explicit nil
### GetCleanupsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCleanupsMaxConcurrentCount() interface{}`

GetCleanupsMaxConcurrentCount returns the CleanupsMaxConcurrentCount field if non-nil, zero value otherwise.

### GetCleanupsMaxConcurrentCountOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCleanupsMaxConcurrentCountOk() (*interface{}, bool)`

GetCleanupsMaxConcurrentCountOk returns a tuple with the CleanupsMaxConcurrentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanupsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCleanupsMaxConcurrentCount(v interface{})`

SetCleanupsMaxConcurrentCount sets CleanupsMaxConcurrentCount field to given value.

### HasCleanupsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasCleanupsMaxConcurrentCount() bool`

HasCleanupsMaxConcurrentCount returns a boolean if a field has been set.

### SetCleanupsMaxConcurrentCountNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCleanupsMaxConcurrentCountNil(b bool)`

 SetCleanupsMaxConcurrentCountNil sets the value for CleanupsMaxConcurrentCount to be an explicit nil

### UnsetCleanupsMaxConcurrentCount
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetCleanupsMaxConcurrentCount()`

UnsetCleanupsMaxConcurrentCount ensures that no value is present for CleanupsMaxConcurrentCount, not even an explicit nil
### GetExportsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetExportsMaxConcurrentCount() interface{}`

GetExportsMaxConcurrentCount returns the ExportsMaxConcurrentCount field if non-nil, zero value otherwise.

### GetExportsMaxConcurrentCountOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetExportsMaxConcurrentCountOk() (*interface{}, bool)`

GetExportsMaxConcurrentCountOk returns a tuple with the ExportsMaxConcurrentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetExportsMaxConcurrentCount(v interface{})`

SetExportsMaxConcurrentCount sets ExportsMaxConcurrentCount field to given value.

### HasExportsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasExportsMaxConcurrentCount() bool`

HasExportsMaxConcurrentCount returns a boolean if a field has been set.

### SetExportsMaxConcurrentCountNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetExportsMaxConcurrentCountNil(b bool)`

 SetExportsMaxConcurrentCountNil sets the value for ExportsMaxConcurrentCount to be an explicit nil

### UnsetExportsMaxConcurrentCount
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetExportsMaxConcurrentCount()`

UnsetExportsMaxConcurrentCount ensures that no value is present for ExportsMaxConcurrentCount, not even an explicit nil
### GetImportsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetImportsMaxConcurrentCount() interface{}`

GetImportsMaxConcurrentCount returns the ImportsMaxConcurrentCount field if non-nil, zero value otherwise.

### GetImportsMaxConcurrentCountOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetImportsMaxConcurrentCountOk() (*interface{}, bool)`

GetImportsMaxConcurrentCountOk returns a tuple with the ImportsMaxConcurrentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetImportsMaxConcurrentCount(v interface{})`

SetImportsMaxConcurrentCount sets ImportsMaxConcurrentCount field to given value.

### HasImportsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasImportsMaxConcurrentCount() bool`

HasImportsMaxConcurrentCount returns a boolean if a field has been set.

### SetImportsMaxConcurrentCountNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetImportsMaxConcurrentCountNil(b bool)`

 SetImportsMaxConcurrentCountNil sets the value for ImportsMaxConcurrentCount to be an explicit nil

### UnsetImportsMaxConcurrentCount
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetImportsMaxConcurrentCount()`

UnsetImportsMaxConcurrentCount ensures that no value is present for ImportsMaxConcurrentCount, not even an explicit nil
### GetImportsPurgeCache

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetImportsPurgeCache() interface{}`

GetImportsPurgeCache returns the ImportsPurgeCache field if non-nil, zero value otherwise.

### GetImportsPurgeCacheOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetImportsPurgeCacheOk() (*interface{}, bool)`

GetImportsPurgeCacheOk returns a tuple with the ImportsPurgeCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportsPurgeCache

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetImportsPurgeCache(v interface{})`

SetImportsPurgeCache sets ImportsPurgeCache field to given value.

### HasImportsPurgeCache

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasImportsPurgeCache() bool`

HasImportsPurgeCache returns a boolean if a field has been set.

### SetImportsPurgeCacheNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetImportsPurgeCacheNil(b bool)`

 SetImportsPurgeCacheNil sets the value for ImportsPurgeCache to be an explicit nil

### UnsetImportsPurgeCache
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetImportsPurgeCache()`

UnsetImportsPurgeCache ensures that no value is present for ImportsPurgeCache, not even an explicit nil
### GetImportsSkipErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetImportsSkipErrors() interface{}`

GetImportsSkipErrors returns the ImportsSkipErrors field if non-nil, zero value otherwise.

### GetImportsSkipErrorsOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetImportsSkipErrorsOk() (*interface{}, bool)`

GetImportsSkipErrorsOk returns a tuple with the ImportsSkipErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportsSkipErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetImportsSkipErrors(v interface{})`

SetImportsSkipErrors sets ImportsSkipErrors field to given value.

### HasImportsSkipErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasImportsSkipErrors() bool`

HasImportsSkipErrors returns a boolean if a field has been set.

### SetImportsSkipErrorsNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetImportsSkipErrorsNil(b bool)`

 SetImportsSkipErrorsNil sets the value for ImportsSkipErrors to be an explicit nil

### UnsetImportsSkipErrors
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetImportsSkipErrors()`

UnsetImportsSkipErrors ensures that no value is present for ImportsSkipErrors, not even an explicit nil
### GetPromotionsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetPromotionsMaxConcurrentCount() interface{}`

GetPromotionsMaxConcurrentCount returns the PromotionsMaxConcurrentCount field if non-nil, zero value otherwise.

### GetPromotionsMaxConcurrentCountOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetPromotionsMaxConcurrentCountOk() (*interface{}, bool)`

GetPromotionsMaxConcurrentCountOk returns a tuple with the PromotionsMaxConcurrentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetPromotionsMaxConcurrentCount(v interface{})`

SetPromotionsMaxConcurrentCount sets PromotionsMaxConcurrentCount field to given value.

### HasPromotionsMaxConcurrentCount

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasPromotionsMaxConcurrentCount() bool`

HasPromotionsMaxConcurrentCount returns a boolean if a field has been set.

### SetPromotionsMaxConcurrentCountNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetPromotionsMaxConcurrentCountNil(b bool)`

 SetPromotionsMaxConcurrentCountNil sets the value for PromotionsMaxConcurrentCount to be an explicit nil

### UnsetPromotionsMaxConcurrentCount
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetPromotionsMaxConcurrentCount()`

UnsetPromotionsMaxConcurrentCount ensures that no value is present for PromotionsMaxConcurrentCount, not even an explicit nil
### GetImportsTriggerWebhooks

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetImportsTriggerWebhooks() interface{}`

GetImportsTriggerWebhooks returns the ImportsTriggerWebhooks field if non-nil, zero value otherwise.

### GetImportsTriggerWebhooksOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetImportsTriggerWebhooksOk() (*interface{}, bool)`

GetImportsTriggerWebhooksOk returns a tuple with the ImportsTriggerWebhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportsTriggerWebhooks

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetImportsTriggerWebhooks(v interface{})`

SetImportsTriggerWebhooks sets ImportsTriggerWebhooks field to given value.

### HasImportsTriggerWebhooks

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasImportsTriggerWebhooks() bool`

HasImportsTriggerWebhooks returns a boolean if a field has been set.

### SetImportsTriggerWebhooksNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetImportsTriggerWebhooksNil(b bool)`

 SetImportsTriggerWebhooksNil sets the value for ImportsTriggerWebhooks to be an explicit nil

### UnsetImportsTriggerWebhooks
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetImportsTriggerWebhooks()`

UnsetImportsTriggerWebhooks ensures that no value is present for ImportsTriggerWebhooks, not even an explicit nil
### GetDiscountEnginesEnabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDiscountEnginesEnabled() interface{}`

GetDiscountEnginesEnabled returns the DiscountEnginesEnabled field if non-nil, zero value otherwise.

### GetDiscountEnginesEnabledOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDiscountEnginesEnabledOk() (*interface{}, bool)`

GetDiscountEnginesEnabledOk returns a tuple with the DiscountEnginesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountEnginesEnabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDiscountEnginesEnabled(v interface{})`

SetDiscountEnginesEnabled sets DiscountEnginesEnabled field to given value.

### HasDiscountEnginesEnabled

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasDiscountEnginesEnabled() bool`

HasDiscountEnginesEnabled returns a boolean if a field has been set.

### SetDiscountEnginesEnabledNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDiscountEnginesEnabledNil(b bool)`

 SetDiscountEnginesEnabledNil sets the value for DiscountEnginesEnabled to be an explicit nil

### UnsetDiscountEnginesEnabled
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetDiscountEnginesEnabled()`

UnsetDiscountEnginesEnabled ensures that no value is present for DiscountEnginesEnabled, not even an explicit nil
### GetDiscountEnginesErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDiscountEnginesErrors() interface{}`

GetDiscountEnginesErrors returns the DiscountEnginesErrors field if non-nil, zero value otherwise.

### GetDiscountEnginesErrorsOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetDiscountEnginesErrorsOk() (*interface{}, bool)`

GetDiscountEnginesErrorsOk returns a tuple with the DiscountEnginesErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountEnginesErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDiscountEnginesErrors(v interface{})`

SetDiscountEnginesErrors sets DiscountEnginesErrors field to given value.

### HasDiscountEnginesErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasDiscountEnginesErrors() bool`

HasDiscountEnginesErrors returns a boolean if a field has been set.

### SetDiscountEnginesErrorsNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetDiscountEnginesErrorsNil(b bool)`

 SetDiscountEnginesErrorsNil sets the value for DiscountEnginesErrors to be an explicit nil

### UnsetDiscountEnginesErrors
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetDiscountEnginesErrors()`

UnsetDiscountEnginesErrors ensures that no value is present for DiscountEnginesErrors, not even an explicit nil
### GetTagsMaxNameLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetTagsMaxNameLength() interface{}`

GetTagsMaxNameLength returns the TagsMaxNameLength field if non-nil, zero value otherwise.

### GetTagsMaxNameLengthOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetTagsMaxNameLengthOk() (*interface{}, bool)`

GetTagsMaxNameLengthOk returns a tuple with the TagsMaxNameLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsMaxNameLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetTagsMaxNameLength(v interface{})`

SetTagsMaxNameLength sets TagsMaxNameLength field to given value.

### HasTagsMaxNameLength

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasTagsMaxNameLength() bool`

HasTagsMaxNameLength returns a boolean if a field has been set.

### SetTagsMaxNameLengthNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetTagsMaxNameLengthNil(b bool)`

 SetTagsMaxNameLengthNil sets the value for TagsMaxNameLength to be an explicit nil

### UnsetTagsMaxNameLength
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetTagsMaxNameLength()`

UnsetTagsMaxNameLength ensures that no value is present for TagsMaxNameLength, not even an explicit nil
### GetTaxCalculatorsErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetTaxCalculatorsErrors() interface{}`

GetTaxCalculatorsErrors returns the TaxCalculatorsErrors field if non-nil, zero value otherwise.

### GetTaxCalculatorsErrorsOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetTaxCalculatorsErrorsOk() (*interface{}, bool)`

GetTaxCalculatorsErrorsOk returns a tuple with the TaxCalculatorsErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCalculatorsErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetTaxCalculatorsErrors(v interface{})`

SetTaxCalculatorsErrors sets TaxCalculatorsErrors field to given value.

### HasTaxCalculatorsErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasTaxCalculatorsErrors() bool`

HasTaxCalculatorsErrors returns a boolean if a field has been set.

### SetTaxCalculatorsErrorsNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetTaxCalculatorsErrorsNil(b bool)`

 SetTaxCalculatorsErrorsNil sets the value for TaxCalculatorsErrors to be an explicit nil

### UnsetTaxCalculatorsErrors
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetTaxCalculatorsErrors()`

UnsetTaxCalculatorsErrors ensures that no value is present for TaxCalculatorsErrors, not even an explicit nil
### GetExternalPromotionsErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetExternalPromotionsErrors() interface{}`

GetExternalPromotionsErrors returns the ExternalPromotionsErrors field if non-nil, zero value otherwise.

### GetExternalPromotionsErrorsOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetExternalPromotionsErrorsOk() (*interface{}, bool)`

GetExternalPromotionsErrorsOk returns a tuple with the ExternalPromotionsErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPromotionsErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetExternalPromotionsErrors(v interface{})`

SetExternalPromotionsErrors sets ExternalPromotionsErrors field to given value.

### HasExternalPromotionsErrors

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasExternalPromotionsErrors() bool`

HasExternalPromotionsErrors returns a boolean if a field has been set.

### SetExternalPromotionsErrorsNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetExternalPromotionsErrorsNil(b bool)`

 SetExternalPromotionsErrorsNil sets the value for ExternalPromotionsErrors to be an explicit nil

### UnsetExternalPromotionsErrors
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetExternalPromotionsErrors()`

UnsetExternalPromotionsErrors ensures that no value is present for ExternalPromotionsErrors, not even an explicit nil
### GetCreatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETOrganizationOrganizationId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


