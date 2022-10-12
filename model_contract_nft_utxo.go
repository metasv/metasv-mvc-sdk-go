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

// ContractNftUtxo Contract non fungible token Utxo belongs to the specified address
type ContractNftUtxo struct {
	// Address string of this utxo
	Address *string `json:"address,omitempty"`
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
	Height *int32 `json:"height,omitempty"`
	// The metanet tx describing the nft.
	MetaTxid *string `json:"metaTxid,omitempty"`
	// Symbol of the token.
	MetaOutputIndex *int32 `json:"metaOutputIndex,omitempty"`
	// The total supply of this NFT.
	TokenSupply *int64 `json:"tokenSupply,omitempty"`
	// The index of this NFT.
	TokenIndex *int64 `json:"tokenIndex,omitempty"`
	// Mvc value of the utxo(Irrelavant to token value)
	Satoshi *int64 `json:"satoshi,omitempty"`
	// Mvc value of the utxo(In string format)
	SatoshiString *string `json:"satoshiString,omitempty"`
	// Flag used for paging
	Flag *string `json:"flag,omitempty"`
}

// NewContractNftUtxo instantiates a new ContractNftUtxo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractNftUtxo() *ContractNftUtxo {
	this := ContractNftUtxo{}
	return &this
}

// NewContractNftUtxoWithDefaults instantiates a new ContractNftUtxo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractNftUtxoWithDefaults() *ContractNftUtxo {
	this := ContractNftUtxo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ContractNftUtxo) SetAddress(v string) {
	o.Address = &v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *ContractNftUtxo) SetTxid(v string) {
	o.Txid = &v
}

// GetTxIndex returns the TxIndex field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetTxIndex() int32 {
	if o == nil || o.TxIndex == nil {
		var ret int32
		return ret
	}
	return *o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetTxIndexOk() (*int32, bool) {
	if o == nil || o.TxIndex == nil {
		return nil, false
	}
	return o.TxIndex, true
}

// HasTxIndex returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasTxIndex() bool {
	if o != nil && o.TxIndex != nil {
		return true
	}

	return false
}

// SetTxIndex gets a reference to the given int32 and assigns it to the TxIndex field.
func (o *ContractNftUtxo) SetTxIndex(v int32) {
	o.TxIndex = &v
}

// GetCodeHash returns the CodeHash field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetCodeHash() string {
	if o == nil || o.CodeHash == nil {
		var ret string
		return ret
	}
	return *o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetCodeHashOk() (*string, bool) {
	if o == nil || o.CodeHash == nil {
		return nil, false
	}
	return o.CodeHash, true
}

// HasCodeHash returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasCodeHash() bool {
	if o != nil && o.CodeHash != nil {
		return true
	}

	return false
}

// SetCodeHash gets a reference to the given string and assigns it to the CodeHash field.
func (o *ContractNftUtxo) SetCodeHash(v string) {
	o.CodeHash = &v
}

// GetGenesis returns the Genesis field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetGenesis() string {
	if o == nil || o.Genesis == nil {
		var ret string
		return ret
	}
	return *o.Genesis
}

// GetGenesisOk returns a tuple with the Genesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetGenesisOk() (*string, bool) {
	if o == nil || o.Genesis == nil {
		return nil, false
	}
	return o.Genesis, true
}

// HasGenesis returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasGenesis() bool {
	if o != nil && o.Genesis != nil {
		return true
	}

	return false
}

// SetGenesis gets a reference to the given string and assigns it to the Genesis field.
func (o *ContractNftUtxo) SetGenesis(v string) {
	o.Genesis = &v
}

// GetSensibleId returns the SensibleId field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetSensibleId() string {
	if o == nil || o.SensibleId == nil {
		var ret string
		return ret
	}
	return *o.SensibleId
}

// GetSensibleIdOk returns a tuple with the SensibleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetSensibleIdOk() (*string, bool) {
	if o == nil || o.SensibleId == nil {
		return nil, false
	}
	return o.SensibleId, true
}

// HasSensibleId returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasSensibleId() bool {
	if o != nil && o.SensibleId != nil {
		return true
	}

	return false
}

// SetSensibleId gets a reference to the given string and assigns it to the SensibleId field.
func (o *ContractNftUtxo) SetSensibleId(v string) {
	o.SensibleId = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ContractNftUtxo) SetHeight(v int32) {
	o.Height = &v
}

// GetMetaTxid returns the MetaTxid field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetMetaTxid() string {
	if o == nil || o.MetaTxid == nil {
		var ret string
		return ret
	}
	return *o.MetaTxid
}

// GetMetaTxidOk returns a tuple with the MetaTxid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetMetaTxidOk() (*string, bool) {
	if o == nil || o.MetaTxid == nil {
		return nil, false
	}
	return o.MetaTxid, true
}

// HasMetaTxid returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasMetaTxid() bool {
	if o != nil && o.MetaTxid != nil {
		return true
	}

	return false
}

// SetMetaTxid gets a reference to the given string and assigns it to the MetaTxid field.
func (o *ContractNftUtxo) SetMetaTxid(v string) {
	o.MetaTxid = &v
}

// GetMetaOutputIndex returns the MetaOutputIndex field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetMetaOutputIndex() int32 {
	if o == nil || o.MetaOutputIndex == nil {
		var ret int32
		return ret
	}
	return *o.MetaOutputIndex
}

