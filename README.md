# crt-cli

Golang package to retrieve subdomains of a domain using [Crt.sh](https://crt.sh)

It will then resolve the domain name and try to retrieve the content on 443/80 if ports are open

## Installation
`go get -u github.com/cptainobvious/crt-cli`

## Usage
`crt-cli find example.com`

### Options
#### Blacklist
`--blacklist *.blacklist.example.com --blacklist production.*.example.com`