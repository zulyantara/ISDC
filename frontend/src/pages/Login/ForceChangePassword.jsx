import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Card, Form, Input, Button, Typography, Space, message, Alert } from 'antd'
import { LockOutlined, KeyOutlined, SafetyOutlined, WarningOutlined } from '@ant-design/icons'
import { useAuth } from '../../context/AuthContext'

const { Title, Text } = Typography

export default function ForceChangePassword() {
  const [loading, setLoading] = useState(false)
  const { changePassword, user } = useAuth()
  const navigate = useNavigate()

  const onFinish = async (values) => {
    if (values.new_password !== values.confirm_password) {
      message.error('Konfirmasi password tidak cocok'); return
    }
    setLoading(true)
    try {
      const result = await changePassword(values.old_password, values.new_password)
      if (result.success) {
        message.success('Password berhasil diubah!')
        navigate('/dashboard')
      } else {
        message.error(result.message || 'Gagal mengubah password')
      }
    } catch (err) {
      message.error(err.message || 'Terjadi kesalahan')
    } finally { setLoading(false) }
  }

  return (
    <div style={{
      minHeight: '100vh',
      display: 'flex',
      background: 'linear-gradient(135deg, #f59e0b 0%, #f97316 100%)',
      alignItems: 'center', justifyContent: 'center', padding: 20,
    }}>
      <Card style={{ width: 440, borderRadius: 16, boxShadow: '0 20px 60px rgba(0,0,0,0.3)', border: 'none' }}>
        <div style={{ textAlign: 'center', marginBottom: 24 }}>
          <div style={{
            width: 64, height: 64, borderRadius: 12,
            background: '#fff7e6', display: 'inline-flex',
            alignItems: 'center', justifyContent: 'center', marginBottom: 12,
          }}>
            <WarningOutlined style={{ fontSize: 32, color: '#f59e0b' }} />
          </div>
          <Title level={4} style={{ margin: 0 }}>Ubah Password</Title>
          <Text type="secondary">Selamat datang, {user?.user_name}</Text>
        </div>

        <Alert
          message="Anda masih menggunakan password default!"
          description="Silakan ubah password sebelum melanjutkan."
          type="warning" showIcon style={{ marginBottom: 24, borderRadius: 8 }}
        />

        <Form onFinish={onFinish} size="large" autoComplete="off">
          <Form.Item name="old_password" rules={[{ required: true, message: 'Masukkan password saat ini' }]}>
            <Input.Password prefix={<LockOutlined />} placeholder="Password saat ini (default)" />
          </Form.Item>
          <Form.Item name="new_password" rules={[
            { required: true }, { min: 8, message: 'Minimal 8 karakter' },
          ]}>
            <Input.Password prefix={<KeyOutlined />} placeholder="Password baru" />
          </Form.Item>
          <Form.Item name="confirm_password" rules={[
            { required: true },
            ({ getFieldValue }) => ({
              validator(_, v) {
                return !v || getFieldValue('new_password') === v
                  ? Promise.resolve() : Promise.reject(new Error('Tidak cocok'))
              },
            }),
          ]}>
            <Input.Password prefix={<KeyOutlined />} placeholder="Ulangi password baru" />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" loading={loading} block
              style={{ height: 46, borderRadius: 8, fontWeight: 600 }}>
              Ubah Password & Lanjutkan
            </Button>
          </Form.Item>
        </Form>
        <div style={{ textAlign: 'center' }}>
          <Button type="link" onClick={() => navigate('/login')}>← Kembali ke Login</Button>
        </div>
      </Card>
    </div>
  )
}
