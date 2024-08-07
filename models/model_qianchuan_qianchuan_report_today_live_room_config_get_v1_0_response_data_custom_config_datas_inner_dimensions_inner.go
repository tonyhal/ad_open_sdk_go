/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanQianchuanReportTodayLiveRoomConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInner struct for QianchuanQianchuanReportTodayLiveRoomConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInner
type QianchuanQianchuanReportTodayLiveRoomConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInner struct {
	//
	Description *string `json:"description,omitempty"`
	//
	ExclusionDims []string `json:"exclusion_dims,omitempty"`
	//
	ExclusionMetrics []string `json:"exclusion_metrics,omitempty"`
	//
	Field        *string                                                                                                         `json:"field,omitempty"`
	FilterConfig *QianchuanQianchuanReportTodayLiveRoomConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInnerFilterConfig `json:"filter_config,omitempty"`
	//
	Filterable *bool `json:"filterable,omitempty"`
	//
	IsRequired *bool `json:"is_required,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Sortable *bool `json:"sortable,omitempty"`
}
