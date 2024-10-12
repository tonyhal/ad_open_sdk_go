/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// EnterpriseCommentListGetV10DataCommentListSource
type EnterpriseCommentListGetV10DataCommentListSource string

// List of enterprise_comment_list_get_v1.0_data_comment_list_source
const (
	FROM_DOUPLUS_EnterpriseCommentListGetV10DataCommentListSource EnterpriseCommentListGetV10DataCommentListSource = "FROM_DOUPLUS"
	FROM_PERFORM_EnterpriseCommentListGetV10DataCommentListSource EnterpriseCommentListGetV10DataCommentListSource = "FROM_PERFORM"
	FROM_OTHER_EnterpriseCommentListGetV10DataCommentListSource   EnterpriseCommentListGetV10DataCommentListSource = "FROM_OTHER"
	FROM_NATURAL_EnterpriseCommentListGetV10DataCommentListSource EnterpriseCommentListGetV10DataCommentListSource = "FROM_NATURAL"
	FROM_BRAND_EnterpriseCommentListGetV10DataCommentListSource   EnterpriseCommentListGetV10DataCommentListSource = "FROM_BRAND"
)

// Ptr returns reference to enterprise_comment_list_get_v1.0_data_comment_list_source value
func (v EnterpriseCommentListGetV10DataCommentListSource) Ptr() *EnterpriseCommentListGetV10DataCommentListSource {
	return &v
}
