package ace

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var innerTemplate string = `
= content inner
  section
    h1 hello
`

func TestFormatter(t *testing.T) {

	inner, _ := ioutil.TempFile("", "inner")
	inner.WriteString(innerTemplate)
	inner.Close()

	os.Rename(inner.Name(), inner.Name()+".ace")
	defer os.Remove(inner.Name() + ".ace")

	for i, this := range []struct {
		template string
		expect   string
	}{
		{
			template: `
a
  span foo
`,
			expect: `<a><span>foo</span></a>`,
		},
		{
			template: `
div
  span foo
`,
			expect: `<div>
 <span>foo</span>
</div>`,
		},
		{
			template: `
section
  div
    {{if true}}
      span foo
    {{end}}
  div text
`,
			expect: `<section>
 <div>
   <span>foo</span>
 </div>
 <div>
  text
 </div>
</section>`,
		},
		{
			template: `
body
  = yield inner
`,
			expect: `<body>
 
 <section>
  <h1>
   hello
  </h1>
 </section>
</body>`,
		},
	} {
		base, _ := ioutil.TempFile("", "base")

		base.WriteString(this.template)
		base.Close()
		os.Rename(base.Name(), base.Name()+".ace")
		defer os.Remove(base.Name() + ".ace")

		tpl, _ := Load(base.Name(), inner.Name(), &Options{
			Indent: " ",
		})
		buf := new(bytes.Buffer)
		tpl.Execute(buf, map[string]string{})

		compiled := buf.String()
		if compiled != this.expect {
			t.Errorf("[%d] Compiler didn't return an expected value, got %v but expected %v", i, compiled, this.expect)
		}
	}
}
