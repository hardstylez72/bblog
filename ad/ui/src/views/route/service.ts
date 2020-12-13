import DefaultService from '../base/services/default';

export interface Service {
  id: number;
  route: string;
  method: string;
  description: string;
}

export default class RouteService extends DefaultService<Service> {

}
