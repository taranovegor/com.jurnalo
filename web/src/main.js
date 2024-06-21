import { createApp } from 'vue'
import PrimeVue from 'primevue/config'
import { updatePreset } from '@primevue/themes'
import Aura from '@primevue/themes/aura'

import './style.css'
import App from './App.vue'

const app = createApp(App)

app.use(PrimeVue, {theme: {preset: Aura,}})
updatePreset({
    semantic: {
        primary: {
            50: '{neutral.50}',
            100: '{neutral.100}',
            200: '{neutral.200}',
            300: '{neutral.300}',
            400: '{neutral.400}',
            500: '{neutral.500}',
            600: '{neutral.600}',
            700: '{neutral.700}',
            800: '{neutral.800}',
            900: '{neutral.900}',
            950: '{neutral.950}',
        },
    },
})

app.mount('#app')
