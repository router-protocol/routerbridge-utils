// Copyright 2020 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package msg

import (
	"fmt"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type ChainId uint8
type TransferType string
type ResourceId [32]byte
type WidgetId string

func (r ResourceId) Hex() string {
	return fmt.Sprintf("%x", r)
}

type Nonce uint64

func (n Nonce) Big() *big.Int {
	return big.NewInt(int64(n))
}

var FungibleTransfer TransferType = "FungibleTransfer"
var NonFungibleTransfer TransferType = "NonFungibleTransfer"
var GenericTransfer TransferType = "GenericTransfer"

// Message is used as a generic format to communicate between chains
type Message struct {
	Source         ChainId      // Source where message was initiated
	Destination    ChainId      // Destination chain of message
	Type           TransferType // type of bridge transfer
	DepositNonce   Nonce        // Nonce for the deposit
	ResourceId     ResourceId
	DestResourceId ResourceId
	WidgetId       WidgetId
	Payload        []interface{} // data associated with event sequence
}

func NewFungibleTransfer(
	source,
	dest ChainId,
	nonce Nonce,
	srcAmount *big.Int,
	resourceId ResourceId,
	recipient ethcommon.Address,
	stableAmount *big.Int,
	destAmount *big.Int,
	srcToken []byte,
	destToken []byte,
	destResourceId ResourceId,
	destStableToken []byte,
	destStableAmount *big.Int,
	isDestNative *big.Int,
	widgetId WidgetId,
) Message {
	return Message{
		Source:         source,
		Destination:    dest,
		Type:           FungibleTransfer,
		DepositNonce:   nonce,
		ResourceId:     resourceId,
		DestResourceId: destResourceId,
		WidgetId:       widgetId,
		Payload: []interface{}{
			srcAmount.Bytes(),
			stableAmount.Bytes(),
			destStableAmount.Bytes(),
			destAmount.Bytes(),
			recipient,
			srcToken,
			destStableToken,
			destToken,
			isDestNative.Bytes(),
		},
	}
}

func NewNonFungibleTransfer(source, dest ChainId, nonce Nonce, resourceId ResourceId, tokenId *big.Int, recipient []byte, widgetId WidgetId, metadata []byte) Message {
	return Message{
		Source:       source,
		Destination:  dest,
		Type:         NonFungibleTransfer,
		DepositNonce: nonce,
		ResourceId:   resourceId,
		WidgetId:     widgetId,
		Payload: []interface{}{
			tokenId.Bytes(),
			recipient,
			metadata,
		},
	}
}

func NewGenericTransfer(source, dest ChainId, nonce Nonce, resourceId ResourceId, widgetId WidgetId, gasLimit *big.Int, gasPrice *big.Int, metadata []byte) Message {
	return Message{
		Source:       source,
		Destination:  dest,
		Type:         GenericTransfer,
		DepositNonce: nonce,
		ResourceId:   resourceId,
		WidgetId:     widgetId,
		Payload: []interface{}{
			metadata,
			gasLimit,
			gasPrice,
		},
	}
}

func ResourceIdFromSlice(in []byte) ResourceId {
	var res ResourceId
	copy(res[:], in)
	return res
}
