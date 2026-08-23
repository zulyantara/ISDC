/**
 * Format currency to Indonesian Rupiah
 */
export function formatRupiah(amount) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(amount)
}

/**
 * Format date to Indonesian format
 */
export function formatDate(dateStr) {
  if (!dateStr) return '-'
  const options = { day: 'numeric', month: 'long', year: 'numeric' }
  return new Date(dateStr).toLocaleDateString('id-ID', options)
}

/**
 * Format datetime
 */
export function formatDateTime(dateStr) {
  if (!dateStr) return '-'
  const options = { day: 'numeric', month: 'long', year: 'numeric', hour: '2-digit', minute: '2-digit' }
  return new Date(dateStr).toLocaleDateString('id-ID', options)
}

/**
 * Get kelamin text from ID
 */
export function getKelaminText(id) {
  const map = { 1: 'Laki-laki', 2: 'Perempuan' }
  return map[id] || '-'
}

/**
 * Generate print-friendly window
 */
export function printWindow(title, htmlContent) {
  const printWindow = window.open('', '_blank')
  printWindow.document.write(`
    <!DOCTYPE html>
    <html>
    <head>
      <title>${title}</title>
      <style>
        body { font-family: Arial, sans-serif; padding: 20px; }
        table { width: 100%; border-collapse: collapse; }
        th, td { padding: 8px 12px; border: 1px solid #ddd; text-align: left; }
        th { background: #f5f5f5; }
        .header { text-align: center; margin-bottom: 20px; }
        @media print { body { padding: 0; } }
      </style>
    </head>
    <body>
      ${htmlContent}
    </body>
    </html>
  `)
  printWindow.document.close()
  printWindow.focus()
  printWindow.print()
}

/**
 * Truncate text
 */
export function truncate(str, maxLen = 50) {
  if (!str) return ''
  return str.length > maxLen ? str.substring(0, maxLen) + '...' : str
}
