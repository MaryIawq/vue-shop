<script setup>
import cardComponent from './card-component.vue'
import {defineEmits} from "vue";

defineProps({
  items: Array,
})
const emit = defineEmits(['addToFavorite', 'addToCart']);


</script>

<template>
  <div class="cards grid grid-cols-4 gap-4">
    <card-component
        v-for="item in items"
        :key="item.id"
        :id="item.id"
        :title="item.title"
        :image-url="item.imageUrl"
        :price="item.price"
        :weight="item.weight"
        :is-sugar-free="item.isSugarFree"
        :is-favorite="item.isFavorite"
        :is-added="item.isAdded"
        :protein="item.protein"
        :fat="item.fat"
        :carbs="item.carbs"
        :kkal="item.kkal"
        :on-click-favorite="() => emit('addToFavorite', item)"
        :on-click-add="() => emit('addToCart', item)">
    </card-component>
  </div>
</template>



<style scoped>
@media (max-width: 1024px) {
  .cards {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }
}

@media (max-width: 768px) {
  .cards {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}

@media (max-width: 425px) {
  .cards {
    display: grid;
    grid-template-columns: repeat(1, minmax(0, 1fr));
  }
}
</style>