# POSTPaymentMethods201ResponseDataRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Market** | Pointer to [**POSTBundles201ResponseDataRelationshipsMarket**](POSTBundles201ResponseDataRelationshipsMarket.md) |  | [optional] 
**PaymentGateway** | Pointer to [**POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway**](POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway.md) |  | [optional] 
**Store** | Pointer to [**POSTOrders201ResponseDataRelationshipsStore**](POSTOrders201ResponseDataRelationshipsStore.md) |  | [optional] 
**Orders** | Pointer to [**POSTCustomers201ResponseDataRelationshipsOrders**](POSTCustomers201ResponseDataRelationshipsOrders.md) |  | [optional] 
**Attachments** | Pointer to [**GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments**](GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments.md) |  | [optional] 
**Versions** | Pointer to [**POSTAddresses201ResponseDataRelationshipsVersions**](POSTAddresses201ResponseDataRelationshipsVersions.md) |  | [optional] 

## Methods

### NewPOSTPaymentMethods201ResponseDataRelationships

`func NewPOSTPaymentMethods201ResponseDataRelationships() *POSTPaymentMethods201ResponseDataRelationships`

NewPOSTPaymentMethods201ResponseDataRelationships instantiates a new POSTPaymentMethods201ResponseDataRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPOSTPaymentMethods201ResponseDataRelationshipsWithDefaults

`func NewPOSTPaymentMethods201ResponseDataRelationshipsWithDefaults() *POSTPaymentMethods201ResponseDataRelationships`

NewPOSTPaymentMethods201ResponseDataRelationshipsWithDefaults instantiates a new POSTPaymentMethods201ResponseDataRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarket

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetMarket() POSTBundles201ResponseDataRelationshipsMarket`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetMarketOk() (*POSTBundles201ResponseDataRelationshipsMarket, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *POSTPaymentMethods201ResponseDataRelationships) SetMarket(v POSTBundles201ResponseDataRelationshipsMarket)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *POSTPaymentMethods201ResponseDataRelationships) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetPaymentGateway

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetPaymentGateway() POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway`

GetPaymentGateway returns the PaymentGateway field if non-nil, zero value otherwise.

### GetPaymentGatewayOk

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetPaymentGatewayOk() (*POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway, bool)`

GetPaymentGatewayOk returns a tuple with the PaymentGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGateway

`func (o *POSTPaymentMethods201ResponseDataRelationships) SetPaymentGateway(v POSTAdyenPayments201ResponseDataRelationshipsPaymentGateway)`

SetPaymentGateway sets PaymentGateway field to given value.

### HasPaymentGateway

`func (o *POSTPaymentMethods201ResponseDataRelationships) HasPaymentGateway() bool`

HasPaymentGateway returns a boolean if a field has been set.

### GetStore

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetStore() POSTOrders201ResponseDataRelationshipsStore`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetStoreOk() (*POSTOrders201ResponseDataRelationshipsStore, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *POSTPaymentMethods201ResponseDataRelationships) SetStore(v POSTOrders201ResponseDataRelationshipsStore)`

SetStore sets Store field to given value.

### HasStore

`func (o *POSTPaymentMethods201ResponseDataRelationships) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetOrders

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetOrders() POSTCustomers201ResponseDataRelationshipsOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetOrdersOk() (*POSTCustomers201ResponseDataRelationshipsOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *POSTPaymentMethods201ResponseDataRelationships) SetOrders(v POSTCustomers201ResponseDataRelationshipsOrders)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *POSTPaymentMethods201ResponseDataRelationships) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetAttachments

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetAttachments() GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetAttachmentsOk() (*GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *POSTPaymentMethods201ResponseDataRelationships) SetAttachments(v GETAuthorizationsAuthorizationId200ResponseDataRelationshipsAttachments)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *POSTPaymentMethods201ResponseDataRelationships) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetVersions

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetVersions() POSTAddresses201ResponseDataRelationshipsVersions`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *POSTPaymentMethods201ResponseDataRelationships) GetVersionsOk() (*POSTAddresses201ResponseDataRelationshipsVersions, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *POSTPaymentMethods201ResponseDataRelationships) SetVersions(v POSTAddresses201ResponseDataRelationshipsVersions)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *POSTPaymentMethods201ResponseDataRelationships) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


