/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CdpBrandGetV30ResponseDataBrandCategoryInfoInnerChildrenInner struct for CdpBrandGetV30ResponseDataBrandCategoryInfoInnerChildrenInner
type CdpBrandGetV30ResponseDataBrandCategoryInfoInnerChildrenInner struct {
	//
	Children []*CdpBrandGetV30ResponseDataBrandCategoryInfoInnerChildrenInnerChildrenInner `json:"children,omitempty"`
	// 二级类别id: yuntu_category_id
	Id *string `json:"id,omitempty"`
	// 二级类别标签
	Lable *string `json:"lable,omitempty"`
}
