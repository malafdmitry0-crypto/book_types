import fs from 'node:fs';
import path from 'node:path';
import {fileURLToPath} from 'node:url';
const root=path.resolve(path.dirname(fileURLToPath(import.meta.url)),'..');
const text=fs.readFileSync(path.join(root,'lessons/journey/results/bench-go1.25.4.txt'),'utf8');
const rows=new Map();
for(const line of text.split('\n')){
 const m=line.match(/^(Benchmark\S+)\s+\d+\s+([\d.]+) ns\/op\s+(\d+) B\/op\s+(\d+) allocs\/op/);
 if(!m)continue;
 const values=rows.get(m[1])??[];values.push([+m[2],+m[3],+m[4]]);rows.set(m[1],values);
}
const names=['concrete','closure','interface','empty-interface','reflection','generation','generic','iterator'];
const stats=name=>{const v=rows.get(name);if(!v||v.length!==3)throw new Error('Expected 3 runs: '+name);const ns=v.map(x=>x[0]).sort((a,b)=>a-b);return {median:ns[1],min:ns[0],max:ns[2],bytes:v[1][1],allocs:v[1][2]};};
let out='| Подход | Медиана, ns/op | Диапазон, ns/op | B/op | allocs/op |\n|---|---|---|---|---|\n';
for(const name of names){const s=stats('BenchmarkJourney/FindCheap/'+name);out+=`| ${name} | ${s.median} | ${s.min}–${s.max} | ${s.bytes} | ${s.allocs} |\n`;}
out+='\n| Подход | FindMath, µs/op | Sum, µs/op | Filter, µs/op |\n|---|---|---|---|\n';
for(const name of names){out+=`| ${name} | `+['FindMath','Sum','Filter'].map(op=>(stats(`BenchmarkJourney/${op}/${name}`).median/1000).toFixed(2)).join(' | ')+' |\n';}
const p=stats('BenchmarkAnyPrepared');out+=`\nAnyFind с заранее подготовленным входом: медиана ${p.median} ns/op, диапазон ${p.min}–${p.max}, ${p.bytes} B/op, ${p.allocs} allocs/op.\n`;
fs.writeFileSync(path.join(root,'lessons/journey/results/tables.md'),out);
console.log(out);
