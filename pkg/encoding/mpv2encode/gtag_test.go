package mpv2encode_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/foomo/gostandards/iso4217"
	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/foomo/sesamy-go/pkg/encoding/gtagencode"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2"
	"github.com/foomo/sesamy-go/pkg/encoding/mpv2encode"
	"github.com/foomo/sesamy-go/pkg/event/params"
	"github.com/foomo/sesamy-go/pkg/sesamy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectItem(t *testing.T) {
	t.Parallel()

	query := "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=2065234266&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=18&cu=USD&sid=1716807360&sct=16&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=remove_from_cart&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&epn.value=30.03&_et=10086&tfd=9253808&richsstsse"
	values, err := url.ParseQuery(query)
	require.NoError(t, err)

	var incoming gtag.Payload

	err = gtag.Decode(values, &incoming)
	require.NoError(t, err)

	var intermediate *mpv2.Payload[params.RemoveFromCart[params.Item]]

	err = gtagencode.MPv2(incoming, &intermediate)
	require.NoError(t, err)
	assert.Equal(t, iso4217.CodeUSD, intermediate.Events[0].Params.Currency)

	intermediate.Events[0].Params.Currency = iso4217.CodeEUR

	{
		out, err := json.MarshalIndent(intermediate, "", "  ")
		require.NoError(t, err)
		fmt.Println(string(out))
	}

	var outgoing gtag.Payload

	err = mpv2encode.GTag(*intermediate, &outgoing)
	require.NoError(t, err)
	assert.Equal(t, iso4217.CodeEUR, gtag.Get(outgoing.Currency))
}

func TestSelectItem_Pointer(t *testing.T) {
	t.Parallel()

	query := "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=2065234266&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=18&cu=USD&sid=1716807360&sct=16&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=remove_from_cart&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&epn.value=30.03&_et=10086&tfd=9253808&richsstsse"
	values, err := url.ParseQuery(query)
	require.NoError(t, err)

	var incoming *gtag.Payload

	err = gtag.Decode(values, &incoming)
	require.NoError(t, err)

	var intermediate *mpv2.Payload[params.RemoveFromCart[params.Item]]

	err = gtagencode.MPv2(*incoming, &intermediate)
	require.NoError(t, err)
	assert.Equal(t, iso4217.CodeUSD, intermediate.Events[0].Params.Currency)

	// override value
	intermediate.Events[0].Params.Currency = iso4217.CodeEUR

	// {
	// 	out, err := json.MarshalIndent(intermediate, "", "  ")
	// 	require.NoError(t, err)
	// 	fmt.Println(string(out))
	// }

	err = mpv2encode.GTag(*intermediate, &incoming)
	require.NoError(t, err)
	assert.Equal(t, iso4217.CodeEUR, gtag.Get(incoming.Currency))
}

func TestGTag_MissingEvents(t *testing.T) {
	t.Parallel()

	var payload gtag.Payload
	assert.Error(t, mpv2encode.GTag(mpv2.Payload[any]{ClientID: "123.456"}, &payload))
}

