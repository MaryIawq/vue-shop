<script>

</script>

<template>

  <h1 class="text-center text-3xl font-bold mb-6 text-neutral-600">About company</h1>
  <div class="img flex flex-col rounded-3xl p-20 text-stone-700 gap-16 overflow-auto">
    <div class="about__block flex flex-col gap-2 rounded-3xl p-4">
      <h2 class="font-bold  text-center text-3xl">About us:</h2>
      <p>"Lakomka - Sweet Workshop" is a unique company specializing in the production of high quality sweets that
        delight
        and inspire our customers since its foundation in 1998. We are proud that our products are not only delicious,
        but
        also bring joy and satisfaction to everyone who tastes them.</p>
    </div>
    <div class="about__block flex flex-col gap-2 rounded-3xl p-4">
      <h2 class="font-bold text-center text-3xl">Our products:</h2>
      <p>We offer a wide range of sweets including candies, chocolates, marmalade, marshmallows, cookies, cakes and
        more.
        Our products are made only from selected natural ingredients without artificial additives and preservatives,
        which
        makes them not only tasty but also safe for health.</p>
    </div>
    <div class="about__block flex flex-col gap-2 rounded-3xl p-4">
      <h2 class="font-bold text-center text-3xl">Our Mission:</h2>
      <p>We strive to make the world a little bit sweeter every day, one treated customer after another. Our mission is
        to delight people with delicious and quality sweets, creating magical moments and unforgettable experiences.</p>
    </div>
  </div>
</template>

<style scoped>
.img {
  background-image: url('/white-background.jpg');
  background-attachment: fixed;
  background-size: cover;
  background-position: center;
  height: 80vh;
}

@keyframes scroll {
  0% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(-100%);
  }
}

.img::-webkit-scrollbar {
  display: none;
}

.about__block {
  background-color: rgba(255, 255, 255, 0.5);
  animation: scroll 30s linear infinite;
}
</style>