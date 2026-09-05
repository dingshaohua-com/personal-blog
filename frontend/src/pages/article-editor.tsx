import { type FormEvent, useState } from 'react';
import { Link, useNavigate, useParams } from 'react-router';
import { useSWRConfig } from 'swr';
import api from '@/api';
import type { ArticleDetailResponse } from '@/api/generated/models';
import MarkdownContent from '@/components/markdown-content';
import { Button } from '@/components/ui/button';
import { apiErrorMessage, articleIdFromParam, isArticleCacheKey } from '@/utils/article';

const fieldClass = 'w-full rounded-lg border bg-background px-3 py-2 outline-none focus-visible:ring-2 focus-visible:ring-ring disabled:opacity-60';

function ArticleForm({ article }: { article?: ArticleDetailResponse }) {
  const [title, setTitle] = useState(article?.title ?? '');
  const [content, setContent] = useState(article?.content ?? '');
  const [typeId, setTypeId] = useState(article?.typeId ? String(article.typeId) : '');
  const [saving, setSaving] = useState(false);
  const [preview, setPreview] = useState(false);
  const [error, setError] = useState('');
  const navigate = useNavigate();
  const { mutate } = useSWRConfig();
  const types = api.articleType.useList({ swr: { revalidateOnFocus: false, shouldRetryOnError: false } });
  const cancelTo = article ? `/articles/${article.id}` : '/';

  async function submit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    if (saving) return;
    const normalizedTitle = title.trim();
    // Matches the current backend ArticleTitle domain rule.
    if (!normalizedTitle || [...normalizedTitle].length > 200) {
      setError('标题不能为空，且不能超过 200 个字符。');
      return;
    }
    setSaving(true);
    setError('');
    try {
      const body = { title: normalizedTitle, content, ...(typeId ? { typeId: Number(typeId) } : {}) };
      let id = article?.id;
      if (id) await api.article.update(id, body);
      else id = (await api.article.create(body)).id;
      await mutate(isArticleCacheKey, undefined, { revalidate: false });
      navigate(`/articles/${id}`, { replace: true });
    } catch (cause) {
      setError(apiErrorMessage(cause, '保存失败，请重试。'));
      setSaving(false);
    }
  }

  return (
    <form className="space-y-5" onSubmit={submit}>
      <fieldset className="space-y-5" disabled={saving}>
        <div className="space-y-2">
          <label htmlFor="article-title" className="block font-medium">
            标题 <span className="text-destructive">*</span>
          </label>
          <input id="article-title" name="title" required value={title} onChange={(event) => setTitle(event.target.value)} aria-describedby="title-hint" className={fieldClass} />
          <p id="title-hint" className="text-sm text-muted-foreground">
            最多 200 个字符（{[...title.trim()].length}/200）
          </p>
        </div>
        <div className="space-y-2">
          <label htmlFor="article-type" className="block font-medium">
            分类
          </label>
          <select id="article-type" name="typeId" value={typeId} onChange={(event) => setTypeId(event.target.value)} className={fieldClass} disabled={types.isLoading}>
            {!article?.typeId && <option value="">未分类</option>}
            {article?.typeId && !types.data?.some((type) => type.id === article.typeId) && <option value={article.typeId}>{article.typeName || '当前分类'}</option>}
            {types.data?.map((type) => (
              <option key={type.id} value={type.id}>
                {type.name}
              </option>
            ))}
          </select>
          {types.isLoading && (
            <p role="status" className="text-sm text-muted-foreground">
              正在加载分类…
            </p>
          )}
          {types.error && (
            <div className="flex flex-wrap items-center gap-2 text-sm">
              <p role="alert" className="text-destructive">
                分类加载失败，可保留当前分类后保存。
              </p>
              <Button type="button" variant="outline" disabled={types.isValidating} onClick={() => void types.mutate()}>
                重试分类
              </Button>
            </div>
          )}
        </div>
        <div className="space-y-2">
          <label htmlFor="article-content" className="block font-medium">
            正文
          </label>
          <div className="flex gap-2" role="group" aria-label="正文模式">
            <Button type="button" size="sm" variant={preview ? 'outline' : 'secondary'} aria-pressed={!preview} onClick={() => setPreview(false)}>
              编辑源码
            </Button>
            <Button type="button" size="sm" variant={preview ? 'secondary' : 'outline'} aria-pressed={preview} onClick={() => setPreview(true)}>
              预览
            </Button>
          </div>
          <textarea id="article-content" name="content" rows={16} hidden={preview} value={content} onChange={(event) => setContent(event.target.value)} className={`${fieldClass} resize-y font-mono leading-7`} aria-describedby="content-hint" placeholder="支持 Markdown，例如 ## 标题、**加粗**、列表和代码块…" />
          {preview && (
            <section aria-label="Markdown 预览" className="min-h-64 rounded-lg border bg-background p-4 sm:p-6">
              <MarkdownContent content={content || '暂无内容，请先在编辑源码中输入正文。'} />
            </section>
          )}
          <p id="content-hint" className="text-sm text-muted-foreground">
            支持 Markdown：标题、列表、引用、代码块、表格、链接和图片。
          </p>
        </div>
      </fieldset>
      {error && (
        <p role="alert" className="text-sm text-destructive">
          {error}
        </p>
      )}
      <div className="flex gap-2">
        <Button type="submit" disabled={saving}>
          {saving ? '保存中…' : article ? '保存修改' : '创建文章'}
        </Button>
        <Button type="button" variant="outline" disabled={saving} onClick={() => navigate(cancelTo)}>
          取消
        </Button>
      </div>
    </form>
  );
}

export default function ArticleEditor() {
  const { id: param } = useParams();
  const editing = param !== undefined;
  const id = articleIdFromParam(param);
  const { data, error, isLoading, isValidating, mutate } = api.article.useGet(id ?? 0, { swr: { enabled: editing && id !== undefined, revalidateOnFocus: false, shouldRetryOnError: false } });

  return (
    <main className="mx-auto max-w-3xl space-y-6 p-4 py-8">
      <Button asChild variant="outline">
        <Link to={id ? `/articles/${id}` : '/'}>返回{editing && id ? '文章详情' : '文章列表'}</Link>
      </Button>
      <h1 className="text-2xl font-semibold">{editing ? '编辑文章' : '新增文章'}</h1>
      {editing && !id && <p role="alert">文章地址无效。</p>}
      {editing && id && isLoading && <p role="status">正在加载文章…</p>}
      {editing && id && error && (
        <div className="space-y-3">
          <p role="alert" className="text-destructive">
            {error.response?.status === 404 ? '文章不存在或已被删除。' : apiErrorMessage(error, '文章加载失败，请重试。')}
          </p>
          <Button variant="outline" disabled={isValidating} onClick={() => void mutate()}>
            重试
          </Button>
        </div>
      )}
      {!editing && <ArticleForm key="new" />}
      {editing && id && data && !error && <ArticleForm key={id} article={data} />}
    </main>
  );
}
