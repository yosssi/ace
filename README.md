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

### Making Use of the Go Standard Template Package

**Ace fully utilizes the strength of the [html/template](http://golang.org/pkg/html/template/) package.** You can embed [actions](http://golang.org/pkg/text/template/#hdr-Actions) of the template package in Ace templates. Ace also uses [nested template definitions](http://golang.org/pkg/text/template/#hdr-Nested_template_definitions) of the template package and Ace templates can pass [pipelines](http://golang.org/pkg/text/template/#hdr-Pipelines) (parameters) to other templates which they include.

### Simple Syntax

Ace has a simple syntax and **this makes template files simple and light**.

### Caching Function

Ace has a caching function which caches the result data of the templates parsing process. **You can omit the templates parsing process and save template parsing time** by using this function.

### Binary Template Load Function

Ace has a binary template load function which loads Ace templates from binary data in memory instead of template files on disk. **You can compile your web application into one binary file** by using this function. [go-bindata](https://github.com/jteeuwen/go-bindata) is the best for generating binary data from template files.

## Getting Started

Please check the following documentation.

* [Getting Started](https://github.com/yosssi/ace/blob/master/documentation/getting-started.md) - shows the getting started guide.
* [Examples](https://github.com/yosssi/ace/tree/master/examples) - shows the examples of the web applications which use the Ace template engine.

## Documentation

You can get the documentation about Ace via the following channels:

* [Documentation](https://github.com/yosssi/ace/tree/master/documentation) - includes the getting started guide and the syntax documentation.
* [GoDoc](https://godoc.org/github.com/yosssi/ace) - includes the API docuemntation.

## Official Information & Announcement

You can get the official information and announcement about Ace via the following channels:

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

## Contributions

**Any contributions are welcome.** Please feel free to [create an issue](https://github.com/yosssi/ace/issues/new) or [send a pull request](https://github.com/yosssi/ace/compare/).

## Renderers for web frameworks

* [Martini Acerender](https://github.com/yosssi/martini-acerender) - For [Martini](http://martini.codegangsta.io/)

## Tools

* [vim-ace](https://github.com/yosssi/vim-ace) - Vim syntax highlighting for Ace templates
