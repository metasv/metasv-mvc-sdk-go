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

// XpubUtxo Utxo belongs to the specified xpub
type XpubUtxo struct {
	// xpub of the utxo
	Xpub *string `json:"xpub,omitempty"`
	// Address string of this utxo
	Address *string `json:"address,omitempty"`
	// Address type, 0 for receive address, 1 for change address. path is {{addressType}}/{{addressIndex}}
	AddressType *int32 `json:"addressType,omitempty"`
	// Address index. Address path in the xpub is {{addressType}}/{{addressIndex}}
	AddressIndex *int32 `json:"addressIndex,omitempty"`
	// Txid for this utxo.
	Txid *string `json:"txid,omitempty"`
	// Output index for the Utxo.
	TxIndex *int32 `json:"txIndex,omitempty"`
	// Satoshi value of the utxo.
	Value *int64 `json:"value,omitempty"`
	// The height of this utxo, -1 for unconfirmed utxo.
	Height *int64 `json:"height,omitempty"`
	// The paging flag of utxo
	Flag *int64 `json:"flag,omitempty"`
}

// NewXpubUtxo instantiates a new XpubUtxo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXpubUtxo() *XpubUtxo {
	this := XpubUtxo{}
	return &this
}

// NewXpubUtxoWithDefaults instantiates a new XpubUtxo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXpubUtxoWithDefaults() *XpubUtxo {
	this := XpubUtxo{}
	return &this
}

// GetXpub returns the Xpub field value if set, zero value otherwise.
func (o *XpubUtxo) GetXpub() string {
	if o == nil || o.Xpub == nil {
		var ret string
		return ret
	}
	return *o.Xpub
}

// GetXpubOk returns a tuple with the Xpub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpubUtxo) GetXpubOk() (*string, bool) {
	if o == nil || o.Xpub == nil {
		return nil, false
	}
	return o.Xpub, true
}

// HasXpub returns a boolean if a field has been set.
func (o *XpubUtxo) HasXpub() bool {
	if o != nil && o.Xpub != nil {
		return true
	}

	return false
}

// SetXpub gets a reference to the given string and assigns it to the Xpub field.
func (o *XpubUtxo) SetXpub(v string) {
	o.Xpub = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *XpubUtxo) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpubUtxo) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *XpubUtxo) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *XpubUtxo) SetAddress(v string) {
	o.Address = &v
}

// GetAddressType returns the AddressType field value if set, zero value otherwise.
func (o *XpubUtxo) GetAddressType() int32 {
	if o == nil || o.AddressType == nil {
		var ret int32
		return ret
	}
	return *o.AddressType
}

// GetAddressTypeOk returns a tuple with the AddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpubUtxo) GetAddressTypeOk() (*int32, bool) {
	if o == nil || o.AddressType == nil {
		return nil, false
	}
	return o.AddressType, true
}

// HasAddressType returns a boolean if a field has been set.
func (o *XpubUtxo) HasAddressType() bool {
	if o != nil && o.AddressType != nil {
		return true
	}

	return false
}

// SetAddressType gets a reference to the given int32 and assigns it to the AddressType field.
func (o *XpubUtxo) SetAddressType(v int32) {
	o.AddressType = &v
}

// GetAddressIndex returns the AddressIndex field value if set, zero value otherwise.
func (o *XpubUtxo) GetAddressIndex() int32 {
	if o == nil || o.AddressIndex == nil {
		var ret int32
		return ret
	}
	return *o.AddressIndex
}

// GetAddressIndexOk returns a tuple with the AddressIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpubUtxo) GetAddressIndexOk() (*int32, bool) {
	if o == nil || o.AddressIndex == nil {
		return nil, false
	}
	return o.AddressIndex, true
}

// HasAddressIndex returns a boolean if a field has been set.
func (o *XpubUtxo) HasAddressIndex() bool {
	if o != nil && o.AddressIndex != nil {
		return true
	}

	return false
}

// SetAddressIndex gets a reference to the given int32 and assigns it to the AddressIndex field.
func (o *XpubUtxo) SetAddressIndex(v int32) {
	o.AddressIndex = &v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *XpubUtxo) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpubUtxo) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *XpubUtxo) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *XpubUtxo) SetTxid(v string) {
	o.Txid = &v
}

// GetTxIndex returns the TxIndex field value if set, zero value otherwise.
func (o *XpubUtxo) GetTxIndex() int32 {
	if o == nil || o.TxIndex == nil {
		var ret int32
		return ret
	}
	return *o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpubUtxo) GetTxIndexOk() (*int32, bool) {
	if o == nil || o.TxIndex == nil {
		return nil, false
	}
	return o.TxIndex, true
}

// HasTxIndex returns a boolean if a field has been set.
func (o *XpubUtxo) HasTxIndex() bool {
	if o != nil && o.TxIndex != nil {
		return true
	}

	return false
}

// SetTxIndex gets a reference to the given int32 and assigns it to the TxIndex field.
func (o *XpubUtxo) SetTxIndex(v int32) {
	o.TxIndex = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *XpubUtxo) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpubUtxo) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *XpubUtxo) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *XpubUtxo) SetValue(v int64) {
	o.Value = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *XpubUtxo) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpubUtxo) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *XpubUtxo) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *XpubUtxo) SetHeight(v int64) {
	o.Height = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *XpubUtxo) GetFlag() int64 {
	if o == nil || o.Flag == nil {
		var ret int64
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XpubUtxo) GetFlagOk() (*int64, bool) {
	if o == nil || o.Flag == nil {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *XpubUtxo) HasFlag() bool {
	if o != nil && o.Flag != nil {
		return true
	}

	return false
}

// SetFlag gets a reference to the given int64 and assigns it to the Flag field.
func (o *XpubUtxo) SetFlag(v int64) {
	o.Flag = &v
}

func (o XpubUtxo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Xpub != nil {
		toSerialize["xpub"] = o.Xpub
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.AddressType != nil {
		toSerialize["addressType"] = o.AddressType
	}
	if o.AddressIndex != nil {
		toSerialize["addressIndex"] = o.AddressIndex
	}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.TxIndex != nil {
		toSerialize["txIndex"] = o.TxIndex
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Flag != nil {
		toSerialize["flag"] = o.Flag
	}
	return json.Marshal(toSerialize)
}

type NullableXpubUtxo struct {
	value *XpubUtxo
	isSet bool
}

func (v NullableXpubUtxo) Get() *XpubUtxo {
	return v.value
}

func (v *NullableXpubUtxo) Set(val *XpubUtxo) {
	v.value = val
	v.isSet = true
}

func (v NullableXpubUtxo) IsSet() bool {
	return v.isSet
}

func (v *NullableXpubUtxo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXpubUtxo(val *XpubUtxo) *NullableXpubUtxo {
	return &NullableXpubUtxo{value: val, isSet: true}
}

func (v NullableXpubUtxo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXpubUtxo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
