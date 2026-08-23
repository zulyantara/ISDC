import CrudPage from '../../components/CrudPage'

const nlColumns = [
  { title: 'ID', dataIndex: 'nl_id', width: 60 },
  { title: 'Nilai Minimum', dataIndex: 'nl_nilai', width: 150, render: v => <b>{v}</b> },
]

const nlFields = [
  { name: 'nl_nilai', label: 'Nilai Minimum Lulus', type: 'number' },
]

export default function MasterNilaiLulus() {
  return (
    <CrudPage
      title="Master Nilai Lulus"
      apiPath="/nilai-lulus"
      menuUrl="mt_nilai_lulus"
      columns={nlColumns}
      formFields={nlFields}
      mapItemToForm={(item) => ({ nl_nilai: item.nl_nilai })}
    />
  )
}
