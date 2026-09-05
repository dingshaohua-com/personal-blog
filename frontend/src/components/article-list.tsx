import { ArrowUpRight, BookOpen, ChevronLeft, ChevronRight, Plus, RefreshCw } from 'lucide-react';
import { Link, useSearchParams } from 'react-router';
import api from '@/api';
import DeleteArticleButton from '@/components/delete-article-button';
import { Button } from '@/components/ui/button';
import { apiErrorMessage } from '@/utils/article';

const PAGE_SIZE = 6;

export default function ArticleList() {
  const [searchParams, setSearchParams] = useSearchParams();
  const requestedPage = Number(searchParams.get('page'));
  const page = Number.isSafeInteger(requestedPage) && requestedPage > 0 ? requestedPage : 1;
  const setPage = (next: number) => {
    const params = new URLSearchParams(searchParams);
    if (next > 1) params.set('page', String(next));
    else params.delete('page');
    setSearchParams(params);
  };
  const { data, error, isLoading, isValidating, mutate } = api.article.useList({ page, pageSize: PAGE_SIZE }, { swr: { revalidateOnFocus: false, shouldRetryOnError: false } });

  return (
    <section className="space-y-6" aria-labelledby="articles-title" aria-busy={isValidating}>
      <div className="flex flex-wrap items-center justify-between gap-4">
        <div>
          <h2 id="articles-title" className="text-lg font-semibold">
            最新文章
          </h2>
          <p className="mt-1 text-xs text-muted-foreground">{data && !error ? `共 ${data.total} 篇 · ` : ''}每页 6 篇</p>
        </div>
        <div className="flex gap-2">
          <Button variant="outline" disabled={isValidating} onClick={() => void mutate()} aria-label="刷新文章">
            <RefreshCw className={isValidating ? 'animate-spin' : ''} aria-hidden="true" />
            刷新
          </Button>
          <Button asChild className="bg-emerald-700 text-white hover:bg-emerald-800">
            <Link to="/articles/new">
              <Plus aria-hidden="true" />
              写文章
            </Link>
          </Button>
        </div>
      </div>
      {isLoading && (
        <p role="status" className="rounded-xl border bg-background p-8 text-center text-muted-foreground">
          正在加载文章…
        </p>
      )}
      {error && (
        <p role="alert" className="rounded-xl border border-destructive/20 bg-background p-6 text-destructive">
          {apiErrorMessage(error, '文章加载失败，请点击刷新重试。')}
        </p>
      )}
      {!isLoading && !error && data && (
        <>
          {data.list?.length ? (
            <ul className="grid gap-4 md:grid-cols-2">
              {data.list.map((article) => (
                <li key={article.id} className="group flex min-w-0 flex-col rounded-2xl border bg-card p-5 text-card-foreground shadow-sm transition-shadow hover:shadow-md sm:p-6">
                  <div className="mb-5 flex flex-wrap items-center justify-between gap-2 text-xs">
                    <span className="rounded-md bg-emerald-700/8 px-2.5 py-1 font-medium text-emerald-800 dark:text-emerald-300">{article.typeName || '未分类'}</span>
                    <time className="text-muted-foreground" dateTime={article.createdAt}>
                      {new Date(article.createdAt).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })}
                    </time>
                  </div>
                  <h3 className="break-words text-xl font-semibold leading-8">
                    <Link to={`/articles/${article.id}`} className="transition-colors hover:text-emerald-700 dark:hover:text-emerald-400">
                      {article.title}
                    </Link>
                  </h3>
                  {article.description && <p className="mt-3 line-clamp-2 text-sm leading-7 text-muted-foreground">{article.description}</p>}
                  <div className="mt-auto flex flex-wrap items-center justify-between gap-3 pt-6">
                    <Link to={`/articles/${article.id}`} className="inline-flex items-center gap-1 text-sm font-medium text-emerald-700 hover:underline dark:text-emerald-400" aria-label={`阅读《${article.title}》`}>
                      阅读全文
                      <ArrowUpRight className="size-4" aria-hidden="true" />
                    </Link>
                    <div className="flex gap-1">
                      <Button asChild variant="ghost">
                        <Link to={`/articles/${article.id}/edit`} aria-label={`编辑《${article.title}》`}>
                          编辑
                        </Link>
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
                  </div>
                </li>
              ))}
            </ul>
          ) : (
            <div className="rounded-2xl border border-dashed bg-background p-12 text-center">
              <BookOpen className="mx-auto mb-4 size-8 text-muted-foreground" aria-hidden="true" />
              <p className="font-medium">{data.total > 0 ? '这一页没有文章' : '还没有文章'}</p>
              <p className="mt-2 text-sm text-muted-foreground">{data.total > 0 ? '返回第一页，继续阅读。' : '点击「写文章」，留下第一篇记录。'}</p>
              {page > 1 && (
                <Button className="mt-4" variant="outline" onClick={() => setPage(1)}>
                  返回第一页
                </Button>
              )}
            </div>
          )}
          <nav aria-label="文章分页" className="flex flex-wrap items-center justify-between gap-3 border-t pt-6">
            <Button variant="outline" disabled={page <= 1 || isValidating} onClick={() => setPage(page - 1)}>
              <ChevronLeft aria-hidden="true" />
              上一页
            </Button>
            <p className="text-xs text-muted-foreground sm:text-sm" aria-live="polite">
              第 {data.page} / {Math.max(1, data.totalPage)} 页 · 共 {data.total} 篇
            </p>
            <Button variant="outline" disabled={page >= data.totalPage || isValidating} onClick={() => setPage(page + 1)}>
              下一页
              <ChevronRight aria-hidden="true" />
            </Button>
          </nav>
        </>
      )}
    </section>
  );
}
