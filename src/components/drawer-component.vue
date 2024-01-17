<script setup>
import cartListComponent from './cart-list-component.vue'
</script>

<template>

  <div class="fixed top-0 left-0 h-full w-full bg-slate-900 z-10 opacity-60">
  </div>
  <div class="bg-sky-50 shadow-2xl w-full h-full fixed top-0 z-20 p-4 md:w-4/12 md:right-0 md:rounded-l-xl md:p-10"
       id="shoppingCartBlock">
    <div class="flex items-center gap-5 mb-5">
      <img class="w-9 h-9 opacity-70 cursor-pointer rotate-180 transition hover:opacity-90 hover:-translate-x-2"
           src="/right-arrow.ico" alt="right arrow">
      <h1 class="text-slate-600 text-2xl text-center font-bold">shopping cart</h1>
    </div>
    <cart-list-component></cart-list-component>
    <div class="flex flex-col gap-4 mt-7">
      <div class="flex gap-2 text-slate-600">
        <span>Discount:</span>
        <div class="flex-1 border-b border-dashed"></div>
        <span>316₽</span>
      </div>
      <div class="flex gap-2 text-slate-600">
        <span>Delivery:</span>
        <div class="flex-1 border-b border-dashed"></div>
        <span>98₽</span>
      </div>
      <div class="flex gap-2 font-bold text-slate-600">
        <span>Total:</span>
        <div class="flex-1 border-b border-dashed"></div>
        <span>1132₽</span>
      </div>
      <button
          disabled=""
          class="bg-lime-300 w-full disabled:bg-slate-300 rounded-2xl py-3 mt-5 transition font-bold text-slate-700 cursor-pointer hover:bg-lime-400 active:bg-lime-500">
        place an order
      </button>
    </div>

  </div>
</template>

<style scoped>
@media (max-width: 768px) {
  #shoppingCartBlock {
    right: 0;
    padding: 4px;
    width: 65%;
    height: 100%;
  }
}

@media (max-width: 425px) {
  #shoppingCartBlock {
    padding: 10px;
    width: 100%;
    height: fit-content;
  }
}
</style>