/*
Wasp API

REST API for the Wasp node

API version: 0.4.0-alpha.2-402-g2adcf666b
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OffLedgerRequest struct for OffLedgerRequest
type OffLedgerRequest struct {
	// The chain id
	ChainId *string `json:"chainId,omitempty"`
	// Offledger Request (Hex)
	Request *string `json:"request,omitempty"`
}

// NewOffLedgerRequest instantiates a new OffLedgerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffLedgerRequest() *OffLedgerRequest {
	this := OffLedgerRequest{}
	return &this
}

// NewOffLedgerRequestWithDefaults instantiates a new OffLedgerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOffLedgerRequestWithDefaults() *OffLedgerRequest {
	this := OffLedgerRequest{}
	return &this
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *OffLedgerRequest) GetChainId() string {
	if o == nil || isNil(o.ChainId) {
		var ret string
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OffLedgerRequest) GetChainIdOk() (*string, bool) {
	if o == nil || isNil(o.ChainId) {
    return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *OffLedgerRequest) HasChainId() bool {
	if o != nil && !isNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given string and assigns it to the ChainId field.
func (o *OffLedgerRequest) SetChainId(v string) {
	o.ChainId = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *OffLedgerRequest) GetRequest() string {
	if o == nil || isNil(o.Request) {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OffLedgerRequest) GetRequestOk() (*string, bool) {
	if o == nil || isNil(o.Request) {
    return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *OffLedgerRequest) HasRequest() bool {
	if o != nil && !isNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *OffLedgerRequest) SetRequest(v string) {
	o.Request = &v
}

func (o OffLedgerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !isNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return json.Marshal(toSerialize)
}

type NullableOffLedgerRequest struct {
	value *OffLedgerRequest
	isSet bool
}

func (v NullableOffLedgerRequest) Get() *OffLedgerRequest {
	return v.value
}

func (v *NullableOffLedgerRequest) Set(val *OffLedgerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOffLedgerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOffLedgerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffLedgerRequest(val *OffLedgerRequest) *NullableOffLedgerRequest {
	return &NullableOffLedgerRequest{value: val, isSet: true}
}

func (v NullableOffLedgerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffLedgerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


