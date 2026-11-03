import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import { AuthService } from '../services/AuthService';
import { clearToken, hasValidToken, setToken } from '../utils/token.utils';

interface AuthState {
    isAuth: boolean;
    userId: string | null;
    email: string | null;
    username: string | null;
    isLoading: boolean;

    login: (email: string, password: string) => Promise<void>;
    register: (email: string, username: string, password: string) => Promise<void>;
    logout: () => void;
}

type AuthData = Pick<AuthState, 'isAuth' | 'userId' | 'email' | 'username' | 'isLoading'>;

const initialState: AuthData = {
    isAuth: false,
    userId: null,
    email: null,
    username: null,
    isLoading: false,
};

export const useAuthStore = create<AuthState>()(
    persist(
        (set, get) => ({
            ...initialState,

            async login(email: string, password: string) {
                set({ isLoading: true });
                try {
                    const response = await AuthService.login({ email, password });
                    if (!response.token) {
                        throw new Error('Сервер не вернул токен');
                    }
                    setToken(response.token);
                    set({
                        isAuth: true,
                        isLoading: false,
                        userId: response.user_id ?? null,
                        username: response.username ?? null,
                        email,
                    });
                } catch (e) {
                    set({ isLoading: false });
                    throw e;
                }
            },

            async register(email: string, username: string, password: string) {
                set({ isLoading: true });
                try {
                    await AuthService.register({ email, username, password, role: 'user' });
                } catch (e) {
                    set({ isLoading: false });
                    throw e;
                }
                await get().login(email, password);
            },

            logout() {
                AuthService.logout();
                set({ ...initialState });
            },
        }),
        {
            name: 'auth-storage',
            partialize: (state) => ({
                userId: state.userId,
                email: state.email,
                username: state.username,
            }),
            merge: (persisted, current) => {
                const valid = hasValidToken();
                if (!valid) {
                    clearToken();
                    return { ...current, ...initialState };
                }
                return { ...current, ...(persisted as Partial<AuthState>), isAuth: true };
            },
        },
    ),
);
