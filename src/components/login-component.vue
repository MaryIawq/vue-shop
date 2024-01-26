<script setup>
import {ref} from 'vue';
const emit = defineEmits(['show-register', 'show-forget-password']);
const authorizedUsers = ref([]);

const login = () => {
  const email = document.getElementById('email').value;
  const password = document.getElementById('password').value;

  if (email && password) {
    const user = {email, password};
    authorizedUsers.value.push(user);
    document.getElementById('email').value = '';
    document.getElementById('password').value = '';
    console.log(user)
  } else {
  }
}

</script>

<template>
  <div class="flex flex-col items-center justify-center">
    <div
        class="main flex shadow-xl flex-col w-96 border-b border-l border-r bg-neutral-100 justify-center rounded-2xl  p-4 items-center">
      <h2 class="uppercase text-3xl font-bold text-neutral-600">login</h2>
      <div class="login__container border bg-white rounded-2xl mt-6 p-14">
        <div class="flex flex-col mb-2 gap-2">
          <div class="flex gap-2">
            <img src="/email.png" alt="email" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">email</p>
          </div>
          <input id="email" class="border rounded-md p-2" placeholder="enter your email" type="email"/>
        </div>
        <div class="flex flex-col mt-4 gap-2">
          <div class="flex gap-2">
            <img src="/password.png" alt="password" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">password</p>
          </div>
          <input id="password" class="border rounded-md p-2" placeholder="enter your password" type="password"/>
          <button @click="login" class="mt-2">OK</button>
        </div>
        <button @click="() => emit('show-forget-password')"
                class="mt-7 transition text-neutral-600 opacity-75 border-b bg-white p-1 hover:opacity-100">forget your
          password?
        </button>
      </div>
      <div class="flex flex-col items-center justify-center border bg-white rounded-2xl mt-6 p-8">

        <p class="text-neutral-600 text-center font-semibold">don't have a account?</p>
        <button @click="() => emit('showRegister')"
                class="transition mt-3 font-bold text-neutral-600 border bg-white p-2 px-8 rounded-2xl hover:-translate-y-0.5 hover:shadow-lg">
          register
        </button>
      </div>

    </div>
  </div>
</template>


<style scoped>
@media (max-width: 500px) {
  .main {
    width: 300px
  }

  .login__container {
    width: 280px
  }
}

@media (max-width: 350px) {
  .main {
    width: 260px
  }

  .login__container {
    width: 250px
  }
}
</style>