<script setup>
import { ref } from 'vue';
import LoginComponent from "@/components/login-component.vue";
import RegisterComponent from "@/components/register-component.vue";
import ForgetPasswordComponent from "@/components/forget-password-component.vue";
const currentComponent = ref('login');

const showLogin = () => {
  currentComponent.value = 'login';
};

const showRegister = () => {
  currentComponent.value = 'register';
};

const showForgetPassword = () => {
  currentComponent.value = 'forget-password';
};
</script>

<template>
  <div v-auto-animate>
    <h1 class="text-4xl font-bold text-neutral-500 text-center mb-10">please login or register to continue shopping</h1>
    <div class="flex justify-center gap-12 mb-3 text-neutral-600 text-xl font-bold">
      <button @click="showLogin" :class="{ 'active-button': currentComponent === 'login' }" class="p-4 translate-y-5 rounded-t-2xl" :disabled="currentComponent === 'login'">login</button>
      <button @click="showRegister" :class="{ 'active-button': currentComponent === 'register' }" class="p-4 translate-y-5 rounded-t-2xl" :disabled="currentComponent === 'register'">register</button>
    </div>
    <login-component @show-forget-password="showForgetPassword"
                     @show-register="showRegister"
                     v-if="currentComponent === 'login'"></login-component>
    <register-component @show-login="showLogin" v-if="currentComponent === 'register'"></register-component>
    <forget-password-component v-auto-animate v-if="currentComponent === 'forget-password'"></forget-password-component>
  </div>
</template>




<style scoped>

button {

  background: linear-gradient(180deg, rgba(255, 255, 255, 0) 0%, rgba(245, 245, 245, 1) 100%);
  transition: transform 0.3s ease; /* добавляем transition к свойству transform */
}
.active-button {

  background-color: rgb(245, 245, 245);

}
</style>