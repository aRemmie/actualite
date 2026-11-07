import React, { useState, useEffect } from 'react';
import Modal from '@components/shared/Modal/Modal';
import Button from '@components/shared/Button/Button';
import SigninForm from './SigninForm';
import SignupForm from './SignupForm';
import styles from './AuthModal.module.scss';

interface AuthModalProps {
    open: boolean;
    handleClose: () => void;
    initialMode?: 'signin' | 'signup';
}

const AuthModal: React.FC<AuthModalProps> = ({ open, handleClose, initialMode = 'signup' }) => {
    const [isSignup, setIsSignup] = useState(initialMode === 'signup');

    useEffect(() => {
        if (open) {
            setIsSignup(initialMode === 'signup');
        }
    }, [open, initialMode]);

    return (
        <Modal open={open} onClose={handleClose}>
            <div className={styles.modalContent}>
                {isSignup ? <SignupForm /> : <SigninForm />}
                <div className={styles.switchSection}>
                    <p className={styles.switchText}>
                        {isSignup ? 'Уже есть аккаунт?' : 'Нет аккаунта?'}
                    </p>
                    <Button variant="outlined" fullWidth onClick={() => setIsSignup(!isSignup)}>
                        {isSignup ? 'Войти' : 'Зарегистрироваться'}
                    </Button>
                </div>
            </div>
        </Modal>
    );
};

export default AuthModal;
