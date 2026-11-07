import { safeGetItem, safeRemoveItem, safeSetItem } from './storage.utils';
import { isTokenExpired } from './string.utils';

export const TOKEN_KEY = 'token';

export const getToken = (): string | null => safeGetItem(TOKEN_KEY);

export const setToken = (token: string): void => {
    if (!safeSetItem(TOKEN_KEY, token)) {
        throw new Error('Не удалось сохранить сессию в localStorage. Возможно, включён приватный режим или заблокированы cookies.');
    }
};

export const clearToken = (): void => safeRemoveItem(TOKEN_KEY);

export const hasValidToken = (): boolean => {
    const token = getToken();
    return Boolean(token) && !isTokenExpired(token as string);
};
