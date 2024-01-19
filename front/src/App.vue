<script setup>
import {onMounted, ref, watch, reactive, provide} from "vue";
import axios from "axios";
import headerComponent from './components/header-component.vue'
import cardListComponent from './components/card-list-component.vue'
import drawerComponent from './components/drawer-component.vue'
import categoriesComponent from './components/categories-component.vue'

const items = ref([]);
let categoriesItems = ref([]);

const filters = reactive({
  sortBy: 'title',
  searchQuery: '',
});

const onChangeSelect = event => {
  filters.sortBy = event.target.value
}
const onChangeSearchInput = event => {
  filters.searchQuery = event.target.value
}

const fetchFavorites = async () => {
  try {
    const {data: favorites} = await axios.get(`https://af3c46b46dc5a452.mokky.dev/favorites`)
    items.value = items.value.map(item => {
      const favorite = favorites.find(favorite => favorite.parentId === item.id);

      if (!favorite) {
        return item;
      }
      return {
        ...item,
        isFavorite: true,
        favoriteId: favorite.id,
      };
    });
  } catch (err) {
    console.log(err);
  }
}

const addToFavorite = async (item) => {
  item.isFavorite = true;
}

const fetchItems = async () => {
  try {
    const params = {
      sortBy: filters.sortBy,
    };
    if (filters.searchQuery) {
      params.title = `*${encodeURI(filters.searchQuery)}*`;
    }

    let response;
    if (params.sortBy === 'isSugarFree=1') {
      const url = 'https://af3c46b46dc5a452.mokky.dev/items?isSugarFree=1';
      response = await axios.get(url);
    } else {
      const url = 'https://af3c46b46dc5a452.mokky.dev/items';
      response = await axios.get(url, {params: params});
    }

    items.value = response.data.map((obj) => ({
      ...obj,
      isFavorite: false,
      isAdded: false,
    }));
  } catch (err) {
    console.error(err);
  }
};

const updateItems = async (newItems) => {
  items.value = newItems;
}

onMounted(async () => {
  await fetchItems();
  await fetchFavorites();
});

onMounted(async () => {
  try {
    const {data} = await axios.get('https://af3c46b46dc5a452.mokky.dev/categories')
    categoriesItems.value = data;
  } catch (err) {
    console.log(err)
  }

})
watch(filters, fetchItems);

provide('addToFavorite', addToFavorite)

</script>


<template>
  <div class="w-5/6 sm:w-11/12 m-auto bg-white rounded-2xl shadow-2xl mt-14">
    <header-component></header-component>
<!--    <drawer-component></drawer-component>-->
    <div class="p-10 main">
      <div class="flex justify-between items-center" id="filterMenu">
        <h2 class="text-xl font-bold text-slate-500 mb-8 mx-5">delivery every day 11 am to 9 pm</h2>
        <div class="flex gap-5 mb-7" id="inputAndSelect">

          <select
              @change="onChangeSelect"
              class="py-1.5 px-3 border rounded-md outline-none text-slate-500 font-bold">
            <option value="title">by name</option>
            <option value="price">ascending in price</option>
            <option value="-price">descending in price</option>
            <option value="isSugarFree=1">sugar free</option>
          </select>
          <div class="relative">
            <img
                src="/search.ico"
                class="absolute left-3 top-2 w-6 h-6"
                alt="search"/>
            <input
                @input="onChangeSearchInput"
                placeholder="search..."
                class="border rounded-md py-1.5 pl-12 pr-4 outline-none focus:border-slate-400"/>
          </div>
        </div>
      </div>

      <categories-component
          :categoriesItems="categoriesItems"
          @updateItems="updateItems"
      ></categories-component>
      <div class="mt-8">
        <card-list-component :items="items"></card-list-component>
      </div>

    </div>

  </div>

</template>


<style scoped>


@media (max-width: 768px) {
  #inputAndSelect {
    flex-direction: column;
    margin-bottom: 20px;
  }

  #inputAndSelect select {
    width: 100%;
    margin-bottom: 1rem;
  }
}

@media (max-width: 600px) {
  #filterMenu {
    flex-direction: column;
    display: flex;

  }

}
</style>
