<script setup>
import { ref } from "vue";
defineProps({
  title: String,
  imageUrl: String,
  weight: Number,
  price: Number,
  isFavorite: Boolean,
  isSugarFree: Boolean,
  isAdded: Boolean,
  protein: Number,
  fat: Number,
  carbs: Number,
  kkal: Number,
  onClickAdd: Function,
  onClickFavorite: Function,
})

let showEnergyValueTooltip = ref(false)
function toggleEnergyValueTooltip () {
  showEnergyValueTooltip.value = !showEnergyValueTooltip.value
}
let showSugarFreeTooltip = ref(false)
function toggleSugarFreeTooltip () {
  showSugarFreeTooltip.value = !showSugarFreeTooltip.value
}
</script>

<template>
  <div
      class="relative bg-white border border-slate-100 rounded-3xl p-8 cursor-pointer transition hover:-translate-y-1.5 hover:shadow-xl flex flex-col">

    <div class="relative">
      <img :src="imageUrl" alt="product" class="w-72 h-auto">
      <img
          src="/sugar-free.png"
          @mouseover="toggleSugarFreeTooltip"
          @mouseleave="toggleSugarFreeTooltip"
          v-if="isSugarFree"
          alt="sugar free"
          class="w-9 border h-auto bottom-0 left-0 opacity-95  bg-white absolute rounded-2xl p-2 m-1">
      <div
          class="absolute ml-8 z-30 rounded-r-2xl rounded-b-2xl bg-slate-200 text-slate-700 font-thin p-2 shadow-2xl"
          v-if="showSugarFreeTooltip"><p>sugar free</p>
      </div>
      <img
          @click="onClickFavorite"
          :src="!isFavorite ? '/favorites.ico' : '/added-to-fav.ico'"
          alt="is favorite"
          class="w-9 bg-white p-1.5 border rounded-2xl h-auto top-0 right-0 mr-1 mt-1 absolute opacity-85 hover:opacity-100">
    </div>

    <div class="flex justify-between">

      <p class="text-slate-600 uppercase mt-1 text-xl mb-5">{{ title }} </p>
      <div class="flex flex-wrap justify-evenly">
        <div class="">
          <img src="/info.ico"
               @mouseover="toggleEnergyValueTooltip"
               @mouseleave="toggleEnergyValueTooltip"
               alt="info"
               class="absolute transition mt-1 w-6 h-6 opacity-60 hover:rotate-45 hover:opacity-90">
        </div>
        <div v-if="showEnergyValueTooltip"
            class="absolute z-30 rounded-l-2xl rounded-b-2xl bg-slate-200 mr-16 text-slate-700 font-thin p-2 mt-10 shadow-2xl">
           <p>protein:{{ protein }}</p><p>fat:{{ fat }}</p><p>carbs:{{ carbs }}</p><p>kkal:{{ kkal }}</p>
        </div>
      </div>
    </div>

    <div class="flex flex-col mt-auto">
      <div class="flex justify-between">
        <div class="flex flex-col ml-1">
          <span class="text-slate-600 mb-1">{{ weight }} g</span>
          <b class="text-slate-600">{{ price }} ₽</b>
        </div>
        <img
            @click="onClickAdd"
            :src="!isAdded ? '/plus-add-to-cart.ico' : '/tick.ico'"
            alt="is added"
            class="w-8 h-8 z-20 opacity-75 hover:opacity-100">
      </div>
    </div>
  </div>
</template>

<style scoped>


</style>