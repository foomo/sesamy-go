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

func TestDecode(t *testing.T) {
	t.Parallel()
	testingx.Tags(t, tagx.Short)

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "page_view",
			args: "v=2&tid=G-F9XM71K45T&gtm=45he5641v9184715813z89184708445za204zb9184708445&_p=1749196701069&gcs=G111&gcd=13r3r3r2r5l1&npa=0&dma_cps=syphamo&dma=1&tag_exp=101509157~103116026~103200004~103233427~103351869~103351871~104653070~104653072~104661466~104661468~104698127~104698129&cid=584335997.1746564151&ecid=1548980841&ul=en-us&sr=1728x1117&lps=1&_fplc=0&ur=DE&uaa=arm&uab=64&uafvl=Chromium%3B136.0.7103.114%7CGoogle%2520Chrome%3B136.0.7103.114%7CNot.A%252FBrand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=15.5.0&uaw=0&are=1&frm=0&pscdl=noapi&ec_mode=a&_eu=AAAAAAQ&sst.rnd=506702095.1749196701&sst.etld=google.de&sst.gcsub=region1&sst.adr=1&sst.tft=1749196701069&sst.lpc=253076363&sst.navt=n&sst.ude=1&sst.sw_exp=1&_s=13&dl=https%3A%2F%2Fsesamy.bestbytes.com%2F%3Futm_source%3Dgoogle%26utm_medium%3Ddemandgen%26utm_campaign%3Ddemandgenprodukte%26utm_id%3D22133718417%26utm_content%3D%26utm_term%3D%26gad_source%3D1%26gad_campaignid%3D22133718417%26gclid%3DCj0KCQjwlrvBBhDnARIsAHEQgOSQ%26gtm_debug%3D1747981544829&dt=Sesamy%20Demo%20Page&sid=1749196701&sct=41&seg=1&dr=https%3A%2F%2Fbestbytes.cloudflareaccess.com%2F&_tu=DA&en=page_view&_et=12397&tfd=285810&richsstsse",
			want: `{"consent":{"google_consent_status":"G111","google_consent_default":"13r3r3r2r5l1"},"campaign":{},"ecommerce":{},"client_hints":{"screen_resolution":"1728x1117","user_language":"en-us","user_agent_architecture":"arm","user_agent_bitness":"64","user_agent_full_version_list":"Chromium;136.0.7103.114|Google%20Chrome;136.0.7103.114|Not.A%2FBrand;99.0.0.0","user_agent_mobile":"0","user_agent_model":"","user_agent_platform":"macOS","user_agent_platform_version":"15.5.0","user_agent_wow_64":"0","user_region":"DE"},"protocol_version":"2","tracking_id":"G-F9XM71K45T","gtmhash_info":"45he5641v9184715813z89184708445za204zb9184708445","client_id":"584335997.1746564151","richsstsse":"","document_location":"https://sesamy.bestbytes.com/?utm_source=google\u0026utm_medium=demandgen\u0026utm_campaign=demandgenprodukte\u0026utm_id=22133718417\u0026utm_content=\u0026utm_term=\u0026gad_source=1\u0026gad_campaignid=22133718417\u0026gclid=Cj0KCQjwlrvBBhDnARIsAHEQgOSQ\u0026gtm_debug=1747981544829","document_title":"Sesamy Demo Page","document_referrer":"https://bestbytes.cloudflareaccess.com/","event_name":"page_view","session_id":"1749196701","non_personalized_ads":"0","sst":{"adr":"1","rnd":"506702095.1749196701","etld":"google.de","gcsub":"region1","tft":"1749196701069","ude":"1","lpc":"253076363","navt":"n","sw_exp":"1"}}`,
		},
		{
			name: "add_to_cart",
			args: "v=2&tid=G-F9XM71K45T&gtm=45he5641v9184715813z89184708445za204zb9184708445&_p=1749196701069&gcs=G111&gcd=13r3r3r2r5l1&npa=0&dma_cps=syphamo&dma=1&tag_exp=101509157~103116026~103200004~103233427~103351869~103351871~104653070~104653072~104661466~104661468~104698127~104698129&cid=584335997.1746564151&ecid=1548980841&ul=en-us&sr=1728x1117&lps=1&_fplc=0&ur=DE&uaa=arm&uab=64&uafvl=Chromium%3B136.0.7103.114%7CGoogle%2520Chrome%3B136.0.7103.114%7CNot.A%252FBrand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=15.5.0&uaw=0&are=1&frm=0&pscdl=noapi&ec_mode=a&_eu=AAAAAAQ&sst.rnd=506702095.1749196701&sst.etld=google.de&sst.gcsub=region1&sst.adr=1&sst.tft=1749196701069&sst.lpc=253076363&sst.navt=n&sst.ude=1&sst.sw_exp=1&_s=14&cu=USD&sid=1749196701&sct=41&seg=1&dl=https%3A%2F%2Fsesamy.bestbytes.com%2F%3Futm_source%3Dgoogle%26utm_medium%3Ddemandgen%26utm_campaign%3Ddemandgenprodukte%26utm_id%3D22133718417%26utm_content%3D%26utm_term%3D%26gad_source%3D1%26gad_campaignid%3D22133718417%26gclid%3DCj0KCQjwlrvBBhDnARIsAHEQgOSQ%26gtm_debug%3D1747981544829&dr=https%3A%2F%2Fbestbytes.cloudflareaccess.com%2F&dt=Sesamy%20Demo%20Page&_tu=DA&en=add_to_cart&_c=1&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&epn.value=30.03&_et=3110&tfd=395426&richsstsse",
			want: `{"consent":{"google_consent_status":"G111","google_consent_default":"13r3r3r2r5l1"},"campaign":{},"ecommerce":{"currency":"USD","items":[{"affiliation":"Google Merchandise Store","coupon":"SUMMER_FUN","discount":"2.22","item_brand":"Google","item_category":"Apparel","item_category2":"Adult","item_category3":"Shirts","item_category4":"Crew","item_category5":"Short sleeve","item_id":"SKU_12345","item_list_id":"related_products","item_list_name":"Related Products","item_name":"Stan and Friends Tee","item_variant":"green","item_list_position":"0","location_id":"ChIJIQBpAG2ahYAR_6128GcTUEo","price":"10.01","quantity":"3"}],"is_conversion":"1"},"client_hints":{"screen_resolution":"1728x1117","user_language":"en-us","user_agent_architecture":"arm","user_agent_bitness":"64","user_agent_full_version_list":"Chromium;136.0.7103.114|Google%20Chrome;136.0.7103.114|Not.A%2FBrand;99.0.0.0","user_agent_mobile":"0","user_agent_model":"","user_agent_platform":"macOS","user_agent_platform_version":"15.5.0","user_agent_wow_64":"0","user_region":"DE"},"protocol_version":"2","tracking_id":"G-F9XM71K45T","gtmhash_info":"45he5641v9184715813z89184708445za204zb9184708445","client_id":"584335997.1746564151","richsstsse":"","document_location":"https://sesamy.bestbytes.com/?utm_source=google\u0026utm_medium=demandgen\u0026utm_campaign=demandgenprodukte\u0026utm_id=22133718417\u0026utm_content=\u0026utm_term=\u0026gad_source=1\u0026gad_campaignid=22133718417\u0026gclid=Cj0KCQjwlrvBBhDnARIsAHEQgOSQ\u0026gtm_debug=1747981544829","document_title":"Sesamy Demo Page","document_referrer":"https://bestbytes.cloudflareaccess.com/","event_name":"add_to_cart","event_parameter_number":{"value":"30.03"},"session_id":"1749196701","non_personalized_ads":"0","sst":{"adr":"1","rnd":"506702095.1749196701","etld":"google.de","gcsub":"region1","tft":"1749196701069","ude":"1","lpc":"253076363","navt":"n","sw_exp":"1"}}`,
		},
		{
			name: "select_item",
			args: "v=2&tid=G-F9XM71K45T&gtm=45he5641v9184715813z89184708445za204zb9184708445&_p=1749196701069&gcs=G111&gcd=13r3r3r2r5l1&npa=0&dma_cps=syphamo&dma=1&tag_exp=101509157~103116026~103200004~103233427~103351869~103351871~104653070~104653072~104661466~104661468~104698127~104698129&cid=584335997.1746564151&ecid=1548980841&ul=en-us&sr=1728x1117&lps=1&_fplc=0&ur=DE&uaa=arm&uab=64&uafvl=Chromium%3B136.0.7103.114%7CGoogle%2520Chrome%3B136.0.7103.114%7CNot.A%252FBrand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=15.5.0&uaw=0&are=1&frm=0&pscdl=noapi&_eu=AAAAAAQ&sst.rnd=506702095.1749196701&sst.etld=google.de&sst.gcsub=region1&sst.adr=1&sst.tft=1749196701069&sst.lpc=253076363&sst.navt=n&sst.ude=1&sst.sw_exp=1&_s=15&sid=1749196701&sct=41&seg=1&dl=https%3A%2F%2Fsesamy.bestbytes.com%2F%3Futm_source%3Dgoogle%26utm_medium%3Ddemandgen%26utm_campaign%3Ddemandgenprodukte%26utm_id%3D22133718417%26utm_content%3D%26utm_term%3D%26gad_source%3D1%26gad_campaignid%3D22133718417%26gclid%3DCj0KCQjwlrvBBhDnARIsAHEQgOSQ%26gtm_debug%3D1747981544829&dr=https%3A%2F%2Fbestbytes.cloudflareaccess.com%2F&dt=Sesamy%20Demo%20Page&_tu=DA&en=select_item&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.item_list_id=related_products&ep.item_list_name=Related%20products&_et=4363&tfd=436740&richsstsse",
			want: `{"consent":{"google_consent_status":"G111","google_consent_default":"13r3r3r2r5l1"},"campaign":{},"ecommerce":{"items":[{"affiliation":"Google Merchandise Store","coupon":"SUMMER_FUN","discount":"2.22","item_brand":"Google","item_category":"Apparel","item_category2":"Adult","item_category3":"Shirts","item_category4":"Crew","item_category5":"Short sleeve","item_id":"SKU_12345","item_list_id":"related_products","item_list_name":"Related Products","item_name":"Stan and Friends Tee","item_variant":"green","item_list_position":"0","location_id":"ChIJIQBpAG2ahYAR_6128GcTUEo","price":"10.01","quantity":"3"}]},"client_hints":{"screen_resolution":"1728x1117","user_language":"en-us","user_agent_architecture":"arm","user_agent_bitness":"64","user_agent_full_version_list":"Chromium;136.0.7103.114|Google%20Chrome;136.0.7103.114|Not.A%2FBrand;99.0.0.0","user_agent_mobile":"0","user_agent_model":"","user_agent_platform":"macOS","user_agent_platform_version":"15.5.0","user_agent_wow_64":"0","user_region":"DE"},"protocol_version":"2","tracking_id":"G-F9XM71K45T","gtmhash_info":"45he5641v9184715813z89184708445za204zb9184708445","client_id":"584335997.1746564151","richsstsse":"","document_location":"https://sesamy.bestbytes.com/?utm_source=google\u0026utm_medium=demandgen\u0026utm_campaign=demandgenprodukte\u0026utm_id=22133718417\u0026utm_content=\u0026utm_term=\u0026gad_source=1\u0026gad_campaignid=22133718417\u0026gclid=Cj0KCQjwlrvBBhDnARIsAHEQgOSQ\u0026gtm_debug=1747981544829","document_title":"Sesamy Demo Page","document_referrer":"https://bestbytes.cloudflareaccess.com/","event_name":"select_item","event_parameter":{"item_list_id":"related_products","item_list_name":"Related products"},"session_id":"1749196701","non_personalized_ads":"0","sst":{"adr":"1","rnd":"506702095.1749196701","etld":"google.de","gcsub":"region1","tft":"1749196701069","ude":"1","lpc":"253076363","navt":"n","sw_exp":"1"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
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
	t.Parallel()
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
			t.Parallel()
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
	t.Parallel()
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
			t.Parallel()
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
	t.Parallel()
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
			t.Parallel()
			if got, err := gtag.DecodeObjectValue(tt.name); assert.NoError(t, err) && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeObjectValue() = %v, want %v", got, tt.want)
				t.Log(got)
			}
		})
	}
}
