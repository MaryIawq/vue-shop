<script setup>

import CurrentOrdersList from "@/components/current-orders-list.vue";
import CompletedOrdersList from "@/components/completed-orders-list.vue";
</script>

<template>

    <div class="orders__parent__container bg-stone-50 relative items-center max-w-96 w-full mb-64 shadow-2xl rounded-2xl p-7">
      <h1 class="uppercase font-bold text-stone-500 text-2xl text-center">my orders</h1>
      <div class="orders__container flex mt-5 gap-3 bg-neutral-200 p-4 rounded-2xl flex-col">
        <h2 class="uppercase font-bold text-stone-600 text-center">current</h2>
        <current-orders-list></current-orders-list>
        <h2 class="uppercase font-bold mt-4 text-stone-600 text-center">completed</h2>
        <div class="bg-white shadow-xl">
          <completed-orders-list></completed-orders-list>
        </div>
      </div>
    </div>

</template>

<style scoped>
@media (max-width: 425px) {
.orders__parent__container {
  padding: 0;
}
  .orders__container {
    padding: 0;
  }
}
</style>