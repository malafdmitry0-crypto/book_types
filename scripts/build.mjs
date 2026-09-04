import fs from 'node:fs';
import path from 'node:path';
import {fileURLToPath} from 'node:url';
const root=path.resolve(path.dirname(fileURLToPath(import.meta.url)),'..');
const read=p=>fs.readFileSync(path.join(root,p),'utf8');
const chapters=JSON.parse(read('book/chapters.json'));
const escape=s=>s.replaceAll('&','&amp;').replaceAll('<','&lt;').replaceAll('>','&gt;').replaceAll('"','&quot;');
function href(url){
 const chapter=chapters.find(c=>c.filename===url);
 if(chapter)return '#'+chapter.id;
 if(url==='../README.md')return '#contents';
 if(url.startsWith('https://')||url.startsWith('#'))return escape(url);
 if(url.startsWith('../../examples/')) {
  const target=path.resolve(root,'book/chapters',url);
  if(target.startsWith(path.join(root,'examples')+path.sep)&&fs.existsSync(target))return escape(path.relative(root,target));
 }
 throw new Error('Unsupported link: '+url);
}
function inline(text){
 const tokens=[];
 const token=html=>`\u0000${tokens.push(html)-1}\u0000`;
 let s=text.replace(/`([^`]+)`/g,(_,code)=>token('<code>'+escape(code)+'</code>'));
 s=s.replace(/\[([^\]]+)\]\(([^)]+)\)/g,(_,label,url)=>token(`<a href="${href(url)}">${escape(label)}</a>`));
 s=escape(s).replace(/\*\*([^*]+)\*\*/g,'<strong>$1</strong>');
 // Nested placeholders can occur when a link label contains code.
 for(let pass=0;pass<3;pass++)s=s.replace(/\u0000(\d+)\u0000/g,(_,n)=>tokens[n]);
 return s;
}
function cells(row){
 // Pipes inside code spans are content, not column boundaries.
 const parts=[];let text='',code=false;
 for(const char of row.trim().replace(/^\|/,'').replace(/\|$/,'')){
  if(char==='`')code=!code;
  if(char==='|'&&!code){parts.push(text.trim());text='';}else text+=char;
 }
 parts.push(text.trim()); return parts;
}
function markdown(source,id){
 const lines=source.split('\n'),out=[];let section=0;
 for(let i=0;i<lines.length;){
  const line=lines[i];
  if(!line.trim()){i++;continue;}
  if(line.startsWith('```')){
   const language=line.slice(3).trim() || 'text'; const code=[];i++;while(i<lines.length&&!lines[i].startsWith('```'))code.push(lines[i++]);
   if(i===lines.length)throw new Error('Unclosed code block in '+id);i++;
   out.push(`<div class="code-block"><div class="code-toolbar"><span>${escape(language === "go" ? "Go" : "Схема")}</span><button class="copy" type="button" aria-label="Копировать пример кода">Копировать</button></div><pre><code>${escape(code.join('\n'))}</code></pre></div>`);continue;
  }
  if(/^#{1,3} /.test(line)){
   const [,hashes,title]=line.match(/^(#{1,3}) (.*)$/);
   if(hashes.length>1){section++;out.push(`<h2 id="${id}-s${section}">${inline(title)}</h2>`);}i++;continue;
  }
  if(line.startsWith('|')){
   const rows=[];while(i<lines.length&&lines[i].startsWith('|'))rows.push(cells(lines[i++]));
   out.push('<div class="table-wrap" role="region" aria-label="Таблица" tabindex="0"><table><thead><tr>'+rows[0].map(x=>'<th scope="col">'+inline(x)+'</th>').join('')+'</tr></thead><tbody>'+rows.slice(2).map(row=>'<tr>'+row.map(x=>'<td>'+inline(x)+'</td>').join('')+'</tr>').join('')+'</tbody></table></div>');continue;
  }
  if(line.startsWith('> ')){out.push('<aside class="objective">'+inline(line.slice(2))+'</aside>');i++;continue;}
  if(/^\d+\. /.test(line)){
   const items=[];while(i<lines.length&&/^\d+\. /.test(lines[i]))items.push('<li>'+inline(lines[i++].replace(/^\d+\. /,''))+'</li>');
   out.push('<ol>'+items.join('')+'</ol>');continue;
  }
  if(line==='---')break; // Markdown chapter navigation is replaced with the HTML pager.
  const paragraph=[];while(i<lines.length&&lines[i].trim()&&!/^(#{1,3} |```|\||> |\d+\. |---$)/.test(lines[i]))paragraph.push(lines[i++]);
  out.push('<p>'+inline(paragraph.join(' '))+'</p>');
 }
 return out.join('\n');
}
const articles=chapters.map((c,i)=>{
 const source=read('book/chapters/'+c.filename);
 c.text=source.replace(/\n---\n[\s\S]*$/,'');
 const body=markdown(source,c.id);
 const minutes=Math.max(1,Math.ceil(source.split(/\s+/).length/150));
 const prev=i?chapters[i-1]:null,next=chapters[i+1];
 return `<article class="chapter" id="${c.id}" aria-labelledby="${c.id}-title"><header class="chapter-heading"><div class="eyebrow">${escape(c.part)} / ${String(c.number).padStart(2,'0')} <span>· ${minutes} мин чтения</span></div><h1 id="${c.id}-title" tabindex="-1">${inline(c.title)}</h1><a class="source-link" href="book/chapters/${c.filename}">Исходник Markdown ↗</a></header>${body}<nav class="pager" aria-label="Переход между главами">${prev?`<a href="#${prev.id}"><small>← Предыдущая</small>${inline(prev.title)}</a>`:'<a href="#contents"><small>← К началу</small>Оглавление</a>'}${next?`<a href="#${next.id}"><small>Следующая →</small>${inline(next.title)}</a>`:'<a href="#contents"><small>Вы прошли учебник</small>Вернуться к оглавлению →</a>'}</nav></article>`;
}).join('\n');
const links=['Основной курс','Справочник'].map(part=>`<details class="nav-group" ${part==='Основной курс'?'open':''}><summary>${part} · ${chapters.filter(c=>c.part===part).length}</summary>${chapters.filter(c=>c.part===part).map(c=>`<a href="#${c.id}" data-chapter="${c.id}"><span>${String(c.number).padStart(2,'0')}</span><span>${inline(c.title)}</span></a>`).join('')}</details>`).join('');
const data=JSON.stringify(chapters.map(c=>({id:c.id,title:c.title,text:c.text}))).replaceAll('<','\\u003c');
const html=`<!doctype html>
<html lang="ru"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><meta name="description" content="Обобщённое программирование в Go на примере Find, All и Map: интерфейсы, reflection, генерация, дженерики и итераторы."><title>Обобщённое программирование в Go — учебник</title><style>${read('web/styles.css')}</style></head>
<body><a class="skip" href="#main">Перейти к тексту</a><header class="topbar"><a class="brand" href="#contents"><span class="brand-symbol">[E]</span><span>Go / Алгоритмы</span></a><div class="top-actions"><span class="edition">Одна библиотека · Восемь подходов</span><button id="print" type="button">Печать / PDF</button><button id="menu" type="button" aria-expanded="false" aria-controls="sidebar">Главы</button></div></header>
<div class="layout"><aside id="sidebar"><div class="sidebar-top"><span class="eyebrow">Содержание</span><span>${chapters.length} разделов</span></div><label class="search-label" for="search">Поиск по учебнику</label><div class="search-box"><input id="search" type="search" placeholder="Например, интерфейсы" autocomplete="off"><button id="clear-search" type="button" aria-label="Очистить поиск" hidden>×</button></div><p id="search-status" role="status" aria-live="polite" hidden></p><nav id="chapter-nav" aria-label="Главы учебника">${links}</nav><div class="sidebar-note">Один алгоритм.<br>Разные границы абстракции.<a href="book/README.md">Читать в Markdown ↗</a></div></aside>
<main id="main" tabindex="-1"><section id="contents">${read('web/cover.html')}</section>${articles}<footer>Go / Алгоритмы <span>Обобщённое программирование · Markdown + HTML</span></footer></main></div><div id="notice" role="status" aria-live="polite"></div><script id="book-data" type="application/json">${data}</script><script>${read('web/app.js')}</script></body></html>`;
fs.writeFileSync(path.join(root,'index.html'),html);
console.log(`Built index.html: ${chapters.length} chapters, ${Buffer.byteLength(html)} bytes`);
