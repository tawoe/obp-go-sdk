/*
 * Open Bank Project API
 *
 * An Open Source API for Banks. (c) TESOBE GmbH. 2011 - 2024. Licensed under the AGPL and commercial licences.
 *
 * API version: v5.1.0
 * Contact: contact@tesobe.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package obp_golang

type ViewJsonV300 struct {
	CanSeeTransactionStartDate           bool   `json:"can_see_transaction_start_date"`
	CanAddUrl                            bool   `json:"can_add_url"`
	CanAddWhereTag                       bool   `json:"can_add_where_tag"`
	CanSeeTransactionThisBankAccount     bool   `json:"can_see_transaction_this_bank_account"`
	CanSeeBankAccountOwners              bool   `json:"can_see_bank_account_owners"`
	CanSeeBankRoutingAddress             bool   `json:"can_see_bank_routing_address"`
	CanSeePrivateAlias                   bool   `json:"can_see_private_alias"`
	CanEditOwnerComment                  bool   `json:"can_edit_owner_comment"`
	IsSystem                             bool   `json:"is_system"`
	CanSeeOtherAccountNationalIdentifier bool   `json:"can_see_other_account_national_identifier"`
	CanSeeBankRoutingScheme              bool   `json:"can_see_bank_routing_scheme"`
	CanSeePublicAlias                    bool   `json:"can_see_public_alias"`
	CanSeePhysicalLocation               bool   `json:"can_see_physical_location"`
	CanSeeOwnerComment                   bool   `json:"can_see_owner_comment"`
	CanSeeBankAccountIban                bool   `json:"can_see_bank_account_iban"`
	CanSeeCorporateLocation              bool   `json:"can_see_corporate_location"`
	CanSeeBankAccountNumber              bool   `json:"can_see_bank_account_number"`
	CanSeeOtherAccountBankName           bool   `json:"can_see_other_account_bank_name"`
	Description                          string `json:"description"`
	CanSeeBankAccountRoutingScheme       bool   `json:"can_see_bank_account_routing_scheme"`
	CanSeeTransactionOtherBankAccount    bool   `json:"can_see_transaction_other_bank_account"`
	CanDeleteCorporateLocation           bool   `json:"can_delete_corporate_location"`
	CanSeeComments                       bool   `json:"can_see_comments"`
	CanSeeBankAccountBankName            bool   `json:"can_see_bank_account_bank_name"`
	CanAddMoreInfo                       bool   `json:"can_add_more_info"`
	CanCreateDirectDebit                 bool   `json:"can_create_direct_debit"`
	CanSeeOtherAccountNumber             bool   `json:"can_see_other_account_number"`
	CanSeeOtherAccountSwiftBic           bool   `json:"can_see_other_account_swift_bic"`
	CanAddOpenCorporatesUrl              bool   `json:"can_add_open_corporates_url"`
	CanSeeOtherAccountKind               bool   `json:"can_see_other_account_kind"`
	CanAddTransactionRequestToOwnAccount bool   `json:"can_add_transaction_request_to_own_account"`
	CanDeletePhysicalLocation            bool   `json:"can_delete_physical_location"`
	CanSeeBankAccountLabel               bool   `json:"can_see_bank_account_label"`
	CanSeeTransactionCurrency            bool   `json:"can_see_transaction_currency"`
	IsPublic                             bool   `json:"is_public"`
	CanSeeTransactionFinishDate          bool   `json:"can_see_transaction_finish_date"`
	CanSeeBankAccountRoutingAddress      bool   `json:"can_see_bank_account_routing_address"`
	CanAddTag                            bool   `json:"can_add_tag"`
	CanSeeImages                         bool   `json:"can_see_images"`
	CanQueryAvailableFunds               bool   `json:"can_query_available_funds"`
	CanSeeBankAccountCreditLimit         bool   `json:"can_see_bank_account_credit_limit"`
	CanSeeBankAccountCurrency            bool   `json:"can_see_bank_account_currency"`
	HideMetadataIfAliasUsed              bool   `json:"hide_metadata_if_alias_used"`
	CanDeleteWhereTag                    bool   `json:"can_delete_where_tag"`
	Alias                                string `json:"alias"`
	CanAddImageUrl                       bool   `json:"can_add_image_url"`
	CanAddComment                        bool   `json:"can_add_comment"`
	CanSeeImageUrl                       bool   `json:"can_see_image_url"`
	Id                                   string `json:"id"`
	CanCreateStandingOrder               bool   `json:"can_create_standing_order"`
	CanSeeBankAccountNationalIdentifier  bool   `json:"can_see_bank_account_national_identifier"`
	CanAddCounterparty                   bool   `json:"can_add_counterparty"`
	CanAddTransactionRequestToAnyAccount bool   `json:"can_add_transaction_request_to_any_account"`
	CanSeeTags                           bool   `json:"can_see_tags"`
	CanSeeOpenCorporatesUrl              bool   `json:"can_see_open_corporates_url"`
	ShortName                            string `json:"short_name"`
	CanDeleteTag                         bool   `json:"can_delete_tag"`
	CanSeeOtherAccountRoutingScheme      bool   `json:"can_see_other_account_routing_scheme"`
	CanSeeMoreInfo                       bool   `json:"can_see_more_info"`
	CanSeeTransactionMetadata            bool   `json:"can_see_transaction_metadata"`
	CanDeleteComment                     bool   `json:"can_delete_comment"`
	CanSeeWhereTag                       bool   `json:"can_see_where_tag"`
	CanAddPrivateAlias                   bool   `json:"can_add_private_alias"`
	IsFirehose                           bool   `json:"is_firehose,omitempty"`
	CanAddPublicAlias                    bool   `json:"can_add_public_alias"`
	CanSeeBankAccountSwiftBic            bool   `json:"can_see_bank_account_swift_bic"`
	CanAddImage                          bool   `json:"can_add_image"`
	CanSeeTransactionType                bool   `json:"can_see_transaction_type"`
	CanSeeOtherAccountRoutingAddress     bool   `json:"can_see_other_account_routing_address"`
	CanSeeOtherAccountIban               bool   `json:"can_see_other_account_iban"`
	CanAddPhysicalLocation               bool   `json:"can_add_physical_location"`
	CanAddCorporateLocation              bool   `json:"can_add_corporate_location"`
	CanDeleteImage                       bool   `json:"can_delete_image"`
	CanSeeUrl                            bool   `json:"can_see_url"`
	CanSeeBankAccountBalance             bool   `json:"can_see_bank_account_balance"`
	CanSeeOtherBankRoutingAddress        bool   `json:"can_see_other_bank_routing_address"`
	CanSeeTransactionBalance             bool   `json:"can_see_transaction_balance"`
	MetadataView                         string `json:"metadata_view"`
	CanSeeTransactionAmount              bool   `json:"can_see_transaction_amount"`
	CanSeeOtherAccountMetadata           bool   `json:"can_see_other_account_metadata"`
	CanSeeBankAccountType                bool   `json:"can_see_bank_account_type"`
	CanSeeOtherBankRoutingScheme         bool   `json:"can_see_other_bank_routing_scheme"`
	CanSeeTransactionDescription         bool   `json:"can_see_transaction_description"`
}
