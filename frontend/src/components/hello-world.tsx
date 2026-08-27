import { useState } from 'react';
import viteLogo from '@/assets/imgs/vite.svg';
import {
  COLOR_THEME_OPTIONS,
  readStoredColorTheme,
  setColorTheme,
  type ColorThemeId,
} from '@/utils/theme-helper';
import { Button } from '@/components/ui/button';

export default function HelloWorld() {
  const [colorTheme, setColorThemeState] = useState<ColorThemeId>(readStoredColorTheme);

  function handleThemeChange(next: ColorThemeId) {
    setColorTheme(next);
    setColorThemeState(next);
  }

  return (
    <div className="bg-background text-foreground min-h-[50vh] p-4">
      <p className="mb-3">你好世界！</p>
      <Button>Click me</Button>
      <div className="mb-4 flex flex-wrap gap-2">
        <span className="text-muted-foreground self-center text-sm">主题：</span>
        {COLOR_THEME_OPTIONS.map(({ id, label }) => (
          <button
            key={id}
            type="button"
            onClick={() => handleThemeChange(id)}
            className={`rounded-md border px-3 py-1 text-sm transition-colors ${
              colorTheme === id
                ? 'border-primary bg-primary text-primary-foreground'
                : 'border-border bg-card text-card-foreground hover:bg-accent'
            }`}
          >
            {label}
          </button>
        ))}
      </div>
      <div className="bg-primary inline-flex rounded-lg p-4 ring-2 ring-ring/40">
        <img src={viteLogo} className="vite block" alt="Vite logo" />
      </div>
      <div className="shadow-soft mt-4">
        <input type="text" className="pure-ipt" />
      </div>
    </div>
  );
}
