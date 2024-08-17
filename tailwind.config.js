/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./frontend/**/*.{html,tmpl,tmpl.html,js,ts,jsx,tsx}",
  ],
  theme: {
    "fontFamily": {
      "sans": ["Inter Tight", "sans-serif"],
    },
    extend: {
      colors: {
        "primary": "#17A34A",
        "secondary": "#00FF00",
        "tertiary": "#0000FF",
      },
    },
  },
  plugins: [],
}

