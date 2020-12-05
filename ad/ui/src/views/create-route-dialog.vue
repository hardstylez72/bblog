<template>

        <v-dialog
          v-model="show"
          max-width="500px"

          persistent
        >

          <v-card>
            <v-card-title>
              <span class="headline">{{title}}</span>
            </v-card-title>

            <v-card-text>

            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="blue darken-1"
                text
                @click="close"
              >
                Cancel
              </v-btn>
              <v-btn
                color="blue darken-1"
                text
                @click="createRoute"
              >
                Save
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>

</template>

<script lang="ts">
import {
  Component, Model, Vue, Watch,
} from 'vue-property-decorator';
import { Route } from '@/services/route';

@Component
export default class RoutesTab extends Vue {
  protected show = false;

  @Model('change', { default: false, type: Boolean })
  protected readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    console.log('Внешка поменялась');
    this.show = value;
  }

  @Watch('show')
  protected onChangeIsShowDialog(show: boolean): void {
    if (show !== this.value) {
      console.log('Внутри поменялась');
      this.$emit('change', show);
    }
  }

  route: Route = {
    description: '',
    id: -1,
    method: 'POST',
    route: '',
  }

  title = 'Новый маршрут'

  createRoute() {
    return this.$store.direct.dispatch.route.Create(this.$data.route);
  }

  close() {
    this.show = false;
  }
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
