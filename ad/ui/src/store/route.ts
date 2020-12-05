/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { moduleActionContext } from './index';
import Service from '../services/service';
import RouteService, { Route } from '../services/route';

export interface State {
  service: Service<Route>;
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new RouteService({ host: '', baseUrl: '/api/v1/route' }),
  } as State,
  getters: {},
  mutations: {},
  actions: {
    GetList(context): Promise<Route[]> {
      const { state } = actionContext(context);
      return state.service.GetList();
    },
    Create(context, route: Route): Promise<Route> {
      const { state } = actionContext(context);
      return state.service.Create(route);
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
