/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudiencePushV10AccountType
type QianchuanAudiencePushV10AccountType string

// List of qianchuan_audience_push_v1.0_account_type
const (
	LOCAL_QianchuanAudiencePushV10AccountType    QianchuanAudiencePushV10AccountType = "LOCAL"
	NO_LOCAL_QianchuanAudiencePushV10AccountType QianchuanAudiencePushV10AccountType = "NO_LOCAL"
)

// Ptr returns reference to qianchuan_audience_push_v1.0_account_type value
func (v QianchuanAudiencePushV10AccountType) Ptr() *QianchuanAudiencePushV10AccountType {
	return &v
}
