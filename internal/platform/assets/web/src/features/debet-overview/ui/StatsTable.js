const HUNDRED_MILLION = 100_000_000;

function formatBillionsRoundedUp(value) {
    const numericValue = Number(value) || 0;
    const roundedUpBillions = Math.ceil(numericValue / HUNDRED_MILLION) / 10;

    return roundedUpBillions.toLocaleString('ru-RU', {
        minimumFractionDigits: 1,
        maximumFractionDigits: 1
    }) + ' млрд ₽';
}

export class StatsTable {
    constructor(container) {
        this.container = container;
    }

    render(data) {
        if (!data || !data.length) {
            this.container.innerHTML = '<div class="empty">Нет данных для таблицы</div>';
            return;
        }

        const table = document.createElement('table');
        table.className = 'stats-table';

        const thead = document.createElement('thead');
        thead.innerHTML = `
            <tr>
                <th>Заказчик</th>
                <th>Цена контрактов</th>
                <th>Дебиторская задолженность</th>
                <th>Соотношение дебиторской задолженности к сумме контрактов</th>
            </tr>
        `;
        table.appendChild(thead);

        const tbody = document.createElement('tbody');
        data.forEach(row => {
            const tr = document.createElement('tr');

            const tdName = document.createElement('td');
            tdName.textContent = row.name;
            tr.appendChild(tdName);

            const tdContract = document.createElement('td');
            tdContract.textContent = formatBillionsRoundedUp(row.contractAmount);
            tr.appendChild(tdContract);

            const tdDebet = document.createElement('td');
            tdDebet.textContent = formatBillionsRoundedUp(row.debetTotal);
            tr.appendChild(tdDebet);

            const tdCoeff = document.createElement('td');
            const percent = Math.min(row.coefficient, 100);
            tdCoeff.innerHTML = `
                <div class="progress-container">
                    <div class="progress-bar">
                        <div class="progress-fill" style="width: ${percent}%;"></div>
                    </div>
                    <span class="coeff-value">${row.coefficient.toFixed(1) + '%'}</span>
                </div>
            `;
            tr.appendChild(tdCoeff);

            tbody.appendChild(tr);
        });
        table.appendChild(tbody);

        this.container.innerHTML = '';
        this.container.appendChild(table);
    }
}
