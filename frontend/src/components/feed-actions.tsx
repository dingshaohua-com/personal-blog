import { Plus, Trash2 } from 'lucide-react';
import { AlertDialog, Dialog } from 'radix-ui';
import { useRef, useState } from 'react';
import api from '@/api';
import type { FeedDTO } from '@/api/generated/models';
import { Button } from '@/components/ui/button';
import { apiErrorMessage } from '@/utils/article';

const dialogClass = 'fixed left-1/2 top-1/2 z-50 w-[calc(100%-2rem)] max-w-md -translate-x-1/2 -translate-y-1/2 space-y-4 rounded-xl border bg-background p-6 shadow-lg';

export function CreateFeedButton({ onCreated }: { onCreated: (feed: FeedDTO) => void }) {
  const [open, setOpen] = useState(false);
  const [content, setContent] = useState('');
  const [pending, setPending] = useState(false);
  const [error, setError] = useState('');
  const busy = useRef(false);
  const length = Array.from(content).length;

  async function submit() {
    if (busy.current || !content.trim() || length > 100) return;
    busy.current = true;
    setPending(true);
    setError('');
    try {
      const feed = await api.feed.create({ content: content.trim() });
      onCreated(feed);
      setContent('');
      setOpen(false);
    } catch (cause) {
      setError(apiErrorMessage(cause, '发布失败，请重试。'));
    } finally {
      busy.current = false;
      setPending(false);
    }
  }

  return (
    <Dialog.Root open={open} onOpenChange={(next) => { if (!busy.current) { setError(''); setOpen(next); } }}>
      <Dialog.Trigger asChild>
        <Button size="icon-lg" className="bg-emerald-700 text-white hover:bg-emerald-800" aria-label="新增说说" title="新增说说">
          <Plus aria-hidden="true" />
        </Button>
      </Dialog.Trigger>
      <Dialog.Portal>
        <Dialog.Overlay className="fixed inset-0 z-40 bg-black/40" />
        <Dialog.Content className={dialogClass}>
          <Dialog.Title className="text-lg font-semibold">写说说</Dialog.Title>
          <Dialog.Description className="text-sm text-muted-foreground">记录此刻的想法，最多 100 字。发布后可删除，不支持编辑。</Dialog.Description>
          <form className="space-y-4" onSubmit={(event) => { event.preventDefault(); void submit(); }}>
            <label htmlFor="feed-content" className="sr-only">说说内容</label>
            <textarea id="feed-content" value={content} onChange={(event) => setContent(event.target.value)} disabled={pending} required rows={5} aria-describedby="feed-length" aria-invalid={length > 100} placeholder="有什么想记下来的？" className="w-full resize-y rounded-lg border bg-background px-3 py-2 text-sm leading-7 outline-none focus-visible:border-ring focus-visible:ring-2 focus-visible:ring-ring/30 disabled:opacity-50" />
            <p id="feed-length" className={`text-right text-xs ${length > 100 ? 'text-destructive' : 'text-muted-foreground'}`}>{length} / 100</p>
            {error && <p role="alert" className="text-sm text-destructive">{error}</p>}
            <div className="flex justify-end gap-2">
              <Dialog.Close asChild><Button type="button" variant="outline" disabled={pending}>取消</Button></Dialog.Close>
              <Button type="submit" disabled={pending || !content.trim() || length > 100} className="bg-emerald-700 text-white hover:bg-emerald-800">{pending ? '发布中…' : '发布'}</Button>
            </div>
          </form>
        </Dialog.Content>
      </Dialog.Portal>
    </Dialog.Root>
  );
}

export function DeleteFeedButton({ feed, onDeleted }: { feed: FeedDTO; onDeleted: (id: number) => void }) {
  const [open, setOpen] = useState(false);
  const [pending, setPending] = useState(false);
  const [error, setError] = useState('');
  const busy = useRef(false);

  async function remove() {
    if (busy.current) return;
    busy.current = true;
    setPending(true);
    setError('');
    try {
      await api.feed.remove(feed.id);
      setOpen(false);
      onDeleted(feed.id);
    } catch (cause) {
      setError(apiErrorMessage(cause, '删除失败，请重试。'));
    } finally {
      busy.current = false;
      setPending(false);
    }
  }

  return (
    <AlertDialog.Root open={open} onOpenChange={(next) => { if (!busy.current) { setError(''); setOpen(next); } }}>
      <AlertDialog.Trigger asChild>
        <Button variant="ghost" size="icon-lg" className="text-muted-foreground hover:bg-destructive/10 hover:text-destructive" aria-label={`删除说说：${feed.content}`} title="删除说说"><Trash2 aria-hidden="true" /></Button>
      </AlertDialog.Trigger>
      <AlertDialog.Portal>
        <AlertDialog.Overlay className="fixed inset-0 z-40 bg-black/40" />
        <AlertDialog.Content className={dialogClass}>
          <AlertDialog.Title className="text-lg font-semibold">删除说说</AlertDialog.Title>
          <AlertDialog.Description className="text-sm text-muted-foreground">确定删除这条说说吗？删除后无法恢复。</AlertDialog.Description>
          <p className="max-h-40 overflow-y-auto whitespace-pre-wrap break-words rounded-lg bg-muted p-3 text-sm leading-6">{feed.content}</p>
          {error && <p role="alert" className="text-sm text-destructive">{error}</p>}
          <div className="flex justify-end gap-2">
            <AlertDialog.Cancel asChild><Button variant="outline" disabled={pending}>取消</Button></AlertDialog.Cancel>
            <Button variant="destructive" disabled={pending} onClick={() => void remove()}>{pending ? '删除中…' : '确认删除'}</Button>
          </div>
        </AlertDialog.Content>
      </AlertDialog.Portal>
    </AlertDialog.Root>
  );
}
