/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentVideoFullScreen 全屏-视频
type BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentVideoFullScreen struct {
	// 背景图
	ImageInfoBkList []*BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentVideoFullScreenImageInfoBkListInner `json:"image_info_bk_list,omitempty"`
	// 视频信息
	VideoList []*BrandMaterialListV30ResponseDataMaterialsInnerMaterialComponentVideoFullScreenVideoListInner `json:"video_list,omitempty"`
}