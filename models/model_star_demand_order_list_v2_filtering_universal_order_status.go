/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOrderListV2FilteringUniversalOrderStatus
type StarDemandOrderListV2FilteringUniversalOrderStatus string

// List of star_demand_order_list_v2_filtering_universal_order_status
const (
	ALL_StarDemandOrderListV2FilteringUniversalOrderStatus           StarDemandOrderListV2FilteringUniversalOrderStatus = "ALL"
	CANCELED_StarDemandOrderListV2FilteringUniversalOrderStatus      StarDemandOrderListV2FilteringUniversalOrderStatus = "CANCELED"
	FINISHED_StarDemandOrderListV2FilteringUniversalOrderStatus      StarDemandOrderListV2FilteringUniversalOrderStatus = "FINISHED"
	ONGOING_StarDemandOrderListV2FilteringUniversalOrderStatus       StarDemandOrderListV2FilteringUniversalOrderStatus = "ONGOING"
	RECEIVEING_StarDemandOrderListV2FilteringUniversalOrderStatus    StarDemandOrderListV2FilteringUniversalOrderStatus = "RECEIVEING"
	WAIT_EVALUATE_StarDemandOrderListV2FilteringUniversalOrderStatus StarDemandOrderListV2FilteringUniversalOrderStatus = "WAIT_EVALUATE"
	WAIT_PAYMENT_StarDemandOrderListV2FilteringUniversalOrderStatus  StarDemandOrderListV2FilteringUniversalOrderStatus = "WAIT_PAYMENT"
)

// Ptr returns reference to star_demand_order_list_v2_filtering_universal_order_status value
func (v StarDemandOrderListV2FilteringUniversalOrderStatus) Ptr() *StarDemandOrderListV2FilteringUniversalOrderStatus {
	return &v
}
