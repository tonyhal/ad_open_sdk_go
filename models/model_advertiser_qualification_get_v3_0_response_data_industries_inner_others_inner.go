/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationGetV30ResponseDataIndustriesInnerOthersInner struct for AdvertiserQualificationGetV30ResponseDataIndustriesInnerOthersInner
type AdvertiserQualificationGetV30ResponseDataIndustriesInnerOthersInner struct {
	// 资质图片附件id
	AttachmentId *string `json:"attachment_id,omitempty"`
	// 资质图片地址
	PictureUrl *string `json:"picture_url,omitempty"`
	// 资质id
	QualificationId *int64 `json:"qualification_id,omitempty"`
	// 拒绝理由
	RejectReason *string                                                  `json:"reject_reason,omitempty"`
	Status       *AdvertiserQualificationGetV30DataIndustriesOthersStatus `json:"status,omitempty"`
}
