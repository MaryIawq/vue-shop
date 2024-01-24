<script setup>
import {inject, onMounted, ref, watch} from "vue";
import axios from "axios";
import cardListComponent from "../components/card-list-component.vue"

/*FETCH FAVORITES START*/
const favorites   = ref([]);
onMounted(async() => {
try {
const {data} = await axios.get('https://af3c46b46dc5a452.mokky.dev/favorites?_relations=items')
  favorites.value = data.map(object => object.item);
} catch (err) {
}});

/*FETCH FAVORITES END*/

/*ADD/REMOVE TO CART START*/
const {cart, addToCart, removeFromCart} = inject('cart')

const onClickAddPlus = (item) => {
  if (!item.isAdded) {
    addToCart(item)
  } else {
    removeFromCart(item)
  }
}

onMounted(async () => {
  const localCart = localStorage.getItem('cart');
  cart.value = localCart ? JSON.parse(localCart) : []

  await fetchItems();
  await fetchFavorites();

  items.value = items.value.map((item) => ({
    ...item,
    isAdded: cart.value.some((cartItem) => cartItem.id === item.id)
  }))
});

watch(cart, () => {
  items.value = items.value.map((item) => ({
    ...item,
    isAdded: false
  }))
})
/*ADD/REMOVE TO CART END*/


</script>

<template>
  <div class="flex justify-between">
    <h1 class="uppercase font-extrabold text-stone-600 pb-4 text-3xl">My Favorites</h1>
    <router-link to="/"><button class="flex opacity-60 font-extrabold cursor-pointer text-2xl uppercase hover:opacity-80"><img class="w-8" src="/public/back.ico" alt=""/>Home</button></router-link>
  </div>
  <card-list-component
      :items="favorites"
      @add-to-cart="onClickAddPlus"
  ></card-list-component>
</template>

<style scoped>
@media (max-width: 425px) {
  h1,
  button {
    font-size: 17px;
    padding: 4px
  }
}
</style>