# SaaS-CLI
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

## Support
Community Supported

The software and templates in the repo are released under an as-is, best effort,
support policy. This software should be seen as community supported and Palo
Alto Networks will contribute our expertise as and when possible. We do not
provide technical support or help in using or troubleshooting the components of
the project through our normal support options such as Palo Alto Networks
support teams, or ASC (Authorized Support Centers) partners and backline support
options. The underlying product used (the VM-Series firewall) by the scripts or
templates are still supported, but the support is only for the product
functionality and not for help in deploying or using the template or script
itself. Unless explicitly tagged, all projects or work posted in our GitHub
repository (at https://github.com/PaloAltoNetworks) or sites other than our
official Downloads page on https://support.paloaltonetworks.com are provided
under the best effort policy.