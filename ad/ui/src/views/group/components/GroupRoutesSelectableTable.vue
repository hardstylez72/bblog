<script lang="ts">
import { Component } from 'vue-property-decorator';
import { Route } from '@/views/route/service';
import SelectableTable from '@/views/base/components/SelectableTable.vue';

@Component
export default class GroupRoutesSelectableTable extends SelectableTable<Route> {
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
