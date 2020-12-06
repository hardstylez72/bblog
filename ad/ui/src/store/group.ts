/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import {
  defineActions, defineModule, defineMutations, defineGetters,
} from 'direct-vuex';
import { moduleActionContext } from './index';
import Service from '../services/default';
import GroupService, { Group } from '../services/group';

export interface State<T>{
  service: Service<T>;
  entities: T[];
}

const state1 = {
  service: new GroupService({ host: '', baseUrl: '/api/v1/group' }),
  entities: [],
} as State<Group>;

const mutations = defineMutations < State < Group >>()({
  setEntities(state, entities) {
    state.entities = entities;
  },
  deleteEntity(state, id: number) {
    state.entities = state.entities.filter((route) => route.id !== id);
  },
  addEntity(state, entities) {
    state.entities.push(entities);
  },
});

const actions = defineActions({
  async GetList(context): Promise<Group[]> {
    const { state, commit } = actionContext(context);
    const entities = await state.service.GetList();
    commit.setEntities(entities);
    return entities;
  },
  async Create(context, route: Group): Promise<Group> {
    const { state, commit } = actionContext(context);
    const createdEntity = await state.service.Create(route);
    commit.addEntity(createdEntity);
    return createdEntity;
  },
  async Delete(context, id: number): Promise<void> {
    const { state, commit } = actionContext(context);

    await state.service.Delete(id);
    commit.deleteEntity(id);
  },
});

const getters = defineGetters<State<Group>>()({
  getEntities(state) {
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
