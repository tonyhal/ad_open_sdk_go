/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalProjectDetailV30DataAudienceRegionLocationType
type LocalProjectDetailV30DataAudienceRegionLocationType string

// List of local_project_detail_v3.0_data_audience_region_location_type
const (
	ALL_LocalProjectDetailV30DataAudienceRegionLocationType     LocalProjectDetailV30DataAudienceRegionLocationType = "ALL"
	CURRENT_LocalProjectDetailV30DataAudienceRegionLocationType LocalProjectDetailV30DataAudienceRegionLocationType = "CURRENT"
	HOME_LocalProjectDetailV30DataAudienceRegionLocationType    LocalProjectDetailV30DataAudienceRegionLocationType = "HOME"
	TRAVEL_LocalProjectDetailV30DataAudienceRegionLocationType  LocalProjectDetailV30DataAudienceRegionLocationType = "TRAVEL"
)

// Ptr returns reference to local_project_detail_v3.0_data_audience_region_location_type value
func (v LocalProjectDetailV30DataAudienceRegionLocationType) Ptr() *LocalProjectDetailV30DataAudienceRegionLocationType {
	return &v
}
