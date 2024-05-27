package gtag_test

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"

	testingx "github.com/foomo/go/testing"
	tagx "github.com/foomo/go/testing/tag"
	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	GTagAddPaymentInfo   = "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=788548699&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=2&cu=USD&sid=1716793773&sct=14&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=add_payment_info&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&ep.coupon=SUMMER_FUN&ep.payment_type=Credit%20Card&epn.value=30.03&_et=169&tfd=28197&richsstsse"
	GTagAddShippingInfo  = "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=788548699&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=3&cu=USD&sid=1716793773&sct=14&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=add_shipping_info&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&ep.coupon=SUMMER_FUN&ep.shipping_tier=Ground&epn.value=30.03&_et=2112&tfd=114276&richsstsse"
	GTagAddToCart        = "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=788548699&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=4&cu=USD&sid=1716793773&sct=14&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=add_to_cart&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&epn.value=30.03&_et=1255&tfd=145479&richsstsse"
	GTagAddToWishlist    = "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=788548699&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=5&cu=USD&sid=1716793773&sct=14&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=add_to_wishlist&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&epn.value=30.03&_et=2081&tfd=175581&richsstsse"
	GTagAddBeginCheckout = "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=788548699&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=6&cu=USD&sid=1716793773&sct=14&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=begin_checkout&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&ep.coupon=SUMMER_FUN&epn.value=30.03&_et=1649&tfd=210309&richsstsse"
	GTagGenerateLead     = "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=788548699&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=7&cu=USD&sid=1716793773&sct=14&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=generate_lead&ep.enable_page_views=false&epn.value=99.99&_et=1680&tfd=247472&richsstsse"
	GTagJoinGroup        = "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=788548699&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=8&sid=1716793773&sct=14&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=join_group&ep.enable_page_views=false&ep.group_id=G_12345&_et=1510&tfd=265924&richsstsse"
	GTagLevelEnd         = "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=788548699&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=9&sid=1716793773&sct=14&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=level_end&ep.enable_page_views=false&ep.level_name=The%20journey%20begins...&ep.success=true&_et=2950&tfd=296976&richsstsse"
	GTagLevelStart       = "v=2&tid=G-F9XM71K45T&gtm=45he45m0v9184715813z89184708445za200zb9184708445&_p=1716795486104&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=788548699&ul=en-us&sr=2056x1329&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716795486104&sst.ude=0&_s=10&sid=1716793773&sct=14&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716795486020&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=level_start&ep.enable_page_views=false&ep.level_name=The%20journey%20begins...&_et=1266&tfd=321602&richsstsse"
)

