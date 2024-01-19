<script setup>
import {inject} from "vue";
import cartComponent from '/src/components/cart-component.vue'

const { cart, removeFromCart } = inject('cart')

</script>

<template>
  <div class="flex flex-col flex-1 justify-between" v-auto-animate>
    <cart-component
        v-for="item in cart"
        :key="item.id"
        :title="item.title"
        :image-url="item.imageUrl"
        :weight="item.weight"
        :price="item.price"
        @on-click-remove="() => removeFromCart(item)"></cart-component>
  </div>

</template>

<style scoped>

</style>