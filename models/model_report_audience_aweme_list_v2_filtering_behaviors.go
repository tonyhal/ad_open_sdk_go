/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceAwemeListV2FilteringBehaviors
type ReportAudienceAwemeListV2FilteringBehaviors string

// List of report_audience_aweme_list_v2_filtering_behaviors
const (
	SHARED_USER_ReportAudienceAwemeListV2FilteringBehaviors          ReportAudienceAwemeListV2FilteringBehaviors = "SHARED_USER"
	LIVE_GOODS_CLICK_ReportAudienceAwemeListV2FilteringBehaviors     ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_ReportAudienceAwemeListV2FilteringBehaviors     ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_GOODS_ORDER"
	COMMENTED_USER_ReportAudienceAwemeListV2FilteringBehaviors       ReportAudienceAwemeListV2FilteringBehaviors = "COMMENTED_USER"
	GOODS_CARTS_ORDER_ReportAudienceAwemeListV2FilteringBehaviors    ReportAudienceAwemeListV2FilteringBehaviors = "GOODS_CARTS_ORDER"
	LIVE_WATCH_ReportAudienceAwemeListV2FilteringBehaviors           ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_WATCH"
	LIVE_EXCEPTIONAL_ReportAudienceAwemeListV2FilteringBehaviors     ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_COMMENT_ReportAudienceAwemeListV2FilteringBehaviors         ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_ReportAudienceAwemeListV2FilteringBehaviors ReportAudienceAwemeListV2FilteringBehaviors = "LIVE_EFFECTIVE_WATCH"
	FOLLOWED_USER_ReportAudienceAwemeListV2FilteringBehaviors        ReportAudienceAwemeListV2FilteringBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_ReportAudienceAwemeListV2FilteringBehaviors    ReportAudienceAwemeListV2FilteringBehaviors = "GOODS_CARTS_CLICK"
	LIKED_USER_ReportAudienceAwemeListV2FilteringBehaviors           ReportAudienceAwemeListV2FilteringBehaviors = "LIKED_USER"
)

// Ptr returns reference to report_audience_aweme_list_v2_filtering_behaviors value
func (v ReportAudienceAwemeListV2FilteringBehaviors) Ptr() *ReportAudienceAwemeListV2FilteringBehaviors {
	return &v
}
