import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

interface UserInfo {
  name: string
  token: string
}
export const useUserStore = defineStore('user', {
  state: ()=>{
    return {
      user: null as UserInfo | null
    }
  }
})
