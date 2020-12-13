/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import GroupRouteService, { GroupRoute } from '@/views/group-route/service';
import { Service } from '@/views/route/service';
import { moduleActionContext } from '../base/store';

export interface State{
  service: GroupRouteService;
  entities: Service[];
  routesNotBelongToGroup: Service[];
  entitiesMap: Map<number, Service>;
  groupId: number;
}

const state1 = {
  service: new GroupRouteService({ host: '', baseUrl: '/api/v1/group/route' }),
  entities: [],
  routesNotBelongToGroup: [],
  entitiesMap: new Map(),
  groupId: -1,
} as State;

const mutations = defineMutations < State >()({

  setGroupId(state, groupId: number) {
    state.groupId = groupId;
  },
  setEntities(state, entities: Service[]) {
    state.entities = entities;
  },
  deleteEntity(state, pairs: GroupRoute[]) {
    state.entities = state.entities.filter((r: Service) => {
      const exist = pairs.some((pair) => pair.routeId === r.id);
      return !exist;
    });
  },
  setRoutesNotBelongToGroup(state, entities: Service[]) {
    state.routesNotBelongToGroup = entities;
  },
  deleteRoutesNotBelongToGroup(state, pairs: GroupRoute[]) {
    state.routesNotBelongToGroup = state.routesNotBelongToGroup.filter((r: Service) => {
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

  async GetListNotBelongToGroup(context, groupId: number): Promise<Service[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(groupId, false);
    commit.setRoutesNotBelongToGroup(entities);
    commit.setGroupId(groupId);
    return entities;
  },
  async GetList(context, payload: {groupId: number; belongToGroup: boolean}): Promise<Service[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList(payload.groupId, payload.belongToGroup);

    if (payload.belongToGroup) {
      commit.setEntities(entities);
      commit.setGroupId(payload.groupId);
    }

    return entities;
  },
  async Create(context, pairs: GroupRoute[]): Promise<Service[]> {
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
  getEntities(state): Service[] {
    return state.entities;
  },
  getRoutesNotBelongToGroup(state): Service[] {
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
