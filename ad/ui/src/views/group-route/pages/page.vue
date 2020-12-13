<template>
  <div>
    <routes-table-select-delete />
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';

@Component({
  components: {
    'routes-table-select-delete': () => import('../../group/pages/routes-table-select-delete.vue'),
  },
  computed: {

  },

})
export default class GroupPage extends Vue {

}
</script>

<style scoped lang="scss">
.routes-tab-container {
  display: flex;
  height: 1000px;
  justify-content: space-between;
}
.create-route-btn {
  margin: 10px;

}
</style>
