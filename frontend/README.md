# 博客前端

React + TypeScript + Vite，使用 Axios、SWR 和 Orval 接入 Go/Huma 后端。
参考：[Huma + Orval 前后端 API 自动生成方案](https://www.cnblogs.com/dingshaohua/p/22339996)。

## 开发

在 `frontend` 目录执行（推荐 Node.js 24；pnpm）：

```sh
pnpm install
cp .env.example .env.local
pnpm dev
```

浏览器默认请求 `/api`，Vite 转发到 `http://localhost:8080` 并移除 `/api` 前缀。
实际读取文章需要启动后端，并配置其 PostgreSQL、Redis；后端启动命令为 `cd ../backend && go run .`。
首页文章列表使用生成的 SWR Hook，支持加载状态、错误提示、刷新及分页。

文章管理入口：列表中点击标题或“查看详情”进入详情页；点击“新增文章”创建文章；列表和详情页均提供编辑、删除按钮。
新增与编辑支持标题、分类和正文，成功后跳转详情并刷新文章缓存；删除需要确认，删除当前页最后一篇时自动返回上一页。
路由为 `#/articles/new`、`#/articles/:id`、`#/articles/:id/edit`。
标题遵循当前后端规则，最多 10 个字符。正文按纯文本编辑和显示并保留换行。
分类来自后端：新增可选“未分类”，已有分类可更换；当前更新接口不支持清除已有分类。

环境变量在 `.env.local` 中设置，修改后重启 Vite：

| 变量 | 默认值 | 作用 |
| --- | --- | --- |
| `VITE_API_BASE_URL` | `/api` | 浏览器请求的 API 基地址，生产构建时也会使用 |
| `API_PROXY_TARGET` | `http://localhost:8080` | Vite 开发代理的后端地址 |
| `OPENAPI_INPUT` | `http://localhost:8080/openapi.json` | 后端 OpenAPI 接口地址 |

生产环境须将 `/api/*` 反向代理到后端并移除 `/api` 前缀，或在构建前设置 `VITE_API_BASE_URL` 为后端地址；使用跨域地址时后端需允许对应的前端来源。
`API_PROXY_TARGET` 仅用于开发，`OPENAPI_INPUT` 仅用于生成，均不会作为前端环境变量公开。

## 一键生成 API

```sh
pnpm gen:api
```

此命令依次执行：

1. 直接请求后端 `http://localhost:8080/openapi.json`，不保存本地 OpenAPI 文件。
2. Orval 按 Tag 生成类型、Axios 函数、SWR Query/Mutation Hooks。
3. 自动生成 `src/api/index.ts`，最后执行 TypeScript 检查。

生成前需要启动后端，并确保 OpenAPI 接口可访问；前端生成命令无需运行 Go。
新增模块须注册到 `backend/internal/bootstrap/api.go`，新增接口使用 `sharedApi.NewGroup` 并设置动作名（如 `list`、`get`、`create`、`update`、`remove`）。
完整 operationId 自动成为 `list-article` 等全局唯一名称；`article-type` 映射为 `api.articleType`。
同一 Tag 内的动作名必须唯一。删除动作使用 `remove`，避免 JavaScript 保留字。

可以通过 `.env.local` 或命令行指定其他后端 OpenAPI 地址（JSON/YAML 均可）：

```sh
OPENAPI_INPUT=http://localhost:8080/openapi.yaml pnpm gen:api
```

`src/api/generated/` 和 `src/api/index.ts` 一起提交到 Git，保证普通前端构建无需 Go 或在线后端。
生成文件不要手工修改，也已排除在 Biome 格式化之外。Orval 版本固定，更新时应重新生成并检查差异。

## 调用方式

普通函数返回后端 JSON 正文，保留分页结构，没有额外的 AxiosResponse 包装：

```ts
import api from '@/api';
import type { CreateArticleRequestBody } from '@/api/generated/models';

const result = await api.article.list({ page: 1, pageSize: 10 });
console.log(result.list, result.total);

const types = await api.articleType.list();
const feeds = await api.feed.list();
const body: CreateArticleRequestBody = { title: '新文章', content: '正文' };
const created = await api.article.create(body);
await api.article.remove(created.id);
```

React 组件内使用 Hooks：

```tsx
const articles = api.article.useList({ page: 1, pageSize: 10 });
// articles.data / error / isLoading / isValidating / mutate

const detail = api.article.useGet(articleId, {
  swr: { enabled: articleId > 0 },
});

const createArticle = api.article.useCreate();
async function submit() {
  await createArticle.trigger({ title: '新文章' });
  await articles.mutate();
}
```

带查询参数的 Hook，第一个参数为查询参数，第二个参数为 `{ swr, request }`；无查询参数的接口，如 `api.feed.useList()`，其配置直接放在第一个参数。
普通函数最后一个参数可传 Axios 请求选项（如 `signal`、`headers`、`timeout`）；Hook 放在 `request` 内。
修改和删除接口成功时为 `204 No Content`，不应读取响应字段。

`src/api/custom-axios.ts` 统一维护基地址、15 秒超时、配置和请求头合并。
Huma 错误通过 `AxiosError` 保留在 `error.response.data`，例如 `detail`；需要鉴权或拦截器时修改 `axiosInstance`。

## 检查

```sh
pnpm typecheck
pnpm test:api
pnpm build
```

`test:api` 使用生成函数和 Axios 适配器验证参数、分页响应、请求头、空响应和错误传递，无需后端。
后端 `go test ./...` 验证 HTTP 提供的 Schema 与注册定义一致、组件名称合法、operationId 唯一及空响应约定。

CI 或提交前检查生成结果是否最新（需可访问后端 OpenAPI 地址）：

```sh
pnpm gen:api
git diff --exit-code -- src/api/generated src/api/index.ts
```

先将新增生成文件加入版本控制，再使用以上差异检查。


## Markdown 正文

文章详情使用 `react-markdown` 和 `remark-gfm` 渲染正文；编辑页可切换源码与预览。支持标题、列表、引用、代码块、表格、任务列表、链接和图片。原始 HTML 不执行，保留链接协议过滤。

运行 `pnpm test:markdown` 验证渲染及外部内容处理，运行 `pnpm build` 检查类型和构建。
