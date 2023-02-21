package gocaptcha

import (
	"net/url"
)

type ICaptchaRequest interface {
	MakeParams()
	ToString() string
	Get(string) string
	Add(string, string)
}

type CaptchaRequest struct {
	Id     string
	Code   string
	Params url.Values
}

type ReCaptchaV2 struct{ CaptchaRequest }
type FunCaptcha struct{ CaptchaRequest }
type CapyCaptcha struct{ CaptchaRequest }
type TikTokCaptcha struct{ CaptchaRequest }

func BuildCaptcha(captcha ICaptchaRequest, method string) ICaptchaRequest {
	captcha.MakeParams()
	captcha.Add("method", method)
	return captcha
}

func NewRecaptchaV2() *ReCaptchaV2 {
	return BuildCaptcha(&ReCaptchaV2{}, "userrecaptcha").(*ReCaptchaV2)
}

func NewFuncaptcha() *FunCaptcha {
	return BuildCaptcha(&FunCaptcha{}, "funcaptcha").(*FunCaptcha)
}

func NewCapyCaptcha() *CapyCaptcha {
	return BuildCaptcha(&CapyCaptcha{}, "capy").(*CapyCaptcha)
}

func NewTikTokCaptcha() *TikTokCaptcha {
	return BuildCaptcha(&TikTokCaptcha{}, "tiktok").(*TikTokCaptcha)
}

func (c *CaptchaRequest) MakeParams()            { c.Params = make(url.Values) }
func (c *CaptchaRequest) ToString() string       { return c.Params.Encode() }
func (c *CaptchaRequest) Get(k string) string    { return c.Params.Get(k) }
func (c *CaptchaRequest) Add(k string, v string) { c.Params.Add(k, v) }

func (r *ReCaptchaV2) FSetSiteKey(sitekey string) { r.CaptchaRequest.Params.Add("googlekey", sitekey) }
func (r *ReCaptchaV2) FSetUrl(uri string)         { r.CaptchaRequest.Params.Add("pageurl", uri) }
func (r *ReCaptchaV2) FSetInvisible(inv bool) {
	r.CaptchaRequest.Params.Add("invisible", boolConv(inv))
}
func (r *ReCaptchaV2) FSetAction(action string) { r.CaptchaRequest.Params.Add("sitekey", action) }
func (r *ReCaptchaV2) FSetProxy(proxy string)   { r.CaptchaRequest.Params.Add("proxy", proxy) }

func (f *FunCaptcha) FSetPublicKey(pkey string)   { f.CaptchaRequest.Params.Add("publickey", pkey) }
func (f *FunCaptcha) FSetServerUrl(surl string)   { f.CaptchaRequest.Params.Add("surl", surl) }
func (f *FunCaptcha) FSetPageUrl(purl string)     { f.CaptchaRequest.Params.Add("pageurl", purl) }
func (f *FunCaptcha) FSetCustomData(cdata string) { f.CaptchaRequest.Params.Add("data[blob]", cdata) }

func (c *CapyCaptcha) FSetCaptchaKey(ckey string) { c.CaptchaRequest.Params.Add("captchakey", ckey) }
func (c *CapyCaptcha) FSetApiServer(aserv string) { c.CaptchaRequest.Params.Add("api_server", aserv) }
func (c *CapyCaptcha) FSetVersion(ver string)     { c.CaptchaRequest.Params.Add("version", ver) }
func (c *CapyCaptcha) FSetPageUrl(purl string)    { c.CaptchaRequest.Params.Add("pageurl", purl) }

func (t *TikTokCaptcha) FSetAid(aid string)       { t.CaptchaRequest.Params.Add("aid", aid) }
func (t *TikTokCaptcha) FSetHost(host string)     { t.CaptchaRequest.Params.Add("host", host) }
func (t *TikTokCaptcha) FSetPageUrl(purl string)  { t.CaptchaRequest.Params.Add("pageurl", purl) }
func (t *TikTokCaptcha) FSetCookies(ckies string) { t.CaptchaRequest.Params.Add("cookies", ckies) }
