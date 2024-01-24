<script setup>

</script>

<template>
  <div class="flex flex-col items-center justify-center">
    <h1 class="text-4xl font-bold text-neutral-500 text-center mb-10">please login or register to continue shopping</h1>
    <div class="main flex shadow-xl flex-col w-96 border bg-neutral-50 justify-center rounded-2xl  p-4 items-center">
      <h2 class="uppercase text-3xl font-bold text-neutral-600">register</h2>
      <div class="register__container gap-5 flex flex-col items-center justify-center border bg-white rounded-2xl mt-6 p-14">
        <div class="flex flex-col mb-2 gap-2">
          <div class="flex gap-2">
            <img src="/name.png" alt="name" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">name</p>
          </div>
          <input class="border rounded-md" placeholder="enter your name" type="text"/>
        </div>
        <div class="flex flex-col mb-2 gap-2">
          <div class="flex gap-2">
            <img src="/phone.png" alt="phone" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">phone number</p>
          </div>
          <input class="border rounded-md" placeholder="enter your phone num" type="number"/>
        </div>

        <div class="flex flex-col mb-2 gap-2">
          <div class="flex gap-2">
            <img src="/email.png" alt="email" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">email</p>
          </div>
          <input class="border rounded-md" placeholder="enter your email" type="email"/>
        </div>
        <div class="flex flex-col mt-4 gap-5">
          <div class="flex gap-2">
            <img src="/password.png" alt="password" class="w-6 h-6 opacity-75"/>
            <p class="text-neutral-600 font-semibold">password</p>
          </div>
          <p class="text-neutral-500 font-semibold">create a password</p>
          <input class="border rounded-md" placeholder="enter your password" type="password"/>
          <p class="text-neutral-500 font-semibold">confirm password</p>
          <input class="border rounded-md" placeholder="confirm your password" type="password"/>
        </div>
        <button class="transition mt-10 font-bold text-neutral-600 border bg-white p-2 px-8 rounded-2xl hover:-translate-y-0.5 hover:shadow-lg">Register</button>
      </div>
      <div class="flex flex-col items-center justify-center border bg-white rounded-2xl mt-6 p-8">

        <p class="text-neutral-600 text-center font-semibold">do you have a account?</p>

          <button class="transition mt-3 font-bold text-neutral-600 border bg-white p-2 px-8 rounded-2xl hover:-translate-y-0.5 hover:shadow-lg">log in</button>

      </div>
    </div>
  </div>
</template>

<style scoped>
@media (max-width: 500px) {
  .main {
    width: 300px
  }
  .register__container {
    width: 280px
  }
}
@media (max-width: 350px) {
  .main {
    width: 260px
  }
  .register__container {
    width: 250px
  }
}
</style>