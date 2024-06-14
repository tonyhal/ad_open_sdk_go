/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductListV2Filtering
type DpaClueProductListV2Filtering struct {
	// 商品审核状态
	AuditStatus []*DpaClueProductListV2FilteringAuditStatus `json:"audit_status,omitempty"`
	// 类目id，会级联查询所有叶子类目
	CategoryIds []int64 `json:"category_ids,omitempty"`
	// 类目名称，支持模糊搜索
	CategoryName *string `json:"category_name,omitempty"`
	// 商品必填字段完整性
	CompletionStatus []*DpaClueProductListV2FilteringCompletionStatus `json:"completion_status,omitempty"`
	// 商品ID或商品名称查询
	ProductIdOrNameSearch *string `json:"product_id_or_name_search,omitempty"`
	// 商品ID精确搜索
	ProductIds []int64 `json:"product_ids,omitempty"`
	// 商品名称模糊搜索
	ProductName *string `json:"product_name,omitempty"`
	// 商品权限关系过滤
	Rels []*DpaClueProductListV2FilteringRels `json:"rels,omitempty"`
	// 可投状态过滤
	Statuses []*DpaClueProductListV2FilteringStatuses `json:"statuses,omitempty"`
}
