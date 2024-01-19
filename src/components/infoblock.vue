<script setup>
const emit = defineEmits(['closeDrawer']);
defineProps({
  title: String,
  description: String,
  imageUrl: String,
})
</script>

<template>
  <div class="flex flex-col items-center text-center gap-2 w-72 mx-auto">
    <img class="w-36 h-36 "
         :src="imageUrl"
         alt="empty"/>
    <h2 class="mt-4 text-gray-600 font-bold text-2xl">{{ title }}</h2>
    <button
        @click="emit('closeDrawer')"
        class="text-slate-600 font-bold bg-lime-100 mt-5 border-lime-500 border-4 transition rounded-3xl p-2 hover:-translate-y-0.5 hover:shadow-xl">
      {{ description }}
    </button>
  </div>
</template>

<style scoped>

</style>