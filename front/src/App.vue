<script setup>

import {onMounted, ref, watch, reactive, provide, computed} from "vue";
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
  cart.value.push(item)
  item.isAdded = true
}
const removeFromCart = (item) => {
  cart.value.splice(cart.value.indexOf(item), 1)
  item.isAdded = false
}

const onClickAddPlus = (item) => {
  if (!item.isAdded) {
    addToCart(item)
  } else {
    removeFromCart(item)
  }
}

const totalPrice = computed(() => cart.value.reduce((acc, item) => acc + item.price, 0))
const discountPrice = computed(() => Math.round((totalPrice.value * 9) / 100))
const deliveryPrice = computed(() => Math.round((totalPrice.value * 5) / 100))

const finalPrice = computed(() => {
  const total = totalPrice.value + deliveryPrice.value - discountPrice.value;
  return Math.max(0, total);
});

watch(cart, () => {
  items.value = items.value.map((item) => ({
    ...item,
    isAdded: false
  }))
})

provide('cart', {
  cart,
  addToCart,
  removeFromCart,
})
watch(cart, () => {
      localStorage.setItem('cart', JSON.stringify(cart.value))
    },
    {deep: true}
)

//orders fun

const cartButtonDisabled = computed(() =>
    isCreatingOrder.value ? true : totalPrice.value ? false : true
)
const createOrder = async () => {
  try {
    isCreatingOrder.value = true
    const {data} = await axios.post(`https://af3c46b46dc5a452.mokky.dev/orders`, {
      items: cart.value,
      totalPrice: totalPrice.value
    })

    cart.value = []

    return data
  } catch (e) {
    console.log(e)
  } finally {
    isCreatingOrder.value = false
  }
}
const isCreatingOrder = ref(false)

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
  const localCart = localStorage.getItem('cart');
  cart.value = localCart ? JSON.parse(localCart) : []


  await fetchItems();
  await fetchFavorites();

  items.value = items.value.map((item) => ({
    ...item,
    isAdded: cart.value.some((cartItem) => cartItem.id === item.id)
  }))
});
watch(filters, fetchItems);


</script>


<template>
  <div class="w-5/6 sm:w-11/12 m-auto bg-white rounded-2xl shadow-2xl mt-14">
    <header-component
        :final-price="finalPrice"
        :total-price="totalPrice"
        @open-drawer="openDrawer"></header-component>
    <div class="cart-container overflow-y-auto overflow-x-hidden">
      <drawer-component v-if="drawerOpen"
                        :final-price="finalPrice"
                        :total-price="totalPrice"
                        :delivery-price="deliveryPrice"
                        :discount-price="discountPrice"
                        @create-order="createOrder"
                        :cart-button-disabled="cartButtonDisabled"
                        @close-drawer="closeDrawer"></drawer-component>
    </div>

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
            @add-to-cart="onClickAddPlus"
            @add-to-favorite="addToFavorite"></card-list-component>
      </div>

    </div>

  </div>

</template>


<style scoped>
.cart-container {
  max-height: 60vh;
}

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
