import React from 'react';
import { useParams } from 'react-router-dom';
import AppLayout from '@components/layout/AppLayout/AppLayout';
import EmptyState from '@components/shared/EmptyState/EmptyState';
import { useAuthStore } from '@stores/useAuthStore';

const ProfilePage = () => {
    const { id } = useParams<{ id: string }>();
    const { userId, username, email } = useAuthStore();
    const isOwn = !id || id === userId;

    return (
        <AppLayout title="Профиль" showBack={Boolean(id)}>
            {isOwn ? (
                <EmptyState title={username ? `@${username}` : 'Мой профиль'}>
                    {email && <p>{email}</p>}
                </EmptyState>
            ) : (
                <EmptyState title="Профиль пользователя">
                    <p>Профиль {id} будет отображаться здесь.</p>
                </EmptyState>
            )}
        </AppLayout>
    );
};

export default ProfilePage;
