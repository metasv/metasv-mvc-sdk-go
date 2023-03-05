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

// BroadcastResult Broadcast result
type BroadcastResult struct {
	// return txid if broadcast success
	Txid *string `json:"txid,omitempty"`
	// return messages if broadcast failed
	Message *string `json:"message,omitempty"`
}

// NewBroadcastResult instantiates a new BroadcastResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastResult() *BroadcastResult {
	this := BroadcastResult{}
	return &this
}

// NewBroadcastResultWithDefaults instantiates a new BroadcastResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastResultWithDefaults() *BroadcastResult {
	this := BroadcastResult{}
	return &this
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *BroadcastResult) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastResult) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *BroadcastResult) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *BroadcastResult) SetTxid(v string) {
	o.Txid = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BroadcastResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BroadcastResult) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BroadcastResult) SetMessage(v string) {
	o.Message = &v
}

func (o BroadcastResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableBroadcastResult struct {
	value *BroadcastResult
	isSet bool
}

func (v NullableBroadcastResult) Get() *BroadcastResult {
	return v.value
}

func (v *NullableBroadcastResult) Set(val *BroadcastResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastResult(val *BroadcastResult) *NullableBroadcastResult {
	return &NullableBroadcastResult{value: val, isSet: true}
}

func (v NullableBroadcastResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
