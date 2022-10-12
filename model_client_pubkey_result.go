/*
 * MetaSV API Spec
 *
 * API definition for MetaSV provided apis
 *
 * API version: 2.2.0
 * Contact: heqiming@metasv.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metasv

import (
	"encoding/json"
)

// ClientPubkeyResult Client public key result
type ClientPubkeyResult struct {
	// return pubkey if created successfully
	Pubkey *string `json:"pubkey,omitempty"`
	// return messages if broadcast failed
	Message *string `json:"message,omitempty"`
}

// NewClientPubkeyResult instantiates a new ClientPubkeyResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientPubkeyResult() *ClientPubkeyResult {
	this := ClientPubkeyResult{}
	return &this
}

// NewClientPubkeyResultWithDefaults instantiates a new ClientPubkeyResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientPubkeyResultWithDefaults() *ClientPubkeyResult {
	this := ClientPubkeyResult{}
	return &this
}

// GetPubkey returns the Pubkey field value if set, zero value otherwise.
func (o *ClientPubkeyResult) GetPubkey() string {
	if o == nil || o.Pubkey == nil {
		var ret string
		return ret
	}
	return *o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPubkeyResult) GetPubkeyOk() (*string, bool) {
	if o == nil || o.Pubkey == nil {
		return nil, false
	}
	return o.Pubkey, true
}

// HasPubkey returns a boolean if a field has been set.
func (o *ClientPubkeyResult) HasPubkey() bool {
	if o != nil && o.Pubkey != nil {
		return true
	}

	return false
}

// SetPubkey gets a reference to the given string and assigns it to the Pubkey field.
func (o *ClientPubkeyResult) SetPubkey(v string) {
	o.Pubkey = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClientPubkeyResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientPubkeyResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClientPubkeyResult) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClientPubkeyResult) SetMessage(v string) {
	o.Message = &v
}

func (o ClientPubkeyResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pubkey != nil {
		toSerialize["pubkey"] = o.Pubkey
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableClientPubkeyResult struct {
	value *ClientPubkeyResult
	isSet bool
}

func (v NullableClientPubkeyResult) Get() *ClientPubkeyResult {
	return v.value
}

func (v *NullableClientPubkeyResult) Set(val *ClientPubkeyResult) {
	v.value = val
	v.isSet = true
}

func (v NullableClientPubkeyResult) IsSet() bool {
	return v.isSet
}

func (v *NullableClientPubkeyResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientPubkeyResult(val *ClientPubkeyResult) *NullableClientPubkeyResult {
	return &NullableClientPubkeyResult{value: val, isSet: true}
}

func (v NullableClientPubkeyResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientPubkeyResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
