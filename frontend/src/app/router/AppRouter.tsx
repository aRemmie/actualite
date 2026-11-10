import { ReactNode, Suspense } from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import NotFoundPageAsync from '@pages/NotFoundPage/NotFoundPage.async';
import { useAuthStore } from '@stores/useAuthStore';
import { AppRouteConfig, routeConfig } from './routeConfig';
import { RoutePath } from './paths';

const applyAccess = (route: AppRouteConfig, isAuth: boolean): ReactNode => {
    if (route.access === 'authOnly' && !isAuth) {
        return <Navigate to={RoutePath.auth} replace />;
    }
    if (route.access === 'guestOnly' && isAuth) {
        return <Navigate to={RoutePath.feed} replace />;
    }
    return route.element;
};

const AppRouter = () => {
    const isAuth = useAuthStore((state) => state.isAuth);
    const home = isAuth ? RoutePath.feed : RoutePath.auth;

    return (
        <BrowserRouter>
            <Suspense fallback={<div />}>
                <Routes>
                    {Object.values(routeConfig).map((route) => (
                        <Route
                            key={route.path}
                            path={route.path}
                            element={applyAccess(route, isAuth)}
                        />
                    ))}
                    <Route path="/" element={<Navigate to={home} replace />} />
                    <Route path="*" element={<NotFoundPageAsync />} />
                </Routes>
            </Suspense>
        </BrowserRouter>
    );
};

export default AppRouter;
