<script setup>
import headerComponent from './components/header-component.vue'
import cardListComponent from './components/card-list-component.vue'
import drawerComponent from './components/drawer-component.vue'
</script>

<template>


  <div class="w-5/6 sm:w-11/12 m-auto bg-white rounded-2xl shadow-2xl mt-14">
    <header-component></header-component>
    <!--    <drawer-component></drawer-component>-->
    <div class="p-10">
      <div class="flex justify-between items-center" id="filterMenu">
        <h2 class="text-3xl font-bold text-slate-600 mb-8">all sweets</h2>

        <div class="flex gap-2" id="inputAndSelect">
          <select class="py-2 px-3 border rounded-md outline-none text-slate-500 font-bold">
            <option>by name</option>
            <option>ascending in price</option>
            <option>descending in price</option>
          </select>
          <div class="relative">
            <img
                src="/search.ico"
                class="absolute left-3 top-2 w-6 h-6"
                alt="search"/>
            <input
                placeholder="search..."
                class="border rounded-md py-1.5 pl-12 pr-4 outline-none focus:border-slate-400"/>
          </div>
        </div>
      </div>

      <card-list-component></card-list-component>
    </div>

  </div>

</template>

<style scoped>
@media (max-width: 768px) {
  #inputAndSelect {
    flex-direction: column;
    margin-bottom: 38px;
  }

  #inputAndSelect select {
    width: 100%;
    margin-bottom: 1rem;
  }
}
@media (max-width: 425px) {
  #filterMenu {
    flex-direction: column;
    display: flex;

  }

}
</style>
