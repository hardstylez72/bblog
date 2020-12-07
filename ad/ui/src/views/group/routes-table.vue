<template>
  <div>
    <v-data-table
      v-model="selected"
      :headers="headers"
      :items="items"
      sort-by="calories"
      class="elevation-1"
      show-select
    >
      <template v-slot:no-data>
        Нет данных
      </template>

      <template v-slot:top>
        <slot name="top" />
      </template>
    </v-data-table>
  </div>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import { Route } from '@/services/route';

@Component
export default class RoutesTab extends Vue {
  @Prop({ default: () => ([]), type: Array }) items: Array<Route>

  selected: Route[] = []

  @Model('change', { default: () => ([]), type: Array })
  readonly value!: []

  @Watch('value')
  protected onChangeValue(value: Route[]): void {
    this.selected = value;
  }

  @Watch('selected')
  protected onChangeIsShowDialog(selected: Route[]): void {
      this.$emit('change', selected);
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
</style>
