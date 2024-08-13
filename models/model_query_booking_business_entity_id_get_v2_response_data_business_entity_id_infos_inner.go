/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryBookingBusinessEntityIdGetV2ResponseDataBusinessEntityIdInfosInner struct for QueryBookingBusinessEntityIdGetV2ResponseDataBusinessEntityIdInfosInner
type QueryBookingBusinessEntityIdGetV2ResponseDataBusinessEntityIdInfosInner struct {
	// 业务录入实体ID
	BusinessEntityId *string `json:"business_entity_id,omitempty"`
	// 排期ID
	CartId *int64 `json:"cart_id,omitempty"`
	// 排期名称
	CartName *string `json:"cart_name,omitempty"`
	// 主订单ID
	OrderId *int64 `json:"order_id,omitempty"`
}
