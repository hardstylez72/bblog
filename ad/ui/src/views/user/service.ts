import DefaultService from '../base/services/default';

export interface User {
  id: number;
  externalId: string;
  isSystem: boolean;
  description: string;
}

export default class UserService extends DefaultService<User> {

}
