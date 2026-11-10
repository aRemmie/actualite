import React from 'react';
import { useThemeStore } from '@stores/useThemeStore';
import styles from './ThemeToggle.module.scss';

const ThemeToggle: React.FC = () => {
    const theme = useThemeStore((state) => state.theme);
    const toggleTheme = useThemeStore((state) => state.toggleTheme);

    return (
        <button className={styles.toggle} onClick={toggleTheme} aria-label="Переключить тему">
            <div className={styles.switch}>
                <div className={`${styles.slider} ${theme === 'light' ? styles.active : ''}`}></div>
            </div>
        </button>
    );
};

export default ThemeToggle;
