/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfo 素材信息
type QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfo struct {
	ImageMaterial *QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoImageMaterial `json:"image_material,omitempty"`
	MaterialType  QianchuanAdMaterialGetV10DataAdMaterialInfosMaterialInfoMaterialType                `json:"material_type"`
	RoomMaterial  *QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoRoomMaterial  `json:"room_material,omitempty"`
	TitleMaterial *QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoTitleMaterial `json:"title_material,omitempty"`
	VideoMaterial *QianchuanAdMaterialGetV10ResponseDataAdMaterialInfosInnerMaterialInfoVideoMaterial `json:"video_material,omitempty"`
}