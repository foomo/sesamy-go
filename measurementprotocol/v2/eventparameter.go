package v2

// EventParameter as string
// See https://support.google.com/analytics/table/13594742?sjid=7861230991468479976-EU
type EventParameter string

const (
	// EventParameterAppVersion The mobile app's versionName (Android) or short bundle version (iOS) in which an event occurred
	EventParameterAppVersion EventParameter = "app_version"
	// EventParameterCancellationReason The reason a cancellation occurred
	EventParameterCancellationReason EventParameter = "cancellation_reason"
	// EventParameterFirebaseErrorValue Additional details corresponding to the error code reported by the Firebase SDK
	EventParameterFirebaseErrorValue EventParameter = "firebase_error_value"
	// EventParameterFirebaseScreen The current screen’s name, as provided by the developer
	EventParameterFirebaseScreen EventParameter = "firebase_screen"
	// EventParameterFirebaseScreenClass The current screen’s class name
	EventParameterFirebaseScreenClass EventParameter = "firebase_screen_class"
	// EventParameterFirebaseScreenID A random identifier for the current screen
	EventParameterFirebaseScreenID EventParameter = "firebase_screen_id"
	// EventParameterMessageID The Firebase Cloud Messaging or Firebase In-App Messaging message identifier, which is unique per message campaign
	EventParameterMessageID EventParameter = "message_id"
	// EventParameterMessageName The Firebase Cloud Messaging or Firebase In-App Messaging message name
	EventParameterMessageName EventParameter = "message_name"
	// EventParameterMessageType The Firebase Cloud Messaging message notification type
	EventParameterMessageType EventParameter = "message_type"
	// EventParameterPreviousAppVersion For the app_update event, the parameter signifies the previous application version
	EventParameterPreviousAppVersion EventParameter = "previous_app_version"
	// EventParameterPreviousOsVersion For the os_update event, the parameter signifies the previous OS version
	EventParameterPreviousOsVersion EventParameter = "previous_os_version"
	// EventParameterCampaignContent The ad content that was associated with the start of a session
	EventParameterCampaignContent EventParameter = "campaign_content"
	// EventParameterCampaignMedium The method for acquiring users to your website or application
	EventParameterCampaignMedium EventParameter = "campaign_medium"
	// EventParameterCampaignSource A representation of the publisher or inventory source from which traffic originated. For example, users who return to your website from Google Search show as "google" in the Session source dimension
	EventParameterCampaignSource EventParameter = "campaign_source"
	// EventParameterCampaignTerm The term that was associated with the start of a session
	EventParameterCampaignTerm EventParameter = "campaign_term"
	// EventParameterCoupon The coupon name or code associated with an event
	EventParameterCoupon EventParameter = "coupon"
	// EventParameterCurrency The currency used in an event, in 3-letter ISO 4217 format. For example, the currency used in a purchase
	EventParameterCurrency EventParameter = "currency"
	// EventParameterShippingTier The shipping tier selected for delivery of a purchased item
	EventParameterShippingTier EventParameter = "shipping_tier"
	// EventParameterTransactionID The unique identifier of a transaction
	EventParameterTransactionID EventParameter = "transaction_id"
	// EventParameterValue The monetary value of the event
	EventParameterValue EventParameter = "value"
	// EventParameterAffiliation A product affiliation to designate a supplying company or brick and mortar store location
	EventParameterAffiliation EventParameter = "affiliation"
	// EventParameterCreativeName The name of a creative used in a promotion. Example value: summer_banner
	EventParameterCreativeName EventParameter = "creative_name"
	// EventParameterCreativeSlot The name of the promotional creative slot associated with an event. Example value: featured_app_1
	EventParameterCreativeSlot EventParameter = "creative_slot"
	// EventParameterDiscount The value of a discount value associated with a purchased item
	EventParameterDiscount EventParameter = "discount"
	// EventParameterIndex The index of the item in a list
	EventParameterIndex EventParameter = "index"
	// EventParameterItemBrand The brand of an item
	EventParameterItemBrand EventParameter = "item_brand"
	// EventParameterItemCategory The category of an item. If used as part of a category hierarchy or taxonomy, then this is the first category
	EventParameterItemCategory EventParameter = "item_category"
	// EventParameterItemCategory2 The second hierarchical category in which you classified an item
	EventParameterItemCategory2 EventParameter = "item_category2"
	// EventParameterItemCategory3 The third hierarchical category in which you classified an item
	EventParameterItemCategory3 EventParameter = "item_category3"
	// EventParameterItemCategory4 The fourth hierarchical category in which you classified an item
	EventParameterItemCategory4 EventParameter = "item_category4"
	// EventParameterItemCategory5 The fifth hierarchical category in which you classified an item
	EventParameterItemCategory5 EventParameter = "item_category5"
	// EventParameterItemID The ID that you specify for an item
	EventParameterItemID EventParameter = "item_id"
	// EventParameterItemListID The name of the list in which an item was presented to a user
	EventParameterItemListID EventParameter = "item_list_id"
	// EventParameterItemListName The ID of the list in which an item was presented to a user
	EventParameterItemListName EventParameter = "item_list_name"
	// EventParameterItemName The name of the event that contains the parameter group
	EventParameterItemName EventParameter = "item_name"
	// EventParameterItemVariant The item variant or unique code or description (e.g., XS, S, M, L for size; Red, Blue, Green, Black for color) for additional item details or options
	EventParameterItemVariant EventParameter = "item_variant"
	// EventParameterLocationID The physical location associated with the item (e.g. the physical store location)
	EventParameterLocationID EventParameter = "location_id"
	// EventParameterPromotionID The ID of the promotion associated with an event
	EventParameterPromotionID EventParameter = "promotion_id"
	// EventParameterPromotionName The name of the promotion associated with an event
	EventParameterPromotionName EventParameter = "promotion_name"
	// EventParameterAchievementID The ID of an achievement that was unlocked in a game
	EventParameterAchievementID EventParameter = "achievement_id"
	// EventParameterCharacter The name of a character in a game
	EventParameterCharacter EventParameter = "character"
	// EventParameterLevelName The name of the level in a game
	EventParameterLevelName EventParameter = "level_name"
	// EventParameterVirtualCurrencyName The name of a virtual currency
	EventParameterVirtualCurrencyName EventParameter = "virtual_currency_name"
	// EventParameterFileExtension The extension of a file download
	EventParameterFileExtension EventParameter = "file_extension"
	// EventParameterFileName The page path of a file download
	EventParameterFileName EventParameter = "file_name"
	// EventParameterFormDestination The URL to which a form is being submitted
	EventParameterFormDestination EventParameter = "form_destination"
	// EventParameterFormID The HTML id attribution of the <form> DOM element
	EventParameterFormID EventParameter = "form_id"
	// EventParameterFormName The HTML name attribute of the <form> DOM element
	EventParameterFormName EventParameter = "form_name"
	// EventParameterFormSubmitText The text of the submit button, if present
	EventParameterFormSubmitText EventParameter = "form_submit_text"
	// EventParameterGroupID The ID of a group
	EventParameterGroupID EventParameter = "group_id"
	// EventParameterLanguage The language setting of a user’s browser or device, displayed as the ISO 639 language code
	EventParameterLanguage EventParameter = "language"
	// EventParameterPercentScrolled The percentage down the page that the user scrolled
	EventParameterPercentScrolled EventParameter = "percent_scrolled"
	// EventParameterSearchTerm The strings or keywords used in a search
	EventParameterSearchTerm EventParameter = "search_term"
	// EventParameterLinkClasses The HTML class attribute for an outbound link or file download
	EventParameterLinkClasses EventParameter = "link_classes"
	// EventParameterLinkDomain The destination domain of an outbound link or file download
	EventParameterLinkDomain EventParameter = "link_domain"
	// EventParameterLinkID The ID for an outbound link or file download
	EventParameterLinkID EventParameter = "link_id"
	// EventParameterLinkUrl The full URL for an outbound link or file download
	EventParameterLinkUrl EventParameter = "link_url"
	// EventParameterOutbound Indicates whether a click was on an outbound link
	EventParameterOutbound EventParameter = "outbound"
	// EventParameterContentGroup The content group associated with a page or screen
	EventParameterContentGroup EventParameter = "content_group"
	// EventParameterContentID An ID for an article of content that a user interacted with
	EventParameterContentID EventParameter = "content_id"
	// EventParameterContentType The type of content that a user interacted with
	EventParameterContentType EventParameter = "content_type"
	// EventParameterPageLocation The complete URL of the webpage that someone visited on your website
	EventParameterPageLocation EventParameter = "page_location"
	// EventParameterPageReferrer The referring URL, which is the user's previous URL and can be your website's domain or other domains
	EventParameterPageReferrer EventParameter = "page_referrer"
	// EventParameterPageTitle The HTML page title that you set on your website
	EventParameterPageTitle EventParameter = "page_title"
	// EventParameterScreenResolution The resolution of a device, in the format (Width)x(Height)
	EventParameterScreenResolution EventParameter = "screen_resolution"
	// EventParameterAdFormat The format of an advertisement in an app
	EventParameterAdFormat EventParameter = "ad_format"
	// EventParameterAdPlatform The platform used to surface an advertisement in an app
	EventParameterAdPlatform EventParameter = "ad_platform"
	// EventParameterAdSource The source network that served an advertisement
	EventParameterAdSource EventParameter = "ad_source"
	// EventParameterAdUnitID The unique identifier for an ad unit
	EventParameterAdUnitID EventParameter = "ad_unit_id"
	// EventParameterAdUnitName The name you choose for an ad unit
	EventParameterAdUnitName EventParameter = "ad_unit_name"
	// EventParameterVideoProvider The source of an embedded video
	EventParameterVideoProvider EventParameter = "video_provider"
	// EventParameterVideoTitle The title of an embedded video
	EventParameterVideoTitle EventParameter = "video_title"
	// EventParameterVideoUrl The url of an embedded video
	EventParameterVideoUrl EventParameter = "video_url"

	// Additional undocumented parameters

	EventParameterMethod      EventParameter = "method"
	EventParameterPaymentType EventParameter = "payment_type"
)

func (s EventParameter) String() string {
	return string(s)
}
