import { ReactNode } from 'react';
import AuthenticationPageAsync from '@pages/AuthenticationPage/AuthenticationPage.async';
import FeedPageAsync from '@pages/FeedPage/FeedPage.async';
import PostDetailPageAsync from '@pages/PostDetailPage/PostDetailPage.async';
import ProfilePageAsync from '@pages/ProfilePage/ProfilePage.async';
import { AppRoutes, RoutePath } from './paths';

export { AppRoutes, RoutePath };

export type RouteAccess = 'any' | 'authOnly' | 'guestOnly';

export interface AppRouteConfig {
    path: string;
    element: ReactNode;
    access: RouteAccess;
}

export const routeConfig: Record<AppRoutes, AppRouteConfig> = {
    [AppRoutes.AUTH]: {
        path: RoutePath.auth,
        element: <AuthenticationPageAsync />,
        access: 'guestOnly',
    },
    [AppRoutes.FEED]: {
        path: RoutePath.feed,
        element: <FeedPageAsync />,
        access: 'authOnly',
    },
    [AppRoutes.POST]: {
        path: RoutePath.post,
        element: <PostDetailPageAsync />,
        access: 'authOnly',
    },
    [AppRoutes.PROFILE]: {
        path: RoutePath.profile,
        element: <ProfilePageAsync />,
        access: 'authOnly',
    },
    [AppRoutes.OTHERS_PROFILE]: {
        path: RoutePath.others_profile,
        element: <ProfilePageAsync />,
        access: 'authOnly',
    },
};
