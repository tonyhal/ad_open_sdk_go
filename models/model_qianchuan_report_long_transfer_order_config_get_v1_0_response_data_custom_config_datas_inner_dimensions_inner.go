/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportLongTransferOrderConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInner struct for QianchuanReportLongTransferOrderConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInner
type QianchuanReportLongTransferOrderConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInner struct {
	//
	Description *string `json:"description,omitempty"`
	//
	ExclusionDims []string `json:"exclusion_dims,omitempty"`
	//
	ExclusionMetrics []string `json:"exclusion_metrics,omitempty"`
	//
	Field        *string                                                                                                    `json:"field,omitempty"`
	FilterConfig *QianchuanReportLongTransferOrderConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInnerFilterConfig `json:"filter_config,omitempty"`
	//
	Filterable *bool `json:"filterable,omitempty"`
	//
	IsRequired *bool `json:"is_required,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Sortable *bool `json:"sortable,omitempty"`
}
