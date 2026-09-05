import { MessageCircle, RefreshCw } from 'lucide-react';
import { useState } from 'react';
import api from '@/api';
import { CreateFeedButton, DeleteFeedButton } from '@/components/feed-actions';
import MarkdownContent from '@/components/markdown-content';
import { Button } from '@/components/ui/button';
import { apiErrorMessage } from '@/utils/article';

export default function FeedList() {
  const [visibleCount, setVisibleCount] = useState(10);
  const { data, error, isLoading, isValidating, mutate } = api.feed.useList({ swr: { revalidateOnFocus: false, shouldRetryOnError: false } });
  const feeds = [...(data ?? [])].sort((a, b) => Date.parse(b.createdAt) - Date.parse(a.createdAt) || b.id - a.id);

  return (
    <section className="space-y-6" aria-labelledby="feeds-title" aria-busy={isValidating}>
      <div className="flex items-center justify-between gap-4">
        <div>
          <h2 id="feeds-title" className="text-lg font-semibold">
            最近说说
          </h2>
          <p className="mt-1 text-xs text-muted-foreground">{data && !error ? `共 ${feeds.length} 条 · ` : ''}简短的想法，随时的分享</p>
        </div>
        <div className="flex gap-2">
        <Button variant="ghost" size="icon-lg" className="text-muted-foreground" disabled={isValidating} onClick={() => void mutate()} aria-label="刷新说说" title="刷新说说">
          <RefreshCw className={isValidating ? 'animate-spin' : ''} aria-hidden="true" />
        </Button>
        <CreateFeedButton onCreated={(feed) => { void mutate((current) => [feed, ...(current ?? []).filter((item) => item.id !== feed.id)]); }} />
        </div>
      </div>
      {isLoading && (
        <p role="status" className="rounded-xl border bg-background p-8 text-center text-muted-foreground">
          正在加载说说…
        </p>
      )}
      {error && (
        <p role="alert" className="rounded-xl border border-destructive/20 bg-background p-6 text-destructive">
          {apiErrorMessage(error, '说说加载失败，请点击刷新重试。')}
        </p>
      )}
      {!isLoading &&
        !error &&
        (feeds.length ? (
          <>
            <ol className="divide-y border-t">
              {feeds.slice(0, visibleCount).map((feed) => (
                <li key={feed.id} className="flex gap-3 py-5 sm:gap-5 sm:py-6">
                  <span className="mt-0.5 flex size-9 shrink-0 items-center justify-center rounded-full bg-emerald-700/8 text-emerald-700 dark:text-emerald-300">
                    <MessageCircle className="size-4" aria-hidden="true" />
                  </span>
                  <article className="min-w-0 flex-1">
                    <div className="mb-3 flex items-center justify-between gap-3">
                    <time dateTime={feed.createdAt} className="text-xs text-muted-foreground">
                      {new Date(feed.createdAt).toLocaleString('zh-CN', { year: 'numeric', month: '2-digit', day: '2-digit', hour: '2-digit', minute: '2-digit', hour12: false })}
                    </time>
                    <DeleteFeedButton feed={feed} onDeleted={(id) => { void mutate((current) => (current ?? []).filter((item) => item.id !== id)); }} />
                    </div>
                    <div className="text-sm leading-7 sm:text-base">
                      <MarkdownContent content={feed.content} />
                    </div>
                  </article>
                </li>
              ))}
            </ol>
            <div className="pt-2 text-center">
              {visibleCount < feeds.length ? (
                <Button variant="outline" onClick={() => setVisibleCount((count) => count + 10)}>
                  加载更多（还有 {feeds.length - visibleCount} 条）
                </Button>
              ) : (
                <p className="text-xs text-muted-foreground">已展示全部 {feeds.length} 条说说</p>
              )}
            </div>
          </>
        ) : (
          <div className="rounded-2xl border border-dashed bg-background p-12 text-center">
            <MessageCircle className="mx-auto mb-4 size-8 text-muted-foreground" aria-hidden="true" />
            <p className="font-medium">还没有说说</p>
            <p className="mt-2 text-sm text-muted-foreground">新的想法与分享会出现在这里。</p>
          </div>
        ))}
    </section>
  );
}
