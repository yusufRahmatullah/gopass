# GoPass - Pasword (and pin Generator)

Password/pin generator using *master key* and *purpose* as the seed for the PRNG (*Pseudo Random Number Generator*) to generate 64-char alphanumeric and some symbols password (minimal secure length of SHA-256 key) and to generate 6-num pin. The generator can be used like a password manager that can be accessed anywhere (currently the web-version is hosted on [here](https://passgen.netlify.com)) and not persist on any storage (cloud or locally).

## Getting Started

This repo is the implementation of [passgen](https://github.com/yusufRahmatullah/passgen) in Golang. The implementation of PRNG is following the javascript version from [David Bau's seedrandom](https://github.com/davidbau/seedrandom) to maintain same result in Golang.

## Acknowledgment

This app doesn't have any warranty. Use this app on your risk.
