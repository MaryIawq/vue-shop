<script setup>
const emit = defineEmits(['openDrawer'])

</script>

<template>
  <header class="flex justify-between border-b border-gray-200 pr-12 py-10 pl-4">
    <div class="flex items-center gap-4" id="logo">
      <img src="/logo.png" alt="logo" class="w-16 h-16">
      <div>
        <h1 class="text-2xl font-bold uppercase text-slate-700">candy store</h1>
        <p class="text-lg text-slate-500">store for the best sweets</p>
<!--        <p class="text-sm font-thin text-slate-400">delivery every day 11 am to 9 pm</p>-->
      </div>
    </div>
    <ul class="flex items-center gap-11">
      <li @click="() => emit('openDrawer')" class="flex items-center cursor-pointer gap-4 mr-5">
        <img src="/cart.svg" class="max-w-7 opacity-70 hover:opacity-100" alt="cart">
        <p class="text-slate-600 hover:text-slate-800">1456₽</p>
      </li>
      <li class="flex items-center cursor-pointer gap-3">
        <img src="/favorites.ico" class="max-w-7 opacity-75 hover:opacity-100" alt="favorites">
        <p class="text-slate-600 hover:text-slate-800">favorites</p>
      </li>
      <li class="flex items-center cursor-pointer gap-3">
        <img src="/profile.svg" class="max-w-7 opacity-75 hover:opacity-100" alt="">
        <p class="text-slate-600 hover:text-slate-800">profile</p>
      </li>
    </ul>
  </header>
</template>

<style scoped>
@media (max-width: 768px) {
  ul {
    display: flex;
    flex-direction: column;
    margin-left: 30px
  }

  #logo {
    margin-right: 50px;
    display: flex;
    justify-content: flex-start;
  }

  @media (max-width: 375px) {
    ul {
      display: flex;
      align-items: flex-start;
      margin-top: 30px;
    }
    header {
      display: flex;
      flex-direction: column;
      height: 100%;
    }
  }
}

</style>