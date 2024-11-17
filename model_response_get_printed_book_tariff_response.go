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

type ResponseGetPrintedBookTariffResponse struct {
	BaseTariff float32 `json:"base_tariff,omitempty"`
	Cgst float32 `json:"cgst,omitempty"`
	Gst float32 `json:"gst,omitempty"`
	Price float32 `json:"price,omitempty"`
	ProductCode string `json:"product_code,omitempty"`
	Sgst float32 `json:"sgst,omitempty"`
	TotalWithTax float32 `json:"total_with_tax,omitempty"`
	// Map to store VAS code and its corresponding price
	VasDetails map[string]float32 `json:"vas_details,omitempty"`
	Weight float32 `json:"weight,omitempty"`
}
