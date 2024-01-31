<script setup>
import CurrentOrdersList from "@/components/current-orders-list.vue";
import CompletedOrder from "@/components/completed-order.vue";
import axios from "axios";
import { ref, onMounted, computed } from "vue";
import CompletedOrdersList from "@/components/completed-orders-list.vue";

const orders = ref([]);

const fetchData = async () => {
  try {
    const response = await axios.get('https://af3c46b46dc5a452.mokky.dev/orders');
    orders.value = response.data;
  } catch (error) {
    console.log(error);
  }
};

onMounted(() => {
  fetchData();
});

const currentOrders = computed(() => orders.value.filter(order => order.status === 0));
const completedOrders = computed(() => orders.value.filter(order => order.status === 1));
</script>

<template>
    <h1 class="uppercase font-bold text-stone-500 text-3xl text-center">my orders</h1>

    <div class="orders__container flex mt-5 gap-10 flex-col">


        <current-orders-list :currentOrders="currentOrders"></current-orders-list>




       <completed-orders-list :completedOrders="completedOrders"></completed-orders-list>
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