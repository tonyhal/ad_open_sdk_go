/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductAvailablesV2ResponseDataListInner struct for DpaProductAvailablesV2ResponseDataListInner
type DpaProductAvailablesV2ResponseDataListInner struct {
	//
	Name *string `json:"name,omitempty"`
	//
	PlatformId      *int64                                         `json:"platform_id,omitempty"`
	ProductIndustry *DpaProductAvailablesV2DataListProductIndustry `json:"product_industry,omitempty"`
}
