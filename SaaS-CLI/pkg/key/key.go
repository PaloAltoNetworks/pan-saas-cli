package apikey

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/zackmacharia/PANOS-GOLANG/SaaS-CLI/pkg/crypto"
	"github.com/zackmacharia/PANOS-GOLANG/SaaS-CLI/pkg/fwinfo"
)

/* GetApiKey: takes user input use provided information to generate an api key.
The APIKey is then encrypted and stored in a file named key.data. The device
hostname is also written to a file(host.data) in clear text.*/
func GetApiKey() {

	fmt.Print("Firewall IP Address or FQDN: ")
	var firewall string
	fmt.Scanln(&firewall)

	fmt.Print("Username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		log.Fatal(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore invalid SSL certs
	}
	client := &http.Client{Transport: tr}

	path := "/api/?type=keygen&user=" + username + "&password=" + string(password) //path to retrieve API Key

	resp, err := client.Get("https://" + firewall + path)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// reads the http body from the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// grep only the part with the API Key string
	re := regexp.MustCompile("<key>.*</key>")
	keyTag := re.FindString(string(body))
	APIKEY := strings.TrimRight(strings.TrimLeft(keyTag, "<key>"), "</key>")

	// write firewall name to a file
	fwinfo.WriteHostname(firewall)

	// Encrypt and write Key to a file
	a := crypto.Encrypt(APIKEY)
	fwinfo.WriteAPIKey(a)

}
