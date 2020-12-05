<template>
  <v-data-table
    :headers="headers"
    :items="routes"
    sort-by="calories"
    class="elevation-1"
  >
    <template v-slot:no-data>
      Нет данных
    </template>

    <template v-slot:top>
      <v-toolbar
        flat
      >
        <v-toolbar-title>Маршруты</v-toolbar-title>
        <v-divider
          class="mx-4"
          inset
          vertical
        ></v-divider>
        <v-spacer></v-spacer>

          <v-btn
            color="primary"
            dark
            class="mb-2"
            @click="createRoute"
          >
            Новый маршрут
          </v-btn>

      </v-toolbar>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-icon
        small
        class="mr-2"
        @click="editItem(item)"
      >
        mdi-pencil
      </v-icon>
      <v-icon
        small
        @click="deleteItem(item)"
      >
        mdi-delete
      </v-icon>
    </template>

    <create-route-dialog
      v-model='showCreateDialog'
    />
  </v-data-table>

</template>

<script lang="ts">
import {
  Component, Prop, Vue, ProvideReactive,
} from 'vue-property-decorator';

@Component({
  components: {
    'create-route-dialog': () => import('./create-route-dialog.vue'),
  },
  mounted() {
    this.$store.direct.dispatch.route.GetList();
  },
})
export default class RoutesTab extends Vue {
  showCreateDialog = false

  valid = true

  methods = ['GET', 'POST', 'PUT', 'DELETE']

  dialogDelete = false

  editedIndex = -1

  readonly headers = [
    {
      text: 'Маршруты',
      align: 'start',
      sortable: false,
      value: 'route',
    },
    {
      text: 'Метод',
      align: 'start',
      sortable: false,
      value: 'method',
    },
    {
      text: 'Описание',
      align: 'start',
      sortable: false,
      value: 'description',
    },
  ]

  routes = []

  protected createRoute(): void {
    this.showCreateDialog = true;
  }

  protected editItem(item: any): void {
    this.createRoute();
    console.log(item);
  }

  protected deleteItem(item: any): void {
    this.createRoute();
    console.log(item);
  }
}
</script>

<style scoped lang="scss">
.routes-tab-container {
  display: flex;
  height: 1000px;
  justify-content: space-between;
}
.create-route-btn {
  margin: 10px;

}
</style>
