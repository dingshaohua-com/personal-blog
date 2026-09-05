import assert from 'node:assert/strict';
import { after, before, test } from 'node:test';
import { AxiosError } from 'axios';
import { createServer } from 'vite';

let server;
let api;
let axiosInstance;
let articleHelpers;

before(async () => {
  server = await createServer({
    configFile: false,
    server: { middlewareMode: true, hmr: false, ws: false, watch: null },
    optimizeDeps: { noDiscovery: true, include: [] },
  });
  ({ api } = await server.ssrLoadModule('/src/api/index.ts'));
  ({ axiosInstance } = await server.ssrLoadModule('/src/api/custom-axios.ts'));
  articleHelpers = await server.ssrLoadModule('/src/utils/article.ts');
});

after(async () => {
  await server?.close();
});

test('article route IDs reject invalid values and cache invalidation stays within article resources', () => {
  assert.equal(articleHelpers.articleIdFromParam('12'), 12);
  for (const value of [undefined, '', '0', '-1', '1.5', '1e2', '12abc', '9007199254740992']) {
    assert.equal(articleHelpers.articleIdFromParam(value), undefined);
  }
  assert.equal(articleHelpers.isArticleCacheKey(api.article.getListKey({ page: 2 })), true);
  assert.equal(articleHelpers.isArticleCacheKey(api.article.getGetKey(12)), true);
  assert.equal(articleHelpers.isArticleCacheKey(api.articleType.getListKey()), false);
  assert.equal(articleHelpers.isArticleCacheKey(api.feed.getListKey()), false);
});

test('article create and update preserve body fields including cleared content', async () => {
  const body = { title: '新文章', typeId: 2, content: '第一行\n第二行' };
  const created = await api.article.create(body, {
    adapter: async (config) => {
      assert.equal(config.url, '/article');
      assert.equal(config.method, 'post');
      assert.deepEqual(JSON.parse(config.data), body);
      return { data: { id: 12 }, status: 200, statusText: 'OK', headers: {}, config };
    },
  });
  await api.article.update(
    created.id,
    { title: '修改文章', content: '' },
    {
      adapter: async (config) => {
        assert.equal(config.url, '/article/12');
        assert.equal(config.method, 'put');
        assert.deepEqual(JSON.parse(config.data), { title: '修改文章', content: '' });
        return { data: undefined, status: 204, statusText: 'No Content', headers: {}, config };
      },
    },
  );
});

test('generated article list preserves pagination, query params and per-call options', async () => {
  const body = { list: [], page: 2, pageSize: 10, total: 12, totalPage: 2 };
  const signal = new AbortController().signal;
  const result = await api.article.list(
    { page: 2, pageSize: 10, title: '测试' },
    {
      signal,
      timeout: 3000,
      adapter: async (config) => {
        assert.equal(config.url, '/article');
        assert.equal(config.baseURL, axiosInstance.defaults.baseURL);
        assert.equal(config.method, 'get');
        assert.deepEqual(config.params, { page: 2, pageSize: 10, title: '测试' });
        assert.equal(config.signal, signal);
        assert.equal(config.timeout, 3000);
        return { data: body, status: 200, statusText: 'OK', headers: {}, config };
      },
    },
  );
  assert.deepEqual(result, body);
  assert.deepEqual(api.article.getListKey({ page: 2 }), ['/article', { page: 2 }]);
});

test('mutation keeps generated JSON headers and merges caller headers', async () => {
  const body = { content: '新说说' };
  const result = await api.feed.create(body, {
    headers: { Authorization: 'Bearer test-only' },
    adapter: async (config) => {
      assert.equal(config.url, '/feed');
      assert.equal(config.method, 'post');
      assert.equal(config.headers.get('Content-Type'), 'application/json');
      assert.equal(config.headers.get('Authorization'), 'Bearer test-only');
      assert.deepEqual(JSON.parse(config.data), body);
      return { data: { id: 1, ...body }, status: 201, statusText: 'Created', headers: {}, config };
    },
  });
  assert.equal(result.id, 1);
});

test('remove uses the resource path and supports an empty 204 response', async () => {
  const result = await api.articleType.remove(7, {
    adapter: async (config) => {
      assert.equal(config.url, '/article-type/7');
      assert.equal(config.method, 'delete');
      return { data: undefined, status: 204, statusText: 'No Content', headers: {}, config };
    },
  });
  assert.equal(result, undefined);
});

test('backend problem details reach the caller as AxiosError', async () => {
  const problem = { status: 404, title: 'Not Found', detail: '文章不存在' };
  await assert.rejects(
    api.article.get(999, {
      adapter: async (config) => {
        throw new AxiosError('Request failed', 'ERR_BAD_REQUEST', config, undefined, {
          data: problem,
          status: 404,
          statusText: 'Not Found',
          headers: {},
          config,
        });
      },
    }),
    (error) => {
      assert.equal(error.isAxiosError, true);
      assert.deepEqual(error.response.data, problem);
      return true;
    },
  );
});
