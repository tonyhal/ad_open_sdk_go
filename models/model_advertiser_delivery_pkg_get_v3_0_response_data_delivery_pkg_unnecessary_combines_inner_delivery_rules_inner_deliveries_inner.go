/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInnerDeliveriesInner struct for AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInnerDeliveriesInner
type AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInnerDeliveriesInner struct {
	// 资质图片附件
	Attachments []*AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInnerDeliveryRulesInnerDeliveriesInnerAttachmentsInner `json:"attachments"`
	// 资质类型id
	QualType int64 `json:"qual_type"`
	// 资质类型名称
	QualTypeName string `json:"qual_type_name"`
	// 资质id
	QualificationId *int64 `json:"qualification_id,omitempty"`
	// 拒绝理由，若资质被拒绝，则会有拒绝理由
	RejectReason *string                                                                                    `json:"reject_reason,omitempty"`
	Status       AdvertiserDeliveryPkgGetV30DataDeliveryPkgUnnecessaryCombinesDeliveryRulesDeliveriesStatus `json:"status"`
}
