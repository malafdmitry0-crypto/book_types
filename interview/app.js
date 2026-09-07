(() => {
 const data=JSON.parse(document.getElementById('task-data').textContent);
 const topic=document.getElementById('topic'),query=document.getElementById('query');
 const collection=document.getElementById('collection');
 const status=document.getElementById('status'),empty=document.getElementById('empty');
 const articles=new Map(data.map(t=>[t.id,document.getElementById('task-'+t.id)]));
 const initialSet=new URLSearchParams(location.search).get('set');
 if(['extra','original','exam'].includes(initialSet))collection.value=initialSet;
 let single=null;
 function inCollection(id){return collection.value==='all'||(collection.value==='original'?id<=400:collection.value==='extra'?id>400&&id<=600:id>600);}
 function matching(){
  const term=query.value.trim().toLowerCase();
  return data.filter(t=>inCollection(t.id)&&(topic.value==='all'||String(t.topic)===topic.value)&&(!term||t.search.includes(term)));
 }
 function render(){
  const visible=matching().filter(t=>single===null||t.id===single),ids=new Set(visible.map(t=>t.id));
  for(const [id,article] of articles)article.hidden=!ids.has(id);
  const label=collection.value==='exam'?' · экзамен 601–800':collection.value==='extra'?' · адаптеры 401–600':collection.value==='original'?' · первые 001–400':'';
  status.textContent=single===null?`Показано задач: ${visible.length} из ${data.length}${label}`:`Задача ${String(single).padStart(3,'0')} из ${data.length}`;
  empty.hidden=visible.length!==0;
  return visible;
 }
 function filter(){
  single=null;const url=new URL(location.href);url.hash='';
  if(collection.value==='all')url.searchParams.delete('set');else url.searchParams.set('set',collection.value);
  history.replaceState(null,'',url);render();
 }
 function revealHash(){
  if(location.hash==='#tasks'){
   single=null;topic.value='all';query.value='';
   const set=new URLSearchParams(location.search).get('set');
   collection.value=['extra','original','exam'].includes(set)?set:'all';
   render();return;
  }
  const m=location.hash.match(/^#task-(\d+)$/);if(!m)return;
  const id=Number(m[1]);if(!articles.has(id))return;
  single=id;topic.value='all';query.value='';
  if(!inCollection(id))collection.value='all';
  render();articles.get(id).scrollIntoView();articles.get(id).querySelector('h2').focus({preventScroll:true});
 }
 topic.addEventListener('change',filter);query.addEventListener('input',filter);
 collection.addEventListener('change',()=>{topic.value='all';query.value='';filter();});
 document.getElementById('reset').addEventListener('click',()=>{collection.value='all';topic.value='all';query.value='';filter();query.focus();});
 document.getElementById('random').addEventListener('click',()=>{
  let pool=matching();if(pool.length>1)pool=pool.filter(t=>t.id!==single);
  if(!pool.length){render();return;}
  single=pool[Math.floor(Math.random()*pool.length)].id;
  articles.get(single).querySelector('details').open=false;render();history.replaceState(null,'','#task-'+single);
  articles.get(single).scrollIntoView({block:'start'});articles.get(single).querySelector('h2').focus({preventScroll:true});
 });
 window.addEventListener('hashchange',revealHash);render();revealHash();
})();
