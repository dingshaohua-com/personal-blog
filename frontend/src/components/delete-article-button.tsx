import { Trash2 } from 'lucide-react';
import { AlertDialog } from 'radix-ui';
import { useState } from 'react';
import { useSWRConfig } from 'swr';
import api from '@/api';
import { Button } from '@/components/ui/button';
import { apiErrorMessage, isArticleCacheKey } from '@/utils/article';

export default function DeleteArticleButton({ id, title, onDeleted, iconOnly = false }: { id: number; title: string; onDeleted: () => void; iconOnly?: boolean }) {
  const [open, setOpen] = useState(false);
  const [deleting, setDeleting] = useState(false);
  const [error, setError] = useState('');
  const { mutate } = useSWRConfig();

  async function remove() {
    if (deleting) return;
    setDeleting(true);
    setError('');
    try {
      await api.article.remove(id);
      await mutate(isArticleCacheKey, undefined, { revalidate: false });
      setOpen(false);
      onDeleted();
    } catch (cause) {
      setError(apiErrorMessage(cause, '删除失败，请重试。'));
    } finally {
      setDeleting(false);
    }
  }

  return (
    <AlertDialog.Root
      open={open}
      onOpenChange={(next) => {
        if (!deleting) {
          setError('');
          setOpen(next);
        }
      }}
    >
      <AlertDialog.Trigger asChild>
        <Button
          variant={iconOnly ? 'ghost' : 'destructive'}
          size={iconOnly ? 'icon-lg' : 'default'}
          className={iconOnly ? 'text-muted-foreground hover:bg-destructive/10 hover:text-destructive' : undefined}
          aria-label={`删除《${title}》`}
          title="删除文章"
        >
          {iconOnly ? <Trash2 aria-hidden="true" /> : '删除'}
        </Button>
      </AlertDialog.Trigger>
      <AlertDialog.Portal>
        <AlertDialog.Overlay className="fixed inset-0 z-40 bg-black/40" />
        <AlertDialog.Content className="fixed left-1/2 top-1/2 z-50 w-[calc(100%-2rem)] max-w-md -translate-x-1/2 -translate-y-1/2 space-y-4 rounded-xl border bg-background p-6 shadow-lg">
          <AlertDialog.Title className="text-lg font-semibold">删除文章</AlertDialog.Title>
          <AlertDialog.Description className="break-words text-muted-foreground">确定删除《{title}》吗？删除后无法恢复。</AlertDialog.Description>
          {error && (
            <p role="alert" className="text-sm text-destructive">
              {error}
            </p>
          )}
          <div className="flex justify-end gap-2">
            <AlertDialog.Cancel asChild>
              <Button variant="outline" disabled={deleting}>
                取消
              </Button>
            </AlertDialog.Cancel>
            <Button variant="destructive" disabled={deleting} onClick={() => void remove()}>
              {deleting ? '删除中…' : '确认删除'}
            </Button>
          </div>
        </AlertDialog.Content>
      </AlertDialog.Portal>
    </AlertDialog.Root>
  );
}
