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
import { Route } from '@/views/route/service';

@Component
export default class SelectableTable<T> extends Vue {
  @Prop({ default: () => ([]), type: Array }) items: Array<T>

  protected selected: T[] = []

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

  protected headers = []
}
</script>

<style scoped lang="scss">
</style>