// TestTrackingURL converts a gtag.js tracking query into a Measurement
// Protocol v2 payload and back into a tracking query.
func TestTrackingURL(t *testing.T) {
	t.Parallel()

	const (
		query = "_dbg=1&cid=123.456&cu=EUR&dl=https%3A%2F%2Fexample.com%2Fcart%3Fa%3D1&dr=https%3A%2F%2Fexample.com%2F&dt=Cart%20Page&en=add_to_cart&ep.coupon=SUMMER&epn.value=21&gcs=G111&npa=0&pr1=idSKU_1~nmShirt%20%26%20Tie~pr10.5~qt2&sid=789&tid=G-123456&uid=U1&up.tier=gold&upn.age=42&v=2&richsstsse"
		// NOTE: mpv2 payloads do not carry v, tid, npa and richsstsse back into gtag.js
		wantQuery = "_dbg=1&cid=123.456&cu=EUR&dl=https%3A%2F%2Fexample.com%2Fcart%3Fa%3D1&dr=https%3A%2F%2Fexample.com%2F&dt=Cart%20Page&en=add_to_cart&ep.coupon=SUMMER&epn.value=21&gcd=13t3t3t2t5l1&gcs=G111&pr1=idSKU_1~nmShirt%20%26%20Tie~pr10.5~qt2&sid=789&uid=U1&up.tier=gold&upn.age=42"
		wantMPv2  = `{
			"client_id": "123.456",
			"user_id": "U1",
			"session_id": "789",
			"debug_mode": true,
			"user_properties": {"age": "42", "tier": "gold"},
			"consent": {
				"ad_storage": "GRANTED",
				"ad_user_data": "GRANTED",
				"ad_personalization": "GRANTED",
				"analytics_storage": "GRANTED"
			},
			"events": [{
				"name": "add_to_cart",
				"params": {
					"coupon": "SUMMER",
					"currency": "EUR",
					"items": [{"item_id": "SKU_1", "item_name": "Shirt & Tie", "price": "10.5", "quantity": "2"}],
					"page_location": "https://example.com/cart?a=1",
					"page_referrer": "https://example.com/",
					"page_title": "Cart Page",
					"value": 21
				}
			}]
		}`
	)

	var incoming gtag.Payload
	require.NoError(t, gtag.DecodeQuery(query, &incoming))

	var intermediate mpv2.Payload[any]
	require.NoError(t, gtagencode.MPv2(incoming, &intermediate))

	out, err := json.Marshal(intermediate)
	require.NoError(t, err)
	assert.JSONEq(t, wantMPv2, string(out))

	var outgoing gtag.Payload
	require.NoError(t, mpv2encode.GTag(intermediate, &outgoing))

	actual, body, err := gtag.Encode(&outgoing)
	require.NoError(t, err)
	assert.Nil(t, body)
	assert.Equal(t, wantQuery, actual)
}

func TestGTag_Consent(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		consent *mpv2.ConsentData
		wantGCS string
		wantGCD string
	}{
		{
			name: "granted",
			consent: &mpv2.ConsentData{
				AdStorage:         new(mpv2.ConsentGranted),
				AdUserData:        new(mpv2.ConsentGranted),
				AdPersonalization: new(mpv2.ConsentGranted),
				AnalyticsStorage:  new(mpv2.ConsentGranted),
			},
			wantGCS: "G111",
			wantGCD: "13t3t3t2t5l1",
		},
		{
			name: "mixed",
			consent: &mpv2.ConsentData{
				AdStorage:         new(mpv2.ConsentDenied),
				AdUserData:        new(mpv2.ConsentDenied),
				AdPersonalization: new(mpv2.ConsentGranted),
				AnalyticsStorage:  new(mpv2.ConsentGranted),
			},
			wantGCS: "G101",
			wantGCD: "13p3t3p2t5l1",
		},
		{
			name:    "unset",
			consent: &mpv2.ConsentData{},
			wantGCS: "G111",
			wantGCD: "13l3l3l2l5l1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			source := mpv2.Payload[any]{
				Consent: tt.consent,
				Events:  []sesamy.Event[any]{{Name: "page_view"}},
			}

			var outgoing gtag.Payload
			require.NoError(t, mpv2encode.GTag(source, &outgoing))
			assert.Equal(t, tt.wantGCS, gtag.Get(outgoing.GoogleConsentStatus))
			assert.Equal(t, tt.wantGCD, gtag.Get(outgoing.GoogleConsentDefault))

			// the gtag consent must convert back into the same mpv2 consent
			var back mpv2.Payload[any]
			require.NoError(t, gtagencode.MPv2(outgoing, &back))

			for _, c := range []struct{ want, got *mpv2.Consent }{
				{tt.consent.AdStorage, back.Consent.AdStorage},
				{tt.consent.AdUserData, back.Consent.AdUserData},
				{tt.consent.AdPersonalization, back.Consent.AdPersonalization},
				{tt.consent.AnalyticsStorage, back.Consent.AnalyticsStorage},
			} {
				if c.want != nil {
					assert.Equal(t, *c.want, *c.got)
				}
			}
		})
	}
}
