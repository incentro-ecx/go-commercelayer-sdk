# GETLineItemsLineItemId200ResponseDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkuCode** | Pointer to **interface{}** | The code of the associated SKU. | [optional] 
**BundleCode** | Pointer to **interface{}** | The code of the associated bundle. | [optional] 
**Quantity** | Pointer to **interface{}** | The line item quantity. | [optional] 
**ExternalPrice** | Pointer to **interface{}** | When creating or updating a new line item, set this attribute to &#39;1&#39; if you want to inject the unit_amount_cents price from an external source. Any successive price computation will be done externally, until the attribute is reset to &#39;0&#39;. | [optional] 
**CurrencyCode** | Pointer to **interface{}** | The international 3-letter currency code as defined by the ISO 4217 standard, automatically inherited from the order&#39;s market. | [optional] 
**UnitAmountCents** | Pointer to **interface{}** | The unit amount of the line item, in cents. Can be specified only via an integration application, or when the item is missing, otherwise is automatically computed by using one of the available methods. Cannot be passed by sales channels. | [optional] 
**UnitAmountFloat** | Pointer to **interface{}** | The unit amount of the line item, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedUnitAmount** | Pointer to **interface{}** | The unit amount of the line item, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**CompareAtAmountCents** | Pointer to **interface{}** | The compared price amount, in cents. Useful to display a percentage discount. | [optional] 
**CompareAtAmountFloat** | Pointer to **interface{}** | The compared price amount, float. | [optional] 
**FormattedCompareAtAmount** | Pointer to **interface{}** | The compared price amount, formatted. | [optional] 
**OptionsAmountCents** | Pointer to **interface{}** | The options amount of the line item, in cents. Cannot be passed by sales channels. | [optional] 
**OptionsAmountFloat** | Pointer to **interface{}** | The options amount of the line item, float. | [optional] 
**FormattedOptionsAmount** | Pointer to **interface{}** | The options amount of the line item, formatted. | [optional] 
**DiscountCents** | Pointer to **interface{}** | The discount applied to the line item, in cents. When you apply a discount to an order, this is automatically calculated basing on the line item total_amount_cents value. | [optional] 
**DiscountFloat** | Pointer to **interface{}** | The discount applied to the line item, float. When you apply a discount to an order, this is automatically calculated basing on the line item total_amount_cents value. | [optional] 
**FormattedDiscount** | Pointer to **interface{}** | The discount applied to the line item, fromatted. When you apply a discount to an order, this is automatically calculated basing on the line item total_amount_cents value. | [optional] 
**TotalAmountCents** | Pointer to **interface{}** | Calculated as unit amount x quantity + options amount, in cents. | [optional] 
**TotalAmountFloat** | Pointer to **interface{}** | Calculated as unit amount x quantity + options amount, float. This can be useful to track the purchase on thrid party systems, e.g Google Analyitcs Enhanced Ecommerce. | [optional] 
**FormattedTotalAmount** | Pointer to **interface{}** | Calculated as unit amount x quantity + options amount, formatted. This can be useful to display the amount with currency in you views. | [optional] 
**TaxAmountCents** | Pointer to **interface{}** | The collected tax amount, otherwise calculated as total amount cents - discount cent * tax rate, in cents. | [optional] 
**TaxAmountFloat** | Pointer to **interface{}** | The collected tax amount, otherwise calculated as total amount cents - discount cent * tax rate, float. | [optional] 
**FormattedTaxAmount** | Pointer to **interface{}** | The collected tax amount, otherwise calculated as total amount cents - discount cent * tax rate, formatted. | [optional] 
**Name** | Pointer to **interface{}** | The name of the line item. When blank, it gets populated with the name of the associated item (if present). | [optional] 
**ImageUrl** | Pointer to **interface{}** | The image_url of the line item. When blank, it gets populated with the image_url of the associated item (if present, SKU only). | [optional] 
**DiscountBreakdown** | Pointer to **interface{}** | The discount breakdown for this line item (if calculated). | [optional] 
**TaxRate** | Pointer to **interface{}** | The tax rate for this line item (if calculated). | [optional] 
**TaxBreakdown** | Pointer to **interface{}** | The tax breakdown for this line item (if calculated). | [optional] 
**ItemType** | Pointer to **interface{}** | The type of the associated item. One of &#39;skus&#39;, &#39;bundles&#39;, &#39;gift_cards&#39;, &#39;shipments&#39;, &#39;payment_methods&#39;, &#39;adjustments&#39;, &#39;discount_engine_items&#39;, &#39;percentage_discount_promotions&#39;, &#39;free_shipping_promotions&#39;, &#39;buy_x_pay_y_promotions&#39;, &#39;free_gift_promotions&#39;, &#39;fixed_price_promotions&#39;, &#39;external_promotions&#39;, &#39;fixed_amount_promotions&#39;, or &#39;flex_promotions&#39;. | [optional] 
**Frequency** | Pointer to **interface{}** | The frequency which generates a subscription. Must be supported by existing associated subscription_model. | [optional] 
**CouponCode** | Pointer to **interface{}** | The coupon code, if any, used to trigger this promotion line item. null for other line item types or if the promotion line item wasn&#39;t triggered by a coupon. | [optional] 
**CircuitState** | Pointer to **interface{}** | The circuit breaker state, by default it is &#39;closed&#39;. It can become &#39;open&#39; once the number of consecutive failures overlaps the specified threshold, in such case no further calls to the failing callback are made. | [optional] 
**CircuitFailureCount** | Pointer to **interface{}** | The number of consecutive failures recorded by the circuit breaker associated to this resource, will be reset on first successful call to callback. | [optional] 
**CreatedAt** | Pointer to **interface{}** | Time at which the resource was created. | [optional] 
**UpdatedAt** | Pointer to **interface{}** | Time at which the resource was last updated. | [optional] 
**Reference** | Pointer to **interface{}** | A string that you can use to add any external identifier to the resource. This can be useful for integrating the resource to an external system, like an ERP, a marketing tool, a CRM, or whatever. | [optional] 
**ReferenceOrigin** | Pointer to **interface{}** | Any identifier of the third party system that defines the reference code. | [optional] 
**Metadata** | Pointer to **interface{}** | Set of key-value pairs that you can attach to the resource. This can be useful for storing additional information about the resource in a structured format. | [optional] 

