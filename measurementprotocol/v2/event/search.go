package event

import (
	mp "github.com/foomo/sesamy-go/measurementprotocol"
	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
)

/*
Search https://developers.google.com/tag-platform/gtagjs/reference/events#search

	gtag('event', 'search', {
		search_term: 't-shirts',
	});

Query: v=2&tid=G-PZ5ELRCR31&gtm=45je42t1v9177778896z89175355532za200&_p=1709325262551&gcd=13l3l3l3l1&npa=0&dma_cps=sypham&dma=1&cid=1220643501.1708014725&ul=en-us&sr=3840x1600&_fplc=0&ur=DE-BY&uaa=arm&uab=64&uafvl=Chromium%3B122.0.6261.69%7CNot(A%253ABrand%3B24.0.0.0%7CGoogle%2520Chrome%3B122.0.6261.69&uamb=0&uam=&uap=macOS&uapv=14.3.1&uaw=0&are=1&pscdl=noapi&_eu=IA&sst.uc=DE&sst.etld=google.de&sst.gcsub=region1&sst.gcd=13l3l3l3l1&sst.tft=1709325262551&_s=3&cu=USD&sid=1709324719&sct=6&seg=1&dl=https%3A%2F%2Fsniffer.local.bestbytes.net%2F&dt=Server%20Side%20Tracking%20Prototype%20(codename%3A%20sniffer)&en=search&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Store~cpSUMMER_FUN~ds2.22~lp5~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&epn.value=30.03&_et=3187&tfd=11387&richsstsse
*/
type Search struct {
	SearchTerm string `json:"search_term,omitempty"`
}

func (e *Search) MarshalMPv2() (*mpv2.Event, error) {
	eventParameter := map[string]string{}
	mp.AddStringMap(eventParameter, mpv2.EventParameterSearchTerm.String(), mp.SetString(e.SearchTerm))
	return &mpv2.Event{
		EventName:      mp.Set(mpv2.EventNameSearch),
		EventParameter: mp.SetStringMap(eventParameter),
	}, nil
}
