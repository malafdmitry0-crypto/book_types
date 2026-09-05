import fs from 'node:fs';
import path from 'node:path';
import {fileURLToPath} from 'node:url';
const root=path.resolve(path.dirname(fileURLToPath(import.meta.url)),'..');
const chapters=JSON.parse(fs.readFileSync(path.join(root,'book/chapters.json'),'utf8'));
const codeMap=JSON.parse(fs.readFileSync(path.join(root,'book/code-map.json'),'utf8'));
const check=process.argv.includes('--check');
const expected=new Map();
const manifest=[];
for(const chapter of chapters){
 const source=fs.readFileSync(path.join(root,'book/chapters',chapter.filename),'utf8');
 for(const file of codeMap[chapter.id]??[]){if(!fs.statSync(path.join(root,file)).isFile())throw new Error('Missing code: '+file);}
 let number=0;
 for(const match of source.matchAll(/```go\n([\s\S]*?)\n```/g)){
  number++;
  const file=`lessons/snippets/testdata/${chapter.id}/${String(number).padStart(2,'0')}.go`;
  const content=`// Фрагмент ${number}: book/chapters/${chapter.filename}\n// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.\n// Это точная выдержка, не самостоятельная единица компиляции.\n${match[1]}\n`;
  expected.set(file,content);
  manifest.push({chapter:chapter.id,number,file,kind:'excerpt',runnableSources:codeMap[chapter.id]});
 }
}
expected.set('lessons/snippets/manifest.json',JSON.stringify(manifest,null,2)+'\n');
for(const [file,content]of expected){const target=path.join(root,file);if(check){if(!fs.existsSync(target)||fs.readFileSync(target,'utf8')!==content)throw new Error('Outdated excerpt: '+file);}else{fs.mkdirSync(path.dirname(target),{recursive:true});fs.writeFileSync(target,content);}}
// Удаляем только устаревшие сгенерированные выдержки внутри собственного каталога.
const base=path.join(root,'lessons/snippets/testdata');
function walk(dir){if(!fs.existsSync(dir))return;for(const name of fs.readdirSync(dir)){const target=path.join(dir,name);if(fs.statSync(target).isDirectory()){walk(target);continue;}const file=path.relative(root,target);if(!expected.has(file)){if(check)throw new Error('Stale excerpt: '+file);fs.unlinkSync(target);}}}
walk(base);
console.log(`${check?'Checked':'Synced'} ${manifest.length} Go excerpts from ${chapters.length} chapters`);
