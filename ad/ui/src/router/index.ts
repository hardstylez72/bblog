import Vue from 'vue';
import VueRouter, { RouteConfig, Route } from 'vue-router';
import Home from '../views/main-page.vue';
import Group from '../views/group-page/page.vue';

Vue.use(VueRouter);
export const generateItemPageProps = (route: Route) => ({
  id: route.params.id,
});

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/group/:id',
    name: 'Group',
    component: Group,
    props: generateItemPageProps,
  },
];

const router = new VueRouter({
  mode: 'history',
  routes,
});

export default router;
