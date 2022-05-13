# SaaS-CLI
<<<<<<< HEAD
This a CLI application that interacts with a Palo Alto Networks Firewall to pull the [SaaS Application Usage Report](https://docs.paloaltonetworks.com/pan-os/10-1/pan-os-admin/monitoring/view-and-manage-reports/view-reports.html#id12e5da3d-d44f-4c9e-9d97-8704151ed103) and lists all the SaaS applicaiton in use. The applications also can tag the predefined Sanctioned tag to applications.

## Prerequisites
* [Install Go](https://go.dev/dl/). This will allow you to edit the Secret function with a key of your choosing that will be used to encrypt and decrypt the API Key.
For security reasons the key is not hard-coded in the code.
* Clone the repo
* Rename the "examplesecret" folder to "secret"
* Inside the secret folder update the package name from "examplesecret" to "secret"
* Change the string returned by the Secret function


## How to run the program
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

=======
SaaS-CLI is a utility that interacts with a Palo Alto Networks Firewall to pull the [SaaS Application Usage Report](https://docs.paloaltonetworks.com/pan-os/10-1/pan-os-admin/monitoring/view-and-manage-reports/view-reports.html#id12e5da3d-d44f-4c9e-9d97-8704151ed103) and outputs all the discovered SaaS applications which you can then review and tag them as Sanctioned if applicable.

## Files contained in this folder:
* cmd - Contains the application entry point code
* pkg - Contains multiple packages used by the application (The secret package is not included in this repository for security reasons).
* saascli-linux - Executable for Linux systems
* saascli-windows - Executable for Windows systems

## Download and Installation
* For Windows users download the "saascli-windows" executable and rename it by adding a ".exe" extension
* For Linux users download the "saascli-linux" executable
* Windows run "saas-cli-windows.exe"
* Linux run "./saas-cli-linux" 
* For Linux users you might have to give the application executable rights
* Run saas-cli-your_operating_system -h to list all the supported commands


## Optional - Encryption/Decryption Key Details
The executable is packaged with a 32-byte array used to encrypt and decrypt the API key. If you want to change this key and recompile the executable use the following steps.
*[Install Go](https://go.dev/doc/install) (if you don't have it already)
* Clone the repository
* Create a secret folder under pkg
* Create a secret.go file inside the secret folder
* Create a function that returns a 32-byte array. Example below:
```
package secret

// Secret: returns a 32 byte array used as the encryption and decryption key. This can be any 32byte array.
func Secret() []byte {
	return []byte("thisshouldbeany32bytearrayyoudig") // CHANGE THIS VALUE TO YOUR OWN VALUE
}
```
* Build the executable 
```
env GOOS=target-OS GOARCH=target-architecture go build package-import-path

Example: Linux OS on amd64 Platform

env GOOS=linux GOARCH=amd64 go build -o saascli-linux cmd/saascli.go
```

# SaaS-CSV-Report-Parser
This script takes a CSV file "ExampleSaaSApps.csv" in this example that was generated from running a [SaaS Application Usage Report](https://docs.paloaltonetworks.com/pan-os/10-1/pan-os-admin/monitoring/view-and-manage-reports/view-reports.html#id12e5da3d-d44f-4c9e-9d97-8704151ed103) on a Palo Alto Networks Firewall.

The output is a text file with only the Applications names, in this example "ExampleSaaSApps.txt"

## Prerequisites
You need to have Go installed in your system. If you don't have Go please visit (https://go.dev/dl/).

## How to run the program
* Navigate to the directory containing the main.go code
* Run "go run main.go"

>>>>>>> dev

## Consideration
In this example the CSV file is named "SaaSApps.csv". Please update the code to match your file or rename your file to "SaaSApps.csv" if you don't want to modify the code.