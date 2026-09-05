import { Link, useNavigate, useParams } from 'react-router';
import api from '@/api';
import DeleteArticleButton from '@/components/delete-article-button';
import MarkdownContent from '@/components/markdown-content';
import { Button } from '@/components/ui/button';
import { apiErrorMessage, articleIdFromParam } from '@/utils/article';

export default function ArticleDetail() {
  const { id: param } = useParams();
  const id = articleIdFromParam(param);
  const navigate = useNavigate();
  const { data, error, isLoading, isValidating, mutate } = api.article.useGet(id ?? 0, { swr: { enabled: id !== undefined, revalidateOnFocus: false, shouldRetryOnError: false } });

  return (
    <main className="mx-auto max-w-3xl space-y-6 p-4 py-8">
      <Button asChild variant="outline">
        <Link to="/">返回文章列表</Link>
      </Button>
      {!id && <p role="alert">文章地址无效。</p>}
      {id && isLoading && <p role="status">正在加载文章…</p>}
      {id && error && (
        <div className="space-y-3">
          <p role="alert" className="text-destructive">
            {error.response?.status === 404 ? '文章不存在或已被删除。' : apiErrorMessage(error, '文章加载失败，请重试。')}
          </p>
          <Button variant="outline" disabled={isValidating} onClick={() => void mutate()}>
            重试
          </Button>
        </div>
      )}
      {id && data && !error && (
        <article className="space-y-6">
          <header className="space-y-4 border-b pb-5">
            <h1 className="break-words text-3xl font-semibold">{data.title}</h1>
            <div className="flex flex-wrap items-center gap-3 text-sm text-muted-foreground">
              <span>{data.typeName || '未分类'}</span>
              <time dateTime={data.createdAt}>{new Date(data.createdAt).toLocaleString('zh-CN')}</time>
            </div>
            <div className="flex gap-2">
              <Button asChild variant="outline">
                <Link to={`/articles/${id}/edit`}>编辑</Link>
              </Button>
              <DeleteArticleButton id={id} title={data.title} onDeleted={() => navigate('/', { replace: true })} />
            </div>
          </header>
          <MarkdownContent content={data.content || '暂无正文。'} />
        </article>
      )}
    </main>
  );
}
