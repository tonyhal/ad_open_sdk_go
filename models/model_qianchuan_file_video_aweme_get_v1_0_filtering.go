/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFileVideoAwemeGetV10Filtering
type QianchuanFileVideoAwemeGetV10Filtering struct {
	//
	AwemeItemIds []int64 `json:"aweme_item_ids,omitempty"`
	//
	AwemeItemUrl *string `json:"aweme_item_url,omitempty"`
	//
	MaterialIds []int64 `json:"material_ids,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
}
