/*
 * MetaSV for MVC API Spec
 *
 * API definition for MetaSV provided apis
 *
 * API version: 3.0.2
 * Contact: heqiming@metasv.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metasv

import (
	"encoding/json"
)

// UtxoPickRequest Request object to batch pick utxo
type UtxoPickRequest struct {
	// The address list to pick utxo from
	Addresses *[]string `json:"addresses,omitempty"`
	// The total amount you want to pick
	Amount *int64 `json:"amount,omitempty"`
}

// NewUtxoPickRequest instantiates a new UtxoPickRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtxoPickRequest() *UtxoPickRequest {
	this := UtxoPickRequest{}
	return &this
}

// NewUtxoPickRequestWithDefaults instantiates a new UtxoPickRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtxoPickRequestWithDefaults() *UtxoPickRequest {
	this := UtxoPickRequest{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *UtxoPickRequest) GetAddresses() []string {
	if o == nil || o.Addresses == nil {
		var ret []string
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtxoPickRequest) GetAddressesOk() (*[]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *UtxoPickRequest) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *UtxoPickRequest) SetAddresses(v []string) {
	o.Addresses = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *UtxoPickRequest) GetAmount() int64 {
	if o == nil || o.Amount == nil {
		var ret int64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtxoPickRequest) GetAmountOk() (*int64, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *UtxoPickRequest) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given int64 and assigns it to the Amount field.
func (o *UtxoPickRequest) SetAmount(v int64) {
	o.Amount = &v
}

func (o UtxoPickRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	return json.Marshal(toSerialize)
}

type NullableUtxoPickRequest struct {
	value *UtxoPickRequest
	isSet bool
}

func (v NullableUtxoPickRequest) Get() *UtxoPickRequest {
	return v.value
}

func (v *NullableUtxoPickRequest) Set(val *UtxoPickRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUtxoPickRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUtxoPickRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtxoPickRequest(val *UtxoPickRequest) *NullableUtxoPickRequest {
	return &NullableUtxoPickRequest{value: val, isSet: true}
}

func (v NullableUtxoPickRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtxoPickRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
