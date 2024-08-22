/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionCreateV30RequestPromotionMaterials
type PromotionCreateV30RequestPromotionMaterials struct {
	//
	AdvancedDcSettings []*PromotionCreateV30PromotionMaterialsAdvancedDcSettings `json:"advanced_dc_settings,omitempty"`
	//
	AnchorMaterialList        []*PromotionCreateV30RequestPromotionMaterialsAnchorMaterialListInner `json:"anchor_material_list,omitempty"`
	BlueFlowMaterialRecommend *PromotionCreateV30PromotionMaterialsBlueFlowMaterialRecommend        `json:"blue_flow_material_recommend,omitempty"`
	//
	CallToActionButtons []string `json:"call_to_action_buttons,omitempty"`
	//
	CarouselMaterialList []*PromotionCreateV30RequestPromotionMaterialsCarouselMaterialListInner `json:"carousel_material_list,omitempty"`
	//
	ComponentMaterialList []*PromotionCreateV30RequestPromotionMaterialsComponentMaterialListInner `json:"component_material_list,omitempty"`
	DecorationMaterial    *PromotionCreateV30RequestPromotionMaterialsDecorationMaterial           `json:"decoration_material,omitempty"`
	DynamicCreativeSwitch *PromotionCreateV30PromotionMaterialsDynamicCreativeSwitch               `json:"dynamic_creative_switch,omitempty"`
	// 落地页链接字段选择，当params_type为DPA时必填
	ExternalUrlField *string `json:"external_url_field,omitempty"`
	//
	ExternalUrlMaterialList []string `json:"external_url_material_list,omitempty"`
	// 落地页检测参数
	ExternalUrlParams *string `json:"external_url_params,omitempty"`
	//
	ImageMaterialList     []*PromotionCreateV30RequestPromotionMaterialsImageMaterialListInner `json:"image_material_list,omitempty"`
	IntelligentGeneration *PromotionCreateV30PromotionMaterialsIntelligentGeneration           `json:"intelligent_generation,omitempty"`
	MiniProgramInfo       *PromotionCreateV30RequestPromotionMaterialsMiniProgramInfo          `json:"mini_program_info,omitempty"`
	//
	OpenUrl *string `json:"open_url,omitempty"`
	// 直达链接字段选择，当dpa_open_url_type为DPA必填
	OpenUrlField *string `json:"open_url_field,omitempty"`
	// 直达链接检测参数(DPA推广目的特有,在“产品库中提取的scheme地址\"后面追加填写的参数)
	OpenUrlParams *string                                          `json:"open_url_params,omitempty"`
	OpenUrlType   *PromotionCreateV30PromotionMaterialsOpenUrlType `json:"open_url_type,omitempty"`
	// 优选直达链接
	OpenUrls         []string                                              `json:"open_urls,omitempty"`
	OriginVideoTitle *PromotionCreateV30PromotionMaterialsOriginVideoTitle `json:"origin_video_title,omitempty"`
	ParamsType       *PromotionCreateV30PromotionMaterialsParamsType       `json:"params_type,omitempty"`
	//
	PlayableUrlMaterialList []string                                                `json:"playable_url_material_list,omitempty"`
	ProductInfo             *PromotionCreateV30RequestPromotionMaterialsProductInfo `json:"product_info,omitempty"`
	//
	TextAbstractList []*PromotionCreateV30RequestPromotionMaterialsTextAbstractListInner `json:"text_abstract_list,omitempty"`
	//
	TitleMaterialList []*PromotionCreateV30RequestPromotionMaterialsTitleMaterialListInner `json:"title_material_list,omitempty"`
	//
	Ulink *string `json:"ulink,omitempty"`
	//
	VideoMaterialList []*PromotionCreateV30RequestPromotionMaterialsVideoMaterialListInner `json:"video_material_list,omitempty"`
	//
	WebUrlMaterialList []string `json:"web_url_material_list,omitempty"`
}
