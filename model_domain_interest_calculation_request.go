/*
 * Master Data Management
 *
 * This is Master Data Management API with Swagger documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DomainInterestCalculationRequest struct {
	AccountType string `json:"account_type,omitempty"`
	Amount float32 `json:"amount"`
	CbNsckvptype string `json:"cb_nsckvptype,omitempty"`
	CbTdtype int32 `json:"cb_tdtype,omitempty"`
	GirlAge int32 `json:"girl_age,omitempty"`
	ProductCode string `json:"product_code"`
}
