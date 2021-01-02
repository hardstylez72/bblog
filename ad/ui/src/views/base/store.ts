/* eslint-disable import/no-cycle */

import Vue from 'vue';
import Vuex from 'vuex';
import { createDirectStore } from 'direct-vuex';
import routeModule from '../route/store';
import groupModule from '../group/store/group';
import userModule from '../user/store/store';
import groupRouteModule from '../group/store/grouproute';
import userGroupModule from '../user/store/usergroup';
import userRouteModule from '../user/store/userroute';

Vue.use(Vuex);

const {
  store,
  rootActionContext,
  moduleActionContext,
  rootGetterContext,
  moduleGetterContext,
} = createDirectStore({

  actions: {},
  modules: {
    route: routeModule,
    group: groupModule,
    groupRoute: groupRouteModule,
    user: userModule,
    userGroup: userGroupModule,
    userRoute: userRouteModule,
  },
});

export default store;

export {
  rootActionContext,
  moduleActionContext,
  rootGetterContext,
  moduleGetterContext,
};

export type AppStore = typeof store
declare module 'vuex' {
  interface Store<S> {
    direct: AppStore;
  }
}
