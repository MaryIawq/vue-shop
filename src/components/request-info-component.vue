<script setup>
import { ref, onMounted } from 'vue';
import axios from "axios";
import RequestsList from "@/components/requests-list.vue";
const selectedTopic = ref([""])
const topics = ["long waiting time", "poor product quality", "the wrong product was delivered" ];
const complaintText = ref("");

const saveComplaint = async () => {
  if (selectedTopic.value && complaintText.value) {
    const complaint = `${selectedTopic.value}: ${complaintText.value}`;
    try {
      await axios.post('https://af3c46b46dc5a452.mokky.dev/requests', { complaint });
      alert('Complaint saved successfully');
      selectedTopic.value = "";
      complaintText.value = "";
    } catch (error) {
      console.error('Error saving complaint:', error.message);
    }
  } else {
    alert("Please fill in all fields");
  }
};



</script>

<template>
  <div class="relative items-center w-full mb-64 shadow-2xl rounded-2xl p-7">
    <h2 class="text-center font-bold text-neutral-600 text-3xl p-2 mb-4">Create request</h2>
    <div class="bg-neutral-100 rounded-2xl border p-4 shadow-lg ">
      <div class="flex flex-col gap-2 mb-5 text-neutral-600 text-xl">
        <label class="p-2 text-center font-bold" >Choose a topic:</label>
        <select class="border-2 bg-white p-1 rounded-2xl" v-model="selectedTopic">
          <option v-for="topic in topics" :key="topic">{{ topic }}</option>
        </select>
      </div>
      <div class="text-neutral-600 text-xl gap-4 flex flex-col">
        <label class="font-bold text-center mt-8"> Write the text of the appeal: </label>
        <textarea class="border-2 bg-white p-2 text-lg font-thin rounded-2xl" v-model="complaintText"></textarea>
      </div>
    </div>
    <button class="text-neutral-600 transition p-2 border rounded-3xl w-28 bg-white mt-8 text-xl hover:-translate-y-0.5 hover:shadow-lg" @click="saveComplaint">Save request</button>
    <div class="my__requests mt-10">
      <h2 class="text-center font-bold text-neutral-600 text-3xl p-2 mb-4">My requests</h2>
      <requests-list></requests-list>
    </div>
  </div>

</template>

<style scoped>

</style>