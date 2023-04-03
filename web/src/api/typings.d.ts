import 'axios';

declare module 'axios' {
  interface AxiosResponse {
    msg: string;
    code: number;
    data: T;
    total: number;
    page: number;
    page_size: number;
  }
}
