<script setup>
defineProps({
  id: Number,
  title: String,
  imageUrl: String,
  weight: Number,
  price: Number,
  onClickFavorite: Function,
  onClickAdd: Function
})
const emit = defineEmits(['onClickRemove'])
</script>

<template>
  <div class="flex items-center gap-4 mb-3 border p-2 border-slate-200 transition rounded-2xl mt-3 shadow-lg hover:translate-x-0.5 hover:-translate-y-0.5">
    <img
        :src="imageUrl"
        class="w-24 h-24"
        alt="product"/>
    <div class="flex flex-col gap-1 ml-3">
      <p
      class="uppercase">{{ title }}</p>
      <b class="opacity-60">{{ weight }}g</b>
      <div class="flex justify-between">
        <b class="opacity-75">{{ price }}₽</b>

      </div>
      <img
          src="/close.ico"
          @click="emit('onClickRemove')"
          class="w-6 h-6 opacity-60 transition cursor-pointer hover:opacity-90"
          alt="close">
    </div>

  </div>
</template>

<style scoped>

</style>