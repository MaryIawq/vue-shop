<script setup>
import OrdersComponent from "@/components/orders-component.vue";

</script>

<template>
  <div class="flex justify-between">
    <h1 class="page__header uppercase font-extrabold text-stone-600 pb-4 text-3xl">My profile</h1>
    <router-link to="/">
      <button class="home__btn flex opacity-60 font-extrabold cursor-pointer text-2xl uppercase hover:opacity-80">
        <img class="w-8"
             src="/public/back.ico"
             alt=""/>Home
      </button>
    </router-link>
  </div>
  <div class="main bg-stone-100 rounded-2xl flex">
    <div class="left__info__block w-3/12 bg-neutral-200 rounded-2xl my-10 ml-7">
      <div class="bg-white gap-3 grid justify-items-start	 place-content-start m-7 p-2 rounded-2xl">
        <h2 class="font-bold text-center text-stone-700">info</h2>
        <button class="text-stone-600 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">for individuals</button>
        <button class="text-stone-600 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">for legal entities</button>
      </div>
      <div class="bg-white m-7 p-2 gap-3 grid justify-items-start rounded-2xl ">
        <button class="font-bold text-stone-700 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">how to cancel points</button>
      </div>
      <div class="bg-white m-7 p-2 gap-3 grid justify-items-start rounded-2xl ">
        <button class="font-bold text-stone-700 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">my delivery addresses</button>
      </div>
    </div>


    <div class="profile w-6/12 flex flex-col items-center gap-10 justify-center">
      <div
          class="profile__card flex flex-col relative items-center max-w-96 w-full bg-stone-50 shadow-2xl mt-10 rounded-2xl p-7">
        <div class="img__container mb-2.5 relative h-36 w-36 rounded-full border-stone-300 border-4 bg-stone-50">
          <img
              class="h-full w-full object-cover p-4"
              src="/10102alienmonster_109997.png"
              alt="profile image">
        </div>
        <div class="text__data gap-2 flex flex-col items-center">
          <span class="name text-2xl text-stone-700 font-bold capitalize">tessi</span>
          <span class="discount text-lg text-stone-500 font-semibold">your discount 5%</span>
        </div>
        <div class="media__btns gap-2 mt-8 flex items-center">
          <button
              class="h-11 w-11 border-2 border-stone-400 bg-neutral-200 rounded-3xl transition hover:brightness-110">
            <img class="p-1" src="/purchases.png" alt="purchases"></button>
          <button class="h-12 w-12 border-2 border-stone-400 bg-white rounded-3xl transition hover:brightness-110"><img
              class="p-1" src="/payment.png" alt="payment"></button>
          <button
              class="h-12 w-12 border-2 border-stone-400 bg-neutral-200 rounded-3xl transition hover:brightness-110">
            <img class="p-1" src="/points.png" alt="points"></button>
          <button class="h-12 w-12 border-2 border-stone-400 bg-white rounded-3xl transition hover:brightness-110"><img
              class="p-1" src="/requests.png" alt="requests"></button>
        </div>
        <div class="mt-6 flex items-center gap-6">
          <router-link to="/auth">
            <button class=" rounded-3xl border-2 border-stone-400 transition p-2 hover:-translate-y-0.5 hover:shadow-lg">
              exit
            </button>
          </router-link>
        </div>
      </div>
      <div class="info__profile">
        <orders-component></orders-component>
      </div>
    </div>


    <div class="right__info__block w-3/12 bg-neutral-200 rounded-2xl my-10 mr-7">
      <div class="bg-white m-7 p-2 gap-3 grid justify-items-start rounded-2xl ">
        <button class="font-bold text-stone-700 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">my promo codes</button>
      </div>
      <div class="bg-white m-7 p-2 gap-3 grid justify-items-start rounded-2xl ">
        <h2 class="font-bold text-center text-stone-700">edit profile</h2>
        <button class="text-stone-600 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">change name</button>
        <button class="text-stone-600 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">change my phone number</button>
        <button class="text-stone-600 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">change the photo</button>
        <button class="text-stone-600 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">delete profile</button>
      </div>
      <div class="bg-white m-7 p-2 gap-3 grid justify-items-start rounded-2xl ">
        <button class="font-bold text-stone-700 border p-1 rounded-2xl transition hover:-translate-y-0.5 hover:shadow-lg">select a city from the list</button>
      </div>

    </div>
  </div>
</template>

<style scoped>
@media (max-width: 960px) {
  .main {
    flex-direction: column;

  }

  .left__info__block,
  .profile,
  .right__info__block {
    width: 100%;
    margin: 0;
  }
  .profile__card,
  .info__profile{
    max-width: 90%;
  }
  .left__info__block {
    margin-top: 10px;
    order: 2;
  }
}
@media (max-width: 425px) {
  .profile {
    background-color: white;
  }
  .media__btns {
    flex-wrap: wrap;
    margin-left: 25px;
  }
  .text__data {
    margin-top: 25px
  }
  .page__header,
  .home__btn{
    font-size: 18px;
  }
}
.profile__card::before {
  content: '';
  position: absolute;
  height: 36%;
  width: 100%;
  background-image: url("/public/cookies-bg.jpg");
  background-size: cover;
  top: 0;
  left: 0;
  border-radius: 16px 16px 0 0;
}
</style>