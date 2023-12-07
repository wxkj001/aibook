<script setup lang="ts">
import { defineComponent, ref } from 'vue'
import { darkTheme } from 'naive-ui'
import { useThemeStore } from './stores/theme';
const theme = useThemeStore()
</script>

<template>
  <n-config-provider :theme="theme.$state.isDark ? darkTheme : null">
    <RouterView />
  </n-config-provider>
</template>

<style scoped>
</style>
