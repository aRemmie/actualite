import { isAxiosError } from 'axios';

export interface ApiErrorInfo {
    status?: number;
    code?: string;
    message: string;
}

interface ApiErrorBody {
    code?: string;
    message?: string;
}

export const parseApiError = (error: unknown): ApiErrorInfo => {
    if (isAxiosError(error)) {
        const body = error.response?.data as ApiErrorBody | undefined;
        return {
            status: error.response?.status,
            code: body?.code,
            message: body?.message || error.message,
        };
    }
    return { message: error instanceof Error ? error.message : '' };
};
