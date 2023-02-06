/*
Wasp API

REST API for the Wasp node

API version: v0.4.0-alpha.8-60-g8b0f75e6d
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
	"time"
)

// checks if the TransactionMetricItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionMetricItem{}

// TransactionMetricItem struct for TransactionMetricItem
type TransactionMetricItem struct {
	LastMessage Transaction `json:"lastMessage"`
	Messages uint32 `json:"messages"`
	Timestamp time.Time `json:"timestamp"`
}

// NewTransactionMetricItem instantiates a new TransactionMetricItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionMetricItem(lastMessage Transaction, messages uint32, timestamp time.Time) *TransactionMetricItem {
	this := TransactionMetricItem{}
	this.LastMessage = lastMessage
	this.Messages = messages
	this.Timestamp = timestamp
	return &this
}

// NewTransactionMetricItemWithDefaults instantiates a new TransactionMetricItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionMetricItemWithDefaults() *TransactionMetricItem {
	this := TransactionMetricItem{}
	return &this
}

// GetLastMessage returns the LastMessage field value
func (o *TransactionMetricItem) GetLastMessage() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.LastMessage
}

// GetLastMessageOk returns a tuple with the LastMessage field value
// and a boolean to check if the value has been set.
func (o *TransactionMetricItem) GetLastMessageOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastMessage, true
}

// SetLastMessage sets field value
func (o *TransactionMetricItem) SetLastMessage(v Transaction) {
	o.LastMessage = v
}

// GetMessages returns the Messages field value
func (o *TransactionMetricItem) GetMessages() uint32 {
	if o == nil {
		var ret uint32
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *TransactionMetricItem) GetMessagesOk() (*uint32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value
func (o *TransactionMetricItem) SetMessages(v uint32) {
	o.Messages = v
}

// GetTimestamp returns the Timestamp field value
func (o *TransactionMetricItem) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *TransactionMetricItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *TransactionMetricItem) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o TransactionMetricItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionMetricItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lastMessage"] = o.LastMessage
	toSerialize["messages"] = o.Messages
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableTransactionMetricItem struct {
	value *TransactionMetricItem
	isSet bool
}

func (v NullableTransactionMetricItem) Get() *TransactionMetricItem {
	return v.value
}

func (v *NullableTransactionMetricItem) Set(val *TransactionMetricItem) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionMetricItem) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionMetricItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionMetricItem(val *TransactionMetricItem) *NullableTransactionMetricItem {
	return &NullableTransactionMetricItem{value: val, isSet: true}
}

func (v NullableTransactionMetricItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionMetricItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


