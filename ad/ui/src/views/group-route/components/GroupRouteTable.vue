<template>
  <c-dialog
    v-model="show"
    max-width="2000px"
  >
    <template v-slot:activator="props">
      <v-btn
        color="primary"

        class="mb-2"
        v-bind="props"
        v-on="props.on"
      >
        Добавить маршруты
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        Добавление маршрутов к группе
      </v-card-title>

      <routes-table
        v-model="selected"
        :items="routes"
      >
        <template v-slot:top>
          <v-toolbar
            flat
          >
            <v-spacer />
            <div>
              <v-btn
                v-if="showAddBtn"
                color="primary"
                class="mb-2"
                @click="addSelectedRoutes"
              >
                Добавить выбранные маршруты
              </v-btn>
            </div>
          </v-toolbar>
        </template>
      </routes-table>
      <v-card-actions>
        <v-spacer />
        <v-btn
          color="blue darken-1"
          text
          @click="close"
        >
          Cancel
        </v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </c-dialog>
</template>

<script lang="ts">
import {
  Component, Vue, Prop,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';

@Component({
  components: {
    'c-dialog': () => import('../../base/components/Dialog.vue'),
    'routes-table': () => import('../../group/components/RoutesTable.vue'),
  },
})
export default class RoutesTableSelectAddDialog extends Vue {
  show = false

  @Prop({ type: Number, default: -1 })
  private readonly groupId!: number

  entities: Route[] =[]

  selected: Route[] =[]

  mounted() {
    this.$store.direct.dispatch.groupRoute.GetListNotBelongToGroup(this.groupId);
  }

  get routes(): readonly Route[] {
    return this.$store.direct.getters.groupRoute.getRoutesNotBelongToGroup;
  }

  get showAddBtn(): boolean {
    return this.selected.length > 0;
  }

  async addSelectedRoutes() {
    const routes = this.selected;
    const params = routes.map((route) => ({
      groupId: this.groupId,
      routeId: route.id,
    }));

    await this.$store.direct.dispatch.groupRoute.Create(params);
    this.selected = [];
    this.show = false;
  }

  close() {
    this.show = false;
  }
}
</script>

<style scoped lang="scss">

</style>
