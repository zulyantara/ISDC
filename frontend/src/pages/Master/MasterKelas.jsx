import CrudPage from '../../components/CrudPage'

const kelasColumns = [
  { title: 'ID', dataIndex: 'kelasid', width: 60 },
  { title: 'Kode', dataIndex: 'kelas_id', width: 80 },
  { title: 'Nama Kelas', dataIndex: 'kelas' },
  { title: 'Tarif', dataIndex: 'tarif', render: v => `Rp ${(v||0).toLocaleString('id-ID')}` },
  { title: 'Teori ID', dataIndex: 'teori_id', width: 80 },
  { title: 'Area ID', dataIndex: 'area_id', width: 80 },
]

const kelasFields = [
  { name: 'kelas_id', label: 'Kode Kelas', type: 'number', isId: true },
  { name: 'kelas', label: 'Nama Kelas', type: 'text' },
  { name: 'tarif', label: 'Tarif', type: 'number' },
  { name: 'teori_id', label: 'Teori ID (Kategori)', type: 'number' },
  { name: 'area_id', label: 'Area ID', type: 'number' },
]

export default function MasterKelas() {
  return (
    <CrudPage
      title="Master Kelas"
      apiPath="/kelas"
      menuUrl="mt_kelas"
      columns={kelasColumns}
      formFields={kelasFields}
      mapItemToForm={(item) => ({
        kelas_id: item.kelas_id,
        kelas: item.kelas,
        tarif: item.tarif,
        teori_id: item.teori_id,
        area_id: item.area_id,
      })}
    />
  )
}
