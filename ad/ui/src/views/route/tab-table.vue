<template>
  <div>
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
          />
          <v-spacer />
          <create-route-dialog />
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
    </v-data-table>
    <delete-route-dialog
      v-model="showDeleteDialog"
      :route-id="routeIdToDelete"
    />
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Route } from '@/services/route';

@Component({
  components: {
    'create-route-dialog': () => import('./create-dialog.vue'),
    'delete-route-dialog': () => import('./delete-dialog.vue'),
  },
  computed: {
    routes(): Route[] {
      return this.$store.direct.getters.route.getRoutes;
    },
  },
  mounted() {
    this.$store.direct.dispatch.route.GetList();
  },
})
export default class RoutesTab extends Vue {
  protected showCreateDialog = false

  protected showDeleteDialog = false

  routeIdToDelete = -1

  methods = ['GET', 'POST', 'PUT', 'DELETE']

  editedIndex = -1

  readonly headers = [
    {
      text: 'ID',
      value: 'id',
    },
    {
      text: 'Маршруты',
      value: 'route',
    },
    {
      text: 'Метод',
      value: 'method',
    },
    {
      text: 'Описание',
      value: 'description',
    },
    { text: 'Actions', value: 'actions', sortable: false },
  ]

  protected createRoute(): void {
    this.showCreateDialog = true;
  }

  protected editItem(item: any): void {
    // todo: create
    this.createRoute();
    console.log(item);
  }

  protected deleteItem(item: Route): void {
    this.showDeleteDialog = true;
   this.routeIdToDelete = item.id;
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
