import { createContext, useContext, useState, useEffect, useMemo } from 'react'
import api from '../api/axios'

const AuthContext = createContext(null)

export function AuthProvider({ children }) {
  const [user, setUser] = useState(null)
  const [loading, setLoading] = useState(true)
  const [mustChangePassword, setMustChangePassword] = useState(false)
  const [permissions, setPermissions] = useState([])

  useEffect(() => {
    const storedUser = localStorage.getItem('user')
    const token = localStorage.getItem('token')
    const storedMustChange = localStorage.getItem('must_change_password')

    if (storedUser && token) {
      setUser(JSON.parse(storedUser))
      setMustChangePassword(storedMustChange === 'true')
      fetchPermissions()
    }
    setLoading(false)
  }, [])

  const fetchPermissions = async () => {
    try {
      const res = await api.get('/auth/permissions')
      if (res.status) setPermissions(res.data || [])
    } catch {
      setPermissions([])
    }
  }

  const login = async (userId, password) => {
    const response = await api.post('/auth/login', {
      user_id: userId,
      user_pwd: password,
    })

    if (response.status) {
      const { token, must_change_password, ...userData } = response.data
      localStorage.setItem('token', token)
      localStorage.setItem('user', JSON.stringify(userData))
      localStorage.setItem('must_change_password', String(must_change_password))
      setUser(userData)
      setMustChangePassword(!!must_change_password)
      if (!must_change_password) {
        await fetchPermissions()
      }
      return { success: true, mustChangePassword: !!must_change_password }
    }

    return { success: false, message: response.message }
  }

  const changePassword = async (oldPassword, newPassword) => {
    const response = await api.post('/auth/change-password', {
      old_password: oldPassword,
      new_password: newPassword,
    })

    if (response.status) {
      setMustChangePassword(false)
      localStorage.setItem('must_change_password', 'false')
      await fetchPermissions()
      return { success: true }
    }

    return { success: false, message: response.message }
  }

  const logout = async () => {
    try { await api.post('/auth/logout') } catch {}
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    localStorage.removeItem('must_change_password')
    setUser(null)
    setMustChangePassword(false)
    setPermissions([])
  }

  // Check permission helper
  const hasPermission = (menuUrl, action = 'view') => {
    const perm = permissions.find(p => p.menu_url === menuUrl)
    if (!perm) return false
    switch (action) {
      case 'view': return perm.view === 1
      case 'insert': return perm.insert === 1
      case 'update': return perm.update === 1
      case 'delete': return perm.delete === 1
      case 'proses': return perm.proses === 1
      default: return false
    }
  }

  // Get accessible menu URLs
  const accessibleMenus = useMemo(() => {
    return permissions.filter(p => p.view === 1).map(p => p.menu_url)
  }, [permissions])

  const value = {
    user, loading, login, logout, changePassword,
    mustChangePassword, permissions, hasPermission, accessibleMenus,
    isAuthenticated: !!user,
  }

  return (
    <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
  )
}

export function useAuth() {
  const context = useContext(AuthContext)
  if (!context) throw new Error('useAuth must be used within an AuthProvider')
  return context
}
