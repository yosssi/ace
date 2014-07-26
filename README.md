# Ace - HTML template engine for Go

[![GoDoc](https://godoc.org/github.com/yosssi/ace?status.svg)](https://godoc.org/github.com/yosssi/ace)

## Overview

Ace is an HTML template engine for Go. This is inspired by [Slim](http://slim-lang.com/) and [Jade](http://jade-lang.com/). This is a refinement of [Gold](http://gold.yoss.si/).

## Example

```ace
= doctype html
html lang=en
  head
    title Hello Ace
    = css
      h1 { color: blue; }
  body
    h1 {{.Msg}}
    #container.wrapper
      p..
        Ace is an HTML template engine for Go.
        This engine simplifies HTML coding in Go web application development.
    = javascript
      var msg = 'Welcome to Ace';
      console.log(msg);
```

becomes

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <title>Hello Ace</title>
    <style type="text/css">
      h1 { color: blue; }
    </style>
  </head>
  <body>
    <h1>Hello Ace</h1>
    <div id="container" class="wrapper">
      <p>
        Ace is an HTML template engine for Go.<br>
        This engine simplifies HTML coding in Go web application development.
      </p>
    </div>
    <script type="text/javascript">
      var msg = 'Welcome to Ace';
      console.log(msg);
    </script>
  </body>
</html>
```

## Installation

```sh
$ go get github.com/yosssi/ace/...
```

## Implementation Example

```go
package main

import (
	"net/http"

	"github.com/yosssi/ace"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.ParseFiles("hello", "", nil) // Parse `hello.ace`.
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

## Documentation

You can get the documentation about Ace via the following channels:

* [Documents](https://github.com/yosssi/ace/tree/master/docs)
* [GoDoc](https://godoc.org/github.com/yosssi/ace)

## Official Information & Announcement

You can get the official information and announcement about Ace via the following channels:

* Website
* [Twitter](https://twitter.com/acehtml)
* [Google Group](https://groups.google.com/forum/#!forum/acehtml)

## Discussion & Contact

You can discuss Ace and contact the Ace development team via the following channels:

* [GitHub Issues](https://github.com/yosssi/ace/issues)
* [Google Group](https://groups.google.com/forum/#!forum/acehtml)
* [Twitter](https://twitter.com/acehtml)
* [Gitter (Chat)](https://gitter.im/yosssi/ace)
* [Keiji Yoshida (Author of Ace)'s Twitter](https://twitter.com/_yosssi)
* [Keiji Yoshida (Author of Ace)'s Email](mailto:yoshida.keiji.84@gmail.com)

## Tools

* HTML2Ace
* Ace2HTML
* vim-ace - Vim syntax highlighting for Ace templates
