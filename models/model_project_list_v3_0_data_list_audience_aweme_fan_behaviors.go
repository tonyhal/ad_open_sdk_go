/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectListV30DataListAudienceAwemeFanBehaviors
type ProjectListV30DataListAudienceAwemeFanBehaviors string

// List of project_list_v3.0_data_list_audience_aweme_fan_behaviors
const (
	COMMENTED_USER_ProjectListV30DataListAudienceAwemeFanBehaviors       ProjectListV30DataListAudienceAwemeFanBehaviors = "COMMENTED_USER"
	FOLLOWED_USER_ProjectListV30DataListAudienceAwemeFanBehaviors        ProjectListV30DataListAudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_ProjectListV30DataListAudienceAwemeFanBehaviors    ProjectListV30DataListAudienceAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_ProjectListV30DataListAudienceAwemeFanBehaviors    ProjectListV30DataListAudienceAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_ProjectListV30DataListAudienceAwemeFanBehaviors           ProjectListV30DataListAudienceAwemeFanBehaviors = "LIKED_USER"
	LIVE_COMMENT_ProjectListV30DataListAudienceAwemeFanBehaviors         ProjectListV30DataListAudienceAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_ProjectListV30DataListAudienceAwemeFanBehaviors ProjectListV30DataListAudienceAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_ProjectListV30DataListAudienceAwemeFanBehaviors     ProjectListV30DataListAudienceAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_ProjectListV30DataListAudienceAwemeFanBehaviors     ProjectListV30DataListAudienceAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_ProjectListV30DataListAudienceAwemeFanBehaviors     ProjectListV30DataListAudienceAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_ProjectListV30DataListAudienceAwemeFanBehaviors           ProjectListV30DataListAudienceAwemeFanBehaviors = "LIVE_WATCH"
	SHARED_USER_ProjectListV30DataListAudienceAwemeFanBehaviors          ProjectListV30DataListAudienceAwemeFanBehaviors = "SHARED_USER"
)

// All allowed values of ProjectListV30DataListAudienceAwemeFanBehaviors enum
var AllowedProjectListV30DataListAudienceAwemeFanBehaviorsEnumValues = []ProjectListV30DataListAudienceAwemeFanBehaviors{
	"COMMENTED_USER",
	"FOLLOWED_USER",
	"GOODS_CARTS_CLICK",
	"GOODS_CARTS_ORDER",
	"LIKED_USER",
	"LIVE_COMMENT",
	"LIVE_EFFECTIVE_WATCH",
	"LIVE_EXCEPTIONAL",
	"LIVE_GOODS_CLICK",
	"LIVE_GOODS_ORDER",
	"LIVE_WATCH",
	"SHARED_USER",
}

// NewProjectListV30DataListAudienceAwemeFanBehaviorsFromValue returns a pointer to a valid ProjectListV30DataListAudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30DataListAudienceAwemeFanBehaviorsFromValue(v string) (*ProjectListV30DataListAudienceAwemeFanBehaviors, error) {
	ev := ProjectListV30DataListAudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30DataListAudienceAwemeFanBehaviors: valid values are %v", v, AllowedProjectListV30DataListAudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30DataListAudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedProjectListV30DataListAudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_data_list_audience_aweme_fan_behaviors value
func (v ProjectListV30DataListAudienceAwemeFanBehaviors) Ptr() *ProjectListV30DataListAudienceAwemeFanBehaviors {
	return &v
}
