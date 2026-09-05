import { defineConfig } from 'orval';
import { loadEnv } from 'vite';

const env = loadEnv(process.env.NODE_ENV || 'development', process.cwd(), '');

export default defineConfig({
  blog: {
    input: env.OPENAPI_INPUT || 'http://localhost:8080/openapi.json',
    output: {
      clean: true,
      tsconfig: './tsconfig.app.json',
      target: './src/api/generated/api.ts',
      schemas: './src/api/generated/models',
      client: 'swr',
      httpClient: 'axios',
      mode: 'tags',
      override: {
        mutator: {
          path: './src/api/custom-axios.ts',
          name: 'customAxios',
        },
        operationName: (operation, route, verb) => {
          const tag = operation.tags?.[0];
          const id = operation.operationId;
          const suffix = `-${tag}`;
          if (operation.tags?.length !== 1 || !tag || !id?.endsWith(suffix)) {
            throw new Error(`${verb.toUpperCase()} ${route}: 需要唯一 Tag 和“动作-Tag”格式的 operationId`);
          }
          const action = id.slice(0, -suffix.length);
          if (!/^[a-z][a-zA-Z0-9]*$/.test(action) || action === 'delete') {
            throw new Error(`无效的 API 动作名: ${id}；删除请使用 remove`);
          }
          return action;
        },
      },
    },
    hooks: {
      afterAllFilesWrite: {
        command: 'node scripts/gen-api-index.mjs',
        injectGeneratedDirsAndFiles: false,
      },
    },
  },
});
