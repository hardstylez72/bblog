<template>
  <div>
    <v-data-table
      v-model="selected"
      :headers="headers"
      :items="items"
      sort-by="calories"
      class="elevation-1"
    >
      <template v-slot:no-data>
        Нет данных
      </template>

      <template v-slot:top>
        <slot name="top" />
      </template>

      <template v-slot:item.actions="{ item }">
        <slot name="item.actions" v-bind="{ item }"/>
      </template>

      <template v-slot:item.statuses="{ item }">
        <slot name="item.statuses" v-bind="{ item }"/>
      </template>

    </v-data-table>
  </div>
</template>

<script lang="ts">
import { Component } from 'vue-property-decorator';
import SelectableTable from '@/views/base/components/SelectableTable.vue';
import { Group } from '@/views/group/services/group';
import { DataTableHeader } from 'vuetify';

@Component
export default class UserRoutesTable extends SelectableTable<Group> {
  readonly headers: DataTableHeader[] = [
      { text: 'ID', value: 'id' },
      { text: 'Маршрут', value: 'route' },
      { text: 'Метод', value: 'method' },
      { text: 'Описание', value: 'description' },
      { text: 'Группы', value: 'groupCodes' },
      { text: 'Действия', value: 'actions', width: '10%' },
      { text: 'Статусы', value: 'statuses', width: '10%' },
    ]

  view(group: Group) {
    return this.$router.push({ name: 'Group', params: { id: group.id.toString() } });
  }
}
</script>

<style scoped lang="scss">
</style>
