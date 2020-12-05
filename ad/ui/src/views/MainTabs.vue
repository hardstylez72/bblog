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
        >
          <v-tabs-slider color="yellow"></v-tabs-slider>

          <v-tab
            v-for="item in items"
            :key="item"
          >
            {{ item }}
          </v-tab>
        </v-tabs>
      </template>
    </v-toolbar>

    <v-tabs-items v-model="tab">
      <v-tab-item
        v-for="item in items"
        :key="item"
      >
        <div v-if="item === 'Routes'">
          <routes-tab></routes-tab>
        </div>
        <v-card flat>
          <v-card-text v-text="text"></v-card-text>
        </v-card>
      </v-tab-item>
    </v-tabs-items>
  </v-card>
</template>

<script lang="ts">
import {
  Component, Prop, Vue,
} from 'vue-property-decorator';

@Component({
  components: {
    'routes-tab': () => import('./RoutesTab.vue'),
  },
})
export default class MainTabs extends Vue {
  @Prop() private msg!: string;

  tab: object| null = {};

  items: string[] = [
    'Routes', 'Groups', 'Users',
  ];

  text = 'bozdo text';
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
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
</style>
