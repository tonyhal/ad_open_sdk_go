/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetComponentV2ResponseDataOrderComponentListInner struct for StarOrderGetComponentV2ResponseDataOrderComponentListInner
type StarOrderGetComponentV2ResponseDataOrderComponentListInner struct {
	//
	CocreateDouyinId *string `json:"cocreate_douyin_id,omitempty"`
	//
	IndustryComponentId *int64 `json:"industry_component_id,omitempty"`
	// 常规组件信息
	LinkComponentList []*StarOrderGetComponentV2ResponseDataOrderComponentListInnerLinkComponentListInner `json:"link_component_list,omitempty"`
	//
	LiveAttractComponentId *int64 `json:"live_attract_component_id,omitempty"`
	// 任务ID
	OrderId *int64 `json:"order_id,omitempty"`
	//
	SearchWord *string `json:"search_word,omitempty"`
}
