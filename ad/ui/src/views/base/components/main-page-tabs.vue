<template>
  <v-card>
    <v-toolbar
      color="cyan"
      dark
      flat
    >
      <template v-slot:extension>
        <v-tabs
          v-model="tab"
          align-with-title
          @change="tabChanged"
        >
          <v-tabs-slider color="yellow" />

          <v-tab
            v-for="item in tabs"
            :key="item"
          >
            {{ item }}
          </v-tab>
        </v-tabs>
      </template>
    </v-toolbar>

    <v-tabs-items
      v-model="tab"
    >
      <v-tab-item
        v-for="item in tabs"
        :key="item"
      >
        <div v-if="item === tabs[0]">
          <routes-tab />
        </div>
        <div v-if="item === tabs[1]">
          <groups-tab />
        </div>
        <div v-if="item === tabs[2]">
          <user-tab />
        </div>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script lang="ts">
import {
  Component, Prop, Vue, Watch,
} from 'vue-property-decorator';

@Component({
  components: {
    routesTab: () => import('@/views/route/tab-table.vue'),
    groupsTab: () => import('@/views/group/components/tap-table.vue'),
    userTab: () => import('@/views/user/components/tap-table.vue'),
  },

})
export default class MainTabs extends Vue {
  @Prop() private msg!: string;

  tab = 0;

  tabs: string[] = [
    'routes', 'groups', 'users',
  ];

  @Watch('$route')
  RouteUpdate() {
    const tabNumber = this.getTabNumberFromUrlQueryParams();
    if (this.tab !== tabNumber) {
      this.tab = tabNumber;
    }
  }

  mounted() {
    const tabNumber = this.getTabNumberFromUrlQueryParams();
    if (this.tab !== tabNumber) {
      this.tab = tabNumber;
    }

    const { tab } = this.$route.query;
    if (tab !== this.tabs[tabNumber]) {
      this.$router.push({ query: { tab: this.tabs[tabNumber] } });
    }
  }

  tabChanged(tabNumber: number) {
    if (tabNumber !== this.getTabNumberFromUrlQueryParams()) {
      this.$router.push({ query: { tab: this.tabs[tabNumber] } });
    }
  }

  getTabNumberFromUrlQueryParams(): number {
    const { tab } = this.$route.query;
    let tabNumber = 0;
    if (tab) {
      this.tabs.some((itemTabName: string, itemTabNumber: number) => {
        if (itemTabName === tab) {
          tabNumber = itemTabNumber;
          return true;
        }
        return false;
      });
    }
    return tabNumber;
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 10px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
routes-tab {
  display: flex;
  flex-direction: column;
}
</style>
