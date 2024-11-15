/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseListArticlesDataEntriesResponse struct {
	ArticleAmount float32 `json:"article_amount,omitempty"`
	ArticleNumber string `json:"article_number,omitempty"`
	ArticleType string `json:"article_type,omitempty"`
	CodValue float32 `json:"cod_value,omitempty"`
	CustomDuty float32 `json:"custom_duty,omitempty"`
	CustomHandlingFee float32 `json:"custom_handling_fee,omitempty"`
	DataentryDate string `json:"dataentry_date,omitempty"`
	DemmurageCharges float32 `json:"demmurage_charges,omitempty"`
	InsuredValue float32 `json:"insured_value,omitempty"`
	IsCod bool `json:"is_cod,omitempty"`
	IsVpp bool `json:"is_vpp,omitempty"`
	OfficeId int32 `json:"office_id,omitempty"`
	PostDue float32 `json:"post_due,omitempty"`
	Reason string `json:"reason,omitempty"`
	ReceiverAddress1 string `json:"receiver_address1,omitempty"`
	ReceiverAddress2 string `json:"receiver_address2,omitempty"`
	ReceiverCity string `json:"receiver_city,omitempty"`
	ReceiverName string `json:"receiver_name,omitempty"`
	ReceiverPincode int32 `json:"receiver_pincode,omitempty"`
	RedirectionFee float32 `json:"redirection_fee,omitempty"`
	SenderAddress1 string `json:"sender_address1,omitempty"`
	SenderAddress2 string `json:"sender_address2,omitempty"`
	SenderCity string `json:"sender_city,omitempty"`
	SenderName string `json:"sender_name,omitempty"`
	SenderPincode int32 `json:"sender_pincode,omitempty"`
	TotalAmountToBeCollected float32 `json:"total_amount_to_be_collected,omitempty"`
	UserId int32 `json:"user_id,omitempty"`
	VppCommission float32 `json:"vpp_commission,omitempty"`
	VppValue float32 `json:"vpp_value,omitempty"`
	WarehouseCharges float32 `json:"warehouse_charges,omitempty"`
}
