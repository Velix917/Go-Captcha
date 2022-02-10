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

// **** Initializers **** //

func NewRecaptchaV2() *ReCaptchaV2 {
	recaptchav2 := &ReCaptchaV2{}
	recaptchav2.Params = make(url.Values)
	recaptchav2.Params.Add("method", "userrecaptcha")
	return recaptchav2
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
