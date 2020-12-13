<template>
  <div>
    <v-data-table
      v-bind="tableAttrs"
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
        <v-icon
          small
          class="mr-2"
          @click="view(item)"
        >
          mdi-eye
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
  Component, Vue, Prop, Inject,
} from 'vue-property-decorator';
import { DataTableHeader } from 'vuetify';
import { User } from '@/views/user/service';

@Component({
  components: {
    'create-dialog': () => import('../../user/components/create-dialog.vue'),
    'delete-dialog': () => import('../../user/components/delete-dialog.vue'),
  },
})
export default class DictTable<T> extends Vue {
  protected items: T[] = []

  protected title = ''

  protected headers: DataTableHeader[] = [];

  protected showDeleteDialog = false

  protected showEditDialog = false

  private activeItemId = -1

  protected get tableAttrs() {
    return {
      headers: this.headers,
      items: this.items,
    };
  }

  protected edit(item: T): void {
    this.showEditDialog = true;
    this.activeItemId = item.id;
  }

  protected view(item: T): void {
    this.activeItemId = item.id;
  }

  protected remove(item: T): void {
    this.showDeleteDialog = true;
   this.activeItemId = item.id;
  }
}
</script>

<style scoped lang="scss">

</style>
