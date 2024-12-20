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

type ResponseGetPliQuoteBaseResponse struct {
	Duration float32 `json:"duration,omitempty"`
	MaturityAge float32 `json:"maturity_age,omitempty"`
	MaturityAmount float32 `json:"maturity_amount,omitempty"`
	NetPremium float32 `json:"net_premium,omitempty"`
	Premium float32 `json:"premium,omitempty"`
	Rebate float32 `json:"rebate,omitempty"`
	Tax float32 `json:"tax,omitempty"`
	TotalBonus float32 `json:"total_bonus,omitempty"`
}
