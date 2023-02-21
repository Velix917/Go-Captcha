package gocaptcha

import "testing"

func TestBuilder(t *testing.T) {
	captcha := NewRecaptchaV2()
	captcha.FSetUrl("google.com")
	captcha.FSetSiteKey("Some-site-key")
	captcha.FSetInvisible(true)
	captcha.FSetAction("verify")

	got := captcha.ToString()
	want := "googlekey=Some-site-key&invisible=1&method=userrecaptcha&pageurl=google.com&sitekey=verify"

	if got != want {
		t.Errorf("error: got %s, want %s", got, want)
	}
}
