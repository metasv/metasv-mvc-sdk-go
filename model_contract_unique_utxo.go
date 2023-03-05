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

// ContractUniqueUtxo Contract unique Utxo indexed by codeHash and genesis.
type ContractUniqueUtxo struct {
	// Txid for this utxo.
	Txid *string `json:"txid,omitempty"`
	// Output index for the Utxo.
	TxIndex *int32 `json:"txIndex,omitempty"`
	// Codehash of this utxo.
	CodeHash *string `json:"codeHash,omitempty"`
	// Genesis of this utxo.
	Genesis *string `json:"genesis,omitempty"`
	// SensibleId of the token
	SensibleId *string `json:"sensibleId,omitempty"`
	// The height of this utxo, -1 for unconfirmed utxo.
	Height *int64 `json:"height,omitempty"`
	// The hex encoded customData
	CustomData *string `json:"customData,omitempty"`
	// Mvc value of the utxo(Irrelavant to token value)
	Satoshi *int64 `json:"satoshi,omitempty"`
	// Mvc value of the utxo(In string format)
	SatoshiString *string `json:"satoshiString,omitempty"`
}

// NewContractUniqueUtxo instantiates a new ContractUniqueUtxo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractUniqueUtxo() *ContractUniqueUtxo {
	this := ContractUniqueUtxo{}
	return &this
}

// NewContractUniqueUtxoWithDefaults instantiates a new ContractUniqueUtxo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractUniqueUtxoWithDefaults() *ContractUniqueUtxo {
	this := ContractUniqueUtxo{}
	return &this
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *ContractUniqueUtxo) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUniqueUtxo) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *ContractUniqueUtxo) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *ContractUniqueUtxo) SetTxid(v string) {
	o.Txid = &v
}

// GetTxIndex returns the TxIndex field value if set, zero value otherwise.
func (o *ContractUniqueUtxo) GetTxIndex() int32 {
	if o == nil || o.TxIndex == nil {
		var ret int32
		return ret
	}
	return *o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUniqueUtxo) GetTxIndexOk() (*int32, bool) {
	if o == nil || o.TxIndex == nil {
		return nil, false
	}
	return o.TxIndex, true
}

// HasTxIndex returns a boolean if a field has been set.
func (o *ContractUniqueUtxo) HasTxIndex() bool {
	if o != nil && o.TxIndex != nil {
		return true
	}

	return false
}

// SetTxIndex gets a reference to the given int32 and assigns it to the TxIndex field.
func (o *ContractUniqueUtxo) SetTxIndex(v int32) {
	o.TxIndex = &v
}

// GetCodeHash returns the CodeHash field value if set, zero value otherwise.
func (o *ContractUniqueUtxo) GetCodeHash() string {
	if o == nil || o.CodeHash == nil {
		var ret string
		return ret
	}
	return *o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUniqueUtxo) GetCodeHashOk() (*string, bool) {
	if o == nil || o.CodeHash == nil {
		return nil, false
	}
	return o.CodeHash, true
}

// HasCodeHash returns a boolean if a field has been set.
func (o *ContractUniqueUtxo) HasCodeHash() bool {
	if o != nil && o.CodeHash != nil {
		return true
	}

	return false
}

// SetCodeHash gets a reference to the given string and assigns it to the CodeHash field.
func (o *ContractUniqueUtxo) SetCodeHash(v string) {
	o.CodeHash = &v
}

// GetGenesis returns the Genesis field value if set, zero value otherwise.
func (o *ContractUniqueUtxo) GetGenesis() string {
	if o == nil || o.Genesis == nil {
		var ret string
		return ret
	}
	return *o.Genesis
}

// GetGenesisOk returns a tuple with the Genesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUniqueUtxo) GetGenesisOk() (*string, bool) {
	if o == nil || o.Genesis == nil {
		return nil, false
	}
	return o.Genesis, true
}

// HasGenesis returns a boolean if a field has been set.
func (o *ContractUniqueUtxo) HasGenesis() bool {
	if o != nil && o.Genesis != nil {
		return true
	}

	return false
}

// SetGenesis gets a reference to the given string and assigns it to the Genesis field.
func (o *ContractUniqueUtxo) SetGenesis(v string) {
	o.Genesis = &v
}

// GetSensibleId returns the SensibleId field value if set, zero value otherwise.
func (o *ContractUniqueUtxo) GetSensibleId() string {
	if o == nil || o.SensibleId == nil {
		var ret string
		return ret
	}
	return *o.SensibleId
}

// GetSensibleIdOk returns a tuple with the SensibleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUniqueUtxo) GetSensibleIdOk() (*string, bool) {
	if o == nil || o.SensibleId == nil {
		return nil, false
	}
	return o.SensibleId, true
}

