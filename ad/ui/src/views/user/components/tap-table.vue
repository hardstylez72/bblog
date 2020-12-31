<template>
  <div>
    <v-data-table
      :items="users"
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
          @click="view(item)"
        >
          mdi-eye
        </v-icon>
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
  Component,
} from 'vue-property-decorator';

import { DataTableHeader } from 'vuetify';
import { User } from '../service';
import DictTable from '../../base/components/DictTable.vue';

@Component
export default class UserDictTable extends DictTable<User> {
  protected title = 'Пользователи'

  get users(): readonly User[] {
    return this.$store.direct.getters.user.getEntities;
  }

  mounted() {
    this.$store.direct.dispatch.user.GetList();
  }

  protected view(item: T): void {
     this.$router.push({ name: 'User', params: { id: item.id.toString() } });
  }

  protected headers: DataTableHeader[] = [
    { text: 'ID', value: 'id' },
    { text: 'Внеший ID', value: 'externalId' },
    { text: 'Системный пользователь', value: 'isSystem' },
    { text: 'Actions', value: 'actions', sortable: false },
  ];
}
</script>

<style scoped lang="scss">

</style>
