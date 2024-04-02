/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AnalyticsAttributionV30RequestProperties
type AnalyticsAttributionV30RequestProperties struct {
	// 订单编号
	OrderId *string `json:"order_id,omitempty"`
	// 订单状态
	OrderState *string `json:"order_state,omitempty"`
	// 下单金额（付款金额，单位为“分”）如下单金额为13.23元，回传数值为1323 下单金额单位为“分”，请注意不要按照元回传⚠️
	PayAmount *int64 `json:"pay_amount,omitempty"`
	// 媒体位id，mm_1_2_3
	Pid *string `json:"pid,omitempty"`
	// 商品id，按照淘宝的sku_id进行回传。若订单中包含多个sku，则可在此字段中填写全部sku_id，英文逗号分隔
	ProductId *string `json:"product_id,omitempty"`
	// 商品图片
	ProductImag *string `json:"product_imag,omitempty"`
	// 商品名称
	ProductName *string `json:"product_name,omitempty"`
	// 商品数量
	ProductNumber *int64 `json:"product_number,omitempty"`
	// 商品单价，单位为“元”
	ProductPrice *int64 `json:"product_price,omitempty"`
	// 商品标题
	ProductTitle *string `json:"product_title,omitempty"`
	// 商品类型
	ProductType *string `json:"product_type,omitempty"`
	// 收货人所在的城市（若城市为直辖市仍然填写市，如北京市）；当phone_num_blurred值为*******1234类型时必填
	ReceiverCity *string `json:"receiver_city,omitempty"`
	// 收货人所在的省份；当phone_num_blurred值为*******1234类型时必填
	ReceiverProvince *string `json:"receiver_province,omitempty"`
	// 店铺名称
	ShopName *string `json:"shop_name,omitempty"`
}
