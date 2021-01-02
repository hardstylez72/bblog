import { Method } from 'axios';

import { Group } from '@/views/group/services/group';
import { Route } from '@/views/route/service';
import { makeRequest, Request } from '../../base/services/utils/requester';

export interface UserRoute {
  routeId: number;
  userId: number;
}

interface Options {
  host: string;
  baseUrl: string;
}

export default class UserGroupService {
  private options: Options

  private readonly methodPost: Method = 'POST'

  private readonly baseUrl: string;

  constructor(options: Options) {
    this.options = options;
    this.baseUrl = `${this.options.host}${this.options.baseUrl}`;
  }

  Create(t: UserRoute[]): Promise<Route[]> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/create`,
    };
    return makeRequest(req);
  }

  Delete(t: UserRoute[]): Promise<void> {
    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.baseUrl}/delete`,
    };
    return makeRequest(req);
  }

  GetList(userId: number, belongToUser: boolean): Promise<Route[]> {
    const req: Request = {
      data: { userId, belongToUser },
      method: this.methodPost,
      url: `${this.baseUrl}/list`,
    };
    return makeRequest(req);
  }
}
