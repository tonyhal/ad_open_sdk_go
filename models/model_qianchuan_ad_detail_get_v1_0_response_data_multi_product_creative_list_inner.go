/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInner struct for QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInner
type QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInner struct {
	CreativeMaterialMode *QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeMaterialMode         `json:"creative_material_mode,omitempty"`
	CreativeSetting      *QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerCreativeSetting `json:"creative_setting,omitempty"`
	//
	ProductId            *int64                                                                                `json:"product_id,omitempty"`
	ProgrammaticCreative *QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreative `json:"programmatic_creative,omitempty"`
}