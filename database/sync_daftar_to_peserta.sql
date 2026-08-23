-- ============================================================
-- SYNC: Copy data dari tb_daftar yang belum ada di tb_peserta
-- Jalankan SETELAH membuat trigger tb_daftar_before_insert
-- ============================================================

INSERT INTO tb_peserta (
    `peserta_id`, `tgl_daftar`, `nama`, `kelamin_id`, `kelas_id`,
    `biaya`, `teori_hadir`, `teori_nilai`, `teori_hasil`, `teori_cetak`,
    `sertif_cetak`, `counter`, `tgl_input`, `user_id`, `area_id`
)
SELECT
    d.`peserta_id`, d.`tgl_daftar`, d.`nama`, d.`kelamin_id`, d.`kelas_id`,
    d.`biaya`, 'T', 0, 'T', 'T',
    'T', 0, d.`tgl_input`, d.`user_id`, d.`area_id`
FROM tb_daftar d
LEFT JOIN tb_peserta p ON d.`peserta_id` = p.`peserta_id`
WHERE p.`peserta_id` IS NULL;
