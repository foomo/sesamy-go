package gtag_test

import (
	"net/url"
	"strings"
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

			if actual, body, err := gtag.Encode(&event); assert.NoError(t, err) {
				expected := gtag.EncodeValues(values)
				if !assert.Len(t, actual, len(expected)) {
					t.Logf("expected: %s", expected)
					t.Logf("actual:   %s", actual)
				}
				// these fixtures fit within the query length limit
				assert.Nil(t, body)

				// the returned query must be fully encoded and parse back cleanly,
				// preserving every parameter key
				parsed, err := url.ParseQuery(actual)
				require.NoError(t, err)
				assert.Equal(t, len(values), len(parsed))
				for k, want := range values {
					// NOTE: EncodeObjectValue sorts item sub-keys, so `pr<n>`
					// values round-trip with reordered (not lost) content
					if gtag.RegexProduct.MatchString(k) {
						assert.ElementsMatch(t,
							strings.Split(want[0], "~"),
							strings.Split(parsed.Get(k), "~"),
							"param %q", k)
						continue
					}
					assert.Equal(t, want, parsed[k], "param %q", k)
				}
			}
		})
	}
}

// TestEncode_Body ensures overflow parameters are moved into a properly
// percent-encoded POST body.
func TestEncode_Body(t *testing.T) {
	t.Parallel()
	testingx.Tags(t, tagx.Short)

	// force an overflow so that Encode has to spill into the body
	remain := map[string]any{
		"filler": strings.Repeat("x", 2048),
		// value contains reserved characters that must not leak into the wire format
		"pr2":     "brWine in a Box~id128573~nmPersonalisierbare Weinkiste Moët & Chandon Rosé Impérial 75cl~pr73.91304347826087~qt1",
		"ep.q":    "a=b&c=d",
		"ep.plus": "1+2 3",
	}
	event := gtag.Payload{Remain: remain}

	query, body, err := gtag.Encode(&event)
	require.NoError(t, err)

	encoded := readBody(t, body)

	// both query and body must be parseable as query strings
	values, err := url.ParseQuery(query)
	require.NoError(t, err)
	parsed, err := url.ParseQuery(encoded)
	require.NoError(t, err)

	// nothing may be lost or duplicated between query and body
	for k, want := range map[string]string{
		"pr2":     remain["pr2"].(string),
		"ep.q":    remain["ep.q"].(string),
		"ep.plus": remain["ep.plus"].(string),
	} {
		if _, inQuery := values[k]; inQuery {
			assert.Equal(t, want, values.Get(k), "query param %q", k)
			continue
		}
		assert.Equal(t, want, parsed.Get(k), "body param %q round-trips", k)
	}

	// the reserved characters must be escaped, not emitted raw
	assert.NotContains(t, encoded, "& Chandon", "raw & must be escaped in body")
	assert.NotContains(t, encoded, "a=b&c=d", "raw &/= must be escaped in body")
	// spaces are encoded as %20, never as "+", to match EncodeValues
	assert.NotContains(t, encoded, "+", "spaces must be %20 and literal + must be %2B")
}

// TestEncode_BodyEscapesAmpersand is a regression test: a value containing "&"
// used to be written to the body unescaped, which split it into a bogus extra
// parameter and truncated the value.
func TestEncode_BodyEscapesAmpersand(t *testing.T) {
	t.Parallel()
	testingx.Tags(t, tagx.Short)

	const want = "Moët & Chandon Rosé Impérial 75cl"

	event := gtag.Payload{
		Remain: map[string]any{
			"filler": strings.Repeat("x", 2048),
			"pr2":    want,
		},
	}

	query, body, err := gtag.Encode(&event)
	require.NoError(t, err)

	encoded := readBody(t, body)
	values, err := url.ParseQuery(query)
	require.NoError(t, err)
	parsed, err := url.ParseQuery(encoded)
	require.NoError(t, err)

	// combine query + body, mirroring what the receiving endpoint sees
	got := values.Get("pr2")
	if got == "" {
		got = parsed.Get("pr2")
	}
	assert.Equal(t, want, got)

	// no bogus parameter created by an unescaped "&"
	for k := range parsed {
		assert.NotContains(t, k, " ", "unescaped & created bogus key %q", k)
	}
}

// TestEncode_NoBodyWhenShort verifies that a short payload produces no body.
func TestEncode_NoBodyWhenShort(t *testing.T) {
	t.Parallel()
	testingx.Tags(t, tagx.Short)

	event := gtag.Payload{Remain: map[string]any{"en": "page_view"}}

	query, body, err := gtag.Encode(&event)
	require.NoError(t, err)
	assert.Nil(t, body)
	assert.Equal(t, "en=page_view", query)
}

// TestEncode_RichsstsseStaysLast verifies richsstsse remains the last query
// parameter and is never spilled into the body.
func TestEncode_RichsstsseStaysLast(t *testing.T) {
	t.Parallel()
	testingx.Tags(t, tagx.Short)

	richsstsse := ""
	event := gtag.Payload{
		Richsstsse: &richsstsse,
		Remain: map[string]any{
			"filler": strings.Repeat("x", 2048),
		},
	}

	query, body, err := gtag.Encode(&event)
	require.NoError(t, err)

	assert.True(t, strings.HasSuffix(query, "richsstsse"), "got: %s", query)
	assert.Equal(t, 1, strings.Count(query, "richsstsse"))
	// bare flag, never "richsstsse="
	assert.NotContains(t, query, "richsstsse=")

	if body != nil {
		assert.NotContains(t, readBody(t, body), "richsstsse")
	}
}
