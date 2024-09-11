/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: ["class"],
    content: [
        "./admin/*.templ",
        "./admin/**/*.templ",
    ],
    theme: {
        container: {},
        extend: {},
    },
    daisyui: {
        themes: ["corporate", "dark"]
    },
    plugins: [require("daisyui")],
    corePlugins: {
        preflight: true,
    }
};
