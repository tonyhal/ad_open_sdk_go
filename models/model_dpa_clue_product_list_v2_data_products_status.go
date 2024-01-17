/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaClueProductListV2DataProductsStatus
type DpaClueProductListV2DataProductsStatus string

// List of dpa_clue_product_list_v2_data_products_status
const (
	STATUS_OFFLINE_DpaClueProductListV2DataProductsStatus DpaClueProductListV2DataProductsStatus = "STATUS_OFFLINE"
	STATUS_ONLINE_DpaClueProductListV2DataProductsStatus  DpaClueProductListV2DataProductsStatus = "STATUS_ONLINE"
)

// All allowed values of DpaClueProductListV2DataProductsStatus enum
var AllowedDpaClueProductListV2DataProductsStatusEnumValues = []DpaClueProductListV2DataProductsStatus{
	"STATUS_OFFLINE",
	"STATUS_ONLINE",
}

// NewDpaClueProductListV2DataProductsStatusFromValue returns a pointer to a valid DpaClueProductListV2DataProductsStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaClueProductListV2DataProductsStatusFromValue(v string) (*DpaClueProductListV2DataProductsStatus, error) {
	ev := DpaClueProductListV2DataProductsStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaClueProductListV2DataProductsStatus: valid values are %v", v, AllowedDpaClueProductListV2DataProductsStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaClueProductListV2DataProductsStatus) IsValid() bool {
	for _, existing := range AllowedDpaClueProductListV2DataProductsStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_clue_product_list_v2_data_products_status value
func (v DpaClueProductListV2DataProductsStatus) Ptr() *DpaClueProductListV2DataProductsStatus {
	return &v
}