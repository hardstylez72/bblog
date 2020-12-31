/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import GroupRouteService, { GroupRoute } from '@/views/group-route/service';
import { Route } from '@/views/route/service';
import { moduleActionContext } from '../base/store';

export interface State{
  service: GroupRouteService;
  entities: Route[];
  routesNotBelongToGroup: Route[];
  groupId: number;
}

const state1 = {
  service: new GroupRouteService({ host: '', baseUrl: '/api/v1/group/route' }),
  entities: [],
  routesNotBelongToGroup: [],
  groupId: -1,
} as State;

const mutations = defineMutations < State >()({

  setGroupId(state, groupId: number) {
    state.groupId = groupId;
  },
  setEntities(state, entities: Route[]) {
    state.entities = entities;
  },
  deleteEntity(state, pairs: GroupRoute[]) {
    state.entities = state.entities.filter((r: Route) => {
      const exist = pairs.some((pair) => pair.routeId === r.id);
      return !exist;
    });
  },
  setRoutesNotBelongToGroup(state, entities: Route[]) {
    state.routesNotBelongToGroup = entities;
  },
  deleteRoutesNotBelongToGroup(state, pairs: GroupRoute[]) {
    state.routesNotBelongToGroup = state.routesNotBelongToGroup.filter((r: Route) => {
      const exist = pairs.some((pair) => pair.routeId === r.id);
      return !exist;
    });
  },
  addEntity(state, entities) {
    state.entities.push(...entities);
  },
  addRoutesNotBelongToGroup(state, entities) {
    state.routesNotBelongToGroup.push(...entities);
  },
});

const actions = defineActions({

  async GetListNotBelongToGroup(context, groupId: number): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(groupId, false);
    commit.setRoutesNotBelongToGroup(entities);
    commit.setGroupId(groupId);
    return entities;
  },
  async GetList(context, payload: {groupId: number; belongToGroup: boolean}): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(payload.groupId, payload.belongToGroup);

    if (payload.belongToGroup) {
      commit.setEntities(entities);
      commit.setGroupId(payload.groupId);
    }

    return entities;
  },
  async Create(context, pairs: GroupRoute[]): Promise<Route[]> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Create(pairs);
    commit.addEntity(createdEntity);
    commit.deleteRoutesNotBelongToGroup(pairs);
    return createdEntity;
  },
  async Delete(context, pairs: GroupRoute[]): Promise<void> {
    const { state, commit } = actionContext(context);
    await state.service.Delete(pairs);
    commit.deleteEntity(pairs);
    commit.addRoutesNotBelongToGroup(pairs);
  },
});

const getters = defineGetters<State>()({
  getEntities(state): Route[] {
    return state.entities;
  },
  getRoutesNotBelongToGroup(state): Route[] {
    return state.routesNotBelongToGroup;
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
