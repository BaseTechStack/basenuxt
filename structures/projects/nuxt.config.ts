// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
 
  future: { compatibilityVersion: 4 },

  // Add Nuxt UI module for this layer
  modules: ['@nuxt/ui'],
  
 
  css: ['~/assets/css/main.css'],
})
