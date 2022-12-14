package main

// TinyGo wasm module
import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.SetHandleHttp(Handle)
}

func Handle(req hf.Request) (resp hf.Response, errResp error) {
	html := `
    <html>
        <head>
            <title>Wasm is fantastic π</title>
        </head>

        <body>
            <h1>π Hola Mundo π</h1>
            <h2>Served with ππ with Capsule π</h2>
            <h3>π π π [BISπ]</h3>
        </body>

    </html>
    `

	headers := map[string]string{
		"Content-Type": "text/html; charset=utf-8",
	}

	return hf.Response{Body: html, Headers: headers}, nil
}
