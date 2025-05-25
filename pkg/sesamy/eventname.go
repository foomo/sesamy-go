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
	EventNameAddToWishlist        EventName = "add_to_wishlist"
	EventNameBeginCheckout        EventName = "begin_checkout"
	EventNameCampaignDetails      EventName = "campaign_details"
	EventNameClick                EventName = "click"
	EventNameCloseConvertLead     EventName = "close_convert_lead"
	EventNameCloseUnconvertLead   EventName = "close_unconvert_lead"
	EventNameDisqualifyLead       EventName = "disqualify_lead"
	EventNameEarnVirtualMoney     EventName = "earn_virtual_money"
	EventNameException            EventName = "exception"
	EventNameFileDownload         EventName = "file_download"
	EventNameFirstVisit           EventName = "first_visit"
	EventNameFormStart            EventName = "form_start"
	EventNameFormSubmit           EventName = "form_submit"
	EventNameGenerateLead         EventName = "generate_lead"
	EventNameJoinGroup            EventName = "join_group"
	EventNameLevelEnd             EventName = "level_end"
	EventNameLevelStart           EventName = "level_start"
	EventNameLevelUp              EventName = "level_up"
	EventNameLogin                EventName = "login"
	EventNameLogout               EventName = "logout"
	EventNamePageView             EventName = "page_view"
	EventNamePostScore            EventName = "post_score"
	EventNamePurchase             EventName = "purchase"
	EventNameQualifyLead          EventName = "qualify_lead"
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
	EventNameUnlockAchievement    EventName = "unlock_achievement"
	EventNameUserEngagement       EventName = "user_engagement"
	EventNameVideoComplete        EventName = "video_complete"
	EventNameVideoProgress        EventName = "video_progress"
	EventNameVideoStart           EventName = "video_start"
	EventNameViewCart             EventName = "view_cart"
	EventNameViewItem             EventName = "view_item"
	EventNameViewItemList         EventName = "view_item_list"
	EventNameViewPromotion        EventName = "view_promotion"
	EventNameViewSearchResults    EventName = "view_search_results"
	EventNameWorkingLead          EventName = "working_lead"
)

func (s EventName) String() string {
	return string(s)
}
