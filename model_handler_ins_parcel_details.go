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

type HandlerInsParcelDetails struct {
	AdditionalAmount float32 `json:"additionalAmount,omitempty"`
	AdditionalPrice float32 `json:"additionalPrice,omitempty"`
	BaseAmount float32 `json:"baseAmount,omitempty"`
	BasePrice float32 `json:"basePrice,omitempty"`
	MaximumAmount float32 `json:"maximumAmount,omitempty"`
}
