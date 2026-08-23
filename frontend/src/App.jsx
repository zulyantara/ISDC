import { Routes, Route, Navigate } from 'react-router-dom'
import { Spin } from 'antd'
import { useAuth } from './context/AuthContext'
import MainLayout from './components/Layout/MainLayout'
import Login from './pages/Login/Login'
import Dashboard from './pages/Dashboard/Dashboard'
import ListPendaftaran from './pages/Pendaftaran/ListPendaftaran'
import FormPendaftaran from './pages/Pendaftaran/FormPendaftaran'
import ForceChangePassword from './pages/Login/ForceChangePassword'
import ListPeserta from './pages/Peserta/ListPeserta'
import FormUjian from './pages/Ujian/FormUjian'
import ListSertifikat from './pages/Sertifikat/ListSertifikat'
import MasterKelas from './pages/Master/MasterKelas'
import MasterUser from './pages/Master/MasterUser'
import MasterArea from './pages/Master/MasterArea'
import MasterNilaiLulus from './pages/Master/MasterNilaiLulus'
import MasterSoal from './pages/Master/MasterSoal'
import MasterRBAC from './pages/Master/MasterRBAC'

// Placeholder pages - will be implemented in later phases
function PlaceholderPage({ title }) {
  return (
    <div style={{ textAlign: 'center', padding: '60px 0' }}>
      <h2>{title}</h2>
      <p style={{ color: '#999', marginTop: 8 }}>Halaman ini akan segera diimplementasi</p>
    </div>
  )
}

// Protected route wrapper
function ProtectedRoute({ children }) {
  const { isAuthenticated, loading, mustChangePassword } = useAuth()

  if (loading) {
    return (
      <div style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', height: '100vh' }}>
        <Spin size="large" />
      </div>
    )
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" replace />
  }

  // Force change password if still using default
  if (mustChangePassword) {
    return <Navigate to="/ubah-password" replace />
  }

  return <MainLayout>{children}</MainLayout>
}

export default function App() {
  return (
    <Routes>
      {/* Public routes */}
      <Route path="/login" element={<Login />} />
      <Route path="/ubah-password" element={<ForceChangePassword />} />

      {/* Protected routes */}
      <Route path="/dashboard" element={
        <ProtectedRoute><Dashboard /></ProtectedRoute>
      } />

      {/* Pendaftaran */}
      <Route path="/pendaftaran" element={
        <ProtectedRoute><ListPendaftaran /></ProtectedRoute>
      } />
      <Route path="/pendaftaran/tambah" element={
        <ProtectedRoute><FormPendaftaran /></ProtectedRoute>
      } />
      <Route path="/pendaftaran/edit/:id" element={
        <ProtectedRoute><FormPendaftaran /></ProtectedRoute>
      } />

      {/* Peserta */}
      <Route path="/peserta" element={
        <ProtectedRoute><ListPeserta /></ProtectedRoute>
      } />

      {/* Ujian */}
      <Route path="/ujian/:id" element={
        <ProtectedRoute><FormUjian /></ProtectedRoute>
      } />

      {/* Sertifikat */}
      <Route path="/sertifikat" element={
        <ProtectedRoute><ListSertifikat /></ProtectedRoute>
      } />

      {/* Master Data */}
      <Route path="/master/kelas" element={
        <ProtectedRoute><MasterKelas /></ProtectedRoute>
      } />
      <Route path="/master/user" element={
        <ProtectedRoute><MasterUser /></ProtectedRoute>
      } />
      <Route path="/master/area" element={
        <ProtectedRoute><MasterArea /></ProtectedRoute>
      } />
      <Route path="/master/nilai-lulus" element={
        <ProtectedRoute><MasterNilaiLulus /></ProtectedRoute>
      } />	      <Route path="/master/soal" element={
        <ProtectedRoute><MasterSoal /></ProtectedRoute>
      } />
      <Route path="/master/rbac" element={
        <ProtectedRoute><MasterRBAC /></ProtectedRoute>
      } />

      {/* Default redirect */}
      <Route path="/" element={<Navigate to="/dashboard" replace />} />
      <Route path="*" element={<Navigate to="/dashboard" replace />} />
    </Routes>
  )
}
