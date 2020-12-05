import axios, { Method } from "axios";
import { uuid } from "uuidv4";

const instance = axios.create();

export interface Request {
  method: Method;
  url: string;
  data: any;
  headers?: any;
}

export const makeRequest = (req: Request) => requester(req)

const requester = (req: Request):Promise<any> => {
  const headers = {
    "x-request-id": uuid(),
    ...req.headers
  };

  return instance({
    method: req.method,
    url: process.env.VUE_APP_SERVER_HOST + req.url,
    data: req.data,
    headers: headers,

    withCredentials: true,
    timeout: 30000
  })
    .then((res: { data: any; }) => {
      return res.data;
    })
    .catch(async (err: Error) => {
      console.error(err);
      throw err;
    });
};
