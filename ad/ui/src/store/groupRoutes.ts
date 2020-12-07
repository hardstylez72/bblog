/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import GroupRouteService, { GroupRoute } from '@/services/groupRoute';
import { Route } from '@/services/route';
import { moduleActionContext } from './index';

export interface State{
  service: GroupRouteService;
  entities: Route[];
  entitiesMap: Map<number, Route>;
  groupId: number;
}

const state1 = {
  service: new GroupRouteService({ host: '', baseUrl: '/api/v1/group/route' }),
  entities: [],
  entitiesMap: new Map(),
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
  addEntity(state, entities) {
    state.entities.push(...entities);
  },
});

const actions = defineActions({

  async GetListNotBelongToGroup(context, groupId: number): Promise<Route[]> {
    const { state } = actionContext(context);
    const entities = await state.service.GetList(groupId, false);

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
    commit.setGroupId(pairs[0].groupId);
    return createdEntity;
  },
  async Delete(context, pairs: GroupRoute[]): Promise<void> {
    const { state, commit } = actionContext(context);

    await state.service.Delete(pairs);
    commit.deleteEntity(pairs);
    commit.setGroupId(pairs[0].groupId);
  },
});

const getters = defineGetters<State>()({
  getEntities(state): Route[] {
    return state.entities;
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
