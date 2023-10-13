/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionListV10ResponseDataAdListInnerAdInfo
type QianchuanUniPromotionListV10ResponseDataAdListInnerAdInfo struct {
	//
	Budget     *float64                                                `json:"budget,omitempty"`
	BudgetMode *QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode `json:"budget_mode,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	DeliverySeconds *int64 `json:"delivery_seconds,omitempty"`
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	Id            *int64                                                     `json:"id,omitempty"`
	MarketingGoal *QianchuanUniPromotionListV10DataAdListAdInfoMarketingGoal `json:"marketing_goal,omitempty"`
	//
	ModifyTime *string                                                `json:"modify_time,omitempty"`
	OptStatus  *QianchuanUniPromotionListV10DataAdListAdInfoOptStatus `json:"opt_status,omitempty"`
	//
	Roi2Goal *float64 `json:"roi2_goal,omitempty"`
	//
	StartTime *string                                             `json:"start_time,omitempty"`
	Status    *QianchuanUniPromotionListV10DataAdListAdInfoStatus `json:"status,omitempty"`
}
