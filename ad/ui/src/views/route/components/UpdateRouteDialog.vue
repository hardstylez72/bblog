<template>
  <Dialog v-model="show">
    <v-card>
      <v-card-title class="headline grey lighten-2">
        Редактирование маршрута
      </v-card-title>
      <v-card-text>
        <EditRouteForm v-model="route">
          <template v-slot:actions="{ref}">
            <v-card-actions>
              <v-spacer />
              <v-btn color="blue darken-1" text @click="close">Отмена</v-btn>
              <v-btn color="blue darken-1" text @click="updateRoute(ref)">Обновить</v-btn>
            </v-card-actions>
          </template>
        </EditRouteForm>
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
import EditRouteForm from './EditRouteForm.vue';

@Component({
  components: {
    Dialog,
    EditRouteForm,
  },
})
export default class CreateRouteDialog extends Vue {
  show = false

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

  @Watch('item', { deep: true })
  protected onChangeItem(value: Route): void {
    this.route = {
      description: value.description,
      id: value.id,
      method: value.method,
      route: value.route,
    };
  }

  async updateRoute(ref: any) {
    if (ref) {
      ref.validate();
    }

    if (!this.route.id) {
      return;
    }
    if (!this.route.description) {
      return;
    }

    if (!this.route.method) {
      return;
    }

    if (!this.route.route) {
      return;
    }

    await this.$store.direct.dispatch.route.Update(this.$data.route);
    this.$emit('change', false);
    this.show = false;
  }

  close() {
    this.$emit('change', false);
    this.show = false;
  }
}
</script>

<style scoped lang="scss">

</style>
