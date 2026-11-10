const warn = (action: string, key: string, error: unknown): void => {
    if (process.env.NODE_ENV !== 'production') {
        console.warn(`[storage] не удалось ${action} "${key}"`, error);
    }
};

export const safeGetItem = (key: string): string | null => {
    try {
        return localStorage.getItem(key);
    } catch (error) {
        warn('прочитать', key, error);
        return null;
    }
};

export const safeSetItem = (key: string, value: string): boolean => {
    try {
        localStorage.setItem(key, value);
        return true;
    } catch (error) {
        warn('записать', key, error);
        return false;
    }
};

export const safeRemoveItem = (key: string): void => {
    try {
        localStorage.removeItem(key);
    } catch (error) {
        warn('удалить', key, error);
    }
};
