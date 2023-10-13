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

// QianchuanAdGetV10DataListStatus
type QianchuanAdGetV10DataListStatus string

// List of qianchuan_ad_get_v1.0_data_list_status
const (
	ADVERTISER_OFFLINE_BUDGET_QianchuanAdGetV10DataListStatus     QianchuanAdGetV10DataListStatus = "ADVERTISER_OFFLINE_BUDGET"
	ADVERTISER_PRE_OFFLINE_BUDGET_QianchuanAdGetV10DataListStatus QianchuanAdGetV10DataListStatus = "ADVERTISER_PRE_OFFLINE_BUDGET"
	AUDIT_QianchuanAdGetV10DataListStatus                         QianchuanAdGetV10DataListStatus = "AUDIT"
	AUDIT_STATUS_ERROR_QianchuanAdGetV10DataListStatus            QianchuanAdGetV10DataListStatus = "AUDIT_STATUS_ERROR"
	AWEME_ACCOUNT_DISABLED_QianchuanAdGetV10DataListStatus        QianchuanAdGetV10DataListStatus = "AWEME_ACCOUNT_DISABLED"
	CAMPAIGN_DISABLE_QianchuanAdGetV10DataListStatus              QianchuanAdGetV10DataListStatus = "CAMPAIGN_DISABLE"
	CAMPAIGN_OFFLINE_BUDGET_QianchuanAdGetV10DataListStatus       QianchuanAdGetV10DataListStatus = "CAMPAIGN_OFFLINE_BUDGET"
	CAMPAIGN_PRE_OFFLINE_BUDGET_QianchuanAdGetV10DataListStatus   QianchuanAdGetV10DataListStatus = "CAMPAIGN_PRE_OFFLINE_BUDGET"
	CREATE_QianchuanAdGetV10DataListStatus                        QianchuanAdGetV10DataListStatus = "CREATE"
	DELETE_QianchuanAdGetV10DataListStatus                        QianchuanAdGetV10DataListStatus = "DELETE"
	DELIVERY_OK_QianchuanAdGetV10DataListStatus                   QianchuanAdGetV10DataListStatus = "DELIVERY_OK"
	DISABLE_QianchuanAdGetV10DataListStatus                       QianchuanAdGetV10DataListStatus = "DISABLE"
	DRAFT_QianchuanAdGetV10DataListStatus                         QianchuanAdGetV10DataListStatus = "DRAFT"
	ERROR_QianchuanAdGetV10DataListStatus                         QianchuanAdGetV10DataListStatus = "ERROR"
	EXTERNAL_URL_DISABLE_QianchuanAdGetV10DataListStatus          QianchuanAdGetV10DataListStatus = "EXTERNAL_URL_DISABLE"
	FROZEN_QianchuanAdGetV10DataListStatus                        QianchuanAdGetV10DataListStatus = "FROZEN"
	LIVE_ROOM_OFF_QianchuanAdGetV10DataListStatus                 QianchuanAdGetV10DataListStatus = "LIVE_ROOM_OFF"
	NO_SCHEDULE_QianchuanAdGetV10DataListStatus                   QianchuanAdGetV10DataListStatus = "NO_SCHEDULE"
	OFFLINE_AUDIT_QianchuanAdGetV10DataListStatus                 QianchuanAdGetV10DataListStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_QianchuanAdGetV10DataListStatus               QianchuanAdGetV10DataListStatus = "OFFLINE_BALANCE"
	OFFLINE_BUDGET_QianchuanAdGetV10DataListStatus                QianchuanAdGetV10DataListStatus = "OFFLINE_BUDGET"
	PRE_OFFLINE_BUDGET_QianchuanAdGetV10DataListStatus            QianchuanAdGetV10DataListStatus = "PRE_OFFLINE_BUDGET"
	PRE_ONLINE_QianchuanAdGetV10DataListStatus                    QianchuanAdGetV10DataListStatus = "PRE_ONLINE"
	QUOTA_DISABLE_QianchuanAdGetV10DataListStatus                 QianchuanAdGetV10DataListStatus = "QUOTA_DISABLE"
	REAUDIT_QianchuanAdGetV10DataListStatus                       QianchuanAdGetV10DataListStatus = "REAUDIT"
	ROI2_DISABLE_QianchuanAdGetV10DataListStatus                  QianchuanAdGetV10DataListStatus = "ROI2_DISABLE"
	SYSTEM_DISABLE_QianchuanAdGetV10DataListStatus                QianchuanAdGetV10DataListStatus = "SYSTEM_DISABLE"
	TIME_DONE_QianchuanAdGetV10DataListStatus                     QianchuanAdGetV10DataListStatus = "TIME_DONE"
	TIME_NO_REACH_QianchuanAdGetV10DataListStatus                 QianchuanAdGetV10DataListStatus = "TIME_NO_REACH"
)

// All allowed values of QianchuanAdGetV10DataListStatus enum
var AllowedQianchuanAdGetV10DataListStatusEnumValues = []QianchuanAdGetV10DataListStatus{
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
	"ROI2_DISABLE",
	"SYSTEM_DISABLE",
	"TIME_DONE",
	"TIME_NO_REACH",
}

// NewQianchuanAdGetV10DataListStatusFromValue returns a pointer to a valid QianchuanAdGetV10DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListStatusFromValue(v string) (*QianchuanAdGetV10DataListStatus, error) {
	ev := QianchuanAdGetV10DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListStatus: valid values are %v", v, AllowedQianchuanAdGetV10DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_status value
func (v QianchuanAdGetV10DataListStatus) Ptr() *QianchuanAdGetV10DataListStatus {
	return &v
}
