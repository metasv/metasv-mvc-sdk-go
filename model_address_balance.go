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

// AddressBalance The current balance for a particular address
type AddressBalance struct {
	// current address
	Address *string `json:"address,omitempty"`
	// confirmed balance
	Confirmed *int64 `json:"confirmed,omitempty"`
	// unconfirmed balance
	Unconfirmed *int64 `json:"unconfirmed,omitempty"`
}

// NewAddressBalance instantiates a new AddressBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressBalance() *AddressBalance {
	this := AddressBalance{}
	return &this
}

// NewAddressBalanceWithDefaults instantiates a new AddressBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressBalanceWithDefaults() *AddressBalance {
	this := AddressBalance{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *AddressBalance) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *AddressBalance) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *AddressBalance) SetAddress(v string) {
	o.Address = &v
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *AddressBalance) GetConfirmed() int64 {
	if o == nil || o.Confirmed == nil {
		var ret int64
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetConfirmedOk() (*int64, bool) {
	if o == nil || o.Confirmed == nil {
		return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *AddressBalance) HasConfirmed() bool {
	if o != nil && o.Confirmed != nil {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given int64 and assigns it to the Confirmed field.
func (o *AddressBalance) SetConfirmed(v int64) {
	o.Confirmed = &v
}

// GetUnconfirmed returns the Unconfirmed field value if set, zero value otherwise.
func (o *AddressBalance) GetUnconfirmed() int64 {
	if o == nil || o.Unconfirmed == nil {
		var ret int64
		return ret
	}
	return *o.Unconfirmed
}

// GetUnconfirmedOk returns a tuple with the Unconfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressBalance) GetUnconfirmedOk() (*int64, bool) {
	if o == nil || o.Unconfirmed == nil {
		return nil, false
	}
	return o.Unconfirmed, true
}

// HasUnconfirmed returns a boolean if a field has been set.
func (o *AddressBalance) HasUnconfirmed() bool {
	if o != nil && o.Unconfirmed != nil {
		return true
	}

	return false
}

// SetUnconfirmed gets a reference to the given int64 and assigns it to the Unconfirmed field.
func (o *AddressBalance) SetUnconfirmed(v int64) {
	o.Unconfirmed = &v
}

func (o AddressBalance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Confirmed != nil {
		toSerialize["confirmed"] = o.Confirmed
	}
	if o.Unconfirmed != nil {
		toSerialize["unconfirmed"] = o.Unconfirmed
	}
	return json.Marshal(toSerialize)
}

type NullableAddressBalance struct {
	value *AddressBalance
	isSet bool
}

func (v NullableAddressBalance) Get() *AddressBalance {
	return v.value
}

func (v *NullableAddressBalance) Set(val *AddressBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressBalance(val *AddressBalance) *NullableAddressBalance {
	return &NullableAddressBalance{value: val, isSet: true}
}

func (v NullableAddressBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
