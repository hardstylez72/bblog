import { Entity } from '@/views/base/services/entity';
import DefaultService from '../base/services/default';

export interface Route extends Entity {
  id: number;
  route: string;
  method: string;
  description: string;
}

export default class RouteService extends DefaultService<Route> {

}
