# Ace - HTML template engine for Go

[![wercker status](https://app.wercker.com/status/8d3c657bcae7f31d10c8f88bbfa966d8/m "wercker status")](https://app.wercker.com/project/bykey/8d3c657bcae7f31d10c8f88bbfa966d8)
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
      console.log('Welcome to Ace');
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
      console.log('Welcome to Ace');
    </script>
  </body>
</html>
```

## Features

* Making use of the Go standard template library : **Ace fully utilizes the strength of the [html/template](http://golang.org/pkg/html/template/) package.** You can embed the [actions](http://golang.org/pkg/text/template/#hdr-Actions) of the template package in the Ace templates. Ace uses the [nested template definitions](http://golang.org/pkg/text/template/#hdr-Nested_template_definitions) of the template package in the Ace template include function and you can have the Ace templates pass the [pipelines](http://golang.org/pkg/text/template/#hdr-Pipelines) / parameters to the other templates which they include.
* Simple Syntax : Ace has a simple syntax. This feature makes the HTML templates simple and light.


## Documentation

You can get the documentation about Ace via the following channels:

* [Documentation](https://github.com/yosssi/ace/tree/master/documentation)
* [GoDoc](https://godoc.org/github.com/yosssi/ace)

## Official Information & Announcement

You can get the official information and announcement about Ace via the following channels:

* Website (Under construction)
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

* HTML2Ace (Under construction)
* [vim-ace](https://github.com/yosssi/vim-ace) - Vim syntax highlighting for Ace templates
