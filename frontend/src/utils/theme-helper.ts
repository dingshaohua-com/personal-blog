export type ColorThemeId = 'default' | 'green' | 'pink';

const STORAGE_KEY = 'app-color-theme';

export const COLOR_THEME_OPTIONS: { id: ColorThemeId; label: string }[] = [
  { id: 'default', label: '默认' },
  { id: 'green', label: '绿色' },
  { id: 'pink', label: '粉色' },
];

export function readStoredColorTheme(): ColorThemeId {
  const raw = localStorage.getItem(STORAGE_KEY);
  if (raw === 'green' || raw === 'pink' || raw === 'default') return raw;
  return 'default';
}

/** 只更新 DOM，不写 storage（用于首屏同步） */
export function applyColorTheme(theme: ColorThemeId): void {
  const root = document.documentElement;
  if (theme === 'default') root.removeAttribute('data-theme');
  else root.setAttribute('data-theme', theme);
}

/** 切换主题并持久化 */
export function setColorTheme(theme: ColorThemeId): void {
  applyColorTheme(theme);
  localStorage.setItem(STORAGE_KEY, theme);
}
