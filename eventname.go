package sesamy

type EventName string

const (
	EventNameAddPaymentInfo       EventName = "add_payment_info"
	EventNameAddShippingInfo      EventName = "add_shipping_info"
	EventNameAddToCart            EventName = "add_to_cart"
	EventNameAddToWishlist        EventName = "add_to_wishlit"
	EventNameAdImpression         EventName = "ad_impression"
	EventNameBeginCheckout        EventName = "begin_checkout"
	EventNameCampaignDetails      EventName = "campaign_details"
	EventNameClick                EventName = "click"
	EventNameEarnVirtualMoney     EventName = "earn_virtual_money"
	EventNameGenerateLead         EventName = "generate_lead"
	EventNameJoinGroup            EventName = "join_group"
	EventNameLevelEnd             EventName = "level_end"
	EventNameLevelStart           EventName = "level_start"
	EventNameLevelUp              EventName = "level_up"
	EventNameLogin                EventName = "login"
	EventNamePageView             EventName = "page_view"
	EventNamePostScore            EventName = "post_score"
	EventNamePurchase             EventName = "purchase"
	EventNameRefund               EventName = "refund"
	EventNameRemoveFromCart       EventName = "remove_from_cart"
	EventNameScreenView           EventName = "screen_view"
	EventNameScroll               EventName = "scroll"
	EventNameSearch               EventName = "search"
	EventNameSelectContent        EventName = "select_content"
	EventNameSelectItem           EventName = "select_item"
	EventNameSelectPromotion      EventName = "select_promotion"
	EventNameShare                EventName = "share"
	EventNameSignUp               EventName = "sign_up"
	EventNameSpendVirtualCurrency EventName = "spend_virtual_currency"
	EventNameTutorialBegin        EventName = "tutorial_begin"
	EventNameTutorialComplete     EventName = "tutorial_complete"
	EventNameUnlockArchievement   EventName = "unlock_archievement"
	EventNameViewCart             EventName = "view_cart"
	EventNameViewItem             EventName = "view_item"
	EventNameViewItemList         EventName = "view_item_list"
	EventNameViewPromotion        EventName = "view_promotion"
	EventNameViewSearchResults    EventName = "view_search_results"
)

func (s EventName) String() string {
	return string(s)
}
