<script setup>
import { ref } from 'vue';

const phoneNumber = ref('');
const email = ref('');
const clickedNoMessage = ref(false);

const currentMethod = ref('emailMethod');

const showIncorrectPhone = ref(false);
const showIncorrectEmail = ref(false);

const showRerequestPhone = ref(false);
const showRerequestEmail = ref(false);

let countRerequestPhone = ref(60);
let countRerequestEmail = ref(60);

let blockBtnPhone = ref(false);
let blockBtnEmail = ref(false);

const showRecoveryByEmail = () => {
  currentMethod.value = 'emailMethod'
}
const showRecoveryByPhone = () => {
  currentMethod.value = 'phoneMethod'
}
const startCountdown = (method) => {
  if (method === 'phone') {
    blockBtnPhone.value = true;
    countRerequestPhone.value = 10;
    const timer = setInterval(() => {
      countRerequestPhone.value--;
      if (countRerequestPhone.value <= 0) {
        clearInterval(timer);
        blockBtnPhone.value = false;
      }
    }, 1000);
  } else if (method === 'email') {
    blockBtnEmail.value = true;
    countRerequestEmail.value = 10;
    const timer = setInterval(() => {
      countRerequestEmail.value--;
      if (countRerequestEmail.value <= 0) {
        clearInterval(timer);
        blockBtnEmail.value = false;
      }
    }, 1000);
  }
};

const checkPhoneNumber = () => {
  const phoneNumberRegex = /^\d{11}$/;
  showIncorrectPhone.value = !phoneNumberRegex.test(phoneNumber.value);
  if (!showIncorrectPhone.value) {
    showRerequestPhone.value = true;
    startCountdown('phone');
  }
};

const checkEmail = () => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  showIncorrectEmail.value = !emailRegex.test(email.value);
  if (!showIncorrectEmail.value) {
    showRerequestEmail.value = true;
    startCountdown('email');
  }
};


</script>


<template>
  <div class="flex flex-col items-center justify-center">
    <div class="main flex shadow-xl flex-col w-96 border bg-neutral-50 justify-center rounded-2xl  p-4 items-center">
      <h2 class="uppercase p-2 text-2xl font-bold text-neutral-600">Password Recovery</h2>

      <div v-if="currentMethod === 'phoneMethod'"
          class="phone__recovery gap-2 flex flex-col items-center justify-center border bg-white rounded-2xl p-14" :class="{ 'hidden': showRerequestEmail }">
        <div class="flex flex-col mb-2 gap-2">
          <p class="text-neutral-600">For password recovery, enter your phone number:</p>
          <input v-model="phoneNumber" class="border mt-2 rounded-md p-2" placeholder="Enter your phone number" type="number"/>
          <button @click="showRecoveryByEmail" class="text-neutral-900 opacity-75 hover:opacity-100">Recover by email</button>
        </div>
        <button @click="checkPhoneNumber" :class="{ 'disabled': blockBtnPhone }" :disabled="blockBtnPhone" class="transition mt-2 font-bold text-neutral-600 border bg-white p-2 px-8 rounded-2xl hover:-translate-y-0.5 hover:shadow-lg">Send code</button>
        <div class="incorrect" v-if="showIncorrectPhone">
          <p class="text-rose-400">Incorrect phone format</p>
        </div>
        <div class="rerequest flex flex-col items-center justify-center gap-2 mt-4" v-if="showRerequestPhone">
          <p class="text-green-700 font-bold text-xl">Code sent</p>
          <input type="number" class="border mt-2 w-40 rounded-md p-2" placeholder="enter code here"/>
          <p class="count text-neutral-600">You can re-request the code via: <span class="font-bold">{{ countRerequestPhone }}</span> sec</p>
        </div>
      </div>

      <div v-if="currentMethod === 'emailMethod'"
          class="email__recovery gap-2 flex flex-col items-center justify-center border bg-white rounded-2xl p-14" :class="{ 'hidden': showRerequestPhone }">
        <div class="flex flex-col mb-2 gap-2">
          <p class="text-neutral-600">For password recovery, enter your email:</p>
          <input v-model="email" class="border mt-2 rounded-md p-2" placeholder="Enter your email" type="email"/>
          <button @click="showRecoveryByPhone" class="text-neutral-900 opacity-75 hover:opacity-100">Recover by phone number</button>
        </div>
        <button @click="checkEmail" :class="{ 'disabled': blockBtnEmail }" :disabled="blockBtnEmail" class="transition mt-2 font-bold text-neutral-600 border bg-white p-2 px-8 rounded-2xl hover:-translate-y-0.5 hover:shadow-lg">Send email</button>
        <div class="incorrect" v-if="showIncorrectEmail">
          <p class="text-rose-400">Incorrect email format</p>
        </div>
        <div class="rerequest flex flex-col items-center justify-center gap-2 mt-4" v-if="showRerequestEmail">
          <p class="text-green-700 font-bold text-xl">An e-mail with further instructions has been sent to your e-mail account. </p>
          <p class="count text-neutral-600">You can re-request the email via: <span class="font-bold">{{ countRerequestEmail }}</span> sec</p>
          <button @click="clickedNoMessage=true" class="pt-2 opacity-75 hover:opacity-100">no message?</button>
        </div>

      </div>

    </div>

  </div>
  <div v-if="clickedNoMessage" class="absolute inset-0 flex items-center justify-center z-50">
    <div class="bg-white shadow-2xl flex flex-col justify-center items-center border-2 p-8 rounded-3xl">
      <ul class="text-stone-600 font-bold text-lg">
        <li class="mt-4">Search for emails in other folders. Email messages don't always end up in your Inbox.</li>
        <li class="mt-4">Cancel "spam" labeling for important emails</li>
        <li class="mt-4">Configure filtering and blocking settings</li>
        <li class="mt-4">Verify that you have free space in Google storage</li>
      </ul>
      <button @click="clickedNoMessage = false" class="mt-4 bg-stone-600 text-white py-2 px-4 rounded-2xl hover:brightness-110">ok</button>
    </div>-
  </div>
</template>
<style scoped>
.disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
@media (max-width: 500px) {
  .main {
    width: 300px
  }
  .phone__recovery,
  .email__recovery {
    width: 280px
  }
}
@media (max-width: 350px) {
  .main {
    width: 260px
  }
  .phone__recovery,
  .email__recovery {
    width: 250px
  }
}
</style>