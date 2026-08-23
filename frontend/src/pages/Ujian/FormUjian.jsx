import { useState, useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import {
  Card, Table, InputNumber, Button, Typography, Space, Row, Col,
  message, Descriptions, Tag, Divider, Spin, Modal, Input, Alert,
} from 'antd'
import {
  ArrowLeftOutlined, SaveOutlined, CommentOutlined, CheckCircleOutlined,
} from '@ant-design/icons'
import api from '../../api/axios'
import { formatDate } from '../../utils/helpers'

const { Title, Text } = Typography
const { TextArea } = Input

export default function FormUjian() {
  const { id } = useParams()
  const navigate = useNavigate()
  const [loading, setLoading] = useState(true)
  const [submitting, setSubmitting] = useState(false)
  const [peserta, setPeserta] = useState(null)
  const [soal, setSoal] = useState([])
  const [scores, setScores] = useState({})
  const [commentModal, setCommentModal] = useState(false)
  const [comment, setComment] = useState({ pengetahuan: '', teknik: '', perilaku: '' })

  useEffect(() => {
    fetchSoal()
  }, [id])

  const fetchSoal = async () => {
    setLoading(true)
    try {
      const res = await api.get(`/peserta/${id}/soal`)
      if (res.status && res.data) {
        setPeserta(res.data.peserta)
        setSoal(res.data.soal || [])

        // Load existing scores
        try {
          const praktekRes = await api.get(`/peserta/${id}/praktek`)
          if (praktekRes.status && praktekRes.data) {
            const existing = {}
            praktekRes.data.forEach(p => {
              existing[p.soal_id] = p.hasil
            })
            setScores(existing)
          }
        } catch { /* no existing scores */ }

        // Load existing comment
        try {
          const commentRes = await api.get(`/peserta/${id}/comment`)
          if (commentRes.status && commentRes.data) {
            setComment({
              pengetahuan: commentRes.data.pengetahuan || '',
              teknik: commentRes.data.teknik || '',
              perilaku: commentRes.data.perilaku || '',
            })
          }
        } catch { /* no existing comment */ }
      }
    } catch {
      message.error('Gagal memuat data ujian')
    } finally {
      setLoading(false)
    }
  }

  const handleScoreChange = (ujianId, value) => {
    if (value < 0 || value > 100) {
      message.warning('Nilai harus antara 0-100')
      return
    }
    setScores(prev => ({ ...prev, [ujianId]: value }))
  }

  const handleSubmit = async () => {
    // Validate all scores filled
    const missing = soal.filter(s => scores[s.ujian_id] === undefined || scores[s.ujian_id] === null)
    if (missing.length > 0) {
      message.warning(`Masih ada ${missing.length} soal yang belum diisi`)
      return
    }

    setSubmitting(true)
    try {
      // Submit scores
      const results = soal.map(s => ({
        soal_id: s.ujian_id,
        hasil: scores[s.ujian_id] || 0,
      }))

      const res = await api.post(`/peserta/${id}/praktek`, { results })
      if (!res.status) {
        message.error('Gagal menyimpan nilai')
        return
      }

      // Submit comment if filled
      if (comment.pengetahuan || comment.teknik || comment.perilaku) {
        await api.post(`/peserta/${id}/comment`, comment)
      }

      message.success('Nilai ujian berhasil disimpan!')
      navigate('/peserta')
    } catch {
      message.error('Gagal menyimpan data ujian')
    } finally {
      setSubmitting(false)
    }
  }

  // Group soal by sesi
  const groupedSoal = soal.reduce((acc, s) => {
    if (!acc[s.sesi]) acc[s.sesi] = []
    acc[s.sesi].push(s)
    return acc
  }, {})

  const totalScore = soal.reduce((sum, s) => sum + (scores[s.ujian_id] || 0), 0)
  const avgScore = soal.length > 0 ? (totalScore / soal.length).toFixed(1) : 0

  if (loading) {
    return <div style={{ textAlign: 'center', padding: 60 }}><Spin size="large" /></div>
  }

  return (
    <div>
      <Row justify="space-between" align="middle" style={{ marginBottom: 16 }}>
        <Col>
          <Space>
            <Button icon={<ArrowLeftOutlined />} onClick={() => navigate('/peserta')} />
            <Title level={4} style={{ margin: 0 }}>📝 Ujian — {peserta?.nama}</Title>
          </Space>
        </Col>
        <Col>
          <Space>
            <Button icon={<CommentOutlined />} onClick={() => setCommentModal(true)}>
              Komentar
            </Button>
            <Button
              type="primary"
              icon={<SaveOutlined />}
              loading={submitting}
              onClick={handleSubmit}
              size="large"
            >
              Simpan Semua Nilai
            </Button>
          </Space>
        </Col>
      </Row>

      {/* Info Peserta */}
      <Card size="small" style={{ marginBottom: 16 }}>
        <Descriptions column={4} size="small">
          <Descriptions.Item label="No. Peserta">{peserta?.peserta_id}</Descriptions.Item>
          <Descriptions.Item label="Nama">{peserta?.nama}</Descriptions.Item>
          <Descriptions.Item label="Kelas">{peserta?.kelas_nama}</Descriptions.Item>
          <Descriptions.Item label="Total Soal">{soal.length}</Descriptions.Item>
        </Descriptions>
      </Card>

      {/* Rata-rata */}
      <Alert
        message={`Rata-rata Nilai: ${avgScore} / 100`}
        type={avgScore >= 70 ? 'success' : avgScore >= 50 ? 'warning' : 'error'}
        showIcon
        style={{ marginBottom: 16 }}
      />

      {/* Form Soal per Sesi */}
      {Object.entries(groupedSoal).map(([sesi, soalList]) => (
        <Card
          key={sesi}
          title={`Sesi ${sesi}`}
          style={{ marginBottom: 16 }}
        >
          <Table
            dataSource={soalList}
            rowKey="ujian_id"
            pagination={false}
            size="small"
            columns={[
              {
                title: 'No',
                dataIndex: 'nomor',
                width: 50,
                align: 'center',
              },
              {
                title: 'Soal',
                dataIndex: 'soal',
              },
              {
                title: 'Nilai (0-100)',
                width: 140,
                align: 'center',
                render: (_, record) => (
                  <InputNumber
                    min={0}
                    max={100}
                    step={5}
                    value={scores[record.ujian_id]}
                    onChange={(val) => handleScoreChange(record.ujian_id, val)}
                    style={{ width: 100 }}
                    placeholder="0"
                  />
                ),
              },
              {
                title: 'Status',
                width: 80,
                align: 'center',
                render: (_, record) => (
                  scores[record.ujian_id] !== undefined && scores[record.ujian_id] !== null
                    ? <Tag color="green"><CheckCircleOutlined /></Tag>
                    : <Tag color="default">-</Tag>
                ),
              },
            ]}
          />
        </Card>
      ))}

      {soal.length === 0 && (
        <Empty description="Tidak ada soal untuk peserta ini" />
      )}

      {/* Comment Modal */}
      <Modal
        title="Komentar Penilaian"
        open={commentModal}
        onCancel={() => setCommentModal(false)}
        onOk={() => { setCommentModal(false); message.success('Komentar disimpan') }}
        okText="Simpan"
      >
        <div style={{ marginBottom: 12 }}>
          <Text strong>Pengetahuan:</Text>
          <TextArea
            rows={3}
            value={comment.pengetahuan}
            onChange={(e) => setComment(prev => ({ ...prev, pengetahuan: e.target.value }))}
            placeholder="Komentar pengetahuan..."
          />
        </div>
        <div style={{ marginBottom: 12 }}>
          <Text strong>Teknik:</Text>
          <TextArea
            rows={3}
            value={comment.teknik}
            onChange={(e) => setComment(prev => ({ ...prev, teknik: e.target.value }))}
            placeholder="Komentar teknik..."
          />
        </div>
        <div>
          <Text strong>Perilaku:</Text>
          <TextArea
            rows={3}
            value={comment.perilaku}
            onChange={(e) => setComment(prev => ({ ...prev, perilaku: e.target.value }))}
            placeholder="Komentar perilaku..."
          />
        </div>
      </Modal>
    </div>
  )
}
