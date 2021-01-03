import { makeRequest, Request } from '@/views/base/services/utils/requester';
import { Entity } from '@/views/base/services/entity';
import DefaultService from '../../base/services/default';

export interface Group extends Entity {
  id: number;
  code: string;
  description: string;
}

export default class GroupService extends DefaultService<Group> {
  GetById(id: number): Promise<Group> {
    const req: Request = {
      data: { id },
      method: this.methodPost,
      url: `${this.baseUrl}/get`,
    };
    return makeRequest(req);
  }
}
