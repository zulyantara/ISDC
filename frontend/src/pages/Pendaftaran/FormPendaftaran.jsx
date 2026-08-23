import { useState, useEffect } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import {
  Form, Input, Select, Button, Card, Row, Col, Typography,
  message, Space, Divider, InputNumber, DatePicker,
} from 'antd'
import { SaveOutlined, ArrowLeftOutlined } from '@ant-design/icons'
import dayjs from 'dayjs'
import api from '../../api/axios'

const { Title } = Typography
const { Option } = Select

export default function FormPendaftaran() {
  const navigate = useNavigate()
  const { id } = useParams()
  const isEdit = !!id

  const [form] = Form.useForm()
  const [loading, setLoading] = useState(false)
  const [kelasList, setKelasList] = useState([])
  const [submitting, setSubmitting] = useState(false)

  useEffect(() => {
    fetchKelas()
    if (isEdit) {
      fetchDaftar()
    }
  }, [id])

  const fetchKelas = async () => {
    try {
      const res = await api.get('/kelas')
      if (res.status) {
        setKelasList(res.data || [])
      }
    } catch {
      message.error('Gagal memuat data kelas')
    }
  }

  const fetchDaftar = async () => {
    setLoading(true)
    try {
      const res = await api.get(`/daftar/${id}`)
      if (res.status && res.data) {
        const d = res.data
        form.setFieldsValue({
          ...d,
          tgl_lahir: d.tgl_lahir ? dayjs(d.tgl_lahir) : null,
          kelas_id: d.kelas_id,
          kelamin_id: d.kelamin_id,
        })
      }
    } catch {
      message.error('Gagal memuat data pendaftaran')
    } finally {
      setLoading(false)
    }
  }

  const onFinish = async (values) => {
    setSubmitting(true)
    try {
      const payload = {
        ...values,
        tgl_lahir: values.tgl_lahir ? values.tgl_lahir.format('YYYY-MM-DD') : null,
      }

      let res
      if (isEdit) {
        res = await api.put(`/daftar/${id}`, payload)
      } else {
        res = await api.post('/daftar', payload)
      }

      if (res.status) {
        message.success(isEdit ? 'Berhasil mengupdate pendaftaran' : 'Berhasil mendaftarkan peserta')
        navigate('/pendaftaran')
      }
    } catch (err) {
      message.error(err.message || 'Gagal menyimpan data')
    } finally {
      setSubmitting(false)
    }
  }

  const onKelasChange = (kelasId) => {
    const selected = kelasList.find(k => k.kelas_id === kelasId)
    if (selected) {
      form.setFieldsValue({ biaya: selected.tarif })
    }
  }

  return (
    <div>
      <Row justify="space-between" align="middle" style={{ marginBottom: 16 }}>
        <Col>
          <Space>
            <Button icon={<ArrowLeftOutlined />} onClick={() => navigate('/pendaftaran')} />
            <Title level={4} style={{ margin: 0 }}>
              {isEdit ? '✏️ Edit Pendaftaran' : '📝 Pendaftaran Baru'}
            </Title>
          </Space>
        </Col>
      </Row>

      <Card loading={loading}>
        <Form
          form={form}
          layout="vertical"
          onFinish={onFinish}
          initialValues={{
            kelamin_id: 1,
            ref_id: 'J-000',
            biaya: 0,
          }}
        >
          <Row gutter={24}>
            {/* Data Diri */}
            <Col xs={24} lg={12}>
              <Divider orientation="left">Data Diri Peserta</Divider>

              <Form.Item
                label="Nama Lengkap"
                name="nama"
                rules={[{ required: true, message: 'Nama wajib diisi' }]}
              >
                <Input placeholder="NAMA LENGKAP" style={{ textTransform: 'uppercase' }} />
              </Form.Item>

              <Row gutter={16}>
                <Col span={12}>
                  <Form.Item
                    label="Jenis Kelamin"
                    name="kelamin_id"
                    rules={[{ required: true }]}
                  >
                    <Select>
                      <Option value={1}>Laki-laki</Option>
                      <Option value={2}>Perempuan</Option>
                    </Select>
                  </Form.Item>
                </Col>
                <Col span={12}>
                  <Form.Item
                    label="Tempat Lahir"
                    name="tempat_lahir"
                  >
                    <Input placeholder="Kota Kelahiran" style={{ textTransform: 'uppercase' }} />
                  </Form.Item>
                </Col>
              </Row>

              <Form.Item
                label="Tanggal Lahir"
                name="tgl_lahir"
              >
                <DatePicker style={{ width: '100%' }} placeholder="Pilih tanggal lahir" />
              </Form.Item>

              <Form.Item
                label="Alamat"
                name="alamat1"
              >
                <Input placeholder="Alamat lengkap" />
              </Form.Item>

              <Row gutter={16}>
                <Col span={16}>
                  <Form.Item
                    label="Alamat 2"
                    name="alamat2"
                  >
                    <Input placeholder="RT/RW, Kelurahan, dll" />
                  </Form.Item>
                </Col>
                <Col span={8}>
                  <Form.Item
                    label="Kota"
                    name="kota"
                  >
                    <Input placeholder="Kota" style={{ textTransform: 'uppercase' }} />
                  </Form.Item>
                </Col>
              </Row>
            </Col>

            {/* Data Pendaftaran */}
            <Col xs={24} lg={12}>
              <Divider orientation="left">Data Pendaftaran</Divider>

              <Form.Item
                label="Kelas / Program Pelatihan"
                name="kelas_id"
                rules={[{ required: true, message: 'Kelas wajib dipilih' }]}
              >
                <Select
                  placeholder="Pilih kelas"
                  showSearch
                  optionFilterProp="children"
                  onChange={onKelasChange}
                >
                  {kelasList.map(k => (
                    <Option key={k.kelas_id} value={k.kelas_id}>
                      {k.kelas} - {new Intl.NumberFormat('id-ID', { style: 'currency', currency: 'IDR', minimumFractionDigits: 0 }).format(k.tarif)}
                    </Option>
                  ))}
                </Select>
              </Form.Item>

              <Form.Item
                label="No. Referensi"
                name="ref_id"
              >
                <Input placeholder="J-000" />
              </Form.Item>

              <Form.Item
                label="Biaya"
                name="biaya"
              >
                <InputNumber
                  style={{ width: '100%' }}
                  formatter={(value) => `Rp ${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ',')}
                  parser={(value) => value.replace(/Rp\s?|(,*)/g, '')}
                  readOnly
                />
              </Form.Item>

              <Card
                size="small"
                style={{ background: '#f6ffed', borderColor: '#b7eb8f' }}
              >
                <Typography.Text type="secondary">
                  ℹ️ No. Pendaftaran akan digenerate otomatis oleh sistem.
                  Biaya akan diisi otomatis berdasarkan kelas yang dipilih.
                </Typography.Text>
              </Card>
            </Col>
          </Row>

          <Divider />

          <Row justify="end">
            <Space>
              <Button onClick={() => navigate('/pendaftaran')}>
                Batal
              </Button>
              <Button
                type="primary"
                htmlType="submit"
                icon={<SaveOutlined />}
                loading={submitting}
                size="large"
              >
                {isEdit ? 'Simpan Perubahan' : 'Daftarkan Peserta'}
              </Button>
            </Space>
          </Row>
        </Form>
      </Card>
    </div>
  )
}
