import { useState, useEffect } from 'react'
import {
  Table, Input, Space, Typography, Card, Row, Col, Button, Tag, message, Empty,
} from 'antd'
import { SearchOutlined, ReloadOutlined, PrinterOutlined, FileTextOutlined } from '@ant-design/icons'
import api from '../../api/axios'
import { useAuth } from '../../context/AuthContext'
import { formatDate } from '../../utils/helpers'

const { Title, Text } = Typography

export default function ListSertifikat() {
  const { hasPermission } = useAuth()
  const canView = hasPermission('sertifikat', 'view')
  const [data, setData] = useState([])
  const [loading, setLoading] = useState(false)
  const [searchText, setSearchText] = useState('')

  useEffect(() => { fetchData() }, [])

  const fetchData = async () => {
    setLoading(true)
    try {
      const res = await api.get('/peserta')
      if (res.status) {
        const sertif = (res.data || []).filter(p => p.sertif_nomor && p.sertif_nomor !== '')
        setData(sertif)
      }
    } catch {
      message.error('Gagal memuat data sertifikat')
    } finally {
      setLoading(false)
    }
  }

  const handlePrint = (record) => {
    const html = `
      <div style="width:210mm; min-height:148mm; margin:auto; padding:15mm; font-family:Arial;">
        <div style="text-align:center; border:3px double #333; padding:20mm 15mm; min-height:100mm;">
          <p style="font-size:14px; color:#666; margin:0;">ISDC</p>
          <h1 style="font-size:28px; margin:5mm 0;">INDONESIA SAFETY DRIVING CENTER</h1>
          <p style="font-size:11px; color:#999; margin:0;">Jl. Daan Mogot Jakarta Barat</p>
          <br>
          <p style="font-size:16px;">SERTIFIKAT KEAHLIAN</p>
          <p style="font-size:12px;">No: <b>${record.sertif_nomor}</b></p>
          <br>
          <p style="font-size:13px;">Diberikan kepada:</p>
          <h2 style="font-size:24px; text-decoration:underline;">${record.nama}</h2>
          <br>
          <p style="font-size:12px;">Telah mengikuti pelatihan</p>
          <p style="font-size:14px;"><b>${record.kelas_nama || '-'}</b></p>
          <br>
          <p style="font-size:12px;">Tanggal: ${formatDate(record.sertif_tanggal)}</p>
          <br><br>
          <table style="width:100%; font-size:12px;">
            <tr>
              <td style="text-align:left; width:50%;">Mengetahui,</td>
              <td style="text-align:right; width:50%;">Jakarta, ${formatDate(record.sertif_tanggal)}</td>
            </tr>
            <tr><td colspan="2"><br><br></td></tr>
            <tr>
              <td style="text-align:left;"><b>Direktur</b></td>
              <td style="text-align:right;"><b>Penanggung Jawab</b></td>
            </tr>
          </table>
        </div>
      </div>
    `
    const w = window.open('', '_blank')
    w.document.write(`<html><head><title>Sertifikat - ${record.nama}</title>
      <style>@media print{body{margin:0;}}</style></head><body>${html}</body></html>`)
    w.document.close()
    w.onload = () => w.print()
  }

  const filteredData = data.filter(item => {
    const s = searchText.toLowerCase()
    return (
      (item.peserta_id || '').toLowerCase().includes(s) ||
      (item.nama || '').toLowerCase().includes(s) ||
      (item.sertif_nomor || '').toLowerCase().includes(s)
    )
  })

  return (
    <div>
      <Row justify="space-between" align="middle" style={{ marginBottom: 20 }}>
        <Col>
          <Title level={4} style={{ margin: 0 }}>📜 Daftar Sertifikat</Title>
          <Text type="secondary" style={{ fontSize: 13 }}>Kelola dan cetak sertifikat peserta</Text>
        </Col>
      </Row>

      <Card style={{ borderRadius: 12, boxShadow: '0 1px 3px rgba(0,0,0,0.06)' }}>
        <Row justify="space-between" style={{ marginBottom: 16 }}>
          <Col>
            <Input
              placeholder="Cari nama, nomor sertifikat..."
              prefix={<SearchOutlined style={{ color: '#bbb' }} />}
              value={searchText}
              onChange={e => setSearchText(e.target.value)}
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
          columns={[
            {
              title: 'No. Peserta',
              dataIndex: 'peserta_id',
              width: 160,
              render: t => <Text code style={{ fontSize: 12 }}>{t}</Text>
            },
            {
              title: 'Nama',
              dataIndex: 'nama',
              render: t => <Text strong>{t}</Text>
            },
            { title: 'Kelas', dataIndex: 'kelas_nama', width: 200 },
            {
              title: 'No. Sertifikat',
              dataIndex: 'sertif_nomor',
              width: 180,
              render: t => <Tag color="blue" style={{ borderRadius: 12 }}>{t}</Tag>
            },
            {
              title: 'Tanggal',
              dataIndex: 'sertif_tanggal',
              width: 120,
              render: t => formatDate(t)
            },
            {
              title: 'Aksi',
              key: 'aksi',
              width: 80,
              align: 'center',
              render: (_, r) => (
                <Button
                  type="primary"
                  size="small"
                  icon={<PrinterOutlined />}
                  onClick={() => handlePrint(r)}
                  style={{ borderRadius: 6 }}
                />
              ),
            },
          ]}
          dataSource={filteredData}
          rowKey="peserta_id"
          loading={loading}
          pagination={{ pageSize: 10, showTotal: t => `Total ${t} sertifikat` }}
          size="middle"
          locale={{ emptyText: <Empty description="Belum ada sertifikat" /> }}
        />
      </Card>
    </div>
  )
}
