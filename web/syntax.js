// Общая подсветка для сборки HTML и интерактивного примера. Без сети и зависимостей.
(() => {
  const escape = text => text.replaceAll('&', '&amp;').replaceAll('<', '&lt;')
    .replaceAll('>', '&gt;').replaceAll('"', '&quot;');
  const keywords = new Set('break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var'.split(' '));
  const types = new Set('any bool byte comparable complex64 complex128 error float32 float64 int int8 int16 int32 int64 rune string uint uint8 uint16 uint32 uint64 uintptr'.split(' '));
  const builtins = new Set('append cap clear close complex copy delete imag len make max min new panic print println real recover'.split(' '));
  const literals = new Set(['true', 'false', 'nil', 'iota']);
  // Строки и комментарии распознаются раньше идентификаторов, содержимое не перекрашивается.
  const goTokens = /\/\/[^\r\n]*|\/\*[\s\S]*?(?:\*\/|$)|"(?:\\[\s\S]|[^"\\])*"?|'(?:\\[\s\S]|[^'\\])*'?|`[^`]*`?|\b0[xX][\da-fA-F_]+(?:\.[\da-fA-F_]*)?(?:[pP][+-]?[\d_]+)?i?|\b0[bB][01_]+|\b0[oO][0-7_]+|(?:\b\d[\d_]*(?:\.[\d_]*)?|\.\d[\d_]*)(?:[eE][+-]?[\d_]+)?i?|[\p{L}_][\p{L}\p{N}_]*|(?:<<|>>|&\^|:=|==|!=|<=|>=|&&|\|\||<-|\+\+|--|[-+*\/%&|^=<>!:])=?|\s+|[^\s]/gu;
  const shellTokens = /#[^\r\n]*|"(?:\\[\s\S]|[^"\\])*"?|'[^']*'?|\$\{[^}]*\}|\$[\w]+|--?[\w-]+|\b(?:go|node|git)\b|\b\d+(?:ms|s)?\b|[|><&]+|\s+|[^\s]/g;
  function highlight(source, language = 'go') {
    const shell = ['sh', 'bash', 'shell'].includes(language);
    if (language !== 'go' && !shell) return escape(source);
    let html = '';
    let previous = '';
    const tokenPattern = shell ? shellTokens : goTokens;
    for (const match of source.matchAll(tokenPattern)) {
      const token = match[0];
      let kind = '';
      if (shell) {
        if (token.startsWith('#')) kind = 'comment';
        else if (/^["']/.test(token)) kind = 'string';
        else if (token.startsWith('$')) kind = 'type';
        else if (/^--?\w/.test(token)) kind = 'keyword';
        else if (/^(go|node|git)$/.test(token)) kind = 'function';
        else if (/^\d/.test(token)) kind = 'number';
        else if (/^[|><&]/.test(token)) kind = 'operator';
      } else {
        if (token.startsWith('//') || token.startsWith('/*')) kind = 'comment';
        else if (/^["'`]/.test(token)) kind = 'string';
        else if (/^(\d|\.\d)/.test(token)) kind = 'number';
        else if (keywords.has(token)) kind = 'keyword';
        else if (types.has(token) || previous === 'type') kind = 'type';
        else if (literals.has(token)) kind = 'literal';
        else if (builtins.has(token)) kind = 'function';
        else if (/^[\p{L}_]/u.test(token)) {
          const rest = source.slice(match.index + token.length);
          if (/^\s*(?:\[[^\]\n]*\]\s*)?\(/.test(rest)) kind = 'function';
          else if (/^\p{Lu}/u.test(token)) kind = 'type';
        } else if (/^[-+*\/%&|^=<>!:]/.test(token)) kind = 'operator';
      }
      html += kind ? `<span class="syntax-${kind}">${escape(token)}</span>` : escape(token);
      if (token.trim() && kind !== 'comment') previous = token;
    }
    return html;
  }
  globalThis.BookSyntax = Object.freeze({highlight});
})();
