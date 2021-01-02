<template>
  <div>
    <UserGroupsSelectableTable
      v-model="selected"
      :items="groups"
    >
      <template v-slot:top>
        <v-toolbar
          flat
        >
          <v-toolbar-title>{{ title }}</v-toolbar-title>
          <v-divider
            class="mx-4"
            inset
            vertical
          />
          <v-spacer />
          <div>
            <v-btn
              v-if="showDeleteBtn"
              color="primary"
              class="mb-2"
              @click="deleteSelectedRoutes"
            >
              Удалить выбранные группы
            </v-btn>
          </div>
          <GroupsNotBelongUserDialogButton :user-id="userIdC" />
        </v-toolbar>
      </template>
    </UserGroupsSelectableTable>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Group } from '@/views/group/services/group';
import UserGroupsSelectableTable from '../components/UserGroupsSelectableTable.vue';
import GroupsNotBelongUserDialogButton from '../components/AddGroupsButton.vue';

@Component({
  components: {
    UserGroupsSelectableTable,
    GroupsNotBelongUserDialogButton,
},
})
export default class UserPage extends Vue {
  userId = Number(this.$route.params.id);

  title = 'Группы'

  mounted() {
    this.$store.direct.dispatch.userGroup.GetListBelongToGroup(this.userId);
  }

  get userIdC(): number {
    return this.userId;
  }

  get showDeleteBtn(): boolean {
    return this.selected.length > 0;
  }

  selected: Group[] = []

  async deleteSelectedRoutes() {
    const groups = this.selected;
    const params = groups.map((group) => ({
      groupId: group.id,
      userId: this.userId,
    }));

    await this.$store.direct.dispatch.userGroup.Delete(params);
    this.selected = [];
  }

  get groups(): readonly Group[] {
    return this.$store.direct.getters.userGroup.getRoutesBelongToGroup;
  }
}
</script>

<style scoped lang="scss">

</style>
