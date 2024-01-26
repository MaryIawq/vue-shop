<script setup>
import { ref } from 'vue';

const emit = defineEmits(['showLogin']);
const newUsers = ref([]);
const isRegistered = ref(false);

const register = () => {
  const name = document.getElementById('name').value;
  const number = document.getElementById('number').value;
  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;
  const confirm = document.getElementById('confirm').value;
  const mismatchDiv = document.getElementById('mismatch');

  if (email && name && number && password === confirm) {
    const user = { email, password, name, number };
    newUsers.value.push(user);
    document.getElementById('email').value = '';
    document.getElementById('name').value = '';
    document.getElementById('number').value = '';
    document.getElementById('password').value = '';
    document.getElementById('confirm').value = '';
    console.log(user);
    mismatchDiv.style.display = 'none';
    isRegistered.value = true;
  } else {
    mismatchDiv.style.display = 'block';
  }
};
</script>

<template>
  <div class="flex flex-col items-center justify-center">
    <div class="main flex shadow-xl flex-col w-96 border bg-neutral-50 justify-center rounded-2xl  p-4 items-center">
      <h2 class="uppercase text-3xl font-bold text-neutral-600">register</h2>
      <div class="register__container gap-5 flex flex-col items-center justify-center border bg-white rounded-2xl mt-6 p-14">

        <div class="flex flex-col mb-2 gap-2">
          <div class="flex gap-2">
            <img src="/name.png" alt="name" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">name</p>
          </div>
          <input id="name" class="border rounded-md p-2" placeholder="enter your name" type="text"/>
        </div>

        <div class="flex flex-col mb-2 gap-2">
          <div class="flex gap-2">
            <img src="/phone.png" alt="phone" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">phone number</p>
          </div>
          <input id="number" class="border rounded-md p-2" placeholder="enter your phone num" type="number"/>
        </div>

        <div class="flex flex-col mb-2 gap-2">
          <div class="flex gap-2">
            <img src="/email.png" alt="email" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">email</p>
          </div>
          <input id="email" class="border rounded-md p-2" placeholder="enter your email" type="email"/>
        </div>

        <div class="flex flex-col mt-4 gap-5">
          <div class="flex gap-2">
            <img src="/password.png" alt="password" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">password</p>
          </div>
          <p class="text-neutral-500 font-semibold">create a password</p>
          <input id="password" class="border rounded-md p-2" placeholder="enter your password" type="password"/>
          <p class="text-neutral-500 font-semibold">confirm password</p>
          <input id="confirm" class="border rounded-md p-2" placeholder="confirm your password" type="password"/>
          <div id="mismatch" style="display: none;">
            <p class="text-rose-400">Passwords do not match.</p>
          </div>
        </div>

        <button @click="register" class="transition mt-10 font-bold text-neutral-600 border bg-white p-2 px-8 rounded-2xl hover:-translate-y-0.5 hover:shadow-lg">Register</button>
      </div>
      <div class="flex flex-col items-center justify-center border bg-white rounded-2xl mt-6 p-8">

        <p class="text-neutral-600 text-center font-semibold">do you have a account?</p>

          <button @click="() => emit('showLogin')" class="transition mt-3 font-bold text-neutral-600 border bg-white p-2 px-8 rounded-2xl hover:-translate-y-0.5 hover:shadow-lg">log in</button>
      </div>
    </div>
    <div v-if="isRegistered" class="absolute inset-0 flex items-center justify-center bg-gray-900 bg-opacity-75 z-50">
      <div class="bg-white flex flex-col justify-center items-center p-8 rounded-xl">
        <p class="text-3xl text-green-600 font-bold">you've successfully registered!</p>
        <router-link to="/">
          <button @click="isRegistered = false" class="mt-4 bg-stone-600 text-white py-2 px-4 rounded-2xl hover:brightness-110">close</button>
        </router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
@media (max-width: 500px) {
  .main {
    width: 300px
  }
  .register__container {
    width: 280px
  }
}
@media (max-width: 350px) {
  .main {
    width: 260px
  }
  .register__container {
    width: 250px
  }
}
</style>