<script setup lang="ts">

import { onBeforeMount, onMounted, ref, Suspense, watch } from 'vue';
import { useDefaultStore } from './store'
import { getSelf } from "./compositions/useUser";
import { useRouter, useRoute } from 'vue-router';
import { User } from './models/User';

import Navbar from './components/Navbar.vue';
import Login from './views/Login.vue';

const isUserLoggedIn = () => store.user !== null;
const store = useDefaultStore();

const route = useRoute()
const router = useRouter();

onMounted(async () => {
  try {
    const user = await getSelf();
    if (user) {
      store.setUser(user);
    }
  } catch (e) {
    console.log(e);
  }
})

</script>

<template>
  <navbar/>
  <main>
    <router-view v-if="isUserLoggedIn()"/>
    <Login v-else/>
  </main>
</template>

<style scoped>
/* main should take up a maximum of 60% width, be centered, yet not center the route-view */
main {
  max-width: 60%;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 64px;
}
</style>
