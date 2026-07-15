import axios from 'axios';

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

export default api;
