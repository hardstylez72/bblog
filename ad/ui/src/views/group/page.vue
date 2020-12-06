<template>
  <div>
    <page-routes />
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';

@Component({
  components: {
    'page-routes': () => import('./page-routes.vue'),
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
