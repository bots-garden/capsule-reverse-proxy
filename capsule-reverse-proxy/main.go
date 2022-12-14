package main

import (
    "flag"
    "fmt"
    "github.com/bots-garden/capsule-reverse-proxy/reverse-proxy"
    "github.com/bots-garden/capsule-reverse-proxy/commons"
    "os"
)

type CapsuleFlags struct {
    httpPort string
    config   string // config file for the reverse proxy
    crt      string // https (certificate)
    key      string // https (key)
    backend  string // backend for reverse proxy (and/or service discovery)

}

func main() {
    //argsWithProg := os.Args
    args := os.Args[1:]

    if len(args) == 0 {
        fmt.Println("😮 no args. Type capsule-reverse-proxy --help")
        os.Exit(0)
    }

    if args[0] == "version" {
        fmt.Println(commons.CapsuleReverseProxyVersion())
    } else {
        //flags
        httpPortPtr := flag.String("httpPort", "8080", "http port")

        configPtr := flag.String("config", "", "config file (🚧 not implemented)")
        backendPtr := flag.String("backend", "memory", "backend for reverse proxy, registration, discovery")

        crtPtr := flag.String("crt", "", "certificate")
        keyPtr := flag.String("key", "", "key")

        flag.Parse()

        flags := CapsuleFlags{
            *httpPortPtr,
            *configPtr,
            *crtPtr,
            *keyPtr,
            *backendPtr,
        }

        reverse_proxy.Serve(flags.httpPort, flags.config, flags.backend, flags.crt, flags.key)

    }

}
