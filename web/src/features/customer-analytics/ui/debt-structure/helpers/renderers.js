import { CONTRACTOR_TABLE_HEADERS, OBJECT_METRIC_DEFS } from './constants.js';
import { escapeHtml, formatDate, formatMoneyInMillions, formatObjectMetric } from './formatters.js';

export function renderContractorTable(container, data, onObjectClick) {
    if (!data || data.length === 0) {
        container.innerHTML = 'Нет данных';
        return;
    }

    let html = '<table class="contractor-table"><thead><tr>';
    Object.values(CONTRACTOR_TABLE_HEADERS).forEach((h) => {
        html += `<th>${h}</th>`;
    });
    html += '</tr></thead><tbody>';

    data.forEach((item) => {
        html += '<tr>';
        if (item.object) {
            const safeObject = escapeHtml(item.object);
            html += `<td><button type="button" class="object-link-btn" data-object="${safeObject}">${safeObject}</button></td>`;
        } else {
            html += '<td>—</td>';
        }
        html += `<td>${formatDate(item.contract_date)}</td>`;
        html += `<td>${formatDate(item.work_end_date)}</td>`;
        html += `<td>${item.number || '—'}</td>`;
        html += `<td>${formatMoneyInMillions(item.amount)}</td>`;
        html += `<td>${formatMoneyInMillions(item.debet_total)}</td>`;
        html += `<td>${formatMoneyInMillions(item.debet_overdue)}</td>`;
        html += '</tr>';
    });

    html += '</tbody></table>';
    container.innerHTML = html;

    const objectButtons = container.querySelectorAll('.object-link-btn');
    objectButtons.forEach((button) => {
        button.addEventListener('click', () => {
            const objectName = button.dataset.object;
            if (!objectName) return;
            onObjectClick(objectName);
        });
    });
}

export function renderObjectMetricCards(cardsGrid, metrics) {
    OBJECT_METRIC_DEFS.forEach((def) => {
        const wrapper = document.createElement('div');
        wrapper.className = 'metric-card-wrapper';

        const card = document.createElement('div');
        card.className = 'metric-card';

        const cardTitle = document.createElement('div');
        cardTitle.className = 'card-title';
        cardTitle.textContent = def.label;

        const cardValue = document.createElement('div');
        cardValue.className = 'card-value';
        cardValue.textContent = formatObjectMetric(metrics[def.key], def.format);

        card.appendChild(cardTitle);
        card.appendChild(cardValue);
        wrapper.appendChild(card);
        cardsGrid.appendChild(wrapper);
    });
}
