import { ArrowLeft, CalendarDays, FolderOpen, Pencil, RefreshCw } from 'lucide-react';
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
    <main className="mx-auto max-w-3xl px-5 pb-12 pt-6 sm:px-8 sm:pb-16 sm:pt-10">
      <nav aria-label="文章操作" className="mb-8 flex items-center justify-between border-b pb-4 sm:mb-10">
        <div className="flex items-center gap-3">
          <Button asChild variant="ghost" size="icon-lg" className="text-muted-foreground">
            <Link to="/" aria-label="返回文章列表" title="返回文章列表">
              <ArrowLeft aria-hidden="true" />
            </Link>
          </Button>
          <Link to="/" className="font-serif text-sm font-semibold tracking-widest text-muted-foreground hover:text-foreground">随记</Link>
        </div>
        {id && data && !error && (
          <div className="flex items-center gap-1">
            <Button asChild variant="ghost" size="icon-lg" className="text-muted-foreground hover:text-emerald-700 dark:hover:text-emerald-400">
              <Link to={`/articles/${id}/edit`} aria-label={`编辑《${data.title}》`} title="编辑文章">
                <Pencil aria-hidden="true" />
              </Link>
            </Button>
            <DeleteArticleButton iconOnly id={id} title={data.title} onDeleted={() => navigate('/', { replace: true })} />
          </div>
        )}
      </nav>
      {!id && <p role="alert">文章地址无效。</p>}
      {id && isLoading && <p role="status">正在加载文章…</p>}
      {id && error && (
        <div className="space-y-3">
          <p role="alert" className="text-destructive">
            {error.response?.status === 404 ? '文章不存在或已被删除。' : apiErrorMessage(error, '文章加载失败，请重试。')}
          </p>
          <Button variant="outline" disabled={isValidating} onClick={() => void mutate()}>
            <RefreshCw className={isValidating ? 'animate-spin' : ''} aria-hidden="true" />
            重试
          </Button>
        </div>
      )}
      {id && data && !error && (
        <article>
          <header className="mb-7 space-y-4 border-b pb-6 sm:mb-8">
            <h1 className="break-words font-serif text-2xl font-semibold leading-snug tracking-tight sm:text-3xl">{data.title}</h1>
            <div className="flex flex-wrap items-center gap-x-5 gap-y-2 text-xs text-muted-foreground">
              <span className="inline-flex items-center gap-1.5 text-emerald-700 dark:text-emerald-400">
                <FolderOpen className="size-3.5" aria-hidden="true" />
                {data.typeName || '未分类'}
              </span>
              <span className="inline-flex items-center gap-1.5">
                <CalendarDays className="size-3.5" aria-hidden="true" />
                <time dateTime={data.createdAt}>{new Date(data.createdAt).toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric' })}</time>
              </span>
            </div>
          </header>
          <div className="text-[15px] sm:text-base">
            <MarkdownContent content={data.content || '暂无正文。'} />
          </div>
          <footer className="mt-10 flex items-center gap-4 text-xs text-muted-foreground" aria-label="文章结束">
            <span className="h-px flex-1 bg-border" />
            <span>全文完</span>
            <span className="h-px flex-1 bg-border" />
          </footer>
        </article>
      )}
    </main>
  );
}
