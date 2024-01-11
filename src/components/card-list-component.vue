<script setup>
import cardComponent from './card-component.vue'
const onClickAdd = () => {
  alert('Add!!!')
}
</script>

<template>
  <div class="grid grid-cols-4 gap-4" id="cards">
    <card-component
        title="orange chocolate"
        image-url="/products/orange-chocolate.jpg"
        :price="120"
        :weight="140"
        :is-added="true"
        :is-favorite="true"
        :on-click-add="onClickAdd"></card-component>
    <card-component title="jelly beans set" image-url="/products/jelly-beans-set.jpg" :price="210" :weight="250"></card-component>
    <card-component></card-component>
    <card-component></card-component>
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