/*
Wasp API

REST API for the Wasp node

API version: 0.4.0-alpha.8-14-gbb23763d3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
)

// checks if the ContractInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractInfoResponse{}

// ContractInfoResponse struct for ContractInfoResponse
type ContractInfoResponse struct {
	// The description of the contract.
	Description string `json:"description"`
	// The id (HName as Hex)) of the contract.
	HName string `json:"hName"`
	// The name of the contract.
	Name string `json:"name"`
	// The hash of the contract.
	ProgramHash []int32 `json:"programHash"`
}

// NewContractInfoResponse instantiates a new ContractInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractInfoResponse(description string, hName string, name string, programHash []int32) *ContractInfoResponse {
	this := ContractInfoResponse{}
	this.Description = description
	this.HName = hName
	this.Name = name
	this.ProgramHash = programHash
	return &this
}

// NewContractInfoResponseWithDefaults instantiates a new ContractInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractInfoResponseWithDefaults() *ContractInfoResponse {
	this := ContractInfoResponse{}
	return &this
}

// GetDescription returns the Description field value
func (o *ContractInfoResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ContractInfoResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ContractInfoResponse) SetDescription(v string) {
	o.Description = v
}

// GetHName returns the HName field value
func (o *ContractInfoResponse) GetHName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HName
}

// GetHNameOk returns a tuple with the HName field value
// and a boolean to check if the value has been set.
func (o *ContractInfoResponse) GetHNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HName, true
}

// SetHName sets field value
func (o *ContractInfoResponse) SetHName(v string) {
	o.HName = v
}

// GetName returns the Name field value
func (o *ContractInfoResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContractInfoResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContractInfoResponse) SetName(v string) {
	o.Name = v
}

// GetProgramHash returns the ProgramHash field value
func (o *ContractInfoResponse) GetProgramHash() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ProgramHash
}

// GetProgramHashOk returns a tuple with the ProgramHash field value
// and a boolean to check if the value has been set.
func (o *ContractInfoResponse) GetProgramHashOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProgramHash, true
}

// SetProgramHash sets field value
func (o *ContractInfoResponse) SetProgramHash(v []int32) {
	o.ProgramHash = v
}

func (o ContractInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["hName"] = o.HName
	toSerialize["name"] = o.Name
	toSerialize["programHash"] = o.ProgramHash
	return toSerialize, nil
}

type NullableContractInfoResponse struct {
	value *ContractInfoResponse
	isSet bool
}

func (v NullableContractInfoResponse) Get() *ContractInfoResponse {
	return v.value
}

func (v *NullableContractInfoResponse) Set(val *ContractInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContractInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContractInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractInfoResponse(val *ContractInfoResponse) *NullableContractInfoResponse {
	return &NullableContractInfoResponse{value: val, isSet: true}
}

func (v NullableContractInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


