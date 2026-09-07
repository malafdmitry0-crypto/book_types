import fs from 'node:fs';
import os from 'node:os';
import path from 'node:path';
import {execFile} from 'node:child_process';
import {promisify} from 'node:util';
import {tasks} from './all-tasks.mjs';
const exec=promisify(execFile);
const base=fs.mkdtempSync(path.join(os.tmpdir(),'gotypes-interview-'));
const env={...process.env,GOCACHE:process.env.GOCACHE||'/tmp/gotypes-review-cache',GOTOOLCHAIN:process.env.GOTOOLCHAIN||'go1.23.0',GOWORK:'off'};
const selected=process.argv.slice(2).filter(x=>/^\d+$/.test(x)).map(Number);
const queue=tasks.filter(t=>!selected.length||selected.includes(t.id));
const results=[];let cursor=0;
try{
 await Promise.all(Array.from({length:4},async()=>{
  while(cursor<queue.length){const t=queue[cursor++];const src=path.join(base,t.id+'.go');const bin=path.join(base,String(t.id));fs.writeFileSync(src,t.source);
   let error;try{await exec('go',['build','-o',bin,src],{env,timeout:60000});}catch(e){error=e;}
   let actual='',ok=false;
   if(t.kind==='compile'){actual=error?.stderr||'';ok=!!error&&actual.includes(t.answer);}
   else if(error){actual=error.stderr||error.message;}
   else{try{const r=await exec(bin,[],{timeout:10000});actual=r.stdout.trimEnd();ok=actual===t.answer;}catch(e){actual=(e.stdout||'')+(e.stderr||e.message);}}
   results.push({id:t.id,ok,kind:t.kind,actual});if(!ok)console.log(JSON.stringify({id:t.id,title:t.title,expected:t.answer,actual}));
  }
 }));
 const version=(await exec('go',['version'],{env})).stdout.trim();
 const report={version,checked:results.length,passed:results.filter(r=>r.ok).length,results:results.sort((a,b)=>a.id-b.id)};
 fs.writeFileSync('/tmp/gotypes-interview-verification.json',JSON.stringify(report,null,2)+'\n');
 if(report.checked===tasks.length&&report.passed===tasks.length)fs.writeFileSync(new URL('verification.json',import.meta.url),JSON.stringify({version,checked:report.checked,passed:report.passed,sourceDigest:(await import('node:crypto')).createHash('sha256').update(JSON.stringify(tasks)).digest('hex')},null,2)+'\n');
 console.log(`${report.passed}/${report.checked} verified (${version})`);if(report.passed!==report.checked)process.exitCode=1;
}finally{fs.rmSync(base,{recursive:true,force:true});}
