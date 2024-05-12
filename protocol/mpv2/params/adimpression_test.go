package params_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/foomo/sesamy-go/protocol/gtag"
	"github.com/foomo/sesamy-go/protocol/mpv2"
	"github.com/foomo/sesamy-go/protocol/mpv2/params"
	"github.com/stretchr/testify/require"
)

func TestSelectItem(t *testing.T) {
	query := "v=2&tid=G-F9XM71K45T&gtm=45he4580v9184715813z89184708445za200&_p=1715430403224&_dbg=1&gcd=13l3l3l2l1&npa=1&dma_cps=sypham&dma=1&cid=179294588.1715353601&ecid=251283723&ul=en-us&sr=3840x1600&_fplc=0&ur=&uaa=arm&uab=64&uafvl=Chromium%3B124.0.6367.119%7CGoogle%2520Chrome%3B124.0.6367.119%7CNot-A.Brand%3B99.0.0.0&uamb=0&uam=&uap=macOS&uapv=14.4.1&uaw=0&are=1&frm=0&pscdl=noapi&sst.gcd=13l3l3l2l1&sst.tft=1715430403224&sst.ude=0&_s=3&sid=1715428762&sct=2&seg=1&dl=https%3A%2F%2Fsesamy.local.bestbytes.net%2F%3Fgtm_debug%3D1715430402906&dr=https%3A%2F%2Ftagassistant.google.com%2F&dt=Home&en=select_item&pr1=idSKU_12345~nmStan%20and%20Friends%20Tee~afGoogle%20Merchandise%20Store~cpSUMMER_FUN~ds2.22~lp0~brGoogle~caApparel~c2Adult~c3Shirts~c4Crew~c5Short%20sleeve~lirelated_products~lnRelated%20Products~vagreen~loChIJIQBpAG2ahYAR_6128GcTUEo~pr10.01~qt3&ep.item_list_id=related_products&ep.item_list_name=Related%20products&_et=89&tfd=8618&richsstsse"
	values, err := url.ParseQuery(query)
	require.NoError(t, err)

	var incoming gtag.Payload
	err = gtag.Decode(values, &incoming)
	require.NoError(t, err)

	var intermediate mpv2.Payload[params.SelectItem[params.Item]]
	err = params.Decode(incoming, &intermediate)
	require.NoError(t, err)

	{
		out, err := json.MarshalIndent(intermediate, "", "  ")
		require.NoError(t, err)
		fmt.Println(string(out))
	}

	// var outgoing gtag.Event
	// err = event.Encode(intermediate, &outgoing)
	// require.NoError(t, err)
	// assert.Equal(t, incoming, outgoing)

	err = params.Encode(intermediate, &incoming)
	require.NoError(t, err)
	// assert.Equal(t, incoming, outgoing)
}
