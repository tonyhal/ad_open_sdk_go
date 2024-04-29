/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportStardeliveryTaskDataGetV30ResponseDataListInner struct for ReportStardeliveryTaskDataGetV30ResponseDataListInner
type ReportStardeliveryTaskDataGetV30ResponseDataListInner struct {
	//
	Active *int64 `json:"active,omitempty"`
	//
	ActivePay *int64 `json:"active_pay,omitempty"`
	//
	ActiveRegister *int64 `json:"active_register,omitempty"`
	//
	GamePayCount *int64 `json:"game_pay_count,omitempty"`
	//
	InstallFinishCnt *int64 `json:"install_finish_cnt,omitempty"`
	//
	NonAdStatCost *float64 `json:"non_ad_stat_cost,omitempty"`
	//
	StarActiveCount *int64 `json:"star_active_count,omitempty"`
	//
	StarActivePayCount *int64 `json:"star_active_pay_count,omitempty"`
	//
	StarActiveRegisterCount *int64 `json:"star_active_register_count,omitempty"`
	//
	StarConvertCnt *int64 `json:"star_convert_cnt,omitempty"`
	//
	StarDeepPurchaseCount *int64 `json:"star_deep_purchase_count,omitempty"`
	//
	StarInstallFinishCount *int64 `json:"star_install_finish_count,omitempty"`
	//
	StarMaterialStatCost *float64 `json:"star_material_stat_cost,omitempty"`
	//
	StarSaleMaterialStatCost *float64 `json:"star_sale_material_stat_cost,omitempty"`
	//
	StarStatCost       *float64                                                    `json:"star_stat_cost,omitempty"`
	StarTaskAnchorType *ReportStardeliveryTaskDataGetV30DataListStarTaskAnchorType `json:"star_task_anchor_type,omitempty"`
	//
	StarTaskBudget *float64 `json:"star_task_budget,omitempty"`
	//
	StarTaskBudgetCostRate *float64 `json:"star_task_budget_cost_rate,omitempty"`
	//
	StarTaskCanDeliveryItemCount *int64                                                          `json:"star_task_can_delivery_item_count,omitempty"`
	StarTaskExternalAction       *ReportStardeliveryTaskDataGetV30DataListStarTaskExternalAction `json:"star_task_external_action,omitempty"`
	//
	StarTaskId *int64 `json:"star_task_id,omitempty"`
	//
	StarTaskName *string `json:"star_task_name,omitempty"`
	//
	StarTaskPostItemCount *int64 `json:"star_task_post_item_count,omitempty"`
	//
	StarTaskRelateProjectCount *int64                                                  `json:"star_task_relate_project_count,omitempty"`
	StarTaskSource             *ReportStardeliveryTaskDataGetV30DataListStarTaskSource `json:"star_task_source,omitempty"`
	//
	StarTaskStartTime *string                                                 `json:"star_task_start_time,omitempty"`
	StarTaskStatus    *ReportStardeliveryTaskDataGetV30DataListStarTaskStatus `json:"star_task_status,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
}