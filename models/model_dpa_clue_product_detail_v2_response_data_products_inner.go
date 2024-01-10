/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductDetailV2ResponseDataProductsInner struct for DpaClueProductDetailV2ResponseDataProductsInner
type DpaClueProductDetailV2ResponseDataProductsInner struct {
	// 定向年龄
	Age         []int64                                        `json:"age,omitempty"`
	AuditStatus *DpaClueProductDetailV2DataProductsAuditStatus `json:"audit_status,omitempty"`
	// 购买量
	Bought    *int64                                                    `json:"bought,omitempty"`
	BrandInfo *DpaClueProductDetailV2ResponseDataProductsInnerBrandInfo `json:"brand_info,omitempty"`
	Category  *DpaClueProductDetailV2ResponseDataProductsInnerCategory  `json:"category,omitempty"`
	// 定向城市
	City []string `json:"city,omitempty"`
	// 评论数
	Comments         *int64                                              `json:"comments,omitempty"`
	CompletionStatus *DpaClueProductDetailV2DataProductsCompletionStatus `json:"completion_status,omitempty"`
	// 商品描述信息
	Description *string `json:"description,omitempty"`
	// 特点
	Feature *string `json:"feature,omitempty"`
	// 是否存在视频;true:存在,false:不存在
	HasVideo *bool `json:"has_video,omitempty"`
	// 商品头图url
	ImageUrl *string `json:"image_url,omitempty"`
	// 商品组图
	ImagesUrl []*DpaClueProductDetailV2ResponseDataProductsInnerImagesUrlInner `json:"images_url,omitempty"`
	//
	Label       []string                                                    `json:"label,omitempty"`
	LandingInfo *DpaClueProductDetailV2ResponseDataProductsInnerLandingInfo `json:"landing_info,omitempty"`
	// 评分
	Mark *float64 `json:"mark,omitempty"`
	// 商品名称
	Name *string `json:"name,omitempty"`
	// 下线时间，精确到秒
	OfflineTime *string `json:"offline_time,omitempty"`
	// 上线时间，精确到秒
	OnlineTime *string `json:"online_time,omitempty"`
	// 自定义商品ID
	OuterId *string `json:"outer_id,omitempty"`
	// 商品poiID
	PoiId     *string                                                   `json:"poi_id,omitempty"`
	PriceInfo *DpaClueProductDetailV2ResponseDataProductsInnerPriceInfo `json:"price_info,omitempty"`
	// 商品ID
	ProductId *int64 `json:"product_id,omitempty"`
	// 行业特定字段
	Profession *map[string]string `json:"profession,omitempty"`
	// 定向省
	Province       []string                                                       `json:"province,omitempty"`
	ShopKeeperInfo *DpaClueProductDetailV2ResponseDataProductsInnerShopKeeperInfo `json:"shop_keeper_info,omitempty"`
	// 商品spuID
	SpuId  *string                                   `json:"spu_id,omitempty"`
	Status *DpaClueProductDetailV2DataProductsStatus `json:"status,omitempty"`
	// 商品标签
	Tags []string `json:"tags,omitempty"`
	// 商品标题
	Title *string `json:"title,omitempty"`
	// 商品视频链接
	VideoUrl *string `json:"video_url,omitempty"`
	// 商品视频列表
	Videos []*DpaClueProductDetailV2ResponseDataProductsInnerVideosInner `json:"videos,omitempty"`
}
