/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{tsx,ts,js}"],
  theme: {
    extend: {
      colors: {
        primary: '#f3ac57',
        secondary: '#fbd09d',
        tertiary: '#f0dabe',
        brown: '#aa8c84',
        beige: '#c1a99d',
        background: '#fbfbe3',
        danger: '#ef4444',
        border: "#754540",
      },
      textColor: {
        default: '#754540',
        danger: '#ef4444',
      },
      fontFamily: {
        roboto: ['Roboto', 'sans-serif'],
        arial: ['Arial', 'sans-serif'],
      },
      fontSize: {
        heading: '24px',
        body: '16px',
        emphasis: '18px',
      },
      fontWeight: {
        bold: '700',
        regular: '400',
        semiBold: '600',
      },
    },
  },
  plugins: [],
}

