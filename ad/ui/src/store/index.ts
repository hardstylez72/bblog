/* eslint-disable import/no-cycle */

import Vue from 'vue';
import Vuex from 'vuex';
import { createDirectStore } from 'direct-vuex';
import routeModule from './route';

Vue.use(Vuex);

const {
  store,
  rootActionContext,
  moduleActionContext,
  rootGetterContext,
  moduleGetterContext,
} = createDirectStore({

  actions: {
  },
  modules: {
    route: routeModule,
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
