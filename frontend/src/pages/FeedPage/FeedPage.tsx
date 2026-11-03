import React from 'react';
import AppLayout from '../../components/layout/AppLayout/AppLayout';
import EmptyState from '../../components/shared/EmptyState/EmptyState';
import { useAuthStore } from '../../stores/useAuthStore';

const FeedPage = () => {
    const username = useAuthStore((state) => state.username);

    return (
        <AppLayout title="Главная">
            <EmptyState title={username ? `Здравствуйте, @${username}` : 'Здравствуйте'}>
                <p>посты пока не реализованы.</p>
            </EmptyState>
        </AppLayout>
    );
};

export default FeedPage;
