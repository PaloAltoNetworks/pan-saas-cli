# SaaS-CLI
This a CLI application that interacts with a Palo Alto Networks Firewall to pull the [SaaS Application Usage Report](https://docs.paloaltonetworks.com/pan-os/10-1/pan-os-admin/monitoring/view-and-manage-reports/view-reports.html#id12e5da3d-d44f-4c9e-9d97-8704151ed103) and lists all the SaaS applicaiton in use. The applications also can tag the predefined Sanctioned tag to applications.

## How to run the program
* Download the minicli executable
* Run ./minicli to see all the avalable options
* Run ./minicli key to generate an API Key
* Run ./minicli dsa to display the SaaS applications
* Run ./minicli wsf to create a file with the SaaS applications
* Run ./minicli ast to apply sanctioned tag to applications

## Files contained in this folder:
* cmd: Contains application entry code
* examplesecret: Contains an example how to hardcode your encryption/decryption key to be used by the secret.Secret()
* pkg: Contains source code

## Consideration
Start the application by generating the API Key.


# SaaS-CSV-Report-Parser
This script takes a CSV file "ExampleSaaSApps.csv" in this example that was generated from running a [SaaS Application Usage Report](https://docs.paloaltonetworks.com/pan-os/10-1/pan-os-admin/monitoring/view-and-manage-reports/view-reports.html#id12e5da3d-d44f-4c9e-9d97-8704151ed103) on a Palo Alto Networks Firewall.

The output is a text file with only the Applications names, in this example "ExampleSaaSApps.txt"

## Prerequisites
You need to have Go installed in your system. If you don't have Go please visit (https://go.dev/dl/).

## How to run the program
* Navigate to the directory containing the main.go code
* Run "go run main.go"


## Consideration
In this example the CSV file is named "SaaSApps.csv". Please update the code to match your file or rename your file to "SaaSApps.csv" if you don't want to modify the code.