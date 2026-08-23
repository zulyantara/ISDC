import { useState, useEffect, useCallback } from 'react'
import {
  Card, Tabs, Table, Button, Modal, Form, Input, InputNumber,
  Switch, Space, Typography, message, Popconfirm, Tag, Row, Col, Select, Tree
} from 'antd'
import {
  PlusOutlined, EditOutlined, DeleteOutlined,
  SafetyCertificateOutlined, MenuOutlined, TeamOutlined,
  SaveOutlined, ReloadOutlined
} from '@ant-design/icons'
import api from '../../api/axios'

const { Title, Text } = Typography

export default function MasterRBAC() {
  return (
    <div>
      <div style={{ marginBottom: 16 }}>
        <Title level={4} style={{ margin: 0 }}>Role-Based Access Control</Title>
        <Text type="secondary">Kelola level, menu, dan hak akses pengguna</Text>
      </div>
      <Card bodyStyle={{ padding: 0 }}>
        <Tabs
          defaultActiveKey="levels"
          style={{ padding: '0 16px' }}
          items={[
            { key: 'levels', label: <span><TeamOutlined /> Level / Role</span>, children: <LevelTab /> },
            { key: 'menus', label: <span><MenuOutlined /> Menu</span>, children: <MenuTab /> },
            { key: 'permissions', label: <span><SafetyCertificateOutlined /> Hak Akses</span>, children: <PermissionTab /> },
          ]}
        />
      </Card>
    </div>
  )
}

