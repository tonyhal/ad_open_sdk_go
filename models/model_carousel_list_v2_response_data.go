/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselListV2ResponseData
type CarouselListV2ResponseData struct {
	//
	Carousels []*CarouselListV2ResponseDataCarouselsInner `json:"carousels,omitempty"`
	PageInfo  *CarouselListV2ResponseDataPageInfo         `json:"page_info,omitempty"`
}
