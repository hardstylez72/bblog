<template>
  <div>
    <v-form ref="refss" v-model="valid" lazy-validation>
      <v-row>
        <v-col cols="12" sm="12" md="10">
          <v-text-field
            v-model="route.route"
            required
            :rules="routeRules"
            label="Маршрут"
          />
        </v-col>
        <v-col cols="15" sm="2" md="4">
          <v-select
            v-model="route.method"
            required
            :rules="methodRules"
            :items="httpMethodList"
            label="Метод"
          />
        </v-col>
        <v-col cols="12" sm="10" md="8">
          <v-textarea
            v-model="route.description"
            outlined
            required
            :rules="descriptionRules"
            label="Описание"
          />
        </v-col>
        <v-col cols="12" sm="15" md="15">
          <v-autocomplete
            ref="autocomplete-input"
            v-model="selectedTags"
            :items="suggestedTags"
            :search-input.sync="searchTags"
            item-text="name"
            item-value="name"
            hide-selected
            :disabled="isSuggestUpdating"
            multiple
            chips
          >
            <template v-slot:no-data>
              <v-list-item>
                <v-list-item-title>
                  Поиск тегов
                </v-list-item-title>
              </v-list-item>
            </template>
            <template v-slot:selection="data">
              <v-chip
                v-bind="data.attrs"
                :input-value="data.selected"
                close
                @click:close="removeTagFromSelected(data.item)"
              >
                {{ data.item.name }}
              </v-chip>
            </template>
          </v-autocomplete>
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
import { Tag } from '@/views/tag/service';

@Component
export default class RouteForm extends Vue {
  valid = true

  selectedTags: string[] = []

  suggestedTags: Tag[] = []

  isSuggestUpdating = false

  searchTags = ''

  route: Route = {
    description: '',
    id: -1,
    method: '',
    route: '/',
    tags: [],
  }

  httpMethodList = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH']

  @Watch('selectedTags')
  async onChangeSelectedTags(pattern: string) {
    this.route.tags = this.selectedTags;
    this.searchTags = '';
  }

  async delay(ms: number) {
    return new Promise((res) => {
      setTimeout(() => {
        res();
      }, ms);
    });
  }

  @Watch('searchTags')
  async onChangeSearchTags(pattern: string) {
    if (!pattern) return;

    if (this.isSuggestUpdating) {
      return;
    }

    await this.delay(400);

    this.isSuggestUpdating = true;

    // api suggest
    this.suggestedTags = await this.$store.direct.dispatch.tag.GetByPattern(pattern)
      .finally(() => {
        this.isSuggestUpdating = false;
      });

    // after suggest request focus fades away
     this.$refs['autocomplete-input'].focus();
    // if not found => create new tag
    if (this.suggestedTags.length === 0) {
      this.suggestedTags.push({ name: pattern, id: 1 });
    }

    // to keep selected tags in input
    const selected: Tag[] = this.selectedTags.map((tagName) => ({
        name: tagName,
        id: -1,
      }));

    this.suggestedTags.push(...selected);
  }

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

  removeTagFromSelected(tag: Tag) {
    this.selectedTags = this.selectedTags.filter((tagName) => tagName !== tag.name);
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
