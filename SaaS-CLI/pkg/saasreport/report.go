package saasreport

import (
	"bufio"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/zackmacharia/example/pkg/crypto"
	"github.com/zackmacharia/example/pkg/fwinfo"
)

//Report: Parent tag of the firewall SaaS Report XML file
type Report struct {
	Report string `xml:"report"`
	Result Result `xml:"result"`
}

//Result: Child to the Report tag.Multiple Entry tags exist within the Result tag; hence set as an array
type Result struct {
	Result string  `xml:"result"`
	Entry  []Entry `xml:"entry"`
}

//Entry: Child to the Result tag
type Entry struct {
	Subcategory      string `xml:"subcategory-of-name"`
	Name             string `xml:"name"`
	Bytes            int    `xml:"nbytes"`
	NumberOfSessions int    `xml:"nsess"`
	NumberOfThreats  int    `xml:"nthreats"`
}

//Client: Defines transport and protocol to connect to the device
func Client(path string) *http.Response {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, //ignore invalid SSL certs
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://" + fwinfo.Hostname() + path)

	if err != nil {
		log.Fatal(err)
	}
	return resp
}

//DisplaySaaSReport: Displays the SaaS applications on the terminal
func DisplaySaaSReport() {

	var r Report

	path := "/api/?key=" + crypto.Decrypt() + "&type=report&async=yes&reporttype=predefined&reportname=SaaS+Application+Usage"

	resp := Client(path)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	xml.Unmarshal(body, &r)

	for i, s := range r.Result.Entry {
		fmt.Println(i+1, s.Name) //Display index start at 1 instead of 0
		i++
	}
}

/*PullSaaSReport: This function makes an API call to the device
and pulls down SaaS report data in bytes form*/
func PullSaaSReport() ([]byte, error) {

	path := "/api/?key=" + crypto.Decrypt() + "&type=report&async=yes&reporttype=predefined&reportname=SaaS+Application+Usage"

	resp := Client(path)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

/*CreateSaaSAppsFile: Parses the SaaS report data pulled from the device.
Writes the application names from the report to a text file.
Text file has the title of "SaaSApps_created:XXXX"
Where XXXX is the current date and time.*/
func CreateSaaSAppsFile(data []byte) *os.File {

	var r Report
	data, _ = PullSaaSReport()

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

//AddSanctionedTag: Adds the predefined 'Sanctioned' tag to an application
func AddSanctionedTag() {

	fmt.Print("Enter File Name full path e.g 'C:\\Documents\\myfile.txt': ")
	var fname string
	fmt.Scanln(&fname)

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal("Unable to read", fname, err)
	}
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		app := scanner.Text()
		path := "/api/?key=" + crypto.Decrypt() + "&type=config&action=set&xpath=/config/devices/entry[@name='localhost.localdomain']/vsys/entry[@name='vsys1']/application-tag/entry[@name=" + "'" + app + "'" + "]/tag&element=<member>Sanctioned</member>"
		resp := Client(path)

		defer resp.Body.Close()

		fmt.Println("Sanctioned Tag Added to: ", app)
	}



}
