<script setup>
import RequestComponent from "@/components/request-component.vue";
import {onMounted, ref} from "vue";
import axios from "axios";
const complaints = ref([]);

const fetchData = async () => {
  try {
    const response = await axios.get('https://af3c46b46dc5a452.mokky.dev/requests');
    complaints.value = response.data;
  } catch (error) {
    console.log(error);
  }
};
onMounted(() => {
  fetchData();
});
</script>

<template>
  <div v-if="complaints.length===0">
    <h2 class="text-center font-bold text-2xl text-neutral-500 p-36" >no requests</h2>
  </div>

  <div v-else class="flex flex-col">
    <request-component :complaints ></request-component>
  </div>

</template>

<style scoped>

</style>