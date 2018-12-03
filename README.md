# Terraform ServiceNow Provider

A custom provider for Terraform to manage objects in a ServiceNow instance for developping system applications outside of an instance.

[![Terraform Logo](https://www.terraform.io/assets/images/og-image-f5bbc98c.png)](https://www.terraform.io/)
[![ServiceNow Logo](https://community.servicenow.com/c4fe846adbb95f0037015e77dc961918.iix)](https://www.servicenow.com/)

[![Travis Report](https://travis-ci.org/coveo/terraform-provider-servicenow.svg?branch=master)](https://travis-ci.org/coveo/terraform-provider-servicenow)
[![Go Report Card](https://goreportcard.com/badge/github.com/coveo/terraform-provider-servicenow)](https://goreportcard.com/report/github.com/coveo/terraform-provider-servicenow)

## Requirements

- [Terraform](https://www.terraform.io/downloads.html)
- [Go](https://golang.org/doc/install) (to build the provider plugin)

## Installation

### Windows

1. Clone repository to: `%GOPATH%/src/github.com/terraform-providers/terraform-provider-servicenow`
1. Build the executable using `go build -o terraform-provider-servicenow.exe`
1. Copy the file to `%APPDATA%\terraform.d\plugins`

### Linux

1. Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-servicenow`
1. Build the executable using `go build -o terraform-provider-servicenow`
1. Copy the file to `~.terraform.d/plugins`
