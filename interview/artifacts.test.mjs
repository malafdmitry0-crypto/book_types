import test from 'node:test';
import assert from 'node:assert/strict';
import fs from 'node:fs';
import {createHash} from 'node:crypto';
import {tasks} from './all-tasks.mjs';
import {answerDiagrams} from './answer-diagrams.mjs';

const read=p=>fs.readFileSync(new URL(p,import.meta.url),'utf8');
const html=read('index.html');
const decode=s=>s.replace(/<[^>]*>/g,'').replaceAll('&quot;','"').replaceAll('&gt;','>').replaceAll('&lt;','<').replaceAll('&amp;','&');

test('selected answer diagrams are offline, inside closed answers, and match Markdown assets',()=>{
 const md=read('TASKS.md');
 assert.equal(answerDiagrams.size,27);
 for(const t of tasks){
  const article=html.slice(html.indexOf(`<article id="task-${t.id}"`)).split('</article>')[0];
  const diagram=answerDiagrams.get(t.id);
  assert.equal(article.includes('<figure class="answer-diagram">'),!!diagram,`task ${t.id}`);
  if(!diagram)continue;
  assert.ok(t.id<=100);
  const answer=article.match(/<details class="answer">([\s\S]*?)<\/details>/)[1];
  const image=answer.match(/<img src="data:image\/svg\+xml;base64,([^"]+)"/);
  assert.ok(image,`diagram must be inside answer ${t.id}`);
  assert.equal(Buffer.from(image[1],'base64').toString(),read(diagram.file));
  assert.equal(read(diagram.file),diagram.svg);
  assert.ok(md.includes(`![${diagram.caption}](${diagram.file})`));
  assert.ok(!/<script\b|<foreignObject\b|(?:href|src)="https?:/.test(diagram.svg));
 }
});

test('exactly 800 distinct tasks, 20 per topic, with full verification',()=>{
 assert.equal(tasks.length,800);
 assert.equal(new Set(tasks.map(t=>t.source)).size,800);
 assert.equal(new Set(tasks.map(t=>t.title)).size,800);
 const topics=[...new Set(tasks.map(t=>t.topic))];
 assert.equal(topics.length,40);
 for(const topic of topics)assert.equal(tasks.filter(t=>t.topic===topic).length,20);
 const proof=JSON.parse(read('verification.json'));
 assert.equal(proof.checked,800);assert.equal(proof.passed,800);
 assert.equal(proof.sourceDigest,createHash('sha256').update(JSON.stringify(tasks)).digest('hex'));
});

test('HTML preserves all displayed Go files and keeps answers closed',()=>{
 const blocks=[...html.matchAll(/<code class="language-go">([\s\S]*?)<\/code>/g)];
 assert.equal(blocks.length,800);
 for(let i=0;i<blocks.length;i++)assert.equal(decode(blocks[i][1]),read(`testdata/${String(i+1).padStart(3,'0')}.go`).trimEnd(),`task ${i+1}`);
 assert.equal([...html.matchAll(/<details class="answer">/g)].length,800);
 assert.equal([...html.matchAll(/<details[^>]*\bopen\b/g)].length,0);
 const ids=[...html.matchAll(/\bid="([^"]+)"/g)].map(m=>m[1]);
 assert.equal(new Set(ids).size,ids.length);
 for(const t of tasks)assert.ok(ids.includes(`task-${t.id}`));
});

test('Markdown has every full program and answer; local links resolve',()=>{
 const md=read('TASKS.md');
 assert.equal([...md.matchAll(/^### \d{3}\./gm)].length,800);
 assert.equal([...md.matchAll(/<summary>Ответ и объяснение<\/summary>/g)].length,800);
 for(const t of tasks){assert.ok(md.includes(t.why));assert.ok(md.includes(read(`testdata/${String(t.id).padStart(3,'0')}.go`).trimEnd()));}
 for(const [,link] of html.matchAll(/href="([^"]+)"/g)){
  if(link.startsWith('#'))assert.ok(html.includes(`id="${link.slice(1)}"`),link);
  else if(!/^(https?:|\?)/.test(link))assert.ok(fs.existsSync(new URL(link,import.meta.url)),link);
 }
});

test('original tasks 001–400 remain byte-for-byte unchanged',()=>{
 assert.equal(createHash('sha256').update(JSON.stringify(tasks.slice(0,400))).digest('hex'),'ebbe8b6f94eb3e18358d396ea8ffe3bad05c5fd849f7d96892decc1f97f3143a');
 assert.deepEqual(tasks.map(t=>t.id),Array.from({length:800},(_,i)=>i+1));
});


test('tasks 001–600 remain unchanged; 200 exam questions have four options and one keyed answer',()=>{
 assert.equal(createHash('sha256').update(JSON.stringify(tasks.slice(0,600))).digest('hex'),'2d784b7debb9cbd0e8d3e45bbe22dd56d2f4d98e3d0a3a691ce9b58cc511485c');
 const exam=tasks.slice(600);
 assert.equal(exam.length,200);
 assert.equal([...html.matchAll(/<fieldset class="choices">/g)].length,200);
 for(const t of exam){
  assert.equal(t.choices.length,4);assert.equal(new Set(t.choices).size,4);
  assert.ok(Number.isInteger(t.correctIndex)&&t.correctIndex>=0&&t.correctIndex<4);
  assert.equal(t.choices[t.correctIndex],t.kind==='compile'?'Не компилируется':t.answer);
  const article=html.slice(html.indexOf('<article id="task-'+t.id+'"')).split('</article>')[0];
  assert.equal([...article.matchAll(/type="radio"/g)].length,4);
  assert.ok(article.includes('Правильный вариант: '+'ABCD'[t.correctIndex]+'.'));
  for(const choice of t.choices)assert.ok(decode(article).includes(choice));
 }
});
