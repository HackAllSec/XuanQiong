import axios from 'axios';

const api = axios.create({
  baseURL: '',
  timeout: 5000,
});

export default api;