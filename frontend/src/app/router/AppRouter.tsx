import { Suspense } from 'react';
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import NotFoundPageAsync from '@pages/NotFoundPage/NotFoundPage.async';
import { routeConfig } from './routeConfig';
import { RoutePath } from './paths';
import { useAuthStore } from '../../stores/useAuthStore';

const AppRouter = () => {
    const isAuth = useAuthStore((state) => state.isAuth);
    const home = isAuth ? RoutePath.feed : RoutePath.auth;

    return (
        <BrowserRouter>
            <Suspense fallback={<div />}>
                <Routes>
                    {Object.values(routeConfig).map((route) => {
                        const element =
                            route.authOnly === isAuth ? (
                                route.element
                            ) : (
                                <Navigate to={home} replace />
                            );
                        return <Route key={route.path} path={route.path} element={element} />;
                    })}
                    <Route path="/" element={<Navigate to={home} replace />} />
                    <Route path="*" element={<NotFoundPageAsync />} />
                </Routes>
            </Suspense>
        </BrowserRouter>
    );
};

export default AppRouter;
