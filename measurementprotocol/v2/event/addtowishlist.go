package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

/*
AddToWishlist https://developers.google.com/tag-platform/gtagjs/reference/events#add_to_wishlist

	gtag('event', 'add_to_wishlist', {
		currency: 'USD',
		value: 30.03,
		items: [
			{
				item_id: 'SKU_12345',
				item_name: 'Stan and Friends Tee',
				affiliation: 'Google Store',
				coupon: 'SUMMER_FUN',
				discount: 2.22,
				index: 5,
				item_brand: 'Google',
				item_category: 'Apparel',
				item_category2: 'Adult',
				item_category3: 'Shirts',
				item_category4: 'Crew',
				item_category5: 'Short sleeve',
				item_list_id: 'related_products',
				item_list_name: 'Related products',
				item_variant: 'green',
				location_id: 'ChIJIQBpAG2ahYAR_6128GcTUEo',
				price: 10.01,
				quantity: 3
			}
		]
	});

Query: v=2&tid=G-PZ5ELRCR31&gtm=45je42t1v9177778896z89175355532za200&_p=1709325262551&gcd=13l3l3l3l1&npa=0&dma_cps=sypham&dma=1&cid=1220643501.1708014725&ul=en-us&sr=3840x1600&_fplc=0&ur=DE-BY&uaa=arm&uab=64&uafvl=Chromium%3B122.0.6261.69%7CNot(A%253ABrand%3B24.0.0.0%7CGoogle%2520Chrome%3B122.0.6261.69&uamb=0&uam=&uap=macOS&uapv=14.3.1&uaw=0&are=1&pscdl=noapi&_eu=IA&sst.uc=DE&sst.etld=google.de&sst.gcsub=region1&sst.gcd=13l3l3l3l1&sst.tft=1709325262551&_s=3&cu=USD&sid=1709324719&sct=6&seg=1&dl=https%3A%2F%2Fsniffer.local.bestbytes.net%2F&dt=Server%20Side%20Tracking%20Prototype%20(codename%3A%20sniffer)&en=add_to_wishlist&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Store~cpSUMMER_FUN~ds2.22~lp5~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&epn.value=30.03&_et=3187&tfd=11387&richsstsse
*/
type AddToWishlist struct {
	Currency string  `json:"currency,omitempty"`
	Value    float64 `json:"value,omitempty"`
	Items    []*Item `json:"items,omitempty"`
}

func (e *AddToWishlist) MarshalMPv2() (*mpv2.Event, error) {
	items := make([]*mpv2.Item, len(e.Items))
	for i, item := range e.Items {
		items[i] = item.MarshalMPv2()
	}
	eventParameterNumber := map[string]string{}
	mp.AddStringMap(eventParameterNumber, mpv2.EventParameterNumberValue.String(), mp.SetFloat64(e.Value))
	return &mpv2.Event{
		EventName:            mp.Set(mpv2.EventNameAddToWishlist),
		Currency:             mp.SetString(e.Currency),
		EventParameterNumber: mp.SetStringMap(eventParameterNumber),
		Items:                items,
	}, nil
}
