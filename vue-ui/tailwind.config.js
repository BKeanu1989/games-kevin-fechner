/** @type {import('tailwindcss').Config} */

import typography from '@tailwindcss/typography'
import forms from '@tailwindcss/forms'
import aspectRatio from '@tailwindcss/aspect-ratio'

export default {
  content: ['./src/**/*.{html,js,vue}'],
  darkMode: 'selector',
  theme: {
    extend: {}
  },
  plugins: [typography, forms, aspectRatio]
}
