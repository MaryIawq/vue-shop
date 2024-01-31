<script setup>
const props = defineProps({
  complaints: Array})

</script>

<template>

  <div v-for="complaint in complaints" :key="complaint.id" class="bg-white p-3 border m-2 transition rounded-2xl shadow-lg hover:-translate-y-0.5 hover:shadow-xl">
    <div class="flex flex-col justify-center items-center">
      <div class="flex gap-2 bg-red-50 p-2 rounded-2xl border mb-2 shadow-lg">
        <p class="font-bold text-stone-600 text-lg">Date: {{complaint.complaintDate}}</p>
      </div>
      <h2 class="font-bold flex mt-2 text-rose-800 border-b border-rose-800 text-lg">{{ complaint.selectedTopic }}</h2>
      <p class="mt-3 text-neutral-600">{{ complaint.complaintText }}</p>
    </div>
    <div class="flex flex-col mt-4 justify-center items-center">
      <h2 class="font-bold text-green-700 text-lg">answer:</h2>
      <p class="mt-5 text-neutral-600">{{ complaint.answerText }}</p>
    </div>
  </div>
</template>

<style scoped>

</style>