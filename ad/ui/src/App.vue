<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >

    </v-app-bar>

    <v-main>
      <MainTabs/>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import MainTabs from '@/views/MainTabs.vue';

export default Vue.extend({
  name: 'App',

  components: {
    MainTabs,

  },

  data: () => ({
    //
  }),
});
</script>

<style scoped lang="css">
v-main {
  height: 100%;
}

v-main {
  display: flex;
  flex-direction: column;
}
</style>
