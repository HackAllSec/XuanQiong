import { useRouter } from 'vue-router';
import { jwtDecode } from 'jwt-decode'
export function formatDate(datetime) {
  const date = new Date(datetime);
  const year = date.getFullYear();
  const month = (date.getMonth() + 1).toString().padStart(2, '0');
  const day = date.getDate().toString().padStart(2, '0');
  return `${year}-${month}-${day}`;
}

export function checkLogin() {
  const router = useRouter();
  const token = sessionStorage.getItem('token');
  if (token) {
      try {
        const decodedToken = jwtDecode(token)
        const currentTime = Math.floor(Date.now() / 1000)
        if (currentTime > decodedToken.exp) {
            sessionStorage.removeItem('token')
            sessionStorage.removeItem('username')
            sessionStorage.removeItem('avatar')
            location.reload();
            return;
        }
      } catch (error) {
        console.log(error)
        sessionStorage.removeItem('token')
        sessionStorage.removeItem('username')
        sessionStorage.removeItem('avatar')
        location.reload();
        return;
      }
  } else {
    router.push('/login');
  }
}