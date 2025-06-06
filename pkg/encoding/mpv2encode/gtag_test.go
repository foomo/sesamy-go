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
	assert.Equal(t, iso4217.USD, intermediate.Events[0].Params.Currency)

	intermediate.Events[0].Params.Currency = iso4217.EUR

	{
		out, err := json.MarshalIndent(intermediate, "", "  ")
		require.NoError(t, err)
		fmt.Println(string(out))
	}

	var outgoing gtag.Payload
	err = mpv2encode.GTag(*intermediate, &outgoing)
	require.NoError(t, err)
	assert.Equal(t, iso4217.EUR, gtag.Get(outgoing.Currency))
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
	assert.Equal(t, iso4217.USD, intermediate.Events[0].Params.Currency)

	// override value
	intermediate.Events[0].Params.Currency = iso4217.EUR

	// {
	// 	out, err := json.MarshalIndent(intermediate, "", "  ")
	// 	require.NoError(t, err)
	// 	fmt.Println(string(out))
	// }

	err = mpv2encode.GTag(*intermediate, &incoming)
	require.NoError(t, err)
	assert.Equal(t, iso4217.EUR, gtag.Get(incoming.Currency))
}
