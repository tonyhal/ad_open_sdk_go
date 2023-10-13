/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanUniPromotionListV10DataAdListAdInfoOptStatus
type QianchuanUniPromotionListV10DataAdListAdInfoOptStatus string

// List of qianchuan_uni_promotion_list_v1.0_data_ad_list_ad_info_opt_status
const (
	DELETE_QianchuanUniPromotionListV10DataAdListAdInfoOptStatus        QianchuanUniPromotionListV10DataAdListAdInfoOptStatus = "DELETE"
	DISABLE_QianchuanUniPromotionListV10DataAdListAdInfoOptStatus       QianchuanUniPromotionListV10DataAdListAdInfoOptStatus = "DISABLE"
	ENABLE_QianchuanUniPromotionListV10DataAdListAdInfoOptStatus        QianchuanUniPromotionListV10DataAdListAdInfoOptStatus = "ENABLE"
	QUOTA_DISABLE_QianchuanUniPromotionListV10DataAdListAdInfoOptStatus QianchuanUniPromotionListV10DataAdListAdInfoOptStatus = "QUOTA_DISABLE"
	ROI2_DISABLE_QianchuanUniPromotionListV10DataAdListAdInfoOptStatus  QianchuanUniPromotionListV10DataAdListAdInfoOptStatus = "ROI2_DISABLE"
)

// All allowed values of QianchuanUniPromotionListV10DataAdListAdInfoOptStatus enum
var AllowedQianchuanUniPromotionListV10DataAdListAdInfoOptStatusEnumValues = []QianchuanUniPromotionListV10DataAdListAdInfoOptStatus{
	"DELETE",
	"DISABLE",
	"ENABLE",
	"QUOTA_DISABLE",
	"ROI2_DISABLE",
}

// NewQianchuanUniPromotionListV10DataAdListAdInfoOptStatusFromValue returns a pointer to a valid QianchuanUniPromotionListV10DataAdListAdInfoOptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionListV10DataAdListAdInfoOptStatusFromValue(v string) (*QianchuanUniPromotionListV10DataAdListAdInfoOptStatus, error) {
	ev := QianchuanUniPromotionListV10DataAdListAdInfoOptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionListV10DataAdListAdInfoOptStatus: valid values are %v", v, AllowedQianchuanUniPromotionListV10DataAdListAdInfoOptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionListV10DataAdListAdInfoOptStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionListV10DataAdListAdInfoOptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_data_ad_list_ad_info_opt_status value
func (v QianchuanUniPromotionListV10DataAdListAdInfoOptStatus) Ptr() *QianchuanUniPromotionListV10DataAdListAdInfoOptStatus {
	return &v
}
