<template>
  <div>
    <h2>Пользователь {{userC.externalId}}</h2>
    <UserGroupsSelectableTable v-model="selectedGroups" :items="groups">
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{ titleGroup }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer/>
          <div>
            <v-btn
              v-if="showDeleteGroupButton"
              color="primary"
              class="mb-2"
              @click="deleteSelectedGroups"
            >
              Удалить выбранные группы
            </v-btn>
          </div>
          <AddGroupsButton :user-id="userIdC" />
        </v-toolbar>
      </template>
    </UserGroupsSelectableTable>

    <UserRoutesSelectableTable v-model="selectedRoutes" :items="routes" >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>{{ titleRoute }}</v-toolbar-title>
          <v-divider class="mx-4" inset vertical/>
          <v-spacer/>
          <div>
            <v-btn
              v-if="showDeleteRouteButton"
              color="primary"
              class="mb-2"
              @click="deleteSelectedRoutes"
            >
              Удалить выбранные группы
            </v-btn>
          </div>
          <AddRoutesButton :user-id="userIdC" />
        </v-toolbar>
      </template>
    </UserRoutesSelectableTable>
  </div>
</template>

<script lang="ts">
import {
  Component, Vue,
} from 'vue-property-decorator';
import { Group } from '@/views/group/services/group';
import { Route } from '@/views/route/service';
import { User } from '@/views/user/services/user';
import UserGroupsSelectableTable from '../components/UserGroupsSelectableTable.vue';
import UserRoutesSelectableTable from '../components/UserRoutesSelectableTable.vue';
import AddGroupsButton from '../components/AddGroupsButton.vue';
import AddRoutesButton from '../components/AddRoutesButton.vue';

@Component({
  components: {
    UserGroupsSelectableTable,
    UserRoutesSelectableTable,
    AddGroupsButton,
    AddRoutesButton,
},
})
export default class UserPage extends Vue {
  userId = Number(this.$route.params.id);

  user: User = {
    externalId: 'Не найден',
    isSystem: false,
    description: '',
    id: -1,
  }

  titleGroup = 'Группы'

  titleRoute = 'Маршруты'

  selectedGroups: Group[] = []

  selectedRoutes: Route[] = []

 async mounted() {
    this.$store.direct.dispatch.userGroup.GetListBelongToGroup(this.userId);
    this.$store.direct.dispatch.userRoute.GetListBelongToGroup(this.userId);
    this.$store.direct.dispatch.user.GetById(this.userId).then((user) => {
       this.user = user;
     });
  }

  get userC(): User {
    return this.user;
  }

  get userIdC(): number {
    return this.userId;
  }

  get showDeleteGroupButton(): boolean {
    return this.selectedGroups.length > 0;
  }

  get showDeleteRouteButton(): boolean {
    return this.selectedRoutes.length > 0;
  }

  async deleteSelectedRoutes() {
    const routes = this.selectedRoutes;
    const params = routes.map((route) => ({
      routeId: route.id,
      userId: this.userId,
    }));

    await this.$store.direct.dispatch.userRoute.Delete(params);
    this.selectedGroups = [];
  }

  async deleteSelectedGroups() {
    const groups = this.selectedGroups;
    const params = groups.map((group) => ({
      groupId: group.id,
      userId: this.userId,
    }));

    await this.$store.direct.dispatch.userGroup.Delete(params);
    this.selectedGroups = [];
  }

  get groups(): readonly Group[] {
    return this.$store.direct.getters.userGroup.getGroupsBelongToUser;
  }

  get routes(): readonly Route[] {
    return this.$store.direct.getters.userRoute.getRoutesBelongToUser;
  }
}
</script>

<style scoped lang="scss">

</style>
