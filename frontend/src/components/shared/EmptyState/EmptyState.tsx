import React, { FC, ReactNode } from 'react';
import styles from './EmptyState.module.scss';

interface EmptyStateProps {
    title: string;
    children?: ReactNode;
}

const EmptyState: FC<EmptyStateProps> = ({ title, children }) => (
    <div className={styles.empty}>
        <h2 className={styles.title}>{title}</h2>
        {children}
    </div>
);

export default EmptyState;
