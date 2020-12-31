<template>
  <div />
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Service } from '@/views/route/service';

@Component({
  components: {
},
})
export default class UserPage extends Vue {

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
