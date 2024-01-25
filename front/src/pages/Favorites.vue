<script setup>
import { onMounted, ref } from "vue";
import axios from "axios";
import cardListComponent from "../components/card-list-component.vue"
const favorites   = ref([]);
onMounted(async() => {
try {
const {data} = await axios.get('https://af3c46b46dc5a452.mokky.dev/favorites?_relations=items')
  favorites.value = data.map(object => object.item);
} catch (err) {

}
});
</script>

<template>
  <div class="flex justify-between">
    <h1 class="uppercase font-extrabold text-slate-600 pb-4 text-3xl">My Favorites</h1>
    <router-link to="/"><button class="flex opacity-60 font-extrabold cursor-pointer text-2xl uppercase hover:opacity-80"><img class="w-8" src="/public/back.ico" alt=""/>Home</button></router-link>
  </div>

  <card-list-component
      :items="favorites"
  ></card-list-component>
</template>

<style scoped>

</style>