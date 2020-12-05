import {makeRequest, Request} from './requester'
import {Method} from "axios";

interface Route {
  id: number;
  route: string;
  method: string;
  description: string;
}

interface Options {
  host: string
  baseUrl: string
}

export default class RouteService implements Service<Route> {
  private options: Options
  readonly methodPost: Method = 'POST'
  constructor(options: Options) {
    this.options = options;
  }

  Create(t: Route): Promise<Route> {

    const req: Request = {
      data: t,
      method: this.methodPost,
      url: `${this.options.baseUrl}/create`
    }
    return makeRequest(req)
  }

  Delete(id: number): Promise<void> {
    const req: Request = {
      data: {id: id},
      method: this.methodPost,
      url: `${this.options.baseUrl}/delete`
    }
    return makeRequest(req)
  }

  GetList(): Promise<Route[]> {
    const req: Request = {
      data: {},
      method: this.methodPost,
      url: `${this.options.baseUrl}/list`
    }
    return makeRequest(req)
  }

  Update(t: Route): Promise<Route> {
    // todo: implement@!!!
    const req: Route = {
      description: '',
      id:1,
      method: this.methodPost,
      route: ''
    }
    return Promise.resolve(req);
  }
}
