# personal-blog

React + Vite 前端，Go + Huma 后端。通过 Taskfile 统一开发和构建；发布时将网页嵌入 Go 可执行文件。

## 准备

安装 Go（版本见 backend/go.mod）、Node.js、pnpm、Task 和 Air。Air 可通过 `go install github.com/air-verse/air@latest` 安装。

演示环境使用 `backend/.env` 配置 PostgreSQL 和 Redis，该文件允许提交到 Git。缺少时可从 `backend/.env.example` 复制。配置不会嵌入程序，由部署包单独携带。

## 常用命令

在仓库根目录执行：

```shell
task install
task dev
task build
```

- `task dev`：并行启动 Air 和 Vite，浏览器打开 Vite 打印的地址。前端 `/api` 请求由 Vite 代理给 Go；开发不需要 dist。
- `task build`：前端构建 → 复制 dist 到 Go 模块内 → 使用 production 标签编译并嵌入网页 → 输出 `dist/personal-blog.exe`（Windows）或 `dist/personal-blog`。

Taskfile 通过 includes 拆分到 `tasks/frontend.yml` 和 `tasks/backend.yml`。可单独执行 `task frontend:build`、`task backend:dev` 等任务。

## 运行构建产物

将可执行文件和运行配置放到部署目录，从该目录运行程序。Windows 示例：

```powershell
cd dist
# 在此目录放置配置好的 .env，或通过进程环境变量配置。
.\personal-blog.exe
```

浏览器打开 `http://localhost:18080`，端口由 `HTTP_PORT` 控制。网页资源已在程序内，运行时无需 Node.js、Task 或额外 dist 网页目录；PostgreSQL 和 Redis 仍需可用。

后端同时支持 `/api/...` 和原有无前缀接口，保留 `/docs`、`/openapi.json`。前端默认使用同源 `/api`；不要在生产构建中将 `VITE_API_BASE_URL` 设为开发机地址。当前使用 Hash 路由，不对缺失的 JS/CSS 做首页兜底。

修改前端后重新执行 `task build`。直接 `go build` 是开发版本，不含网页；单独执行 `task backend:build` 会嵌入上一次准备的网页资源，发布优先使用根任务。

构建只生成本机平台产物，不执行上传、服务重启或数据库迁移。

## GitHub Actions

`.github/workflows/deploy.yml` 在 main 推送或手动触发时执行：安装 Node.js、pnpm、Go、Task → `task install` → `task build` → 上传程序 → 重启服务并检查 OpenAPI 接口。

CI 与本地共用 Taskfile，前端资源随程序交付；Huma 自动提供 OpenAPI，无需安装 Swag 或复制 docs。

仓库 Secrets 需要配置 `REMOTE_HOST`、`SSH_PRIVATE_KEY`，可选 `REMOTE_PORT`（默认 22）。演示部署将 Git 中的 `backend/.env` 随程序上传，每次部署覆盖服务器 `/home/apps/blog-2026-server/.env`。服务器需有 tar、lsof、curl。

服务器继续使用工作流中的 `APP_PORT=8080`，启动时通过 `HTTP_PORT` 显式传入；本地默认 18080 不影响服务器。部署目录和程序名由工作流的 `DEPLOY_DIR`、`APP_NAME` 控制。
