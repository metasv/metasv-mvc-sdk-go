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

// TxRaw Raw hex formatted Tx returned.
type TxRaw struct {
	// hex formatted raw tx.
	Hex *string `json:"hex,omitempty"`
}

// NewTxRaw instantiates a new TxRaw object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxRaw() *TxRaw {
	this := TxRaw{}
	return &this
}

// NewTxRawWithDefaults instantiates a new TxRaw object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxRawWithDefaults() *TxRaw {
	this := TxRaw{}
	return &this
}

// GetHex returns the Hex field value if set, zero value otherwise.
func (o *TxRaw) GetHex() string {
	if o == nil || o.Hex == nil {
		var ret string
		return ret
	}
	return *o.Hex
}

// GetHexOk returns a tuple with the Hex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TxRaw) GetHexOk() (*string, bool) {
	if o == nil || o.Hex == nil {
		return nil, false
	}
	return o.Hex, true
}

// HasHex returns a boolean if a field has been set.
func (o *TxRaw) HasHex() bool {
	if o != nil && o.Hex != nil {
		return true
	}

	return false
}

// SetHex gets a reference to the given string and assigns it to the Hex field.
func (o *TxRaw) SetHex(v string) {
	o.Hex = &v
}

func (o TxRaw) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hex != nil {
		toSerialize["hex"] = o.Hex
	}
	return json.Marshal(toSerialize)
}

type NullableTxRaw struct {
	value *TxRaw
	isSet bool
}

func (v NullableTxRaw) Get() *TxRaw {
	return v.value
}

func (v *NullableTxRaw) Set(val *TxRaw) {
	v.value = val
	v.isSet = true
}

func (v NullableTxRaw) IsSet() bool {
	return v.isSet
}

func (v *NullableTxRaw) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxRaw(val *TxRaw) *NullableTxRaw {
	return &NullableTxRaw{value: val, isSet: true}
}

func (v NullableTxRaw) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxRaw) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
