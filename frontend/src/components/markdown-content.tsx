import ReactMarkdown, { type Components } from 'react-markdown';
import remarkGfm from 'remark-gfm';
import '@/assets/styles/markdown.css';

const components: Components = {
  a: ({ children, href, title }) => {
    const external = /^(https?:)?\/\//i.test(href ?? '');
    return (
      <a href={href} title={title} target={external ? '_blank' : undefined} rel={external ? 'noopener noreferrer' : undefined}>
        {children}
      </a>
    );
  },
  img: ({ src, alt, title }) => (src ? <img src={src} alt={alt ?? ''} title={title} loading="lazy" referrerPolicy="no-referrer" /> : <span>{alt}</span>),
  table: ({ children }) => (
    <div className="markdown-table-scroll" role="region" aria-label="表格" tabIndex={0}>
      <table>{children}</table>
    </div>
  ),
};

export default function MarkdownContent({ content }: { content: string }) {
  return (
    <div className="markdown-content">
      <ReactMarkdown remarkPlugins={[remarkGfm]} components={components}>
        {content}
      </ReactMarkdown>
    </div>
  );
}
