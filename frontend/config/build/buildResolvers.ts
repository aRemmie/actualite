import path from 'path';
import webpack from 'webpack';
import { BuildOptions } from './types/config';

export function buildResolvers(options: BuildOptions): webpack.ResolveOptions {
    const { src } = options.paths;

    return {
        extensions: ['.tsx', '.ts', '.js'],
        modules: [src, 'node_modules'],
        alias: {
            '@components': path.resolve(src, 'components'),
            '@assets': path.resolve(src, 'assets'),
            '@models': path.resolve(src, 'models'),
            '@services': path.resolve(src, 'services'),
            '@api': path.resolve(src, 'app', 'api'),
            '@pages': path.resolve(src, 'pages'),
        },
    };
}
