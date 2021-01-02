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
        Добавить группы
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        Добавление групп к пользователю
      </v-card-title>

      <UserGroupsSelectableTable
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
                Добавить выбранные группы
              </v-btn>
            </div>
          </v-toolbar>
        </template>
      </UserGroupsSelectableTable>
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
import { Group } from '@/views/group/services/group';

import UserGroupsSelectableTable from './UserGroupsSelectableTable.vue';

@Component({
  components: {
    'c-dialog': () => import('../../base/components/Dialog.vue'),
    UserGroupsSelectableTable,
  },
})
export default class RoutesTableSelectAddDialog extends Vue {
  show = false

  @Prop({ type: Number, default: -1 })
  private readonly userId!: number

  entities: Group[] =[]

  selected: Group[] =[]

  mounted() {
    this.$store.direct.dispatch.userGroup.GetListNotBelongToGroup(this.userId);
  }

  get routes(): readonly Group[] {
    return this.$store.direct.getters.userGroup.getRoutesNotBelongToGroup;
  }

  get showAddBtn(): boolean {
    return this.selected.length > 0;
  }

  async addSelectedRoutes() {
    const groups = this.selected;
    const params = groups.map((group) => ({
      userId: this.userId,
      groupId: group.id,
    }));

    await this.$store.direct.dispatch.userGroup.Create(params);
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
