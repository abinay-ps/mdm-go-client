/*
 * Master Data Management
 *
 * This is Master Data Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DomainModifyAccountCodeActiveDeactive struct {
	GlCode int32 `json:"gl_code,omitempty"`
	GlStatus string `json:"gl_status,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	OrderNumber string `json:"order_number,omitempty"`
	Remarks string `json:"remarks,omitempty"`
	SupportedDocPath string `json:"supported_doc_path,omitempty"`
}
