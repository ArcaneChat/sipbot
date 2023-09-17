#  sipbot

![Latest release](https://img.shields.io/github/v/tag/deltalab-org/sipbot?label=release)
[![CI](https://github.com/deltalab-org/sipbot/actions/workflows/ci.yml/badge.svg)](https://github.com/deltalab-org/sipbot/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-16.9%25-red)
[![Go Report Card](https://goreportcard.com/badge/github.com/deltalab-org/sipbot)](https://goreportcard.com/report/github.com/deltalab-org/sipbot)

Bot to manage registrations in a [Flexisip](https://www.linphone.org/technical-corner/flexisip) server

## Install

Binary releases can be found at: https://github.com/deltalab-org/sipbot/releases

To install from source:

```sh
go install github.com/deltalab-org/sipbot@latest
```

### Installing deltachat-rpc-server

This program depends on a standalone Delta Chat RPC server `deltachat-rpc-server` program that must be
available in your `PATH`. For installation instructions check:
https://github.com/deltachat/deltachat-core-rust/tree/master/deltachat-rpc-server

## Running the bot

Configure the bot:

```sh
sipbot init bot@example.com PASSWORD
```

To start listening to incoming messages, you must set the environment variable `SIPBOT_DBDSN`
with the [DSN (Data Source Name)](https://github.com/go-sql-driver/mysql/#dsn-data-source-name)
to connect to the Flexisip database, and set `SIPBOT_DOMAIN` to the domain name of the SIP server:

```sh
export SIPBOT_DBDSN="flexisip:<PASSWORD>@tcp(localhost:3306)/flexisip_accounts"
export SIPBOT_DOMAIN="example.com"
sipbot serve
```

Run `sipbot --help` to see all available options.

## Contributing

Pull requests are welcome! check [CONTRIBUTING.md](CONTRIBUTING.md)
