package eventgrid

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// EndpointType enumerates the values for endpoint type.
type EndpointType string

const (
	// WebHook ...
	WebHook EndpointType = "WebHook"
)

// EventSubscriptionProvisioningState enumerates the values for event subscription provisioning state.
type EventSubscriptionProvisioningState string

const (
	// Canceled ...
	Canceled EventSubscriptionProvisioningState = "Canceled"
	// Creating ...
	Creating EventSubscriptionProvisioningState = "Creating"
	// Deleting ...
	Deleting EventSubscriptionProvisioningState = "Deleting"
	// Failed ...
	Failed EventSubscriptionProvisioningState = "Failed"
	// Succeeded ...
	Succeeded EventSubscriptionProvisioningState = "Succeeded"
	// Updating ...
	Updating EventSubscriptionProvisioningState = "Updating"
)

// OperationOrigin enumerates the values for operation origin.
type OperationOrigin string

const (
	// System ...
	System OperationOrigin = "System"
	// User ...
	User OperationOrigin = "User"
	// UserAndSystem ...
	UserAndSystem OperationOrigin = "UserAndSystem"
)

// ResourceRegionType enumerates the values for resource region type.
type ResourceRegionType string

const (
	// GlobalResource ...
	GlobalResource ResourceRegionType = "GlobalResource"
	// RegionalResource ...
	RegionalResource ResourceRegionType = "RegionalResource"
)

// TopicProvisioningState enumerates the values for topic provisioning state.
type TopicProvisioningState string

const (
	// TopicProvisioningStateCanceled ...
	TopicProvisioningStateCanceled TopicProvisioningState = "Canceled"
	// TopicProvisioningStateCreating ...
	TopicProvisioningStateCreating TopicProvisioningState = "Creating"
	// TopicProvisioningStateDeleting ...
	TopicProvisioningStateDeleting TopicProvisioningState = "Deleting"
	// TopicProvisioningStateFailed ...
	TopicProvisioningStateFailed TopicProvisioningState = "Failed"
	// TopicProvisioningStateSucceeded ...
	TopicProvisioningStateSucceeded TopicProvisioningState = "Succeeded"
	// TopicProvisioningStateUpdating ...
	TopicProvisioningStateUpdating TopicProvisioningState = "Updating"
)

// TopicTypeProvisioningState enumerates the values for topic type provisioning state.
type TopicTypeProvisioningState string

const (
	// TopicTypeProvisioningStateCanceled ...
	TopicTypeProvisioningStateCanceled TopicTypeProvisioningState = "Canceled"
	// TopicTypeProvisioningStateCreating ...
	TopicTypeProvisioningStateCreating TopicTypeProvisioningState = "Creating"
	// TopicTypeProvisioningStateDeleting ...
	TopicTypeProvisioningStateDeleting TopicTypeProvisioningState = "Deleting"
	// TopicTypeProvisioningStateFailed ...
	TopicTypeProvisioningStateFailed TopicTypeProvisioningState = "Failed"
	// TopicTypeProvisioningStateSucceeded ...
	TopicTypeProvisioningStateSucceeded TopicTypeProvisioningState = "Succeeded"
	// TopicTypeProvisioningStateUpdating ...
	TopicTypeProvisioningStateUpdating TopicTypeProvisioningState = "Updating"
)

// EventSubscription event Subscription
type EventSubscription struct {
	autorest.Response `json:"-"`
	// ID - Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - Type of the resource
	Type *string `json:"type,omitempty"`
	// EventSubscriptionProperties - Properties of the event subscription
	*EventSubscriptionProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for EventSubscription struct.
func (es *EventSubscription) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties EventSubscriptionProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		es.EventSubscriptionProperties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		es.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		es.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		es.Type = &typeVar
	}

	return nil
}

// EventSubscriptionDestination information about the destination for an event subscription
type EventSubscriptionDestination struct {
	// EndpointType - Type of the endpoint for the event subscription destination. Possible values include: 'WebHook'
	EndpointType EndpointType `json:"endpointType,omitempty"`
	// EventSubscriptionDestinationProperties - Properties of the event subscription destination
	*EventSubscriptionDestinationProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for EventSubscriptionDestination struct.
func (esd *EventSubscriptionDestination) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["endpointType"]
	if v != nil {
		var endpointType EndpointType
		err = json.Unmarshal(*m["endpointType"], &endpointType)
		if err != nil {
			return err
		}
		esd.EndpointType = endpointType
	}

	v = m["properties"]
	if v != nil {
		var properties EventSubscriptionDestinationProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		esd.EventSubscriptionDestinationProperties = &properties
	}

	return nil
}

