import React, { FC } from 'react';
import styles from './TextField.module.scss';

interface TextFieldProps {
    label?: string;
    error?: string;
    touched?: boolean;
    textarea?: boolean;
    rows?: number;
    name?: string;
    type?: string;
    placeholder?: string;
    value?: string | number;
    onChange?: (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => void;
    onBlur?: (e: React.FocusEvent<HTMLInputElement | HTMLTextAreaElement>) => void;
    disabled?: boolean;
    required?: boolean;
    autoFocus?: boolean;
    maxLength?: number;
    minLength?: number;
    pattern?: string;
    readOnly?: boolean;
    autoComplete?: string;
    className?: string;
}

const TextField: FC<TextFieldProps> = ({
    label,
    error,
    touched,
    textarea = false,
    rows = 3,
    className = '',
    ...props
}) => {
    const hasError = Boolean(touched && error);
    const inputClassName = `${styles.input} ${textarea ? styles.textarea : ''} ${hasError ? styles.error : ''}`;
    const ariaProps = { 'aria-invalid': hasError ? true : undefined };

    return (
        <div className={`${styles.container} ${className}`}>
            {label && <label className={styles.label}>{label}</label>}
            {textarea ? (
                <textarea className={inputClassName} rows={rows} {...ariaProps} {...(props as any)} />
            ) : (
                <input className={inputClassName} {...ariaProps} {...(props as any)} />
            )}
            {hasError && <span className={styles.errorMessage}>{error}</span>}
        </div>
    );
};

export default TextField;
