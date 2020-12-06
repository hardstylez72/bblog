/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable import/no-cycle */

import { defineModule } from 'direct-vuex';
import { moduleActionContext } from './index';
import Service from '../services/default';
import RouteService, { Route } from '../services/route';

export interface State {
  service: Service<Route>;
  routes: Route[];
}

const module = defineModule({
  namespaced: true as true,
  state: {
    service: new RouteService({ host: '', baseUrl: '/api/v1/route' }),
    routes: [],
  } as State,
  getters: {
    getRoutes(state) {
      return state.routes;
    },
  },
  mutations: {
    setRoutes(state, routes: Route[]) {
      state.routes = routes;
    },
    deleteRoute(state, id: number) {
      state.routes = state.routes.filter((route) => route.id !== id);
    },
    addRoute(state, routes: Route) {
      state.routes.push(routes);
    },
  },
  actions: {
    async GetList(context): Promise<Route[]> {
      const { state, commit } = actionContext(context);
      const routes = await state.service.GetList();
      commit.setRoutes(routes);
      return routes;
    },
    async Create(context, route: Route): Promise<Route> {
      const { state, commit } = actionContext(context);
      const createdRoute = await state.service.Create(route);
      commit.addRoute(createdRoute);
      return createdRoute;
    },
    async Delete(context, id: number): Promise<void> {
      const { state, commit } = actionContext(context);

      await state.service.Delete(id);
      commit.deleteRoute(id);
    },
  },
});

export default module;

const actionContext = (context: any) => moduleActionContext(context, module);