## Methods

### NewGETLineItemsLineItemId200ResponseDataAttributes

`func NewGETLineItemsLineItemId200ResponseDataAttributes() *GETLineItemsLineItemId200ResponseDataAttributes`

NewGETLineItemsLineItemId200ResponseDataAttributes instantiates a new GETLineItemsLineItemId200ResponseDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETLineItemsLineItemId200ResponseDataAttributesWithDefaults

`func NewGETLineItemsLineItemId200ResponseDataAttributesWithDefaults() *GETLineItemsLineItemId200ResponseDataAttributes`

NewGETLineItemsLineItemId200ResponseDataAttributesWithDefaults instantiates a new GETLineItemsLineItemId200ResponseDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkuCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetSkuCode() interface{}`

GetSkuCode returns the SkuCode field if non-nil, zero value otherwise.

### GetSkuCodeOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetSkuCodeOk() (*interface{}, bool)`

GetSkuCodeOk returns a tuple with the SkuCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetSkuCode(v interface{})`

SetSkuCode sets SkuCode field to given value.

### HasSkuCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasSkuCode() bool`

HasSkuCode returns a boolean if a field has been set.

### SetSkuCodeNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetSkuCodeNil(b bool)`

 SetSkuCodeNil sets the value for SkuCode to be an explicit nil

### UnsetSkuCode
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetSkuCode()`

UnsetSkuCode ensures that no value is present for SkuCode, not even an explicit nil
### GetBundleCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetBundleCode() interface{}`

GetBundleCode returns the BundleCode field if non-nil, zero value otherwise.

### GetBundleCodeOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetBundleCodeOk() (*interface{}, bool)`

GetBundleCodeOk returns a tuple with the BundleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetBundleCode(v interface{})`

SetBundleCode sets BundleCode field to given value.

### HasBundleCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasBundleCode() bool`

