/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseListBulkCustomerMappingsRespone struct {
	// New Addition
	BatchName string `json:"batch_name,omitempty"`
	// New Addition
	BeatName string `json:"beat_name,omitempty"`
	BulkAddresseeAddress1 string `json:"bulk_addressee_address1,omitempty"`
	BulkAddresseeAddress2 string `json:"bulk_addressee_address2,omitempty"`
	BulkAddresseeAddress3 string `json:"bulk_addressee_address3,omitempty"`
	BulkAddresseeCity string `json:"bulk_addressee_city,omitempty"`
	BulkAddresseeId string `json:"bulk_addressee_id,omitempty"`
	BulkAddresseeName string `json:"bulk_addressee_name,omitempty"`
	BulkAddresseeState string `json:"bulk_addressee_state,omitempty"`
	// New Addition
	CreatedBy string `json:"created_by,omitempty"`
	// New Addition
	CreatedDate string `json:"created_date,omitempty"`
	Id int32 `json:"id,omitempty"`
	OfficeId int32 `json:"office_id,omitempty"`
}