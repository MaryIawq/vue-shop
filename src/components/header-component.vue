<script setup>

import EmailFormComponent from "@/components/email-form-component.vue";
import {ref} from "vue";

const emit = defineEmits(['openDrawer'])
defineProps({
  totalPrice: Number
})
const emailVisible = ref(false)
const showEmail = () => {
  emailVisible.value = !emailVisible.value;
}
</script>

<template>
  <header>
    <div class="top__header flex gap-3 border-b pb-2 justify-between mt-2">
        <div class="flex transition gap-1 opacity-75 pt-5 px-4 hover:opacity-100">
          <img src="/whatsapp_icon.png" class="w-7 h-7" alt="whatsapp">
          <span class="mt-1">+71234567890</span>
        </div>
        <div @click="showEmail" class="flex transition gap-1 opacity-75 pt-5 px-4 hover:opacity-100">
          <img src="/email.png" class="w-7 h-7" alt="email">
          <span class="mt-1">candystore@gmail.com</span>
        </div>
        <div class="flex transition gap-1 opacity-75 pt-5 px-4 hover:opacity-100">
          <img src="/phone-callback.png" class="w-6 h-6" alt="email">
          <span class="mt-0.5">callback</span>
        </div>
    </div>


    <div class="bottom__header flex justify-between border-b border-gray-200 py-10 px-4 md:pr-12">
      <router-link to="/">
        <div class="flex mr-6 items-center gap-4" id="logo">
          <img src="/logo.png" alt="logo" class="w-16 h-16">
          <div>
            <h1 class="text-2xl font-bold uppercase text-slate-700">candy store</h1>
            <p class="text-lg text-slate-500">great selection of sweets</p>
          </div>
        </div>
      </router-link>
      <ul class="flex items-center gap-6 mt-1 mr-3">
        <router-link to="/">
          <li class="flex opacity-75 transition hover:opacity-100 hover:-translate-y-0.5 items-center cursor-pointer gap-0.5">
            <img src="/home.png" class="w-7 opacity-75" alt="home">
            <p class="text-sm md:text-base text-slate-700">home</p>
          </li>
        </router-link>
        <li @click="() => emit('openDrawer')" class="flex opacity-75 transition hover:opacity-100 hover:-translate-y-0.5 items-center cursor-pointer gap-0.5">
          <img src="/cart.svg" class="w-7" alt="cart">
          <p class="text-sm md:text-base text-slate-700">{{ totalPrice }}₽</p>
        </li>
        <router-link to="/Favorites">
          <li class="flex opacity-75 transition hover:opacity-100 hover:-translate-y-0.5 items-center cursor-pointer gap-0.5">
            <img src="/favorites.ico" class="w-7" alt="favorites">
            <p class="text-sm md:text-base text-slate-700">favorites</p>
          </li>
        </router-link>
        <router-link to="/Profile">
          <li class="flex opacity-75 transition hover:opacity-100 hover:-translate-y-0.5 items-center cursor-pointer gap-0.5">
            <img src="/profile.svg" class="w-7" alt="profile">
            <p class="text-sm md:text-base text-slate-700">profile</p>
          </li>
        </router-link>
        <router-link to="/info">
          <li class="flex opacity-75 transition hover:opacity-100 hover:-translate-y-0.5 items-center cursor-pointer gap-0.5">
            <img src="/info.ico" class="w-7 opacity-90" alt="info">
            <span class="mt-0.5">info</span>
          </li>
        </router-link>
      </ul>
    </div>
  </header>
<div v-if="emailVisible" class="z-50 email__form shadow-xl border-2 rounded-2xl p-5 bg-neutral-50 absolute">
  <button class="rounded-full" @click="showEmail"><img class="w-6 h-6" src="/close.ico"></button>
  <email-form-component></email-form-component>
</div>
</template>

<style scoped>
@media (max-width: 768px) {
  ul {
    display: flex;
    align-items: flex-start;
    flex-wrap: wrap;
    height: 100%;
  }
  .top__header {
    display: flex;
    flex-wrap: wrap;
    height: 100%;
  }
}
@media (max-width: 425px) {
  .bottom__header {
    display: flex;
    flex-direction: column;
  }
}
.email__form {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 50;

}
</style>