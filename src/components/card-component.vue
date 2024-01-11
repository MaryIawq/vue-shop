<script setup>
defineProps({
  title: String,
  imageUrl: String,
  weight: Number,
  price: Number,
  isFavorite: Boolean,
  isAdded: Boolean,
})
</script>

<template>
  <div class="relative bg-white border border-slate-100 rounded-3xl p-8 cursor-pointer transition hover:-translate-y-1.5 hover:shadow-xl flex flex-col">
    <img src="/added-to-fav.ico" alt="added to favorites" class="absolute top-8 left-8 max-w-9">

    <img :src="imageUrl" alt="orange chocolate" class="w-72">

    <p class="text-slate-600 uppercase mt-1 text-xl mb-5">{{ title }}</p>

    <div class="flex flex-col mt-auto">
      <div class="flex justify-between">
        <div class="flex flex-col ml-1">
          <span class="text-slate-600 mb-1">{{ weight }} g</span>
          <b class="text-slate-600">{{ price }} ₽</b>
        </div>
        <img :src="isAdded ? '/plus.svg' : '/tick.ico'" alt="plus" class="max-w-8 max-h-8">
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>