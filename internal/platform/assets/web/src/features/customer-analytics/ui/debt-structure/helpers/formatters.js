export function escapeHtml(value) {
    return String(value)
        .replaceAll('&', '&amp;')
        .replaceAll('<', '&lt;')
        .replaceAll('>', '&gt;')
        .replaceAll('"', '&quot;')
        .replaceAll("'", '&#39;');
}

export function formatDate(dateStr) {
    if (!dateStr) return '—';
    const date = new Date(dateStr);
    if (Number.isNaN(date.getTime())) return dateStr;
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
}

export function formatMoneyInMillions(value) {
    if (value === null || value === undefined) return '—';
    const num = Number(value);
    if (Number.isNaN(num)) return '—';
    const mln = num / 1_000_000;
    return mln.toLocaleString('ru-RU', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
}

export function formatObjectMetric(value, format) {
    if (value === null || value === undefined) return '—';

    if (format === 'money') {
        const number = Number(value);
        if (Number.isNaN(number)) return '—';
        const inMillions = number / 1_000_000;
        const rounded = Math.round(inMillions * 10) / 10;
        return `${rounded.toLocaleString('ru-RU').replace('.', ',')} млн ₽`;
    }

    if (format === 'boolean') return value ? 'Да' : 'Нет';
    return String(value);
}
