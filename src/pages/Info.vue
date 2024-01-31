<script setup>
import {ref} from 'vue';
import {useRouter} from 'vue-router';
import NewsPromotionsComponent from "@/components/news-promotions-component.vue";
import ReviewsComponent from "@/components/reviews-component.vue";
import AboutComponent from "@/components/about-component.vue";


const router = useRouter();
const isOpen = ref(false);
const toggleMenu = () => {
  isOpen.value = !isOpen.value;
};
const scrollToSection = (sectionId) => {
  const sectionElement = document.getElementById(sectionId);
  if (sectionElement) {
    sectionElement.scrollIntoView({behavior: 'smooth'});
  }
};

</script>

<template>
  <div class="flex" id="menu">
    <div class="flex items-center gap-2 font-bold text-neutral-600 mb-14 p-2 border rounded-3xl">
      <button @click="toggleMenu" class="burger-button transition opacity-75 hover:opacity-100 hover:rotate-12">
        <img src="/burger-button.png" class="w-14 h-14" alt="Burger icon">
      </button>
      <p>Navigation menu</p>
    </div>
    <div class="bg-neutral-100 rounded-2xl absolute z-30">
      <div class="" v-if="isOpen">
        <div class="flex border-b border-stone-300 justify-between p-3">
          <h2 class="font-bold text-neutral-600 text-xl">Menu</h2>
          <button class="opacity-75 transition hover:opacity-100" @click="toggleMenu"><img class="w-7 h-7"
                                                                                           src="/close.ico" alt="close">
          </button>
        </div>
        <div class="flex flex-col p-4">
          <button
              class="p-1 border rounded-2xl mb-8 bg-white transition hover:-translate-y-0.5 hover:shadow-lg"
              @click="scrollToSection('about')">
            about the company
          </button>
          <button
              class="p-1 border rounded-2xl mb-8 bg-white transition hover:-translate-y-0.5 hover:shadow-lg"
              @click="scrollToSection('news')">
            news and promotions
          </button>
          <button
              class="p-1 border rounded-2xl mb-8 bg-white transition hover:-translate-y-0.5 hover:shadow-lg"
              @click="scrollToSection('reviews')">
            reviews and comments
          </button>
        </div>
      </div>
    </div>
  </div>

  <div class="content">
    <div id="about" class="mb-32">
      <about-component>
      </about-component>
    </div>
    <div id="news" class="">
      <news-promotions-component></news-promotions-component>
    </div>
    <div id="reviews" class="mt-32">
      <reviews-component></reviews-component>
    </div>
  </div>

  <button @click="scrollToSection('menu')"
          class="fixed right-4 bottom-4 bg-neutral-300 rounded-full transition py-2 px-2 shadow-lg hover:-translate-y-2">
    <img class="-rotate-90 w-10 h-10" src="/right-arrow.ico">
  </button>

</template>

<style scoped>
@media (min-width: 768px) {
  .content {
    margin-left: 40px;
    margin-right: 40px;
  }

}
</style>