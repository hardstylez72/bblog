<template>
  <div>
    <routes-table
      v-model="selected"
      :items="routes"
    >
      <template v-slot:top>
        <v-toolbar
          flat
        >
          <v-toolbar-title>{{ dict.title }}</v-toolbar-title>
          <v-divider
            class="mx-4"
            inset
            vertical
          />
          <v-spacer />
          <div>
            <v-btn
              v-if="showDeleteBtn"
              color="primary"
              class="mb-2"
              @click="deleteSelectedRoutes"
            >
              Удалить выбранные маршруты
            </v-btn>
          </div>
          <routes-table-select-add-dialog :group-id="groupIdC" />
        </v-toolbar>
      </template>
    </routes-table>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';

@Component({
  components: {
    'routes-table': () => import('../components/GroupRoutesSelectableTable.vue'),
    routesTableSelectAddDialog: () => import('../components/RoutesNotBelongGroupDialogButton.vue'),
  },
})
export default class RoutesTab extends Vue {
  dict = {
    title: `Маршруты группы ${this.$route.params.id}`,
  }

  get groupIdC(): number {
    return this.groupId;
  }

  groupId = Number(this.$route.params.id);

  mounted() {
    this.$store.direct.dispatch.groupRoute.GetListBelongToGroup(this.groupId);
  }

  get routes(): readonly Route[] {
    return this.$store.direct.getters.groupRoute.getRoutesBelongToGroup;
  }

  get showDeleteBtn(): boolean {
    return this.selected.length > 0;
  }

  selected: Route[] = []

  entities: Route[] = []

  valid = true

  editedIndex = -1

  async deleteSelectedRoutes() {
    const routes = this.selected;
    const groupId = Number(this.$route.params.id);
    const params = routes.map((route) => ({
        groupId,
        routeId: route.id,
      }));

    await this.$store.direct.dispatch.groupRoute.Delete(params);
    this.selected = [];
  }

  readonly headers = [
    {
      text: 'ID',
      value: 'id',
    },
    {
      text: 'Маршрут',
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
