/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CustomerCenterAdvertiserTransferableListV2TransferAccountType
type CustomerCenterAdvertiserTransferableListV2TransferAccountType string

// List of customer_center_advertiser_transferable_list_v2_transfer_account_type
const (
	PAYEE_CustomerCenterAdvertiserTransferableListV2TransferAccountType    CustomerCenterAdvertiserTransferableListV2TransferAccountType = "PAYEE"
	REMITTER_CustomerCenterAdvertiserTransferableListV2TransferAccountType CustomerCenterAdvertiserTransferableListV2TransferAccountType = "REMITTER"
)

// All allowed values of CustomerCenterAdvertiserTransferableListV2TransferAccountType enum
var AllowedCustomerCenterAdvertiserTransferableListV2TransferAccountTypeEnumValues = []CustomerCenterAdvertiserTransferableListV2TransferAccountType{
	"PAYEE",
	"REMITTER",
}

// NewCustomerCenterAdvertiserTransferableListV2TransferAccountTypeFromValue returns a pointer to a valid CustomerCenterAdvertiserTransferableListV2TransferAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomerCenterAdvertiserTransferableListV2TransferAccountTypeFromValue(v string) (*CustomerCenterAdvertiserTransferableListV2TransferAccountType, error) {
	ev := CustomerCenterAdvertiserTransferableListV2TransferAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomerCenterAdvertiserTransferableListV2TransferAccountType: valid values are %v", v, AllowedCustomerCenterAdvertiserTransferableListV2TransferAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomerCenterAdvertiserTransferableListV2TransferAccountType) IsValid() bool {
	for _, existing := range AllowedCustomerCenterAdvertiserTransferableListV2TransferAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to customer_center_advertiser_transferable_list_v2_transfer_account_type value
func (v CustomerCenterAdvertiserTransferableListV2TransferAccountType) Ptr() *CustomerCenterAdvertiserTransferableListV2TransferAccountType {
	return &v
}
