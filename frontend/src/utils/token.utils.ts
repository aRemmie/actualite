import { isTokenExpired } from './string.utils';

export const TOKEN_KEY = 'token';

export const getToken = (): string | null => {
    try {
        return localStorage.getItem(TOKEN_KEY);
    } catch {
        return null;
    }
};

export const setToken = (token: string): void => {
    localStorage.setItem(TOKEN_KEY, token);
};

export const clearToken = (): void => {
    try {
        localStorage.removeItem(TOKEN_KEY);
    } catch {
    }
};

export const hasValidToken = (): boolean => {
    const token = getToken();
    return Boolean(token) && !isTokenExpired(token as string);
};
