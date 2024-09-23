/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './public/index.html',
    './src/**/*.{js,jsx,ts,tsx}',  // Include all JS/JSX/TS/TSX files in the src directory
  ],
  theme: {
    extend: {
      // Custom theme extensions like colors, spacing, etc.
      colors: {
        primary: '#4F46E5',  // Example: Add a custom primary color
      },
    },
  },
  plugins: [
    // Add any Tailwind plugins here, if needed
    require('@tailwindcss/forms'),    // Example: Form styles plugin
    require('@tailwindcss/typography'), // Example: Typography plugin
  ],
}


