<script setup>
import {inject} from "vue";
import cartComponent from '/src/components/cart-component.vue'

const { cart } = inject('cart')
</script>

<template>
  <div class="flex flex-col">
    <cart-component
        v-for="item in cart"
        :key="item.id"
        :title="item.title"
        :image-url="item.imageUrl"
        :weight="item.weight"
        :price="item.price"></cart-component>
  </div>

</template>

<style scoped>

</style>