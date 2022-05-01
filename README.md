# SaaS-CLI
This folder contains a program that interacts with a Palo Alto Networks Firewall to pull the [SaaS Application Usage Report](https://docs.paloaltonetworks.com/pan-os/10-1/pan-os-admin/monitoring/view-and-manage-reports/view-reports.html#id12e5da3d-d44f-4c9e-9d97-8704151ed103) and outputs all the discovered SaaS applications which you can then review and tag them as Sanctioned if applicable.

## How to run the program
* [Install Go](https://go.dev/doc/install) (if you don't have it already)
* Import the project
* Issue "go run main.go" on CLI
* Review generated file

## Files contained in this folder:
* sample.xml - Example file showing the raw format of the report.
* SaaSApps_created:XXXX - Example file showing the output after running this program; XXXX reflects the current time in the system
* main.go - Program Source Code

## Consideration
Add/Edit the following strings to your systems environmental variables:
* Username
* PANOUSER - Firewall's Username
* PANOPWD - Firewall's Password
* FIREWALL - Firewall's IP Address or FQDN
