/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetVideoV2ResponseDataOrderVideoListInner struct for StarOrderGetVideoV2ResponseDataOrderVideoListInner
type StarOrderGetVideoV2ResponseDataOrderVideoListInner struct {
	// 任务ID
	OrderId *int64 `json:"order_id,omitempty"`
	// 视频列表
	VideoList []*StarOrderGetVideoV2ResponseDataOrderVideoListInnerVideoListInner `json:"video_list,omitempty"`
}
