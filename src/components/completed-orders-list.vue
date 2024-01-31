<script setup>
import CompletedOrder from "@/components/completed-order.vue";
import {defineProps} from 'vue';

const props = defineProps({
  completedOrders: Array,
});
console.log(props)

</script>

<template>
  <h2 class="uppercase font-bold mt-4 text-stone-600 text-xl text-center">completed</h2>
  <div class="bg-white shadow-xl rounded-2xl" v-for="order in props.completedOrders">
    <completed-order :items="order.items"></completed-order>
  </div>

</template>

<style scoped>

</style>