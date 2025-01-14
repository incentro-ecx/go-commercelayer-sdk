# NotificationUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Id** | **interface{}** | The resource&#39;s id | 
**Attributes** | [**PATCHNotificationsNotificationId200ResponseDataAttributes**](PATCHNotificationsNotificationId200ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**NotificationUpdateDataRelationships**](NotificationUpdateDataRelationships.md) |  | [optional] 

## Methods

### NewNotificationUpdateData

`func NewNotificationUpdateData(type_ interface{}, id interface{}, attributes PATCHNotificationsNotificationId200ResponseDataAttributes, ) *NotificationUpdateData`

NewNotificationUpdateData instantiates a new NotificationUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationUpdateDataWithDefaults

`func NewNotificationUpdateDataWithDefaults() *NotificationUpdateData`

NewNotificationUpdateDataWithDefaults instantiates a new NotificationUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationUpdateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationUpdateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationUpdateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *NotificationUpdateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NotificationUpdateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetId

`func (o *NotificationUpdateData) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationUpdateData) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationUpdateData) SetId(v interface{})`

SetId sets Id field to given value.


### SetIdNil

`func (o *NotificationUpdateData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *NotificationUpdateData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetAttributes

`func (o *NotificationUpdateData) GetAttributes() PATCHNotificationsNotificationId200ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotificationUpdateData) GetAttributesOk() (*PATCHNotificationsNotificationId200ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotificationUpdateData) SetAttributes(v PATCHNotificationsNotificationId200ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *NotificationUpdateData) GetRelationships() NotificationUpdateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *NotificationUpdateData) GetRelationshipsOk() (*NotificationUpdateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *NotificationUpdateData) SetRelationships(v NotificationUpdateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *NotificationUpdateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


