import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const useThemeStore = defineStore('theme', {
  state: ()=>{
    return {
      isDark: false 
    }
  }
})