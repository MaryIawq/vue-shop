<script setup>
import { onMounted, ref } from "vue";
import axios from "axios";
import CardListComponent from "@/components/card-list-component.vue";
import CurrentOrdersList from "@/components/current-orders-list.vue";
const orders = ref([]);
onMounted(async() => {
  try {
    const {data} = await axios.get('https://af3c46b46dc5a452.mokky.dev/orders?_relations=items')
    orders.value = data.map(object => object.item);
  } catch (err) {

  }
});
</script>

<template>
  <div class="main flex flex-col items-center justify-center bg-stone-100 gap-10 rounded-2xl">
    <div class="profile__card flex flex-col relative items-center max-w-96 w-full bg-stone-50 shadow-2xl mt-10 rounded-2xl p-7">
      <div class="img__container mb-2.5 relative h-36 w-36 rounded-full border-stone-300 border-4 bg-stone-50">
        <img
            class="h-full w-full object-cover p-4"
            src="/10102alienmonster_109997.png"
            alt="profile image">
      </div>
      <div class="text__data gap-2 flex flex-col items-center">
        <span class="name text-2xl text-stone-700 font-bold capitalize">tessi</span>
        <span class="discount text-lg text-stone-500 font-semibold">your discount 5%</span>
      </div>
      <div class="media__btns gap-2 mt-8 flex items-center">
        <button class="h-11 w-11 border-2 border-stone-400 bg-neutral-200 rounded-3xl transition hover:brightness-110"><img class="p-1" src="/purchases.png" alt="purchases"></button>
        <button class="h-12 w-12 border-2 border-stone-400 bg-white rounded-3xl transition hover:brightness-110"><img class="p-1" src="/payment.png" alt="payment"></button>
        <button class="h-12 w-12 border-2 border-stone-400 bg-neutral-200 rounded-3xl transition hover:brightness-110"><img class="p-1" src="/points.png" alt="points"></button>
        <button class="h-12 w-12 border-2 border-stone-400 bg-white rounded-3xl transition hover:brightness-110"><img class="p-1" src="/requests.png" alt="requests"></button>
      </div>
      <div class="mt-6 flex items-center gap-6">
        <button class=" rounded-3xl border-2 border-stone-400 transition p-2 hover:-translate-y-0.5 hover:shadow-lg">exit</button>
        <button class=" rounded-3xl border-2 border-stone-400 transition p-2 hover:-translate-y-0.5 hover:shadow-lg">change profile</button>
      </div>
    </div>
    <div class="bg-stone-50 relative items-center max-w-96 w-full mb-64 shadow-2xl rounded-2xl p-7">
      <h1 class="uppercase font-bold text-stone-500 text-2xl">my orders</h1>
      <div class="flex mt-5 gap-3 bg-neutral-200 p-4 rounded-2xl flex-col">
        <h2 class="uppercase font-bold text-stone-600 text-center">current</h2>
        <current-orders-list></current-orders-list>
        <div class="bg-white shadow-xl">
          <h2 class="uppercase font-bold text-stone-600">latest</h2>
          <card-list-component :items="orders"></card-list-component>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile__card::before{
  content: '';
  position: absolute;
  height: 36%;
  width: 100%;
  background-image: url("/public/cookies-bg.jpg");
  background-size: cover;
  top: 0;
  left: 0;
  border-radius: 16px 16px 0 0;
}
</style>