HasBundleCode returns a boolean if a field has been set.

### SetBundleCodeNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetBundleCodeNil(b bool)`

 SetBundleCodeNil sets the value for BundleCode to be an explicit nil

### UnsetBundleCode
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetBundleCode()`

UnsetBundleCode ensures that no value is present for BundleCode, not even an explicit nil
### GetQuantity

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetQuantity() interface{}`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetQuantityOk() (*interface{}, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetQuantity(v interface{})`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetExternalPrice

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetExternalPrice() interface{}`

GetExternalPrice returns the ExternalPrice field if non-nil, zero value otherwise.

### GetExternalPriceOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetExternalPriceOk() (*interface{}, bool)`

GetExternalPriceOk returns a tuple with the ExternalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrice

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetExternalPrice(v interface{})`

SetExternalPrice sets ExternalPrice field to given value.

### HasExternalPrice

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasExternalPrice() bool`

HasExternalPrice returns a boolean if a field has been set.

### SetExternalPriceNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetExternalPriceNil(b bool)`

 SetExternalPriceNil sets the value for ExternalPrice to be an explicit nil

### UnsetExternalPrice
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetExternalPrice()`

UnsetExternalPrice ensures that no value is present for ExternalPrice, not even an explicit nil
### GetCurrencyCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCurrencyCode() interface{}`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCurrencyCodeOk() (*interface{}, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCurrencyCode(v interface{})`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetUnitAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetUnitAmountCents() interface{}`

GetUnitAmountCents returns the UnitAmountCents field if non-nil, zero value otherwise.

### GetUnitAmountCentsOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetUnitAmountCentsOk() (*interface{}, bool)`

GetUnitAmountCentsOk returns a tuple with the UnitAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetUnitAmountCents(v interface{})`

SetUnitAmountCents sets UnitAmountCents field to given value.

### HasUnitAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasUnitAmountCents() bool`

HasUnitAmountCents returns a boolean if a field has been set.

### SetUnitAmountCentsNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetUnitAmountCentsNil(b bool)`

 SetUnitAmountCentsNil sets the value for UnitAmountCents to be an explicit nil

### UnsetUnitAmountCents
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetUnitAmountCents()`

UnsetUnitAmountCents ensures that no value is present for UnitAmountCents, not even an explicit nil
### GetUnitAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetUnitAmountFloat() interface{}`

GetUnitAmountFloat returns the UnitAmountFloat field if non-nil, zero value otherwise.

### GetUnitAmountFloatOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetUnitAmountFloatOk() (*interface{}, bool)`

GetUnitAmountFloatOk returns a tuple with the UnitAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetUnitAmountFloat(v interface{})`

SetUnitAmountFloat sets UnitAmountFloat field to given value.

### HasUnitAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasUnitAmountFloat() bool`

HasUnitAmountFloat returns a boolean if a field has been set.

### SetUnitAmountFloatNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetUnitAmountFloatNil(b bool)`

 SetUnitAmountFloatNil sets the value for UnitAmountFloat to be an explicit nil

### UnsetUnitAmountFloat
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetUnitAmountFloat()`

UnsetUnitAmountFloat ensures that no value is present for UnitAmountFloat, not even an explicit nil
### GetFormattedUnitAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedUnitAmount() interface{}`

GetFormattedUnitAmount returns the FormattedUnitAmount field if non-nil, zero value otherwise.

### GetFormattedUnitAmountOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedUnitAmountOk() (*interface{}, bool)`

GetFormattedUnitAmountOk returns a tuple with the FormattedUnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedUnitAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedUnitAmount(v interface{})`

SetFormattedUnitAmount sets FormattedUnitAmount field to given value.

### HasFormattedUnitAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasFormattedUnitAmount() bool`

HasFormattedUnitAmount returns a boolean if a field has been set.

### SetFormattedUnitAmountNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedUnitAmountNil(b bool)`

 SetFormattedUnitAmountNil sets the value for FormattedUnitAmount to be an explicit nil

### UnsetFormattedUnitAmount
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetFormattedUnitAmount()`

UnsetFormattedUnitAmount ensures that no value is present for FormattedUnitAmount, not even an explicit nil
### GetCompareAtAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCompareAtAmountCents() interface{}`

GetCompareAtAmountCents returns the CompareAtAmountCents field if non-nil, zero value otherwise.

### GetCompareAtAmountCentsOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCompareAtAmountCentsOk() (*interface{}, bool)`

GetCompareAtAmountCentsOk returns a tuple with the CompareAtAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCompareAtAmountCents(v interface{})`

SetCompareAtAmountCents sets CompareAtAmountCents field to given value.

### HasCompareAtAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasCompareAtAmountCents() bool`

HasCompareAtAmountCents returns a boolean if a field has been set.

### SetCompareAtAmountCentsNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCompareAtAmountCentsNil(b bool)`

 SetCompareAtAmountCentsNil sets the value for CompareAtAmountCents to be an explicit nil

### UnsetCompareAtAmountCents
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetCompareAtAmountCents()`

UnsetCompareAtAmountCents ensures that no value is present for CompareAtAmountCents, not even an explicit nil
### GetCompareAtAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCompareAtAmountFloat() interface{}`

GetCompareAtAmountFloat returns the CompareAtAmountFloat field if non-nil, zero value otherwise.

### GetCompareAtAmountFloatOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCompareAtAmountFloatOk() (*interface{}, bool)`

GetCompareAtAmountFloatOk returns a tuple with the CompareAtAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCompareAtAmountFloat(v interface{})`

SetCompareAtAmountFloat sets CompareAtAmountFloat field to given value.

### HasCompareAtAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasCompareAtAmountFloat() bool`

HasCompareAtAmountFloat returns a boolean if a field has been set.

### SetCompareAtAmountFloatNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCompareAtAmountFloatNil(b bool)`

 SetCompareAtAmountFloatNil sets the value for CompareAtAmountFloat to be an explicit nil

### UnsetCompareAtAmountFloat
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetCompareAtAmountFloat()`

UnsetCompareAtAmountFloat ensures that no value is present for CompareAtAmountFloat, not even an explicit nil
### GetFormattedCompareAtAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedCompareAtAmount() interface{}`

GetFormattedCompareAtAmount returns the FormattedCompareAtAmount field if non-nil, zero value otherwise.

### GetFormattedCompareAtAmountOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedCompareAtAmountOk() (*interface{}, bool)`

GetFormattedCompareAtAmountOk returns a tuple with the FormattedCompareAtAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedCompareAtAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedCompareAtAmount(v interface{})`

SetFormattedCompareAtAmount sets FormattedCompareAtAmount field to given value.

### HasFormattedCompareAtAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasFormattedCompareAtAmount() bool`

HasFormattedCompareAtAmount returns a boolean if a field has been set.

### SetFormattedCompareAtAmountNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedCompareAtAmountNil(b bool)`

 SetFormattedCompareAtAmountNil sets the value for FormattedCompareAtAmount to be an explicit nil

### UnsetFormattedCompareAtAmount
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetFormattedCompareAtAmount()`

UnsetFormattedCompareAtAmount ensures that no value is present for FormattedCompareAtAmount, not even an explicit nil
### GetOptionsAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetOptionsAmountCents() interface{}`

GetOptionsAmountCents returns the OptionsAmountCents field if non-nil, zero value otherwise.

### GetOptionsAmountCentsOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetOptionsAmountCentsOk() (*interface{}, bool)`

GetOptionsAmountCentsOk returns a tuple with the OptionsAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetOptionsAmountCents(v interface{})`

SetOptionsAmountCents sets OptionsAmountCents field to given value.

### HasOptionsAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasOptionsAmountCents() bool`

HasOptionsAmountCents returns a boolean if a field has been set.

### SetOptionsAmountCentsNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetOptionsAmountCentsNil(b bool)`

 SetOptionsAmountCentsNil sets the value for OptionsAmountCents to be an explicit nil

### UnsetOptionsAmountCents
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetOptionsAmountCents()`

UnsetOptionsAmountCents ensures that no value is present for OptionsAmountCents, not even an explicit nil
### GetOptionsAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetOptionsAmountFloat() interface{}`

GetOptionsAmountFloat returns the OptionsAmountFloat field if non-nil, zero value otherwise.

### GetOptionsAmountFloatOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetOptionsAmountFloatOk() (*interface{}, bool)`

GetOptionsAmountFloatOk returns a tuple with the OptionsAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionsAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetOptionsAmountFloat(v interface{})`

SetOptionsAmountFloat sets OptionsAmountFloat field to given value.

### HasOptionsAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasOptionsAmountFloat() bool`

HasOptionsAmountFloat returns a boolean if a field has been set.

### SetOptionsAmountFloatNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetOptionsAmountFloatNil(b bool)`

 SetOptionsAmountFloatNil sets the value for OptionsAmountFloat to be an explicit nil

### UnsetOptionsAmountFloat
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetOptionsAmountFloat()`

UnsetOptionsAmountFloat ensures that no value is present for OptionsAmountFloat, not even an explicit nil
### GetFormattedOptionsAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedOptionsAmount() interface{}`

GetFormattedOptionsAmount returns the FormattedOptionsAmount field if non-nil, zero value otherwise.

### GetFormattedOptionsAmountOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedOptionsAmountOk() (*interface{}, bool)`

GetFormattedOptionsAmountOk returns a tuple with the FormattedOptionsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedOptionsAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedOptionsAmount(v interface{})`

SetFormattedOptionsAmount sets FormattedOptionsAmount field to given value.

### HasFormattedOptionsAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasFormattedOptionsAmount() bool`

HasFormattedOptionsAmount returns a boolean if a field has been set.

### SetFormattedOptionsAmountNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedOptionsAmountNil(b bool)`

 SetFormattedOptionsAmountNil sets the value for FormattedOptionsAmount to be an explicit nil

### UnsetFormattedOptionsAmount
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetFormattedOptionsAmount()`

UnsetFormattedOptionsAmount ensures that no value is present for FormattedOptionsAmount, not even an explicit nil
### GetDiscountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetDiscountCents() interface{}`

GetDiscountCents returns the DiscountCents field if non-nil, zero value otherwise.

### GetDiscountCentsOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetDiscountCentsOk() (*interface{}, bool)`

GetDiscountCentsOk returns a tuple with the DiscountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetDiscountCents(v interface{})`

SetDiscountCents sets DiscountCents field to given value.

### HasDiscountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasDiscountCents() bool`

HasDiscountCents returns a boolean if a field has been set.

### SetDiscountCentsNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetDiscountCentsNil(b bool)`

 SetDiscountCentsNil sets the value for DiscountCents to be an explicit nil

### UnsetDiscountCents
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetDiscountCents()`

UnsetDiscountCents ensures that no value is present for DiscountCents, not even an explicit nil
### GetDiscountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetDiscountFloat() interface{}`

GetDiscountFloat returns the DiscountFloat field if non-nil, zero value otherwise.

### GetDiscountFloatOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetDiscountFloatOk() (*interface{}, bool)`

GetDiscountFloatOk returns a tuple with the DiscountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetDiscountFloat(v interface{})`

SetDiscountFloat sets DiscountFloat field to given value.

### HasDiscountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasDiscountFloat() bool`

HasDiscountFloat returns a boolean if a field has been set.

### SetDiscountFloatNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetDiscountFloatNil(b bool)`

 SetDiscountFloatNil sets the value for DiscountFloat to be an explicit nil

### UnsetDiscountFloat
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetDiscountFloat()`

UnsetDiscountFloat ensures that no value is present for DiscountFloat, not even an explicit nil
### GetFormattedDiscount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedDiscount() interface{}`

GetFormattedDiscount returns the FormattedDiscount field if non-nil, zero value otherwise.

### GetFormattedDiscountOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedDiscountOk() (*interface{}, bool)`

GetFormattedDiscountOk returns a tuple with the FormattedDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedDiscount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedDiscount(v interface{})`

SetFormattedDiscount sets FormattedDiscount field to given value.

### HasFormattedDiscount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasFormattedDiscount() bool`

HasFormattedDiscount returns a boolean if a field has been set.

### SetFormattedDiscountNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedDiscountNil(b bool)`

 SetFormattedDiscountNil sets the value for FormattedDiscount to be an explicit nil

### UnsetFormattedDiscount
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetFormattedDiscount()`

UnsetFormattedDiscount ensures that no value is present for FormattedDiscount, not even an explicit nil
### GetTotalAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTotalAmountCents() interface{}`

GetTotalAmountCents returns the TotalAmountCents field if non-nil, zero value otherwise.

### GetTotalAmountCentsOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTotalAmountCentsOk() (*interface{}, bool)`

GetTotalAmountCentsOk returns a tuple with the TotalAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTotalAmountCents(v interface{})`

SetTotalAmountCents sets TotalAmountCents field to given value.

### HasTotalAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasTotalAmountCents() bool`

HasTotalAmountCents returns a boolean if a field has been set.

### SetTotalAmountCentsNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTotalAmountCentsNil(b bool)`

 SetTotalAmountCentsNil sets the value for TotalAmountCents to be an explicit nil

### UnsetTotalAmountCents
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetTotalAmountCents()`

UnsetTotalAmountCents ensures that no value is present for TotalAmountCents, not even an explicit nil
### GetTotalAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTotalAmountFloat() interface{}`

GetTotalAmountFloat returns the TotalAmountFloat field if non-nil, zero value otherwise.

### GetTotalAmountFloatOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTotalAmountFloatOk() (*interface{}, bool)`

GetTotalAmountFloatOk returns a tuple with the TotalAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTotalAmountFloat(v interface{})`

SetTotalAmountFloat sets TotalAmountFloat field to given value.

### HasTotalAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasTotalAmountFloat() bool`

HasTotalAmountFloat returns a boolean if a field has been set.

### SetTotalAmountFloatNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTotalAmountFloatNil(b bool)`

 SetTotalAmountFloatNil sets the value for TotalAmountFloat to be an explicit nil

### UnsetTotalAmountFloat
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetTotalAmountFloat()`

UnsetTotalAmountFloat ensures that no value is present for TotalAmountFloat, not even an explicit nil
### GetFormattedTotalAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedTotalAmount() interface{}`

GetFormattedTotalAmount returns the FormattedTotalAmount field if non-nil, zero value otherwise.

### GetFormattedTotalAmountOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedTotalAmountOk() (*interface{}, bool)`

GetFormattedTotalAmountOk returns a tuple with the FormattedTotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTotalAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedTotalAmount(v interface{})`

SetFormattedTotalAmount sets FormattedTotalAmount field to given value.

### HasFormattedTotalAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasFormattedTotalAmount() bool`

HasFormattedTotalAmount returns a boolean if a field has been set.

### SetFormattedTotalAmountNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedTotalAmountNil(b bool)`

 SetFormattedTotalAmountNil sets the value for FormattedTotalAmount to be an explicit nil

### UnsetFormattedTotalAmount
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetFormattedTotalAmount()`

UnsetFormattedTotalAmount ensures that no value is present for FormattedTotalAmount, not even an explicit nil
### GetTaxAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTaxAmountCents() interface{}`

GetTaxAmountCents returns the TaxAmountCents field if non-nil, zero value otherwise.

### GetTaxAmountCentsOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTaxAmountCentsOk() (*interface{}, bool)`

GetTaxAmountCentsOk returns a tuple with the TaxAmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTaxAmountCents(v interface{})`

SetTaxAmountCents sets TaxAmountCents field to given value.

### HasTaxAmountCents

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasTaxAmountCents() bool`

HasTaxAmountCents returns a boolean if a field has been set.

### SetTaxAmountCentsNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTaxAmountCentsNil(b bool)`

 SetTaxAmountCentsNil sets the value for TaxAmountCents to be an explicit nil

### UnsetTaxAmountCents
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetTaxAmountCents()`

UnsetTaxAmountCents ensures that no value is present for TaxAmountCents, not even an explicit nil
### GetTaxAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTaxAmountFloat() interface{}`

GetTaxAmountFloat returns the TaxAmountFloat field if non-nil, zero value otherwise.

### GetTaxAmountFloatOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTaxAmountFloatOk() (*interface{}, bool)`

GetTaxAmountFloatOk returns a tuple with the TaxAmountFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTaxAmountFloat(v interface{})`

SetTaxAmountFloat sets TaxAmountFloat field to given value.

### HasTaxAmountFloat

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasTaxAmountFloat() bool`

HasTaxAmountFloat returns a boolean if a field has been set.

### SetTaxAmountFloatNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTaxAmountFloatNil(b bool)`

 SetTaxAmountFloatNil sets the value for TaxAmountFloat to be an explicit nil

### UnsetTaxAmountFloat
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetTaxAmountFloat()`

UnsetTaxAmountFloat ensures that no value is present for TaxAmountFloat, not even an explicit nil
### GetFormattedTaxAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedTaxAmount() interface{}`

GetFormattedTaxAmount returns the FormattedTaxAmount field if non-nil, zero value otherwise.

### GetFormattedTaxAmountOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFormattedTaxAmountOk() (*interface{}, bool)`

GetFormattedTaxAmountOk returns a tuple with the FormattedTaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormattedTaxAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedTaxAmount(v interface{})`

SetFormattedTaxAmount sets FormattedTaxAmount field to given value.

### HasFormattedTaxAmount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasFormattedTaxAmount() bool`

HasFormattedTaxAmount returns a boolean if a field has been set.

### SetFormattedTaxAmountNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFormattedTaxAmountNil(b bool)`

 SetFormattedTaxAmountNil sets the value for FormattedTaxAmount to be an explicit nil

### UnsetFormattedTaxAmount
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetFormattedTaxAmount()`

UnsetFormattedTaxAmount ensures that no value is present for FormattedTaxAmount, not even an explicit nil
### GetName

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetName() interface{}`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetNameOk() (*interface{}, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetName(v interface{})`

SetName sets Name field to given value.

### HasName

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImageUrl

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetImageUrl() interface{}`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetImageUrlOk() (*interface{}, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetImageUrl(v interface{})`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetDiscountBreakdown

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetDiscountBreakdown() interface{}`

GetDiscountBreakdown returns the DiscountBreakdown field if non-nil, zero value otherwise.

### GetDiscountBreakdownOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetDiscountBreakdownOk() (*interface{}, bool)`

GetDiscountBreakdownOk returns a tuple with the DiscountBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountBreakdown

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetDiscountBreakdown(v interface{})`

SetDiscountBreakdown sets DiscountBreakdown field to given value.

### HasDiscountBreakdown

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasDiscountBreakdown() bool`

HasDiscountBreakdown returns a boolean if a field has been set.

### SetDiscountBreakdownNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetDiscountBreakdownNil(b bool)`

 SetDiscountBreakdownNil sets the value for DiscountBreakdown to be an explicit nil

### UnsetDiscountBreakdown
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetDiscountBreakdown()`

UnsetDiscountBreakdown ensures that no value is present for DiscountBreakdown, not even an explicit nil
### GetTaxRate

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTaxRate() interface{}`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTaxRateOk() (*interface{}, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTaxRate(v interface{})`

SetTaxRate sets TaxRate field to given value.

### HasTaxRate

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasTaxRate() bool`

HasTaxRate returns a boolean if a field has been set.

### SetTaxRateNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetTaxBreakdown

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTaxBreakdown() interface{}`

GetTaxBreakdown returns the TaxBreakdown field if non-nil, zero value otherwise.

### GetTaxBreakdownOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetTaxBreakdownOk() (*interface{}, bool)`

GetTaxBreakdownOk returns a tuple with the TaxBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBreakdown

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTaxBreakdown(v interface{})`

SetTaxBreakdown sets TaxBreakdown field to given value.

### HasTaxBreakdown

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasTaxBreakdown() bool`

HasTaxBreakdown returns a boolean if a field has been set.

### SetTaxBreakdownNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetTaxBreakdownNil(b bool)`

 SetTaxBreakdownNil sets the value for TaxBreakdown to be an explicit nil

### UnsetTaxBreakdown
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetTaxBreakdown()`

UnsetTaxBreakdown ensures that no value is present for TaxBreakdown, not even an explicit nil
### GetItemType

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetItemType() interface{}`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetItemTypeOk() (*interface{}, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetItemType(v interface{})`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetFrequency

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFrequency() interface{}`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetFrequencyOk() (*interface{}, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFrequency(v interface{})`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### SetFrequencyNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetFrequencyNil(b bool)`

 SetFrequencyNil sets the value for Frequency to be an explicit nil

### UnsetFrequency
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetFrequency()`

UnsetFrequency ensures that no value is present for Frequency, not even an explicit nil
### GetCouponCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCouponCode() interface{}`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCouponCodeOk() (*interface{}, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCouponCode(v interface{})`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.

### SetCouponCodeNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCouponCodeNil(b bool)`

 SetCouponCodeNil sets the value for CouponCode to be an explicit nil

### UnsetCouponCode
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetCouponCode()`

UnsetCouponCode ensures that no value is present for CouponCode, not even an explicit nil
### GetCircuitState

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCircuitState() interface{}`

GetCircuitState returns the CircuitState field if non-nil, zero value otherwise.

### GetCircuitStateOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCircuitStateOk() (*interface{}, bool)`

GetCircuitStateOk returns a tuple with the CircuitState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitState

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCircuitState(v interface{})`

SetCircuitState sets CircuitState field to given value.

### HasCircuitState

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasCircuitState() bool`

HasCircuitState returns a boolean if a field has been set.

### SetCircuitStateNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCircuitStateNil(b bool)`

 SetCircuitStateNil sets the value for CircuitState to be an explicit nil

### UnsetCircuitState
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetCircuitState()`

UnsetCircuitState ensures that no value is present for CircuitState, not even an explicit nil
### GetCircuitFailureCount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCircuitFailureCount() interface{}`

GetCircuitFailureCount returns the CircuitFailureCount field if non-nil, zero value otherwise.

### GetCircuitFailureCountOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCircuitFailureCountOk() (*interface{}, bool)`

GetCircuitFailureCountOk returns a tuple with the CircuitFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitFailureCount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCircuitFailureCount(v interface{})`

SetCircuitFailureCount sets CircuitFailureCount field to given value.

### HasCircuitFailureCount

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasCircuitFailureCount() bool`

HasCircuitFailureCount returns a boolean if a field has been set.

### SetCircuitFailureCountNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCircuitFailureCountNil(b bool)`

 SetCircuitFailureCountNil sets the value for CircuitFailureCount to be an explicit nil

### UnsetCircuitFailureCount
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetCircuitFailureCount()`

UnsetCircuitFailureCount ensures that no value is present for CircuitFailureCount, not even an explicit nil
### GetCreatedAt

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetUpdatedAt() interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetUpdatedAtOk() (*interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetUpdatedAt(v interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetReference

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetReference() interface{}`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetReferenceOk() (*interface{}, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetReference(v interface{})`

SetReference sets Reference field to given value.

### HasReference

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasReference() bool`

HasReference returns a boolean if a field has been set.

### SetReferenceNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetReferenceOrigin

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetReferenceOrigin() interface{}`

GetReferenceOrigin returns the ReferenceOrigin field if non-nil, zero value otherwise.

### GetReferenceOriginOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetReferenceOriginOk() (*interface{}, bool)`

GetReferenceOriginOk returns a tuple with the ReferenceOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOrigin

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetReferenceOrigin(v interface{})`

SetReferenceOrigin sets ReferenceOrigin field to given value.

### HasReferenceOrigin

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasReferenceOrigin() bool`

HasReferenceOrigin returns a boolean if a field has been set.

### SetReferenceOriginNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetReferenceOriginNil(b bool)`

 SetReferenceOriginNil sets the value for ReferenceOrigin to be an explicit nil

### UnsetReferenceOrigin
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetReferenceOrigin()`

UnsetReferenceOrigin ensures that no value is present for ReferenceOrigin, not even an explicit nil
### GetMetadata

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetMetadata() interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) GetMetadataOk() (*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetMetadata(v interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GETLineItemsLineItemId200ResponseDataAttributes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GETLineItemsLineItemId200ResponseDataAttributes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


