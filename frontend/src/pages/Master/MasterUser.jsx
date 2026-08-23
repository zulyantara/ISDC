import CrudPage from '../../components/CrudPage'

const userColumns = [
  { title: 'User ID', dataIndex: 'user_id', width: 120 },
  { title: 'Nama', dataIndex: 'user_name' },
  { title: 'Level', dataIndex: 'user_level', width: 100, render: v => {
    const map = { 1: 'Admin', 2: 'Super Admin', 3: 'Kasir', 4: 'Staff', 5: 'Superuser', 6: 'Instruktur' }
    return map[v] || v
  }},
  { title: 'Area ID', dataIndex: 'area_id', width: 80 },
  { title: 'Flag', dataIndex: 'flag', width: 60 },
  { title: 'Aktif', dataIndex: 'aktif', width: 60, render: v => v === -1 ? '✅' : '❌' },
]

const userFields = [
  { name: 'user_id', label: 'User ID', type: 'text', isId: true },
  { name: 'user_name', label: 'Nama User', type: 'text' },
  { name: 'user_level', label: 'Level (1=Admin,2=SuperAdmin,3=Kasir,6=Instruktur)', type: 'number' },
  { name: 'area_id', label: 'Area ID', type: 'number' },
]

export default function MasterUser() {
  return (
    <CrudPage
      title="Master User"
      apiPath="/users"
      menuUrl="mt_user"
      columns={userColumns}
      formFields={userFields}
      mapItemToForm={(item) => ({
        user_id: item.user_id,
        user_name: item.user_name,
        user_level: item.user_level,
        area_id: item.area_id,
      })}
    />
  )
}