// HasSensibleId returns a boolean if a field has been set.
func (o *ContractUniqueUtxo) HasSensibleId() bool {
	if o != nil && o.SensibleId != nil {
		return true
	}

	return false
}

// SetSensibleId gets a reference to the given string and assigns it to the SensibleId field.
func (o *ContractUniqueUtxo) SetSensibleId(v string) {
	o.SensibleId = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ContractUniqueUtxo) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUniqueUtxo) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ContractUniqueUtxo) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *ContractUniqueUtxo) SetHeight(v int64) {
	o.Height = &v
}

// GetCustomData returns the CustomData field value if set, zero value otherwise.
func (o *ContractUniqueUtxo) GetCustomData() string {
	if o == nil || o.CustomData == nil {
		var ret string
		return ret
	}
	return *o.CustomData
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUniqueUtxo) GetCustomDataOk() (*string, bool) {
	if o == nil || o.CustomData == nil {
		return nil, false
	}
	return o.CustomData, true
}

// HasCustomData returns a boolean if a field has been set.
func (o *ContractUniqueUtxo) HasCustomData() bool {
	if o != nil && o.CustomData != nil {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given string and assigns it to the CustomData field.
func (o *ContractUniqueUtxo) SetCustomData(v string) {
	o.CustomData = &v
}

// GetSatoshi returns the Satoshi field value if set, zero value otherwise.
func (o *ContractUniqueUtxo) GetSatoshi() int64 {
	if o == nil || o.Satoshi == nil {
		var ret int64
		return ret
	}
	return *o.Satoshi
}

// GetSatoshiOk returns a tuple with the Satoshi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUniqueUtxo) GetSatoshiOk() (*int64, bool) {
	if o == nil || o.Satoshi == nil {
		return nil, false
	}
	return o.Satoshi, true
}

// HasSatoshi returns a boolean if a field has been set.
func (o *ContractUniqueUtxo) HasSatoshi() bool {
	if o != nil && o.Satoshi != nil {
		return true
	}

	return false
}

// SetSatoshi gets a reference to the given int64 and assigns it to the Satoshi field.
func (o *ContractUniqueUtxo) SetSatoshi(v int64) {
	o.Satoshi = &v
}

// GetSatoshiString returns the SatoshiString field value if set, zero value otherwise.
func (o *ContractUniqueUtxo) GetSatoshiString() string {
	if o == nil || o.SatoshiString == nil {
		var ret string
		return ret
	}
	return *o.SatoshiString
}

// GetSatoshiStringOk returns a tuple with the SatoshiString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractUniqueUtxo) GetSatoshiStringOk() (*string, bool) {
	if o == nil || o.SatoshiString == nil {
		return nil, false
	}
	return o.SatoshiString, true
}

// HasSatoshiString returns a boolean if a field has been set.
func (o *ContractUniqueUtxo) HasSatoshiString() bool {
	if o != nil && o.SatoshiString != nil {
		return true
	}

	return false
}

// SetSatoshiString gets a reference to the given string and assigns it to the SatoshiString field.
func (o *ContractUniqueUtxo) SetSatoshiString(v string) {
	o.SatoshiString = &v
}

func (o ContractUniqueUtxo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	if o.TxIndex != nil {
		toSerialize["txIndex"] = o.TxIndex
	}
	if o.CodeHash != nil {
		toSerialize["codeHash"] = o.CodeHash
	}
	if o.Genesis != nil {
		toSerialize["genesis"] = o.Genesis
	}
	if o.SensibleId != nil {
		toSerialize["sensibleId"] = o.SensibleId
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.CustomData != nil {
		toSerialize["customData"] = o.CustomData
	}
	if o.Satoshi != nil {
		toSerialize["satoshi"] = o.Satoshi
	}
	if o.SatoshiString != nil {
		toSerialize["satoshiString"] = o.SatoshiString
	}
	return json.Marshal(toSerialize)
}

type NullableContractUniqueUtxo struct {
	value *ContractUniqueUtxo
	isSet bool
}

func (v NullableContractUniqueUtxo) Get() *ContractUniqueUtxo {
	return v.value
}

func (v *NullableContractUniqueUtxo) Set(val *ContractUniqueUtxo) {
	v.value = val
	v.isSet = true
}

func (v NullableContractUniqueUtxo) IsSet() bool {
	return v.isSet
}

func (v *NullableContractUniqueUtxo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractUniqueUtxo(val *ContractUniqueUtxo) *NullableContractUniqueUtxo {
	return &NullableContractUniqueUtxo{value: val, isSet: true}
}

func (v NullableContractUniqueUtxo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractUniqueUtxo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
