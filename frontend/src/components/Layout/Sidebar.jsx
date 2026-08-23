import { useNavigate, useLocation } from 'react-router-dom'
import { Layout, Menu, Typography } from 'antd'
import {
  DashboardOutlined, UserAddOutlined, TeamOutlined,
  FileProtectOutlined, SettingOutlined, BookOutlined,
  EnvironmentOutlined, AimOutlined, FileTextOutlined,
  TrophyOutlined, SafetyOutlined, SafetyCertificateOutlined,
} from '@ant-design/icons'
import { useAuth } from '../../context/AuthContext'

const { Sider } = Layout
const { Text } = Typography

const allMenuItems = [
  {
    key: '/dashboard',
    icon: <DashboardOutlined />,
    label: 'Dashboard',
    menuUrl: 'welcome',
  },
  {
    key: '/pendaftaran',
    icon: <UserAddOutlined />,
    label: 'Pendaftaran',
    menuUrl: 'pendaftaran',
  },
  {
    key: '/peserta',
    icon: <TeamOutlined />,
    label: 'Peserta',
    menuUrl: 'petugas',
  },
  {
    key: '/sertifikat',
    icon: <FileProtectOutlined />,
    label: 'Sertifikat',
    menuUrl: 'sertifikat',
  },
  {
    type: 'divider',
  },
  {
    key: 'master',
    icon: <SettingOutlined />,
    label: 'Master Data',
    children: [
      { key: '/master/kelas', icon: <BookOutlined />, label: 'Kelas', menuUrl: 'mt_kelas' },
      { key: '/master/user', icon: <TeamOutlined />, label: 'User', menuUrl: 'mt_user' },
      { key: '/master/area', icon: <EnvironmentOutlined />, label: 'Area', menuUrl: 'mt_area' },
      { key: '/master/nilai-lulus', icon: <AimOutlined />, label: 'Nilai Lulus', menuUrl: 'mt_nilai_lulus' },
      { key: '/master/soal', icon: <FileTextOutlined />, label: 'Soal', menuUrl: 'ujian' },
      { key: '/master/rbac', icon: <SafetyCertificateOutlined />, label: 'RBAC / Hak Akses', menuUrl: 'mt_user' },
    ],
  },
]

export default function Sidebar({ collapsed }) {
  const navigate = useNavigate()
  const location = useLocation()
  const { accessibleMenus } = useAuth()

  const filterMenu = (items) => {
    return items
      .map(item => {
        if (item.type === 'divider') return item
        if (item.children) {
          const filteredChildren = item.children.filter(child =>
            accessibleMenus.includes(child.menuUrl)
          )
          if (filteredChildren.length === 0) return null
          return { ...item, children: filteredChildren }
        }
        if (!accessibleMenus.includes(item.menuUrl)) return null
        return item
      })
      .filter(Boolean)
  }

  const filteredMenuItems = filterMenu(allMenuItems)

  const findMenuKey = (path) => {
    if (path.startsWith('/master/')) return path
    return '/' + path.split('/')[1]
  }

  const selectedKey = findMenuKey(location.pathname)
  const openKeys = location.pathname.startsWith('/master/') ? ['master'] : []

  return (
    <Sider
      trigger={null}
      collapsible
      collapsed={collapsed}
      width={240}
      style={{
        overflow: 'auto',
        height: '100vh',
        position: 'fixed',
        left: 0,
        top: 0,
        bottom: 0,
        zIndex: 100,
        background: 'linear-gradient(180deg, #001529 0%, #002140 100%)',
      }}
    >
      {/* Logo */}
      <div style={{
        height: 64,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        borderBottom: '1px solid rgba(255,255,255,0.08)',
        gap: 10,
        padding: '0 16px',
      }}>
        <div style={{
          width: 36,
          height: 36,
          borderRadius: 10,
          background: 'linear-gradient(135deg, #1677ff 0%, #4096ff 100%)',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          boxShadow: '0 4px 12px rgba(22,119,255,0.4)',
        }}>
          <SafetyOutlined style={{ fontSize: 18, color: '#fff' }} />
        </div>
        {!collapsed && (
          <div>
            <Text strong style={{ color: '#fff', fontSize: 16, display: 'block', lineHeight: 1.2 }}>
              ISDC
            </Text>
            <Text style={{ color: 'rgba(255,255,255,0.4)', fontSize: 10 }}>
              Safety Driving Center
            </Text>
          </div>
        )}
      </div>

      <Menu
        theme="dark"
        mode="inline"
        selectedKeys={[selectedKey]}
        defaultOpenKeys={openKeys}
        items={filteredMenuItems}
        onClick={({ key }) => navigate(key)}
        style={{
          background: 'transparent',
          borderRight: 0,
          padding: '8px 0',
        }}
      />
    </Sider>
  )
}