// EventSubscriptionDestinationProperties properties of the event subscription destination
type EventSubscriptionDestinationProperties struct {
	// EndpointURL - The URL that represents the endpoint of the destination of an event subscription.
	EndpointURL *string `json:"endpointUrl,omitempty"`
	// EndpointBaseURL - The base URL that represents the endpoint of the destination of an event subscription.
	EndpointBaseURL *string `json:"endpointBaseUrl,omitempty"`
}

// EventSubscriptionFilter filter for the Event Subscription
type EventSubscriptionFilter struct {
	// SubjectBeginsWith - An optional string to filter events for an event subscription based on a resource path prefix.
	// The format of this depends on the publisher of the events.
	// Wildcard characters are not supported in this path.
	SubjectBeginsWith *string `json:"subjectBeginsWith,omitempty"`
	// SubjectEndsWith - An optional string to filter events for an event subscription based on a resource path suffix.
	// Wildcard characters are not supported in this path.
	SubjectEndsWith *string `json:"subjectEndsWith,omitempty"`
	// IncludedEventTypes - A list of applicable event types that need to be part of the event subscription.
	// If it is desired to subscribe to all event types, the string "all" needs to be specified as an element in this list.
	IncludedEventTypes *[]string `json:"includedEventTypes,omitempty"`
	// IsSubjectCaseSensitive - Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter
	// should be compared in a case sensitive manner.
	IsSubjectCaseSensitive *bool `json:"isSubjectCaseSensitive,omitempty"`
}

// EventSubscriptionFullURL full endpoint url of an event subscription
type EventSubscriptionFullURL struct {
	autorest.Response `json:"-"`
	// EndpointURL - The URL that represents the endpoint of the destination of an event subscription.
	EndpointURL *string `json:"endpointUrl,omitempty"`
}

// EventSubscriptionProperties properties of the Event Subscription
type EventSubscriptionProperties struct {
	// Topic - Name of the topic of the event subscription.
	Topic *string `json:"topic,omitempty"`
	// ProvisioningState - Provisioning state of the event subscription. Possible values include: 'Creating', 'Updating', 'Deleting', 'Succeeded', 'Canceled', 'Failed'
	ProvisioningState EventSubscriptionProvisioningState `json:"provisioningState,omitempty"`
	// Destination - Information about the destination where events have to be delivered for the event subscription.
	Destination *EventSubscriptionDestination `json:"destination,omitempty"`
	// Filter - Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter `json:"filter,omitempty"`
	// Labels - List of user defined labels.
	Labels *[]string `json:"labels,omitempty"`
}

// EventSubscriptionsCreateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type EventSubscriptionsCreateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future EventSubscriptionsCreateFuture) Result(client EventSubscriptionsClient) (es EventSubscription, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return es, autorest.NewError("eventgrid.EventSubscriptionsCreateFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		es, err = client.CreateResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	es, err = client.CreateResponder(resp)
	return
}

// EventSubscriptionsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type EventSubscriptionsDeleteFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future EventSubscriptionsDeleteFuture) Result(client EventSubscriptionsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return ar, autorest.NewError("eventgrid.EventSubscriptionsDeleteFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ar, err = client.DeleteResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	ar, err = client.DeleteResponder(resp)
	return
}

// EventSubscriptionsListResult result of the List EventSubscriptions operation
type EventSubscriptionsListResult struct {
	autorest.Response `json:"-"`
	// Value - A collection of EventSubscriptions
	Value *[]EventSubscription `json:"value,omitempty"`
}

// EventSubscriptionsUpdateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type EventSubscriptionsUpdateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future EventSubscriptionsUpdateFuture) Result(client EventSubscriptionsClient) (es EventSubscription, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return es, autorest.NewError("eventgrid.EventSubscriptionsUpdateFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		es, err = client.UpdateResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	es, err = client.UpdateResponder(resp)
	return
}

