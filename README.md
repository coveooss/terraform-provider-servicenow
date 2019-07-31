# Terraform ServiceNow Provider

A custom provider for Terraform to manage objects in a ServiceNow instance for developping system applications outside of an instance. This is especially useful when you want to create an Application with proper source control and continuous development integration.

<img src="https://www.terraform.io/assets/images/og-image-8b3e4f7d.png" height="200" alt="Terraform Logo"/><img src="https://community.servicenow.com/c4fe846adbb95f0037015e77dc961918.iix" height="200" alt="ServiceNow Logo"/>

[![Travis Report](https://travis-ci.org/coveooss/terraform-provider-servicenow.svg?branch=master)](https://travis-ci.org/coveooss/terraform-provider-servicenow) [![Coverage Status](https://coveralls.io/repos/github/coveooss/terraform-provider-servicenow/badge.svg)](https://coveralls.io/github/coveooss/terraform-provider-servicenow)
[![Go Report Card](https://goreportcard.com/badge/github.com/coveooss/terraform-provider-servicenow)](https://goreportcard.com/report/github.com/coveooss/terraform-provider-servicenow) 
[![Release Badge](https://img.shields.io/github/release/coveooss/terraform-provider-servicenow.svg)](https://github.com/coveooss/terraform-provider-servicenow/releases/latest)
[![License Badge](https://img.shields.io/github/license/coveooss/terraform-provider-servicenow.svg)](LICENSE)

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

### Other

You can also download the [latest release](https://github.com/coveooss/terraform-provider-servicenow/releases) binaries and place them in your working directory, since Terraform will look for providers in the working directory also.

## Supported Resources

Check out the [Wiki](https://github.com/coveooss/terraform-provider-servicenow/wiki) !