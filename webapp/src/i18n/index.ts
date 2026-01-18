import { createI18n } from 'vue-i18n';
import zhCN from './locales/zh-CN';
import enUS from './locales/en-US';

export type AppLocale = 'zh-CN' | 'en-US';

export const supportedLocales: AppLocale[] = ['zh-CN', 'en-US'];
export const defaultLocale: AppLocale = 'zh-CN';
export const languageStorageKey = 'nginxpulse_locale';

const messages = {
  'zh-CN': zhCN,
  'en-US': enUS,
};

const numberFormats = {
  'zh-CN': {
    decimal: {
      maximumFractionDigits: 2,
    },
    percent: {
      style: 'percent',
      minimumFractionDigits: 2,
      maximumFractionDigits: 2,
    },
  },
  'en-US': {
    decimal: {
      maximumFractionDigits: 2,
    },
    percent: {
      style: 'percent',
      minimumFractionDigits: 2,
      maximumFractionDigits: 2,
    },
  },
};

export function normalizeLocale(value?: string | null): AppLocale {
  if (!value) {
    return defaultLocale;
  }
  const normalized = value.trim().toLowerCase();
  if (normalized === 'en' || normalized === 'en-us' || normalized === 'en_us') {
    return 'en-US';
  }
  if (normalized === 'zh' || normalized === 'zh-cn' || normalized === 'zh_cn' || normalized === 'zh-hans') {
    return 'zh-CN';
  }
  return defaultLocale;
}

export function getLocaleFromQuery(): AppLocale | null {
  if (typeof window === 'undefined') {
    return null;
  }
  const params = new URLSearchParams(window.location.search);
  const value = params.get('lang') || params.get('locale');
  return value ? normalizeLocale(value) : null;
}

export function getStoredLocale(): AppLocale | null {
  if (typeof window === 'undefined') {
    return null;
  }
  const stored = window.localStorage.getItem(languageStorageKey);
  return stored ? normalizeLocale(stored) : null;
}

export function storeLocale(value: AppLocale) {
  if (typeof window === 'undefined') {
    return;
  }
  window.localStorage.setItem(languageStorageKey, value);
}

const initialLocale = normalizeLocale(getLocaleFromQuery() || getStoredLocale() || defaultLocale);

export const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  locale: initialLocale,
  fallbackLocale: defaultLocale,
  messages,
  numberFormats,
});

export function setLocale(nextLocale: AppLocale, persist = true) {
  const normalized = normalizeLocale(nextLocale);
  i18n.global.locale.value = normalized;
  if (persist) {
    storeLocale(normalized);
  }
  if (typeof document !== 'undefined') {
    document.documentElement.lang = normalized;
    document.title = String(i18n.global.t('app.title'));
  }
}

export function getCurrentLocale(): AppLocale {
  return normalizeLocale(i18n.global.locale.value as string);
}
