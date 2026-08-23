import CrudPage from '../../components/CrudPage'

const soalColumns = [
  { title: 'ID', dataIndex: 'ujian_id', width: 60 },
  { title: 'Sesi', dataIndex: 'sesi', width: 60 },
  { title: 'No.', dataIndex: 'nomor', width: 60 },
  { title: 'Kategori', dataIndex: 'category', width: 80 },
  { title: 'Soal', dataIndex: 'soal' },
]

const soalFields = [
  { name: 'sesi', label: 'Sesi', type: 'number' },
  { name: 'nomor', label: 'Nomor Soal', type: 'number' },
  { name: 'category', label: 'Kategori (teori_id)', type: 'number' },
  { name: 'soal', label: 'Soal', type: 'text' },
]

export default function MasterSoal() {
  return (
    <CrudPage
      title="Master Soal"
      apiPath="/soal"
      menuUrl="ujian"
      columns={soalColumns}
      formFields={soalFields}
      mapItemToForm={(item) => ({
        sesi: item.sesi,
        nomor: item.nomor,
        category: item.category,
        soal: item.soal,
      })}
    />
  )
}
