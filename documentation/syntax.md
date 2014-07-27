# Syntax

## Indent

A unit of an indent must be 2 spaces.

```ace
html
  body
    div
    p
```

becomes

```html
<html>
  <body>
    <div></div>
    <p></p>
  </body>
</html>
```

## HTML Tags

A head word of a line is interpreted as an HTML tag. The rest words of the same line are interpreted as attributes or a text. An attribute value which contains spaces must be surrounded by double quotes. An attribute without value (like "check" and "required") can be defined by specifying no value and ending with an equal (=).

```ace
div id=container style="font-size: 12px; color: blue;"
  p class=row This is interpreted as a text.
  a href=https://github.com/ Go to GitHub
  input type=checkbox checked=
```

becomes

```html
<div id="container" style="font-size: 12px; color: blue;">
  <p class="row">This is interpreted as a text.</p>
  <a href="https://github.com/">Go to GitHub</a>
  <input type="checkbox" checked>
</div>
```

ID and classes can be defined with a head word of a line.

```ace
p#foo.bar
#container
.wrapper
```

becomes

```html
<p id="foo" class="bar"></p>
<div id="container"></div>
<div class="wrapper"></div>
```

Block texts can be defined as a child element of an HTML tag by appending a dot (.) or double dot (..) to the head word of a line. BR tags are inserted to each line except for the last line by appending a double dot (..) to the head word of a line.

```ace
script.
  var msg = 'Hello Ace';
  alert(msg);
p..
  This is a block text.
  BR tags are inserted
  automatically.
```

becomes

```html
<script>
  var msg = 'Hello Ace';
  alert(msg);
</script>
<p>
  This is a block text.<br>
  BR tags are inserted<br>
  automatically.
</p>
```

## Plain Texts

A line which starts with a pipe (|) or double pipe (||) is interpreted as plain texts. BR tags are inserted to each line except for the last line by having a line start with a double pipe (||).

```ace
div
  | This is a single line.
div
  |
    This is a
    block line.
div
  ||
    This is a
    block line
    with BR tags.
```

becomes

```html
<div>This is a single line.</div>
<div>
  This is a
  block line.
</div>
<div>
  This is a<br>
  block line<br>
  with BR tags.
</div>
```

## Helper Methods

## Comments

## Actions
