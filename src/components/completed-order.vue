<script setup>
import OrderCard from "@/components/order-card.vue";
import { defineProps } from 'vue';

const props = defineProps({
  items: Array,
});
console.log(props)
</script>

<template>
  <div v-for="item in props.items">
    <div class="bg-white p-3 rounded-2xl">
      <order-card
          :title="item.title"
          :imageUrl="item.imageUrl"
          :weight="item.weight"></order-card>
    </div>
  </div>


  <div class="text-slate-600 bg-neutral-100 p-2 pl-4 flex flex-col rounded-2xl mt-3">
    <div class="order__info flex justify-between border-b-2 mb-2 border-stone-400">
      <b>order status: </b>
      <p>end</p>
    </div>
    <div class="order__info flex justify-between border-b-2 mb-2 border-stone-400">
      <b>payable: </b>
      <p>1983 ₽</p>
    </div>
    <div class="order__info flex justify-between border-b-2 mb-2 border-stone-400">
      <b>delivered in: </b>
      <p>31 min</p>
    </div>
    <button
        class="p-2 border border-stone-300 mx-16 mt-4 transition rounded-2xl bg-white hover:-translate-y-0.5 hover:shadow-lg">
      repeat
    </button>
  </div>
</template>

<style scoped>
@media (max-width: 425px) {
  button {
    width: fit-content;
  }
}
</style>