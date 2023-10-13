/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationSelectV2V2ResponseDataListInner struct for AdvertiserQualificationSelectV2V2ResponseDataListInner
type AdvertiserQualificationSelectV2V2ResponseDataListInner struct {
	//
	AuditTime string `json:"audit_time"`
	//
	ImageList []*AdvertiserQualificationSelectV2V2ResponseDataListInnerImageListInner `json:"image_list,omitempty"`
	//
	PicturePreviewUrl string `json:"picture_preview_url"`
	//
	QualificationId int64 `json:"qualification_id"`
	//
	QualificationType *int64 `json:"qualification_type,omitempty"`
	//
	RejectReason string `json:"reject_reason"`
	//
	Status int64 `json:"status"`
}
