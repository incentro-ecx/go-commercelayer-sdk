# NotificationCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** | The resource&#39;s type | 
**Attributes** | [**POSTNotifications201ResponseDataAttributes**](POSTNotifications201ResponseDataAttributes.md) |  | 
**Relationships** | Pointer to [**NotificationCreateDataRelationships**](NotificationCreateDataRelationships.md) |  | [optional] 

## Methods

### NewNotificationCreateData

`func NewNotificationCreateData(type_ interface{}, attributes POSTNotifications201ResponseDataAttributes, ) *NotificationCreateData`

NewNotificationCreateData instantiates a new NotificationCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationCreateDataWithDefaults

`func NewNotificationCreateDataWithDefaults() *NotificationCreateData`

NewNotificationCreateDataWithDefaults instantiates a new NotificationCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationCreateData) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationCreateData) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationCreateData) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *NotificationCreateData) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NotificationCreateData) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAttributes

`func (o *NotificationCreateData) GetAttributes() POSTNotifications201ResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NotificationCreateData) GetAttributesOk() (*POSTNotifications201ResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NotificationCreateData) SetAttributes(v POSTNotifications201ResponseDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *NotificationCreateData) GetRelationships() NotificationCreateDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *NotificationCreateData) GetRelationshipsOk() (*NotificationCreateDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *NotificationCreateData) SetRelationships(v NotificationCreateDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *NotificationCreateData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


