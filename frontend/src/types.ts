export interface LoginPayload {
    username: string;
    password: string;
  }
  
export interface LoginResponse {
    msg: string;
    token: string;
    username: string;
}