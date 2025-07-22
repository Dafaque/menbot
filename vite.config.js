import { dirname, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'
import { defineConfig } from 'vite'

const __dirname = dirname(fileURLToPath(import.meta.url))

export default defineConfig({
    base: '/',
    build: {
        minify: true,
        manifest: false,
        rollupOptions: {
            output: [{
                inlineDynamicImports: true,
                dir: 'assets/web',
            }],
        },
    },
})