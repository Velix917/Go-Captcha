package gocaptcha

import (
	"net/url"
)

// **** Structs ***** //

type ICaptchaRequest interface {
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

// **** Initializers **** //

func NewRecaptchaV2() *ReCaptchaV2 {
	recaptchav2 := &ReCaptchaV2{}
	recaptchav2.Params = make(url.Values)
	recaptchav2.Params.Add("method", "userrecaptcha")
	return recaptchav2
}

func NewFuncaptcha() *FunCaptcha {
	funcaptcha := &FunCaptcha{}
	funcaptcha.Params = make(url.Values)
	funcaptcha.Params.Add("method", "funcaptcha")
	return funcaptcha
}

func NewCapyCaptcha() *CapyCaptcha {
	capycaptcha := &CapyCaptcha{}
	capycaptcha.Params = make(url.Values)
	capycaptcha.Params.Add("method", "capy")
	return capycaptcha
}

func NewTikTokCaptcha() *TikTokCaptcha {
	tiktokcaptcha := &TikTokCaptcha{}
	tiktokcaptcha.Params = make(url.Values)
	tiktokcaptcha.Params.Add("method", "tiktok")
	return tiktokcaptcha
}

// **** Methods ***** //

func (c *CaptchaRequest) ToString() string       { return c.Params.Encode() }
func (c *CaptchaRequest) Get(k string) string    { return c.Params.Get(k) }
func (c *CaptchaRequest) Add(k string, v string) { c.Params.Add(k, v) }

func (r *ReCaptchaV2) SetSiteKey(sitekey string) { r.CaptchaRequest.Params.Add("googlekey", sitekey) }
func (r *ReCaptchaV2) SetUrl(uri string)         { r.CaptchaRequest.Params.Add("pageurl", uri) }
func (r *ReCaptchaV2) SetInvisible(inv bool)     { r.CaptchaRequest.Params.Add("invisible", boolConv(inv)) }
func (r *ReCaptchaV2) SetAction(action string)   { r.CaptchaRequest.Params.Add("sitekey", action) }
func (r *ReCaptchaV2) SetProxy(proxy string)     { r.CaptchaRequest.Params.Add("proxy", proxy) }

func (f *FunCaptcha) SetPublicKey(pkey string)   { f.CaptchaRequest.Params.Add("publickey", pkey) }
func (f *FunCaptcha) SetServerUrl(surl string)   { f.CaptchaRequest.Params.Add("surl", surl) }
func (f *FunCaptcha) SetPageUrl(purl string)     { f.CaptchaRequest.Params.Add("pageurl", purl) }
func (f *FunCaptcha) SetCustomData(cdata string) { f.CaptchaRequest.Params.Add("data[key]", cdata) }

func (c *CapyCaptcha) SetCaptchaKey(ckey string) { c.CaptchaRequest.Params.Add("captchakey", ckey) }
func (c *CapyCaptcha) SetApiServer(aserv string) { c.CaptchaRequest.Params.Add("api_server", aserv) }
func (c *CapyCaptcha) SetVersion(ver string)     { c.CaptchaRequest.Params.Add("version", ver) }
func (c *CapyCaptcha) SetPageUrl(purl string)    { c.CaptchaRequest.Params.Add("pageurl", purl) }

func (t *TikTokCaptcha) SetAid(aid string)       { t.CaptchaRequest.Params.Add("aid", aid) }
func (t *TikTokCaptcha) SetHost(host string)     { t.CaptchaRequest.Params.Add("host", host) }
func (t *TikTokCaptcha) SetPageUrl(purl string)  { t.CaptchaRequest.Params.Add("pageurl", purl) }
func (t *TikTokCaptcha) SetCookies(ckies string) { t.CaptchaRequest.Params.Add("cookies", ckies) }
