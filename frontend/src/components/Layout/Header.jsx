import { Layout, Dropdown, Avatar, Space, Typography, Badge } from 'antd'
import { UserOutlined, LogoutOutlined, MenuFoldOutlined, MenuUnfoldOutlined, BellOutlined } from '@ant-design/icons'
import { useAuth } from '../../context/AuthContext'

const { Header: AntHeader } = Layout
const { Text } = Typography

const levelMap = { 1: 'Admin', 2: 'Super Admin', 3: 'Kasir', 4: 'Staff', 5: 'Superuser', 6: 'Instruktur' }

export default function Header({ collapsed, onToggle }) {
  const { user, logout } = useAuth()

  const menuItems = [
    {
      key: 'info',
      label: (
        <div style={{ padding: '4px 0', minWidth: 160 }}>
          <Text strong style={{ display: 'block' }}>{user?.user_name}</Text>
          <Text type="secondary" style={{ fontSize: 12 }}>
            {levelMap[user?.user_level] || 'User'}
          </Text>
        </div>
      ),
      disabled: true,
    },
    { type: 'divider' },
    { key: 'logout', icon: <LogoutOutlined />, label: 'Logout', onClick: logout },
  ]

  return (
    <AntHeader style={{
      padding: '0 24px',
      background: '#fff',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'space-between',
      boxShadow: '0 1px 4px rgba(0,0,0,0.05)',
      position: 'sticky',
      top: 0,
      zIndex: 99,
      height: 56,
    }}>
      <div onClick={onToggle} style={{ fontSize: 18, cursor: 'pointer', color: '#555' }}>
        {collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
      </div>

      <Space size={16}>
        <Badge dot>
          <BellOutlined style={{ fontSize: 18, color: '#888', cursor: 'pointer' }} />
        </Badge>
        <Dropdown menu={{ items: menuItems }} placement="bottomRight" trigger={['click']}>
          <Space style={{ cursor: 'pointer' }}>
            <Avatar size={32} icon={<UserOutlined />}
              style={{ background: 'linear-gradient(135deg, #1677ff, #4096ff)' }} />
            <Text style={{ fontSize: 13 }}>{user?.user_name}</Text>
          </Space>
        </Dropdown>
      </Space>
    </AntHeader>
  )
}
