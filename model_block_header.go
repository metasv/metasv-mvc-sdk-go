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

// BlockHeader struct for BlockHeader
type BlockHeader struct {
	// Block height.
	Height *int64 `json:"height,omitempty"`
	// Block hash.
	BlockHash *string `json:"blockHash,omitempty"`
	// Block version.
	Version *int64 `json:"version,omitempty"`
	// The block hash of the previous block.
	PrevBlockHash *string `json:"prevBlockHash,omitempty"`
	// Hex formatted block merkle root.
	MerkleRoot *string `json:"merkleRoot,omitempty"`
	// Block timestamp.
	Timestamp *int64 `json:"timestamp,omitempty"`
	// Block median timestamp.
	MedianTime *int64 `json:"medianTime,omitempty"`
	// Miner total rewards, including miner fee.
	Reward *int64 `json:"reward,omitempty"`
	// Guessed miner name.
	Miner *string `json:"miner,omitempty"`
	// Miner address that received rewards.
	MinerAddress *string `json:"minerAddress,omitempty"`
	// Total txs count included in the block.
	TxCount *int64 `json:"txCount,omitempty"`
	// Total input count in the block.
	InputCount *int64 `json:"inputCount,omitempty"`
	// Total output count in the block
	OutputCount *int64 `json:"outputCount,omitempty"`
	// Size of the block
	Size *int64 `json:"size,omitempty"`
	// Target bits.
	Bits *int64 `json:"bits,omitempty"`
	// Block nonce
	Nonce *int64 `json:"nonce,omitempty"`
	// Hex formated coinbase info.
	Coinbase *string `json:"coinbase,omitempty"`
}

// NewBlockHeader instantiates a new BlockHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockHeader() *BlockHeader {
	this := BlockHeader{}
	return &this
}

// NewBlockHeaderWithDefaults instantiates a new BlockHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockHeaderWithDefaults() *BlockHeader {
	this := BlockHeader{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *BlockHeader) GetHeight() int64 {
	if o == nil || o.Height == nil {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetHeightOk() (*int64, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *BlockHeader) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *BlockHeader) SetHeight(v int64) {
	o.Height = &v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *BlockHeader) GetBlockHash() string {
	if o == nil || o.BlockHash == nil {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetBlockHashOk() (*string, bool) {
	if o == nil || o.BlockHash == nil {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *BlockHeader) HasBlockHash() bool {
	if o != nil && o.BlockHash != nil {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *BlockHeader) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BlockHeader) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BlockHeader) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *BlockHeader) SetVersion(v int64) {
	o.Version = &v
}

// GetPrevBlockHash returns the PrevBlockHash field value if set, zero value otherwise.
func (o *BlockHeader) GetPrevBlockHash() string {
	if o == nil || o.PrevBlockHash == nil {
		var ret string
		return ret
	}
	return *o.PrevBlockHash
}

// GetPrevBlockHashOk returns a tuple with the PrevBlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetPrevBlockHashOk() (*string, bool) {
	if o == nil || o.PrevBlockHash == nil {
		return nil, false
	}
	return o.PrevBlockHash, true
}

// HasPrevBlockHash returns a boolean if a field has been set.
func (o *BlockHeader) HasPrevBlockHash() bool {
	if o != nil && o.PrevBlockHash != nil {
		return true
	}

	return false
}

// SetPrevBlockHash gets a reference to the given string and assigns it to the PrevBlockHash field.
func (o *BlockHeader) SetPrevBlockHash(v string) {
	o.PrevBlockHash = &v
}

// GetMerkleRoot returns the MerkleRoot field value if set, zero value otherwise.
func (o *BlockHeader) GetMerkleRoot() string {
	if o == nil || o.MerkleRoot == nil {
		var ret string
		return ret
	}
	return *o.MerkleRoot
}

// GetMerkleRootOk returns a tuple with the MerkleRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetMerkleRootOk() (*string, bool) {
	if o == nil || o.MerkleRoot == nil {
		return nil, false
	}
	return o.MerkleRoot, true
}

// HasMerkleRoot returns a boolean if a field has been set.
func (o *BlockHeader) HasMerkleRoot() bool {
	if o != nil && o.MerkleRoot != nil {
		return true
	}

	return false
}

// SetMerkleRoot gets a reference to the given string and assigns it to the MerkleRoot field.
func (o *BlockHeader) SetMerkleRoot(v string) {
	o.MerkleRoot = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BlockHeader) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BlockHeader) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *BlockHeader) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetMedianTime returns the MedianTime field value if set, zero value otherwise.
func (o *BlockHeader) GetMedianTime() int64 {
	if o == nil || o.MedianTime == nil {
		var ret int64
		return ret
	}
	return *o.MedianTime
}

