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

// AsyncBroadcastResult Broadcast result
type AsyncBroadcastResult struct {
	// the txid of this tx
	Txid *string `json:"txid,omitempty"`
	// the state of this tx, possible values PENDING, BROADCASTED, INVALID, UNKNOWN
	State *string `json:"state,omitempty"`
	// return messages if broadcast failed
	Message *string `json:"message,omitempty"`
}

// NewAsyncBroadcastResult instantiates a new AsyncBroadcastResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsyncBroadcastResult() *AsyncBroadcastResult {
	this := AsyncBroadcastResult{}
	return &this
}

// NewAsyncBroadcastResultWithDefaults instantiates a new AsyncBroadcastResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsyncBroadcastResultWithDefaults() *AsyncBroadcastResult {
	this := AsyncBroadcastResult{}
	return &this
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *AsyncBroadcastResult) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncBroadcastResult) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *AsyncBroadcastResult) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *AsyncBroadcastResult) SetTxid(v string) {
	o.Txid = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AsyncBroadcastResult) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncBroadcastResult) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AsyncBroadcastResult) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AsyncBroadcastResult) SetState(v string) {
	o.State = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AsyncBroadcastResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsyncBroadcastResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AsyncBroadcastResult) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AsyncBroadcastResult) SetMessage(v string) {
	o.Message = &v
}

func (o AsyncBroadcastResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableAsyncBroadcastResult struct {
	value *AsyncBroadcastResult
	isSet bool
}

func (v NullableAsyncBroadcastResult) Get() *AsyncBroadcastResult {
	return v.value
}

func (v *NullableAsyncBroadcastResult) Set(val *AsyncBroadcastResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAsyncBroadcastResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAsyncBroadcastResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsyncBroadcastResult(val *AsyncBroadcastResult) *NullableAsyncBroadcastResult {
	return &NullableAsyncBroadcastResult{value: val, isSet: true}
}

func (v NullableAsyncBroadcastResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsyncBroadcastResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
