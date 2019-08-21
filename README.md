# Nexus CLI

[![DepShield Badge](https://depshield.sonatype.org/badges/sonatype-nexus-community/nexus-cli/depshield.svg)](https://depshield.github.io)

*** This project is in its early stages of development and configuration.

## Overview

`nexus` is a command line tool used to interact with Nexus IQ and Repository Manager. Our intention is that it will help expand on functionality currently available in the existing Java based [Nexus IQ CLI](https://help.sonatype.com/integrations/nexus-iq-cli). Every effort has been made to ensure the commands and flags exposed are clear, understandable, and unambiguous.

We've written this utility in `go` so it can be compiled for multiple platforms and doesn't require any special runtime. This is especially important for integration with other tools or shell stages in a pipeline.

## Uses

The following scenarios are but a few examples of what can be done with this CLI:

* Writing infrastructure code to automatically install licenses
* Managing users and groups
* Scanning artifacts for known vulnerabilities from a command line
* More!

## Installation

* Build from source
* Download release binary from [here on GitHub](https://github.com/sonatype-nexus-community/nexus-cli/releases)

### Build from source

//TODO: Instructions

### Download release binary

//TODO: Instructions

## Development

`nexus` is written using Golang 1.12.7 and uses `go mod` for dependencies.

Tests can be run like `go test ./... -v`

## The Fine Print

It is worth noting that this is **NOT SUPPORTED** by [Sonatype](//www.sonatype.com), and is a contribution  to the open source community (read: you!)

Remember:

- Use this contribution at the risk tolerance that you have
- Do **NOT** file Sonatype support tickets related to this
- **DO** file issues here on GitHub, so that the community can pitch in

## Getting help

Looking to contribute to our code but need some help? There's a few ways to get information:

* Chat with us on [Gitter](https://gitter.im/sonatype/nexus-developers)

