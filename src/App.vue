<script setup>

import {onMounted, ref, watch, reactive, provide} from "vue";
import axios from "axios";
import headerComponent from './components/header-component.vue'
import cardListComponent from './components/card-list-component.vue'
import drawerComponent from './components/drawer-component.vue'
import categoriesComponent from './components/categories-component.vue'

const items = ref([]);
//drawer functional
const drawerOpen = ref(false)
const closeDrawer = () => {
  drawerOpen.value = false
}
const openDrawer = () => {
  drawerOpen.value = true
}
const cart = ref([])
const addToCart = (item) => {
  if (!item.isAdded) {
    cart.value.push(item)
    item.isAdded = true
  } else {
    cart.value.splice(cart.value.indexOf(item), 1)
    item.isAdded = false
  }
}
provide('cart', {
  cart
})

//favorites functional
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
  try {
    if (!item.isFavorite) {
      const obj = {
        parentId: item.id
      };
      item.isFavorite = true;
      const {data} = await axios.post(`https://af3c46b46dc5a452.mokky.dev/favorites`, obj);
      item.favoriteId = data.id;
    } else {
      item.isFavorite = false;
      await axios.delete(`https://af3c46b46dc5a452.mokky.dev/favorites/${item.favoriteId}`)
      item.favoriteId = null;
    }
  } catch (err) {
    console.log(err);
  }
}
const filters = reactive({
  sortBy: 'title',
  searchQuery: '',
});


//categories functional
let categoriesItems = ref([]);
onMounted(async () => {
  try {
    const {data} = await axios.get('https://af3c46b46dc5a452.mokky.dev/categories')
    categoriesItems.value = data;
  } catch (err) {
    console.log(err)
  }
})


//filters functional

const onChangeSelect = event => {
  filters.sortBy = event.target.value
  fetchFavorites()
}
const onChangeSearchInput = event => {
  filters.searchQuery = event.target.value
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
    const url = 'https://af3c46b46dc5a452.mokky.dev/items';
    response = await axios.get(url, {params: params});
    items.value = response.data.map((obj) => ({
      ...obj,
      isFavorite: false,
      favoriteId: null,
      isAdded: false,
    }));
  } catch (err) {
    console.error(err);
  }
  await fetchFavorites()
};
const updateItems = async (newItems) => {
  items.value = newItems;

  await fetchFavorites()
}
onMounted(async () => {
  await fetchItems();
  await fetchFavorites();
});
watch(filters, fetchItems);


</script>


<template>
  <div class="w-5/6 sm:w-11/12 m-auto bg-white rounded-2xl shadow-2xl mt-14">
    <header-component @open-drawer="openDrawer"></header-component>
    <drawer-component v-if="drawerOpen" @close-drawer="closeDrawer"></drawer-component>
    <div class="p-10 main">
      <div class="filter__menu flex justify-between items-center">
        <div class="input__and__select flex gap-5 mb-7">

          <select
              @change="onChangeSelect"
              class="py-1.5 px-3 border rounded-md outline-none text-slate-500 font-bold cursor-pointer">
            <option value="title">by name</option>
            <option value="price">ascending in price</option>
            <option value="-price">descending in price</option>
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
          :categories-items="categoriesItems"
          @update-items="updateItems"
          @fetch-items="fetchItems"
      ></categories-component>
      <div class="mt-8">
        <card-list-component
            :items="items"
            @add-to-cart="addToCart"
            @add-to-favorite="addToFavorite"></card-list-component>
      </div>

    </div>

  </div>

</template>


<style scoped>


@media (max-width: 768px) {
  .input__and__select {
    flex-direction: column;
    margin-bottom: 20px;
  }

  select {
    width: 100%;
    margin-bottom: 1rem;
  }
}

@media (max-width: 600px) {
  .filter__menu {
    flex-direction: column;
    display: flex;

  }

}
</style>
