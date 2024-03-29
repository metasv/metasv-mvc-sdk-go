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

// ContractFtAddressTx Transaction history of the address for specific ft.
type ContractFtAddressTx struct {
	// Paging flag
	Flag *string `json:"flag,omitempty"`
	// The address of the record
	Address *string `json:"address,omitempty"`
	// Specific codeHash
	CodeHash *string `json:"codeHash,omitempty"`
	// specific genesis
	Genesis *string `json:"genesis,omitempty"`
	// timestamp of the tx
	Time *int64 `json:"time,omitempty"`
	// Block Height of the tx, -1 if not confirmed
	Height *int64 `json:"height,omitempty"`
	// total income of the address from this tx
	Income *int64 `json:"income,omitempty"`
	// total outcome of the address from this tx
	Outcome *int64 `json:"outcome,omitempty"`
	// txid of this record
	Txid *string `json:"txid,omitempty"`
}

// NewContractFtAddressTx instantiates a new ContractFtAddressTx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractFtAddressTx() *ContractFtAddressTx {
	this := ContractFtAddressTx{}
	return &this
}

// NewContractFtAddressTxWithDefaults instantiates a new ContractFtAddressTx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractFtAddressTxWithDefaults() *ContractFtAddressTx {
	this := ContractFtAddressTx{}
	return &this
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *ContractFtAddressTx) GetFlag() string {
	if o == nil || o.Flag == nil {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtAddressTx) GetFlagOk() (*string, bool) {
	if o == nil || o.Flag == nil {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *ContractFtAddressTx) HasFlag() bool {
	if o != nil && o.Flag != nil {
		return true
	}

	return false
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *ContractFtAddressTx) SetFlag(v string) {
	o.Flag = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ContractFtAddressTx) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtAddressTx) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ContractFtAddressTx) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ContractFtAddressTx) SetAddress(v string) {
	o.Address = &v
}

// GetCodeHash returns the CodeHash field value if set, zero value otherwise.
func (o *ContractFtAddressTx) GetCodeHash() string {
	if o == nil || o.CodeHash == nil {
		var ret string
		return ret
	}
	return *o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtAddressTx) GetCodeHashOk() (*string, bool) {
	if o == nil || o.CodeHash == nil {
		return nil, false
	}
	return o.CodeHash, true
}

// HasCodeHash returns a boolean if a field has been set.
func (o *ContractFtAddressTx) HasCodeHash() bool {
	if o != nil && o.CodeHash != nil {
		return true
	}

	return false
}

// SetCodeHash gets a reference to the given string and assigns it to the CodeHash field.
func (o *ContractFtAddressTx) SetCodeHash(v string) {
	o.CodeHash = &v
}

// GetGenesis returns the Genesis field value if set, zero value otherwise.
func (o *ContractFtAddressTx) GetGenesis() string {
	if o == nil || o.Genesis == nil {
		var ret string
		return ret
	}
	return *o.Genesis
}

// GetGenesisOk returns a tuple with the Genesis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtAddressTx) GetGenesisOk() (*string, bool) {
	if o == nil || o.Genesis == nil {
		return nil, false
	}
	return o.Genesis, true
}

// HasGenesis returns a boolean if a field has been set.
func (o *ContractFtAddressTx) HasGenesis() bool {
	if o != nil && o.Genesis != nil {
		return true
	}

	return false
}

// SetGenesis gets a reference to the given string and assigns it to the Genesis field.
func (o *ContractFtAddressTx) SetGenesis(v string) {
	o.Genesis = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *ContractFtAddressTx) GetTime() int64 {
	if o == nil || o.Time == nil {
		var ret int64
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtAddressTx) GetTimeOk() (*int64, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *ContractFtAddressTx) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given int64 and assigns it to the Time field.
func (o *ContractFtAddressTx) SetTime(v int64) {
	o.Time = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ContractFtAddressTx) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtAddressTx) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ContractFtAddressTx) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *ContractFtAddressTx) SetHeight(v int64) {
	o.Height = &v
}

// GetIncome returns the Income field value if set, zero value otherwise.
func (o *ContractFtAddressTx) GetIncome() int64 {
	if o == nil || o.Income == nil {
		var ret int64
		return ret
	}
	return *o.Income
}

// GetIncomeOk returns a tuple with the Income field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtAddressTx) GetIncomeOk() (*int64, bool) {
	if o == nil || o.Income == nil {
		return nil, false
	}
	return o.Income, true
}

// HasIncome returns a boolean if a field has been set.
func (o *ContractFtAddressTx) HasIncome() bool {
	if o != nil && o.Income != nil {
		return true
	}

	return false
}

// SetIncome gets a reference to the given int64 and assigns it to the Income field.
func (o *ContractFtAddressTx) SetIncome(v int64) {
	o.Income = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *ContractFtAddressTx) GetOutcome() int64 {
	if o == nil || o.Outcome == nil {
		var ret int64
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtAddressTx) GetOutcomeOk() (*int64, bool) {
	if o == nil || o.Outcome == nil {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *ContractFtAddressTx) HasOutcome() bool {
	if o != nil && o.Outcome != nil {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given int64 and assigns it to the Outcome field.
func (o *ContractFtAddressTx) SetOutcome(v int64) {
	o.Outcome = &v
}

// GetTxid returns the Txid field value if set, zero value otherwise.
func (o *ContractFtAddressTx) GetTxid() string {
	if o == nil || o.Txid == nil {
		var ret string
		return ret
	}
	return *o.Txid
}

// GetTxidOk returns a tuple with the Txid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractFtAddressTx) GetTxidOk() (*string, bool) {
	if o == nil || o.Txid == nil {
		return nil, false
	}
	return o.Txid, true
}

// HasTxid returns a boolean if a field has been set.
func (o *ContractFtAddressTx) HasTxid() bool {
	if o != nil && o.Txid != nil {
		return true
	}

	return false
}

// SetTxid gets a reference to the given string and assigns it to the Txid field.
func (o *ContractFtAddressTx) SetTxid(v string) {
	o.Txid = &v
}

func (o ContractFtAddressTx) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Flag != nil {
		toSerialize["flag"] = o.Flag
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.CodeHash != nil {
		toSerialize["codeHash"] = o.CodeHash
	}
	if o.Genesis != nil {
		toSerialize["genesis"] = o.Genesis
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Income != nil {
		toSerialize["income"] = o.Income
	}
	if o.Outcome != nil {
		toSerialize["outcome"] = o.Outcome
	}
	if o.Txid != nil {
		toSerialize["txid"] = o.Txid
	}
	return json.Marshal(toSerialize)
}

type NullableContractFtAddressTx struct {
	value *ContractFtAddressTx
	isSet bool
}

func (v NullableContractFtAddressTx) Get() *ContractFtAddressTx {
	return v.value
}

func (v *NullableContractFtAddressTx) Set(val *ContractFtAddressTx) {
	v.value = val
	v.isSet = true
}

func (v NullableContractFtAddressTx) IsSet() bool {
	return v.isSet
}

func (v *NullableContractFtAddressTx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractFtAddressTx(val *ContractFtAddressTx) *NullableContractFtAddressTx {
	return &NullableContractFtAddressTx{value: val, isSet: true}
}

func (v NullableContractFtAddressTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractFtAddressTx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
