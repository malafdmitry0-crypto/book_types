import test from 'node:test';
import assert from 'node:assert/strict';
import fs from 'node:fs';
import '../web/syntax.js';
const {highlight} = globalThis.BookSyntax;
const plain = html => html.replace(/<span class="syntax-[a-z]+">|<\/span>/g, '')
  .replaceAll('&quot;', '"').replaceAll('&gt;', '>').replaceAll('&lt;', '<').replaceAll('&amp;', '&');

test('preserves every code example in the textbook', () => {
  for (const chapter of JSON.parse(fs.readFileSync(new URL('../book/chapters.json', import.meta.url)))) {
    const source = fs.readFileSync(new URL('../book/chapters/' + chapter.filename, import.meta.url), 'utf8');
    for (const [,language,code] of source.matchAll(/```([^\n]*)\n([\s\S]*?)\n```/g)) {
      assert.equal(plain(highlight(code, language)), code, chapter.filename);
    }
  }
});
test('comments and strings stay single tokens; markup is escaped', () => {
  const source = 'func main() {\n // return 42\n s := "<script>&</script>"\n raw := `func return`\n}';
  const html = highlight(source);
  assert.match(html, /syntax-keyword">func/);
  assert.match(html, /syntax-comment">\/\/ return 42<\/span>/);
  assert.match(html, /syntax-string">`func return`<\/span>/);
  assert.ok(!html.includes('<script>'));
  assert.equal(plain(html), source);
});
test('Go numeric forms, Unicode and commands preserve their text', () => {
  for (const source of ['число := 0x1.fp+2 + 0b101 + 1_000 + .5 + 2i', "r := '\\n' /* func */", 'go test ./... -run "Example" # test']) {
    assert.equal(plain(highlight(source, source.startsWith('go test') ? 'sh' : 'go')), source);
  }
});
test('diagram text is escaped without Go highlighting', () => {
  assert.equal(highlight('A < B → C', 'text'), 'A &lt; B → C');
});
