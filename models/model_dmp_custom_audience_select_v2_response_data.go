/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpCustomAudienceSelectV2ResponseData
type DmpCustomAudienceSelectV2ResponseData struct {
	//
	CustomAudienceList []*DmpCustomAudienceSelectV2ResponseDataCustomAudienceListInner `json:"custom_audience_list,omitempty"`
	//
	Offset *int64 `json:"offset,omitempty"`
	//
	TotalNum *int64 `json:"total_num,omitempty"`
}
