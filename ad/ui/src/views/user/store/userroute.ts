/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import UserRouteService, { UserRoute } from '@/views/user/services/userroute';

import { Route } from '@/views/route/service';
import { moduleActionContext } from '../../base/store';

export interface State{
  service: UserRouteService;
  routesBelongToUser: Route[];
  routesNotBelongToUser: Route[];
  userId: number;
}

const state1 = {
  service: new UserRouteService({ host: '', baseUrl: '/api/v1/user/route' }),
  routesBelongToUser: [],
  routesNotBelongToUser: [],
  userId: -1,
} as State;

const mutations = defineMutations < State >()({

  setUserId(state, userId: number) {
    state.userId = userId;
  },
  setRoutesBelongToGroup(state, entities: Route[]) {
    state.routesBelongToUser = entities;
  },
  deleteRoutesBelongToGroup(state, pairs: UserRoute[]) {
    state.routesBelongToUser = state.routesBelongToUser.filter((r: Route) => {
      const exist = pairs.some((pair) => pair.routeId === r.id);
      return !exist;
    });
  },
  addRoutesBelongToGroup(state, entities) {
    state.routesBelongToUser.push(...entities);
  },

  setRoutesNotBelongToGroup(state, entities: Route[]) {
    state.routesNotBelongToUser = entities;
  },
  deleteRoutesNotBelongToGroup(state, pairs: UserRoute[]) {
    state.routesNotBelongToUser = state.routesNotBelongToUser.filter((r: Route) => {
      const exist = pairs.some((pair) => pair.routeId === r.id);
      return !exist;
    });
  },
  addRoutesNotBelongToGroup(state, entities) {
    state.routesNotBelongToUser.push(...entities);
  },

});

const actions = defineActions({

  async GetListNotBelongToGroup(context, groupId: number): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(groupId, false);
    commit.setRoutesNotBelongToGroup(entities);
    commit.setUserId(groupId);
    return entities;
  },
  async GetListBelongToGroup(context, groupId: number): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(groupId, true);
    commit.setRoutesBelongToGroup(entities);
    commit.setUserId(groupId);
    return entities;
  },
  async Create(context, pairs: UserRoute[]): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Create(pairs);
    commit.addRoutesBelongToGroup(createdEntity);
    commit.deleteRoutesNotBelongToGroup(pairs);
    return createdEntity;
  },
  async Delete(context, pairs: UserRoute[]): Promise<void> {
    const { state, commit } = actionContext(context);
    await state.service.Delete(pairs);
    commit.deleteRoutesBelongToGroup(pairs);
    commit.addRoutesNotBelongToGroup(pairs);
  },
});

const getters = defineGetters<State>()({
  getRoutesBelongToUser(state): Route[] {
    return state.routesBelongToUser;
  },
  getRoutesNotBelongToGroup(state): Route[] {
    return state.routesNotBelongToUser;
  },
});

const module = defineModule({
  namespaced: true as true,
  state: state1,
  getters,
  mutations,
  actions,
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
