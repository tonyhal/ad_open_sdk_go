/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceCreateByFileV10Request struct for QianchuanAudienceCreateByFileV10Request
type QianchuanAudienceCreateByFileV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	AudienceGroup string `json:"audience_group"`
	//
	AudienceName string `json:"audience_name"`
	//
	DataType int64 `json:"data_type"`
	//
	FileKey   string                                    `json:"file_key"`
	MatchType QianchuanAudienceCreateByFileV10MatchType `json:"match_type"`
}
