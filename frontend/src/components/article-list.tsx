import { BookOpen, ChevronLeft, ChevronRight, Pencil, Plus, RefreshCw } from 'lucide-react';
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
          <Button variant="ghost" size="icon-lg" className="text-muted-foreground" disabled={isValidating} onClick={() => void mutate()} aria-label="刷新文章" title="刷新文章">
            <RefreshCw className={isValidating ? 'animate-spin' : ''} aria-hidden="true" />
          </Button>
          <Button asChild size="icon-lg" className="bg-emerald-700 text-white hover:bg-emerald-800">
            <Link to="/articles/new" aria-label="写文章" title="写文章">
              <Plus aria-hidden="true" />
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
            <ul className="divide-y border-t">
              {data.list.map((article) => (
                <li key={article.id} className="flex min-w-0 flex-col py-5 sm:py-6">
                  <div className="mb-2 flex flex-wrap items-center gap-3 text-xs">
                    <time className="text-muted-foreground" dateTime={article.createdAt}>
                      {new Date(article.createdAt).toLocaleDateString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit' })}
                    </time>
                    <span aria-hidden="true" className="text-muted-foreground/50">/</span>
                    <span className="font-medium text-emerald-800 dark:text-emerald-300">{article.typeName || '未分类'}</span>
                  </div>
                  <h3 className="break-words font-serif text-lg font-semibold leading-7 sm:text-xl">
                    <Link to={`/articles/${article.id}`} className="transition-colors hover:text-emerald-700 dark:hover:text-emerald-400">
                      {article.title}
                    </Link>
                  </h3>
                  {article.description && <p className="mt-2 line-clamp-3 break-words text-sm leading-6 text-muted-foreground sm:text-base sm:leading-7">{article.description}</p>}
                  <div className="flex flex-wrap items-center justify-between gap-3 pt-3">
                    <Link
                      to={`/articles/${article.id}`}
                      className="inline-flex min-h-9 items-center rounded-sm text-sm text-muted-foreground underline-offset-4 transition-colors hover:text-emerald-700 hover:underline focus-visible:outline-2 focus-visible:outline-ring focus-visible:outline-offset-4 dark:hover:text-emerald-400"
                      aria-label={`阅读《${article.title}》`}
                    >
                      阅读全文
                    </Link>
                    <div className="flex gap-1 text-muted-foreground">
                      <Button asChild variant="ghost" size="icon-lg" className="hover:text-emerald-700 dark:hover:text-emerald-400">
                        <Link to={`/articles/${article.id}/edit`} aria-label={`编辑《${article.title}》`} title="编辑文章">
                          <Pencil aria-hidden="true" />
                        </Link>
                      </Button>
                      <DeleteArticleButton
                        iconOnly
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
            <Button variant="outline" size="icon-lg" aria-label="上一页" title="上一页" disabled={page <= 1 || isValidating} onClick={() => setPage(page - 1)}>
              <ChevronLeft aria-hidden="true" />
            </Button>
            <p className="text-xs text-muted-foreground sm:text-sm" aria-live="polite">
              第 {data.page} / {Math.max(1, data.totalPage)} 页 · 共 {data.total} 篇
            </p>
            <Button variant="outline" size="icon-lg" aria-label="下一页" title="下一页" disabled={page >= data.totalPage || isValidating} onClick={() => setPage(page + 1)}>
              <ChevronRight aria-hidden="true" />
            </Button>
          </nav>
        </>
      )}
    </section>
  );
}
