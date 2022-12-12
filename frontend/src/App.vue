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
  <navbar class="bg-grey-lighten-1"/>
  <main>
    <router-view v-if="isUserLoggedIn()"/>
    <Login v-else/>
  </main>
</template>

<style scoped>
</style>
