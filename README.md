# gRPC Go Demo

In this application I test my (and maybe your) skills, to learn gRPC the hard way.

By "hard way", I mean the RTFM way, or in simple terms, by just reading text based resources and avoid using codding assistants.

## But why?

- I have never used gRPC before this project, so to learn it I guess
- Because we all got lazy and let copilot, youtube and random ass sites divide or distract our attention
- To prove myself I'm still able to read docs

## Used resources

### Just

- https://github.com/casey/just
- https://just.systems/man/en/
- Without integrations since the docs states there is and [up-to-date integration, yet it was archived since.](https://just.systems/man/en/chapter_13.html)
- The arch and os magic comes from previous experiences, there are no clear docs for those.
  - ARM macs usually reports `arm64` in `uname -a`, while linux distros tends to use `aarch64` thats why I used just's `arch()` function
  - [just's os function](https://just.systems/man/en/chapter_31.html#system-information) reports macs as `macos` so I need to change it to osx

### Protobuf and ProtoC

- https://grpc.io/docs/protoc-installation/#install-pre-compiled-binaries-any-os
- https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code
- https://pkg.go.dev/google.golang.org/grpc/test/bufconn


### gRPC

- https://github.com/grpc/grpc-go/blob/master/examples/helloworld/greeter_server/main.go
- https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc#section-readme
- https://grpc.io/docs/languages/go/quickstart/

### Go

- The folder layout, interface implementations, table tests and the error handling,
  basically everything comes from my previous experiences, saddly I cannot unlearn those, netiher link the docs.
- https://stackoverflow.com/a/52080545
- https://pkg.go.dev/testing#hdr-Main
