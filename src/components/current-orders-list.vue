<script setup>
import CurrentOrder from "@/components/current-order.vue";
import {defineProps} from 'vue';

const props = defineProps({
  currentOrders: Array,
});
console.log(props)

</script>

<template>
  <h2 class="uppercase font-bold mt-4 text-stone-600 text-xl text-center">current</h2>
  <div class="bg-white shadow-xl rounded-2xl" v-for="order in props.currentOrders">
    <current-order :items="order.items"></current-order>
  </div>

</template>

<style scoped>

</style>