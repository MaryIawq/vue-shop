<script setup>

import OrderCardComponent from "@/components/order-card-component.vue";
</script>

<template>
  <div class="bg-white p-3 rounded-2xl">

    <order-card-component></order-card-component>
    <order-card-component></order-card-component>
    <order-card-component></order-card-component>
    <order-card-component></order-card-component>
    <div class="text-slate-600 bg-neutral-100 p-2 pl-4 rounded-2xl font-semibold mt-3">
      <p class="">order status: in delivery</p>
      <p>payable: 764 ₽</p>
      <p></p>
    </div>
  </div>

</template>

<style scoped>

</style>