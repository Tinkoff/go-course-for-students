package lecture10

import (
	"net/http"
)

func HTTPReq(addr string) (string, error) {
	var body []byte

	resp, err := http.DefaultClient.Get(addr)
	if err != nil {
		return "", err
	}

	defer func() { _ = resp.Body.Close() }()

	_, err = resp.Body.Read(body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
