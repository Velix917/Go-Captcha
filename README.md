# Go-Captcha
Another 2captcha wrapper but made in Golang. Currently only supports google recaptcha v2

Usage
-----

```golang
client := NewClient("Your-api-key")

captcha := NewRecaptchaV2()
captcha.SetSiteKey("6LeIxboZAAAAAFQy7d8GPzgRZu2bV0GwKS8ue_cH")
captcha.SetUrl("https://2captcha.com/demo/recaptcha-v2")
captcha.SetAction("verify")

res, _ := client.Solve(captcha)
fmt.Println(res)
```

License
-------
Project under MIT license.