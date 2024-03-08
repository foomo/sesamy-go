package v2_test

import (
	"net/url"
	"testing"

	mpv2 "github.com/foomo/sesamy-go/measurementprotocol/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
		args string
		want error
	}{
		{
			name: "page_view",
			args: "v=2&tid=G-123456&gtm=45je42s0v893689383z8894072534za200&_p=1709286636310&gcd=13l3l3l3l1&npa=0&dma_cps=sypham&dma=1&cid=292234677.1707898933&ul=en-us&sr=3840x1600&uaa=arm&uab=64&uafvl=Chromium&uamb=0&uam=&uap=macOS&uapv=14.3.1&uaw=0&are=1&pae=1&pscdl=noapi&_s=1&sid=1709286452&sct=7&seg=1&dl=https%3A%2F%2Fwww.homemade.ch%2Fprodukt%2Fkuhn-rikon-wiegemesser-900047100%3Fid%3D900047100&dr=https%3A%2F%2Fwww.homemade.ch%2Fmesser-besteck&dt=Wiegemesser&en=page_view&tfd=5682",
			want: nil,
		},
		// {
		//	name: "add_to_cart",
		//	args: "v=2&tid=G-123456&gtm=45je42s0v9175354889z89175348963za200&_p=1709297934217&_dbg=1&gcd=13l3l3l3l1&npa=0&dma_cps=sypham&dma=1&cid=1220643501.1708014725&ul=en-us&sr=3840x1600&_fplc=0&ur=DE-BY&uaa=arm&uab=64&uafvl=Chromium&uamb=0&uam=&uap=macOS&uapv=14.3.1&uaw=0&are=1&pscdl=noapi&_eu=IA&sst.uc=DE&sst.etld=google.de&sst.gcsub=region1&sst.gcd=13l3l3l3l1&sst.tft=1709297934217&_s=8&cu=USD&sid=1709296380&sct=7&seg=1&dl=https%3A%2F%2Fsniffer.cloud.bestbytes.net%2F%3Fgtm_debug%3D1709297933868&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Server%20Side%20Tracking%20Prototype%20(codename%3A%20sniffer)&en=add_to_cart&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Store~cpSUMMER_FUN~ds2.22~lp5~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&epn.value=30.03&tfd=15129&richsstsse",
		//	want: nil,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := url.ParseQuery(tt.args)
			require.NoError(t, err)
			e := &mpv2.Event{}
			require.NoError(t, mpv2.Decode(u, e))
			if !assert.Empty(t, e.Unknown) {
				t.Errorf("Encode() = %v, want %v", e.Unknown, nil)
			}
			if out, _, err := mpv2.Encode(e); assert.NoError(t, err) {
				assert.EqualValues(t, u, out)
			}
		})
	}
}
