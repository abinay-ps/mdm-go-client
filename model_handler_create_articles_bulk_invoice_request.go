/*
 * Postman Delivery Management
 *
 * This is Postman Delivery Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type HandlerCreateArticlesBulkInvoiceRequest struct {
	ArticleArray []string `json:"article_array"`
	BatchName string `json:"batch_name"`
	BeatName string `json:"beat_name"`
	BranchOffcieId int32 `json:"branch_offcie_id,omitempty"`
	BranchOfficeName string `json:"branch_office_name,omitempty"`
	BulkCustomerId string `json:"bulk_customer_id,omitempty"`
	BulkCustomerName string `json:"bulk_customer_name,omitempty"`
	CreatedBy string `json:"created_by"`
	OfficeId int32 `json:"office_id,omitempty"`
	PersonnelId int32 `json:"personnel_id"`
	PersonnelName string `json:"personnel_name"`
}
