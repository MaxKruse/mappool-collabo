<script setup lang="ts">

import { onMounted, watch } from 'vue';
import { useDefaultStore } from './store'
import { getSelf } from "./compositions/useUser";
import { useRouter } from 'vue-router';
import { User } from './models/User';

import Navbar from './components/Navbar.vue';

 
const store = useDefaultStore()
const router = useRouter();

const isUserLoggedIn = () => {
    return store.user !== null && store.user !== undefined && store.user.avatar_url !== "";
}

onMounted( async () => {
  // get the user from backend
  let user: User | null = null;
  
  try {
    user = await getSelf();
  } catch (e) {
    console.log(e);
    router.push('/login');
  }

  // if the user is null, redirect to login page
  if (!user) {
    return;
  }
  
  // set the user to store
  store.setUser(user);
})
</script>

<template>
  <v-app>
    <v-main>
      <navbar v-if="isUserLoggedIn()"/>
      <router-view/>
    </v-main>
  </v-app>
</template>

<style scoped>
</style>
