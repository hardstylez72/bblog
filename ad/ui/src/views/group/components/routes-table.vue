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
import { Service } from '@/views/route/service';

@Component
export default class RoutesTab extends Vue {
  @Prop({ default: () => ([]), type: Array }) items: Array<Service>

  selected: Service[] = []

  @Model('change', { default: () => ([]), type: Array })
  readonly value!: []

  @Watch('value')
  protected onChangeValue(value: Service[]): void {
    this.selected = value;
  }

  @Watch('selected')
  protected onChangeIsShowDialog(selected: Service[]): void {
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
