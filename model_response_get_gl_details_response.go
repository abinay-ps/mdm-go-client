/*
 * Master Data Management API
 *
 * A comprehensive API suite for managing various aspects of the postal system, including office management, product management, tariffs (domestic and international), product country mapping, local pin codes, country details, and currency exchange rates. Additionally, the API supports insurance quote generation for PLI/RPLI, and interest rate calculations for savings schemes such as SB, SSA, PPF, MIS, TD, NSC/KVP, SCSS, and MSS.
 *
 * API version: 1.0
 * Contact: support_cept@indiapost.gov.in
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseGetGlDetailsResponse struct {
	ApprovedBy string `json:"approved_by,omitempty"`
	ApprovedDate string `json:"approved_date,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	GlCode int32 `json:"gl_code,omitempty"`
	GlId int32 `json:"gl_id,omitempty"`
	GlName string `json:"gl_name,omitempty"`
	GlStatus string `json:"gl_status,omitempty"`
	Hoa string `json:"hoa,omitempty"`
	HoaName string `json:"hoa_name,omitempty"`
	OrderNumber string `json:"order_number,omitempty"`
	PaymentFlag bool `json:"payment_flag,omitempty"`
	ProductCode string `json:"product_code,omitempty"`
	ProductId int32 `json:"product_id,omitempty"`
	ReceiptFlag bool `json:"receipt_flag,omitempty"`
	Remarks string `json:"remarks,omitempty"`
	SupportingDocPath string `json:"supporting_doc_path,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	UpdatedDate string `json:"updated_date,omitempty"`
}
