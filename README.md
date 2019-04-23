
# Why?
This repo is just a copy of github.com/emitter-io/emitter but allows easy integration into your own single binary.

# What?
The following commands are executed on the original source:
<pre>
$ mv internal/* -r .
$ find ./ -type f -print0 | xargs -0 sed -i 's/github.com\/emitter-io\/emitter\/internal\//github.com\/emitter-io\/emitter\//g'
$ find ./ -type f -print0 | xargs -0 sed -i 's/github.com\/emitter-io\/emitter\//mheers\/emitter\//g'
</pre>

# How to use?
<pre>
$ go get github.com/mheers/emitter
</pre>

then you can start the broker from your own Go-file:
<pre>
package main

import (
	"context"

	"github.com/mheers/emitter/broker"
	emitterConfig "github.com/mheers/emitter/config"
)

func startEmitterSvc() {
	cfg := &emitterConfig.Config{
		ListenAddr: ":8081",
		License:    "J9VGSRu7nCDkU8VPhRPYG2cJm4n6V-9bAAAAAAAAAAI",
	}

	// Start new service
	svc, err := broker.NewService(context.Background(), cfg)
	if err != nil {
		panic(err.Error())
	}

	// Listen and serve
	svc.Listen()
}

func main() {
    go startEmitterSvc()
}
</pre>