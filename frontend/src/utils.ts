import { useRouter } from 'vue-router';
import { jwtDecode } from 'jwt-decode'
import { clearAuthSession } from './auth';
export function formatDate(datetime) {
  const date = new Date(datetime);
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  return `${year}-${month}-${day}`;
}

export function checkLogin() {
  const router = useRouter();
  if (sessionStorage.getItem('force_password_change') === '1') {
    router.push('/modifypwd');
    return;
  }
  const token = sessionStorage.getItem('token');
  if (token) {
      try {
        const decodedToken = jwtDecode(token)
        const currentTime = Math.floor(Date.now() / 1000)
        if (currentTime > decodedToken.exp) {
            clearAuthSession()
            location.reload();
            return;
        }
      } catch {
        clearAuthSession()
        location.reload();
        return;
      }
  } else {
    sessionStorage.removeItem('force_password_change')
    router.push('/login');
  }
}
