import CrudPage from '../../components/CrudPage'

const areaColumns = [
  { title: 'ID', dataIndex: 'area_id', width: 60 },
  { title: 'Kode', dataIndex: 'area_kode', width: 80 },
  { title: 'Nama Area', dataIndex: 'area_nama' },
  { title: 'Alamat', dataIndex: 'area_alamat' },
  { title: 'Telepon', dataIndex: 'area_telp', width: 120 },
]

const areaFields = [
  { name: 'area_kode', label: 'Kode Area', type: 'text' },
  { name: 'area_nama', label: 'Nama Area', type: 'text' },
  { name: 'area_alamat', label: 'Alamat', type: 'text' },
  { name: 'area_telp', label: 'Telepon', type: 'text', required: false },
]

export default function MasterArea() {
  return (
    <CrudPage
      title="Master Area"
      apiPath="/area"
      menuUrl="mt_area"
      columns={areaColumns}
      formFields={areaFields}
      mapItemToForm={(item) => ({
        area_kode: item.area_kode,
        area_nama: item.area_nama,
        area_alamat: item.area_alamat,
        area_telp: item.area_telp,
      })}
    />
  )
}
