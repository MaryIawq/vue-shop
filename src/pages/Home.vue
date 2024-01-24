<script setup>
import cardListComponent from '../components/card-list-component.vue'
import categoriesComponent from '../components/categories-component.vue'
import carouselComponent from '../components/carousel-component.vue'
import axios from "axios";
import {inject, onMounted, provide, reactive, ref, watch} from "vue";

/*ADD TO CART START*/
const items = ref([]);

const {cart, addToCart, removeFromCart} = inject('cart')
const onClickAddPlus = (item) => {
  if (!item.isAdded) {
    addToCart(item)
  } else {
    removeFromCart(item)
  }
}

onMounted(async () => {
  const localCart = localStorage.getItem('cart');
  cart.value = localCart ? JSON.parse(localCart) : []

  await fetchItems();
  await fetchFavorites();

  items.value = items.value.map((item) => ({
    ...item,
    isAdded: cart.value.some((cartItem) => cartItem.id === item.id)
  }))
});

watch(cart, () => {
  items.value = items.value.map((item) => ({
    ...item,
    isAdded: false
  }))
})
/*ADD TO CART END*/

/*FILTERS START*/
const filters = reactive({
  sortBy: 'title',
  searchQuery: '',
});
const onChangeSelect = event => {
  filters.sortBy = event.target.value
  fetchFavorites()
}
const onChangeSearchInput = event => {
  filters.searchQuery = event.target.value
}
/*FILTERS END*/

/*FAVORITES AND ITEMS START*/
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
const addToFavorite = async (item) => {
  try {
    if (!item.isFavorite) {
      const obj = {
        item_id: item.id
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

const fetchFavorites = async () => {
  try {
    const {data: favorites} = await axios.get(`https://af3c46b46dc5a452.mokky.dev/favorites`)
    items.value = items.value.map(item => {
      const favorite = favorites.find(favorite => favorite.item_id === item.id);

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

watch(filters, fetchItems);


/*FAVORITES AND ITEMS END*/

/*CATEGORIES START*/
let categoriesItems = ref([]);

onMounted(async () => {
  try {
    const {data} = await axios.get('https://af3c46b46dc5a452.mokky.dev/categories')
    categoriesItems.value = data;
  } catch (err) {
    console.log(err)
  }
})
/*CATEGORIES END*/
</script>

<template>
  <div class="content__block flex gap-3 mb-4">
    <carousel-component class="carousel"></carousel-component>
    <video class="video-fluid z-depth-1 rounded-2xl" loop autoplay muted playsinline>
      <source src="/public/carousel/video.mp4" type="video/mp4">
      <p>Sorry, there's a problem playing this video. Please try using a different browser.</p>
    </video>
  </div>

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
        @add-to-cart="onClickAddPlus"
        @add-to-favorite="addToFavorite"></card-list-component>
  </div>
</template>

<style scoped>
.content__block > video {
  width: 48%;
  height: 20%;
}
.carousel {
  width: 52%;
  height: 20%;
}
@media (max-width: 780px) {
  .input__and__select {
    flex-direction: column;
    margin-bottom: 20px;
  }

  .content__block {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
  }

  .content__block > video {
    width: 100%;
    height: 50%;
  }

  .carousel {
    width: 100%;
    height: 50%;
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

@media (max-width: 340px) {
  input {
    width: 95%;
  }

  select {
    width: 95%;
    margin-bottom: 1rem;
  }
}
</style>