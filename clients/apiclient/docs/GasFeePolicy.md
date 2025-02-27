# GasFeePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvmGasRatio** | [**Ratio32**](Ratio32.md) |  | 
**GasFeeTokenId** | **string** | The gas fee token id. Empty if base token. | 
**GasPerToken** | [**Ratio32**](Ratio32.md) |  | 
**ValidatorFeeShare** | **int32** | The validator fee share. | 

## Methods

### NewGasFeePolicy

`func NewGasFeePolicy(evmGasRatio Ratio32, gasFeeTokenId string, gasPerToken Ratio32, validatorFeeShare int32, ) *GasFeePolicy`

NewGasFeePolicy instantiates a new GasFeePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGasFeePolicyWithDefaults

`func NewGasFeePolicyWithDefaults() *GasFeePolicy`

NewGasFeePolicyWithDefaults instantiates a new GasFeePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvmGasRatio

`func (o *GasFeePolicy) GetEvmGasRatio() Ratio32`

GetEvmGasRatio returns the EvmGasRatio field if non-nil, zero value otherwise.

### GetEvmGasRatioOk

`func (o *GasFeePolicy) GetEvmGasRatioOk() (*Ratio32, bool)`

GetEvmGasRatioOk returns a tuple with the EvmGasRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvmGasRatio

`func (o *GasFeePolicy) SetEvmGasRatio(v Ratio32)`

SetEvmGasRatio sets EvmGasRatio field to given value.


### GetGasFeeTokenId

`func (o *GasFeePolicy) GetGasFeeTokenId() string`

GetGasFeeTokenId returns the GasFeeTokenId field if non-nil, zero value otherwise.

### GetGasFeeTokenIdOk

`func (o *GasFeePolicy) GetGasFeeTokenIdOk() (*string, bool)`

GetGasFeeTokenIdOk returns a tuple with the GasFeeTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFeeTokenId

`func (o *GasFeePolicy) SetGasFeeTokenId(v string)`

SetGasFeeTokenId sets GasFeeTokenId field to given value.


### GetGasPerToken

`func (o *GasFeePolicy) GetGasPerToken() Ratio32`

GetGasPerToken returns the GasPerToken field if non-nil, zero value otherwise.

### GetGasPerTokenOk

`func (o *GasFeePolicy) GetGasPerTokenOk() (*Ratio32, bool)`

GetGasPerTokenOk returns a tuple with the GasPerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPerToken

`func (o *GasFeePolicy) SetGasPerToken(v Ratio32)`

SetGasPerToken sets GasPerToken field to given value.


### GetValidatorFeeShare

`func (o *GasFeePolicy) GetValidatorFeeShare() int32`

GetValidatorFeeShare returns the ValidatorFeeShare field if non-nil, zero value otherwise.

### GetValidatorFeeShareOk

`func (o *GasFeePolicy) GetValidatorFeeShareOk() (*int32, bool)`

GetValidatorFeeShareOk returns a tuple with the ValidatorFeeShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorFeeShare

`func (o *GasFeePolicy) SetValidatorFeeShare(v int32)`

SetValidatorFeeShare sets ValidatorFeeShare field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


