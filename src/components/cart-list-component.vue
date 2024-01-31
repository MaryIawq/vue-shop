<script setup>
import {inject} from "vue";
import cartComponent from '/src/components/cart-component.vue'

const { cart, removeFromCart } = inject('cart')

</script>

<template>
  <div class="flex cart__container flex-col flex-1 justify-between overflow-y-auto" v-auto-animate>
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
.cart__container::-webkit-scrollbar {
  display: none;
}
.cart__container{
  max-height: 500px;
}
</style>