import DefaultService from '../base/services/default';

export interface Group {
  id: number;
  code: string;
  description: string;
}

export default class GroupService extends DefaultService<Group> {

}
