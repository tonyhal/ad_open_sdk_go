/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetComponentV2ResponseData
type StarOrderGetComponentV2ResponseData struct {
	// 任务组件信息
	OrderComponentList []*StarOrderGetComponentV2ResponseDataOrderComponentListInner `json:"order_component_list,omitempty"`
}
