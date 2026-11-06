import { create } from 'zustand';
import { safeGetItem, safeSetItem } from '@utils/storage.utils';

export type Theme = 'dark' | 'light';

const THEME_KEY = 'theme';
const DEFAULT_THEME: Theme = 'dark';

const isTheme = (value: unknown): value is Theme => value === 'dark' || value === 'light';

const readInitialTheme = (): Theme => {
    const saved = safeGetItem(THEME_KEY);
    if (isTheme(saved)) {
        return saved;
    }
    const fromDom = document.documentElement.getAttribute('data-theme');
    if (isTheme(fromDom)) {
        return fromDom;
    }
    return DEFAULT_THEME;
};

const applyTheme = (theme: Theme): void => {
    document.documentElement.setAttribute('data-theme', theme);
    safeSetItem(THEME_KEY, theme);
};

interface ThemeState {
    theme: Theme;
    setTheme: (theme: Theme) => void;
    toggleTheme: () => void;
}

export const useThemeStore = create<ThemeState>()((set, get) => {
    const initial = readInitialTheme();
    document.documentElement.setAttribute('data-theme', initial);

    return {
        theme: initial,
        setTheme: (theme) => {
            applyTheme(theme);
            set({ theme });
        },
        toggleTheme: () => get().setTheme(get().theme === 'dark' ? 'light' : 'dark'),
    };
});