// EventSubscriptionUpdateParameters properties of the Event Subscription update
type EventSubscriptionUpdateParameters struct {
	// Destination - Information about the destination where events have to be delivered for the event subscription.
	Destination *EventSubscriptionDestination `json:"destination,omitempty"`
	// Filter - Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter `json:"filter,omitempty"`
	// Labels - List of user defined labels.
	Labels *[]string `json:"labels,omitempty"`
}

// EventType event Type for a subject under a topic
type EventType struct {
	// ID - Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - Type of the resource
	Type *string `json:"type,omitempty"`
	// EventTypeProperties - Properties of the event type.
	*EventTypeProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for EventType struct.
func (et *EventType) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties EventTypeProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		et.EventTypeProperties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		et.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		et.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		et.Type = &typeVar
	}

	return nil
}

// EventTypeProperties properties of the event type
type EventTypeProperties struct {
	// DisplayName - Display name of the event type.
	DisplayName *string `json:"displayName,omitempty"`
	// Description - Description of the event type.
	Description *string `json:"description,omitempty"`
	// SchemaURL - Url of the schema for this event type.
	SchemaURL *string `json:"schemaUrl,omitempty"`
}

// EventTypesListResult result of the List Event Types operation
type EventTypesListResult struct {
	autorest.Response `json:"-"`
	// Value - A collection of event types
	Value *[]EventType `json:"value,omitempty"`
}

// Operation represents an operation returned by the GetOperations request
type Operation struct {
	// Name - Name of the operation
	Name *string `json:"name,omitempty"`
	// Display - Display name of the operation
	Display *OperationInfo `json:"display,omitempty"`
	// Origin - Origin of the operation. Possible values include: 'User', 'System', 'UserAndSystem'
	Origin OperationOrigin `json:"origin,omitempty"`
	// Properties - Properties of the operation
	Properties *map[string]interface{} `json:"properties,omitempty"`
}

// OperationInfo information about an operation
type OperationInfo struct {
	// Provider - Name of the provider
	Provider *string `json:"provider,omitempty"`
	// Resource - Name of the resource type
	Resource *string `json:"resource,omitempty"`
	// Operation - Name of the operation
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation
	Description *string `json:"description,omitempty"`
}

// OperationsListResult result of the List Operations operation
type OperationsListResult struct {
	autorest.Response `json:"-"`
	// Value - A collection of operations
	Value *[]Operation `json:"value,omitempty"`
}

// Resource definition of a Resource
type Resource struct {
	// ID - Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - Type of the resource
	Type *string `json:"type,omitempty"`
}

// Topic eventGrid Topic
type Topic struct {
	autorest.Response `json:"-"`
	// ID - Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - Type of the resource
	Type *string `json:"type,omitempty"`
	// Location - Location of the resource
	Location *string `json:"location,omitempty"`
	// Tags - Tags of the resource
	Tags *map[string]*string `json:"tags,omitempty"`
	// TopicProperties - Properties of the topic
	*TopicProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for Topic struct.
func (t *Topic) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties TopicProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		t.TopicProperties = &properties
	}

	v = m["location"]
	if v != nil {
		var location string
		err = json.Unmarshal(*m["location"], &location)
		if err != nil {
			return err
		}
		t.Location = &location
	}

	v = m["tags"]
	if v != nil {
		var tags map[string]*string
		err = json.Unmarshal(*m["tags"], &tags)
		if err != nil {
			return err
		}
		t.Tags = &tags
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		t.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		t.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		t.Type = &typeVar
	}

	return nil
}

// TopicProperties properties of the Topic
type TopicProperties struct {
	// ProvisioningState - Provisioning state of the topic. Possible values include: 'TopicProvisioningStateCreating', 'TopicProvisioningStateUpdating', 'TopicProvisioningStateDeleting', 'TopicProvisioningStateSucceeded', 'TopicProvisioningStateCanceled', 'TopicProvisioningStateFailed'
	ProvisioningState TopicProvisioningState `json:"provisioningState,omitempty"`
	// Endpoint - Endpoint for the topic.
	Endpoint *string `json:"endpoint,omitempty"`
}

