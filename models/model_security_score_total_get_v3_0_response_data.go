/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SecurityScoreTotalGetV30ResponseData
type SecurityScoreTotalGetV30ResponseData struct {
	PageInfo SecurityScoreTotalGetV30ResponseDataPageInfo `json:"page_info"`
	// 积分详情
	ScoreInfoList []*SecurityScoreTotalGetV30ResponseDataScoreInfoListInner `json:"score_info_list"`
}