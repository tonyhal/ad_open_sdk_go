/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionCardRecommendTitleGetV2Response struct for ToolsPromotionCardRecommendTitleGetV2Response
type ToolsPromotionCardRecommendTitleGetV2Response struct {
	//
	Code *int64                                             `json:"code,omitempty"`
	Data *ToolsPromotionCardRecommendTitleGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
