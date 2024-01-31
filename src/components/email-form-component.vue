<script setup>
import { ref } from 'vue'


const email = ref('');
const subject = ref('');
const message = ref('');

const submitForm = () => {
  console.log('Email:', email.value);
  console.log('Subject:', subject.value);
  console.log('Message:', message.value);
};
</script>

<template>
  <div class=" max-w-96 bg-white rounded-2xl border p-5 shadow-xl">

      <h2 class="text-center text-xl text-stone-700 font-bold mb-5">Send email</h2>



    <form @submit.prevent="submitForm" class="form">
      <div class="">
        <label for="email" class="text-neutral-700 text-lg">Your email:</label>
        <input type="email" id="email" v-model="email" class="w-full p-1 text-stone-700 mb-4 rounded-2xl border-2 focus:outline-none" required>
      </div>
      <div class="form-group">
        <label for="subject" class="text-neutral-700 text-lgl">Theme:</label>
        <input type="text" id="subject" v-model="subject" class="w-full p-1 text-stone-700 mb-4 rounded-2xl border-2 focus:outline-none" required>
      </div>
      <div class="form-group">
        <label for="message" class="text-neutral-700 text-lg">Message:</label>
        <textarea id="message" v-model="message" rows="4" class="w-full p-1 text-stone-700 mb-4 rounded-2xl border-2 focus:outline-none" required></textarea>
      </div>
      <button type="submit" class="bg-orange-300 p-3 transition hover:rotate-12 rounded-2xl text-white font-bold">Отправить</button>
    </form>
  </div>
</template>



<style scoped>

</style>