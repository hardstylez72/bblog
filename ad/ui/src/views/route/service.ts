import { Entity } from '@/views/base/services/entity';
import { makeRequest, Request } from '@/views/base/services/utils/requester';
import DefaultService from '../base/services/default';

export interface Route extends Entity {
  id: number;
  route: string;
  method: string;
  description: string;
}

export default class RouteService extends DefaultService<Route> {
  Update(t: Route): Promise<Route> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/update`,
    };
    return makeRequest(req);
  }
}
