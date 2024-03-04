package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

/*
BeginCheckout https://developers.google.com/tag-platform/gtagjs/reference/events#begin_checkout

	gtag('event', 'begin_checkout', {
		currency: 'USD',
		value: 30.03,
		coupon: 'SUMMER_FUN',
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

Query: v=2&tid=G-123456&gtm=45je42t1v9177778896z89175355532za200&_p=1709325262551&gcd=13l3l3l3l1&npa=0&dma_cps=sypham&dma=1&cid=1220643501.1708014725&ul=en-us&sr=3840x1600&_fplc=0&ur=DE-BY&uaa=arm&uab=64&uafvl=Chromium%3B122.0.6261.69%7CNot(A%253ABrand%3B24.0.0.0%7CGoogle%2520Chrome%3B122.0.6261.69&uamb=0&uam=&uap=macOS&uapv=14.3.1&uaw=0&are=1&pscdl=noapi&_eu=IA&sst.uc=DE&sst.etld=google.de&sst.gcsub=region1&sst.gcd=13l3l3l3l1&sst.tft=1709325262551&_s=5&cu=USD&sid=1709445696&sct=8&seg=0&dl=https%3A%2F%2Fsniffer.local.bestbytes.net%2F&dt=Server%20Side%20Tracking%20Prototype%20(codename%3A%20sniffer)&en=begin_checkout&_ss=1&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Store~cpSUMMER_FUN~ds2.22~lp5~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&epn.value=30.03&ep.coupon=SUMMER_FUN&tfd=120434230&richsstsse
*/
type BeginCheckout struct {
	Currency string
	Value    float64
	Coupon   string
	Items    []*Item
}

func (e *BeginCheckout) MPv2() *mpv2.Event {
	items := make([]*mpv2.Item, len(e.Items))
	for i, item := range e.Items {
		items[i] = item.MPv2()
	}
	eventParameter := map[string]string{}
	mp.AddStringMap(eventParameter, mpv2.EventParameterCoupon.String(), mp.SetString(e.Coupon))
	eventParameterNumber := map[string]string{}
	mp.AddStringMap(eventParameterNumber, mpv2.EventParameterNumberValue.String(), mp.SetFloat64(e.Value))
	return &mpv2.Event{
		Currency:             mp.SetString(e.Currency),
		EventParameter:       mp.SetStringMap(eventParameter),
		EventParameterNumber: mp.SetStringMap(eventParameterNumber),
		Items:                items,
	}
}
