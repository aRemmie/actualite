import axios, { AxiosError } from 'axios';
import { toast } from 'react-toastify';
import { clearToken, getToken } from '@utils/token.utils';

export const API_URL: string = __API_URL__;

const PUBLIC_ENDPOINTS = ['/login', '/register'];

let isSessionExpiredHandled = false;

let onUnauthorized: (() => void) | null = null;

export const setUnauthorizedHandler = (handler: () => void): void => {
    onUnauthorized = handler;
};

export const resetSessionExpiredFlag = (): void => {
    isSessionExpiredHandled = false;
};

const $api = axios.create({
    baseURL: API_URL,
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json',
    },
    withCredentials: false,
});

$api.interceptors.request.use((config) => {
    const token = getToken();
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

$api.interceptors.response.use(
    (response) => response,
    (error: AxiosError) => {
        if (error.code === 'ERR_NETWORK' || error.message === 'Network Error') {
            toast.error('Не удаётся подключиться к серверу', {
                autoClose: 8000,
                toastId: 'network-error',
            });
        }

        const url = error.config?.url ?? '';
        const isPublic = PUBLIC_ENDPOINTS.some((endpoint) => url.endsWith(endpoint));

        if (error.response?.status === 401 && !isPublic && !isSessionExpiredHandled) {
            isSessionExpiredHandled = true;
            clearToken();
            onUnauthorized?.();
            toast.error('Сессия истекла. Войдите снова', { toastId: 'session-expired' });
        }
        return Promise.reject(error);
    },
);

export default $api;
