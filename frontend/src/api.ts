import axios from 'axios';

const api = axios.create({
  baseURL: '',
  timeout: 5000,
});

api.interceptors.request.use((config) => {
  const headers: any = config.headers || {};
  const explicitAuthorization =
    (typeof headers.get === 'function' && (headers.get('Authorization') || headers.get('authorization'))) ||
    headers.Authorization ||
    headers.authorization;
  const sessionToken = sessionStorage.getItem('token');

  let token = '';
  if (typeof explicitAuthorization === 'string' && explicitAuthorization.startsWith('Bearer ')) {
    token = explicitAuthorization.slice('Bearer '.length);
  } else if (sessionToken) {
    token = sessionToken;
  }

  if (token) {
    if (typeof headers.set === 'function') {
      headers.set('X-Auth-Token', token);
      headers.delete('Authorization');
      headers.delete('authorization');
    } else {
      headers['X-Auth-Token'] = token;
      delete headers.Authorization;
      delete headers.authorization;
    }
  }

  config.headers = headers;
  return config;
});

export default api;
