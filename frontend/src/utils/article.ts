import { isAxiosError } from 'axios';
import type { ErrorModel } from '@/api/generated/models';

export function articleIdFromParam(value: string | undefined): number | undefined {
  if (!value || !/^[1-9]\d*$/.test(value)) return undefined;
  const id = Number(value);
  return Number.isSafeInteger(id) ? id : undefined;
}

export function apiErrorMessage(error: unknown, fallback: string): string {
  if (!isAxiosError<ErrorModel>(error)) return fallback;
  const problem = error.response?.data;
  return (
    problem?.errors
      ?.map((item) => item.message)
      .filter(Boolean)
      .join('；') ||
    problem?.detail ||
    fallback
  );
}

// Clear article caches after writes so returning to another page also loads fresh data.
export function isArticleCacheKey(key: unknown): boolean {
  return Array.isArray(key) && typeof key[0] === 'string' && (key[0] === '/article' || /^\/article\/\d+$/.test(key[0]));
}
