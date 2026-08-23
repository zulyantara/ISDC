import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import {
  Table, Button, Input, Space, Tag, Popconfirm, message,
  Typography, Card, Row, Col, Tooltip, Modal, QRCode,
} from 'antd'
import {
  PlusOutlined, EditOutlined, DeleteOutlined,
  PrinterOutlined, SearchOutlined, ReloadOutlined,
  EyeOutlined,
} from '@ant-design/icons'
import api from '../../api/axios'
import { useAuth } from '../../context/AuthContext'
import { formatRupiah, formatDate, getKelaminText } from '../../utils/helpers'

const { Title } = Typography

export default function ListPendaftaran() {
  const navigate = useNavigate()
  const { hasPermission } = useAuth()
  const canInsert = hasPermission('pendaftaran', 'insert')
  const canUpdate = hasPermission('pendaftaran', 'update')
  const canDelete = hasPermission('pendaftaran', 'delete')
  const [data, setData] = useState([])
  const [loading, setLoading] = useState(false)
  const [searchText, setSearchText] = useState('')
  const [printModal, setPrintModal] = useState({ open: false, record: null })

  useEffect(() => {
    fetchData()
  }, [])

  const fetchData = async () => {
    setLoading(true)
    try {
      const res = await api.get('/daftar')
      if (res.status) {
        setData(res.data || [])
      }
    } catch {
      message.error('Gagal memuat data pendaftaran')
    } finally {
      setLoading(false)
    }
  }

  const handleDelete = async (peserta_id) => {
    try {
      const res = await api.delete(`/daftar/${peserta_id}`)
      if (res.status) {
        message.success('Berhasil menghapus pendaftaran')
        fetchData()
      }
    } catch {
      message.error('Gagal menghapus pendaftaran')
    }
  }

  const handlePrint = (record) => {
    const html = `
      <div style="width: 300px; margin: 0 auto; font-size: 12px;">
        <div class="header">
          <h3 style="margin:0;">ISDC</h3>
          <p style="margin:0;">Indonesia Safety Driving Center</p>
          <p style="margin:0;">Jl. Daan Mogot Jakarta Barat</p>
        </div>
        <hr style="border: 1px dashed #000;">
        <table style="width:100%; border:none; font-size:11px;">
          <tr><td style="border:none; padding:2px 0;">No. Pendaftaran</td><td style="border:none; padding:2px 0; font-weight:bold;">${record.peserta_id}</td></tr>
          <tr><td style="border:none; padding:2px 0;">Tanggal</td><td style="border:none; padding:2px 0;">${formatDate(record.tgl_daftar)}</td></tr>
          <tr><td style="border:none; padding:2px 0;">Nama</td><td style="border:none; padding:2px 0;">${record.nama}</td></tr>
          <tr><td style="border:none; padding:2px 0;">Kelamin</td><td style="border:none; padding:2px 0;">${getKelaminText(record.kelamin_id)}</td></tr>
          <tr><td style="border:none; padding:2px 0;">Tempat, Tgl Lahir</td><td style="border:none; padding:2px 0;">${record.tempat_lahir || '-'}, ${formatDate(record.tgl_lahir)}</td></tr>
          <tr><td style="border:none; padding:2px 0;">Alamat</td><td style="border:none; padding:2px 0;">${record.alamat1 || ''} ${record.alamat2 || ''}</td></tr>
          <tr><td style="border:none; padding:2px 0;">Kota</td><td style="border:none; padding:2px 0;">${record.kota || '-'}</td></tr>
          <tr><td style="border:none; padding:2px 0;">Kelas</td><td style="border:none; padding:2px 0;">${record.kelas || '-'}</td></tr>
          <tr><td style="border:none; padding:2px 0;">Biaya</td><td style="border:none; padding:2px 0; font-weight:bold;">${formatRupiah(record.biaya || 0)}</td></tr>
        </table>
        <hr style="border: 1px dashed #000;">
        <div style="text-align:center; margin-top:10px;">
          <div style="display:inline-block;">
            <img src="https://api.qrserver.com/v1/create-qr-code/?size=100x100&data=${record.peserta_id}" alt="QR" width="80" height="80">
          </div>
        </div>
        <p style="text-align:center; margin-top:10px; font-size:10px; color:#666;">
          Scan QR Code untuk verifikasi
        </p>
      </div>
    `
    const printWindow = window.open('', '_blank')
    printWindow.document.write(`
      <!DOCTYPE html>
      <html><head><title>Struk Pendaftaran</title>
      <style>
        @media print { body { margin: 0; padding: 10mm; } }
        body { font-family: Arial, sans-serif; }
      </style>
      </head><body>${html}</body></html>
    `)
    printWindow.document.close()
    printWindow.onload = () => { printWindow.print() }
  }

  const filteredData = data.filter(item => {
    const search = searchText.toLowerCase()
    return (
      (item.peserta_id || '').toLowerCase().includes(search) ||
      (item.nama || '').toLowerCase().includes(search) ||
      (item.kelas || '').toLowerCase().includes(search) ||
      (item.kota || '').toLowerCase().includes(search)
    )
  })

  const columns = [
    {
      title: 'No. Pendaftaran',
      dataIndex: 'peserta_id',
      key: 'peserta_id',
      width: 160,
      render: (text) => <Typography.Text code>{text}</Typography.Text>,
    },
    {
      title: 'Tanggal',
      dataIndex: 'tgl_daftar',
      key: 'tgl_daftar',
      width: 120,
      render: (text) => formatDate(text),
    },
    {
      title: 'Nama Peserta',
      dataIndex: 'nama',
      key: 'nama',
      render: (text) => <Typography.Text strong>{text}</Typography.Text>,
    },
    {
      title: 'Kelas',
      dataIndex: 'kelas',
      key: 'kelas',
      width: 200,
      render: (text) => text || '-',
    },
    {
      title: 'Kota',
      dataIndex: 'kota',
      key: 'kota',
      width: 120,
    },
    {
      title: 'Biaya',
      dataIndex: 'biaya',
      key: 'biaya',
      width: 120,
      align: 'right',
      render: (val) => formatRupiah(val || 0),
    },
    {
      title: 'User',
      dataIndex: 'user_name',
      key: 'user_name',
      width: 120,
    },
    {
      title: 'Aksi',
      key: 'aksi',
      width: 180,
      align: 'center',
      render: (_, record) => (
        <Space size="small">
          <Tooltip title="Print Struk">
            <Button
              type="primary"
              size="small"
              icon={<PrinterOutlined />}
              onClick={() => handlePrint(record)}
            />
          </Tooltip>
          {canUpdate && (
            <Tooltip title="Edit">
              <Button
                size="small"
                icon={<EditOutlined />}
                onClick={() => navigate(`/pendaftaran/edit/${record.peserta_id}`)}
              />
            </Tooltip>
          )}
          {canDelete && (
            <Popconfirm
              title="Hapus pendaftaran ini?"
              onConfirm={() => handleDelete(record.peserta_id)}
              okText="Ya"
              cancelText="Batal"
            >
              <Tooltip title="Hapus">
                <Button
                  danger
                  size="small"
                  icon={<DeleteOutlined />}
                />
              </Tooltip>
            </Popconfirm>
          )}
        </Space>
      ),
    },
  ]

  return (
    <div>
      <Row justify="space-between" align="middle" style={{ marginBottom: 16 }}>
        <Col>
          <Title level={4} style={{ margin: 0 }}>Daftar Pendaftaran</Title>
        </Col>
        <Col>
          {canInsert && (
            <Button
              type="primary"
              icon={<PlusOutlined />}
              onClick={() => navigate('/pendaftaran/tambah')}
            >
              Pendaftaran Baru
            </Button>
          )}
        </Col>
      </Row>

      <Card>
        <Row justify="space-between" style={{ marginBottom: 16 }}>
          <Col>
            <Input
              placeholder="Cari nama, ID, kelas, kota..."
              prefix={<SearchOutlined />}
              value={searchText}
              onChange={(e) => setSearchText(e.target.value)}
              style={{ width: 300 }}
              allowClear
            />
          </Col>
          <Col>
            <Button icon={<ReloadOutlined />} onClick={fetchData}>
              Refresh
            </Button>
          </Col>
        </Row>

        <Table
          columns={columns}
          dataSource={filteredData}
          rowKey="peserta_id"
          loading={loading}
          pagination={{
            pageSize: 10,
            showSizeChanger: true,
            pageSizeOptions: ['10', '25', '50'],
            showTotal: (total) => `Total ${total} data`,
          }}
          size="middle"
          scroll={{ x: 1200 }}
        />
      </Card>
    </div>
  )
}
