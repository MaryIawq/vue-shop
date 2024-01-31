<script setup>
import axios from "axios";
import {onMounted, ref} from "vue";

const posts = ref([])
const fetchData = async() => {
  try {
    const resp = await axios.get('https://af3c46b46dc5a452.mokky.dev/news');
    posts.value = resp.data;
  }catch (err) {
    console.error(err)
  }
}
onMounted(() => {
  fetchData()
})
</script>

<template>
  <div class="flex flex-col items-center justify-center">
    <h1 class="text-center text-3xl font-bold text-neutral-600">News and promotions</h1>
    <button class="border text-neutral-600 font-bold rounded-2xl p-2 mt-2 w-fit">create post (visible only for the
      administrator)
    </button>
  </div>
  <div class="mt-6 border p-3 shadow-lg rounded-2xl gap-8 flex flex-col justify-center items-center" v-for="post in posts">
    <div class="flex text-xl text-stone-700 font-bold gap-3">
      <h2 class="text-orange-900">{{post.title}}</h2>
      <p >{{post.data}}</p>
    </div>
    <p >{{post.body}}</p>
  </div>
</template>

<style scoped>

</style>