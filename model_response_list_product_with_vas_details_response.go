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

type ResponseListProductWithVasDetailsResponse struct {
	Ack bool `json:"ack,omitempty"`
	AirMailSurchargeIndicator bool `json:"air_mail_surcharge_indicator,omitempty"`
	ApprovedBy string `json:"approved_by,omitempty"`
	ApprovedDate string `json:"approved_date,omitempty"`
	CashOnDeliveryIndicator bool `json:"cash_on_delivery_indicator,omitempty"`
	CgstPercentage float32 `json:"cgst_percentage,omitempty"`
	Colour string `json:"colour,omitempty"`
	CreatedBy string `json:"created_by,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	DeliveryAdviceRequired bool `json:"delivery_advice_required,omitempty"`
	DeliveryToAddresseInPersonRequired bool `json:"delivery_to_addresse_in_person_required,omitempty"`
	Denomination int32 `json:"denomination,omitempty"`
	DenominationType string `json:"denomination_type,omitempty"`
	DimensionInNonRoll int32 `json:"dimension_in_non_roll,omitempty"`
	DoorDeliveryCharges bool `json:"door_delivery_charges,omitempty"`
	EIod bool `json:"e_iod,omitempty"`
	EValuePayablePost bool `json:"e_value_payable_post,omitempty"`
	EnableToPortalFlag bool `json:"enable_to_portal_flag,omitempty"`
	Fac bool `json:"fac,omitempty"`
	GlCode string `json:"gl_code,omitempty"`
	IgstPercentage float32 `json:"igst_percentage,omitempty"`
	Insertion bool `json:"insertion,omitempty"`
	Insurance bool `json:"insurance,omitempty"`
	IsDeletedFlag bool `json:"is_deleted_flag,omitempty"`
	IsModifiedFlag bool `json:"is_modified_flag,omitempty"`
	IsStampFlag bool `json:"is_stamp_flag,omitempty"`
	LengthOfRollFormInCms int32 `json:"length_of_roll_form_in_cms,omitempty"`
	LengthPlus2TimesDiameterInRoll int32 `json:"length_plus_2_times_diameter_in_roll,omitempty"`
	Ltf bool `json:"ltf,omitempty"`
	MaxBreadthOfNonRollFormInCms int32 `json:"max_breadth_of_non_roll_form_in_cms,omitempty"`
	MaxHeightOfNonRollFormInCms int32 `json:"max_height_of_non_roll_form_in_cms,omitempty"`
	MaxLengthOfNonRollFormInCms int32 `json:"max_length_of_non_roll_form_in_cms,omitempty"`
	MaxQuantity int32 `json:"max_quantity,omitempty"`
	MaxWeightInGms int32 `json:"max_weight_in_gms,omitempty"`
	MinBreadthInNonRoll int32 `json:"min_breadth_in_non_roll,omitempty"`
	MinHeightInNonRoll int32 `json:"min_height_in_non_roll,omitempty"`
	MinLengthInNonRoll int32 `json:"min_length_in_non_roll,omitempty"`
	MinLengthInRoll int32 `json:"min_length_in_roll,omitempty"`
	MinQuantity int32 `json:"min_quantity,omitempty"`
	MinWeightInGms int32 `json:"min_weight_in_gms,omitempty"`
	PanIndiaAccessibilityFlag bool `json:"pan_india_accessibility_flag,omitempty"`
	PhilatelicBureauQuantity int32 `json:"philatelic_bureau_quantity,omitempty"`
	PremiumFlag bool `json:"premium_flag,omitempty"`
	Price float32 `json:"price,omitempty"`
	ProductCode string `json:"product_code,omitempty"`
	ProductGroup string `json:"product_group,omitempty"`
	ProductId int32 `json:"product_id,omitempty"`
	ProductMasterStatus string `json:"product_master_status,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	ProductSubgroup string `json:"product_subgroup,omitempty"`
	ProductTypeCode string `json:"product_type_code,omitempty"`
	ProductTypeId int32 `json:"product_type_id,omitempty"`
	ProofOfDelivery bool `json:"proof_of_delivery,omitempty"`
	ProponentsQuantity int32 `json:"proponents_quantity,omitempty"`
	Quantity int32 `json:"quantity,omitempty"`
	Reg bool `json:"reg,omitempty"`
	ServicableProductFlag string `json:"servicable_product_flag,omitempty"`
	ServiceTaxPercentage float32 `json:"service_tax_percentage,omitempty"`
	SgstPercentage float32 `json:"sgst_percentage,omitempty"`
	Sms bool `json:"sms,omitempty"`
	StampCategory string `json:"stamp_category,omitempty"`
	StampCategoryDiscription string `json:"stamp_category_discription,omitempty"`
	SurfaceAirLiftedIndicator bool `json:"surface_air_lifted_indicator,omitempty"`
	UnitOfMeasure string `json:"unit_of_measure,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	UpdatedDate string `json:"updated_date,omitempty"`
	ValidFrom string `json:"valid_from,omitempty"`
	ValidTo string `json:"valid_to,omitempty"`
	ValuePayablePostIndicator bool `json:"value_payable_post_indicator,omitempty"`
	Variant string `json:"variant,omitempty"`
	VasDetails []ResponseGetVasDetailsResponse `json:"vas_details,omitempty"`
}
