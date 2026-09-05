import assert from 'node:assert/strict';
import { after, before, test } from 'node:test';
import { createElement } from 'react';
import { renderToStaticMarkup } from 'react-dom/server';
import { createServer } from 'vite';

let server;
let MarkdownContent;

before(async () => {
  server = await createServer({ server: { middlewareMode: true, hmr: false, watch: null }, appType: 'custom' });
  MarkdownContent = (await server.ssrLoadModule('/src/components/markdown-content.tsx')).default;
});

after(async () => {
  await server?.close();
});

function render(content) {
  return renderToStaticMarkup(createElement(MarkdownContent, { content }));
}

test('renders Markdown headings, code, lists, quotes and GFM tables/tasks', () => {
  const html = render('## 标题\n\n**重点**与`inline`\n\n> 引用\n\n- 第一项\n- 第二项\n\n```go\nfmt.Println("hello")\n```\n\n| 名称 | 数值 |\n| --- | ---: |\n| 月球 | 1 |\n\n- [x] 完成\n- [ ] 待办\n\n~~旧文字~~');
  for (const expected of ['<h2>标题</h2>', '<strong>重点</strong>', '<code>inline</code>', '<blockquote>', '<ul>', '<pre>', 'language-go', '<table>', '<th', '<td', 'type="checkbox"', 'disabled=""', 'checked=""', '<del>旧文字</del>']) {
    assert.ok(html.includes(expected), expected);
  }
});

test('escapes raw HTML and prevents executable URL protocols', () => {
  const html = render('<script>alert(1)</script>\n\n<img src=x onerror="alert(2)">\n\n[bad](javascript:alert%281%29)\n\n![bad](data:text/html;base64,abcd)');
  assert.doesNotMatch(html, /<script|<img[^>]+onerror=|(?:href|src)="(?:javascript|data):/i);
  assert.match(html, /&lt;script&gt;/);
  assert.doesNotMatch(html, /src=""/);
});

test('external links retain sources and internal links stay in the same tab', () => {
  const external = render('[原文](https://go.dev/blog/context)');
  assert.match(external, /href="https:\/\/go.dev\/blog\/context"/);
  assert.match(external, /target="_blank"/);
  assert.match(external, /rel="noopener noreferrer"/);
  const internal = render('[文章](/articles/1)');
  assert.match(internal, /href="\/articles\/1"/);
  assert.doesNotMatch(internal, /target=/);
});

test('keeps plain text and Unicode, renders image alt text without forwarding referrers', () => {
  const html = render('普通中文正文 🌙\n\n第二段\n\n![月球](https://example.com/moon.png)');
  assert.match(html, /普通中文正文 🌙/);
  assert.match(html, /<p>第二段<\/p>/);
  assert.match(html, /alt="月球"/);
  assert.match(html, /loading="lazy"/);
  assert.match(html, /referrerPolicy="no-referrer"/i);
});
