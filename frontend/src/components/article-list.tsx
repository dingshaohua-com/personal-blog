import { Link, useSearchParams } from 'react-router';
import api from '@/api';
import DeleteArticleButton from '@/components/delete-article-button';
import { Button } from '@/components/ui/button';

export default function ArticleList() {
  const [searchParams, setSearchParams] = useSearchParams();
  const requestedPage = Number(searchParams.get('page'));
  const page = Number.isSafeInteger(requestedPage) && requestedPage > 0 ? requestedPage : 1;
  const setPage = (next: number) => setSearchParams(next > 1 ? { page: String(next) } : {});
  const { data, error, isLoading, isValidating, mutate } = api.article.useList({ page, pageSize: 10 }, { swr: { revalidateOnFocus: false, shouldRetryOnError: false } });

  return (
    <section className="mx-auto max-w-3xl space-y-4 p-4" aria-labelledby="articles-title">
      <div className="flex items-center justify-between gap-4">
        <h1 id="articles-title" className="text-2xl font-semibold">
          最新文章
        </h1>
        <div className="flex gap-2">
          <Button asChild>
            <Link to="/articles/new">新增文章</Link>
          </Button>
          <Button variant="outline" disabled={isValidating} onClick={() => void mutate()}>
            {isValidating ? '刷新中…' : '刷新'}
          </Button>
        </div>
      </div>
      {isLoading && <p role="status">正在加载文章…</p>}
      {error && (
        <p role="alert" className="text-destructive">
          {error.response?.data.detail || '文章加载失败，请稍后重试。'}
        </p>
      )}
      {!isLoading && !error && data && (
        <>
          {data.list?.length ? (
            <ul className="space-y-3">
              {data.list.map((article) => (
                <li key={article.id} className="rounded-lg border bg-card p-4 text-card-foreground">
                  <h2 className="break-words text-lg font-medium">
                    <Link to={`/articles/${article.id}`} className="hover:underline">
                      {article.title}
                    </Link>
                  </h2>
                  {article.description && <p className="mt-2 text-muted-foreground">{article.description}</p>}
                  <div className="mt-3 flex gap-3 text-sm text-muted-foreground">
                    {article.typeName && <span>{article.typeName}</span>}
                    <time dateTime={article.createdAt}>{new Date(article.createdAt).toLocaleDateString('zh-CN')}</time>
                  </div>
                  <div className="mt-3 flex gap-2">
                    <Button asChild variant="outline">
                      <Link to={`/articles/${article.id}`}>查看详情</Link>
                    </Button>
                    <Button asChild variant="outline">
                      <Link to={`/articles/${article.id}/edit`}>编辑</Link>
                    </Button>
                    <DeleteArticleButton
                      id={article.id}
                      title={article.title}
                      onDeleted={() => {
                        if (data.list?.length === 1 && page > 1) setPage(page - 1);
                        else void mutate();
                      }}
                    />
                  </div>
                </li>
              ))}
            </ul>
          ) : (
            <p className="text-muted-foreground">暂无文章。</p>
          )}
          <div className="flex items-center justify-between gap-3">
            <Button variant="outline" disabled={page <= 1} onClick={() => setPage(page - 1)}>
              上一页
            </Button>
            <span className="text-sm text-muted-foreground">
              第 {data.page} 页 · 共 {data.total} 篇
            </span>
            <Button variant="outline" disabled={page >= data.totalPage} onClick={() => setPage(page + 1)}>
              下一页
            </Button>
          </div>
        </>
      )}
    </section>
  );
}
