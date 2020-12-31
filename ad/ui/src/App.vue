<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    />

    <v-main>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import MainTabs from '@/views/base/MainPage.vue';

export default Vue.extend({
  name: 'App',
});
</script>

<style scoped lang="css">

v-main {
  height: 100%;
  display: flex;
  flex-direction: column;
}
</style>
