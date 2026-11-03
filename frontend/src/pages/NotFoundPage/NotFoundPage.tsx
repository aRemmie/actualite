import React from 'react';
import { useNavigate } from 'react-router-dom';
import Button from '../../components/shared/Button/Button';
import { RoutePath } from '../../app/router/paths';
import styles from './NotFoundPage.module.scss';

const NotFoundPage = () => {
    const navigate = useNavigate();

    return (
        <div className={styles.container}>
            <div className={styles.brand}>actualite</div>
            <div className={styles.code}>404</div>
            <h1 className={styles.title}>Страница не найдена</h1>
            <p className={styles.text}>Возможно, она была удалена или никогда не существовала.</p>
            <div>
                <Button onClick={() => navigate(RoutePath.feed, { replace: true })}>
                    На главную
                </Button>
            </div>
        </div>
    );
};

export default NotFoundPage;
