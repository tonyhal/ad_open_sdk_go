/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferCreateV30TransferDirection
type CgTransferWalletTransferCreateV30TransferDirection string

// List of cg_transfer_wallet_transfer_create_v3.0_transfer_direction
const (
	TRANSFER_IN_CgTransferWalletTransferCreateV30TransferDirection  CgTransferWalletTransferCreateV30TransferDirection = "TRANSFER_IN"
	TRANSFER_OUT_CgTransferWalletTransferCreateV30TransferDirection CgTransferWalletTransferCreateV30TransferDirection = "TRANSFER_OUT"
)

// Ptr returns reference to cg_transfer_wallet_transfer_create_v3.0_transfer_direction value
func (v CgTransferWalletTransferCreateV30TransferDirection) Ptr() *CgTransferWalletTransferCreateV30TransferDirection {
	return &v
}
