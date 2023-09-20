/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanUniPromotionListV10DataAdListAdInfoStatus
type QianchuanUniPromotionListV10DataAdListAdInfoStatus string

// List of qianchuan_uni_promotion_list_v1.0_data_ad_list_ad_info_status
const (
	ADVERTISER_OFFLINE_BUDGET_QianchuanUniPromotionListV10DataAdListAdInfoStatus     QianchuanUniPromotionListV10DataAdListAdInfoStatus = "ADVERTISER_OFFLINE_BUDGET"
	ADVERTISER_PRE_OFFLINE_BUDGET_QianchuanUniPromotionListV10DataAdListAdInfoStatus QianchuanUniPromotionListV10DataAdListAdInfoStatus = "ADVERTISER_PRE_OFFLINE_BUDGET"
	AUDIT_QianchuanUniPromotionListV10DataAdListAdInfoStatus                         QianchuanUniPromotionListV10DataAdListAdInfoStatus = "AUDIT"
	AUDIT_STATUS_ERROR_QianchuanUniPromotionListV10DataAdListAdInfoStatus            QianchuanUniPromotionListV10DataAdListAdInfoStatus = "AUDIT_STATUS_ERROR"
	AWEME_ACCOUNT_DISABLED_QianchuanUniPromotionListV10DataAdListAdInfoStatus        QianchuanUniPromotionListV10DataAdListAdInfoStatus = "AWEME_ACCOUNT_DISABLED"
	CAMPAIGN_DISABLE_QianchuanUniPromotionListV10DataAdListAdInfoStatus              QianchuanUniPromotionListV10DataAdListAdInfoStatus = "CAMPAIGN_DISABLE"
	CAMPAIGN_OFFLINE_BUDGET_QianchuanUniPromotionListV10DataAdListAdInfoStatus       QianchuanUniPromotionListV10DataAdListAdInfoStatus = "CAMPAIGN_OFFLINE_BUDGET"
	CAMPAIGN_PRE_OFFLINE_BUDGET_QianchuanUniPromotionListV10DataAdListAdInfoStatus   QianchuanUniPromotionListV10DataAdListAdInfoStatus = "CAMPAIGN_PRE_OFFLINE_BUDGET"
	CREATE_QianchuanUniPromotionListV10DataAdListAdInfoStatus                        QianchuanUniPromotionListV10DataAdListAdInfoStatus = "CREATE"
	DELETE_QianchuanUniPromotionListV10DataAdListAdInfoStatus                        QianchuanUniPromotionListV10DataAdListAdInfoStatus = "DELETE"
	DELIVERY_OK_QianchuanUniPromotionListV10DataAdListAdInfoStatus                   QianchuanUniPromotionListV10DataAdListAdInfoStatus = "DELIVERY_OK"
	DISABLE_QianchuanUniPromotionListV10DataAdListAdInfoStatus                       QianchuanUniPromotionListV10DataAdListAdInfoStatus = "DISABLE"
	DRAFT_QianchuanUniPromotionListV10DataAdListAdInfoStatus                         QianchuanUniPromotionListV10DataAdListAdInfoStatus = "DRAFT"
	ERROR_QianchuanUniPromotionListV10DataAdListAdInfoStatus                         QianchuanUniPromotionListV10DataAdListAdInfoStatus = "ERROR"
	EXTERNAL_URL_DISABLE_QianchuanUniPromotionListV10DataAdListAdInfoStatus          QianchuanUniPromotionListV10DataAdListAdInfoStatus = "EXTERNAL_URL_DISABLE"
	FROZEN_QianchuanUniPromotionListV10DataAdListAdInfoStatus                        QianchuanUniPromotionListV10DataAdListAdInfoStatus = "FROZEN"
	LIVE_ROOM_OFF_QianchuanUniPromotionListV10DataAdListAdInfoStatus                 QianchuanUniPromotionListV10DataAdListAdInfoStatus = "LIVE_ROOM_OFF"
	NO_SCHEDULE_QianchuanUniPromotionListV10DataAdListAdInfoStatus                   QianchuanUniPromotionListV10DataAdListAdInfoStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_QianchuanUniPromotionListV10DataAdListAdInfoStatus                 QianchuanUniPromotionListV10DataAdListAdInfoStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_QianchuanUniPromotionListV10DataAdListAdInfoStatus               QianchuanUniPromotionListV10DataAdListAdInfoStatus = "OFFLINE_BALANCE"
	OFFLINE_BUDGET_QianchuanUniPromotionListV10DataAdListAdInfoStatus                QianchuanUniPromotionListV10DataAdListAdInfoStatus = "OFFLINE_BUDGET"
	PRE_OFFLINE_BUDGET_QianchuanUniPromotionListV10DataAdListAdInfoStatus            QianchuanUniPromotionListV10DataAdListAdInfoStatus = "PRE_OFFLINE_BUDGET"
	PRE_ONLINE_QianchuanUniPromotionListV10DataAdListAdInfoStatus                    QianchuanUniPromotionListV10DataAdListAdInfoStatus = "PRE_ONLINE"
	QUOTA_DISABLE_QianchuanUniPromotionListV10DataAdListAdInfoStatus                 QianchuanUniPromotionListV10DataAdListAdInfoStatus = "QUOTA_DISABLE"
	REAUDIT_QianchuanUniPromotionListV10DataAdListAdInfoStatus                       QianchuanUniPromotionListV10DataAdListAdInfoStatus = "REAUDIT"
	SYSTEM_DISABLE_QianchuanUniPromotionListV10DataAdListAdInfoStatus                QianchuanUniPromotionListV10DataAdListAdInfoStatus = "SYSTEM_DISABLE"
	TIME_DONE_QianchuanUniPromotionListV10DataAdListAdInfoStatus                     QianchuanUniPromotionListV10DataAdListAdInfoStatus = "TIME_DONE"
	TIME_NO_REACH_QianchuanUniPromotionListV10DataAdListAdInfoStatus                 QianchuanUniPromotionListV10DataAdListAdInfoStatus = "TIME_NO_REACH"
)