// GetMetaOutputIndexOk returns a tuple with the MetaOutputIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetMetaOutputIndexOk() (*int32, bool) {
	if o == nil || o.MetaOutputIndex == nil {
		return nil, false
	}
	return o.MetaOutputIndex, true
}

// HasMetaOutputIndex returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasMetaOutputIndex() bool {
	if o != nil && o.MetaOutputIndex != nil {
		return true
	}

	return false
}

// SetMetaOutputIndex gets a reference to the given int32 and assigns it to the MetaOutputIndex field.
func (o *ContractNftUtxo) SetMetaOutputIndex(v int32) {
	o.MetaOutputIndex = &v
}

// GetTokenSupply returns the TokenSupply field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetTokenSupply() int64 {
	if o == nil || o.TokenSupply == nil {
		var ret int64
		return ret
	}
	return *o.TokenSupply
}

// GetTokenSupplyOk returns a tuple with the TokenSupply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetTokenSupplyOk() (*int64, bool) {
	if o == nil || o.TokenSupply == nil {
		return nil, false
	}
	return o.TokenSupply, true
}

// HasTokenSupply returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasTokenSupply() bool {
	if o != nil && o.TokenSupply != nil {
		return true
	}

	return false
}

// SetTokenSupply gets a reference to the given int64 and assigns it to the TokenSupply field.
func (o *ContractNftUtxo) SetTokenSupply(v int64) {
	o.TokenSupply = &v
}

// GetTokenIndex returns the TokenIndex field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetTokenIndex() int64 {
	if o == nil || o.TokenIndex == nil {
		var ret int64
		return ret
	}
	return *o.TokenIndex
}

// GetTokenIndexOk returns a tuple with the TokenIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetTokenIndexOk() (*int64, bool) {
	if o == nil || o.TokenIndex == nil {
		return nil, false
	}
	return o.TokenIndex, true
}

// HasTokenIndex returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasTokenIndex() bool {
	if o != nil && o.TokenIndex != nil {
		return true
	}

	return false
}

// SetTokenIndex gets a reference to the given int64 and assigns it to the TokenIndex field.
func (o *ContractNftUtxo) SetTokenIndex(v int64) {
	o.TokenIndex = &v
}

// GetSatoshi returns the Satoshi field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetSatoshi() int64 {
	if o == nil || o.Satoshi == nil {
		var ret int64
		return ret
	}
	return *o.Satoshi
}

// GetSatoshiOk returns a tuple with the Satoshi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetSatoshiOk() (*int64, bool) {
	if o == nil || o.Satoshi == nil {
		return nil, false
	}
	return o.Satoshi, true
}

// HasSatoshi returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasSatoshi() bool {
	if o != nil && o.Satoshi != nil {
		return true
	}

	return false
}

// SetSatoshi gets a reference to the given int64 and assigns it to the Satoshi field.
func (o *ContractNftUtxo) SetSatoshi(v int64) {
	o.Satoshi = &v
}

// GetSatoshiString returns the SatoshiString field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetSatoshiString() string {
	if o == nil || o.SatoshiString == nil {
		var ret string
		return ret
	}
	return *o.SatoshiString
}

// GetSatoshiStringOk returns a tuple with the SatoshiString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetSatoshiStringOk() (*string, bool) {
	if o == nil || o.SatoshiString == nil {
		return nil, false
	}
	return o.SatoshiString, true
}

// HasSatoshiString returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasSatoshiString() bool {
	if o != nil && o.SatoshiString != nil {
		return true
	}

	return false
}

// SetSatoshiString gets a reference to the given string and assigns it to the SatoshiString field.
func (o *ContractNftUtxo) SetSatoshiString(v string) {
	o.SatoshiString = &v
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *ContractNftUtxo) GetFlag() string {
	if o == nil || o.Flag == nil {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractNftUtxo) GetFlagOk() (*string, bool) {
	if o == nil || o.Flag == nil {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *ContractNftUtxo) HasFlag() bool {
	if o != nil && o.Flag != nil {
		return true
	}

	return false
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *ContractNftUtxo) SetFlag(v string) {
	o.Flag = &v
}

func (o ContractNftUtxo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
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
	if o.MetaTxid != nil {
		toSerialize["metaTxid"] = o.MetaTxid
	}
	if o.MetaOutputIndex != nil {
		toSerialize["metaOutputIndex"] = o.MetaOutputIndex
	}
	if o.TokenSupply != nil {
		toSerialize["tokenSupply"] = o.TokenSupply
	}
	if o.TokenIndex != nil {
		toSerialize["tokenIndex"] = o.TokenIndex
	}
	if o.Satoshi != nil {
		toSerialize["satoshi"] = o.Satoshi
	}
	if o.SatoshiString != nil {
		toSerialize["satoshiString"] = o.SatoshiString
	}
	if o.Flag != nil {
		toSerialize["flag"] = o.Flag
	}
	return json.Marshal(toSerialize)
}

type NullableContractNftUtxo struct {
	value *ContractNftUtxo
	isSet bool
}

func (v NullableContractNftUtxo) Get() *ContractNftUtxo {
	return v.value
}

func (v *NullableContractNftUtxo) Set(val *ContractNftUtxo) {
	v.value = val
	v.isSet = true
}

func (v NullableContractNftUtxo) IsSet() bool {
	return v.isSet
}

func (v *NullableContractNftUtxo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractNftUtxo(val *ContractNftUtxo) *NullableContractNftUtxo {
	return &NullableContractNftUtxo{value: val, isSet: true}
}

func (v NullableContractNftUtxo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractNftUtxo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
