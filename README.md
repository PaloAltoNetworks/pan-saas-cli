# SaaSAppsReport
This folder contains a program that interacts with a Palo Alto Networks Firewall to pull the [SaaS Application Usage Report](https://docs.paloaltonetworks.com/pan-os/10-1/pan-os-admin/monitoring/view-and-manage-reports/view-reports.html#id12e5da3d-d44f-4c9e-9d97-8704151ed103) and outputs all the discovered SaaS applications which you can then review and tag them as Sanctioned if applicable.

## How to run the program
* [Install Go](https://go.dev/doc/install) (if you don't have it already)
* Import the project
* Issue "go run main.go" on CLI
* Review generated file

## Files contained in this folder:
* sample.xml - Represents the format of the raw XML file
* SaaSApps_created: XXXX - Output file after running this program; XXXX reflects the current time in the system
* main.go - Program Source Code

## Consideration
Edit your environmental variables for the following strings:
* Username
* PANOUSER - Firewall's Username
* PANOPWD - Firewall's Password
* FIREWALL - Firewall's IP Address or FQDN
