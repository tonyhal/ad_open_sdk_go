/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeCreateV30Request struct for BrandCreativeCreateV30Request
type BrandCreativeCreateV30Request struct {
	// 计划id
	AdId int64 `json:"ad_id"`
	// 广告主ID
	AdvertiserId        int64                                     `json:"advertiser_id"`
	CategoryInfo        BrandCreativeCreateV30RequestCategoryInfo `json:"category_info"`
	CreativeDisplayMode BrandCreativeCreateV30CreativeDisplayMode `json:"creative_display_mode"`
	// 创意列表
	CreativeList []*BrandCreativeCreateV30RequestCreativeListInner `json:"creative_list"`
	// 代理商员工id
	StaffId      *int64                                    `json:"staff_id,omitempty"`
	TrackUrlInfo BrandCreativeCreateV30RequestTrackUrlInfo `json:"track_url_info"`
}
