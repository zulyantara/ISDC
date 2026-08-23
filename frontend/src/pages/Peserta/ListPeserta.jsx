import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import {
  Table, Input, Space, Tag, Typography, Card, Row, Col,
  Tooltip, Button, Descriptions, Modal, message, Empty, Avatar,
} from 'antd'
import {
  SearchOutlined, EyeOutlined, ReadOutlined, ReloadOutlined, UserOutlined,
} from '@ant-design/icons'
import api from '../../api/axios'
import { useAuth } from '../../context/AuthContext'
import { formatDate, formatRupiah, getKelaminText } from '../../utils/helpers'

const { Title, Text } = Typography

export default function ListPeserta() {
  const navigate = useNavigate()
  const { hasPermission } = useAuth()
  const canInsert = hasPermission('pendaftaran', 'insert')
  const [data, setData] = useState([])
  const [loading, setLoading] = useState(false)
  const [searchText, setSearchText] = useState('')
  const [detailModal, setDetailModal] = useState({ open: false, record: null })

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    setLoading(true)
    try {
      const res = await api.get('/peserta')
      if (res.status) setData(res.data || [])
    } catch {
      message.error('Gagal memuat data peserta')
    } finally {
      setLoading(false)
    }
  }

  const filteredData = data.filter(item => {
    const s = searchText.toLowerCase()
    return (
      (item.peserta_id || '').toLowerCase().includes(s) ||
      (item.nama || '').toLowerCase().includes(s) ||
      (item.kelas_nama || '').toLowerCase().includes(s)
    )
  })

  const getHasilTag = (hasil) => {
    if (!hasil || hasil === '') return <Tag>-</Tag>
    const val = parseInt(hasil)
    if (val >= 70) return <Tag color="success" style={{ borderRadius: 12, padding: '2px 10px' }}>{hasil}</Tag>
    if (val >= 50) return <Tag color="warning" style={{ borderRadius: 12, padding: '2px 10px' }}>{hasil}</Tag>
    return <Tag color="error" style={{ borderRadius: 12, padding: '2px 10px' }}>{hasil}</Tag>
  }

  const columns = [
    {
      title: 'No. Peserta',
      dataIndex: 'peserta_id',
      key: 'peserta_id',
      width: 160,
      render: (text) => <Text code style={{ fontSize: 12 }}>{text}</Text>,
    },
    {
      title: 'Nama',
      dataIndex: 'nama',
      key: 'nama',
      render: (text) => (
        <Space>
          <Avatar size={28} icon={<UserOutlined />} style={{ background: 'linear-gradient(135deg, #667eea, #764ba2)' }} />
          <Text strong>{text}</Text>
        </Space>
      ),
    },
    {
      title: 'Kelas',
      dataIndex: 'kelas_nama',
      key: 'kelas_nama',
      width: 200,
      render: (text) => text || '-',
    },
    {
      title: 'Nilai Praktek',
      dataIndex: 'praktek_hasil',
      key: 'praktek_hasil',
      width: 120,
      align: 'center',
      render: (val) => getHasilTag(val),
    },
    {
      title: 'Sertifikat',
      dataIndex: 'sertif_nomor',
      key: 'sertif_nomor',
      width: 160,
      render: (text) => text ? <Tag color="blue" style={{ borderRadius: 12 }}>{text}</Tag> : <Tag style={{ borderRadius: 12 }}>-</Tag>,
    },
    {
      title: 'Aksi',
      key: 'aksi',
      width: 150,
      align: 'center',
      render: (_, record) => (
        <Space size="small">
          <Tooltip title="Lihat Detail">
            <Button
              type="text"
              size="small"
              icon={<EyeOutlined />}
              onClick={() => setDetailModal({ open: true, record })}
              style={{ color: '#1677ff' }}
            />
          </Tooltip>
          {canInsert && (
            <Tooltip title="Input Ujian">
              <Button
                type="primary"
                size="small"
                icon={<ReadOutlined />}
                onClick={() => navigate(`/ujian/${record.peserta_id}`)}
                style={{ borderRadius: 6 }}
              />
            </Tooltip>
          )}
        </Space>
      ),
    },
  ]

  const record = detailModal.record

  return (
    <div>
      <Row justify="space-between" align="middle" style={{ marginBottom: 20 }}>
        <Col>
          <Title level={4} style={{ margin: 0 }}>👨‍🎓 Daftar Peserta</Title>
          <Text type="secondary" style={{ fontSize: 13 }}>Kelola data peserta pelatihan</Text>
        </Col>
      </Row>

      <Card
        style={{ borderRadius: 12, boxShadow: '0 1px 3px rgba(0,0,0,0.06)' }}
        styles={{ body: { padding: '16px 20px' } }}
      >
        <Row justify="space-between" style={{ marginBottom: 16 }}>
          <Col>
            <Input
              placeholder="Cari nama, ID, kelas..."
              prefix={<SearchOutlined style={{ color: '#bbb' }} />}
              value={searchText}
              onChange={(e) => setSearchText(e.target.value)}
              style={{ width: 300, borderRadius: 8 }}
              allowClear
            />
          </Col>
          <Col>
            <Button icon={<ReloadOutlined />} onClick={fetchData} style={{ borderRadius: 8 }}>
              Refresh
            </Button>
          </Col>
        </Row>

        <Table
          columns={columns}
          dataSource={filteredData}
          rowKey="peserta_id"
          loading={loading}
          pagination={{ pageSize: 10, showSizeChanger: true, showTotal: (t) => `Total ${t} data` }}
          size="middle"
          scroll={{ x: 1000 }}
          style={{ borderRadius: 8 }}
        />
      </Card>

      {/* Detail Modal */}
      <Modal
        title={
          <Space>
            <Avatar size={32} icon={<UserOutlined />} style={{ background: 'linear-gradient(135deg, #667eea, #764ba2)' }} />
            <span>Detail Peserta</span>
          </Space>
        }
        open={detailModal.open}
        onCancel={() => setDetailModal({ open: false, record: null })}
        footer={null}
        width={700}
        style={{ borderRadius: 12 }}
      >
        {record && (
          <Descriptions bordered column={2} size="small" style={{ marginTop: 16 }}>
            <Descriptions.Item label="No. Peserta" span={2}>
              <Text code>{record.peserta_id}</Text>
            </Descriptions.Item>
            <Descriptions.Item label="Nama">{record.nama}</Descriptions.Item>
            <Descriptions.Item label="Kelamin">{getKelaminText(record.kelamin_id)}</Descriptions.Item>
            <Descriptions.Item label="Kelas">{record.kelas_nama || '-'}</Descriptions.Item>
            <Descriptions.Item label="Biaya">{formatRupiah(record.biaya || 0)}</Descriptions.Item>
            <Descriptions.Item label="Tanggal Daftar">{formatDate(record.tgl_daftar)}</Descriptions.Item>
            <Descriptions.Item label="User">{record.user_id}</Descriptions.Item>
            <Descriptions.Item label="Praktek">{getHasilTag(record.praktek_hasil)}</Descriptions.Item>
            <Descriptions.Item label="Sertifikat">{record.sertif_nomor || '-'}</Descriptions.Item>
            <Descriptions.Item label="Tgl Sertifikat" span={2}>{formatDate(record.sertif_tanggal)}</Descriptions.Item>
          </Descriptions>
        )}
      </Modal>
    </div>
  )
}
