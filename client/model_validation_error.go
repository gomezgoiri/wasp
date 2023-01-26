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

// ValidationError struct for ValidationError
type ValidationError struct {
	Error *string `json:"Error,omitempty"`
	MissingPermission *string `json:"MissingPermission,omitempty"`
}

// NewValidationError instantiates a new ValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationError() *ValidationError {
	this := ValidationError{}
	return &this
}

// NewValidationErrorWithDefaults instantiates a new ValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorWithDefaults() *ValidationError {
	this := ValidationError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ValidationError) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ValidationError) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ValidationError) SetError(v string) {
	o.Error = &v
}

// GetMissingPermission returns the MissingPermission field value if set, zero value otherwise.
func (o *ValidationError) GetMissingPermission() string {
	if o == nil || isNil(o.MissingPermission) {
		var ret string
		return ret
	}
	return *o.MissingPermission
}

// GetMissingPermissionOk returns a tuple with the MissingPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationError) GetMissingPermissionOk() (*string, bool) {
	if o == nil || isNil(o.MissingPermission) {
    return nil, false
	}
	return o.MissingPermission, true
}

// HasMissingPermission returns a boolean if a field has been set.
func (o *ValidationError) HasMissingPermission() bool {
	if o != nil && !isNil(o.MissingPermission) {
		return true
	}

	return false
}

// SetMissingPermission gets a reference to the given string and assigns it to the MissingPermission field.
func (o *ValidationError) SetMissingPermission(v string) {
	o.MissingPermission = &v
}

func (o ValidationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	if !isNil(o.MissingPermission) {
		toSerialize["MissingPermission"] = o.MissingPermission
	}
	return json.Marshal(toSerialize)
}

type NullableValidationError struct {
	value *ValidationError
	isSet bool
}

func (v NullableValidationError) Get() *ValidationError {
	return v.value
}

func (v *NullableValidationError) Set(val *ValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationError(val *ValidationError) *NullableValidationError {
	return &NullableValidationError{value: val, isSet: true}
}

func (v NullableValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


