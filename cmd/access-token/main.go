package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"math/big"
	"net/http"
	u "net/url"
	"os"
	"strings"

	"crypto/tls"

	"github.com/tidwall/gjson"
)

const codeAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~"
const codeLength = 128

func generateCodeChallenge() (map[string]string, error) {
	code := make([]byte, codeLength)
	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(codeAlphabet))))
		if err != nil {
			return nil, err
		}
		code[i] = codeAlphabet[num.Int64()]
	}

	hash := sha256.Sum256(code)
	hashBase64 := base64.RawURLEncoding.EncodeToString(hash[:])

	fmt.Printf("CODE: %s\nHASH: %s\n", string(code), hashBase64)

	return map[string]string{
		"code": string(code),
		"hash": hashBase64,
	}, nil
}

func getOauthCode(host, hash string) (string, error) {
	url := fmt.Sprintf("https://%s:8443/v1/oauth/authorize?audience=homesmart.local&response_type=code&code_challenge=%s&code_challenge_method=S256", host, hash)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	code := gjson.Get(string(body), "code").String()
	fmt.Printf("oauth code: %s\n", code)
	return code, nil
}

func getOauthToken(host, code, challenge string) (string, error) {
	url := fmt.Sprintf("https://%s:8443/v1/oauth/token", host)

	data := u.Values{}
	data.Set("code", code)
	data.Set("name", hostname())
	data.Set("grant_type", "authorization_code")
	data.Set("code_verifier", challenge)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	// fmt.Println(string(body))

	token := gjson.Get(string(body), "access_token").String()
	// fmt.Printf("oauth token: %s\n", token)

	return token, nil
}

func main() {
	host := flag.String("host", "", "Dirigera hub hostname or IP address (required)")
	flag.Parse()
	if *host == "" {
		fmt.Fprintln(os.Stderr, "Error: -host flag is required (e.g. -host 192.168.1.x or -host my-hub.local)")
		os.Exit(1)
	}
	dirigeraHost := *host

	codeChallenge, err := generateCodeChallenge()
	if err != nil {
		panic(err)
	}

	code, err := getOauthCode(dirigeraHost, codeChallenge["hash"])
	if err != nil {
		panic(err)
	}

	fmt.Println("Press the action button on the dirigera and then press enter")
	fmt.Scanln()

	token, err := getOauthToken(dirigeraHost, code, codeChallenge["code"])
	if err != nil {
		panic(err)
	}
	fmt.Printf("access token: %s\n", token)
}

func hostname() string {
	host, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return host
}
