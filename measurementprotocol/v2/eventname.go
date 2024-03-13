package v2

type EventName string

const (
	EventNameAddPaymentInfo     EventName = "add_payment_info"
	EventNameAddShippingInfo    EventName = "add_shipping_info"
	EventNameAddToCart          EventName = "add_to_cart"
	EventNameAddToWishlist      EventName = "add_to_wishlit"
	EventNameBeginCheckout      EventName = "begin_checkout"
	EventNameEarnVirtualMoney   EventName = "earn_virtual_money"
	EventNameGenerateLead       EventName = "generate_lead"
	EventNameJoinGroup          EventName = "join_group"
	EventNameLevelEnd           EventName = "level_end"
	EventNameLevelStart         EventName = "level_start"
	EventNameLevelUp            EventName = "level_up"
	EventNameLogin              EventName = "login"
	EventNamePostScore          EventName = "post_score"
	EventNamePurchase           EventName = "purchase"
	EventNameRefund             EventName = "refund"
	EventNameRemoveFromCart     EventName = "remove_from_cart"
	EventNameSearch             EventName = "search"
	EventNameSelectContent      EventName = "select_content"
	EventNameSelectItem         EventName = "select_item"
	EventNameSelectPromotion    EventName = "select_promotion"
	EventNameShare              EventName = "share"
	EventNameSignUp             EventName = "sign_up"
	EventNameSpendVirtualMoney  EventName = "spend_virtual_money"
	EventNameTutorialBegin      EventName = "tutorial_begin"
	EventNameTuturialComplete   EventName = "tutorial_complete"
	EventNameUnlockArchievement EventName = "unlock_archievement"
	EventNameViewCart           EventName = "view_cart"
	EventNameViewItem           EventName = "view_item"
	EventNameViewItemList       EventName = "view_item_list"
	EventNameViewPromotion      EventName = "view_promotion"
)

func (s EventName) String() string {
	return string(s)
}
