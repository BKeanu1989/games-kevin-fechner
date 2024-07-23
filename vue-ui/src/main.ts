import { createApp } from 'vue'
import { createPinia } from 'pinia'

import './assets/main.css'
// import './input.css'
import './assets/input.css'
// import './output.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
