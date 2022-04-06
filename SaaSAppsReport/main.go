/*
"""
SaaS Applications Generator

The program runs a SaaS Application Report on a Palo Alto Networks firewall, parses the data, and outputs a text file.

These applications can then be reviewed and tagged as Sanctioned where applicable.

PLEASE NOTE:  This script is still a Work in Progress and may contain bugs.

Author: Zack Macharia
Email: zmacharia@paloaltonetworks.com
© 2019 Palo Alto Networks, Inc.  All rights reserved.
Licensed under SCRIPT SOFTWARE AGREEMENT, Palo Alto Networks, Inc.,
at https://www.paloaltonetworks.com/legal/script-software-license-1-0.pdf
"""
*/
package main

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

//var: Holds environmental variable and a string from the getApiKey function
var (
	Firewall = os.Getenv("FIREWALL")
	ApiKey   = getApiKey()
)

//Report: Parent tag of the XML document
type Report struct {
	Report string `xml:"report"`
	Result Result `xml:"result"`
}

/*Result: Child to the Report tag. 
There are multiple Entry tags within the Result tag; hence set as an array*/
type Result struct {
	Result string  `xml:"result"`
	Entry  []Entry `xml:"entry"`
}

/*Entry: Child to the Result tag.
There are multiple Entry tags in the document*/
type Entry struct {
	Subcategory      string `xml:"subcategory-of-name"`
	Name             string `xml:"name"`
	Bytes            int    `xml:"nbytes"`
	NumberOfSessions int    `xml:"nsess"`
	NumberOfThreats  int    `xml:"nthreats"`
}

//main: Program's entry point
func main() {

	data, err := saasReport()
	if err != nil {
		log.Fatal(err)
	}
	fname := createSaaSAppsFile(data)
	displayFileData(fname.Name())

}

/*Client: Defines a client that connects to a device
This function is used by other functions where different
path URL is passed depending on the use case*/
func Client(path string) *http.Response {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore invalid SSL certs
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://" + Firewall + path)

	if err != nil {
		log.Fatal(err)
	}
	return resp
}

//getApiKey: This function connects to a device and generates an API Key
func getApiKey() string {
	var (
		Username = os.Getenv("PANOUSER")
		Password = os.Getenv("PANOPWD")
	)

	path := "/api/?type=keygen&user=" + Username + "&password=" + Password //path to retrieve API Key

	resp := Client(path)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("<key>.*</key>")
	keyTag := re.FindString(string(body))
	key := strings.TrimRight(strings.TrimLeft(keyTag, "<key>"), "</key>")

	return key
}

/*saasReport: This function makes an API call to the device
 and pulls down SaaS report data in bytes form*/
func saasReport() ([]byte, error) {

	path := "/api/?key=" + ApiKey + "&type=report&async=yes&reporttype=predefined&reportname=SaaS+Application+Usage"

	resp := Client(path)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

/*createSaaSAppsFile: Parses the SaaS report data pulled from the device.
Writes the application names from the report to a text file.
Text file has the title of "SaaSApps_created:XXXX" 
Where XXXX is the current date and time.*/
func createSaaSAppsFile(data []byte) *os.File {

	var r Report
	data, _ = saasReport()

	xml.Unmarshal(data, &r)

	fname := "SaaSApps_created:" + time.Now().Format("2006-01-02 15:04") + ".txt"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range r.Result.Entry {
		f.WriteString(s.Name + "\n")
	}
	return f
}

//displayFileData: Opens a file for reading
//reads the file's content and displays it on terminal
func displayFileData(fname string) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Println(string(data))
}

