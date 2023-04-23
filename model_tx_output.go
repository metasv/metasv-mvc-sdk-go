/*
 * MicrovisionChain API Document
 *
 * API definition for MicrovisionChain provided apis
 *
 * API version: 3.0.8
 * Contact: heqiming@mvcapi.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mvcapi

import (
	"encoding/json"
)

// TxOutput Parsed outputs from raw tx.
type TxOutput struct {
	// Output index of the tx.
	Index *int32 `json:"index,omitempty"`
	// Bitcoin Value in this output.
	Value *int64 `json:"value,omitempty"`
	// Parsed address from output
	Address *string `json:"address,omitempty"`
	// Hex formatted lockScript
	LockScript *string `json:"lockScript,omitempty"`
}

// NewTxOutput instantiates a new TxOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxOutput() *TxOutput {
	this := TxOutput{}
	return &this
}

// NewTxOutputWithDefaults instantiates a new TxOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxOutputWithDefaults() *TxOutput {
	this := TxOutput{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *TxOutput) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutput) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *TxOutput) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *TxOutput) SetIndex(v int32) {
	o.Index = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TxOutput) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutput) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TxOutput) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *TxOutput) SetValue(v int64) {
	o.Value = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TxOutput) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutput) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TxOutput) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TxOutput) SetAddress(v string) {
	o.Address = &v
}

// GetLockScript returns the LockScript field value if set, zero value otherwise.
func (o *TxOutput) GetLockScript() string {
	if o == nil || o.LockScript == nil {
		var ret string
		return ret
	}
	return *o.LockScript
}

// GetLockScriptOk returns a tuple with the LockScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxOutput) GetLockScriptOk() (*string, bool) {
	if o == nil || o.LockScript == nil {
		return nil, false
	}
	return o.LockScript, true
}

// HasLockScript returns a boolean if a field has been set.
func (o *TxOutput) HasLockScript() bool {
	if o != nil && o.LockScript != nil {
		return true
	}

	return false
}

// SetLockScript gets a reference to the given string and assigns it to the LockScript field.
func (o *TxOutput) SetLockScript(v string) {
	o.LockScript = &v
}

func (o TxOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.LockScript != nil {
		toSerialize["lockScript"] = o.LockScript
	}
	return json.Marshal(toSerialize)
}

type NullableTxOutput struct {
	value *TxOutput
	isSet bool
}

func (v NullableTxOutput) Get() *TxOutput {
	return v.value
}

func (v *NullableTxOutput) Set(val *TxOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableTxOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableTxOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxOutput(val *TxOutput) *NullableTxOutput {
	return &NullableTxOutput{value: val, isSet: true}
}

func (v NullableTxOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
