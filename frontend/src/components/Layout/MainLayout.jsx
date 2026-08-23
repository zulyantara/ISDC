import { useState } from 'react'
import { Layout, theme } from 'antd'
import { useLocation } from 'react-router-dom'
import Sidebar from './Sidebar'
import Header from './Header'

const { Content } = Layout

export default function MainLayout({ children }) {
  const [collapsed, setCollapsed] = useState(false)
  const location = useLocation()
  const { token: { borderRadiusLG } } = theme.useToken()

  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sidebar collapsed={collapsed} />
      <Layout style={{ marginLeft: collapsed ? 80 : 240, transition: 'all 0.2s', background: '#f0f2f5' }}>
        <Header collapsed={collapsed} onToggle={() => setCollapsed(!collapsed)} />
        <Content style={{ margin: 0, padding: 20 }}>
          <div style={{
            background: '#fff',
            borderRadius: borderRadiusLG,
            padding: 24,
            minHeight: 'calc(100vh - 56px - 40px)',
            boxShadow: '0 1px 2px rgba(0,0,0,0.03)',
          }}>
            {children}
          </div>
        </Content>
      </Layout>
    </Layout>
  )
}
