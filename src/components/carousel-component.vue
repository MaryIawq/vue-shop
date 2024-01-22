<script setup>
import {Carousel, Navigation, Pagination, Slide} from 'vue3-carousel'
import 'vue3-carousel/dist/carousel.css'

const name = 'Autoplay';
const components = {
  Carousel,
  Slide,
  Pagination,
  Navigation,
}

const pics = [
  '../../public/carousel/carousel_straw_cake_for_new.png',
  '../../public/carousel/carousel_15_700.png',

]
</script>

<template>
  <div class="carousel__container">
    <Carousel :autoplay="10000" :wrap-around="true" class="w-full">
      <Slide v-for="slide in pics" :key="slide">
        <img :src=slide
             class="w-full"
             alt=""/>
      </Slide>
      <template #addons>
        <Pagination/>

      </template>
    </Carousel>
  </div>

</template>

<style scoped>

</style>