import type { LoginResponse } from './types'

const PERMISSIONS_KEY = 'permissions'
const ROLE_IDS_KEY = 'role_ids'
const ROLES_KEY = 'roles'

export function persistLoginSession(response: LoginResponse) {
  sessionStorage.setItem('token', response.token)
  sessionStorage.setItem('username', response.username)
  if (response.avatar) {
    sessionStorage.setItem('avatar', '/download/file?id=' + response.avatar)
  } else {
    sessionStorage.setItem('avatar', '/avatar.svg')
  }
  sessionStorage.setItem(PERMISSIONS_KEY, JSON.stringify(response.permissions || []))
  sessionStorage.setItem(ROLES_KEY, JSON.stringify(response.roles || []))
  sessionStorage.setItem(ROLE_IDS_KEY, JSON.stringify(response.role_ids || []))
  if (response.force_password_change) {
    sessionStorage.setItem('force_password_change', '1')
  } else {
    sessionStorage.removeItem('force_password_change')
  }
}

export function clearAuthSession() {
  sessionStorage.removeItem('token')
  sessionStorage.removeItem('username')
  sessionStorage.removeItem('avatar')
  sessionStorage.removeItem(PERMISSIONS_KEY)
  sessionStorage.removeItem(ROLES_KEY)
  sessionStorage.removeItem(ROLE_IDS_KEY)
  sessionStorage.removeItem('force_password_change')
}

export function getAuthToken(): string {
  return sessionStorage.getItem('token') || ''
}

export function getUploadHeaders(): Record<string, string> {
  const token = getAuthToken()
  if (!token) {
    return {}
  }
  return {
    'X-Auth-Token': token,
  }
}
