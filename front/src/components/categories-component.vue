<script setup>
import {defineEmits, ref} from 'vue';
import axios from 'axios';

const emits = defineEmits();
const props = defineProps({
  categoriesItems: Array,
})
let currentCategoryIdx = ref(-1);
const fetchCategories = async (idx) => {
  console.log(props)
  console.log(props.categoriesItems)
  console.log(props.categoriesItems[idx])
  console.log(props.categoriesItems[idx].title)
  try {
    let url = 'https://af3c46b46dc5a452.mokky.dev/items?category=' + props.categoriesItems[idx].title
    const response = await axios.get(url);
    console.log(response.data);
    emits('updateItems', response.data);
    currentCategoryIdx.value = idx;
  } catch (error) {
    console.log(error);
  }
};

</script>

<template>
  <div class="border rounded-2xl p-2 flex flex-wrap gap-5 justify-around text-slate-600 hover:text-slate-900">
    <button v-for="(categoryItem, idx) in categoriesItems"
            @click="fetchCategories(idx)"
            :class="{
              'flex gap-1 border border-slate-200 rounded-2xl transition p-1.5 active:bg-orange-100 hover:shadow-md hover:-translate-y-1': true,
              'bg-orange-100': idx === currentCategoryIdx
            }"
            :key="idx"
    >
      <span>{{ categoryItem.title }}</span>
      <img :src="categoryItem.imageUrl" alt="bake" class="w-5 h-5">
    </button>
  </div>
</template>

<style scoped>
</style>