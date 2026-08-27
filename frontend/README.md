# React + TypeScript + Vite (Biome 版)

这个模板提供了一个在 Vite 中运行 React 的最小化配置，支持热更新（HMR）以及基于 **Biome** 的超快速代码检查与格式化规则。

## 官方插件

目前支持两种官方插件：

*   [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) 使用 [Oxc](https://oxc.rs) 进行转换。
*   [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) 使用 [SWC](https://swc.rs/) 进行转换。

## React Compiler

本模板已启用 **React Compiler (React 编译器)**。更多信息请参阅 [官方文档](https://react.dev/learn/react-compiler)。

*注意：启用编译器会对 Vite 的开发和构建性能产生一定影响。*


## 常用命令
```
开发环境启动：npm run dev
生产环境构建：npm run build
本地预览构建结果：npm run preview
```


## 使用 Biome 进行代码规范管理

本项目使用 [Biome](https://biomejs.dev/) 取代了传统的 ESLint 和 Prettier，提供秒级的代码检查、格式化及导入排序。执行全面检查并自动修复： `npm run biome`

该命令会同时进行代码检查（Lint）、格式化（Format）以及导入语句排序（Organize Imports），并直接写入文件。

Biome 会自动读取 `tsconfig` 中的 `paths`。如果需要手动忽略某些路径，请在 `linter.ignore` 中添加。

如果你需要调整规则（例如禁用 `rel="noopener"` 的自动添加），请修改根目录下的 `biome.json`：

为了获得最佳开发体验，建议在 VS Code 中安装 [Biome 扩展](https://marketplace.visualstudio.com/items?itemName=biomejs.biome)，并在设置中开启 **"Format on Save"** (保存时自动格式化)。