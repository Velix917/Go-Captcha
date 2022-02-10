# Go-Captcha
Another 2captcha wrapper but made in Golang.

Currently only supports google recaptcha v2 and funcaptcha (the two ones i mostly use). For solving other captchas use the official library (which i recently realized it already exists).

Usage
-----

### Basic example for Google Recaptcha v2

```golang
package main

import (
    "log"
    "fmt"
    "github.com/Grapphy/Go-Captcha"
)

func main() {
    client := gocaptcha.NewClient("Your-api-key")

    captcha := gocaptcha.NewRecaptchaV2()
    captcha.SetSiteKey("6LeIxboZAAAAAFQy7d8GPzgRZu2bV0GwKS8ue_cH")
    captcha.SetUrl("https://2captcha.com/demo/recaptcha-v2")
    captcha.SetAction("verify")

    solution, err := client.Solve(captcha)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(solution)
}
```

### Funcaptcha

```golang
captcha := NewFuncaptcha()
captcha.SetPublicKey("20782B4C-05D0-45D7-97A0-41641055B6F6")
captcha.SetPageUrl("https://github.com/signup")
captcha.SetServerUrl("https://api.funcaptcha.com")
```


License
-------
Project under MIT license.