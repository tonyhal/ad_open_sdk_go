/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionAdStatusUpdateV10ResponseDataResultsInner struct for QianchuanUniPromotionAdStatusUpdateV10ResponseDataResultsInner
type QianchuanUniPromotionAdStatusUpdateV10ResponseDataResultsInner struct {
	//
	AdId  *int64                                                               `json:"ad_id,omitempty"`
	Error *QianchuanUniPromotionAdStatusUpdateV10ResponseDataResultsInnerError `json:"error,omitempty"`
	//
	Flag *bool `json:"flag,omitempty"`
}