// GetMedianTimeOk returns a tuple with the MedianTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetMedianTimeOk() (*int64, bool) {
	if o == nil || o.MedianTime == nil {
		return nil, false
	}
	return o.MedianTime, true
}

// HasMedianTime returns a boolean if a field has been set.
func (o *BlockHeader) HasMedianTime() bool {
	if o != nil && o.MedianTime != nil {
		return true
	}

	return false
}

// SetMedianTime gets a reference to the given int64 and assigns it to the MedianTime field.
func (o *BlockHeader) SetMedianTime(v int64) {
	o.MedianTime = &v
}

// GetReward returns the Reward field value if set, zero value otherwise.
func (o *BlockHeader) GetReward() int64 {
	if o == nil || o.Reward == nil {
		var ret int64
		return ret
	}
	return *o.Reward
}

// GetRewardOk returns a tuple with the Reward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetRewardOk() (*int64, bool) {
	if o == nil || o.Reward == nil {
		return nil, false
	}
	return o.Reward, true
}

// HasReward returns a boolean if a field has been set.
func (o *BlockHeader) HasReward() bool {
	if o != nil && o.Reward != nil {
		return true
	}

	return false
}

// SetReward gets a reference to the given int64 and assigns it to the Reward field.
func (o *BlockHeader) SetReward(v int64) {
	o.Reward = &v
}

// GetMiner returns the Miner field value if set, zero value otherwise.
func (o *BlockHeader) GetMiner() string {
	if o == nil || o.Miner == nil {
		var ret string
		return ret
	}
	return *o.Miner
}

// GetMinerOk returns a tuple with the Miner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetMinerOk() (*string, bool) {
	if o == nil || o.Miner == nil {
		return nil, false
	}
	return o.Miner, true
}

// HasMiner returns a boolean if a field has been set.
func (o *BlockHeader) HasMiner() bool {
	if o != nil && o.Miner != nil {
		return true
	}

	return false
}

// SetMiner gets a reference to the given string and assigns it to the Miner field.
func (o *BlockHeader) SetMiner(v string) {
	o.Miner = &v
}

// GetMinerAddress returns the MinerAddress field value if set, zero value otherwise.
func (o *BlockHeader) GetMinerAddress() string {
	if o == nil || o.MinerAddress == nil {
		var ret string
		return ret
	}
	return *o.MinerAddress
}

// GetMinerAddressOk returns a tuple with the MinerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetMinerAddressOk() (*string, bool) {
	if o == nil || o.MinerAddress == nil {
		return nil, false
	}
	return o.MinerAddress, true
}

// HasMinerAddress returns a boolean if a field has been set.
func (o *BlockHeader) HasMinerAddress() bool {
	if o != nil && o.MinerAddress != nil {
		return true
	}

	return false
}

// SetMinerAddress gets a reference to the given string and assigns it to the MinerAddress field.
func (o *BlockHeader) SetMinerAddress(v string) {
	o.MinerAddress = &v
}

// GetTxCount returns the TxCount field value if set, zero value otherwise.
func (o *BlockHeader) GetTxCount() int64 {
	if o == nil || o.TxCount == nil {
		var ret int64
		return ret
	}
	return *o.TxCount
}

// GetTxCountOk returns a tuple with the TxCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetTxCountOk() (*int64, bool) {
	if o == nil || o.TxCount == nil {
		return nil, false
	}
	return o.TxCount, true
}

// HasTxCount returns a boolean if a field has been set.
func (o *BlockHeader) HasTxCount() bool {
	if o != nil && o.TxCount != nil {
		return true
	}

	return false
}

// SetTxCount gets a reference to the given int64 and assigns it to the TxCount field.
func (o *BlockHeader) SetTxCount(v int64) {
	o.TxCount = &v
}

// GetInputCount returns the InputCount field value if set, zero value otherwise.
func (o *BlockHeader) GetInputCount() int64 {
	if o == nil || o.InputCount == nil {
		var ret int64
		return ret
	}
	return *o.InputCount
}

// GetInputCountOk returns a tuple with the InputCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetInputCountOk() (*int64, bool) {
	if o == nil || o.InputCount == nil {
		return nil, false
	}
	return o.InputCount, true
}

// HasInputCount returns a boolean if a field has been set.
func (o *BlockHeader) HasInputCount() bool {
	if o != nil && o.InputCount != nil {
		return true
	}

	return false
}

// SetInputCount gets a reference to the given int64 and assigns it to the InputCount field.
func (o *BlockHeader) SetInputCount(v int64) {
	o.InputCount = &v
}

// GetOutputCount returns the OutputCount field value if set, zero value otherwise.
func (o *BlockHeader) GetOutputCount() int64 {
	if o == nil || o.OutputCount == nil {
		var ret int64
		return ret
	}
	return *o.OutputCount
}

