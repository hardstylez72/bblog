<template>
  <c-dialog
    v-model="show"
  >
    <template v-slot:activator="props">
      <v-btn
        color="primary"

        class="mb-2"
        v-bind="props"
        v-on="props.on"
      >
        Новый маршрут
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="headline grey lighten-2">
        Создание маршрута
      </v-card-title>
      <v-card-text>
        <v-form
          ref="form"
          v-model="valid"
          lazy-validation
        >
          <v-row>
            <v-col
              cols="12"
              sm="10"
              md="10"
            >
              <v-text-field
                v-model="route.route"
                required
                :rules="routeRules"
                label="Маршрут"
              />
            </v-col>
            <v-col
              cols="12"
              sm="2"
              md="2"
            >
              <v-select
                v-model="route.method"
                required
                :rules="methodRules"
                :items="httpMethodList"
                label="Метод"
              />
            </v-col>
            <v-col
              cols="12"
              sm="10"
              md="10"
            >
              <v-textarea
                v-model="route.description"
                outlined
                required
                :rules="descriptionRules"
                label="Описание"
              />
            </v-col>
          </v-row>
        </v-form>

        <v-card-actions>
          <v-spacer />
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
      </v-card-text>
    </v-card>
  </c-dialog>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';

@Component({
  components: {
    'c-dialog': () => import('../base/components/dialog.vue'),
  },
})
export default class CreateRouteDialog extends Vue {
  show = false

  valid = true

  route: Route = {
    description: '',
    id: -1,
    method: '',
    route: '',
  }

  httpMethodList = ['GET', 'POST', 'PUT', 'DELETE']

  validate() {
    this.$refs.form.validate();
  }

  routeRules = [
    (v) => !!v || 'Обязательное поле',
]

  methodRules = this.routeRules

  descriptionRules = this.routeRules

  async createRoute() {
    this.validate();
    if (!this.valid) {
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

    await this.$store.direct.dispatch.route.Create(this.$data.route);
    this.show = false;
  }

  close() {
    this.show = false;
  }
}
</script>

<style scoped lang="scss">

</style>
