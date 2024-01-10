/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductUpdateV2RequestProductInfo
type DpaProductUpdateV2RequestProductInfo struct {
	//
	Age []int64 `json:"age,omitempty"`
	//
	Bought    *int64                                         `json:"bought,omitempty"`
	BrandInfo *DpaProductUpdateV2RequestProductInfoBrandInfo `json:"brand_info,omitempty"`
	//
	City []string `json:"city,omitempty"`
	//
	Comments *int64 `json:"comments,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	Feature *string `json:"feature,omitempty"`
	//
	FirstCategory *string `json:"first_category,omitempty"`
	//
	FirstCategoryId *string                                  `json:"first_category_id,omitempty"`
	Geo             *DpaProductUpdateV2RequestProductInfoGeo `json:"geo,omitempty"`
	//
	ImageUrl string `json:"image_url"`
	//
	ImageUrls []string `json:"image_urls,omitempty"`
	//
	Label       []string                                         `json:"label,omitempty"`
	LandingInfo *DpaProductUpdateV2RequestProductInfoLandingInfo `json:"landing_info,omitempty"`
	//
	Mark *float64 `json:"mark,omitempty"`
	//
	Name string `json:"name"`
	//
	OfflineTime *int64 `json:"offline_time,omitempty"`
	//
	OnlineTime *int64 `json:"online_time,omitempty"`
	//
	OuterId   *string                                        `json:"outer_id,omitempty"`
	PriceInfo *DpaProductUpdateV2RequestProductInfoPriceInfo `json:"price_info,omitempty"`
	//
	Profession *map[string]string `json:"profession,omitempty"`
	//
	Province       []string                                            `json:"province,omitempty"`
	ShopKeeperInfo *DpaProductUpdateV2RequestProductInfoShopKeeperInfo `json:"shop_keeper_info,omitempty"`
	//
	SpuId  *string                             `json:"spu_id,omitempty"`
	Status DpaProductUpdateV2ProductInfoStatus `json:"status"`
	Stock  DpaProductUpdateV2ProductInfoStock  `json:"stock"`
	//
	SubCategory *string `json:"sub_category,omitempty"`
	//
	SubCategoryId *string `json:"sub_category_id,omitempty"`
	//
	Tags []string `json:"tags,omitempty"`
	//
	ThirdCategory *string `json:"third_category,omitempty"`
	//
	ThirdCategoryId *string `json:"third_category_id,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	Titles []string `json:"titles,omitempty"`
	//
	Video *string `json:"video,omitempty"`
	//
	Videos []*DpaProductCreateV2RequestProductInfoVideosInner `json:"videos,omitempty"`
}
