package gocaptcha

import "testing"

func TestClient(t *testing.T) {
	client := NewClient("Your-key", "https://api.funpass.services/in.php", "https://api.funpass.services/res.php")

	captcha := NewRecaptchaV2()
	captcha.FSetSiteKey("6LeIxboZAAAAAFQy7d8GPzgRZu2bV0GwKS8ue_cH")
	captcha.FSetUrl("https://2captcha.com/demo/recaptcha-v2")
	captcha.FSetAction("verify")

	_, err := client.Solve(captcha)
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
}
