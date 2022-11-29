<script setup lang="ts">

import { onBeforeMount, onMounted, Suspense, watch } from 'vue';
import { useDefaultStore } from './store'
import { getSelf } from "./compositions/useUser";
import { useRouter, useRoute } from 'vue-router';
import { User } from './models/User';

import Navbar from './components/Navbar.vue';
import Login from './views/Login.vue';

 
const store = useDefaultStore()

</script>

<template>
  <v-app>
    <navbar class="bg-grey-lighten-1"/>
    <v-main>
      <router-view/>
    </v-main>
  <v-footer class="bg-grey-lighten-1">Footer</v-footer>
  </v-app>
</template>

<style scoped>
</style>
