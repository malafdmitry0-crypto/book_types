(() => {
  const chapters = JSON.parse(document.getElementById('book-data').textContent);
  const articles = [...document.querySelectorAll('.chapter')];
  const links = [...document.querySelectorAll('[data-chapter]')];
  const contents = document.getElementById('contents');
  const sidebar = document.getElementById('sidebar');
  const menu = document.getElementById('menu');
  const search = document.getElementById('search');
  const clear = document.getElementById('clear-search');
  const status = document.getElementById('search-status');
  let noticeTimer;
  function closeMenu() { sidebar.classList.remove('open'); menu.setAttribute('aria-expanded', 'false'); }
  function showChapter(focus = false) {
    let id;
    try { id = decodeURIComponent(location.hash.slice(1)); } catch { id = ''; }
    if (id === 'main') {
      document.getElementById('main').focus({preventScroll:true});
      return;
    }
    const target = document.getElementById(id);
    const article = target?.closest('.chapter');
    articles.forEach(a => { a.hidden = a !== article; });
    contents.hidden = !!article;
    links.forEach(a => {
      if (a.dataset.chapter === article?.id) {a.setAttribute('aria-current', 'page'); a.closest('details').open = true;}
      else a.removeAttribute('aria-current');
    });
    document.title = article ? `${chapters.find(c => c.id === article.id).title} — Go / Алгоритмы` : 'Обобщённое программирование в Go — учебник';
    closeMenu();
    if (focus) {
      const heading = article?.querySelector('h1') || contents.querySelector('h1');
      heading.setAttribute('tabindex', '-1');
      heading.focus({preventScroll:true});
      if (target && target !== article && article) target.scrollIntoView();
      else window.scrollTo({top:0,behavior:'instant'});
    }
  }
  menu.addEventListener('click', () => {
    const open = sidebar.classList.toggle('open'); menu.setAttribute('aria-expanded', String(open));
    if(open) search.focus();
  });
  document.addEventListener('keydown', e => { if(e.key === 'Escape' && sidebar.classList.contains('open')) {closeMenu();menu.focus();} });
  document.getElementById('main').addEventListener('click', closeMenu);
  function filter() {
    const query = search.value.trim().toLocaleLowerCase('ru');
    let count = 0;
    links.forEach((a,i) => {
      const match = !query || chapters[i].text.toLocaleLowerCase('ru').includes(query);
      a.hidden = !match; if(match)count++;
    });
    document.querySelectorAll('.nav-group').forEach(group => {group.hidden = !!query && !group.querySelector('[data-chapter]:not([hidden])'); if(query) group.open = true;});
    clear.hidden = !query; status.hidden = !query;
    status.textContent = count ? `Найдено разделов: ${count}` : 'Ничего не найдено. Попробуйте другой термин.';
  }
  search.addEventListener('input', filter);
  clear.addEventListener('click', () => {search.value='';filter();search.focus();});
  function notify(text) {const notice=document.getElementById('notice');notice.textContent=text;clearTimeout(noticeTimer);noticeTimer=setTimeout(()=>{notice.textContent='';},2600);}
  document.querySelectorAll('.copy').forEach(button => button.addEventListener('click', async () => {
    const code=button.closest('.code-block').querySelector('code');
    try {await navigator.clipboard.writeText(code.textContent);notify('Код скопирован');}
    catch {const range=document.createRange();range.selectNodeContents(code);const selection=window.getSelection();selection.removeAllRanges();selection.addRange(range);notify('Код выделен. Нажмите ⌘C или Ctrl+C.');}
  }));
  document.getElementById('print').addEventListener('click', () => window.print());
  
  const approaches = {
    find: ['index, err := algorithms.Find(query)\n// query: fractionQuery\n// Результат — индекс в исходном []fraction.Fraction', 'Алгоритм знает Len и Match. Адаптер хранит дроби и условие; математическое равенство проверяет Fraction.Equal. Тип элемента не проходит через API алгоритма.', 'gp-early-search'],
    map: ['err := algorithms.Map(job)\n// job: fractionText\n// Source: []fraction.Fraction, Target: []string', 'Алгоритм вызывает Apply по индексам. Адаптер преобразует дробь и записывает строку в заранее выделенный результат. При ошибке уже заполненный префикс сохраняется.', 'gp-early-map-sum'],
    sum: ['err := algorithms.Sum(&sum)\n// sum: moneySum\n// Начальное Total: money.Money{Currency: "USD"}', 'Алгоритм знает Len и Add. Адаптер хранит типизированный аккумулятор; Money.Add проверяет валюту и переполнение. Начальное значение выбирает вызывающий.', 'gp-early-map-sum'],
    sort: ['err := algorithms.Sort(boxes(values))\n// values: []boxint.BoxInt\n// Адаптер предоставляет Len, Compare и Swap', 'Алгоритм сравнивает и переставляет элементы по индексам. boxes связывает эти операции с BoxInt. Учебная сортировка вставками стабильна; ошибка может оставить частично изменённый порядок.', 'gp-early-contracts']
  };
  document.getElementById('approach').addEventListener('change', e => {
    const [code,note,id] = approaches[e.target.value];
    document.getElementById('approach-code').textContent = code;
    document.getElementById('approach-note').textContent = note;
    document.getElementById('approach-link').href = '#' + id;
  });

  window.addEventListener('hashchange', () => showChapter(true));
  showChapter();
})();
