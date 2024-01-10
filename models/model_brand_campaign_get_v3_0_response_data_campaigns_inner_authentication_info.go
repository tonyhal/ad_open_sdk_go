/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignGetV30ResponseDataCampaignsInnerAuthenticationInfo 认证信息
type BrandCampaignGetV30ResponseDataCampaignsInnerAuthenticationInfo struct {
	// 附件信息
	Accessory []*BrandCampaignGetV30ResponseDataCampaignsInnerAuthenticationInfoAccessoryInner `json:"accessory,omitempty"`
	// 关联合同
	ProjectId   *string                                                                     `json:"project_id,omitempty"`
	ProjectInfo *BrandCampaignGetV30ResponseDataCampaignsInnerAuthenticationInfoProjectInfo `json:"project_info,omitempty"`
	// 内广签约主体ID
	SubjectId   *string                                                                     `json:"subject_id,omitempty"`
	SubjectInfo *BrandCampaignGetV30ResponseDataCampaignsInnerAuthenticationInfoSubjectInfo `json:"subject_info,omitempty"`
}
