/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportLiveRoomProductGetV2ResponseDataListInner struct for ReportLiveRoomProductGetV2ResponseDataListInner
type ReportLiveRoomProductGetV2ResponseDataListInner struct {
	//
	Fields map[string]interface{} `json:"fields,omitempty"`
	//
	ProductCategory *string `json:"product_category,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
	//
	ProductName *string `json:"product_name,omitempty"`
	//
	ProductPicture *string `json:"product_picture,omitempty"`
	//
	ProductUrl *string `json:"product_url,omitempty"`
}
