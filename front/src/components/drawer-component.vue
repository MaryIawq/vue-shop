<script setup>
import cartListComponent from './cart-list-component.vue'
import infoBlock from "@/components/infoblock.vue";

const emit = defineEmits(['closeDrawer, create-order']);
const props = defineProps({
  finalPrice: Number,
  totalPrice: Number,
  deliveryPrice: Number,
  discountPrice: Number,
  cartButtonDisabled: Boolean
})


let confirmation_token = ''; 
const checkout = async () => {     
  console.log("init checkout")
  const uuid = crypto.randomUUID()
  console.log("init UUID: " + uuid)

  fetch('http://127.0.0.1:8080/payment/getToken', {
    method: "POST",
    body: JSON.stringify({
      price: props.finalPrice,
      idempotence_key: uuid,
    })
  })
  .then(response => response.json())
  .then(data => {
    confirmation_token = data
    console.log("about to render form")
    renderPayForm()
  })
}

const renderPayForm = async () => {
  const checkout = new window.YooMoneyCheckoutWidget({

  confirmation_token: confirmation_token, 
    error_callback: function(error) {
      console.log(error)
    }
  });
  checkout.on('success', () => {
    checkout.destroy();
  });

  checkout.on('fail', () => {  
    checkout.destroy();
  });

  //Отображение платежной формы в контейнере
  checkout.render('payment-form');
}

  const script = document.createElement('script');
  script.src = 'https://yookassa.ru/checkout-widget/v1/checkout-widget.js';
  script.async = true;
  document.head.appendChild(script);
</script>

<template>
  <div @click="() => emit('closeDrawer')"
      class="fixed top-0 left-0 h-full w-full bg-slate-900 z-10 opacity-60">
  </div>
  <div class="bg-sky-50 shadow-2xl w-full h-full fixed top-0 z-20 p-4 md:w-4/12 md:right-0 md:rounded-l-xl md:p-10"
       id="shoppingCartBlock">
    <div class="drawer__head flex items-center gap-5 mb-5">
      <img class="w-9 h-9 opacity-70 cursor-pointer rotate-180 transition hover:opacity-90 hover:-translate-x-2"
           src="/right-arrow.ico"
           @click="() => emit('closeDrawer')"
           alt="right arrow">
      <h1 class="text-slate-600 text-2xl text-center font-bold">shopping cart</h1>
    </div>
    <div v-if="!totalPrice"
        class="h-full flex items-center">
      <infoBlock
          class="mb-36"
          @close-drawer="() => emit('closeDrawer')"
          title="your shopping cart is empty"
          description="go to store"
          image-url="/sad-smile.png"></infoBlock>
    </div>

    <cart-list-component v-if="totalPrice"></cart-list-component>

    <div
        v-if="totalPrice"
        class="costing__container flex flex-col gap-4 mt-7">
      <div class="flex gap-2 text-slate-500 font-bold uppercase mb-3">
        <span>Promocode:</span>
        <input type="text"
               class="w-28 mb-2 text-green-700 border uppercase rounded-2xl outline-none focus:border-slate-400"/>
        <img class="w-7 h-7 cursor-pointer transition opacity-75 hover:opacity-100" src="/apply.png" alt="apply"/>
      </div>
      <div class="flex gap-2 text-slate-600">
        <span>Order amount:</span>
        <div class="flex-1 border-b border-dashed"></div>
        <span>{{ totalPrice }}₽</span>
      </div>
      <div class="flex gap-2 text-slate-600">
        <span>Discount:</span>
        <div class="flex-1 border-b border-dashed"></div>
        <span>{{ discountPrice }}₽</span>
      </div>
      <div class="flex gap-2 text-slate-600">
        <span>Delivery:</span>
        <div class="flex-1 border-b border-dashed"></div>
        <span>{{ deliveryPrice }}₽</span>
      </div>
      <div class="flex gap-2 font-bold text-slate-600">
        <span>Total:</span>
        <div class="flex-1 border-b border-dashed"></div>
        <span>{{ finalPrice }}₽</span>
      </div>
      <button
          @click="checkout()"
          :disabled="cartButtonDisabled"
          class="bg-lime-300 w-full disabled:bg-slate-300 rounded-2xl py-3 mt-5 transition font-bold text-slate-700 cursor-pointer hover:bg-lime-400 active:bg-lime-500">
        place an order
      </button>
      <div id="payment-form"></div>
    </div>

  </div>
</template>

<style scoped>

@media (max-width: 900px) {
  #shoppingCartBlock {
    right: 0;
    padding: 4px;
    width: 65%;
    height: 100%;
    border-radius: 20px 0px 0px 20px / 30px 0px 0px 20px;
  }

  .drawer__head {
    margin-left: 15px;
    margin-top: 15px;
  }

  .costing__container {
    margin-left: 20px;
    margin-right: 20px;
  }
}

@media (max-width: 425px) {
  #shoppingCartBlock {
    padding: 10px;
    width: 100%;
    height: fit-content;
    border-radius: 0px 0px 20px 20px / 0px 0px 20px 20px;
  }

  .costing__container {
    margin-bottom: 20px;
  }
}
</style>