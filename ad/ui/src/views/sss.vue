<template>
  <v-container>
    <v-row>

      <v-col
        cols="12"
        sm="6"
        md="4"
      >
        <v-text-field
          v-model="route.route"
          label="Маршрут"
        ></v-text-field>
      </v-col>
      <v-col
        cols="12"
        sm="6"
        md="4"
      >
        <v-text-field
          v-model="route.method"
          label="Метод"
        ></v-text-field>
      </v-col>
      <v-col
        cols="12"
        sm="6"
        md="4"
      >
        <v-text-field
          v-model="route.description"
          label="Описание"
        ></v-text-field>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'sss',
};
</script>

<style scoped>

</style>
