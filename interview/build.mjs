import fs from 'node:fs';
import {execFileSync} from 'node:child_process';
import {createHash} from 'node:crypto';
import {tasks} from './all-tasks.mjs';
import {answerDiagrams} from './answer-diagrams.mjs';
import '../web/syntax.js';
const root=new URL('./',import.meta.url);
const read=p=>fs.readFileSync(new URL(p,root),'utf8');
const write=(p,s)=>fs.writeFileSync(new URL(p,root),s);
const escape=s=>String(s).replaceAll('&','&amp;').replaceAll('<','&lt;').replaceAll('>','&gt;').replaceAll('"','&quot;');
const topics=[...new Set(tasks.map(t=>t.topic))];
if(tasks.length!==800||topics.length!==40||topics.some(topic=>tasks.filter(t=>t.topic===topic).length!==20))throw Error('Expected 40 topics × 20 tasks');
if(new Set(tasks.map(t=>t.source)).size!==800||new Set(tasks.map(t=>t.title)).size!==800)throw Error('Duplicate task');
const proof=JSON.parse(read('verification.json'));
const digest=createHash('sha256').update(JSON.stringify(tasks)).digest('hex');
if(proof.checked!==800||proof.passed!==800||proof.sourceDigest!==digest)throw Error('Run node interview/verify.mjs for this catalog first');
fs.mkdirSync(new URL('testdata/',root),{recursive:true});
const formatted=tasks.map(t=>{
 // Разносим операторы, не меняя литералы и не теряя явные точки с запятой
 // (в частности, в заголовках for). gofmt затем убирает лишние разделители.
 let source='',quote='',escaped=false;
 for(const c of t.source){
  source+=c;
  if(quote){if(escaped){escaped=false;continue;}if(c==='\\'&&quote!=='`'){escaped=true;continue;}if(c===quote)quote='';}
  else if(c==='"'||c==="'"||c==='`')quote=c;
  else if(c===';')source+='\n';
 }
 source=source.replace('func main(){','func main(){\n').replace(/}\n$/,'\n}\n');
 try{source=execFileSync('gofmt',[],{input:source,encoding:'utf8',stdio:['pipe','pipe','pipe']});}
 catch(e){
  if(t.kind!=='compile')throw e;
  // Старый gofmt не разбирает намеренно недопустимый generic-метод.
  // Форматируем его без списка параметров и возвращаем ровно тот же список.
  if(source.includes('Map[R any]'))source=execFileSync('gofmt',[],{input:source.replace('Map[R any]','Map'),encoding:'utf8'}).replace(' Map(',' Map[R any](');
 }
 const file=`testdata/${String(t.id).padStart(3,'0')}.go`;
 write(file,source);
 return {...t,source,file,topicIndex:topics.indexOf(t.topic)};
});
fs.mkdirSync(new URL('assets/answers/',root),{recursive:true});
for(const [id,diagram] of answerDiagrams){
 if(id<1||id>100||!tasks.some(t=>t.id===id))throw Error('Diagram outside reviewed tasks: '+id);
 write(diagram.file,diagram.svg);
}
const diagramMarkdown=t=>{const d=answerDiagrams.get(t.id);return d?'\n\n!['+d.caption+']('+d.file+')':'';};
const diagramHTML=t=>{const d=answerDiagrams.get(t.id);return d?`<figure class="answer-diagram"><img src="data:image/svg+xml;base64,${Buffer.from(d.svg).toString('base64')}" alt="${escape(d.caption)}" width="600" height="${d.height}"><figcaption>${escape(d.caption)}</figcaption></figure>`:'';};
const prompt='Скомпилируется ли программа? Если да, что она выведет и возникнет ли panic? Объясните правило типов и момент его проверки.';
let markdown='# 800 задач на интервью: код — что будет?\n\n'+read('intro.md')+'\n\n';
for(const topic of topics){markdown+=`## ${topic}\n\n`;for(const t of formatted.filter(t=>t.topic===topic))markdown+=`### ${String(t.id).padStart(3,'0')}. ${t.title}\n\n${prompt}\n\n\x60\x60\x60go\n${t.source.trimEnd()}\n\x60\x60\x60\n\n${t.choices?t.choices.map((c,i)=>`${'ABCD'[i]}.\n\n\x60\x60\x60text\n${c}\n\x60\x60\x60`).join('\n\n')+'\n\n':''}<details>\n<summary>Ответ и объяснение</summary>\n\n${t.choices?`**Правильный вариант: ${'ABCD'[t.correctIndex]}.**\n\n`:''}${t.kind==='compile'?'**Не компилируется.** Характерный фрагмент диагностики:':'**Вывод программы:**'}\n\n\x60\x60\x60text\n${t.answer}\n\x60\x60\x60\n\n${t.why}${diagramMarkdown(t)}\n\n</details>\n\n`;}
write('TASKS.md',markdown);
const letter=i=>'ABCD'[i];
const choicesHTML=t=>t.choices?`<fieldset class="choices"><legend>Выберите один ответ</legend>${t.choices.map((c,i)=>`<label><input type="radio" name="choice-${t.id}" value="${i}"><span class="choice-letter">${letter(i)}.</span><code>${escape(c)}</code></label>`).join('')}</fieldset>`:'';
const correctHTML=t=>t.choices?`<p class="correct-choice"><strong>Правильный вариант: ${letter(t.correctIndex)}.</strong></p>`:'';
const articles=formatted.map(t=>`<article id="task-${t.id}" class="task" data-topic="${t.topicIndex}" aria-labelledby="title-${t.id}"><header><a class="task-number" href="#task-${t.id}" aria-label="Ссылка на задачу ${t.id}">${String(t.id).padStart(3,'0')}</a><div><p class="topic-name">${escape(t.topic)}</p><h2 id="title-${t.id}" tabindex="-1">${escape(t.title)}</h2></div></header><p class="question">${prompt}</p><div class="code-block"><div class="code-toolbar"><span>Go 1.23</span><a href="${t.file}">Полный исходник ↗</a></div><pre><code class="language-go">${globalThis.BookSyntax.highlight(t.source.trimEnd(),'go')}</code></pre></div>${choicesHTML(t)}<details class="answer"><summary>Ответ и объяснение</summary><div class="answer-body">${correctHTML(t)}<p><strong>${t.kind==='compile'?'Не компилируется.':'Вывод программы:'}</strong>${t.kind==='compile'?' Характерный фрагмент диагностики:':''}</p><pre><code>${escape(t.answer)}</code></pre><p>${escape(t.why)}</p>${diagramHTML(t)}</div></details></article>`).join('\n');
const options=topics.map((t,i)=>`<option value="${i}">${String(i+1).padStart(2,'0')} · ${escape(t)}</option>`).join('');
const data=JSON.stringify(formatted.map(t=>({id:t.id,topic:t.topicIndex,search:`${String(t.id).padStart(3,'0')} ${t.title} ${t.source}`.toLowerCase()}))).replaceAll('<','\\u003c');
const html=`<!doctype html><html lang="ru"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>800 задач по типизации Go — код и что будет</title><style>${read('../web/styles.css')}\n${read('style.css')}</style></head><body><a class="skip" href="#tasks">Перейти к задачам</a><header class="topbar"><a class="brand" href="../index.html"><span class="brand-symbol">Go</span><span>Go / Интервью</span></a><a href="../index.html">К учебнику →</a></header><main class="interview"><section class="intro"><p class="eyebrow">Типизация Go · 800 задач · 40 тем</p><h1>Код есть.<br>Что будет дальше?</h1><p>Предскажите вывод, найдите отказ компилятора или объясните панику. Ответ открывайте после собственной версии. В новом блоке №601–800 — четыре варианта, один правильный и разбор. Это авторские задачи, а не официальный сертификационный экзамен.</p><p class="version-note">Все задачи проверены на <strong>Go 1.23.0, darwin/arm64</strong>. Это базовая версия сборника: на новых версиях отдельные правила меняются. Полные программы независимы; паники в задачах с recover перехватываются самим кодом.</p><p><a class="start" href="?set=exam#tasks">Новые 200: экзаменационные задачи →</a></p><p><a href="TASKS.md">Все задачи в Markdown ↗</a> · <a href="README.md">Проверка и правила решения ↗</a></p></section><section class="controls" aria-label="Выбор задач"><div><label for="collection">Набор задач</label><select id="collection"><option value="all">Все 800 задач</option><option value="original">Первые 400 · типизация</option><option value="extra">401–600 · адаптеры и композиция</option><option value="exam">601–800 · экзамен с вариантами</option></select></div><div><label for="topic">Тема</label><select id="topic"><option value="all">Все 40 тем</option>${options}</select></div><div class="search-control"><label for="query">Поиск по номеру, названию или коду</label><input id="query" type="search" placeholder="Например, 274 или comparable" autocomplete="off"></div><div class="control-actions"><button id="random" type="button">Случайная задача</button><button id="reset" type="button">Сбросить</button></div></section><p id="status" role="status" aria-live="polite">Показано задач: 800</p><p id="empty" hidden>Совпадений нет. Измените запрос или сбросьте фильтры.</p><section id="tasks" aria-label="Задачи" tabindex="-1">${articles}</section><footer><a href="../index.html">Вернуться к учебнику</a><span>800 самостоятельных программ · Ответы с объяснениями</span></footer></main><script id="task-data" type="application/json">${data}</script><script>${read('app.js')}</script></body></html>`;
write('index.html',html);
console.log(`Built interview/index.html and TASKS.md: ${tasks.length} tasks, ${topics.length} topics`);
