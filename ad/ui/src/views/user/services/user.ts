import { makeRequest, Request } from '@/views/base/services/utils/requester';
import DefaultService, { T } from '../../base/services/default';

export interface User {
  id: number;
  externalId: string;
  isSystem: boolean;
  description: string;
}

export default class UserService extends DefaultService<User> {
  GetById(id: number): Promise<User> {
    const req: Request = {
      data: { id },
      method: this.methodPost,
      url: `${this.baseUrl}/get`,
    };
    return makeRequest(req);
  }
}
