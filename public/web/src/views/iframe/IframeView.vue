<script setup lang="ts">
import { ref } from 'vue';
import { useRouter, useRoute } from 'vue-router'
const route=useRoute()
console.log("r",route.params.href);

const href=ref<string>(route.params.href as string)
</script>

<template>
    <iframe :src="href" frameborder="0"></iframe>
</template>