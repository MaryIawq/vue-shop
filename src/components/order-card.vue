<script setup>

const props = defineProps({
  title: String,
  weight: Number,
  imageUrl: String
});

</script>

<template>
  <div class="order__card flex p-2 gap-2 transition border rounded-2xl mb-2 hover:-translate-y-0.5 hover:shadow-md">
    <img :src="imageUrl" class="w-20 h-20 rounded-xl" alt="order image">
    <div class="flex flex-col">
      <p>{{ props.title }}</p>
      <p>{{ props.weight }} g</p>
    </div>

  </div>
</template>

<style scoped>
@media (max-width:760px) {
  .order__card {
    flex-direction: column;
    align-items: center;
    text-align: center;
  }
}
</style>