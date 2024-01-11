<script setup>
import headerComponent from './components/header-component.vue'
import cardList from './components/card-list-component.vue'
</script>

<template>


  <div class="w-5/6 m-auto bg-white rounded-2xl shadow-2xl mt-14">
    <header-component></header-component>
    <div class="p-10">
      <h2 class="text-3xl font-bold text-slate-700 mb-8">all sweets</h2>
      <card-list></card-list>
    </div>

  </div>

</template>

<style scoped>

</style>
