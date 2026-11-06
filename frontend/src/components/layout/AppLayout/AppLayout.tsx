import React, { FC, ReactNode } from 'react';
import { NavLink, Link, useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import { useAuthStore } from '@stores/useAuthStore';
import { RoutePath } from '@router/paths';
import styles from './AppLayout.module.scss';

interface AppLayoutProps {
    title: string;
    showBack?: boolean;
    children: ReactNode;
}

const navClass = ({ isActive }: { isActive: boolean }) =>
    `${styles.navItem} ${isActive ? styles.active : ''}`;

const AppLayout: FC<AppLayoutProps> = ({ title, showBack = false, children }) => {
    const navigate = useNavigate();
    const logout = useAuthStore((state) => state.logout);

    const handleLogout = () => {
        logout();
        toast.info('Вы вышли из аккаунта');
    };

    return (
        <div className={styles.layout}>
            <aside className={styles.sidebar}>
                <Link to={RoutePath.feed} className={styles.brand}>
                    actualite
                </Link>
                <nav className={styles.nav}>
                    <NavLink to={RoutePath.feed} end className={navClass}>
                        <span className={styles.navIcon}>⌂</span>
                        <span>Главная</span>
                    </NavLink>
                    <span
                        className={`${styles.navItem} ${styles.disabled}`}
                        aria-disabled="true"
                        title="Скоро"
                    >
                        <span className={styles.navIcon}>⌕</span>
                        <span>Поиск</span>
                    </span>
                    <span
                        className={`${styles.navItem} ${styles.disabled}`}
                        aria-disabled="true"
                        title="Скоро"
                    >
                        <span className={styles.navIcon}>🕭</span>
                        <span>Уведомления</span>
                    </span>
                    <NavLink to={RoutePath.profile} end className={navClass}>
                        <span className={styles.navIcon}>♙</span>
                        <span>Мой профиль</span>
                    </NavLink>
                </nav>
                <div className={styles.sidebarBottom}>
                    <button type="button" className={styles.navItem} onClick={handleLogout}>
                        <span className={styles.navIcon}>⎋</span>
                        <span>Выйти</span>
                    </button>
                </div>
            </aside>

            <main className={styles.main}>
                <div className={styles.header}>
                    {showBack && (
                        <button
                            type="button"
                            className={styles.back}
                            onClick={() => navigate(-1)}
                            aria-label="Назад"
                        >
                            ←
                        </button>
                    )}
                    <h1 className={styles.title}>{title}</h1>
                </div>
                {children}
            </main>

            <aside className={styles.aside}>
                <h2 className={styles.asideTitle}>Актуальное</h2>
                <p className={styles.asideText}>Здесь появятся популярные темы и хэштеги.</p>
            </aside>
        </div>
    );
};

export default AppLayout;
