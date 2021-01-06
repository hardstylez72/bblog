<template>
  <div>
    <v-form ref="refss" v-model="valid" lazy-validation>
      <v-row>
        <v-col cols="12" sm="10" md="10">
          <v-text-field
            v-model="route.route"
            required
            :rules="routeRules"
            label="Маршрут"
          />
        </v-col>
        <v-col cols="12" sm="2" md="2">
          <v-select
            v-model="route.method"
            required
            :rules="methodRules"
            :items="httpMethodList"
            label="Метод"
          />
        </v-col>
        <v-col cols="12" sm="10" md="10">
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
    <slot name="actions" v-bind="{ ref: this.$refs.refss }"/>
  </div>

</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';
import { Route } from '@/views/route/service';

@Component
export default class RouteForm extends Vue {
  valid = true

  route: Route = {
    description: '',
    id: -1,
    method: '',
    route: '/',
  }

  httpMethodList = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH']

  @Model('change', { default: {} })
  readonly value!: Route

  @Watch('value', { immediate: true })
  protected onChangeValue(route: Route): void {
    this.route = route;
  }

  @Watch('route', { deep: true })
  protected onChangeRoute(route: Route): void {
    if (JSON.stringify(route) !== JSON.stringify(this.value)) {
      this.$emit('change', route);
    }
  }

  routeRules = [
    (v: string) => !!v || 'Обязательное поле',
    (v: string) => {
      if (v) {
        if (v.length) {
          if (v[0] !== '/') {
            return 'Маршрут должен начинаться со знака `/`';
          }
        }
      }
      return true;
    },
    (v: string) => {
      if (v) {
        if (v.length) {
          if (v.includes('//')) {
            return 'Маршрут не должен содержать повторяющиеся знаки типа //';
          }
        }
      }
      return true;
    },
]

  methodRules = [
    (v: string) => !!v || 'Обязательное поле',
  ]

  descriptionRules = this.methodRules
}
</script>

<style scoped lang="scss">

</style>
