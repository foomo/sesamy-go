package gtag_test

import (
	"net/url"
	"testing"

	testingx "github.com/foomo/go/testing"
	tagx "github.com/foomo/go/testing/tag"
	"github.com/foomo/sesamy-go/pkg/encoding/gtag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	t.Parallel()
	testingx.Tags(t, tagx.Short)

	tests := []struct {
		name string
		args string
	}{
		{
			name: "page_view",
			args: "v=2&tid=G-F9XM71K45T&gtm=45he5641v9184715813z89184708445za204zb9184708445&_p=1749196701069&gcs=G100&gcd=13p3p3p2p5l1&npa=1&dma_cps=-&dma=1&tag_exp=101509157~103116026~103200004~103233427~103351869~103351871~104653070~104653072~104661466~104661468~104698127~104698129&gtm_up=1&cid=1174285007.1749196701&ecid=1548980841&ul=en-us&sr=1728x1117&lps=1&_fplc=0&ur=DE&uaa=arm&uab=64&uafvl=Chromium%3B136.0.7103.114%7CGoogle%2520Chrome%3B136.0.7103.114%7CNot.A%252FBrand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=15.5.0&uaw=0&are=1&frm=0&pscdl=denied&sst.rnd=506702095.1749196701&sst.etld=google.de&sst.gcsub=region1&sst.adr=1&sst.tft=1749196701069&sst.lpc=253076363&sst.navt=n&sst.ude=1&_s=3&sid=1749196701&sct=1&seg=0&dl=https%3A%2F%2Fsesamy.bestbytes.com%2F%3Futm_source%3Dgoogle%26utm_medium%3Ddemandgen%26utm_campaign%3Ddemandgenprodukte%26utm_id%3D22133718417%26utm_content%3D%26utm_term%3D%26gad_source%3D1%26gad_campaignid%3D22133718417%26gclid%3DCj0KCQjwlrvBBhDnARIsAHEQgOSQ%26gtm_debug%3D1747981544829&dr=https%3A%2F%2Fbestbytes.cloudflareaccess.com%2F&dt=Sesamy&_tu=DA&en=scroll&_et=NaN&epn.percent_scrolled=90&tfd=8412&richsstsse",
		},
		{
			name: "add_to_cart",
			args: "v=2&tid=G-F9XM71K45T&gtm=45he5641v9184715813z89184708445za204zb9184708445&_p=1749196701069&gcs=G100&gcd=13p3p3p2p5l1&npa=1&dma_cps=-&dma=1&tag_exp=101509157~103116026~103200004~103233427~103351869~103351871~104653070~104653072~104661466~104661468~104698127~104698129&gtm_up=1&cid=1174285007.1749196701&ecid=1548980841&ul=en-us&sr=1728x1117&lps=1&_fplc=0&ur=DE&uaa=arm&uab=64&uafvl=Chromium%3B136.0.7103.114%7CGoogle%2520Chrome%3B136.0.7103.114%7CNot.A%252FBrand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=15.5.0&uaw=0&are=1&frm=0&pscdl=denied&ec_mode=a&sst.rnd=506702095.1749196701&sst.etld=google.de&sst.gcsub=region1&sst.adr=1&sst.tft=1749196701069&sst.lpc=253076363&sst.navt=n&sst.ude=1&_s=4&cu=USD&sid=1749196701&sct=1&seg=0&dl=https%3A%2F%2Fsesamy.bestbytes.com%2F%3Futm_source%3Dgoogle%26utm_medium%3Ddemandgen%26utm_campaign%3Ddemandgenprodukte%26utm_id%3D22133718417%26utm_content%3D%26utm_term%3D%26gad_source%3D1%26gad_campaignid%3D22133718417%26gclid%3DCj0KCQjwlrvBBhDnARIsAHEQgOSQ%26gtm_debug%3D1747981544829&dr=https%3A%2F%2Fbestbytes.cloudflareaccess.com%2F&dt=Sesamy&_tu=DA&en=add_to_cart&_c=1&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&epn.value=30.03&_et=8715&tfd=53322&richsstsse",
		},
		{
			name: "purchase",
			args: "v=2&tid=G-F9XM71K45T&gtm=45he5641v9184715813z89184708445za204zb9184708445&_p=1749196701069&gcs=G100&gcd=13p3p3p2p5l1&npa=1&dma_cps=-&dma=1&tag_exp=101509157~103116026~103200004~103233427~103351869~103351871~104653070~104653072~104661466~104661468~104698127~104698129&gtm_up=1&cid=1174285007.1749196701&ecid=1548980841&ul=en-us&sr=1728x1117&lps=1&_fplc=0&ur=DE&uaa=arm&uab=64&uafvl=Chromium%3B136.0.7103.114%7CGoogle%2520Chrome%3B136.0.7103.114%7CNot.A%252FBrand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=15.5.0&uaw=0&are=1&frm=0&pscdl=denied&ec_mode=a&sst.rnd=506702095.1749196701&sst.etld=google.de&sst.gcsub=region1&sst.adr=1&sst.tft=1749196701069&sst.lpc=253076363&sst.navt=n&sst.ude=1&_s=4&cu=USD&sid=1749196701&sct=1&seg=0&dl=https%3A%2F%2Fsesamy.bestbytes.com%2F%3Futm_source%3Dgoogle%26utm_medium%3Ddemandgen%26utm_campaign%3Ddemandgenprodukte%26utm_id%3D22133718417%26utm_content%3D%26utm_term%3D%26gad_source%3D1%26gad_campaignid%3D22133718417%26gclid%3DCj0KCQjwlrvBBhDnARIsAHEQgOSQ%26gtm_debug%3D1747981544829&dr=https%3A%2F%2Fbestbytes.cloudflareaccess.com%2F&dt=Sesamy&_tu=DA&en=add_to_cart&_c=1&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&epn.value=30.03&_et=8715&tfd=53322&richsstsse",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			values, err := url.ParseQuery(tt.args)
			require.NoError(t, err)
			var event gtag.Payload

			require.NoError(t, gtag.Decode(values, &event))
			assert.NotEmpty(t, event.Remain)

			if actual, _, err := gtag.Encode(&event); assert.NoError(t, err) {
				if !assert.Len(t, actual.Encode(), len(values.Encode())) {
					t.Logf("expected: %s", values.Encode())
					t.Logf("actual:   %s", actual.Encode())
				}
			}
		})
	}
}
