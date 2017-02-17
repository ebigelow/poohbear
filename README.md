# poohbear

<a href="http://imgur.com/6KSV3YT"><img src="http://i.imgur.com/6KSV3YT.png" width="200" /></a>

<img src="https://img.shields.io/travis/clownpriest/poohbear.svg">
<img src="https://img.shields.io/aur/license/yaourt.svg">


poohbear is a database specifically built to collect and store cryptocoin market data.
It does both the collecting *and* the storing out of the box. Events (trades, orderbook
updates, etc...) are streamed into poohbear from different exchanges (Bitfinex, Poloniex, BTC-e,
etc..), and access to its warehouse of market history is exposed through a gRPC interface.

This repo also contains a Python interface to the poohbear daemon, in the lib/python folder.
It makes gRPC requests to the daemon and recieves protobuf messages in return.


**NOTE: this is still a very early work in progress. That being said, it already basically works, for a very limited
subset of what "working" means.**


## Building/Installing

For the main poohbear system:
```
$: git clone https://github.com/clownpriest/poohbear
$: cd poohbear
$: make install
$: poohbear
```

for the python gRPC interface:
```
$: cd lib/python
$: python setup.py install
```




### Currently implemented exchanges/pairs
Right now poohbear collects market data from
these exchanges, and for these trading pairs:

- Bitfinex (almost ready)
 - BTC/USD
- Poloniex
 - LTC/BTC
 - ETH/BTC
- **more to come!**
