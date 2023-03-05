/*
 * MetaSV for MVC API Spec
 *
 * API definition for MetaSV provided apis
 *
 * API version: 3.0.3
 * Contact: heqiming@metasv.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metasv

import (
	"encoding/json"
)

// TxInput Parsed inputs from raw tx. Use output api to get value and spent info.
type TxInput struct {
	// Input index of the tx.
	Index *int32 `json:"index,omitempty"`
	// The outpoint utxo txid that this input spent
	UtxoTxid *string `json:"utxoTxid,omitempty"`
	// The outpoint utxo index that this input spent
	UtxoIndex *int32 `json:"utxoIndex,omitempty"`
	// Parsed address from pubkey(P2PKH input only).
	Address *string `json:"address,omitempty"`
	// satoshi value of this input.
	Value *int64 `json:"value,omitempty"`
	// The hex of the input script.
	UnlockScript *string `json:"unlockScript,omitempty"`
}

// NewTxInput instantiates a new TxInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxInput() *TxInput {
	this := TxInput{}
	return &this
}

// NewTxInputWithDefaults instantiates a new TxInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxInputWithDefaults() *TxInput {
	this := TxInput{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *TxInput) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *TxInput) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *TxInput) SetIndex(v int32) {
	o.Index = &v
}

// GetUtxoTxid returns the UtxoTxid field value if set, zero value otherwise.
func (o *TxInput) GetUtxoTxid() string {
	if o == nil || o.UtxoTxid == nil {
		var ret string
		return ret
	}
	return *o.UtxoTxid
}

// GetUtxoTxidOk returns a tuple with the UtxoTxid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetUtxoTxidOk() (*string, bool) {
	if o == nil || o.UtxoTxid == nil {
		return nil, false
	}
	return o.UtxoTxid, true
}

// HasUtxoTxid returns a boolean if a field has been set.
func (o *TxInput) HasUtxoTxid() bool {
	if o != nil && o.UtxoTxid != nil {
		return true
	}

	return false
}

// SetUtxoTxid gets a reference to the given string and assigns it to the UtxoTxid field.
func (o *TxInput) SetUtxoTxid(v string) {
	o.UtxoTxid = &v
}

// GetUtxoIndex returns the UtxoIndex field value if set, zero value otherwise.
func (o *TxInput) GetUtxoIndex() int32 {
	if o == nil || o.UtxoIndex == nil {
		var ret int32
		return ret
	}
	return *o.UtxoIndex
}

// GetUtxoIndexOk returns a tuple with the UtxoIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetUtxoIndexOk() (*int32, bool) {
	if o == nil || o.UtxoIndex == nil {
		return nil, false
	}
	return o.UtxoIndex, true
}

// HasUtxoIndex returns a boolean if a field has been set.
func (o *TxInput) HasUtxoIndex() bool {
	if o != nil && o.UtxoIndex != nil {
		return true
	}

	return false
}

// SetUtxoIndex gets a reference to the given int32 and assigns it to the UtxoIndex field.
func (o *TxInput) SetUtxoIndex(v int32) {
	o.UtxoIndex = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TxInput) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TxInput) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TxInput) SetAddress(v string) {
	o.Address = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TxInput) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TxInput) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *TxInput) SetValue(v int64) {
	o.Value = &v
}

// GetUnlockScript returns the UnlockScript field value if set, zero value otherwise.
func (o *TxInput) GetUnlockScript() string {
	if o == nil || o.UnlockScript == nil {
		var ret string
		return ret
	}
	return *o.UnlockScript
}

// GetUnlockScriptOk returns a tuple with the UnlockScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxInput) GetUnlockScriptOk() (*string, bool) {
	if o == nil || o.UnlockScript == nil {
		return nil, false
	}
	return o.UnlockScript, true
}

// HasUnlockScript returns a boolean if a field has been set.
func (o *TxInput) HasUnlockScript() bool {
	if o != nil && o.UnlockScript != nil {
		return true
	}

	return false
}

// SetUnlockScript gets a reference to the given string and assigns it to the UnlockScript field.
func (o *TxInput) SetUnlockScript(v string) {
	o.UnlockScript = &v
}

func (o TxInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.UtxoTxid != nil {
		toSerialize["utxoTxid"] = o.UtxoTxid
	}
	if o.UtxoIndex != nil {
		toSerialize["utxoIndex"] = o.UtxoIndex
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.UnlockScript != nil {
		toSerialize["unlockScript"] = o.UnlockScript
	}
	return json.Marshal(toSerialize)
}

type NullableTxInput struct {
	value *TxInput
	isSet bool
}

func (v NullableTxInput) Get() *TxInput {
	return v.value
}

func (v *NullableTxInput) Set(val *TxInput) {
	v.value = val
	v.isSet = true
}

func (v NullableTxInput) IsSet() bool {
	return v.isSet
}

func (v *NullableTxInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxInput(val *TxInput) *NullableTxInput {
	return &NullableTxInput{value: val, isSet: true}
}

func (v NullableTxInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
