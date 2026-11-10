import path from 'path';
import webpack from 'webpack';
import { buildWebpackConfig } from './config/build/buildWebpackConfig';
import { BuildEnv, BuildMode, BuildPaths } from './config/build/types/config';

export default (env: BuildEnv) => {
    const paths: BuildPaths = {
        entry: path.resolve(__dirname, 'src', 'app', 'index.tsx'),
        output: path.resolve(__dirname, 'dist'),
        html: path.resolve(__dirname, 'public', 'index.html'),
        src: path.resolve(__dirname, 'src'),
    };

    const mode: BuildMode = env.mode || 'development';
    const isDev = mode === 'development';
    const port = Number(env.port) || 3000;
    const backendUrl = process.env.API_URL || env.proxyTarget || 'http://127.0.0.1:8080';
    const apiUrl = isDev ? '/api' : process.env.API_URL || '/api';

    const config: webpack.Configuration = buildWebpackConfig({
        mode,
        paths,
        isDev,
        port,
        apiUrl,
        proxyTarget: backendUrl,
    });

    return config;
};
