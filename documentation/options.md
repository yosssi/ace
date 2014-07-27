# Options

Here is a sample Go code which calls an Ace template engine.

```go
package main

import (
    "net/http"

    "github.com/yosssi/ace"
)

func handler(w http.ResponseWriter, r *http.Request) {
    tpl, err := ace.ParseFiles("base", "inner", nil)
    if err != nil {
        panic(err)
    }
    if err := tpl.Execute(w, map[string]string{"Msg": "Hello Ace"}); err != nil {
        panic(err)
    }
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

You can pass parsing options to the Ace template engine via `ace.ParseFiles`'s third argument.

```go
tpl, err := ace.ParseFiles("base", "inner", &ace.Options{Cache: true})
```

Please check [GoDoc](https://godoc.org/github.com/yosssi/ace#Options) for more detail about options.
