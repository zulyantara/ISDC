import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { Card, Form, Input, Button, Typography, Space, message } from 'antd'
import { UserOutlined, LockOutlined } from '@ant-design/icons'
import { useAuth } from '../../context/AuthContext'

const { Title, Text } = Typography

export default function Login() {
  const [loading, setLoading] = useState(false)
  const { login, isAuthenticated } = useAuth()
  const navigate = useNavigate()

  if (isAuthenticated) { navigate('/dashboard'); return null }

  const onFinish = async (values) => {
    setLoading(true)
    try {
      const result = await login(values.user_id, values.pwd)
      if (result.success) {
        if (result.mustChangePassword) {
          message.warning('Silakan ubah password default Anda')
          navigate('/ubah-password')
        } else {
          message.success('Selamat datang!')
          navigate('/dashboard')
        }
      } else {
        message.error(result.message || 'Login gagal')
      }
    } catch (err) {
      message.error(err.message || 'Terjadi kesalahan')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div style={{
      minHeight: '100vh',
      display: 'flex',
      background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
      alignItems: 'center',
      justifyContent: 'center',
      padding: 20,
    }}>
      <Card style={{
        width: 420,
        borderRadius: 16,
        boxShadow: '0 20px 60px rgba(0,0,0,0.3)',
        border: 'none',
      }}>
        <div style={{ textAlign: 'center', marginBottom: 32 }}>
          <img src="/logo_isdc.jpeg" alt="ISDC Logo"
            style={{ width: 80, height: 80, borderRadius: 16, objectFit: 'cover', marginBottom: 16 }} />
          <Title level={3} style={{ margin: 0 }}>ISDC</Title>
          <Text type="secondary">Indonesia Safety Driving Center</Text>
        </div>

        <Form onFinish={onFinish} size="large" autoComplete="off">
          <Form.Item name="user_id" rules={[{ required: true, message: 'Masukkan User ID' }]}>
            <Input prefix={<UserOutlined style={{ color: '#bbb' }} />} placeholder="User ID" />
          </Form.Item>
          <Form.Item name="pwd" rules={[{ required: true, message: 'Masukkan Password' }]}>
            <Input.Password prefix={<LockOutlined style={{ color: '#bbb' }} />} placeholder="Password" />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" loading={loading} block
              style={{ height: 46, borderRadius: 8, fontWeight: 600, fontSize: 15 }}>
              Masuk
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </div>
  )
}
