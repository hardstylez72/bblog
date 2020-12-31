import DefaultService from '../base/services/default';

export interface Route {
  id: number;
  route: string;
  method: string;
  description: string;
}

export default class RouteService extends DefaultService<Route> {

}
