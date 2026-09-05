import { BookOpen, MessageCircle, NotebookPen } from 'lucide-react';
import { Link, useSearchParams } from 'react-router';
import ArticleList from '@/components/article-list';
import FeedList from '@/components/feed-list';

export default function Home() {
  const [searchParams] = useSearchParams();
  const showFeeds = searchParams.get('view') === 'feeds';
  const articleParams = new URLSearchParams(searchParams);
  articleParams.delete('view');
  const feedParams = new URLSearchParams(searchParams);
  feedParams.set('view', 'feeds');

  return (
    <div className="min-h-screen bg-background">
      <header className="border-b bg-background">
        <div className="mx-auto flex max-w-4xl items-center gap-3 px-5 py-5 sm:px-8">
          <span className="flex size-10 items-center justify-center rounded-xl bg-emerald-700 text-white">
            <NotebookPen className="size-5" aria-hidden="true" />
          </span>
          <Link to="/" className="text-lg font-semibold tracking-wide hover:text-emerald-700">
            随记
          </Link>
          <span className="ml-auto text-xs text-muted-foreground sm:text-sm">记录 · 阅读 · 分享</span>
        </div>
      </header>
      <main className="mx-auto max-w-4xl px-5 py-10 sm:px-8 sm:py-16">
        <div className="mb-10 space-y-4 sm:mb-12">
          <p className="text-xs font-medium tracking-widest text-emerald-700 dark:text-emerald-400">文字与片刻</p>
          <h1 className="font-serif text-3xl font-semibold tracking-tight sm:text-4xl">把日子写成文字。</h1>
          <p className="max-w-xl text-sm leading-7 text-muted-foreground sm:text-base">关于技术、生活，以及值得记下的想法。</p>
        </div>
        <nav aria-label="内容分类" className="mb-8 flex gap-1 border-b">
          <Link
            to={{ pathname: '/', search: articleParams.toString() }}
            aria-current={!showFeeds ? 'page' : undefined}
            className={`inline-flex items-center gap-2 border-b-2 px-5 py-3 text-sm font-medium transition-colors ${!showFeeds ? 'border-emerald-700 text-emerald-700 dark:text-emerald-400' : 'border-transparent text-muted-foreground hover:text-foreground'}`}
          >
            <BookOpen className="size-4" aria-hidden="true" />
            文章
          </Link>
          <Link
            to={{ pathname: '/', search: feedParams.toString() }}
            aria-current={showFeeds ? 'page' : undefined}
            className={`inline-flex items-center gap-2 border-b-2 px-5 py-3 text-sm font-medium transition-colors ${showFeeds ? 'border-emerald-700 text-emerald-700 dark:text-emerald-400' : 'border-transparent text-muted-foreground hover:text-foreground'}`}
          >
            <MessageCircle className="size-4" aria-hidden="true" />
            说说
          </Link>
        </nav>
        {showFeeds ? <FeedList /> : <ArticleList />}
      </main>
      <footer className="mx-auto max-w-4xl px-5 pb-8 pt-4 text-xs text-muted-foreground sm:px-8">
        <div className="border-t pt-6">随记 · 记录生活，保持好奇。</div>
      </footer>
    </div>
  );
}
