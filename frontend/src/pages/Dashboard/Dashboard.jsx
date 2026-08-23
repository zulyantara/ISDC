import { useState, useEffect } from 'react'
import { Row, Col, Card, Typography, Table, Tag, Spin, message, Space, Statistic } from 'antd'
import {
  UserAddOutlined, TeamOutlined, TrophyOutlined,
  SafetyOutlined, CalendarOutlined, ClockCircleOutlined,
} from '@ant-design/icons'
import { useAuth } from '../../context/AuthContext'
import api from '../../api/axios'
import { formatDate } from '../../utils/helpers'

const { Title, Text } = Typography

export default function Dashboard() {
  const { user } = useAuth()
  const [loading, setLoading] = useState(true)
  const [stats, setStats] = useState({ totalPendaftar: 0, totalPeserta: 0, rataRata: 0 })
  const [recentDaftar, setRecentDaftar] = useState([])

  useEffect(() => { fetchData() }, [])

  const fetchData = async () => {
    setLoading(true)
    try {
      const [daftarRes, pesertaRes] = await Promise.all([
        api.get('/daftar').catch(() => ({ status: false, data: [] })),
        api.get('/peserta').catch(() => ({ status: false, data: [] })),
      ])

      const daftar = daftarRes.status ? (daftarRes.data || []) : []
      const peserta = pesertaRes.status ? (pesertaRes.data || []) : []
      const withScore = peserta.filter(p => p.praktek_hasil && p.praktek_hasil !== '')
      const avg = withScore.length > 0
        ? withScore.reduce((s, p) => s + parseInt(p.praktek_hasil || 0), 0) / withScore.length
        : 0

      setStats({ totalPendaftar: daftar.length, totalPeserta: peserta.length, rataRata: Math.round(avg) })
      setRecentDaftar(daftar.slice(0, 5))
    } catch { message.error('Gagal memuat dashboard') }
    finally { setLoading(false) }
  }

  const statCards = [
    {
      title: 'Total Pendaftar',
      value: stats.totalPendaftar,
      icon: <UserAddOutlined style={{ fontSize: 28 }} />,
      color: '#fff',
      bg: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
      shadow: '0 8px 32px rgba(102,126,234,0.4)',
    },
    {
      title: 'Total Peserta',
      value: stats.totalPeserta,
      icon: <TeamOutlined style={{ fontSize: 28 }} />,
      color: '#fff',
      bg: 'linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)',
      shadow: '0 8px 32px rgba(67,233,123,0.4)',
    },
    {
      title: 'Rata-rata Nilai',
      value: stats.rataRata,
      suffix: '/ 100',
      icon: <TrophyOutlined style={{ fontSize: 28 }} />,
      color: '#fff',
      bg: 'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)',
      shadow: '0 8px 32px rgba(245,87,108,0.4)',
    },
  ]

  return (
    <Spin spinning={loading}>
      {/* Welcome Header */}
      <div className="fade-in" style={{
        background: 'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
        borderRadius: 16,
        padding: '28px 32px',
        marginBottom: 24,
        color: '#fff',
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        flexWrap: 'wrap',
        gap: 16,
      }}>
        <div>
          <Title level={3} style={{ margin: 0, color: '#fff' }}>
            <SafetyOutlined style={{ marginRight: 10 }} />
            Selamat Datang, {user?.user_name}
          </Title>
          <Text style={{ color: 'rgba(255,255,255,0.8)', fontSize: 14 }}>
            Indonesia Safety Driving Center — Dashboard
          </Text>
        </div>
        <Space style={{ color: 'rgba(255,255,255,0.7)', fontSize: 13 }}>
          <CalendarOutlined />
          <span>{new Date().toLocaleDateString('id-ID', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })}</span>
        </Space>
      </div>

      {/* Stat Cards */}
      <Row gutter={[16, 16]} style={{ marginBottom: 24 }}>
        {statCards.map((card, i) => (
          <Col xs={24} sm={12} lg={8} key={i}>
            <div
              className="stat-card fade-in"
              style={{
                background: card.bg,
                borderRadius: 16,
                padding: '24px 28px',
                color: '#fff',
                boxShadow: card.shadow,
                cursor: 'default',
              }}
            >
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                <div>
                  <Text style={{ color: 'rgba(255,255,255,0.85)', fontSize: 14, fontWeight: 500 }}>
                    {card.title}
                  </Text>
                  <div style={{ fontSize: 32, fontWeight: 700, lineHeight: 1.2, marginTop: 4 }}>
                    {card.value?.toLocaleString('id-ID')} {card.suffix || ''}
                  </div>
                </div>
                <div style={{
                  width: 56, height: 56, borderRadius: 14,
                  background: 'rgba(255,255,255,0.2)',
                  display: 'flex', alignItems: 'center', justifyContent: 'center',
                  fontSize: 24,
                }}>
                  {card.icon}
                </div>
              </div>
            </div>
          </Col>
        ))}
      </Row>

      {/* Recent Table */}
      <Card
        title={
          <Space>
            <ClockCircleOutlined style={{ color: '#1677ff' }} />
            <Text strong>Pendaftaran Terbaru</Text>
          </Space>
        }
        bodyStyle={{ padding: 0 }}
      >
        <Table
          columns={[
            {
              title: 'No. Pendaftaran',
              dataIndex: 'peserta_id',
              render: t => <Text code style={{ fontSize: 12 }}>{t}</Text>,
            },
            { title: 'Nama', dataIndex: 'nama', render: t => <Text strong>{t}</Text> },
            { title: 'Kelas', dataIndex: 'kelas', render: t => t || '-' },
            { title: 'Tanggal', dataIndex: 'tgl_daftar', render: t => formatDate(t) },
            { title: 'Status', key: 's', render: () => <Tag color="green">Terdaftar</Tag> },
          ]}
          dataSource={recentDaftar}
          rowKey="peserta_id"
          pagination={false}
          size="small"
        />
      </Card>
    </Spin>
  )
}
