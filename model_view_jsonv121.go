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

type ViewJsonv121 struct {
	CanSeeTransactionStartDate           bool   `json:"can_see_transaction_start_date"`
	CanAddUrl                            bool   `json:"can_add_url"`
	CanAddWhereTag                       bool   `json:"can_add_where_tag"`
	CanSeeTransactionThisBankAccount     bool   `json:"can_see_transaction_this_bank_account"`
	CanSeeBankAccountOwners              bool   `json:"can_see_bank_account_owners"`
	CanSeePrivateAlias                   bool   `json:"can_see_private_alias"`
	CanEditOwnerComment                  bool   `json:"can_edit_owner_comment"`
	CanSeeOtherAccountNationalIdentifier bool   `json:"can_see_other_account_national_identifier"`
	CanSeePublicAlias                    bool   `json:"can_see_public_alias"`
	CanSeePhysicalLocation               bool   `json:"can_see_physical_location"`
	CanSeeOwnerComment                   bool   `json:"can_see_owner_comment"`
	CanSeeBankAccountIban                bool   `json:"can_see_bank_account_iban"`
	CanSeeCorporateLocation              bool   `json:"can_see_corporate_location"`
	CanSeeBankAccountNumber              bool   `json:"can_see_bank_account_number"`
	CanSeeOtherAccountBankName           bool   `json:"can_see_other_account_bank_name"`
	Description                          string `json:"description"`
	CanSeeTransactionOtherBankAccount    bool   `json:"can_see_transaction_other_bank_account"`
	CanDeleteCorporateLocation           bool   `json:"can_delete_corporate_location"`
	CanSeeComments                       bool   `json:"can_see_comments"`
	CanSeeBankAccountBankName            bool   `json:"can_see_bank_account_bank_name"`
	CanAddMoreInfo                       bool   `json:"can_add_more_info"`
	CanSeeOtherAccountNumber             bool   `json:"can_see_other_account_number"`
	CanSeeOtherAccountSwiftBic           bool   `json:"can_see_other_account_swift_bic"`
	CanAddOpenCorporatesUrl              bool   `json:"can_add_open_corporates_url"`
	CanSeeOtherAccountKind               bool   `json:"can_see_other_account_kind"`
	CanDeletePhysicalLocation            bool   `json:"can_delete_physical_location"`
	CanSeeBankAccountLabel               bool   `json:"can_see_bank_account_label"`
	CanSeeTransactionCurrency            bool   `json:"can_see_transaction_currency"`
	IsPublic                             bool   `json:"is_public"`
	CanSeeTransactionFinishDate          bool   `json:"can_see_transaction_finish_date"`
	CanAddTag                            bool   `json:"can_add_tag"`
	CanSeeImages                         bool   `json:"can_see_images"`
	CanSeeBankAccountCurrency            bool   `json:"can_see_bank_account_currency"`
	HideMetadataIfAliasUsed              bool   `json:"hide_metadata_if_alias_used"`
	CanDeleteWhereTag                    bool   `json:"can_delete_where_tag"`
	Alias                                string `json:"alias"`
	CanAddImageUrl                       bool   `json:"can_add_image_url"`
	CanAddComment                        bool   `json:"can_add_comment"`
	CanSeeImageUrl                       bool   `json:"can_see_image_url"`
	Id                                   string `json:"id"`
	CanSeeBankAccountNationalIdentifier  bool   `json:"can_see_bank_account_national_identifier"`
	CanSeeTags                           bool   `json:"can_see_tags"`
	CanSeeOpenCorporatesUrl              bool   `json:"can_see_open_corporates_url"`
	ShortName                            string `json:"short_name"`
	CanDeleteTag                         bool   `json:"can_delete_tag"`
	CanSeeMoreInfo                       bool   `json:"can_see_more_info"`
	CanSeeTransactionMetadata            bool   `json:"can_see_transaction_metadata"`
	CanDeleteComment                     bool   `json:"can_delete_comment"`
	CanSeeWhereTag                       bool   `json:"can_see_where_tag"`
	CanAddPrivateAlias                   bool   `json:"can_add_private_alias"`
	CanAddPublicAlias                    bool   `json:"can_add_public_alias"`
	CanSeeBankAccountSwiftBic            bool   `json:"can_see_bank_account_swift_bic"`
	CanAddImage                          bool   `json:"can_add_image"`
	CanSeeTransactionType                bool   `json:"can_see_transaction_type"`
	CanSeeOtherAccountIban               bool   `json:"can_see_other_account_iban"`
	CanAddPhysicalLocation               bool   `json:"can_add_physical_location"`
	CanAddCorporateLocation              bool   `json:"can_add_corporate_location"`
	CanDeleteImage                       bool   `json:"can_delete_image"`
	CanSeeUrl                            bool   `json:"can_see_url"`
	CanSeeBankAccountBalance             bool   `json:"can_see_bank_account_balance"`
	CanSeeTransactionBalance             bool   `json:"can_see_transaction_balance"`
	CanSeeTransactionAmount              bool   `json:"can_see_transaction_amount"`
	CanSeeOtherAccountMetadata           bool   `json:"can_see_other_account_metadata"`
	CanSeeBankAccountType                bool   `json:"can_see_bank_account_type"`
	CanSeeTransactionDescription         bool   `json:"can_see_transaction_description"`
}
