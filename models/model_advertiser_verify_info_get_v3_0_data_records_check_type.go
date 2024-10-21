/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserVerifyInfoGetV30DataRecordsCheckType
type AdvertiserVerifyInfoGetV30DataRecordsCheckType string

// List of advertiser_verify_info_get_v3.0_data_records_check_type
const (
	COMPANY_AdvertiserVerifyInfoGetV30DataRecordsCheckType         AdvertiserVerifyInfoGetV30DataRecordsCheckType = "COMPANY"
	GOVERNMENT_AdvertiserVerifyInfoGetV30DataRecordsCheckType      AdvertiserVerifyInfoGetV30DataRecordsCheckType = "GOVERNMENT"
	HK_MACAO_TAIWAN_AdvertiserVerifyInfoGetV30DataRecordsCheckType AdvertiserVerifyInfoGetV30DataRecordsCheckType = "HK_MACAO_TAIWAN"
	INDIVIDUAL_AdvertiserVerifyInfoGetV30DataRecordsCheckType      AdvertiserVerifyInfoGetV30DataRecordsCheckType = "INDIVIDUAL"
	OTHERS_AdvertiserVerifyInfoGetV30DataRecordsCheckType          AdvertiserVerifyInfoGetV30DataRecordsCheckType = "OTHERS"
	OVERSEA_AdvertiserVerifyInfoGetV30DataRecordsCheckType         AdvertiserVerifyInfoGetV30DataRecordsCheckType = "OVERSEA"
	SELF_EMPLOY_AdvertiserVerifyInfoGetV30DataRecordsCheckType     AdvertiserVerifyInfoGetV30DataRecordsCheckType = "SELF_EMPLOY"
)

// Ptr returns reference to advertiser_verify_info_get_v3.0_data_records_check_type value
func (v AdvertiserVerifyInfoGetV30DataRecordsCheckType) Ptr() *AdvertiserVerifyInfoGetV30DataRecordsCheckType {
	return &v
}