<template>
  <c-dialog
    v-model="show"
    max-width="450px"
  >
    <v-card>
      <v-card-title>
        Вы уверены что хотите удлить группу?
      </v-card-title>
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
          @click="deleteGroup"
        >
          OK
        </v-btn>
        <v-spacer />
      </v-card-actions>
    </v-card>
  </c-dialog>
</template>

<script lang="ts">
import {
  Component, Model, Prop, Vue, Watch,
} from 'vue-property-decorator';

@Component({
  components: {
    'c-dialog': () => import('../../base/components/Dialog.vue'),
  },
})
export default class DeleteRouteDialog extends Vue {
  show = false

  @Prop() id!: number

  @Model('change', { default: false, type: Boolean })
  readonly value!: boolean

  @Watch('value')
  protected onChangeValue(value: boolean): void {
    this.show = value;
  }

  close() {
    this.$emit('change', false);
    this.show = false;
  }

  async deleteGroup() {
    await this.$store.direct.dispatch.group.Delete(this.id);
    this.$emit('change', false);
  }
}
</script>

<style scoped lang="scss">

</style>
