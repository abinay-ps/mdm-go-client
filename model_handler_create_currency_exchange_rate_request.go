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

type HandlerCreateCurrencyExchangeRateRequest struct {
	// User who created the entry (required)
	CreatedBy string `json:"created_by"`
	// Code of the currency (required)
	CurrencyCode string `json:"currency_code"`
	// Name of the currency (required)
	CurrencyName string `json:"currency_name"`
	// Exchange rate for exported goods (required)
	RateOfExchangeForExportGoods float32 `json:"rate_of_exchange_for_export_goods"`
	// Exchange rate for imported goods (required)
	RateOfExchangeForImportedGoods float32 `json:"rate_of_exchange_for_imported_goods"`
	// Schedule for currency exchange rates (required)
	ScheduleName string `json:"schedule_name"`
	// Effective date of the exchange rate (required)
	WithEffectFrom string `json:"with_effect_from"`
}