<script setup>

import {ref, watch, provide, computed} from "vue";
import axios from "axios";
import headerComponent from './components/header-component.vue'
import drawerComponent from './components/drawer-component.vue'
import Home from './pages/home.vue'


/* Drawer (START) */
const cart = ref([])
const isCreatingOrder = ref(false)
const drawerOpen = ref(false)

const totalPrice = computed(() => cart.value.reduce((acc, item) => acc + item.price, 0))
const discountPrice = computed(() => Math.round((totalPrice.value * 9) / 100))
const deliveryPrice = computed(() => Math.round((totalPrice.value * 5) / 100))
const finalPrice = computed(() => {
  const total = totalPrice.value + deliveryPrice.value - discountPrice.value;
  return Math.max(0, total);
});

const cartButtonDisabled = computed(() =>
    isCreatingOrder.value ? true : totalPrice.value ? false : true
)

const closeDrawer = () => {
  drawerOpen.value = false
}
const openDrawer = () => {
  drawerOpen.value = true
}

const addToCart = (item) => {
  cart.value.push(item)
  item.isAdded = true
}

const removeFromCart = (item) => {
  cart.value.splice(cart.value.indexOf(item), 1)
  item.isAdded = false
}


const createOrder = async () => {
  try {
    isCreatingOrder.value = true
    const {data} = await axios.post(`https://af3c46b46dc5a452.mokky.dev/orders`, {
      items: cart.value,
      totalPrice: totalPrice.value
    })
    cart.value = []
    return data
  } catch (e) {
    console.log(e)
  } finally {
    isCreatingOrder.value = false
  }
}
/* Drawer (END) */



watch(cart, () => {
      localStorage.setItem('cart', JSON.stringify(cart.value))
    },
    {deep: true}
)

provide('cart', {
  cart,
  addToCart,
  removeFromCart,
})
</script>


<template>
  <div class="w-5/6 sm:w-11/12 m-auto bg-white rounded-2xl shadow-2xl mt-14">
    <header-component
        :final-price="finalPrice"
        :total-price="totalPrice"
        @open-drawer="openDrawer"></header-component>
    <div class="cart-container overflow-y-auto overflow-x-hidden">
      <drawer-component v-if="drawerOpen"
                        :final-price="finalPrice"
                        :total-price="totalPrice"
                        :delivery-price="deliveryPrice"
                        :discount-price="discountPrice"
                        @create-order="createOrder"
                        :cart-button-disabled="cartButtonDisabled"
                        @close-drawer="closeDrawer"></drawer-component>
    </div>
    <div class="p-7 main">
      <home></home>
    </div>
  </div>
</template>

<style scoped>
.cart-container {
  max-height: 60vh;
}


</style>
