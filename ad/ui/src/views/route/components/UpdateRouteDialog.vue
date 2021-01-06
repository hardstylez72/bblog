<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">
        Редактирование маршрута
      </v-card-title>
      <v-card-text>
        <RouteForm v-model="route">
          <template v-slot:actions="{ref}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">Отмена</v-btn>
              <v-btn color="blue darken-1" text :disabled="disableUpdateButton" @click="updateRoute(ref)">Обновить</v-btn>
            </v-card-actions>
          </template>
        </RouteForm>
      </v-card-text>
    </v-card>
  </Dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';
import Dialog from '@/views/base/components/Dialog.vue';
import RouteForm from './RouteForm.vue';

@Component({
  components: {
    Dialog,
    RouteForm,
  },
})
export default class UpdateRouteDialog extends Vue {
  show = false

  disableUpdateButton = true

  valid = true

  @Prop({ required: true }) item!: Route

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  route: Route = {
    description: '',
    id: -1,
    method: '',
    route: '/',
  }

   initialRouteState: Route | undefined

  @Watch('item', { deep: true })
  protected onChangeItem(item: Route): void {
    this.route = {
      description: item.description,
      id: item.id,
      method: item.method,
      route: item.route,
    };
    this.initialRouteState = {
      description: item.description,
      id: item.id,
      method: item.method,
      route: item.route,
    };
  }

  @Watch('route', { deep: true })
  protected onChangeRoute(route: Route): void {
    if (!this.routesSame(this.item, route)) {
      this.disableUpdateButton = false;
    }

    if (this.initialRouteState) {
      if (this.routesSame(this.initialRouteState, route)) {
        this.disableUpdateButton = true;
      }
    }
  }

  async updateRoute(ref: any) {
    if (ref) {
      ref.validate(); // vee-validate specifics
    }

    if (!this.route.id || !this.route.description || !this.route.method || !this.route.route) {
      return;
    }

    await this.$store.direct.dispatch.route.Update(this.$data.route);
    this.$emit('change', false);
    this.show = false;
  }

  routesSame(a: Route, b: Route): boolean {
    return (a.description === b.description
      && a.method === b.method
      && a.route === b.route);
  }

  close() {
    this.$emit('change', false);
    this.show = false;
  }
}
</script>

<style scoped lang="scss">

</style>
