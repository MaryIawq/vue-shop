<script setup>

import OrderCardComponent from "@/components/order-card-component.vue";
</script>

<template>
  <div class="bg-white p-3 rounded-2xl">
    <order-card-component></order-card-component>
    <order-card-component></order-card-component>
    <order-card-component></order-card-component>
    <div class="text-slate-600 bg-neutral-100 p-2 pl-4 rounded-2xl mt-3">
      <div class="flex justify-between border-b-2 mb-2 border-stone-400">
        <b>order status: </b>
        <p>in delivery</p>
      </div>
      <div class="flex justify-between border-b-2 mb-2 border-stone-400">
        <b>payable: </b>
        <p>764 ₽</p>
      </div>
      <div class="flex justify-between border-b-2 mb-2 border-stone-400">
        <b>delivered in: </b>
        <p>20 min</p>
      </div>
      <div class="btn__group flex gap-2 mt-4">
        <button class="p-2 border rounded-2xl bg-white transition hover:-translate-y-0.5 hover:shadow-lg">contact the courier</button>
        <button class="p-2 border rounded-2xl bg-white transition hover:-translate-y-0.5 hover:shadow-lg">map</button>
      </div>
    </div>
  </div>

</template>

<style scoped>
@media (max-width:760px) {
  .btn__group {
    flex-direction: column;
  }
}
</style>