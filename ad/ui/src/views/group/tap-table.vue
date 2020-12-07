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
          <v-toolbar-title>Группы</v-toolbar-title>
          <v-divider
            class="mx-4"
            inset
            vertical
          />
          <v-spacer />
          <create-group-dialog />
        </v-toolbar>
      </template>

      <template v-slot:item.code="{ item }">
        <div
          class="group-list-item"
          @click="groupClick(item)"
        >
          {{ item.code }}
        </div>
      </template>

      <template v-slot:item.actions="{ item }">
        <v-icon
          small
          class="mr-2"
          @click="editGroup(item)"
        >
          mdi-pencil
        </v-icon>
        <v-icon
          small
          @click="deleteGroup(item)"
        >
          mdi-delete
        </v-icon>
      </template>
    </v-data-table>
    <delete-route-dialog
      :id="groupIdToDelete"
      v-model="showDeleteDialog"
    />
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Group } from '@/services/group';

@Component({
  components: {
    'create-group-dialog': () => import('./create-dialog.vue'),
    'delete-route-dialog': () => import('./delete-dialog.vue'),
  },
    computed: {
    routes(): Group[] {
      return this.$store.direct.getters.group.getEntities;
    },
  },
  mounted() {
    this.$store.direct.dispatch.group.GetList();
  },
})
export default class RoutesTab extends Vue {
  protected showCreateDialog = false

  protected showDeleteDialog = false

  groupIdToDelete = -1

  valid = true

  editedIndex = -1

  readonly headers = [
    {
      text: 'ID',
      value: 'id',
    },
    {
      text: 'Код',
      value: 'code',
    },
    {
      text: 'Описание',
      value: 'description',
    },
    { text: 'Actions', value: 'actions', sortable: false },
  ]

 async groupClick(group: Group) {
    await this.$router.push({ name: 'Group', params: { id: group.id.toString() } });
  }

  protected createGroup(): void {
    this.showCreateDialog = true;
  }

  protected editGroup(item: any): void {
    // todo: create
    this.createGroup();
  }

  protected deleteGroup(item: Group): void {
    this.showDeleteDialog = true;
   this.groupIdToDelete = item.id;
  }
}
</script>

<style scoped lang="scss">
.group-list-item {
  cursor: pointer;
}
.group-list-item:hover {
  color: blue;
}

.routes-tab-container {
  display: flex;
  height: 1000px;
  justify-content: space-between;
}
.create-route-btn {
  margin: 10px;

}
</style>
