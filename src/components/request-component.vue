<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';

const complaints = ref([]);

const fetchData = async () => {
  try {
    const response = await axios.get('https://af3c46b46dc5a452.mokky.dev/requests');
    complaints.value = response.data;
  } catch (error) {
    console.error('There was an error!', error);
  }
};

onMounted(() => {
  fetchData();
});
</script>

<template>
  <h2 v-if="complaints.length === 0">no requests</h2>
  <div v-else class="bg-white p-3 border transition rounded-2xl shadow-lg hover:-translate-y-0.5 hover:shadow-xl">
    <div class="flex flex-col justify-center items-center">
      <div class="flex gap-5 bg-red-50 p-2 rounded-2xl border mb-2 shadow-lg">
        <h2 class="uppercase font-bold text-stone-600 text-lg">Request</h2>
        <p class="font-bold text-stone-600 text-lg">Date: 03.01.2024</p>
      </div>
      <h2 class="font-bold flex mt-2 text-rose-800 border-b border-rose-800 text-lg" v-for="complaint in complaints" :key="complaint.id">{{ complaint.complaint }}</h2>
    </div>
    <div class="flex flex-col justify-center items-center">
      <h2 class="font-bold flex text-green-700 text-lg">answer<img src="/green-checkmark.png" class="w-5 mt-1 h-5"
                                                                   alt="checkmark"></h2>
      <p class="self-start">anscsdcswerdcxass</p>
    </div>
  </div>
</template>

<style scoped>

</style>