func TestDecode(t *testing.T) {
	testingx.Tags(t, tagx.Short)

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "page_view",
			args: "v=2&tid=G-123456&gtm=45je42s0v893689383z8894072534za200&_p=1709286636310&gcd=13l3l3l3l1&npa=0&dma_cps=sypham&dma=1&cid=292234677.1707898933&ul=en-us&sr=3840x1600&uaa=arm&uab=64&uafvl=Chromium&uamb=0&uam=&uap=macOS&uapv=14.3.1&uaw=0&are=1&pae=1&pscdl=noapi&_s=1&sid=1709286452&sct=7&seg=1&dl=https%3A%2F%2Fwww.homemade.ch%2Fprodukt%2Fkuhn-rikon-wiegemesser-900047100%3Fid%3D900047100&dr=https%3A%2F%2Fwww.homemade.ch%2Fmesser-besteck&dt=Wiegemesser&en=page_view&tfd=5682",
			want: `{"consent":{"google_consent_default":"13l3l3l3l1"},"campaign":{},"ecommerce":{},"client_hints":{"screen_resolution":"3840x1600","user_language":"en-us","user_agent_architecture":"arm","user_agent_bitness":"64","user_agent_full_version_list":"Chromium","user_agent_mobile":"0","user_agent_model":"","user_agent_platform":"macOS","user_agent_platform_version":"14.3.1","user_agent_wow_64":"0"},"protocol_version":"2","tracking_id":"G-123456","gtmhash_info":"45je42s0v893689383z8894072534za200","client_id":"292234677.1707898933","document_location":"https://www.homemade.ch/produkt/kuhn-rikon-wiegemesser-900047100?id=900047100","document_title":"Wiegemesser","document_referrer":"https://www.homemade.ch/messer-besteck","event_name":"page_view","session_id":"1709286452","non_personalized_ads":"0"}`,
		},
		{
			name: "add_to_cart",
			args: GTagAddToCart,
			want: `{"consent":{"google_consent_default":"13l3l3l2l1"},"campaign":{},"ecommerce":{"currency":"USD","items":[{"affiliation":"Google Merchandise Store","coupon":"SUMMER_FUN","discount":"2.22","item_brand":"Google","item_category":"Apparel","item_category2":"Adult","item_category3":"Shirts","item_category4":"Crew","item_category5":"Short sleeve","item_id":"SKU_12345","item_list_id":"related_products","item_list_name":"Related Products","item_name":"Stan and Friends Tee","item_variant":"green","item_list_position":"0","location_id":"ChIJIQBpAG2ahYAR_6128GcTUEo","price":"10.01","quantity":"3"}]},"client_hints":{"screen_resolution":"2056x1329","user_language":"en-us","user_agent_architecture":"arm","user_agent_bitness":"64","user_agent_full_version_list":"Chromium;124.0.6367.201|Google%20Chrome;124.0.6367.201|Not-A.Brand;99.0.0.0","user_agent_mobile":"0","user_agent_model":"","user_agent_platform":"macOS","user_agent_platform_version":"14.4.1","user_agent_wow_64":"0","user_region":""},"protocol_version":"2","tracking_id":"G-F9XM71K45T","gtmhash_info":"45he45m0v9184715813z89184708445za200zb9184708445","client_id":"179294588.1715353601","richsstsse":"","document_location":"https://sesamy.local.bestbytes.net/?gtm_debug=1716795486020","document_title":"Home","document_referrer":"https://tagassistant.google.com/","is_debug":"1","event_name":"add_to_cart","event_parameter":{"enable_page_views":"false"},"event_parameter_number":{"value":"30.03"},"session_id":"1716793773","non_personalized_ads":"1","sst":{"tft":"1716795486104","gcd":"13l3l3l2l1","ude":"0"}}`,
		},
		{
			name: "select_item",
			args: "v=2&tid=G-F9XM71K45T&gtm=45he4580v9184715813z89184708445za200&_p=1715430403224&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=251283723&ul=en-us&sr=3840x1600&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.119%7CGoogle%2520Chrome%3B124.0.6367.119%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1715430403224&sst.ude=0&_s=3&sid=1715428762&sct=2&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1715430402906&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=select_item&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&ep.item_list_id=related_products&ep.item_list_name=Related%20products&_et=89&tfd=8618&richsstsse",
			want: `{"consent":{"google_consent_default":"13l3l3l2l1"},"campaign":{},"ecommerce":{"items":[{"affiliation":"Google Merchandise Store","coupon":"SUMMER_FUN","discount":"2.22","item_brand":"Google","item_category":"Apparel","item_category2":"Adult","item_category3":"Shirts","item_category4":"Crew","item_category5":"Short sleeve","item_id":"SKU_12345","item_list_id":"related_products","item_list_name":"Related Products","item_name":"Stan and Friends Tee","item_variant":"green","item_list_position":"0","location_id":"ChIJIQBpAG2ahYAR_6128GcTUEo","price":"10.01","quantity":"3"}]},"client_hints":{"screen_resolution":"3840x1600","user_language":"en-us","user_agent_architecture":"arm","user_agent_bitness":"64","user_agent_full_version_list":"Chromium;124.0.6367.119|Google%20Chrome;124.0.6367.119|Not-A.Brand;99.0.0.0","user_agent_mobile":"0","user_agent_model":"","user_agent_platform":"macOS","user_agent_platform_version":"14.4.1","user_agent_wow_64":"0","user_region":""},"protocol_version":"2","tracking_id":"G-F9XM71K45T","gtmhash_info":"45he4580v9184715813z89184708445za200","is_debug":"1","client_id":"179294588.1715353601","richsstsse":"","document_location":"https://sesamy.local.bestbytes.net/?gtm_debug=1715430402906","document_title":"Home","document_referrer":"https://tagassistant.google.com/","event_name":"select_item","event_parameter":{"enable_page_views":"false","item_list_id":"related_products","item_list_name":"Related products"},"session_id":"1715428762","non_personalized_ads":"1","sst":{"tft":"1715430403224", "ude":"0", "gcd":"13l3l3l2l1"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var event gtag.Payload
			require.NoError(t, gtag.DecodeQuery(tt.args, &event))

			out, err := json.Marshal(event)
			require.NoError(t, err)

			if !assert.JSONEq(t, tt.want, string(out)) {
				t.Log(string(out))
			}
		})
	}
}

func TestDecodeMapValue(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "v=2",
			want: false,
		},
		{
			name: "ep.foo=bar",
			want: true,
		},
		{
			name: "ep.foo.bar=bar",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := gtag.Data{}
			values, err := url.ParseQuery(tt.name)
			require.NoError(t, err)

			for k, v := range values {
				if got, err := gtag.DecodeMapValue(k, v, d); assert.NoError(t, err) && got != tt.want {
					t.Errorf("DecodeMapValue() = %v, want %v", got, tt.want)
					t.Log(d)
				}
			}
		})
	}
}

func TestDecodeProductValue(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "v=2",
			want: false,
		},
		{
			name: "pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr9.99~qt1",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := gtag.Data{}
			values, err := url.ParseQuery(tt.name)
			require.NoError(t, err)
			for k, v := range values {
				if got, err := gtag.DecodeRegexValue(k, v, gtag.RegexProduct, d, "pr"); assert.NoError(t, err) && got != tt.want {
					t.Errorf("decodeMapValue() = %v, want %v", got, tt.want)
					t.Log(d)
				}
			}
		})
	}
}

func TestDecodeObjectValue(t *testing.T) {
	tests := []struct {
		name string
		args string
		want map[string]any
	}{
		{
			name: "",
			want: nil,
		},
		{
			name: "idSKU_12345~nmStan%20and%20Friends%20Tee~qt1",
			want: map[string]any{
				"id": "SKU_12345",
				"nm": "Stan and Friends Tee",
				"qt": "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := gtag.DecodeObjectValue(tt.name); assert.NoError(t, err) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeObjectValue() = %v, want %v", got, tt.want)
				t.Log(got)
			}
		})
	}
}
