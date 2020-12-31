<template>
  <div>
    <v-data-table
      :items="routes"
      :headers="headers"
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
          <v-toolbar-title>{{ title }}</v-toolbar-title>
          <v-divider
            class="mx-4"
            inset
            vertical
          />
          <v-spacer />
          <create-dialog />
        </v-toolbar>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-icon
          small
          class="mr-2"
          @click="edit(item)"
        >
          mdi-pencil
        </v-icon>
        <v-icon
          small
          @click="remove(item)"
        >
          mdi-delete
        </v-icon>
      </template>
    </v-data-table>
    <delete-dialog
      :id="activeItemId"
      v-model="showDeleteDialog"
    />
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';
import DictTable from '@/views/base/components/tap-table.vue';

@Component({
  components: {
    createRouteDialog: () => import('./create-dialog.vue'),
    deleteRouteDialog: () => import('./delete-dialog.vue'),
  },
  computed: {
    routes(): readonly Route[] {
      return this.$store.direct.getters.route.getRoutes;
    },
  },
  mounted() {
    this.$store.direct.dispatch.route.GetList();
  },
})
export default class RoutesTab extends DictTable<Route> {
  readonly title = 'Маршруты'

  mounted() {
    this.$store.direct.dispatch.route.GetList();
  }

  get routes(): readonly Route[] {
    return this.$store.direct.getters.route.getRoutes;
  }

  methods = ['GET', 'POST', 'PUT', 'DELETE']

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
