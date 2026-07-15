export interface LoginPayload {
    username: string;
    password: string;
  }
  
  export interface LoginResponse {
    msg: string;
    code?: number;
    token: string;
    username: string;
    avatar?: string;
    times?: number;
    force_password_change?: boolean;
  }
