import { defineConfig } from 'tailwindcss-patch'
export default defineConfig({
    patch: {
        output: {
            filename: './cmd/classes.json',
        },
        tailwindcss: {
            config: './tailwind.config.js',
        }
    }
})