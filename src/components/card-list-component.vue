<script setup>
import cardComponent from './card-component.vue'
defineProps({
  items: Array
})
const onClickAdd = () => {
  alert('Add!!!')
}
</script>

<template>
  <div class="grid grid-cols-4 gap-4" id="cards">
    <card-component
        v-for="item in items"
        :key="item.id"
        :title="item.title"
        :image-url="item.imageUrl"
        :price="item.price"
        :weight="item.weight"
        :on-click-add="onClickAdd"></card-component>
  </div>
</template>

<style scoped>
@media (max-width:1024px) {
  #cards {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }
}
@media (max-width:768px) {
  #cards {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}
@media (max-width:425px) {
  #cards {
    display: grid;
    grid-template-columns: repeat(1, minmax(0, 1fr));
  }
}
</style>