import { readdir, writeFile } from 'node:fs/promises';

const apiDir = new URL('../src/api/', import.meta.url);
const entries = await readdir(new URL('generated/', apiDir), { withFileTypes: true });
const tags = entries
  .filter((entry) => entry.isFile() && entry.name.endsWith('.ts') && !entry.name.endsWith('.d.ts') && !['api.ts', 'index.ts'].includes(entry.name))
  .map((entry) => entry.name.slice(0, -3))
  .sort();

if (tags.length === 0) throw new Error('Orval 未生成任何 API 模块');
const names = new Set();
const modules = tags.map((tag) => {
  const name = tag.replace(/-([a-z])/g, (_, letter) => letter.toUpperCase());
  if (!/^[a-z][a-zA-Z0-9]*$/.test(name) || names.has(name)) {
    throw new Error(`API 模块名无效或重复: ${tag}`);
  }
  names.add(name);
  return { tag, name };
});

// Prefix import bindings so a tag cannot collide with the exported api binding.
const imports = modules.map(({ tag, name }) => `import * as _${name} from './generated/${tag}';`).join('\n');
const properties = modules.map(({ name }) => `  ${name}: _${name},`).join('\n');
await writeFile(new URL('index.ts', apiDir), `// 此文件由 pnpm gen:api 自动生成，请勿手动修改。\n${imports}\n\nexport const api = {\n${properties}\n} as const;\n\nexport default api;\n`);
