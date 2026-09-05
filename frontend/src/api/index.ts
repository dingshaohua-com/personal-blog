// 此文件由 pnpm gen:api 自动生成，请勿手动修改。
import * as _article from './generated/article';
import * as _articleType from './generated/article-type';
import * as _feed from './generated/feed';

export const api = {
  article: _article,
  articleType: _articleType,
  feed: _feed,
} as const;

export default api;
