/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCompensateStatusGetV10ResponseDataListInner struct for QianchuanAdCompensateStatusGetV10ResponseDataListInner
type QianchuanAdCompensateStatusGetV10ResponseDataListInner struct {
	//
	AdId             *int64                                                     `json:"ad_id,omitempty"`
	CompensateStatus *QianchuanAdCompensateStatusGetV10DataListCompensateStatus `json:"compensate_status,omitempty"`
	//
	Reason *string                                          `json:"reason,omitempty"`
	Status *QianchuanAdCompensateStatusGetV10DataListStatus `json:"status,omitempty"`
	//
	Url *string `json:"url,omitempty"`
}
