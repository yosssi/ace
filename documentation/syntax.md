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

A word of a line head is interpreted as an HTML tag. The rest words of the same line is interpreted as attributes or a text. An attribute value which contains spaces must be surrounded by double quotes.

```ace
div id=container style="font-size: 12px; color: blue;"
  p class=row This is interpreted as a text.
  a href=https://github.com/ Go to GitHub
```

becomes

```html
<div id="container" style="font-size: 12px; color: blue;">
  <p class="row">This is interpreted as a text.</p>
  <a href="https://github.com/">Go to GitHub</a>
</div>
```

## Plain Texts

## Helper Methods

## Comments

## Actions
