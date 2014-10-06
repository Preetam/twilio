package twilio

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Text(account, token, to, from, body string) (io.ReadCloser, error) {
	request, err := http.NewRequest("POST", "https://api.twilio.com/2010-04-01/Accounts/"+account+"/Messages.json",
		strings.NewReader((url.Values{
			"To":   []string{to},
			"From": []string{from},
			"Body": []string{body},
		}).Encode()))

	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(account, token)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(request)

	return res.Body, err
}