// TopicRegenerateKeyRequest topic regenerate share access key key request
type TopicRegenerateKeyRequest struct {
	// KeyName - Key name to regenerate key1 or key2
	KeyName *string `json:"keyName,omitempty"`
}

// TopicsCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type TopicsCreateOrUpdateFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future TopicsCreateOrUpdateFuture) Result(client TopicsClient) (t Topic, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return t, autorest.NewError("eventgrid.TopicsCreateOrUpdateFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		t, err = client.CreateOrUpdateResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	t, err = client.CreateOrUpdateResponder(resp)
	return
}

// TopicsDeleteFuture an abstraction for monitoring and retrieving the results of a long-running operation.
type TopicsDeleteFuture struct {
	azure.Future
	req *http.Request
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future TopicsDeleteFuture) Result(client TopicsClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		return
	}
	if !done {
		return ar, autorest.NewError("eventgrid.TopicsDeleteFuture", "Result", "asynchronous operation has not completed")
	}
	if future.PollingMethod() == azure.PollingLocation {
		ar, err = client.DeleteResponder(future.Response())
		return
	}
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, autorest.ChangeToGet(future.req),
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	ar, err = client.DeleteResponder(resp)
	return
}

// TopicSharedAccessKeys shared access keys of the Topic
type TopicSharedAccessKeys struct {
	autorest.Response `json:"-"`
	// Key1 - Shared access key1 for the topic.
	Key1 *string `json:"key1,omitempty"`
	// Key2 - Shared access key2 for the topic.
	Key2 *string `json:"key2,omitempty"`
}

// TopicsListResult result of the List Topics operation
type TopicsListResult struct {
	autorest.Response `json:"-"`
	// Value - A collection of Topics
	Value *[]Topic `json:"value,omitempty"`
}

// TopicTypeInfo properties of a topic type info.
type TopicTypeInfo struct {
	autorest.Response `json:"-"`
	// ID - Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - Type of the resource
	Type *string `json:"type,omitempty"`
	// TopicTypeProperties - Properties of the topic type info
	*TopicTypeProperties `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for TopicTypeInfo struct.
func (tti *TopicTypeInfo) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["properties"]
	if v != nil {
		var properties TopicTypeProperties
		err = json.Unmarshal(*m["properties"], &properties)
		if err != nil {
			return err
		}
		tti.TopicTypeProperties = &properties
	}

	v = m["id"]
	if v != nil {
		var ID string
		err = json.Unmarshal(*m["id"], &ID)
		if err != nil {
			return err
		}
		tti.ID = &ID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		tti.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar string
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		tti.Type = &typeVar
	}

	return nil
}

// TopicTypeProperties properties of a topic type.
type TopicTypeProperties struct {
	// Provider - Namespace of the provider of the topic type.
	Provider *string `json:"provider,omitempty"`
	// DisplayName - Display Name for the topic type.
	DisplayName *string `json:"displayName,omitempty"`
	// Description - Description of the topic type.
	Description *string `json:"description,omitempty"`
	// ResourceRegionType - Region type of the resource. Possible values include: 'RegionalResource', 'GlobalResource'
	ResourceRegionType ResourceRegionType `json:"resourceRegionType,omitempty"`
	// ProvisioningState - Provisioning state of the topic type. Possible values include: 'TopicTypeProvisioningStateCreating', 'TopicTypeProvisioningStateUpdating', 'TopicTypeProvisioningStateDeleting', 'TopicTypeProvisioningStateSucceeded', 'TopicTypeProvisioningStateCanceled', 'TopicTypeProvisioningStateFailed'
	ProvisioningState TopicTypeProvisioningState `json:"provisioningState,omitempty"`
}

// TopicTypesListResult result of the List Topic Types operation
type TopicTypesListResult struct {
	autorest.Response `json:"-"`
	// Value - A collection of topic types
	Value *[]TopicTypeInfo `json:"value,omitempty"`
}

// TrackedResource definition of a Tracked Resource
type TrackedResource struct {
	// ID - Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - Type of the resource
	Type *string `json:"type,omitempty"`
	// Location - Location of the resource
	Location *string `json:"location,omitempty"`
	// Tags - Tags of the resource
	Tags *map[string]*string `json:"tags,omitempty"`
}
