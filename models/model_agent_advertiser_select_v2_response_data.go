/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserSelectV2ResponseData
type AgentAdvertiserSelectV2ResponseData struct {
	//
	AccountSource *string `json:"account_source,omitempty"`
	//
	AdvertiserIds  []int64                                            `json:"advertiser_ids,omitempty"`
	CursorPageInfo *AgentAdvertiserSelectV2ResponseDataCursorPageInfo `json:"cursor_page_info,omitempty"`
	//
	List     []int64                                      `json:"list,omitempty"`
	PageInfo *AgentAdvertiserSelectV2ResponseDataPageInfo `json:"page_info,omitempty"`
}
