/*
Wasp API

REST API for the Wasp node

API version: 0.4.0-alpha.2-402-g2adcf666b
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// TransactionIDMetricItem struct for TransactionIDMetricItem
type TransactionIDMetricItem struct {
	LastMessage *Transaction `json:"lastMessage,omitempty"`
	Messages *int32 `json:"messages,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewTransactionIDMetricItem instantiates a new TransactionIDMetricItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionIDMetricItem() *TransactionIDMetricItem {
	this := TransactionIDMetricItem{}
	return &this
}

// NewTransactionIDMetricItemWithDefaults instantiates a new TransactionIDMetricItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionIDMetricItemWithDefaults() *TransactionIDMetricItem {
	this := TransactionIDMetricItem{}
	return &this
}

// GetLastMessage returns the LastMessage field value if set, zero value otherwise.
func (o *TransactionIDMetricItem) GetLastMessage() Transaction {
	if o == nil || isNil(o.LastMessage) {
		var ret Transaction
		return ret
	}
	return *o.LastMessage
}

// GetLastMessageOk returns a tuple with the LastMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionIDMetricItem) GetLastMessageOk() (*Transaction, bool) {
	if o == nil || isNil(o.LastMessage) {
    return nil, false
	}
	return o.LastMessage, true
}

// HasLastMessage returns a boolean if a field has been set.
func (o *TransactionIDMetricItem) HasLastMessage() bool {
	if o != nil && !isNil(o.LastMessage) {
		return true
	}

	return false
}

// SetLastMessage gets a reference to the given Transaction and assigns it to the LastMessage field.
func (o *TransactionIDMetricItem) SetLastMessage(v Transaction) {
	o.LastMessage = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *TransactionIDMetricItem) GetMessages() int32 {
	if o == nil || isNil(o.Messages) {
		var ret int32
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionIDMetricItem) GetMessagesOk() (*int32, bool) {
	if o == nil || isNil(o.Messages) {
    return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *TransactionIDMetricItem) HasMessages() bool {
	if o != nil && !isNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given int32 and assigns it to the Messages field.
func (o *TransactionIDMetricItem) SetMessages(v int32) {
	o.Messages = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TransactionIDMetricItem) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionIDMetricItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
    return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TransactionIDMetricItem) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *TransactionIDMetricItem) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o TransactionIDMetricItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LastMessage) {
		toSerialize["lastMessage"] = o.LastMessage
	}
	if !isNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !isNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionIDMetricItem struct {
	value *TransactionIDMetricItem
	isSet bool
}

func (v NullableTransactionIDMetricItem) Get() *TransactionIDMetricItem {
	return v.value
}

func (v *NullableTransactionIDMetricItem) Set(val *TransactionIDMetricItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionIDMetricItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionIDMetricItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionIDMetricItem(val *TransactionIDMetricItem) *NullableTransactionIDMetricItem {
	return &NullableTransactionIDMetricItem{value: val, isSet: true}
}

func (v NullableTransactionIDMetricItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionIDMetricItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


