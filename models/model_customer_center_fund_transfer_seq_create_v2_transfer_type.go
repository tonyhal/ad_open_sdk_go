/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CustomerCenterFundTransferSeqCreateV2TransferType
type CustomerCenterFundTransferSeqCreateV2TransferType string

// List of customer_center_fund_transfer_seq_create_v2_transfer_type
const (
	CREDIT_BRAND_CustomerCenterFundTransferSeqCreateV2TransferType     CustomerCenterFundTransferSeqCreateV2TransferType = "CREDIT_BRAND"
	CREDIT_UNIVERSAL_CustomerCenterFundTransferSeqCreateV2TransferType CustomerCenterFundTransferSeqCreateV2TransferType = "CREDIT_UNIVERSAL"
	GRANT_CustomerCenterFundTransferSeqCreateV2TransferType            CustomerCenterFundTransferSeqCreateV2TransferType = "GRANT"
	PREPAY_BID_CustomerCenterFundTransferSeqCreateV2TransferType       CustomerCenterFundTransferSeqCreateV2TransferType = "PREPAY_BID"
	PREPAY_UNIVERSAL_CustomerCenterFundTransferSeqCreateV2TransferType CustomerCenterFundTransferSeqCreateV2TransferType = "PREPAY_UNIVERSAL"
	PREPAY_BRAND_CustomerCenterFundTransferSeqCreateV2TransferType     CustomerCenterFundTransferSeqCreateV2TransferType = "PREPAY_BRAND"
	CREDIT_BID_CustomerCenterFundTransferSeqCreateV2TransferType       CustomerCenterFundTransferSeqCreateV2TransferType = "CREDIT_BID"
)

// Ptr returns reference to customer_center_fund_transfer_seq_create_v2_transfer_type value
func (v CustomerCenterFundTransferSeqCreateV2TransferType) Ptr() *CustomerCenterFundTransferSeqCreateV2TransferType {
	return &v
}
