/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOrderListV2DataListUniversalOrderStatus
type StarDemandOrderListV2DataListUniversalOrderStatus string

// List of star_demand_order_list_v2_data_list_universal_order_status
const (
	ALL_StarDemandOrderListV2DataListUniversalOrderStatus           StarDemandOrderListV2DataListUniversalOrderStatus = "ALL"
	CANCELED_StarDemandOrderListV2DataListUniversalOrderStatus      StarDemandOrderListV2DataListUniversalOrderStatus = "CANCELED"
	FINISHED_StarDemandOrderListV2DataListUniversalOrderStatus      StarDemandOrderListV2DataListUniversalOrderStatus = "FINISHED"
	ONGOING_StarDemandOrderListV2DataListUniversalOrderStatus       StarDemandOrderListV2DataListUniversalOrderStatus = "ONGOING"
	RECEIVEING_StarDemandOrderListV2DataListUniversalOrderStatus    StarDemandOrderListV2DataListUniversalOrderStatus = "RECEIVEING"
	WAIT_EVALUATE_StarDemandOrderListV2DataListUniversalOrderStatus StarDemandOrderListV2DataListUniversalOrderStatus = "WAIT_EVALUATE"
	WAIT_PAYMENT_StarDemandOrderListV2DataListUniversalOrderStatus  StarDemandOrderListV2DataListUniversalOrderStatus = "WAIT_PAYMENT"
)

// Ptr returns reference to star_demand_order_list_v2_data_list_universal_order_status value
func (v StarDemandOrderListV2DataListUniversalOrderStatus) Ptr() *StarDemandOrderListV2DataListUniversalOrderStatus {
	return &v
}
