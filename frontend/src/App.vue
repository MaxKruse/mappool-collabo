<script setup lang="ts">

import { onMounted } from 'vue';
import { useDefaultStore } from './store'
import { getSelf } from "./compositions";
import { useRouter } from 'vue-router';

 
const store = useDefaultStore()
const router = useRouter();

onMounted( async () => {
  // get the user from backend
  const user = await getSelf();

  // if the user is null, redirect to login page
  if (!user) {
    router.push('/login');
    return;
  }
  
  // set the user to store
  store.setUser(user);
})

</script>

<template>
  <router-view></router-view>
</template>

<style scoped>
</style>
