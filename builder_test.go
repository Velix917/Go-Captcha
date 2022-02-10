package gocaptcha

import "testing"

func TestBuilder(t *testing.T) {
	captcha := NewRecaptchaV2()
	captcha.SetUrl("google.com")
	captcha.SetSiteKey("Some-site-key")
	captcha.SetInvisible(true)
	captcha.SetAction("verify")

	got := captcha.ToString()
	want := "googlekey=Some-site-key&invisible=1&method=userrecaptcha&pageurl=google.com&sitekey=verify"

	if got != want {
		t.Errorf("error: got %s, want %s", got, want)
	}
}
