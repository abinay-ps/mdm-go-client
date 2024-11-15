/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseCreateEmOsBoDeliverySlipGenerationResponse struct {
	BookingDate string `json:"booking_date,omitempty"`
	DestinationOfficeId int32 `json:"destination_office_id,omitempty"`
	DestinationOfficeName string `json:"destination_office_name,omitempty"`
	DestinationPincode int32 `json:"destination_pincode,omitempty"`
	EmoMessage string `json:"emo_message,omitempty"`
	EmoNumber string `json:"emo_number,omitempty"`
	EmoValue float32 `json:"emo_value,omitempty"`
	IsInvoiced bool `json:"is_invoiced,omitempty"`
	IsPrinted bool `json:"is_printed,omitempty"`
	ReceiverAddress1 string `json:"receiver_address1,omitempty"`
	ReceiverAddress2 string `json:"receiver_address2,omitempty"`
	ReceiverAddress3 string `json:"receiver_address3,omitempty"`
	ReceiverCity string `json:"receiver_city,omitempty"`
	ReceiverName string `json:"receiver_name,omitempty"`
	ReceiverState string `json:"receiver_state,omitempty"`
	SenderAddress1 string `json:"sender_address1,omitempty"`
	SenderAddress2 string `json:"sender_address2,omitempty"`
	SenderAddress3 string `json:"sender_address3,omitempty"`
	SenderCity string `json:"sender_city,omitempty"`
	SenderName string `json:"sender_name,omitempty"`
	SenderPincode int32 `json:"sender_pincode,omitempty"`
	SenderState string `json:"sender_state,omitempty"`
	SourceOfficeId int32 `json:"source_office_id,omitempty"`
	SourceOfficeName string `json:"source_office_name,omitempty"`
	SourcePincode int32 `json:"source_pincode,omitempty"`
}