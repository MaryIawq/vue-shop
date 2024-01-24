<script setup>

</script>

<template>


  <div class="order__card flex p-2 gap-2 transition border rounded-2xl mb-2 hover:-translate-y-0.5 hover:shadow-md">
    <img src="/products/christmas-sweets-set.jpg" class="w-20 h-20 rounded-xl" alt="order image">
    <div class="flex flex-col">
      <p>CHRISTMAS SWEETS SET</p>
      <p>800 g</p>
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