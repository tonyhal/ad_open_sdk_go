/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeEstimateProfitV10AudienceBehaviors
type QianchuanAwemeEstimateProfitV10AudienceBehaviors string

// List of qianchuan_aweme_estimate_profit_v1.0_audience_behaviors
const (
	ALL_QianchuanAwemeEstimateProfitV10AudienceBehaviors           QianchuanAwemeEstimateProfitV10AudienceBehaviors = "ALL"
	FOLLOWED_USER_QianchuanAwemeEstimateProfitV10AudienceBehaviors QianchuanAwemeEstimateProfitV10AudienceBehaviors = "FOLLOWED_USER"
	GOODS_SHARE_QianchuanAwemeEstimateProfitV10AudienceBehaviors   QianchuanAwemeEstimateProfitV10AudienceBehaviors = "GOODS_SHARE"
	LIKED_USER_QianchuanAwemeEstimateProfitV10AudienceBehaviors    QianchuanAwemeEstimateProfitV10AudienceBehaviors = "LIKED_USER"
	LIVE_WATCH_QianchuanAwemeEstimateProfitV10AudienceBehaviors    QianchuanAwemeEstimateProfitV10AudienceBehaviors = "LIVE_WATCH"
)

// All allowed values of QianchuanAwemeEstimateProfitV10AudienceBehaviors enum
var AllowedQianchuanAwemeEstimateProfitV10AudienceBehaviorsEnumValues = []QianchuanAwemeEstimateProfitV10AudienceBehaviors{
	"ALL",
	"FOLLOWED_USER",
	"GOODS_SHARE",
	"LIKED_USER",
	"LIVE_WATCH",
}

// NewQianchuanAwemeEstimateProfitV10AudienceBehaviorsFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10AudienceBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10AudienceBehaviorsFromValue(v string) (*QianchuanAwemeEstimateProfitV10AudienceBehaviors, error) {
	ev := QianchuanAwemeEstimateProfitV10AudienceBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10AudienceBehaviors: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10AudienceBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10AudienceBehaviors) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10AudienceBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_audience_behaviors value
func (v QianchuanAwemeEstimateProfitV10AudienceBehaviors) Ptr() *QianchuanAwemeEstimateProfitV10AudienceBehaviors {
	return &v
}
