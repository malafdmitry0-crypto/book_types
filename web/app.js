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
    concrete: ['names := MapUserNames(users, name) // []string', 'Компилятор проверяет User → string. Для Order или другого результата нужна новая специализация.', 'gp-concrete'],
    closure: ['names := make([]string, len(users))\nMapInto(len(users), func(i int) {\n    names[i] = name(users[i])\n})', 'Типы сохранены внутри замыкания. Создание результата и согласование размеров лежат на вызывающем. Здесь показан ненулевой вход.', 'gp-closures'],
    interface: ['names := make([]string, len(users))\nRunMap(UserNameJob{users, names, name})', 'Общий алгоритм знает Len и Apply. Отношение User → string скрыто в конкретном адаптере. Здесь показан ненулевой вход.', 'gp-interfaces'],
    any: ['raw := MapAny(boxed, func(x interface{}) interface{} {\n    return name(x.(User))\n}) // []interface{}, ещё не []string', 'До вызова нужен []interface{}, после — поэлементное извлечение строк. Callback и содержимое контейнера не связаны статически.', 'gp-any'],
    reflection: ['raw, err := MapReflect(users, name)\n// После проверки err:\nnames := raw.([]string)', 'Валидация связывает типы во время исполнения. Динамический результат — []string, статический — interface{}.', 'gp-reflection'],
    generation: ['names := MapUsersGenerated(users, name) // []string', 'Типобезопасная специализация выпущена из шаблона до компиляции. Новая комбинация типов требует генерации.', 'gp-generation'],
    generic: ['names := Map(users, name) // []string\n// Map[E, R any]([]E, func(E) R) []R', 'Одна сигнатура выражает связь E → R. Компилятор проверяет callback и сохраняет конкретный тип результата.', 'gp-generics'],
    iterator: ['names := MapSeq(slices.Values(users), name)\n// iter.Seq[string]; работа начнётся при обходе', 'Сохранена связь E → R и скрыт источник элементов. Срез результата пока не создан: вычисление ленивое.', 'gp-iterators']
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