// All allowed values of QianchuanUniPromotionListV10DataAdListAdInfoStatus enum
var AllowedQianchuanUniPromotionListV10DataAdListAdInfoStatusEnumValues = []QianchuanUniPromotionListV10DataAdListAdInfoStatus{
	"ADVERTISER_OFFLINE_BUDGET",
	"ADVERTISER_PRE_OFFLINE_BUDGET",
	"AUDIT",
	"AUDIT_STATUS_ERROR",
	"AWEME_ACCOUNT_DISABLED",
	"CAMPAIGN_DISABLE",
	"CAMPAIGN_OFFLINE_BUDGET",
	"CAMPAIGN_PRE_OFFLINE_BUDGET",
	"CREATE",
	"DELETE",
	"DELIVERY_OK",
	"DISABLE",
	"DRAFT",
	"ERROR",
	"EXTERNAL_URL_DISABLE",
	"FROZEN",
	"LIVE_ROOM_OFF",
	"NO_SCHEDULE",
	"OFFLINE_AUDIT",
	"OFFLINE_BALANCE",
	"OFFLINE_BUDGET",
	"PRE_OFFLINE_BUDGET",
	"PRE_ONLINE",
	"QUOTA_DISABLE",
	"REAUDIT",
	"SYSTEM_DISABLE",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewQianchuanUniPromotionListV10DataAdListAdInfoStatusFromValue returns a pointer to a valid QianchuanUniPromotionListV10DataAdListAdInfoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionListV10DataAdListAdInfoStatusFromValue(v string) (*QianchuanUniPromotionListV10DataAdListAdInfoStatus, error) {
	ev := QianchuanUniPromotionListV10DataAdListAdInfoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionListV10DataAdListAdInfoStatus: valid values are %v", v, AllowedQianchuanUniPromotionListV10DataAdListAdInfoStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionListV10DataAdListAdInfoStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionListV10DataAdListAdInfoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_data_ad_list_ad_info_status value
func (v QianchuanUniPromotionListV10DataAdListAdInfoStatus) Ptr() *QianchuanUniPromotionListV10DataAdListAdInfoStatus {
	return &v
}