// ==================== LEVEL TAB ====================
function LevelTab() {
  const [data, setData] = useState([])
  const [loading, setLoading] = useState(false)
  const [modal, setModal] = useState({ open: false, edit: null })
  const [form] = Form.useForm()

  const fetch = useCallback(async () => {
    setLoading(true)
    try {
      const res = await api.get('/levels')
      setData(res.data || [])
    } catch { message.error('Gagal memuat data level') }
    setLoading(false)
  }, [])

  useEffect(() => { fetch() }, [fetch])

  const handleSave = async () => {
    try {
      const values = await form.validateFields()
      if (modal.edit) {
        await api.put(`/levels/${modal.edit.level_id}`, values)
        message.success('Level berhasil diupdate')
      } else {
        await api.post('/levels', values)
        message.success('Level berhasil dibuat')
      }
      setModal({ open: false, edit: null })
      form.resetFields()
      fetch()
    } catch (err) {
      if (err?.errorFields) return
      message.error(err?.message || 'Gagal menyimpan level')
    }
  }

  const handleDelete = async (id) => {
    try {
      await api.delete(`/levels/${id}`)
      message.success('Level berhasil dihapus')
      fetch()
    } catch { message.error('Gagal menghapus level') }
  }

  const columns = [
    { title: 'ID', dataIndex: 'level_id', width: 80, align: 'center' },
    { title: 'Nama Level', dataIndex: 'level_desc' },
    {
      title: 'Aksi', width: 120, align: 'center',
      render: (_, row) => (
        <Space>
          <Button size="small" icon={<EditOutlined />} onClick={() => { form.setFieldsValue(row); setModal({ open: true, edit: row }) }} />
          <Popconfirm title="Hapus level ini?" onConfirm={() => handleDelete(row.level_id)} okText="Ya" cancelText="Batal">
            <Button size="small" danger icon={<DeleteOutlined />} />
          </Popconfirm>
        </Space>
      )
    }
  ]

  return (
    <div style={{ padding: '16px 0' }}>
      <div style={{ marginBottom: 12, display: 'flex', justifyContent: 'space-between' }}>
        <Button type="primary" icon={<PlusOutlined />} onClick={() => { form.resetFields(); setModal({ open: true, edit: null }) }}>
          Tambah Level
        </Button>
      </div>
      <Table dataSource={data} columns={columns} rowKey="level_id" size="middle" pagination={{ pageSize: 20 }} />
      <Modal
        title={modal.edit ? 'Edit Level' : 'Tambah Level'}
        open={modal.open}
        onOk={handleSave}
        onCancel={() => { setModal({ open: false, edit: null }); form.resetFields() }}
        okText="Simpan"
        cancelText="Batal"
      >
        <Form form={form} layout="vertical">
          <Form.Item name="level_desc" label="Nama Level" rules={[{ required: true, message: 'Wajib diisi' }]}>
            <Input placeholder="Contoh: Admin, Kasir, Instruktur" />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}

// ==================== MENU TAB ====================
function MenuTab() {
  const [data, setData] = useState([])
  const [loading, setLoading] = useState(false)
  const [modal, setModal] = useState({ open: false, edit: null })
  const [form] = Form.useForm()

  const fetch = useCallback(async () => {
    setLoading(true)
    try {
      const res = await api.get('/menus')
      setData(res.data || [])
    } catch { message.error('Gagal memuat data menu') }
    setLoading(false)
  }, [])

  useEffect(() => { fetch() }, [fetch])

  const parentMenus = data.filter(m => m.menu_parent === 0 || m.menu_url === '#')

  const getParentName = (parentId) => {
    if (parentId === 0) return '-'
    const p = data.find(m => m.menu_id === parentId)
    return p ? p.menu_ket : parentId
  }

  const handleSave = async () => {
    try {
      const values = await form.validateFields()
      values.menu_parent = values.menu_parent || 0
      values.menu_order = values.menu_order || 0
      if (modal.edit) {
        await api.put(`/menus/${modal.edit.menu_id}`, values)
        message.success('Menu berhasil diupdate')
      } else {
        await api.post('/menus', values)
        message.success('Menu berhasil dibuat')
      }
      setModal({ open: false, edit: null })
      form.resetFields()
      fetch()
    } catch (err) {
      if (err?.errorFields) return
      message.error(err?.message || 'Gagal menyimpan menu')
    }
  }

  const handleDelete = async (id) => {
    try {
      await api.delete(`/menus/${id}`)
      message.success('Menu berhasil dihapus')
      fetch()
    } catch { message.error('Gagal menghapus menu') }
  }

  const columns = [
    { title: 'ID', dataIndex: 'menu_id', width: 60, align: 'center' },
    { title: 'Nama Menu', dataIndex: 'menu_ket' },
    { title: 'Parent', render: (_, row) => getParentName(row.menu_parent) },
    { title: 'URL', dataIndex: 'menu_url', render: v => <Tag>{v}</Tag> },
    { title: 'Order', dataIndex: 'menu_order', width: 60, align: 'center' },
    {
      title: 'Aksi', width: 120, align: 'center',
      render: (_, row) => (
        <Space>
          <Button size="small" icon={<EditOutlined />} onClick={() => { form.setFieldsValue(row); setModal({ open: true, edit: row }) }} />
          <Popconfirm title="Hapus menu ini?" onConfirm={() => handleDelete(row.menu_id)} okText="Ya" cancelText="Batal">
            <Button size="small" danger icon={<DeleteOutlined />} />
          </Popconfirm>
        </Space>
      )
    }
  ]

  return (
    <div style={{ padding: '16px 0' }}>
      <div style={{ marginBottom: 12, display: 'flex', justifyContent: 'space-between' }}>
        <Button type="primary" icon={<PlusOutlined />} onClick={() => { form.resetFields(); setModal({ open: true, edit: null }) }}>
          Tambah Menu
        </Button>
      </div>
      <Table dataSource={data} columns={columns} rowKey="menu_id" size="middle" pagination={{ pageSize: 50 }} />
      <Modal
        title={modal.edit ? 'Edit Menu' : 'Tambah Menu'}
        open={modal.open}
        onOk={handleSave}
        onCancel={() => { setModal({ open: false, edit: null }); form.resetFields() }}
        okText="Simpan"
        cancelText="Batal"
      >
        <Form form={form} layout="vertical">
          <Form.Item name="menu_ket" label="Nama Menu" rules={[{ required: true, message: 'Wajib diisi' }]}>
            <Input placeholder="Contoh: Dashboard" />
          </Form.Item>
          <Form.Item name="menu_parent" label="Menu Parent">
            <Select allowClear placeholder="Root (tidak ada parent)" options={parentMenus.map(m => ({ value: m.menu_id, label: m.menu_ket }))} />
          </Form.Item>
          <Form.Item name="menu_url" label="URL / Route" rules={[{ required: true, message: 'Wajib diisi' }]}>
            <Input placeholder="Contoh: dashboard, pendaftaran" />
          </Form.Item>
          <Form.Item name="menu_order" label="Urutan">
            <InputNumber min={0} max={999} style={{ width: '100%' }} />
          </Form.Item>
        </Form>
      </Modal>
    </div>
  )
}

// ==================== PERMISSION TAB ====================
function PermissionTab() {
  const [levels, setLevels] = useState([])
  const [menus, setMenus] = useState([])
  const [permissions, setPermissions] = useState([])
  const [selectedLevel, setSelectedLevel] = useState(null)
  const [loading, setLoading] = useState(false)
  const [saving, setSaving] = useState(false)

  useEffect(() => {
    Promise.all([
      api.get('/levels').then(r => setLevels(r.data || [])),
      api.get('/menus').then(r => setMenus(r.data || [])),
    ])
  }, [])

  const fetchPermissions = useCallback(async (levelId) => {
    if (!levelId) { setPermissions([]); return }
    setLoading(true)
    try {
      const res = await api.get(`/hak-akses/role/${levelId}`)
      setPermissions(res.data || [])
    } catch { message.error('Gagal memuat hak akses') }
    setLoading(false)
  }, [])

  useEffect(() => { if (selectedLevel) fetchPermissions(selectedLevel) }, [selectedLevel, fetchPermissions])

  // Toggle permission flag
  const togglePermission = (menuId, field) => {
    setPermissions(prev => {
      const existing = prev.find(p => p.ha_menu === menuId)
      if (existing) {
        return prev.map(p => p.ha_menu === menuId ? { ...p, [field]: p[field] ? 0 : 1 } : p)
      } else {
        return [...prev, {
          ha_menu: menuId, ha_ur: selectedLevel,
          ha_view: 0, ha_insert: 0, ha_update: 0, ha_delete: 0, ha_proses: 0,
          [field]: 1
        }]
      }
    })
  }

  // Set all permissions for a menu
  const setAllPerms = (menuId, value) => {
    const v = value ? 1 : 0
    setPermissions(prev => {
      const existing = prev.find(p => p.ha_menu === menuId)
      if (existing) {
        return prev.map(p => p.ha_menu === menuId ? { ...p, ha_view: v, ha_insert: v, ha_update: v, ha_delete: v, ha_proses: v } : p)
      } else if (value) {
        return [...prev, {
          ha_menu: menuId, ha_ur: selectedLevel,
          ha_view: 1, ha_insert: 1, ha_update: 1, ha_delete: 1, ha_proses: 1
        }]
      }
      return prev
    })
  }

  // Set all menus permission at once
  const setAllMenus = (value) => {
    const v = value ? 1 : 0
    const newPerms = menus.map(m => ({
      ha_menu: m.menu_id, ha_ur: selectedLevel,
      ha_view: v, ha_insert: v, ha_update: v, ha_delete: v, ha_proses: v
    }))
    setPermissions(value ? newPerms : [])
  }

  const handleSave = async () => {
    if (!selectedLevel) { message.warning('Pilih level terlebih dahulu'); return }
    setSaving(true)
    try {
      await api.put(`/hak-akses/role/${selectedLevel}`, permissions)
      message.success('Hak akses berhasil disimpan')
    } catch { message.error('Gagal menyimpan hak akses') }
    setSaving(false)
  }

  const getPerm = (menuId) => permissions.find(p => p.ha_menu === menuId) || { ha_view: 0, ha_insert: 0, ha_update: 0, ha_delete: 0, ha_proses: 0 }

  const permColumns = [
    {
      title: 'Menu', dataIndex: 'menu_ket', width: 200,
      render: (text, row) => (
        <div style={{ paddingLeft: row.menu_parent ? 20 : 0 }}>
          <Text strong={!row.menu_parent || row.menu_parent === 0}>{text}</Text>
        </div>
      )
    },
    {
      title: <span style={{ fontSize: 12 }}>Semua</span>, width: 60, align: 'center',
      render: (_, row) => {
        const p = getPerm(row.menu_id)
        const all = p.ha_view && p.ha_insert && p.ha_update && p.ha_delete
        return <Switch size="small" checked={all} onChange={(v) => setAllPerms(row.menu_id, v)} />
      }
    },
    {
      title: <span style={{ fontSize: 12 }}>View</span>, width: 50, align: 'center',
      render: (_, row) => <Switch size="small" checked={getPerm(row.menu_id).ha_view} onChange={() => togglePermission(row.menu_id, 'ha_view')} />
    },
    {
      title: <span style={{ fontSize: 12 }}>Insert</span>, width: 50, align: 'center',
      render: (_, row) => <Switch size="small" checked={getPerm(row.menu_id).ha_insert} onChange={() => togglePermission(row.menu_id, 'ha_insert')} />
    },
    {
      title: <span style={{ fontSize: 12 }}>Update</span>, width: 50, align: 'center',
      render: (_, row) => <Switch size="small" checked={getPerm(row.menu_id).ha_update} onChange={() => togglePermission(row.menu_id, 'ha_update')} />
    },
    {
      title: <span style={{ fontSize: 12 }}>Delete</span>, width: 50, align: 'center',
      render: (_, row) => <Switch size="small" checked={getPerm(row.menu_id).ha_delete} onChange={() => togglePermission(row.menu_id, 'ha_delete')} />
    },
    {
      title: <span style={{ fontSize: 12 }}>Proses</span>, width: 50, align: 'center',
      render: (_, row) => <Switch size="small" checked={getPerm(row.menu_id).ha_proses} onChange={() => togglePermission(row.menu_id, 'ha_proses')} />
    },
  ]

  return (
    <div style={{ padding: '16px 0' }}>
      <Row gutter={16} style={{ marginBottom: 16 }}>
        <Col span={8}>
          <Text strong>Pilih Level/Role:</Text>
          <Select
            style={{ width: '100%', marginTop: 4 }}
            placeholder="Pilih level"
            value={selectedLevel}
            onChange={setSelectedLevel}
            options={levels.map(l => ({ value: l.level_id, label: `${l.level_desc} (${l.level_id})` }))}
          />
        </Col>
        <Col span={16} style={{ display: 'flex', alignItems: 'flex-end', gap: 8, justifyContent: 'flex-end' }}>
          <Button
            icon={<SaveOutlined />}
            type="primary"
            loading={saving}
            onClick={handleSave}
            disabled={!selectedLevel}
          >
            Simpan Hak Akses
          </Button>
          <Button
            icon={<ReloadOutlined />}
            onClick={() => selectedLevel && fetchPermissions(selectedLevel)}
            disabled={!selectedLevel}
          >
            Reload
          </Button>
        </Col>
      </Row>

      {selectedLevel && (
        <div style={{ marginBottom: 8, display: 'flex', gap: 8, alignItems: 'center' }}>
          <Text type="secondary">Set semua menu:</Text>
          <Button size="small" onClick={() => setAllMenus(true)}>Aktifkan Semua</Button>
          <Button size="small" danger onClick={() => setAllMenus(false)}>Nonaktifkan Semua</Button>
        </div>
      )}

      <Table
        dataSource={menus}
        columns={permColumns}
        rowKey="menu_id"
        loading={loading}
        size="small"
        pagination={false}
        bordered
        rowClassName={(row) => {
          const p = getPerm(row.menu_id)
          return p.ha_view ? 'row-perm-active' : ''
        }}
      />
    </div>
  )
}
