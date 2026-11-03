import React, { useState } from 'react';
import Button from '../../components/shared/Button/Button';
import AuthModal from '../../components/auth/AuthModal';
import styles from './AuthenticationPage.module.scss';

const AuthenticationPage = () => {
    const [openAuthModal, setOpenAuthModal] = useState(false);
    const [isSignup, setIsSignup] = useState(true);

    const handleOpenAuthModal = (signup: boolean) => {
        setIsSignup(signup);
        setOpenAuthModal(true);
    };

    const handleCloseAuthModal = () => setOpenAuthModal(false);

    return (
        <div className={styles.container}>
            <section className={styles.copy}>
                <div className={styles.brand}>actualite</div>
                <h1 className={styles.title}>
                    Мысли
                    <br />
                    находят
                    <br />
                    свою аудиторию
                </h1>
                <p className={styles.subtitle}>Блог-платформа для тех, кто ценит мысли.</p>
                <div className={styles.actions}>
                    <Button onClick={() => handleOpenAuthModal(true)}>Начать</Button>
                    <Button variant="outlined" onClick={() => handleOpenAuthModal(false)}>
                        Войти
                    </Button>
                </div>
            </section>

            <AuthModal
                open={openAuthModal}
                handleClose={handleCloseAuthModal}
                initialMode={isSignup ? 'signup' : 'signin'}
            />
        </div>
    );
};

export default AuthenticationPage;
