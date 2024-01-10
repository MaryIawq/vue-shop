<script setup>

</script>

<template>
  <header class="flex justify-between border-b border-gray-300 pr-12 py-10 pl-4">
      <div class="flex items-center gap-4" id="logo">
        <img src="/logo.png" alt="logo" class="w-16 h-16">
        <div>
          <h1 class="text-2xl font-bold uppercase text-slate-700">Vue Candies</h1>
          <p class="text-lg text-slate-500">store for the best sweets</p>
        </div>
      </div>
      <ul class="flex items-center gap-11">
        <li class="flex items-center cursor-pointer gap-4 mr-5">
          <img src="/cart.svg" class="max-w-7" alt="cart">
          <p class="text-slate-600 hover:text-slate-800">1456₽</p>
        </li>
        <li class="flex items-center cursor-pointer gap-3">
          <img src="/favorites.ico" class="max-w-7" alt="favorites">
          <p class="text-slate-600 hover:text-slate-800">favorites</p>
        </li>
        <li class="flex items-center cursor-pointer gap-3">
          <img src="/profile.svg" class="max-w-7" alt="">
          <p class="text-slate-600 hover:text-slate-800">profile</p>
        </li>
      </ul>
    </header>
</template>

<style scoped>
@media (max-width: 800px) {
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
}

</style>