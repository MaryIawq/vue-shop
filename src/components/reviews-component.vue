<script setup>
import {onMounted, ref} from "vue";
import axios from "axios";

const reviews = ref([])
const fetchData = async() => {
  try {
    const resp = await axios.get('https://af3c46b46dc5a452.mokky.dev/reviews');
    reviews.value = resp.data;
  }catch (err) {
    console.error(err)
  }
}
const getRatingColor = (rating) => {
  if (rating >= 4) {
    return 'green';
  } else if (rating >= 3) {
    return 'orange';
  } else {
    return 'red';
  }
};

onMounted(() => {
  fetchData()
})
</script>

<template>
  <div class="flex flex-col items-center justify-center">
    <h1 class="text-center text-3xl font-bold text-neutral-600">Reviews and comments</h1>
    <button class="border text-neutral-600 font-bold rounded-2xl p-2 mt-2 w-fit">manage (visible only for the
      administrator)
    </button>
  </div>
<div class="reviews__container grid-cols-4 grid mt-8  gap-4">
  <div v-for="review in reviews" class="p-2 border rounded-2xl">
    <div class="flex text-xl  text-neutral-600 items-center gap-1">
      <img class="w-16 h-16 border-2 rounded-full" :src="review.imageUrl" alt="users photo"/>
      <p class="font-bold">{{review.name}},</p>
      <p >{{review.city}}</p>
    </div>
    <div class="">
      <p class="text-neutral-600 mt-2 font-bold border-b mb-2">{{review.data}}</p>
      <p>{{review.text}}</p>
      <div class="flex items-center font-bold text-neutral-600 gap-1 mt-4">
        <p>rating:</p>
        <p :style="{ color: getRatingColor(review.rating) }" class="text-xl">{{review.rating}}</p>
        <img v-for="star in 5"
             class="w-6 h-6 opacity-75"
             :src="star <= review.rating ? '/rating-star.svg' : '/rating-star-filled.svg'"
             alt="Rating star"
             :style="{ display: star > review.rating ? 'none' : 'inline' }" />
      </div>
    </div>
  </div>
</div>
</template>

<style scoped>
@media (max-width:1700px) {
  .reviews__container {
    display: grid;
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }
}
@media (max-width:1284px) {
  .reviews__container {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }
}
@media (max-width:884px) {
  .reviews__container {
    display: grid;
    grid-template-columns: repeat(1, minmax(0, 1fr));
  }
}
</style>