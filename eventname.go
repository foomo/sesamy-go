package sesamy

type EventName string

// List of standard events
// - [[GA4] Automatically collected events](https://support.google.com/analytics/answer/9234069
// - [[GA4] Enhanced measurement events](https://support.google.com/analytics/answer/9216061)
// - [[GA4] Recommended events](https://support.google.com/analytics/answer/9267735)
const (
	EventNameAdImpression         EventName = "ad_impression"
	EventNameAddPaymentInfo       EventName = "add_payment_info"
	EventNameAddShippingInfo      EventName = "add_shipping_info"
	EventNameAddToCart            EventName = "add_to_cart"
	EventNameAddToWishlist        EventName = "add_to_wishlit"
	EventNameBeginCheckout        EventName = "begin_checkout"
	EventNameCampaignDetails      EventName = "campaign_details"
	EventNameClick                EventName = "click"
	EventNameEarnVirtualMoney     EventName = "earn_virtual_money"
	EventNameFileDownload         EventName = "file_download"
	EventNameFormStart            EventName = "form_start"
	EventNameFormSubmit           EventName = "form_submit"
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
	EventNameSessionStart         EventName = "session_start"
	EventNameShare                EventName = "share"
	EventNameSignUp               EventName = "sign_up"
	EventNameSpendVirtualCurrency EventName = "spend_virtual_currency"
	EventNameTutorialBegin        EventName = "tutorial_begin"
	EventNameTutorialComplete     EventName = "tutorial_complete"
	EventNameUnlockArchievement   EventName = "unlock_archievement"
	EventNameUserEngagement       EventName = "user_engagement"
	EventNameVideoComplete        EventName = "video_complete"
	EventNameVideoProgress        EventName = "video_progress"
	EventNameVideoStart           EventName = "video_start"
	EventNameViewCart             EventName = "view_cart"
	EventNameViewItem             EventName = "view_item"
	EventNameViewItemList         EventName = "view_item_list"
	EventNameViewPromotion        EventName = "view_promotion"
	EventNameViewSearchResults    EventName = "view_search_results"
)

func (s EventName) String() string {
	return string(s)
}
