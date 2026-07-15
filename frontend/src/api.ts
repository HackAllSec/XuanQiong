import axios from 'axios';
import { clearAuthSession } from './auth';

const api = axios.create({
  baseURL: '',
  timeout: 5000,
});

api.interceptors.request.use((config) => {
  const headers: any = config.headers || {};
  const sessionToken = sessionStorage.getItem('token');
  const hasLegacyAuthHeader =
    typeof headers.get === 'function'
      ? headers.get('Authorization') || headers.get('authorization')
      : headers.Authorization || headers.authorization;

  if (hasLegacyAuthHeader) {
    throw new Error('Authorization header is not allowed. Use X-Auth-Token only.');
  }

  if (sessionToken) {
    if (typeof headers.set === 'function') {
      headers.set('X-Auth-Token', sessionToken);
    } else {
      headers['X-Auth-Token'] = sessionToken;
    }
  }

  config.headers = headers;
  return config;
});

function rejectSanitized(error: any) {
  return Promise.reject({
    status: error?.response?.status,
    code: error?.response?.data?.code,
    message: error?.response?.data?.msg || error?.message || 'Request failed',
  });
}

api.interceptors.response.use(
  (response) => {
    const code = response.data?.code;
    const msg = String(response.data?.msg || '').toLowerCase();
    if (code === 9 || (code === 0 && msg.includes('permission'))) {
      clearAuthSession();
    }
    return response;
  },
  (error) => {
    if (error?.response?.status === 401 || error?.response?.status === 403) {
      clearAuthSession();
    }
    return rejectSanitized(error);
  },
);

export default api;
