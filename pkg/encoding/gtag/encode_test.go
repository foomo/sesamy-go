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
	testingx.Tags(t, tagx.Short)

	tests := []struct {
		name string
		args string
	}{
		{
			name: "page_view",
			args: "v=2&tid=G-123456&gtm=45je42s0v893689383z8894072534za200&_p=1709286636310&gcd=13l3l3l3l1&npa=0&dma_cps=sypham&dma=1&cid=292234677.1707898933&ul=en-us&sr=3840x1600&uaa=arm&uab=64&uafvl=Chromium&uamb=0&uam=&uap=macOS&uapv=14.3.1&uaw=0&are=1&pae=1&pscdl=noapi&_s=1&sid=1709286452&sct=7&seg=1&dl=https%3A%2F%2Fwww.homemade.ch%2Fprodukt%2Fkuhn-rikon-wiegemesser-900047100%3Fid%3D900047100&dr=https%3A%2F%2Fwww.homemade.ch%2Fmesser-besteck&dt=Wiegemesser&en=page_view&tfd=5682",
		},
		{
			name: "select_item",
			args: "v=2&tid=G-F9XM71K45T&gtm=45he45k0v9184715813z89184708445za200zb9184708445&_p=1716358119533&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=666528402&ul=en-us&sr=3840x1600&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.201%7CGoogle%2520Chrome%3B124.0.6367.201%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1716358119533&sst.ude=0&_s=5&sid=1716358118&sct=10&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716352834644&dr=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1716352834644&dt=Home&en=select_item&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.enable_page_views=false&ep.item_list_id=related_products&ep.item_list_name=Related%20products&_et=9888&tfd=38911&richsstsse",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, err := url.ParseQuery(tt.args)
			require.NoError(t, err)
			var event gtag.Payload

			require.NoError(t, gtag.Decode(values, &event))
			assert.NotEmpty(t, event.Remain)

			if actual, _, err := gtag.Encode(&event); assert.NoError(t, err) {
				if assert.Len(t, actual.Encode(), len(values.Encode())) {
					t.Logf("expected: %s", values.Encode())
					t.Logf("actual:   %s", actual.Encode())
				}
			}
		})
	}
}
