package gtag_test

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"

	"github.com/foomo/sesamy-go/protocol/gtag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "page_view",
			args: "v=2&tid=G-123456&gtm=45je42s0v893689383z8894072534za200&_p=1709286636310&gcd=13l3l3l3l1&npa=0&dma_cps=sypham&dma=1&cid=292234677.1707898933&ul=en-us&sr=3840x1600&uaa=arm&uab=64&uafvl=Chromium&uamb=0&uam=&uap=macOS&uapv=14.3.1&uaw=0&are=1&pae=1&pscdl=noapi&_s=1&sid=1709286452&sct=7&seg=1&dl=https%3A%2F%2Fwww.homemade.ch%2Fprodukt%2Fkuhn-rikon-wiegemesser-900047100%3Fid%3D900047100&dr=https%3A%2F%2Fwww.homemade.ch%2Fmesser-besteck&dt=Wiegemesser&en=page_view&tfd=5682",
			want: `{"protocol_version":"2","tracking_id":"G-123456","gtmhash_info":"45je42s0v893689383z8894072534za200","client_id":"292234677.1707898933","client_hints":{"screen_resolution":"3840x1600","user_language":"en-us","user_agent_architecture":"arm","user_agent_bitness":"64","user_agent_full_version_list":"Chromium","user_agent_mobile":"0","user_agent_model":"","user_agent_platform":"macOS","user_agent_platform_version":"14.3.1","user_agent_wow_64":"0"},"document_location":"https://www.homemade.ch/produkt/kuhn-rikon-wiegemesser-900047100?id=900047100","document_title":"Wiegemesser","document_referrer":"https://www.homemade.ch/messer-besteck","event_name":"page_view","session_id":"1709286452","ecommerce":{}}`,
		},
		{
			name: "add_to_cart",
			args: "v=2&tid=G-123456&gtm=45je42s0v9175354889z89175348963za200&_p=1709297934217&_dbg=1&gcd=13l3l3l3l1&npa=0&dma_cps=sypham&dma=1&cid=1220643501.1708014725&ul=en-us&sr=3840x1600&_fplc=0&ur=DE-BY&uaa=arm&uab=64&uafvl=Chromium&uamb=0&uam=&uap=macOS&uapv=14.3.1&uaw=0&are=1&pscdl=noapi&_eu=IA&sst.uc=DE&sst.etld=google.de&sst.gcsub=region1&sst.gcd=13l3l3l3l1&sst.tft=1709297934217&_s=8&cu=USD&sid=1709296380&sct=7&seg=1&dl=https%3A%2F%2Fsniffer.cloud.bestbytes.net%2F%3Fgtm_debug%3D1709297933868&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Server%20Side%20Tracking%20Prototype%20(codename%3A%20sniffer)&en=add_to_cart&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Store~cpSUMMER_FUN~ds2.22~lp5~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&epn.value=30.03&tfd=15129&richsstsse",
			want: `{"protocol_version":"2","tracking_id":"G-123456","gtmhash_info":"45je42s0v9175354889z89175348963za200","client_id":"1220643501.1708014725","richsstsse":"","client_hints":{"screen_resolution":"3840x1600","user_language":"en-us","user_agent_architecture":"arm","user_agent_bitness":"64","user_agent_full_version_list":"Chromium","user_agent_mobile":"0","user_agent_model":"","user_agent_platform":"macOS","user_agent_platform_version":"14.3.1","user_agent_wow_64":"0","user_region":"DE-BY"},"document_location":"https://sniffer.cloud.bestbytes.net/?gtm_debug=1709297933868","document_title":"Server Side Tracking Prototype (codename: sniffer)","document_referrer":"https://tagassistant.google.com/","event_name":"add_to_cart","event_parameter_number":{"value":"30.03"},"session_id":"1709296380","ecommerce":{"currency":"USD","items":[{"item_id":"SKU_12345","item_name":"Stan and Friends Tee","item_brand":"Google","item_category":"Apparel","item_category2":"Adult","item_category3":"Shirts","item_category4":"Crew","item_category5":"Short sleeve","item_variant":"green","price":"10.01","quantity":"3","coupon":"SUMMER_FUN","item_list_name":"Related products","item_list_position":"5","item_list_id":"related_products","discount":"2.22","affiliation":"Google Store","location_id":"ChIJIQBpAG2ahYAR_6128GcTUEo"}]},"sst":{"etld":"google.de","gcsub":"region1","uc":"DE","tft":"1709297934217","gcd":"13l3l3l3l1"}}`,
		},
		{
			name: "select_item",
			args: "v=2&tid=G-F9XM71K45T&gtm=45he4580v9184715813z89184708445za200&_p=1715430403224&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=251283723&ul=en-us&sr=3840x1600&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.119%7CGoogle%2520Chrome%3B124.0.6367.119%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1715430403224&sst.ude=0&_s=3&sid=1715428762&sct=2&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1715430402906&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=select_item&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&ep.item_list_id=related_products&ep.item_list_name=Related%20products&_et=89&tfd=8618&richsstsse",
			want: `{"protocol_version":"2","tracking_id":"G-F9XM71K45T","gtmhash_info":"45he4580v9184715813z89184708445za200","client_id":"179294588.1715353601","richsstsse":"","client_hints":{"screen_resolution":"3840x1600","user_language":"en-us","user_agent_architecture":"arm","user_agent_bitness":"64","user_agent_full_version_list":"Chromium;124.0.6367.119|Google Chrome;124.0.6367.119|Not-A.Brand;99.0.0.0","user_agent_mobile":"0","user_agent_model":"","user_agent_platform":"macOS","user_agent_platform_version":"14.4.1","user_agent_wow_64":"0","user_region":""},"document_location":"https://sesamy.local.bestbytes.net/?gtm_debug=1715430402906","document_title":"Home","document_referrer":"https://tagassistant.google.com/","event_name":"select_item","event_parameter":{"enable_page_views":"false","item_list_id":"related_products","item_list_name":"Related products"},"session_id":"1715428762","ecommerce":{"items":[{"item_id":"SKU_12345","item_name":"Stan and Friends Tee","item_brand":"Google","item_category":"Apparel","item_category2":"Adult","item_category3":"Shirts","item_category4":"Crew","item_category5":"Short sleeve","item_variant":"green","price":"10.01","quantity":"3","coupon":"SUMMER_FUN","item_list_name":"Related Products","item_list_position":"0","item_list_id":"related_products","discount":"2.22","affiliation":"Google Merchandise Store","location_id":"ChIJIQBpAG2ahYAR_6128GcTUEo"}]},"sst":{"tft":"1715430403224","gcd":"13l3l3l2l1"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var event gtag.Payload
			require.NoError(t, gtag.DecodeQuery(tt.args, &event))

			out, err := json.Marshal(event)
			require.NoError(t, err)

			assert.Equal(t, tt.want, string(out))
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
