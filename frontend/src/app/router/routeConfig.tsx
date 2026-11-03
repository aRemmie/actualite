import { ReactNode } from 'react';
import AuthenticationPageAsync from '@pages/AuthenticationPage/AuthenticationPage.async';
import FeedPageAsync from '@pages/FeedPage/FeedPage.async';
import PostDetailPageAsync from '@pages/PostDetailPage/PostDetailPage.async';
import ProfilePageAsync from '@pages/ProfilePage/ProfilePage.async';
import { AppRoutes, RoutePath } from './paths';

export { AppRoutes, RoutePath };

export interface AppRouteConfig {
    path: string;
    element: ReactNode;
    authOnly: boolean;
}

export const routeConfig: Record<AppRoutes, AppRouteConfig> = {
    [AppRoutes.AUTH]: {
        path: RoutePath.auth,
        element: <AuthenticationPageAsync />,
        authOnly: false,
    },
    [AppRoutes.FEED]: {
        path: RoutePath.feed,
        element: <FeedPageAsync />,
        authOnly: true,
    },
    [AppRoutes.POST]: {
        path: RoutePath.post,
        element: <PostDetailPageAsync />,
        authOnly: true,
    },
    [AppRoutes.PROFILE]: {
        path: RoutePath.profile,
        element: <ProfilePageAsync />,
        authOnly: true,
    },
    [AppRoutes.OTHERS_PROFILE]: {
        path: RoutePath.others_profile,
        element: <ProfilePageAsync />,
        authOnly: true,
    },
};
