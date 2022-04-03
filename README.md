# SaaSAppsReport
This folder contains a program that interacts with a Palo Alto Networks Firewall to pull the [SaaS Application Usage Report](https://docs.paloaltonetworks.com/pan-os/10-1/pan-os-admin/monitoring/view-and-manage-reports/view-reports.html#id12e5da3d-d44f-4c9e-9d97-8704151ed103).

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
