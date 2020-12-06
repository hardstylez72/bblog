<template>
  <div>
    <MainTabs />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import MainTabs from '@/views/main-page-tabs.vue'; // @ is an alias to /src

@Component({
  components: {
    MainTabs,
  },
})
export default class Home extends Vue {

}
</script>
