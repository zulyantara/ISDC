import { useState, useEffect } from 'react'
import {
  Table, Button, Input, Space, Popconfirm, Modal, Form,
  Typography, Card, Row, Col, message, Tag, Empty,
} from 'antd'
import { PlusOutlined, EditOutlined, DeleteOutlined, SearchOutlined, ReloadOutlined } from '@ant-design/icons'
import { useAuth } from '../context/AuthContext'
import api from '../api/axios'

const { Title, Text } = Typography

/**
 * Reusable CRUD Table component with RBAC guard
 * @param {string} title - Page title
 * @param {string} apiPath - API path e.g. '/kelas'
 * @param {string} menuUrl - RBAC menu URL for permission check
 * @param {Array} columns - Table columns config
 * @param {Array} formFields - Form field configs [{name, label, type, options, required}]
 * @param {Function} mapItemToForm - Transform API item to form values
 */
export default function CrudPage({ title, apiPath, menuUrl, columns, formFields, mapItemToForm }) {
  const { hasPermission } = useAuth()
  const [data, setData] = useState([])
  const [loading, setLoading] = useState(false)
  const [searchText, setSearchText] = useState('')
  const [modalOpen, setModalOpen] = useState(false)
  const [editingItem, setEditingItem] = useState(null)
  const [submitting, setSubmitting] = useState(false)
  const [form] = Form.useForm()

  // RBAC permissions
  const canInsert = menuUrl ? hasPermission(menuUrl, 'insert') : true
  const canUpdate = menuUrl ? hasPermission(menuUrl, 'update') : true
  const canDelete = menuUrl ? hasPermission(menuUrl, 'delete') : true
  const canView = menuUrl ? hasPermission(menuUrl, 'view') : true

  useEffect(() => { fetchData() }, [])

  const fetchData = async () => {
    setLoading(true)
    try {
      const res = await api.get(apiPath)
      if (res.status) setData(res.data || [])
    } catch {
      message.error(`Gagal memuat data ${title}`)
    } finally {
      setLoading(false)
    }
  }

  const handleCreate = () => {
    setEditingItem(null)
    form.resetFields()
    setModalOpen(true)
  }

  const handleEdit = (record) => {
    setEditingItem(record)
    form.setFieldsValue(mapItemToForm ? mapItemToForm(record) : record)
    setModalOpen(true)
  }

  const handleDelete = async (id) => {
    try {
      const res = await api.delete(`${apiPath}/${id}`)
      if (res.status) {
        message.success(`Berhasil menghapus data`)
        fetchData()
      }
    } catch {
      message.error(`Gagal menghapus data`)
    }
  }

  const handleSubmit = async () => {
    try {
      const values = await form.validateFields()
      setSubmitting(true)

      let res
      const idField = formFields.find(f => f.isId)
      const idValue = idField ? editingItem?.[idField.name] : editingItem?.id || editingItem?.[`${apiPath.split('/').pop()}_id`]

      if (editingItem && idValue) {
        res = await api.put(`${apiPath}/${idValue}`, values)
      } else {
        res = await api.post(apiPath, values)
      }

      if (res.status) {
        message.success(editingItem ? 'Berhasil mengupdate data' : 'Berhasil menambah data')
        setModalOpen(false)
        fetchData()
      } else {
        message.error(res.message || 'Gagal menyimpan')
      }
    } catch (err) {
      if (err.errorFields) return
      message.error('Gagal menyimpan data')
    } finally {
      setSubmitting(false)
    }
  }

  const filteredData = data.filter(item => {
    const s = searchText.toLowerCase()
    return Object.values(item).some(v =>
      String(v || '').toLowerCase().includes(s)
    )
  })

  // Build action column based on permissions
  const actionColumn = (canUpdate || canDelete) ? {
    title: 'Aksi',
    key: 'aksi',
    width: 120,
    align: 'center',
    render: (_, record) => {
      const idField = formFields.find(f => f.isId)
      const id = idField ? record[idField.name] : record.id || record[`${apiPath.split('/').pop()}_id`]
      return (
        <Space size="small">
          {canUpdate && (
            <Button size="small" icon={<EditOutlined />} onClick={() => handleEdit(record)} />
          )}
          {canDelete && (
            <Popconfirm title="Hapus data ini?" onConfirm={() => handleDelete(id)} okText="Ya" cancelText="Batal">
              <Button danger size="small" icon={<DeleteOutlined />} />
            </Popconfirm>
          )}
        </Space>
      )
    },
  } : null

  const tableColumns = actionColumn ? [...columns, actionColumn] : columns

  return (
    <div>
      <Row justify="space-between" align="middle" style={{ marginBottom: 16 }}>
        <Col>
          <Title level={4} style={{ margin: 0 }}>{title}</Title>
        </Col>
        <Col>
          {canInsert && (
            <Button type="primary" icon={<PlusOutlined />} onClick={handleCreate}>
              Tambah Baru
            </Button>
          )}
        </Col>
      </Row>

      <Card bodyStyle={{ padding: 20 }}>
        <Row justify="space-between" style={{ marginBottom: 16 }}>
          <Col>
            <Input
              placeholder="Cari..."
              prefix={<SearchOutlined style={{ color: '#bbb' }} />}
              value={searchText}
              onChange={(e) => setSearchText(e.target.value)}
              style={{ width: 280, borderRadius: 8 }}
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
          columns={tableColumns}
          dataSource={filteredData}
          rowKey={(r) => {
            const idField = formFields.find(f => f.isId)
            return idField ? r[idField.name] : r.id || r[`${apiPath.split('/').pop()}_id`]
          }}
          loading={loading}
          pagination={{ pageSize: 10, showSizeChanger: true, showTotal: (t) => `Total ${t} data` }}
          size="middle"
          scroll={{ x: 800 }}
          locale={{ emptyText: <Empty description="Tidak ada data" /> }}
        />
      </Card>

      <Modal
        title={editingItem ? `Edit Data` : `Tambah Data`}
        open={modalOpen}
        onCancel={() => setModalOpen(false)}
        onOk={handleSubmit}
        confirmLoading={submitting}
        okText={editingItem ? 'Simpan Perubahan' : 'Simpan'}
        cancelText="Batal"
        width={600}
        okButtonProps={{ style: { borderRadius: 8 } }}
        cancelButtonProps={{ style: { borderRadius: 8 } }}
      >
        <Form form={form} layout="vertical" autoComplete="off">
          {formFields.map(field => (
            <Form.Item
              key={field.name}
              name={field.name}
              label={field.label}
              rules={field.required !== false ? [{ required: true, message: `${field.label} wajib diisi` }] : []}
            >
              {field.type === 'select' ? (
                <select style={{ width: '100%', padding: '6px 11px', borderRadius: 8, border: '1px solid #d9d9d9', height: 36 }}>
                  <option value="">Pilih {field.label}</option>
                  {field.options.map(opt => (
                    <option key={opt.value} value={opt.value}>{opt.label}</option>
                  ))}
                </select>
              ) : field.type === 'number' ? (
                <input type="number" min={0} style={{ width: '100%', padding: '6px 11px', borderRadius: 8, border: '1px solid #d9d9d9', height: 36 }} />
              ) : field.type === 'textarea' ? (
                <textarea rows={3} style={{ width: '100%', padding: '6px 11px', borderRadius: 8, border: '1px solid #d9d9d9' }} />
              ) : (
                <input type="text" style={{ width: '100%', padding: '6px 11px', borderRadius: 8, border: '1px solid #d9d9d9', height: 36 }} />
              )}
            </Form.Item>
          ))}
        </Form>
      </Modal>
    </div>
  )
}
