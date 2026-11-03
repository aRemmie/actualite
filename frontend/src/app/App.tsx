import React, { useEffect, useState } from 'react';
import { ToastContainer } from 'react-toastify';
import AppRouter from './router/AppRouter';
import ThemeToggle from '../components/shared/ThemeToggle/ThemeToggle';
import './styles/global.scss';

type Theme = 'dark' | 'light';

const App = () => {
    const [theme, setTheme] = useState<Theme>('dark');

    useEffect(() => {
        const savedTheme = localStorage.getItem('theme') as Theme | null;
        const initialTheme: Theme = savedTheme === 'light' ? 'light' : 'dark';
        document.documentElement.setAttribute('data-theme', initialTheme);
        setTheme(initialTheme);

        const observer = new MutationObserver(() => {
            const current = document.documentElement.getAttribute('data-theme');
            if (current === 'dark' || current === 'light') setTheme(current);
        });
        observer.observe(document.documentElement, {
            attributes: true,
            attributeFilter: ['data-theme'],
        });
        return () => observer.disconnect();
    }, []);

    return (
        <div className="app">
            <ToastContainer
                position="bottom-right"
                hideProgressBar={true}
                theme={theme}
                autoClose={3000}
                closeOnClick
                pauseOnHover
                draggable={false}
                toastClassName="actualite-toast"
            />
            <AppRouter />
            <ThemeToggle />
        </div>
    );
};

export default App;