// GetOutputCountOk returns a tuple with the OutputCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetOutputCountOk() (*int64, bool) {
	if o == nil || o.OutputCount == nil {
		return nil, false
	}
	return o.OutputCount, true
}

// HasOutputCount returns a boolean if a field has been set.
func (o *BlockHeader) HasOutputCount() bool {
	if o != nil && o.OutputCount != nil {
		return true
	}

	return false
}

// SetOutputCount gets a reference to the given int64 and assigns it to the OutputCount field.
func (o *BlockHeader) SetOutputCount(v int64) {
	o.OutputCount = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *BlockHeader) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *BlockHeader) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *BlockHeader) SetSize(v int64) {
	o.Size = &v
}

// GetBits returns the Bits field value if set, zero value otherwise.
func (o *BlockHeader) GetBits() int64 {
	if o == nil || o.Bits == nil {
		var ret int64
		return ret
	}
	return *o.Bits
}

// GetBitsOk returns a tuple with the Bits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetBitsOk() (*int64, bool) {
	if o == nil || o.Bits == nil {
		return nil, false
	}
	return o.Bits, true
}

// HasBits returns a boolean if a field has been set.
func (o *BlockHeader) HasBits() bool {
	if o != nil && o.Bits != nil {
		return true
	}

	return false
}

// SetBits gets a reference to the given int64 and assigns it to the Bits field.
func (o *BlockHeader) SetBits(v int64) {
	o.Bits = &v
}

// GetNonce returns the Nonce field value if set, zero value otherwise.
func (o *BlockHeader) GetNonce() int64 {
	if o == nil || o.Nonce == nil {
		var ret int64
		return ret
	}
	return *o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetNonceOk() (*int64, bool) {
	if o == nil || o.Nonce == nil {
		return nil, false
	}
	return o.Nonce, true
}

// HasNonce returns a boolean if a field has been set.
func (o *BlockHeader) HasNonce() bool {
	if o != nil && o.Nonce != nil {
		return true
	}

	return false
}

// SetNonce gets a reference to the given int64 and assigns it to the Nonce field.
func (o *BlockHeader) SetNonce(v int64) {
	o.Nonce = &v
}

// GetCoinbase returns the Coinbase field value if set, zero value otherwise.
func (o *BlockHeader) GetCoinbase() string {
	if o == nil || o.Coinbase == nil {
		var ret string
		return ret
	}
	return *o.Coinbase
}

// GetCoinbaseOk returns a tuple with the Coinbase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockHeader) GetCoinbaseOk() (*string, bool) {
	if o == nil || o.Coinbase == nil {
		return nil, false
	}
	return o.Coinbase, true
}

// HasCoinbase returns a boolean if a field has been set.
func (o *BlockHeader) HasCoinbase() bool {
	if o != nil && o.Coinbase != nil {
		return true
	}

	return false
}

// SetCoinbase gets a reference to the given string and assigns it to the Coinbase field.
func (o *BlockHeader) SetCoinbase(v string) {
	o.Coinbase = &v
}

func (o BlockHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.BlockHash != nil {
		toSerialize["blockHash"] = o.BlockHash
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.PrevBlockHash != nil {
		toSerialize["prevBlockHash"] = o.PrevBlockHash
	}
	if o.MerkleRoot != nil {
		toSerialize["merkleRoot"] = o.MerkleRoot
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.MedianTime != nil {
		toSerialize["medianTime"] = o.MedianTime
	}
	if o.Reward != nil {
		toSerialize["reward"] = o.Reward
	}
	if o.Miner != nil {
		toSerialize["miner"] = o.Miner
	}
	if o.MinerAddress != nil {
		toSerialize["minerAddress"] = o.MinerAddress
	}
	if o.TxCount != nil {
		toSerialize["txCount"] = o.TxCount
	}
	if o.InputCount != nil {
		toSerialize["inputCount"] = o.InputCount
	}
	if o.OutputCount != nil {
		toSerialize["outputCount"] = o.OutputCount
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Bits != nil {
		toSerialize["bits"] = o.Bits
	}
	if o.Nonce != nil {
		toSerialize["nonce"] = o.Nonce
	}
	if o.Coinbase != nil {
		toSerialize["coinbase"] = o.Coinbase
	}
	return json.Marshal(toSerialize)
}

type NullableBlockHeader struct {
	value *BlockHeader
	isSet bool
}

func (v NullableBlockHeader) Get() *BlockHeader {
	return v.value
}

func (v *NullableBlockHeader) Set(val *BlockHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockHeader(val *BlockHeader) *NullableBlockHeader {
	return &NullableBlockHeader{value: val, isSet: true}
}

func (v NullableBlockHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
