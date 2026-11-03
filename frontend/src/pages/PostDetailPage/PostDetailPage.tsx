import React from 'react';
import { useParams } from 'react-router-dom';
import AppLayout from '../../components/layout/AppLayout/AppLayout';
import EmptyState from '../../components/shared/EmptyState/EmptyState';

const PostDetailPage = () => {
    const { id } = useParams<{ id: string }>();

    return (
        <AppLayout title="Пост" showBack>
            <EmptyState title="Страница поста">
                <p>Пост {id} будет отображаться здесь.</p>
            </EmptyState>
        </AppLayout>
    );
};

export default PostDetailPage;
